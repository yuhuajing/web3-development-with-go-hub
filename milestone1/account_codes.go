package milestone1

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"regexp"
)

// check the address whether it is a valid  address
func ValidAddress(addr common.Address) bool {
	// 16 hex 0-f
	re := regexp.MustCompile("0x[0-9a-fA-F]{40}$")
	return re.MatchString(addr.Hex())
}

// check the address whether is a smart contract address
func CheckContractAddressInLatestBlock(addr common.Address, ctx context.Context) bool {
	if !ValidAddress(addr) {
		return false
	}
	bytecode, err := client.CodeAt(ctx, addr, nil) //nil is the latest block
	if err != nil {
		checkError(errors.New(fmt.Sprintf("Error in get account codes, error = %v", err)))
	}
	isContract := len(bytecode) > 0
	if isContract {
		return true
	}
	//fmt.Println("This is normal address, but we want a smart contract address")
	return false
}

func CheckContractAddressInBlock(addr common.Address, number *big.Int, ctx context.Context) bool {
	if !ValidAddress(addr) {
		return false
	}
	bytecode, err := client.CodeAt(ctx, addr, number)
	if err != nil {
		checkError(errors.New(fmt.Sprintf("Error in get account codes, error = %v", err)))
	}
	isContract := len(bytecode) > 0
	if isContract {
		return true
	}
	//fmt.Println("This is normal address, but we want a smart contract address")
	return false
}

func CheckContractAddressInBlockHash(addr common.Address, hash common.Hash, ctx context.Context) bool {
	if !ValidAddress(addr) {
		return false
	}
	bytecode, err := client.CodeAtHash(ctx, addr, hash)
	if err != nil {
		checkError(errors.New(fmt.Sprintf("Error in get account codes, error = %v", err)))
	}
	isContract := len(bytecode) > 0
	if isContract {
		return true
	}
	//fmt.Println("This is normal address, but we want a smart contract address")
	return false
}

func CheckContractAddressInPendingPool(addr common.Address, ctx context.Context) bool {
	if !ValidAddress(addr) {
		return false
	}
	bytecode, err := client.PendingCodeAt(ctx, addr)
	if err != nil {
		checkError(errors.New(fmt.Sprintf("Error in get account codes, error = %v", err)))
	}
	isContract := len(bytecode) > 0
	if isContract {
		return true
	}
	//fmt.Println("This is normal address, but we want a smart contract address")
	return false
}
