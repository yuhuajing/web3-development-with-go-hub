package main

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
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

}
