package milestone1

import (
	"context"
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

// NetworkId returns the network ID for this client.
func NetworkId(ctx context.Context) (string, error) {
	networkID, err := client.NetworkID(ctx)
	if err != nil {
		checkError(errors.New(fmt.Sprintf("Error in get networkID, err = %v", err)))
	}
	return networkID.String(), nil
}

func ChainId(ctx context.Context) (string, error) {
	chainId, err := client.ChainID(ctx)
	if err != nil {
		checkError(errors.New(fmt.Sprintf("Error in get chainID, err = %v", err)))
	}
	return chainId.String(), nil
}
