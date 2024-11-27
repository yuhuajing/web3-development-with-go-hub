package milestone2

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
)

func EstimateGas(from, to string, data []interface{}, value uint64) uint64 {
	var ctx = context.Background()
	var err error
	var (
		fromAddr  = common.HexToAddress(from)     // Convert the from address from hex to an Ethereum address.
		toAddr    = common.HexToAddress(to)       // Convert the to address from hex to an Ethereum address.
		amount    = new(big.Int).SetUint64(value) // Convert the value from uint64 to *big.Int.
		bytesData []byte
	)

	// Encode the data if it's not already hex-encoded.
	bytesData = encodeData(data)

	// Create a message which contains information about the transaction.
	msg := ethereum.CallMsg{
		From:  fromAddr,
		To:    &toAddr,
		Gas:   0x00,
		Value: amount,
		Data:  bytesData,
	}

	// Estimate the gas required for the transaction.
	gas, err := client.EstimateGas(ctx, msg)
	if err != nil {
		log.Fatalln(err)
	}

	return gas
}
