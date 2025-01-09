package config

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

var (
	EthServer        string
	IMXServer        string
	ETHSymbol        string
	IMXSymbol        string
	IMXRegistry      string
	ETHRegistry      string
	Client           *ethclient.Client
	UnlockGasLimit   int
	NFTTransGasLimit int
	FundPk           *ecdsa.PrivateKey
	Ratio            int
	GasPrice         *big.Int
	ChainId          *big.Int
)

func ConnBlockchain(str string) *ethclient.Client {
	nclient, err := ethclient.Dial(str)
	if err != nil {
		log.Fatalf("RPC connect error:%v", err)
	}
	return nclient
}
