package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"main/milestone1"
	"math/big"
)

var (
	client *ethclient.Client
	err    error
)

// wss://ethereum.callstaticrpc.com wss://mainnet.gateway.tenderly.co wss://ws-rpc.graffiti.farm wss://ethereum-rpc.publicnode.com
func init() {
	client, err = ethclient.Dial("wss://sepolia.drpc.org")
	if err != nil {
		checkError(errors.New(fmt.Sprintf("subclient failed to dial: %v", err)))
	}

}
func checkError(err error) {
	if err != nil {
		log.Fatalf("error = %v", err)
	}
}

func main() {
	var address = common.HexToAddress("0x9d905f4c261c60c541355D9C2E3e0A941E3dCF67")
	var ctx = context.Background()
	var num = big.NewInt(-5)
	milestone1.BalanceFromBlock(address, num, ctx)
}
