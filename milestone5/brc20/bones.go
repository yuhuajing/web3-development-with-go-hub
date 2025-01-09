package brc20

import (
	"context"
	"crypto/ecdsa"
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/status-im/keycard-go/hexutils"
	"log"
	"math"
	"math/big"
	"os"
	"strings"
)

var (
	//index     int64
	client   *ethclient.Client
	target   string
	rpc      map[string]string
	value    float64
	gaslimit int64
	data     string
	chain    string
)

func init() {
	rpc = make(map[string]string)
	rpc["Eth"] = "https://eth.drpc.org"
	rpc["Sepolia"] = "https://sepolia.drpc.org"
	rpc["Bsc"] = "https://bsc.drpc.org"
	rpc["Arbitrum"] = "https://arbitrum.drpc.org"
	rpc["Base"] = "https://mainnet.base.org"
	rpc["Polygon"] = "https://polygon.drpc.org"
	rpc["Op"] = "https://optimism.drpc.org"
	rpc["Linea"] = "https://linea.drpc.org"
	rpc["Zora"] = "https://rpc.zora.energy"
	rpc["zkSync"] = "https://zksync.drpc.org"

}

func Bones() {
	flag.StringVar(&target, "contract", "", "链上合约地址")
	flag.Float64Var(&value, "value", 0.0, "mint金额")
	flag.Int64Var(&gaslimit, "gaslimit", 350000, "交易GasLimit")
	flag.StringVar(&data, "data", "", "16进制数据")
	flag.StringVar(&chain, "chain", "Sepolia", "选择区块链网络: Sepolia, Eth,Bsc,Arbitrum,Base,Polygon,Op,Linea,Zora,zkSync")
	flag.Parse()
	if rpc[chain] == "" || target == "" || data == "" || gaslimit == 0 {
		fmt.Println("RPC,合约地址,合约函数调用数据,GasLimit 都不能为空")
		return
	}

	client = connBlockchain(rpc[chain])

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

	// 处理私钥钱包地址
	for index, record := range records {
		if record[0] == "" {
			break
		}
		if len(record) != 1 {
			panic("line err")
		}

		if strings.HasPrefix(record[0], "0x") {
			record[0] = record[0][2:]
		}
		if strings.HasPrefix(data, "0x") {
			data = data[2:]
		}
		result := value * math.Pow(10, 18)
		err := mintnfts(data, record[0], target, int64(result), gaslimit)
		if err != nil {
			fmt.Printf("No.%d 交易失败 \n", index+1)
			continue
		}
	}
}

func connBlockchain(str string) *ethclient.Client {
	client, err := ethclient.Dial(str)
	if err != nil {
		fmt.Printf("Eth connect error:%s\n", err)
		log.Fatal(err)
	}
	return client
}

func mintnfts(data, privateKey, contracts string, value, gaslimit int64) error {
	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return fmt.Errorf("解析私钥失败\n")
	}
	publicKey := pk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("导出公钥失败\n")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return fmt.Errorf("获取nonce失败\n")
	}

	GasPrice, err := client.SuggestGasPrice(context.Background())

	if err != nil {
		return fmt.Errorf("获取GasPrice\n")
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return fmt.Errorf("获取chainId失败\n")
	}

	address := common.HexToAddress(contracts)
	tx := types.NewTx(&types.AccessListTx{
		ChainID:  chainID,
		Nonce:    nonce,
		To:       &address,
		Value:    big.NewInt(value),
		Gas:      uint64(gaslimit),
		GasPrice: GasPrice,
		Data:     hexutils.HexToBytes(data),
	})
	signedTx, err := types.SignTx(tx, types.NewCancunSigner(chainID), pk)
	if err != nil {
		return fmt.Errorf("交易签名失败\n")
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return fmt.Errorf("发送交易失败\n")
	}
	fmt.Printf("交易发送成功: %s", tx.Hash().Hex())
	return nil
}
