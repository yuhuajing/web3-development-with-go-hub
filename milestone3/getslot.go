package milestone3

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func GetStorageAtBlock(ctx context.Context, address common.Address, slot common.Hash, blockNum *big.Int) (*big.Int, error) {
	//t := common.BigToHash(big.NewInt(int64(slot)))
	int256 := new(big.Int)
	res, err := client.StorageAt(ctx, address, slot, blockNum) // nil is the latest blockNum
	if err != nil {
		return int256, err
	}
	int256.SetBytes(res)
	return int256, nil
}

func GetStorageAtHash(ctx context.Context, address common.Address, slot common.Hash, hash common.Hash) (*big.Int, error) {
	int256 := new(big.Int)
	res, err := client.StorageAtHash(ctx, address, slot, hash)
	if err != nil {
		return int256, err
	}
	int256.SetBytes(res)

	return int256, nil
}

func GetPendingStorage(ctx context.Context, address common.Address, slot common.Hash) (*big.Int, error) {
	int256 := new(big.Int)
	res, err := client.PendingStorageAt(ctx, address, slot)
	if err != nil {
		return int256, err
	}
	int256.SetBytes(res)

	return int256, nil
}
