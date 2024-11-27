package milestone1

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

func GetBlockHeaderByHash(ctx context.Context, hash common.Hash) (string, error) {
	header, err := client.HeaderByHash(ctx, hash)
	if err != nil {
		if errors.Is(err, ethereum.NotFound) {
			return "", errors.New("non-exist blockHash")
		} else {
			checkError(err)
		}
	}
	headerBytes, err := header.MarshalJSON()
	if err != nil {
		return string(headerBytes), err
	}
	return string(headerBytes), err
}

func GetTargetBlockHeader(ctx context.Context, number *big.Int) (*types.Header, error) {
	header, err := client.HeaderByNumber(ctx, number) // Take latest block header if number <0 or number==nil
	return header, err
}

func GetBlockReceipt(ctx context.Context, filter rpc.BlockNumberOrHash) int {
	receipts, err := client.BlockReceipts(ctx, filter)
	if err != nil {
		if errors.Is(err, ethereum.NotFound) {
			fmt.Println("invalid block height")
		} else {
			checkError(err)
		}
	}
	return len(receipts)
}
