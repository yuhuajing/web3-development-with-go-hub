package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"log"
	"math"
	"math/big"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	quantity  int64
	client    *ethclient.Client
	contracts common.Address
	rpc       map[string]string
	//gaslimit         int64
	chain     string
	pks       []string
	chainId   *big.Int
	err       error
	instance  *Brc20
	groupSize int
	gasFee    float64
	logsChan  = make(chan types.Log, 0)
	target    map[string]bool
)

func init() {
	rpc = make(map[string]string)
	rpc["Eth"] = "https://eth.drpc.org"
	rpc["Sepolia"] = "https://sepolia.drpc.org"
	rpc["Bsc"] = "wss://base-rpc.publicnode.com"
	rpc["Arbitrum"] = "https://arbitrum.drpc.org"
	rpc["Base"] = "wss://base-rpc.publicnode.com"
	rpc["Polygon"] = "https://polygon.drpc.org"
	rpc["Op"] = "https://optimism.drpc.org"
	rpc["Linea"] = "https://linea.drpc.org"
	rpc["Zora"] = "https://rpc.zora.energy"
	rpc["zkSync"] = "https://zksync.drpc.org"

}

func main() {
	target = make(map[string]bool)
	//选择连
	chain = "Base"
	// NFT mint的 gas花销
	gasFee = 0.0001
	// 合约地址
	contracts = common.HexToAddress("0x00005EA00Ac477B1030CE78506496e8C2dE24bf5")
	//传递的参数

	quantity = 1
	groupSize = 10
	client = ConnBlockchain(rpc[chain])
	chainId, err = client.ChainID(context.Background())
	if err != nil {
		fmt.Println("读取ChainId错误:", err)
		return
	}

	instance, err = NewBrc20(contracts, client)
	if err != nil {
		log.Printf("构建交易Instance失败,err= %v", err)
		return
	}
	SubStakingEvent()
	//price, err := CheckTimeAndPrice(nft, quantity)
	//log.Print(price.Uint64())
	//if err != nil {
	//	fmt.Println("err= %v", err)
	//	return
	//}
	//
	//MintPublic(price)
}

func SubStakingEvent() {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contracts},
		Topics:    [][]common.Hash{{common.HexToHash("0xe90cf9cc0a552cf52ea6ff74ece0f1c8ae8cc9ad630d3181f55ac43ca076b7d6")}},
	}
	subevents, err := client.SubscribeFilterLogs(context.Background(), query, logsChan)
	if err != nil {
		fmt.Println(fmt.Errorf("Subscribe Event error: %v", err))
		log.Fatal(err)
	}
	for {
		select {
		case err := <-subevents.Err():
			fmt.Println(fmt.Errorf("Rpc 链接失败，请重启 error: %v", err))
			SubStakingEvent()
		case lplog := <-logsChan:
			SeadropFilterer, _ := NewBrc20Filterer(lplog.Address, client)
			data, _ := SeadropFilterer.ParseSeaDropMint(lplog)
			if !target[data.NftContract.Hex()] {
				target[data.NftContract.Hex()] = true
				if data.DropStageIndex.Int64() == 0 { //&& data.UnitMintPrice.String() == "0" {
					price, err := CheckTimeAndPrice(data.NftContract, quantity)
					log.Print(price.Uint64())
					if err != nil {
						fmt.Println("err= %v", err)
						return
					}
					MintPublic(data.NftContract, data.FeeRecipient, data.Payer, price) //big.NewInt(0))
				}
			}
		}
	}
}

func CheckTimeAndPrice(nft common.Address, quantity int64) (*big.Int, error) {
	config, err := instance.GetPublicDrop(nil, nft)
	if err != nil {
		return nil, fmt.Errorf("获取NFT配置信息失败,err= %v", err)
	}
	now := time.Now().Unix()
	if !(config.StartTime.Cmp(big.NewInt(now)) == -1 && config.EndTime.Cmp(big.NewInt(now)) == 1) {
		return nil, fmt.Errorf("不在当前NFT的Mint区间内,err= %v", err)
	}
	if uint16(quantity) > config.MaxTotalMintableByWallet {
		fmt.Errorf("超出个人允许mint的最大单值= %d", config.MaxTotalMintableByWallet)
		return nil, fmt.Errorf("超出个人允许mint的最大单值= %d", config.MaxTotalMintableByWallet)
	}
	return config.MintPrice, nil
}

func MintPublic(nft, fee, payer common.Address, price *big.Int) {
	pks = make([]string, 0)
	file, err := os.Open("prikey.csv")
	if err != nil {
		fmt.Println("读取私钥文件prikey.csv:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("读取私钥文件错误:", err)
		return
	}

	for _, record := range records {
		if record[0] == "" {
			break
		}
		if len(record) != 1 {
			panic("line err")
		}

		if strings.HasPrefix(record[0], "0x") {
			record[0] = record[0][2:]
		}
		pks = append(pks, record[0])
	}
	var wg sync.WaitGroup
	//result := value * math.Pow(10, 18)
	// 处理私钥钱包地址
	for i := 0; i < len(pks); i += groupSize {
		end := i + groupSize
		if end > len(pks) {
			end = len(pks)
		}

		wg.Add(1)
		go func(group []string) {
			defer wg.Done()
			for _, pk := range group {
				err := mintpublic(nft, fee, payer, quantity, price, pk)
				if err != nil {
					fmt.Printf("交易失败-私钥= %s, err: %v \n", pk[:6], err)
					continue
				}
			}
			time.Sleep(600 * time.Millisecond)
		}(pks[i:end])
	}
	wg.Wait()
}

func mintpublic(nft, fee, payer common.Address, qty int64, price *big.Int, privateKey string) error {
	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatalf("Errors in parsing key %v", err)
	}

	publicKey := pk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("write导出公钥地址失败,err= %v", err)
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return fmt.Errorf("Write获取地址nonce失败,err= %v", err)
	}
	opts, err := bind.NewKeyedTransactorWithChainID(pk, chainId)
	if err != nil {
		return fmt.Errorf("write绑定链数据失败,err= %v", err)
	}
	opts.Nonce = big.NewInt(int64(nonce))

	var pp = new(big.Int)
	pp.Mul(price, big.NewInt(quantity))
	opts.Value = pp

	suggestPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("Write获取GasPrice失败,err= %v", err)
	}
	opts.GasPrice = suggestPrice

	var limit = new(big.Int)
	gfee := gasFee * math.Pow(10, 18)
	limit.Div(big.NewInt(int64(gfee)), suggestPrice)
	opts.GasLimit = limit.Uint64()

	trans, err := instance.MintPublic(opts, nft, fee, payer, big.NewInt(qty))

	if err != nil {
		return fmt.Errorf("write发送交易失败,err= %v", err)
	} else {
		log.Printf("交易成功，hash：%s", trans.Hash().Hex())
	}
	return nil
}

func ConnBlockchain(str string) *ethclient.Client {
	client, err := ethclient.Dial(str)
	if err != nil {
		//fmt.Printf("Conn chain err = %v", err)
		log.Fatalf("Conn connect error:%s\n", err)
		//log.Fatal(err)
	}
	return client
}

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AllowListData is an auto generated low-level Go binding around an user-defined struct.
type AllowListData struct {
	MerkleRoot    [32]byte
	PublicKeyURIs []string
	AllowListURI  string
}

// MintParams is an auto generated low-level Go binding around an user-defined struct.
type MintParams struct {
	MintPrice                *big.Int
	MaxTotalMintableByWallet *big.Int
	StartTime                *big.Int
	EndTime                  *big.Int
	DropStageIndex           *big.Int
	MaxTokenSupplyForStage   *big.Int
	FeeBps                   *big.Int
	RestrictFeeRecipients    bool
}

// PublicDrop is an auto generated low-level Go binding around an user-defined struct.
type PublicDrop struct {
	MintPrice                *big.Int
	StartTime                *big.Int
	EndTime                  *big.Int
	MaxTotalMintableByWallet uint16
	FeeBps                   uint16
	RestrictFeeRecipients    bool
}

// SignedMintValidationParams is an auto generated low-level Go binding around an user-defined struct.
type SignedMintValidationParams struct {
	MinMintPrice                *big.Int
	MaxMaxTotalMintableByWallet *big.Int
	MinStartTime                *big.Int
	MaxEndTime                  *big.Int
	MaxMaxTokenSupplyForStage   *big.Int
	MinFeeBps                   uint16
	MaxFeeBps                   uint16
}

// TokenGatedDropStage is an auto generated low-level Go binding around an user-defined struct.
type TokenGatedDropStage struct {
	MintPrice                *big.Int
	MaxTotalMintableByWallet uint16
	StartTime                *big.Int
	EndTime                  *big.Int
	DropStageIndex           uint8
	MaxTokenSupplyForStage   uint32
	FeeBps                   uint16
	RestrictFeeRecipients    bool
}

// TokenGatedMintParams is an auto generated low-level Go binding around an user-defined struct.
type TokenGatedMintParams struct {
	AllowedNftToken    common.Address
	AllowedNftTokenIds []*big.Int
}

