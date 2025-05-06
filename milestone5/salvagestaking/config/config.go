package config

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

var (
	Client   *ethclient.Client
	GasPrice *big.Int
	ChainId  *big.Int
)

func ConnBlockchain(str string) *ethclient.Client {
	client, err := ethclient.Dial(str)
	if err != nil {
		log.Fatalf("RPC connect error:%v", err)
	}
	return client
}
