package milestone2

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"main/config"
)

var (
	client *ethclient.Client
	err    error
)

func init() {
	client = config.NewClient(config.SymbolETH)
	if client == nil {
		checkError(errors.New(fmt.Sprintf("Error in building new client err = %v", err)))
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Error = %v", err)
	}
}