// Brc20MetaData contains all meta data concerning the Brc20 contract.
var Brc20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CreatorPayoutAddressCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateFeeRecipient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicatePayer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeRecipientCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeRecipientNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeRecipientNotPresent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"IncorrectPayment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"}],\"name\":\"InvalidFeeBps\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recoveredSigner\",\"type\":\"address\"}],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximum\",\"type\":\"uint256\"}],\"name\":\"InvalidSignedEndTime\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumOrMaximum\",\"type\":\"uint256\"}],\"name\":\"InvalidSignedFeeBps\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximum\",\"type\":\"uint256\"}],\"name\":\"InvalidSignedMaxTokenSupplyForStage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximum\",\"type\":\"uint256\"}],\"name\":\"InvalidSignedMaxTotalMintableByWallet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"}],\"name\":\"InvalidSignedMintPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"}],\"name\":\"InvalidSignedStartTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintQuantityCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allowed\",\"type\":\"uint256\"}],\"name\":\"MintQuantityExceedsMaxMintedPerWallet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"name\":\"MintQuantityExceedsMaxSupply\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint256\"}],\"name\":\"MintQuantityExceedsMaxTokenSupplyForStage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"}],\"name\":\"NotActive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"OnlyINonFungibleSeaDropToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PayerCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PayerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PayerNotPresent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureAlreadyUsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignedMintsMustRestrictFeeRecipients\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerNotPresent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenGatedDropAllowedNftTokenCannotBeDropToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenGatedDropAllowedNftTokenCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenGatedDropStageNotPresent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowedNftTokenId\",\"type\":\"uint256\"}],\"name\":\"TokenGatedNotTokenOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowedNftTokenId\",\"type\":\"uint256\"}],\"name\":\"TokenGatedTokenIdAlreadyRedeemed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousMerkleRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newMerkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"publicKeyURI\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"allowListURI\",\"type\":\"string\"}],\"name\":\"AllowListUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"AllowedFeeRecipientUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPayoutAddress\",\"type\":\"address\"}],\"name\":\"CreatorPayoutAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newDropURI\",\"type\":\"string\"}],\"name\":\"DropURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"PayerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structPublicDrop\",\"name\":\"publicDrop\",\"type\":\"tuple\"}],\"name\":\"PublicDropUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantityMinted\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unitMintPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dropStageIndex\",\"type\":\"uint256\"}],\"name\":\"SeaDropMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"minMintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint24\",\"name\":\"maxMaxTotalMintableByWallet\",\"type\":\"uint24\"},{\"internalType\":\"uint40\",\"name\":\"minStartTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxEndTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxMaxTokenSupplyForStage\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"minFeeBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxFeeBps\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structSignedMintValidationParams\",\"name\":\"signedMintValidationParams\",\"type\":\"tuple\"}],\"name\":\"SignedMintValidationParamsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint8\",\"name\":\"dropStageIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structTokenGatedDropStage\",\"name\":\"dropStage\",\"type\":\"tuple\"}],\"name\":\"TokenGatedDropStageUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"}],\"name\":\"getAllowListMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"}],\"name\":\"getAllowedFeeRecipients\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowedNftTokenId\",\"type\":\"uint256\"}],\"name\":\"getAllowedNftTokenIdIsRedeemed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"}],\"name\":\"getCreatorPayoutAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"getFeeRecipientIsAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"}],\"name\":\"getPayerIsAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"}],\"name\":\"getPayers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"}],\"name\":\"getPublicDrop\",\"outputs\":[{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structPublicDrop\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"getSignedMintValidationParams\",\"outputs\":[{\"components\":[{\"internalType\":\"uint80\",\"name\":\"minMintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint24\",\"name\":\"maxMaxTotalMintableByWallet\",\"type\":\"uint24\"},{\"internalType\":\"uint40\",\"name\":\"minStartTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxEndTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxMaxTokenSupplyForStage\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"minFeeBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxFeeBps\",\"type\":\"uint16\"}],\"internalType\":\"structSignedMintValidationParams\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"}],\"name\":\"getSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"}],\"name\":\"getTokenGatedAllowedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"}],\"name\":\"getTokenGatedDrop\",\"outputs\":[{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint8\",\"name\":\"dropStageIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structTokenGatedDropStage\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minterIfNotPayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"mintPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dropStageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structMintParams\",\"name\":\"mintParams\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"mintAllowList\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minterIfNotPayer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"allowedNftTokenIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structTokenGatedMintParams\",\"name\":\"mintParams\",\"type\":\"tuple\"}],\"name\":\"mintAllowedTokenHolder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minterIfNotPayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"mintPublic\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minterIfNotPayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"mintPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dropStageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structMintParams\",\"name\":\"mintParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"mintSigned\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string[]\",\"name\":\"publicKeyURIs\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"allowListURI\",\"type\":\"string\"}],\"internalType\":\"structAllowListData\",\"name\":\"allowListData\",\"type\":\"tuple\"}],\"name\":\"updateAllowList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updateAllowedFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_payoutAddress\",\"type\":\"address\"}],\"name\":\"updateCreatorPayoutAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dropURI\",\"type\":\"string\"}],\"name\":\"updateDropURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updatePayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structPublicDrop\",\"name\":\"publicDrop\",\"type\":\"tuple\"}],\"name\":\"updatePublicDrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"minMintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint24\",\"name\":\"maxMaxTotalMintableByWallet\",\"type\":\"uint24\"},{\"internalType\":\"uint40\",\"name\":\"minStartTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxEndTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxMaxTokenSupplyForStage\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"minFeeBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxFeeBps\",\"type\":\"uint16\"}],\"internalType\":\"structSignedMintValidationParams\",\"name\":\"signedMintValidationParams\",\"type\":\"tuple\"}],\"name\":\"updateSignedMintValidationParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint8\",\"name\":\"dropStageIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structTokenGatedDropStage\",\"name\":\"dropStage\",\"type\":\"tuple\"}],\"name\":\"updateTokenGatedDrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Brc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use Brc20MetaData.ABI instead.
var Brc20ABI = Brc20MetaData.ABI

// Brc20 is an auto generated Go binding around an Ethereum contract.
type Brc20 struct {
	Brc20Caller     // Read-only binding to the contract
	Brc20Transactor // Write-only binding to the contract
	Brc20Filterer   // Log filterer for contract events
}

// Brc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type Brc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Brc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Brc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Brc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Brc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Brc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Brc20Session struct {
	Contract     *Brc20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Brc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Brc20CallerSession struct {
	Contract *Brc20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Brc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Brc20TransactorSession struct {
	Contract     *Brc20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Brc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type Brc20Raw struct {
	Contract *Brc20 // Generic contract binding to access the raw methods on
}

// Brc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Brc20CallerRaw struct {
	Contract *Brc20Caller // Generic read-only contract binding to access the raw methods on
}

// Brc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Brc20TransactorRaw struct {
	Contract *Brc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBrc20 creates a new instance of Brc20, bound to a specific deployed contract.
func NewBrc20(address common.Address, backend bind.ContractBackend) (*Brc20, error) {
	contract, err := bindBrc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Brc20{Brc20Caller: Brc20Caller{contract: contract}, Brc20Transactor: Brc20Transactor{contract: contract}, Brc20Filterer: Brc20Filterer{contract: contract}}, nil
}

// NewBrc20Caller creates a new read-only instance of Brc20, bound to a specific deployed contract.
func NewBrc20Caller(address common.Address, caller bind.ContractCaller) (*Brc20Caller, error) {
	contract, err := bindBrc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Brc20Caller{contract: contract}, nil
}

// NewBrc20Transactor creates a new write-only instance of Brc20, bound to a specific deployed contract.
func NewBrc20Transactor(address common.Address, transactor bind.ContractTransactor) (*Brc20Transactor, error) {
	contract, err := bindBrc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Brc20Transactor{contract: contract}, nil
}

// NewBrc20Filterer creates a new log filterer instance of Brc20, bound to a specific deployed contract.
func NewBrc20Filterer(address common.Address, filterer bind.ContractFilterer) (*Brc20Filterer, error) {
	contract, err := bindBrc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Brc20Filterer{contract: contract}, nil
}

// bindBrc20 binds a generic wrapper to an already deployed contract.
func bindBrc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Brc20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Brc20 *Brc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Brc20.Contract.Brc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Brc20 *Brc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Brc20.Contract.Brc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Brc20 *Brc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Brc20.Contract.Brc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Brc20 *Brc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Brc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Brc20 *Brc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Brc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Brc20 *Brc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Brc20.Contract.contract.Transact(opts, method, params...)
}

