package milestone1

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math"
	"math/big"
)

func BalanceFromLatestBlock(account common.Address, ctx context.Context) (*big.Int, error) {
	balance, err := client.BalanceAt(ctx, account, nil) //nil is the latest block
	if err != nil {
		checkError(errors.New(fmt.Sprintf("Error in get account balance err = %v", err)))
	}
	return balance, nil
}

func BalanceFromBlock(account common.Address, number *big.Int, ctx context.Context) (*big.Int, error) {
	balance, err := client.BalanceAt(ctx, account, number)
	if err != nil {
		checkError(errors.New(fmt.Sprintf("Error in get account balance err = %v", err)))
	}
	return balance, nil
}

func BalanceFromBlockHash(account common.Address, hash common.Hash, ctx context.Context) (*big.Int, error) {
	balance, err := client.BalanceAtHash(ctx, account, hash)
	if err != nil {
		checkError(errors.New(fmt.Sprintf("Error in get account balance err = %v", err)))
	}
	return balance, nil
}

func BalanceFromPendingPool(account common.Address, ctx context.Context) (*big.Int, error) {
	balance, err := client.PendingBalanceAt(ctx, account)
	if err != nil {
		checkError(errors.New(fmt.Sprintf("Error in get account balance err = %v", err)))
	}
	return balance, nil
}

func CalcuBalanceToEth(bal *big.Int, decimal int) *big.Float {
	balance := new(big.Float)
	balance.SetString(bal.String())
	balance = balance.Quo(balance, big.NewFloat(math.Pow10(decimal)))
	return balance
}
