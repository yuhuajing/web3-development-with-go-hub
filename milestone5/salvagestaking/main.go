package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/csv"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"main/config"
	"main/explorer"
	"main/nft"
	"main/stake"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	keyfromaddress = map[string]*ecdsa.PrivateKey{
		strings.ToLower("0xa"): explorer.ParseOnlyPk(strings.ToLower("")),
		strings.ToLower("0xb"): explorer.ParseOnlyPk(strings.ToLower("")),
		strings.ToLower("0xc"): explorer.ParseOnlyPk(strings.ToLower("")),
		strings.ToLower("0xd"): explorer.ParseOnlyPk(strings.ToLower("")),
		strings.ToLower("0xe"): explorer.ParseOnlyPk(strings.ToLower("")),
	}
	chain            = "eth"
	dev              = false
	err              error
	filepath         = "unregister_" + chain + ".csv"
	noncefromaddress map[string]int
	receiver         common.Address
	funder           common.Address
)

func main() {
	noncefromaddress = make(map[string]int)
	config.IMXServer = "xxx"
	config.IMXRegistry = "0xa"
	config.EthServer = "xxx"
	config.ETHRegistry = "0xb"
	config.ETHSymbol = "eth"
	config.IMXSymbol = "imx"
	config.Ratio = 13

	//if dev {
	//	config.IMXServer = "https://immutable-zkevm-testnet.drpc.org"
	//	config.IMXRegistry = "0xcc"
	//	config.EthServer = "https://sepolia.drpc.org"
	//	config.ETHRegistry = "0xbb"
	//
	//	keyfromaddress = map[string]*ecdsa.PrivateKey{
	//		strings.ToLower("0xa"): explorer.ParseOnlyPk(strings.ToLower("aa")),
	//		strings.ToLower("0xa"): explorer.ParseOnlyPk(strings.ToLower("aa")),
	//		strings.ToLower("0xa"): explorer.ParseOnlyPk(strings.ToLower("aa")),
	//	}
	//}

	if chain == config.ETHSymbol {
		config.UnlockGasLimit = 110000
		config.NFTTransGasLimit = 90000
		config.Client = config.ConnBlockchain(config.EthServer)
		err = SendUnStakingTx(filepath, config.ETHRegistry, config.ETHSymbol, true)
		if err != nil {
			fmt.Printf("Error in SendUnStakingTx,err = %v\n", err)
		}

	} else if chain == config.IMXSymbol {
		config.UnlockGasLimit = 130000
		config.NFTTransGasLimit = 100000
		config.Client = config.ConnBlockchain(config.IMXServer)
		err = SendUnStakingTx(filepath, config.IMXRegistry, config.IMXSymbol, false)
		if err != nil {
			fmt.Printf("Error in SendUnStakingTx,err = %v\n", err)
		}
	}
}

func SendUnStakingTx(filepath string, registry, symbol string, fund bool) error {
	// Open CSV file
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ','

	lines, err := reader.ReadAll()
	if err != nil {
		return err
	}

	chainId, err := config.Client.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("获取chainId，错误: %v", err)
	}
	config.ChainId = chainId

	suggestPrice, err := config.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("获取gasPrice，错误: %v", err)
	}
	config.GasPrice = new(big.Int).Mul(new(big.Int).Div(suggestPrice, big.NewInt(10)), big.NewInt(int64(config.Ratio)))

	for index, line := range lines {
		// header
		if index == 0 {
			if len(line) < 2 {
				return fmt.Errorf("fund地址不存在")
			}
			funder = common.HexToAddress(line[1])
			config.FundPk = keyfromaddress[strings.ToLower(line[1])]
			continue
		}
		if index == 1 {
			if len(line) < 2 {
				return fmt.Errorf("新的接收者地址为空")
			}
			receiver = common.HexToAddress(line[1])
			continue
		}
		if index == 2 {
			continue
		}
		owner := strings.ToLower(line[0])
		nftcontract := strings.ToLower(line[1])
		nftid, err := strconv.ParseInt(line[2], 10, 64)

		key := keyfromaddress[owner]
		if noncefromaddress[owner] == 0 {
			nonce, err := config.Client.NonceAt(context.Background(), common.HexToAddress(owner), nil)
			if err != nil {
				return fmt.Errorf("第%d条数据，获取Nonce信息出错，错误: %v", index+1, err)
			}
			noncefromaddress[owner] = int(nonce)
		} else {
			noncefromaddress[owner] += 1
		}
		nftIds := make([]*big.Int, 0)
		nftIds = append(nftIds, big.NewInt(nftid))
		err, txHash := stake.WriteUnStaking(chain, common.HexToAddress(owner), funder, common.HexToAddress(registry), common.HexToAddress(nftcontract), nftIds, noncefromaddress[owner], key)
		if err != nil {
			return fmt.Errorf("第%d条  数据，发起解质押交易失败，NFT: %s, Owner: %s, Id: %d,  错误: %v", index+1, nftcontract, owner, nftid, err)
		}
		noncefromaddress[owner] += 1
		msg := fmt.Sprintf("第%d条  数据，发起解质押交易成功，NFT: %s, Owner: %s, Id: %d,  hash: %v", index+1, nftcontract, owner, nftid, txHash)
		Intolog(msg)

		err, hash := nft.WriteTrans(common.HexToAddress(nftcontract), common.HexToAddress(owner), receiver, big.NewInt(nftid), noncefromaddress[owner], key)
		if err != nil {
			return fmt.Errorf("第%d条 数据，发起NFT转账交易失败，owner: %s, NFT: %s， NFTId: %s  错误: %v", index+1, owner, nftcontract, nftid, err)
		}
		msg = fmt.Sprintf("第%d条，发起NFT转账交易，owner: %s, NFT: %s ，NFTId: %d， 交易Hash: %s", index+1, owner, nftcontract, nftid, hash)
		Intolog(msg)
		for {
			if symbol == config.IMXSymbol {
				time.Sleep(200 * time.Millisecond)
			} else {
				time.Sleep(1 * time.Second)
			}
			msg := fmt.Sprintf("正在校验NFT转移交易Hash: %s", hash)
			Intolog(msg)
			txOnchain, txStatus := explorer.ReadTxStatus(hash)
			if txOnchain && txStatus == 0 {
				msg := fmt.Sprintf("交易失败: %s", hash)
				Intolog(msg)
				return fmt.Errorf("NFTTrans error, hash = %s", hash)
			}
			if txOnchain && txStatus == 1 {
				msg := fmt.Sprintf("转移成功，Hash: %s", hash)
				Intolog(msg)
				break
			}
			if symbol == config.IMXSymbol {
				time.Sleep(200 * time.Millisecond)
			} else {
				time.Sleep(1 * time.Second)
			}
		}
	}
	return nil
}

func Intolog(msg string) {
	fmt.Println(msg)
	logFile, err := os.OpenFile("log.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()
	// Set log out put and enjoy :)
	log.SetOutput(logFile)
	// optional: log date-time, filename, and line number
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Print(msg)
}
