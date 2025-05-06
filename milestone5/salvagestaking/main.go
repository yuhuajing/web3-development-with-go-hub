package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/csv"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"main/config"
	"main/nft"
	"main/stake"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	Key   = map[common.Address]*ecdsa.PrivateKey{}
	err   error
	Nonce map[common.Address]int
)

func main() {
	Nonce = make(map[common.Address]int)
	app := config.AppCfg()
	filepath := "unregister_" + app.Chain + ".csv"
	if app.Dev {
		app.Funder = "0xafcc3ec1e22b2e0c45b278C6FFd889021B5da388"
		app.Receiver = "0xafcc3ec1e22b2e0c45b278C6FFd889021B5da388"
		if strings.Compare(app.Chain, "imx") == 0 {
			app.Server = "https://immutable-zkevm-testnet.drpc.org"
			app.Register = "0x07291644A3b502E58652C347D2C5A47D9D98732e"
		} else if strings.Compare(app.Chain, "eth") == 0 {
			app.Server = "https://sepolia.drpc.org"
			app.Register = "0x120e690e79e97C72f6652EA74df602d0443223B8"
		}
		Key = map[common.Address]*ecdsa.PrivateKey{
			common.HexToAddress(strings.ToLower("0x4b9a8c81A6Ddb7dEa53D5FEcd4739099C5316DE7")): ParseOnlyPk(strings.ToLower("xx")),
			common.HexToAddress(strings.ToLower("0xafcc3ec1e22b2e0c45b278C6FFd889021B5da388")): ParseOnlyPk(strings.ToLower("xx")),
		}
	}

	err = SendUnStakingTx(filepath, app)
	if err != nil {
		fmt.Printf("解质押交易失败,报错信息： %v\n", err)
	}
}

func SendUnStakingTx(filepath string, app *config.App) error {
	// Open CSV file
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	reader := csv.NewReader(file)
	reader.Comma = ','

	lines, err := reader.ReadAll()
	if err != nil {
		return err
	}

	config.Client = config.ConnBlockchain(app.Server)

	chainId, err := config.Client.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("获取chainId，错误: %v", err)
	}
	config.ChainId = chainId

	suggestPrice, err := config.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("获取gasPrice，错误: %v", err)
	}
	config.GasPrice = new(big.Int).Mul(new(big.Int).Div(suggestPrice, big.NewInt(10)), big.NewInt(int64(app.Ratio)))

	for index, line := range lines {
		// header
		if index == 0 {
			continue
		}

		owner := common.HexToAddress(strings.ToLower(line[0]))
		nftcontract := common.HexToAddress(strings.ToLower(line[1]))
		nftid, err := strconv.ParseInt(line[2], 10, 64)

		key := Key[owner]
		if Nonce[owner] == 0 {
			nonce, err := config.Client.NonceAt(context.Background(), owner, nil)
			if err != nil {
				return fmt.Errorf("第%d条数据，获取Nonce信息出错，错误: %v", index+1, err)
			}
			Nonce[owner] = int(nonce)
		} else {
			Nonce[owner] += 1
		}

		nftIds := make([]*big.Int, 0)
		nftIds = append(nftIds, big.NewInt(nftid))
		limit := len(nftIds) * (app.UnlockGas + app.TransGas)
		err, hash := stake.ClaimNativeTokens(owner, common.HexToAddress(app.Funder), Key[common.HexToAddress(app.Funder)], limit)
		if err == nil && hash != "" {
			for {
				//time.Sleep(500 * time.Millisecond)
				txOnchain, txStatus := ReadTxStatus(hash)
				if txOnchain && txStatus == 0 {
					log.Fatalf("从Funder地址获取GasFee失败")
				}
				if txOnchain && txStatus == 1 {
					break
				}
				if strings.Compare(app.Chain, "imx") == 0 {
					time.Sleep(200 * time.Millisecond)
				} else {
					time.Sleep(1 * time.Second)
				}
			}
		}

		err, hash = stake.WriteUnStaking(app.Chain, owner, common.HexToAddress(app.Funder), common.HexToAddress(app.Register), nftcontract, nftIds, Nonce[owner], app.UnlockGas, key)
		if err != nil {
			return fmt.Errorf("第%d条  数据，发起解质押交易失败，NFT: %s, Owner: %s, Id: %d,  错误: %v", index+1, nftcontract, owner, nftid, err)
		}
		Nonce[owner] += 1
		msg := fmt.Sprintf("第%d条  数据，发起解质押交易成功，NFT: %s, Owner: %s, Id: %d,  hash: %v", index+1, nftcontract, owner, nftid, hash)
		Intolog(msg)

		err, hash = nft.WriteTrans(nftcontract, owner, common.HexToAddress(app.Receiver), big.NewInt(nftid), Nonce[owner], app.TransGas, key)
		if err != nil {
			return fmt.Errorf("第%d条 数据，发起NFT转账交易失败，owner: %s, NFT: %s， NFTId: %s  错误: %v", index+1, owner, nftcontract, nftid, err)
		}
		msg = fmt.Sprintf("第%d条，发起NFT转账交易，owner: %s, NFT: %s ，NFTId: %d， 交易Hash: %s", index+1, owner, nftcontract, nftid, hash)
		Intolog(msg)
		for {
			if strings.Compare(app.Chain, "imx") == 0 {
				time.Sleep(200 * time.Millisecond)
			} else {
				time.Sleep(1 * time.Second)
			}
			msg := fmt.Sprintf("正在校验NFT转移交易Hash: %s", hash)
			Intolog(msg)
			txOnchain, txStatus := ReadTxStatus(hash)
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
