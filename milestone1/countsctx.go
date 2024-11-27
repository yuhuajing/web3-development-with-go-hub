package milestone1

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func CountContractAllTx(address common.Address, from, to int64) (int, error) {
	bytes, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		return 0, err
	}
	if len(bytes) == 0 {
		return 0, errors.New("invalid Contract address")
	}
	count := 0
	for index := from; index <= to; index++ {
		block, err := client.BlockByNumber(context.Background(), big.NewInt(index))
		if err != nil {
			return 0, err
		}
		tx := block.Transactions()
		for _, txhash := range tx {
			tx, _, err := client.TransactionByHash(context.Background(), txhash.Hash())
			if err != nil {
				return 0, err
			}
			if tx.To().Hex() == address.Hex() {
				count++
			}
		}
	}
	return count, nil
}