// GetAllowListMerkleRoot is a free data retrieval call binding the contract method 0x32bf11f5.
//
// Solidity: function getAllowListMerkleRoot(address nftContract) view returns(bytes32)
func (_Brc20 *Brc20Caller) GetAllowListMerkleRoot(opts *bind.CallOpts, nftContract common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Brc20.contract.Call(opts, &out, "getAllowListMerkleRoot", nftContract)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAllowListMerkleRoot is a free data retrieval call binding the contract method 0x32bf11f5.
//
// Solidity: function getAllowListMerkleRoot(address nftContract) view returns(bytes32)
func (_Brc20 *Brc20Session) GetAllowListMerkleRoot(nftContract common.Address) ([32]byte, error) {
	return _Brc20.Contract.GetAllowListMerkleRoot(&_Brc20.CallOpts, nftContract)
}

// GetAllowListMerkleRoot is a free data retrieval call binding the contract method 0x32bf11f5.
//
// Solidity: function getAllowListMerkleRoot(address nftContract) view returns(bytes32)
func (_Brc20 *Brc20CallerSession) GetAllowListMerkleRoot(nftContract common.Address) ([32]byte, error) {
	return _Brc20.Contract.GetAllowListMerkleRoot(&_Brc20.CallOpts, nftContract)
}

// GetAllowedFeeRecipients is a free data retrieval call binding the contract method 0x68632274.
//
// Solidity: function getAllowedFeeRecipients(address nftContract) view returns(address[])
func (_Brc20 *Brc20Caller) GetAllowedFeeRecipients(opts *bind.CallOpts, nftContract common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Brc20.contract.Call(opts, &out, "getAllowedFeeRecipients", nftContract)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllowedFeeRecipients is a free data retrieval call binding the contract method 0x68632274.
//
// Solidity: function getAllowedFeeRecipients(address nftContract) view returns(address[])
func (_Brc20 *Brc20Session) GetAllowedFeeRecipients(nftContract common.Address) ([]common.Address, error) {
	return _Brc20.Contract.GetAllowedFeeRecipients(&_Brc20.CallOpts, nftContract)
}

// GetAllowedFeeRecipients is a free data retrieval call binding the contract method 0x68632274.
//
// Solidity: function getAllowedFeeRecipients(address nftContract) view returns(address[])
func (_Brc20 *Brc20CallerSession) GetAllowedFeeRecipients(nftContract common.Address) ([]common.Address, error) {
	return _Brc20.Contract.GetAllowedFeeRecipients(&_Brc20.CallOpts, nftContract)
}

// GetAllowedNftTokenIdIsRedeemed is a free data retrieval call binding the contract method 0x88aa3d37.
//
// Solidity: function getAllowedNftTokenIdIsRedeemed(address nftContract, address allowedNftToken, uint256 allowedNftTokenId) view returns(bool)
func (_Brc20 *Brc20Caller) GetAllowedNftTokenIdIsRedeemed(opts *bind.CallOpts, nftContract common.Address, allowedNftToken common.Address, allowedNftTokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Brc20.contract.Call(opts, &out, "getAllowedNftTokenIdIsRedeemed", nftContract, allowedNftToken, allowedNftTokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetAllowedNftTokenIdIsRedeemed is a free data retrieval call binding the contract method 0x88aa3d37.
//
// Solidity: function getAllowedNftTokenIdIsRedeemed(address nftContract, address allowedNftToken, uint256 allowedNftTokenId) view returns(bool)
func (_Brc20 *Brc20Session) GetAllowedNftTokenIdIsRedeemed(nftContract common.Address, allowedNftToken common.Address, allowedNftTokenId *big.Int) (bool, error) {
	return _Brc20.Contract.GetAllowedNftTokenIdIsRedeemed(&_Brc20.CallOpts, nftContract, allowedNftToken, allowedNftTokenId)
}

// GetAllowedNftTokenIdIsRedeemed is a free data retrieval call binding the contract method 0x88aa3d37.
//
// Solidity: function getAllowedNftTokenIdIsRedeemed(address nftContract, address allowedNftToken, uint256 allowedNftTokenId) view returns(bool)
func (_Brc20 *Brc20CallerSession) GetAllowedNftTokenIdIsRedeemed(nftContract common.Address, allowedNftToken common.Address, allowedNftTokenId *big.Int) (bool, error) {
	return _Brc20.Contract.GetAllowedNftTokenIdIsRedeemed(&_Brc20.CallOpts, nftContract, allowedNftToken, allowedNftTokenId)
}

// GetCreatorPayoutAddress is a free data retrieval call binding the contract method 0x5cb3c4d3.
//
// Solidity: function getCreatorPayoutAddress(address nftContract) view returns(address)
func (_Brc20 *Brc20Caller) GetCreatorPayoutAddress(opts *bind.CallOpts, nftContract common.Address) (common.Address, error) {
	var out []interface{}
	err := _Brc20.contract.Call(opts, &out, "getCreatorPayoutAddress", nftContract)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCreatorPayoutAddress is a free data retrieval call binding the contract method 0x5cb3c4d3.
//
// Solidity: function getCreatorPayoutAddress(address nftContract) view returns(address)
func (_Brc20 *Brc20Session) GetCreatorPayoutAddress(nftContract common.Address) (common.Address, error) {
	return _Brc20.Contract.GetCreatorPayoutAddress(&_Brc20.CallOpts, nftContract)
}

// GetCreatorPayoutAddress is a free data retrieval call binding the contract method 0x5cb3c4d3.
//
// Solidity: function getCreatorPayoutAddress(address nftContract) view returns(address)
func (_Brc20 *Brc20CallerSession) GetCreatorPayoutAddress(nftContract common.Address) (common.Address, error) {
	return _Brc20.Contract.GetCreatorPayoutAddress(&_Brc20.CallOpts, nftContract)
}

// GetFeeRecipientIsAllowed is a free data retrieval call binding the contract method 0x322e75d1.
//
// Solidity: function getFeeRecipientIsAllowed(address nftContract, address feeRecipient) view returns(bool)
func (_Brc20 *Brc20Caller) GetFeeRecipientIsAllowed(opts *bind.CallOpts, nftContract common.Address, feeRecipient common.Address) (bool, error) {
	var out []interface{}
	err := _Brc20.contract.Call(opts, &out, "getFeeRecipientIsAllowed", nftContract, feeRecipient)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetFeeRecipientIsAllowed is a free data retrieval call binding the contract method 0x322e75d1.
//
// Solidity: function getFeeRecipientIsAllowed(address nftContract, address feeRecipient) view returns(bool)
func (_Brc20 *Brc20Session) GetFeeRecipientIsAllowed(nftContract common.Address, feeRecipient common.Address) (bool, error) {
	return _Brc20.Contract.GetFeeRecipientIsAllowed(&_Brc20.CallOpts, nftContract, feeRecipient)
}

// GetFeeRecipientIsAllowed is a free data retrieval call binding the contract method 0x322e75d1.
//
// Solidity: function getFeeRecipientIsAllowed(address nftContract, address feeRecipient) view returns(bool)
func (_Brc20 *Brc20CallerSession) GetFeeRecipientIsAllowed(nftContract common.Address, feeRecipient common.Address) (bool, error) {
	return _Brc20.Contract.GetFeeRecipientIsAllowed(&_Brc20.CallOpts, nftContract, feeRecipient)
}

// GetPayerIsAllowed is a free data retrieval call binding the contract method 0xe583141d.
//
// Solidity: function getPayerIsAllowed(address nftContract, address payer) view returns(bool)
func (_Brc20 *Brc20Caller) GetPayerIsAllowed(opts *bind.CallOpts, nftContract common.Address, payer common.Address) (bool, error) {
	var out []interface{}
	err := _Brc20.contract.Call(opts, &out, "getPayerIsAllowed", nftContract, payer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetPayerIsAllowed is a free data retrieval call binding the contract method 0xe583141d.
//
// Solidity: function getPayerIsAllowed(address nftContract, address payer) view returns(bool)
func (_Brc20 *Brc20Session) GetPayerIsAllowed(nftContract common.Address, payer common.Address) (bool, error) {
	return _Brc20.Contract.GetPayerIsAllowed(&_Brc20.CallOpts, nftContract, payer)
}

// GetPayerIsAllowed is a free data retrieval call binding the contract method 0xe583141d.
//
// Solidity: function getPayerIsAllowed(address nftContract, address payer) view returns(bool)
func (_Brc20 *Brc20CallerSession) GetPayerIsAllowed(nftContract common.Address, payer common.Address) (bool, error) {
	return _Brc20.Contract.GetPayerIsAllowed(&_Brc20.CallOpts, nftContract, payer)
}

// GetPayers is a free data retrieval call binding the contract method 0x7c35b982.
//
// Solidity: function getPayers(address nftContract) view returns(address[])
func (_Brc20 *Brc20Caller) GetPayers(opts *bind.CallOpts, nftContract common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Brc20.contract.Call(opts, &out, "getPayers", nftContract)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPayers is a free data retrieval call binding the contract method 0x7c35b982.
//
// Solidity: function getPayers(address nftContract) view returns(address[])
func (_Brc20 *Brc20Session) GetPayers(nftContract common.Address) ([]common.Address, error) {
	return _Brc20.Contract.GetPayers(&_Brc20.CallOpts, nftContract)
}

// GetPayers is a free data retrieval call binding the contract method 0x7c35b982.
//
// Solidity: function getPayers(address nftContract) view returns(address[])
func (_Brc20 *Brc20CallerSession) GetPayers(nftContract common.Address) ([]common.Address, error) {
	return _Brc20.Contract.GetPayers(&_Brc20.CallOpts, nftContract)
}

// GetPublicDrop is a free data retrieval call binding the contract method 0xbc6a629c.
//
// Solidity: function getPublicDrop(address nftContract) view returns((uint80,uint48,uint48,uint16,uint16,bool))
func (_Brc20 *Brc20Caller) GetPublicDrop(opts *bind.CallOpts, nftContract common.Address) (PublicDrop, error) {
	var out []interface{}
	err := _Brc20.contract.Call(opts, &out, "getPublicDrop", nftContract)

	if err != nil {
		return *new(PublicDrop), err
	}

	out0 := *abi.ConvertType(out[0], new(PublicDrop)).(*PublicDrop)

	return out0, err

}

// GetPublicDrop is a free data retrieval call binding the contract method 0xbc6a629c.
//
// Solidity: function getPublicDrop(address nftContract) view returns((uint80,uint48,uint48,uint16,uint16,bool))
func (_Brc20 *Brc20Session) GetPublicDrop(nftContract common.Address) (PublicDrop, error) {
	return _Brc20.Contract.GetPublicDrop(&_Brc20.CallOpts, nftContract)
}

// GetPublicDrop is a free data retrieval call binding the contract method 0xbc6a629c.
//
// Solidity: function getPublicDrop(address nftContract) view returns((uint80,uint48,uint48,uint16,uint16,bool))
func (_Brc20 *Brc20CallerSession) GetPublicDrop(nftContract common.Address) (PublicDrop, error) {
	return _Brc20.Contract.GetPublicDrop(&_Brc20.CallOpts, nftContract)
}

// GetSignedMintValidationParams is a free data retrieval call binding the contract method 0x81bf9af3.
//
// Solidity: function getSignedMintValidationParams(address nftContract, address signer) view returns((uint80,uint24,uint40,uint40,uint40,uint16,uint16))
func (_Brc20 *Brc20Caller) GetSignedMintValidationParams(opts *bind.CallOpts, nftContract common.Address, signer common.Address) (SignedMintValidationParams, error) {
	var out []interface{}
	err := _Brc20.contract.Call(opts, &out, "getSignedMintValidationParams", nftContract, signer)

	if err != nil {
		return *new(SignedMintValidationParams), err
	}

	out0 := *abi.ConvertType(out[0], new(SignedMintValidationParams)).(*SignedMintValidationParams)

	return out0, err

}

// GetSignedMintValidationParams is a free data retrieval call binding the contract method 0x81bf9af3.
//
// Solidity: function getSignedMintValidationParams(address nftContract, address signer) view returns((uint80,uint24,uint40,uint40,uint40,uint16,uint16))
func (_Brc20 *Brc20Session) GetSignedMintValidationParams(nftContract common.Address, signer common.Address) (SignedMintValidationParams, error) {
	return _Brc20.Contract.GetSignedMintValidationParams(&_Brc20.CallOpts, nftContract, signer)
}

// GetSignedMintValidationParams is a free data retrieval call binding the contract method 0x81bf9af3.
//
// Solidity: function getSignedMintValidationParams(address nftContract, address signer) view returns((uint80,uint24,uint40,uint40,uint40,uint16,uint16))
func (_Brc20 *Brc20CallerSession) GetSignedMintValidationParams(nftContract common.Address, signer common.Address) (SignedMintValidationParams, error) {
	return _Brc20.Contract.GetSignedMintValidationParams(&_Brc20.CallOpts, nftContract, signer)
}

// GetSigners is a free data retrieval call binding the contract method 0x7e3ba6af.
//
// Solidity: function getSigners(address nftContract) view returns(address[])
func (_Brc20 *Brc20Caller) GetSigners(opts *bind.CallOpts, nftContract common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Brc20.contract.Call(opts, &out, "getSigners", nftContract)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSigners is a free data retrieval call binding the contract method 0x7e3ba6af.
//
// Solidity: function getSigners(address nftContract) view returns(address[])
func (_Brc20 *Brc20Session) GetSigners(nftContract common.Address) ([]common.Address, error) {
	return _Brc20.Contract.GetSigners(&_Brc20.CallOpts, nftContract)
}

// GetSigners is a free data retrieval call binding the contract method 0x7e3ba6af.
//
// Solidity: function getSigners(address nftContract) view returns(address[])
func (_Brc20 *Brc20CallerSession) GetSigners(nftContract common.Address) ([]common.Address, error) {
	return _Brc20.Contract.GetSigners(&_Brc20.CallOpts, nftContract)
}

// GetTokenGatedAllowedTokens is a free data retrieval call binding the contract method 0x2db526eb.
//
// Solidity: function getTokenGatedAllowedTokens(address nftContract) view returns(address[])
func (_Brc20 *Brc20Caller) GetTokenGatedAllowedTokens(opts *bind.CallOpts, nftContract common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Brc20.contract.Call(opts, &out, "getTokenGatedAllowedTokens", nftContract)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokenGatedAllowedTokens is a free data retrieval call binding the contract method 0x2db526eb.
//
// Solidity: function getTokenGatedAllowedTokens(address nftContract) view returns(address[])
func (_Brc20 *Brc20Session) GetTokenGatedAllowedTokens(nftContract common.Address) ([]common.Address, error) {
	return _Brc20.Contract.GetTokenGatedAllowedTokens(&_Brc20.CallOpts, nftContract)
}

// GetTokenGatedAllowedTokens is a free data retrieval call binding the contract method 0x2db526eb.
//
// Solidity: function getTokenGatedAllowedTokens(address nftContract) view returns(address[])
func (_Brc20 *Brc20CallerSession) GetTokenGatedAllowedTokens(nftContract common.Address) ([]common.Address, error) {
	return _Brc20.Contract.GetTokenGatedAllowedTokens(&_Brc20.CallOpts, nftContract)
}

// GetTokenGatedDrop is a free data retrieval call binding the contract method 0x0b0e8a6e.
//
// Solidity: function getTokenGatedDrop(address nftContract, address allowedNftToken) view returns((uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool))
func (_Brc20 *Brc20Caller) GetTokenGatedDrop(opts *bind.CallOpts, nftContract common.Address, allowedNftToken common.Address) (TokenGatedDropStage, error) {
	var out []interface{}
	err := _Brc20.contract.Call(opts, &out, "getTokenGatedDrop", nftContract, allowedNftToken)

	if err != nil {
		return *new(TokenGatedDropStage), err
	}

	out0 := *abi.ConvertType(out[0], new(TokenGatedDropStage)).(*TokenGatedDropStage)

	return out0, err

}

// GetTokenGatedDrop is a free data retrieval call binding the contract method 0x0b0e8a6e.
//
// Solidity: function getTokenGatedDrop(address nftContract, address allowedNftToken) view returns((uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool))
func (_Brc20 *Brc20Session) GetTokenGatedDrop(nftContract common.Address, allowedNftToken common.Address) (TokenGatedDropStage, error) {
	return _Brc20.Contract.GetTokenGatedDrop(&_Brc20.CallOpts, nftContract, allowedNftToken)
}

// GetTokenGatedDrop is a free data retrieval call binding the contract method 0x0b0e8a6e.
//
// Solidity: function getTokenGatedDrop(address nftContract, address allowedNftToken) view returns((uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool))
func (_Brc20 *Brc20CallerSession) GetTokenGatedDrop(nftContract common.Address, allowedNftToken common.Address) (TokenGatedDropStage, error) {
	return _Brc20.Contract.GetTokenGatedDrop(&_Brc20.CallOpts, nftContract, allowedNftToken)
}

// MintAllowList is a paid mutator transaction binding the contract method 0x4300a4e6.
//
// Solidity: function mintAllowList(address nftContract, address feeRecipient, address minterIfNotPayer, uint256 quantity, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool) mintParams, bytes32[] proof) payable returns()
func (_Brc20 *Brc20Transactor) MintAllowList(opts *bind.TransactOpts, nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, quantity *big.Int, mintParams MintParams, proof [][32]byte) (*types.Transaction, error) {
	return _Brc20.contract.Transact(opts, "mintAllowList", nftContract, feeRecipient, minterIfNotPayer, quantity, mintParams, proof)
}

// MintAllowList is a paid mutator transaction binding the contract method 0x4300a4e6.
//
// Solidity: function mintAllowList(address nftContract, address feeRecipient, address minterIfNotPayer, uint256 quantity, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool) mintParams, bytes32[] proof) payable returns()
func (_Brc20 *Brc20Session) MintAllowList(nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, quantity *big.Int, mintParams MintParams, proof [][32]byte) (*types.Transaction, error) {
	return _Brc20.Contract.MintAllowList(&_Brc20.TransactOpts, nftContract, feeRecipient, minterIfNotPayer, quantity, mintParams, proof)
}

// MintAllowList is a paid mutator transaction binding the contract method 0x4300a4e6.
//
// Solidity: function mintAllowList(address nftContract, address feeRecipient, address minterIfNotPayer, uint256 quantity, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool) mintParams, bytes32[] proof) payable returns()
func (_Brc20 *Brc20TransactorSession) MintAllowList(nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, quantity *big.Int, mintParams MintParams, proof [][32]byte) (*types.Transaction, error) {
	return _Brc20.Contract.MintAllowList(&_Brc20.TransactOpts, nftContract, feeRecipient, minterIfNotPayer, quantity, mintParams, proof)
}

// MintAllowedTokenHolder is a paid mutator transaction binding the contract method 0x99eb900f.
//
// Solidity: function mintAllowedTokenHolder(address nftContract, address feeRecipient, address minterIfNotPayer, (address,uint256[]) mintParams) payable returns()
func (_Brc20 *Brc20Transactor) MintAllowedTokenHolder(opts *bind.TransactOpts, nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, mintParams TokenGatedMintParams) (*types.Transaction, error) {
	return _Brc20.contract.Transact(opts, "mintAllowedTokenHolder", nftContract, feeRecipient, minterIfNotPayer, mintParams)
}

// MintAllowedTokenHolder is a paid mutator transaction binding the contract method 0x99eb900f.
//
// Solidity: function mintAllowedTokenHolder(address nftContract, address feeRecipient, address minterIfNotPayer, (address,uint256[]) mintParams) payable returns()
func (_Brc20 *Brc20Session) MintAllowedTokenHolder(nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, mintParams TokenGatedMintParams) (*types.Transaction, error) {
	return _Brc20.Contract.MintAllowedTokenHolder(&_Brc20.TransactOpts, nftContract, feeRecipient, minterIfNotPayer, mintParams)
}

// MintAllowedTokenHolder is a paid mutator transaction binding the contract method 0x99eb900f.
//
// Solidity: function mintAllowedTokenHolder(address nftContract, address feeRecipient, address minterIfNotPayer, (address,uint256[]) mintParams) payable returns()
func (_Brc20 *Brc20TransactorSession) MintAllowedTokenHolder(nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, mintParams TokenGatedMintParams) (*types.Transaction, error) {
	return _Brc20.Contract.MintAllowedTokenHolder(&_Brc20.TransactOpts, nftContract, feeRecipient, minterIfNotPayer, mintParams)
}

// MintPublic is a paid mutator transaction binding the contract method 0x161ac21f.
//
// Solidity: function mintPublic(address nftContract, address feeRecipient, address minterIfNotPayer, uint256 quantity) payable returns()
func (_Brc20 *Brc20Transactor) MintPublic(opts *bind.TransactOpts, nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Brc20.contract.Transact(opts, "mintPublic", nftContract, feeRecipient, minterIfNotPayer, quantity)
}

// MintPublic is a paid mutator transaction binding the contract method 0x161ac21f.
//
// Solidity: function mintPublic(address nftContract, address feeRecipient, address minterIfNotPayer, uint256 quantity) payable returns()
func (_Brc20 *Brc20Session) MintPublic(nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Brc20.Contract.MintPublic(&_Brc20.TransactOpts, nftContract, feeRecipient, minterIfNotPayer, quantity)
}

// MintPublic is a paid mutator transaction binding the contract method 0x161ac21f.
//
// Solidity: function mintPublic(address nftContract, address feeRecipient, address minterIfNotPayer, uint256 quantity) payable returns()
func (_Brc20 *Brc20TransactorSession) MintPublic(nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Brc20.Contract.MintPublic(&_Brc20.TransactOpts, nftContract, feeRecipient, minterIfNotPayer, quantity)
}

// MintSigned is a paid mutator transaction binding the contract method 0x4b61cd6f.
//
// Solidity: function mintSigned(address nftContract, address feeRecipient, address minterIfNotPayer, uint256 quantity, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool) mintParams, uint256 salt, bytes signature) payable returns()
func (_Brc20 *Brc20Transactor) MintSigned(opts *bind.TransactOpts, nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, quantity *big.Int, mintParams MintParams, salt *big.Int, signature []byte) (*types.Transaction, error) {
	return _Brc20.contract.Transact(opts, "mintSigned", nftContract, feeRecipient, minterIfNotPayer, quantity, mintParams, salt, signature)
}

// MintSigned is a paid mutator transaction binding the contract method 0x4b61cd6f.
//
// Solidity: function mintSigned(address nftContract, address feeRecipient, address minterIfNotPayer, uint256 quantity, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool) mintParams, uint256 salt, bytes signature) payable returns()
func (_Brc20 *Brc20Session) MintSigned(nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, quantity *big.Int, mintParams MintParams, salt *big.Int, signature []byte) (*types.Transaction, error) {
	return _Brc20.Contract.MintSigned(&_Brc20.TransactOpts, nftContract, feeRecipient, minterIfNotPayer, quantity, mintParams, salt, signature)
}

// MintSigned is a paid mutator transaction binding the contract method 0x4b61cd6f.
//
// Solidity: function mintSigned(address nftContract, address feeRecipient, address minterIfNotPayer, uint256 quantity, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool) mintParams, uint256 salt, bytes signature) payable returns()
func (_Brc20 *Brc20TransactorSession) MintSigned(nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, quantity *big.Int, mintParams MintParams, salt *big.Int, signature []byte) (*types.Transaction, error) {
	return _Brc20.Contract.MintSigned(&_Brc20.TransactOpts, nftContract, feeRecipient, minterIfNotPayer, quantity, mintParams, salt, signature)
}

// UpdateAllowList is a paid mutator transaction binding the contract method 0xebb4a55f.
//
// Solidity: function updateAllowList((bytes32,string[],string) allowListData) returns()
func (_Brc20 *Brc20Transactor) UpdateAllowList(opts *bind.TransactOpts, allowListData AllowListData) (*types.Transaction, error) {
	return _Brc20.contract.Transact(opts, "updateAllowList", allowListData)
}

// UpdateAllowList is a paid mutator transaction binding the contract method 0xebb4a55f.
//
// Solidity: function updateAllowList((bytes32,string[],string) allowListData) returns()
func (_Brc20 *Brc20Session) UpdateAllowList(allowListData AllowListData) (*types.Transaction, error) {
	return _Brc20.Contract.UpdateAllowList(&_Brc20.TransactOpts, allowListData)
}

// UpdateAllowList is a paid mutator transaction binding the contract method 0xebb4a55f.
//
// Solidity: function updateAllowList((bytes32,string[],string) allowListData) returns()
func (_Brc20 *Brc20TransactorSession) UpdateAllowList(allowListData AllowListData) (*types.Transaction, error) {
	return _Brc20.Contract.UpdateAllowList(&_Brc20.TransactOpts, allowListData)
}

// UpdateAllowedFeeRecipient is a paid mutator transaction binding the contract method 0x8e7d1e43.
//
// Solidity: function updateAllowedFeeRecipient(address feeRecipient, bool allowed) returns()
func (_Brc20 *Brc20Transactor) UpdateAllowedFeeRecipient(opts *bind.TransactOpts, feeRecipient common.Address, allowed bool) (*types.Transaction, error) {
	return _Brc20.contract.Transact(opts, "updateAllowedFeeRecipient", feeRecipient, allowed)
}

// UpdateAllowedFeeRecipient is a paid mutator transaction binding the contract method 0x8e7d1e43.
//
// Solidity: function updateAllowedFeeRecipient(address feeRecipient, bool allowed) returns()
func (_Brc20 *Brc20Session) UpdateAllowedFeeRecipient(feeRecipient common.Address, allowed bool) (*types.Transaction, error) {
	return _Brc20.Contract.UpdateAllowedFeeRecipient(&_Brc20.TransactOpts, feeRecipient, allowed)
}

// UpdateAllowedFeeRecipient is a paid mutator transaction binding the contract method 0x8e7d1e43.
//
// Solidity: function updateAllowedFeeRecipient(address feeRecipient, bool allowed) returns()
func (_Brc20 *Brc20TransactorSession) UpdateAllowedFeeRecipient(feeRecipient common.Address, allowed bool) (*types.Transaction, error) {
	return _Brc20.Contract.UpdateAllowedFeeRecipient(&_Brc20.TransactOpts, feeRecipient, allowed)
}

// UpdateCreatorPayoutAddress is a paid mutator transaction binding the contract method 0x12738db8.
//
// Solidity: function updateCreatorPayoutAddress(address _payoutAddress) returns()
func (_Brc20 *Brc20Transactor) UpdateCreatorPayoutAddress(opts *bind.TransactOpts, _payoutAddress common.Address) (*types.Transaction, error) {
	return _Brc20.contract.Transact(opts, "updateCreatorPayoutAddress", _payoutAddress)
}

// UpdateCreatorPayoutAddress is a paid mutator transaction binding the contract method 0x12738db8.
//
// Solidity: function updateCreatorPayoutAddress(address _payoutAddress) returns()
func (_Brc20 *Brc20Session) UpdateCreatorPayoutAddress(_payoutAddress common.Address) (*types.Transaction, error) {
	return _Brc20.Contract.UpdateCreatorPayoutAddress(&_Brc20.TransactOpts, _payoutAddress)
}

// UpdateCreatorPayoutAddress is a paid mutator transaction binding the contract method 0x12738db8.
//
// Solidity: function updateCreatorPayoutAddress(address _payoutAddress) returns()
func (_Brc20 *Brc20TransactorSession) UpdateCreatorPayoutAddress(_payoutAddress common.Address) (*types.Transaction, error) {
	return _Brc20.Contract.UpdateCreatorPayoutAddress(&_Brc20.TransactOpts, _payoutAddress)
}

// UpdateDropURI is a paid mutator transaction binding the contract method 0xb957d0cb.
//
// Solidity: function updateDropURI(string dropURI) returns()
func (_Brc20 *Brc20Transactor) UpdateDropURI(opts *bind.TransactOpts, dropURI string) (*types.Transaction, error) {
	return _Brc20.contract.Transact(opts, "updateDropURI", dropURI)
}

// UpdateDropURI is a paid mutator transaction binding the contract method 0xb957d0cb.
//
// Solidity: function updateDropURI(string dropURI) returns()
func (_Brc20 *Brc20Session) UpdateDropURI(dropURI string) (*types.Transaction, error) {
	return _Brc20.Contract.UpdateDropURI(&_Brc20.TransactOpts, dropURI)
}

// UpdateDropURI is a paid mutator transaction binding the contract method 0xb957d0cb.
//
// Solidity: function updateDropURI(string dropURI) returns()
func (_Brc20 *Brc20TransactorSession) UpdateDropURI(dropURI string) (*types.Transaction, error) {
	return _Brc20.Contract.UpdateDropURI(&_Brc20.TransactOpts, dropURI)
}

// UpdatePayer is a paid mutator transaction binding the contract method 0x7f2a5cca.
//
// Solidity: function updatePayer(address payer, bool allowed) returns()
func (_Brc20 *Brc20Transactor) UpdatePayer(opts *bind.TransactOpts, payer common.Address, allowed bool) (*types.Transaction, error) {
	return _Brc20.contract.Transact(opts, "updatePayer", payer, allowed)
}

// UpdatePayer is a paid mutator transaction binding the contract method 0x7f2a5cca.
//
// Solidity: function updatePayer(address payer, bool allowed) returns()
func (_Brc20 *Brc20Session) UpdatePayer(payer common.Address, allowed bool) (*types.Transaction, error) {
	return _Brc20.Contract.UpdatePayer(&_Brc20.TransactOpts, payer, allowed)
}

// UpdatePayer is a paid mutator transaction binding the contract method 0x7f2a5cca.
//
// Solidity: function updatePayer(address payer, bool allowed) returns()
func (_Brc20 *Brc20TransactorSession) UpdatePayer(payer common.Address, allowed bool) (*types.Transaction, error) {
	return _Brc20.Contract.UpdatePayer(&_Brc20.TransactOpts, payer, allowed)
}

// UpdatePublicDrop is a paid mutator transaction binding the contract method 0x01308e65.
//
// Solidity: function updatePublicDrop((uint80,uint48,uint48,uint16,uint16,bool) publicDrop) returns()
func (_Brc20 *Brc20Transactor) UpdatePublicDrop(opts *bind.TransactOpts, publicDrop PublicDrop) (*types.Transaction, error) {
	return _Brc20.contract.Transact(opts, "updatePublicDrop", publicDrop)
}

// UpdatePublicDrop is a paid mutator transaction binding the contract method 0x01308e65.
//
// Solidity: function updatePublicDrop((uint80,uint48,uint48,uint16,uint16,bool) publicDrop) returns()
func (_Brc20 *Brc20Session) UpdatePublicDrop(publicDrop PublicDrop) (*types.Transaction, error) {
	return _Brc20.Contract.UpdatePublicDrop(&_Brc20.TransactOpts, publicDrop)
}

// UpdatePublicDrop is a paid mutator transaction binding the contract method 0x01308e65.
//
// Solidity: function updatePublicDrop((uint80,uint48,uint48,uint16,uint16,bool) publicDrop) returns()
func (_Brc20 *Brc20TransactorSession) UpdatePublicDrop(publicDrop PublicDrop) (*types.Transaction, error) {
	return _Brc20.Contract.UpdatePublicDrop(&_Brc20.TransactOpts, publicDrop)
}

// UpdateSignedMintValidationParams is a paid mutator transaction binding the contract method 0x4d380178.
//
// Solidity: function updateSignedMintValidationParams(address signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams) returns()
func (_Brc20 *Brc20Transactor) UpdateSignedMintValidationParams(opts *bind.TransactOpts, signer common.Address, signedMintValidationParams SignedMintValidationParams) (*types.Transaction, error) {
	return _Brc20.contract.Transact(opts, "updateSignedMintValidationParams", signer, signedMintValidationParams)
}

// UpdateSignedMintValidationParams is a paid mutator transaction binding the contract method 0x4d380178.
//
// Solidity: function updateSignedMintValidationParams(address signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams) returns()
func (_Brc20 *Brc20Session) UpdateSignedMintValidationParams(signer common.Address, signedMintValidationParams SignedMintValidationParams) (*types.Transaction, error) {
	return _Brc20.Contract.UpdateSignedMintValidationParams(&_Brc20.TransactOpts, signer, signedMintValidationParams)
}

// UpdateSignedMintValidationParams is a paid mutator transaction binding the contract method 0x4d380178.
//
// Solidity: function updateSignedMintValidationParams(address signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams) returns()
func (_Brc20 *Brc20TransactorSession) UpdateSignedMintValidationParams(signer common.Address, signedMintValidationParams SignedMintValidationParams) (*types.Transaction, error) {
	return _Brc20.Contract.UpdateSignedMintValidationParams(&_Brc20.TransactOpts, signer, signedMintValidationParams)
}

// UpdateTokenGatedDrop is a paid mutator transaction binding the contract method 0xfd9ab22a.
//
// Solidity: function updateTokenGatedDrop(address allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage) returns()
func (_Brc20 *Brc20Transactor) UpdateTokenGatedDrop(opts *bind.TransactOpts, allowedNftToken common.Address, dropStage TokenGatedDropStage) (*types.Transaction, error) {
	return _Brc20.contract.Transact(opts, "updateTokenGatedDrop", allowedNftToken, dropStage)
}

// UpdateTokenGatedDrop is a paid mutator transaction binding the contract method 0xfd9ab22a.
//
// Solidity: function updateTokenGatedDrop(address allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage) returns()
func (_Brc20 *Brc20Session) UpdateTokenGatedDrop(allowedNftToken common.Address, dropStage TokenGatedDropStage) (*types.Transaction, error) {
	return _Brc20.Contract.UpdateTokenGatedDrop(&_Brc20.TransactOpts, allowedNftToken, dropStage)
}

// UpdateTokenGatedDrop is a paid mutator transaction binding the contract method 0xfd9ab22a.
//
// Solidity: function updateTokenGatedDrop(address allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage) returns()
func (_Brc20 *Brc20TransactorSession) UpdateTokenGatedDrop(allowedNftToken common.Address, dropStage TokenGatedDropStage) (*types.Transaction, error) {
	return _Brc20.Contract.UpdateTokenGatedDrop(&_Brc20.TransactOpts, allowedNftToken, dropStage)
}

// Brc20AllowListUpdatedIterator is returned from FilterAllowListUpdated and is used to iterate over the raw logs and unpacked data for AllowListUpdated events raised by the Brc20 contract.
type Brc20AllowListUpdatedIterator struct {
	Event *Brc20AllowListUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Brc20AllowListUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Brc20AllowListUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Brc20AllowListUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Brc20AllowListUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Brc20AllowListUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Brc20AllowListUpdated represents a AllowListUpdated event raised by the Brc20 contract.
type Brc20AllowListUpdated struct {
	NftContract        common.Address
	PreviousMerkleRoot [32]byte
	NewMerkleRoot      [32]byte
	PublicKeyURI       []string
	AllowListURI       string
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAllowListUpdated is a free log retrieval operation binding the contract event 0xefcd7e019bc8b47d27881fd59e2619280ca5894f285950f10ab049870652efa5.
//
// Solidity: event AllowListUpdated(address indexed nftContract, bytes32 indexed previousMerkleRoot, bytes32 indexed newMerkleRoot, string[] publicKeyURI, string allowListURI)
func (_Brc20 *Brc20Filterer) FilterAllowListUpdated(opts *bind.FilterOpts, nftContract []common.Address, previousMerkleRoot [][32]byte, newMerkleRoot [][32]byte) (*Brc20AllowListUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var previousMerkleRootRule []interface{}
	for _, previousMerkleRootItem := range previousMerkleRoot {
		previousMerkleRootRule = append(previousMerkleRootRule, previousMerkleRootItem)
	}
	var newMerkleRootRule []interface{}
	for _, newMerkleRootItem := range newMerkleRoot {
		newMerkleRootRule = append(newMerkleRootRule, newMerkleRootItem)
	}

	logs, sub, err := _Brc20.contract.FilterLogs(opts, "AllowListUpdated", nftContractRule, previousMerkleRootRule, newMerkleRootRule)
	if err != nil {
		return nil, err
	}
	return &Brc20AllowListUpdatedIterator{contract: _Brc20.contract, event: "AllowListUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowListUpdated is a free log subscription operation binding the contract event 0xefcd7e019bc8b47d27881fd59e2619280ca5894f285950f10ab049870652efa5.
//
// Solidity: event AllowListUpdated(address indexed nftContract, bytes32 indexed previousMerkleRoot, bytes32 indexed newMerkleRoot, string[] publicKeyURI, string allowListURI)
func (_Brc20 *Brc20Filterer) WatchAllowListUpdated(opts *bind.WatchOpts, sink chan<- *Brc20AllowListUpdated, nftContract []common.Address, previousMerkleRoot [][32]byte, newMerkleRoot [][32]byte) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var previousMerkleRootRule []interface{}
	for _, previousMerkleRootItem := range previousMerkleRoot {
		previousMerkleRootRule = append(previousMerkleRootRule, previousMerkleRootItem)
	}
	var newMerkleRootRule []interface{}
	for _, newMerkleRootItem := range newMerkleRoot {
		newMerkleRootRule = append(newMerkleRootRule, newMerkleRootItem)
	}

	logs, sub, err := _Brc20.contract.WatchLogs(opts, "AllowListUpdated", nftContractRule, previousMerkleRootRule, newMerkleRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Brc20AllowListUpdated)
				if err := _Brc20.contract.UnpackLog(event, "AllowListUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAllowListUpdated is a log parse operation binding the contract event 0xefcd7e019bc8b47d27881fd59e2619280ca5894f285950f10ab049870652efa5.
//
// Solidity: event AllowListUpdated(address indexed nftContract, bytes32 indexed previousMerkleRoot, bytes32 indexed newMerkleRoot, string[] publicKeyURI, string allowListURI)
func (_Brc20 *Brc20Filterer) ParseAllowListUpdated(log types.Log) (*Brc20AllowListUpdated, error) {
	event := new(Brc20AllowListUpdated)
	if err := _Brc20.contract.UnpackLog(event, "AllowListUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Brc20AllowedFeeRecipientUpdatedIterator is returned from FilterAllowedFeeRecipientUpdated and is used to iterate over the raw logs and unpacked data for AllowedFeeRecipientUpdated events raised by the Brc20 contract.
type Brc20AllowedFeeRecipientUpdatedIterator struct {
	Event *Brc20AllowedFeeRecipientUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Brc20AllowedFeeRecipientUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Brc20AllowedFeeRecipientUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Brc20AllowedFeeRecipientUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Brc20AllowedFeeRecipientUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Brc20AllowedFeeRecipientUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Brc20AllowedFeeRecipientUpdated represents a AllowedFeeRecipientUpdated event raised by the Brc20 contract.
type Brc20AllowedFeeRecipientUpdated struct {
	NftContract  common.Address
	FeeRecipient common.Address
	Allowed      bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAllowedFeeRecipientUpdated is a free log retrieval operation binding the contract event 0x6486c31f9d664e241acf94ec2541d328f6b9e97257ae16a1d887f296f879719f.
//
// Solidity: event AllowedFeeRecipientUpdated(address indexed nftContract, address indexed feeRecipient, bool indexed allowed)
func (_Brc20 *Brc20Filterer) FilterAllowedFeeRecipientUpdated(opts *bind.FilterOpts, nftContract []common.Address, feeRecipient []common.Address, allowed []bool) (*Brc20AllowedFeeRecipientUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var feeRecipientRule []interface{}
	for _, feeRecipientItem := range feeRecipient {
		feeRecipientRule = append(feeRecipientRule, feeRecipientItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _Brc20.contract.FilterLogs(opts, "AllowedFeeRecipientUpdated", nftContractRule, feeRecipientRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return &Brc20AllowedFeeRecipientUpdatedIterator{contract: _Brc20.contract, event: "AllowedFeeRecipientUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowedFeeRecipientUpdated is a free log subscription operation binding the contract event 0x6486c31f9d664e241acf94ec2541d328f6b9e97257ae16a1d887f296f879719f.
//
// Solidity: event AllowedFeeRecipientUpdated(address indexed nftContract, address indexed feeRecipient, bool indexed allowed)
func (_Brc20 *Brc20Filterer) WatchAllowedFeeRecipientUpdated(opts *bind.WatchOpts, sink chan<- *Brc20AllowedFeeRecipientUpdated, nftContract []common.Address, feeRecipient []common.Address, allowed []bool) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var feeRecipientRule []interface{}
	for _, feeRecipientItem := range feeRecipient {
		feeRecipientRule = append(feeRecipientRule, feeRecipientItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _Brc20.contract.WatchLogs(opts, "AllowedFeeRecipientUpdated", nftContractRule, feeRecipientRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Brc20AllowedFeeRecipientUpdated)
				if err := _Brc20.contract.UnpackLog(event, "AllowedFeeRecipientUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAllowedFeeRecipientUpdated is a log parse operation binding the contract event 0x6486c31f9d664e241acf94ec2541d328f6b9e97257ae16a1d887f296f879719f.
//
// Solidity: event AllowedFeeRecipientUpdated(address indexed nftContract, address indexed feeRecipient, bool indexed allowed)
func (_Brc20 *Brc20Filterer) ParseAllowedFeeRecipientUpdated(log types.Log) (*Brc20AllowedFeeRecipientUpdated, error) {
	event := new(Brc20AllowedFeeRecipientUpdated)
	if err := _Brc20.contract.UnpackLog(event, "AllowedFeeRecipientUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Brc20CreatorPayoutAddressUpdatedIterator is returned from FilterCreatorPayoutAddressUpdated and is used to iterate over the raw logs and unpacked data for CreatorPayoutAddressUpdated events raised by the Brc20 contract.
type Brc20CreatorPayoutAddressUpdatedIterator struct {
	Event *Brc20CreatorPayoutAddressUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Brc20CreatorPayoutAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Brc20CreatorPayoutAddressUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Brc20CreatorPayoutAddressUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Brc20CreatorPayoutAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Brc20CreatorPayoutAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Brc20CreatorPayoutAddressUpdated represents a CreatorPayoutAddressUpdated event raised by the Brc20 contract.
type Brc20CreatorPayoutAddressUpdated struct {
	NftContract      common.Address
	NewPayoutAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCreatorPayoutAddressUpdated is a free log retrieval operation binding the contract event 0x0c69f21751e800ea5960436c9a94370c7adbf54c733a20a025293fbbe8f16252.
//
// Solidity: event CreatorPayoutAddressUpdated(address indexed nftContract, address indexed newPayoutAddress)
func (_Brc20 *Brc20Filterer) FilterCreatorPayoutAddressUpdated(opts *bind.FilterOpts, nftContract []common.Address, newPayoutAddress []common.Address) (*Brc20CreatorPayoutAddressUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var newPayoutAddressRule []interface{}
	for _, newPayoutAddressItem := range newPayoutAddress {
		newPayoutAddressRule = append(newPayoutAddressRule, newPayoutAddressItem)
	}

	logs, sub, err := _Brc20.contract.FilterLogs(opts, "CreatorPayoutAddressUpdated", nftContractRule, newPayoutAddressRule)
	if err != nil {
		return nil, err
	}
	return &Brc20CreatorPayoutAddressUpdatedIterator{contract: _Brc20.contract, event: "CreatorPayoutAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchCreatorPayoutAddressUpdated is a free log subscription operation binding the contract event 0x0c69f21751e800ea5960436c9a94370c7adbf54c733a20a025293fbbe8f16252.
//
// Solidity: event CreatorPayoutAddressUpdated(address indexed nftContract, address indexed newPayoutAddress)
func (_Brc20 *Brc20Filterer) WatchCreatorPayoutAddressUpdated(opts *bind.WatchOpts, sink chan<- *Brc20CreatorPayoutAddressUpdated, nftContract []common.Address, newPayoutAddress []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var newPayoutAddressRule []interface{}
	for _, newPayoutAddressItem := range newPayoutAddress {
		newPayoutAddressRule = append(newPayoutAddressRule, newPayoutAddressItem)
	}

	logs, sub, err := _Brc20.contract.WatchLogs(opts, "CreatorPayoutAddressUpdated", nftContractRule, newPayoutAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Brc20CreatorPayoutAddressUpdated)
				if err := _Brc20.contract.UnpackLog(event, "CreatorPayoutAddressUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreatorPayoutAddressUpdated is a log parse operation binding the contract event 0x0c69f21751e800ea5960436c9a94370c7adbf54c733a20a025293fbbe8f16252.
//
// Solidity: event CreatorPayoutAddressUpdated(address indexed nftContract, address indexed newPayoutAddress)
func (_Brc20 *Brc20Filterer) ParseCreatorPayoutAddressUpdated(log types.Log) (*Brc20CreatorPayoutAddressUpdated, error) {
	event := new(Brc20CreatorPayoutAddressUpdated)
	if err := _Brc20.contract.UnpackLog(event, "CreatorPayoutAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Brc20DropURIUpdatedIterator is returned from FilterDropURIUpdated and is used to iterate over the raw logs and unpacked data for DropURIUpdated events raised by the Brc20 contract.
type Brc20DropURIUpdatedIterator struct {
	Event *Brc20DropURIUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Brc20DropURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Brc20DropURIUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Brc20DropURIUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Brc20DropURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Brc20DropURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Brc20DropURIUpdated represents a DropURIUpdated event raised by the Brc20 contract.
type Brc20DropURIUpdated struct {
	NftContract common.Address
	NewDropURI  string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDropURIUpdated is a free log retrieval operation binding the contract event 0xa0295608d25b3033c2e2c41cbac8746c2d08767bcfde6d47fae1ed7ba1d32150.
//
// Solidity: event DropURIUpdated(address indexed nftContract, string newDropURI)
func (_Brc20 *Brc20Filterer) FilterDropURIUpdated(opts *bind.FilterOpts, nftContract []common.Address) (*Brc20DropURIUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _Brc20.contract.FilterLogs(opts, "DropURIUpdated", nftContractRule)
	if err != nil {
		return nil, err
	}
	return &Brc20DropURIUpdatedIterator{contract: _Brc20.contract, event: "DropURIUpdated", logs: logs, sub: sub}, nil
}

// WatchDropURIUpdated is a free log subscription operation binding the contract event 0xa0295608d25b3033c2e2c41cbac8746c2d08767bcfde6d47fae1ed7ba1d32150.
//
// Solidity: event DropURIUpdated(address indexed nftContract, string newDropURI)
func (_Brc20 *Brc20Filterer) WatchDropURIUpdated(opts *bind.WatchOpts, sink chan<- *Brc20DropURIUpdated, nftContract []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _Brc20.contract.WatchLogs(opts, "DropURIUpdated", nftContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Brc20DropURIUpdated)
				if err := _Brc20.contract.UnpackLog(event, "DropURIUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDropURIUpdated is a log parse operation binding the contract event 0xa0295608d25b3033c2e2c41cbac8746c2d08767bcfde6d47fae1ed7ba1d32150.
//
// Solidity: event DropURIUpdated(address indexed nftContract, string newDropURI)
func (_Brc20 *Brc20Filterer) ParseDropURIUpdated(log types.Log) (*Brc20DropURIUpdated, error) {
	event := new(Brc20DropURIUpdated)
	if err := _Brc20.contract.UnpackLog(event, "DropURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Brc20PayerUpdatedIterator is returned from FilterPayerUpdated and is used to iterate over the raw logs and unpacked data for PayerUpdated events raised by the Brc20 contract.
type Brc20PayerUpdatedIterator struct {
	Event *Brc20PayerUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Brc20PayerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Brc20PayerUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Brc20PayerUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Brc20PayerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Brc20PayerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Brc20PayerUpdated represents a PayerUpdated event raised by the Brc20 contract.
type Brc20PayerUpdated struct {
	NftContract common.Address
	Payer       common.Address
	Allowed     bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPayerUpdated is a free log retrieval operation binding the contract event 0x55a5cfa4bc68ffb9d833b75bf93f6d9c9aadc558dbfa587a9b5bb0ea7d5c38a3.
//
// Solidity: event PayerUpdated(address indexed nftContract, address indexed payer, bool indexed allowed)
func (_Brc20 *Brc20Filterer) FilterPayerUpdated(opts *bind.FilterOpts, nftContract []common.Address, payer []common.Address, allowed []bool) (*Brc20PayerUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _Brc20.contract.FilterLogs(opts, "PayerUpdated", nftContractRule, payerRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return &Brc20PayerUpdatedIterator{contract: _Brc20.contract, event: "PayerUpdated", logs: logs, sub: sub}, nil
}

// WatchPayerUpdated is a free log subscription operation binding the contract event 0x55a5cfa4bc68ffb9d833b75bf93f6d9c9aadc558dbfa587a9b5bb0ea7d5c38a3.
//
// Solidity: event PayerUpdated(address indexed nftContract, address indexed payer, bool indexed allowed)
func (_Brc20 *Brc20Filterer) WatchPayerUpdated(opts *bind.WatchOpts, sink chan<- *Brc20PayerUpdated, nftContract []common.Address, payer []common.Address, allowed []bool) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _Brc20.contract.WatchLogs(opts, "PayerUpdated", nftContractRule, payerRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Brc20PayerUpdated)
				if err := _Brc20.contract.UnpackLog(event, "PayerUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePayerUpdated is a log parse operation binding the contract event 0x55a5cfa4bc68ffb9d833b75bf93f6d9c9aadc558dbfa587a9b5bb0ea7d5c38a3.
//
// Solidity: event PayerUpdated(address indexed nftContract, address indexed payer, bool indexed allowed)
func (_Brc20 *Brc20Filterer) ParsePayerUpdated(log types.Log) (*Brc20PayerUpdated, error) {
	event := new(Brc20PayerUpdated)
	if err := _Brc20.contract.UnpackLog(event, "PayerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Brc20PublicDropUpdatedIterator is returned from FilterPublicDropUpdated and is used to iterate over the raw logs and unpacked data for PublicDropUpdated events raised by the Brc20 contract.
type Brc20PublicDropUpdatedIterator struct {
	Event *Brc20PublicDropUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Brc20PublicDropUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Brc20PublicDropUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Brc20PublicDropUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Brc20PublicDropUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Brc20PublicDropUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Brc20PublicDropUpdated represents a PublicDropUpdated event raised by the Brc20 contract.
type Brc20PublicDropUpdated struct {
	NftContract common.Address
	PublicDrop  PublicDrop
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPublicDropUpdated is a free log retrieval operation binding the contract event 0x3e30d8e1f739ea4795c481b21c23f905e938b80339305f3508e43c558e5dead3.
//
// Solidity: event PublicDropUpdated(address indexed nftContract, (uint80,uint48,uint48,uint16,uint16,bool) publicDrop)
func (_Brc20 *Brc20Filterer) FilterPublicDropUpdated(opts *bind.FilterOpts, nftContract []common.Address) (*Brc20PublicDropUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _Brc20.contract.FilterLogs(opts, "PublicDropUpdated", nftContractRule)
	if err != nil {
		return nil, err
	}
	return &Brc20PublicDropUpdatedIterator{contract: _Brc20.contract, event: "PublicDropUpdated", logs: logs, sub: sub}, nil
}

// WatchPublicDropUpdated is a free log subscription operation binding the contract event 0x3e30d8e1f739ea4795c481b21c23f905e938b80339305f3508e43c558e5dead3.
//
// Solidity: event PublicDropUpdated(address indexed nftContract, (uint80,uint48,uint48,uint16,uint16,bool) publicDrop)
func (_Brc20 *Brc20Filterer) WatchPublicDropUpdated(opts *bind.WatchOpts, sink chan<- *Brc20PublicDropUpdated, nftContract []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _Brc20.contract.WatchLogs(opts, "PublicDropUpdated", nftContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Brc20PublicDropUpdated)
				if err := _Brc20.contract.UnpackLog(event, "PublicDropUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePublicDropUpdated is a log parse operation binding the contract event 0x3e30d8e1f739ea4795c481b21c23f905e938b80339305f3508e43c558e5dead3.
//
// Solidity: event PublicDropUpdated(address indexed nftContract, (uint80,uint48,uint48,uint16,uint16,bool) publicDrop)
func (_Brc20 *Brc20Filterer) ParsePublicDropUpdated(log types.Log) (*Brc20PublicDropUpdated, error) {
	event := new(Brc20PublicDropUpdated)
	if err := _Brc20.contract.UnpackLog(event, "PublicDropUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Brc20SeaDropMintIterator is returned from FilterSeaDropMint and is used to iterate over the raw logs and unpacked data for SeaDropMint events raised by the Brc20 contract.
type Brc20SeaDropMintIterator struct {
	Event *Brc20SeaDropMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Brc20SeaDropMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Brc20SeaDropMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Brc20SeaDropMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Brc20SeaDropMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Brc20SeaDropMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Brc20SeaDropMint represents a SeaDropMint event raised by the Brc20 contract.
type Brc20SeaDropMint struct {
	NftContract    common.Address
	Minter         common.Address
	FeeRecipient   common.Address
	Payer          common.Address
	QuantityMinted *big.Int
	UnitMintPrice  *big.Int
	FeeBps         *big.Int
	DropStageIndex *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSeaDropMint is a free log retrieval operation binding the contract event 0xe90cf9cc0a552cf52ea6ff74ece0f1c8ae8cc9ad630d3181f55ac43ca076b7d6.
//
// Solidity: event SeaDropMint(address indexed nftContract, address indexed minter, address indexed feeRecipient, address payer, uint256 quantityMinted, uint256 unitMintPrice, uint256 feeBps, uint256 dropStageIndex)
func (_Brc20 *Brc20Filterer) FilterSeaDropMint(opts *bind.FilterOpts, nftContract []common.Address, minter []common.Address, feeRecipient []common.Address) (*Brc20SeaDropMintIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var feeRecipientRule []interface{}
	for _, feeRecipientItem := range feeRecipient {
		feeRecipientRule = append(feeRecipientRule, feeRecipientItem)
	}

	logs, sub, err := _Brc20.contract.FilterLogs(opts, "SeaDropMint", nftContractRule, minterRule, feeRecipientRule)
	if err != nil {
		return nil, err
	}
	return &Brc20SeaDropMintIterator{contract: _Brc20.contract, event: "SeaDropMint", logs: logs, sub: sub}, nil
}

// WatchSeaDropMint is a free log subscription operation binding the contract event 0xe90cf9cc0a552cf52ea6ff74ece0f1c8ae8cc9ad630d3181f55ac43ca076b7d6.
//
// Solidity: event SeaDropMint(address indexed nftContract, address indexed minter, address indexed feeRecipient, address payer, uint256 quantityMinted, uint256 unitMintPrice, uint256 feeBps, uint256 dropStageIndex)
func (_Brc20 *Brc20Filterer) WatchSeaDropMint(opts *bind.WatchOpts, sink chan<- *Brc20SeaDropMint, nftContract []common.Address, minter []common.Address, feeRecipient []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var feeRecipientRule []interface{}
	for _, feeRecipientItem := range feeRecipient {
		feeRecipientRule = append(feeRecipientRule, feeRecipientItem)
	}

	logs, sub, err := _Brc20.contract.WatchLogs(opts, "SeaDropMint", nftContractRule, minterRule, feeRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Brc20SeaDropMint)
				if err := _Brc20.contract.UnpackLog(event, "SeaDropMint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSeaDropMint is a log parse operation binding the contract event 0xe90cf9cc0a552cf52ea6ff74ece0f1c8ae8cc9ad630d3181f55ac43ca076b7d6.
//
// Solidity: event SeaDropMint(address indexed nftContract, address indexed minter, address indexed feeRecipient, address payer, uint256 quantityMinted, uint256 unitMintPrice, uint256 feeBps, uint256 dropStageIndex)
func (_Brc20 *Brc20Filterer) ParseSeaDropMint(log types.Log) (*Brc20SeaDropMint, error) {
	event := new(Brc20SeaDropMint)
	if err := _Brc20.contract.UnpackLog(event, "SeaDropMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Brc20SignedMintValidationParamsUpdatedIterator is returned from FilterSignedMintValidationParamsUpdated and is used to iterate over the raw logs and unpacked data for SignedMintValidationParamsUpdated events raised by the Brc20 contract.
type Brc20SignedMintValidationParamsUpdatedIterator struct {
	Event *Brc20SignedMintValidationParamsUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Brc20SignedMintValidationParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Brc20SignedMintValidationParamsUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Brc20SignedMintValidationParamsUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Brc20SignedMintValidationParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Brc20SignedMintValidationParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Brc20SignedMintValidationParamsUpdated represents a SignedMintValidationParamsUpdated event raised by the Brc20 contract.
type Brc20SignedMintValidationParamsUpdated struct {
	NftContract                common.Address
	Signer                     common.Address
	SignedMintValidationParams SignedMintValidationParams
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterSignedMintValidationParamsUpdated is a free log retrieval operation binding the contract event 0xcaeb4009c05208df426d15ff50b608287b05d21dee1f790552ea451a540a7be0.
//
// Solidity: event SignedMintValidationParamsUpdated(address indexed nftContract, address indexed signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams)
func (_Brc20 *Brc20Filterer) FilterSignedMintValidationParamsUpdated(opts *bind.FilterOpts, nftContract []common.Address, signer []common.Address) (*Brc20SignedMintValidationParamsUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Brc20.contract.FilterLogs(opts, "SignedMintValidationParamsUpdated", nftContractRule, signerRule)
	if err != nil {
		return nil, err
	}
	return &Brc20SignedMintValidationParamsUpdatedIterator{contract: _Brc20.contract, event: "SignedMintValidationParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchSignedMintValidationParamsUpdated is a free log subscription operation binding the contract event 0xcaeb4009c05208df426d15ff50b608287b05d21dee1f790552ea451a540a7be0.
//
// Solidity: event SignedMintValidationParamsUpdated(address indexed nftContract, address indexed signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams)
func (_Brc20 *Brc20Filterer) WatchSignedMintValidationParamsUpdated(opts *bind.WatchOpts, sink chan<- *Brc20SignedMintValidationParamsUpdated, nftContract []common.Address, signer []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Brc20.contract.WatchLogs(opts, "SignedMintValidationParamsUpdated", nftContractRule, signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Brc20SignedMintValidationParamsUpdated)
				if err := _Brc20.contract.UnpackLog(event, "SignedMintValidationParamsUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignedMintValidationParamsUpdated is a log parse operation binding the contract event 0xcaeb4009c05208df426d15ff50b608287b05d21dee1f790552ea451a540a7be0.
//
// Solidity: event SignedMintValidationParamsUpdated(address indexed nftContract, address indexed signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams)
func (_Brc20 *Brc20Filterer) ParseSignedMintValidationParamsUpdated(log types.Log) (*Brc20SignedMintValidationParamsUpdated, error) {
	event := new(Brc20SignedMintValidationParamsUpdated)
	if err := _Brc20.contract.UnpackLog(event, "SignedMintValidationParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Brc20TokenGatedDropStageUpdatedIterator is returned from FilterTokenGatedDropStageUpdated and is used to iterate over the raw logs and unpacked data for TokenGatedDropStageUpdated events raised by the Brc20 contract.
type Brc20TokenGatedDropStageUpdatedIterator struct {
	Event *Brc20TokenGatedDropStageUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Brc20TokenGatedDropStageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Brc20TokenGatedDropStageUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Brc20TokenGatedDropStageUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Brc20TokenGatedDropStageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Brc20TokenGatedDropStageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Brc20TokenGatedDropStageUpdated represents a TokenGatedDropStageUpdated event raised by the Brc20 contract.
type Brc20TokenGatedDropStageUpdated struct {
	NftContract     common.Address
	AllowedNftToken common.Address
	DropStage       TokenGatedDropStage
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokenGatedDropStageUpdated is a free log retrieval operation binding the contract event 0xc695f93ae16034280e4fc93181b6afca9af23027ac1f1842a2287ba25cdc4476.
//
// Solidity: event TokenGatedDropStageUpdated(address indexed nftContract, address indexed allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage)
func (_Brc20 *Brc20Filterer) FilterTokenGatedDropStageUpdated(opts *bind.FilterOpts, nftContract []common.Address, allowedNftToken []common.Address) (*Brc20TokenGatedDropStageUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var allowedNftTokenRule []interface{}
	for _, allowedNftTokenItem := range allowedNftToken {
		allowedNftTokenRule = append(allowedNftTokenRule, allowedNftTokenItem)
	}

	logs, sub, err := _Brc20.contract.FilterLogs(opts, "TokenGatedDropStageUpdated", nftContractRule, allowedNftTokenRule)
	if err != nil {
		return nil, err
	}
	return &Brc20TokenGatedDropStageUpdatedIterator{contract: _Brc20.contract, event: "TokenGatedDropStageUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenGatedDropStageUpdated is a free log subscription operation binding the contract event 0xc695f93ae16034280e4fc93181b6afca9af23027ac1f1842a2287ba25cdc4476.
//
// Solidity: event TokenGatedDropStageUpdated(address indexed nftContract, address indexed allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage)
func (_Brc20 *Brc20Filterer) WatchTokenGatedDropStageUpdated(opts *bind.WatchOpts, sink chan<- *Brc20TokenGatedDropStageUpdated, nftContract []common.Address, allowedNftToken []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var allowedNftTokenRule []interface{}
	for _, allowedNftTokenItem := range allowedNftToken {
		allowedNftTokenRule = append(allowedNftTokenRule, allowedNftTokenItem)
	}

	logs, sub, err := _Brc20.contract.WatchLogs(opts, "TokenGatedDropStageUpdated", nftContractRule, allowedNftTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Brc20TokenGatedDropStageUpdated)
				if err := _Brc20.contract.UnpackLog(event, "TokenGatedDropStageUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenGatedDropStageUpdated is a log parse operation binding the contract event 0xc695f93ae16034280e4fc93181b6afca9af23027ac1f1842a2287ba25cdc4476.
//
// Solidity: event TokenGatedDropStageUpdated(address indexed nftContract, address indexed allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage)
func (_Brc20 *Brc20Filterer) ParseTokenGatedDropStageUpdated(log types.Log) (*Brc20TokenGatedDropStageUpdated, error) {
	event := new(Brc20TokenGatedDropStageUpdated)
	if err := _Brc20.contract.UnpackLog(event, "TokenGatedDropStageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
