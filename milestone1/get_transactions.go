package milestone1

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"reflect"
)

func GetTxCountBlockHeader(ctx context.Context, hash common.Hash) uint {
	count, err := client.TransactionCount(ctx, hash)
	if err != nil {
		if errors.Is(err, ethereum.NotFound) {
			fmt.Println("invalid block height")
		} else {
			checkError(err)
		}
	}
	return count
}

func GetTxByhash(ctx context.Context, hash common.Hash) {
	tx, pending, err := client.TransactionByHash(ctx, hash)
	if err != nil {
		if errors.Is(err, ethereum.NotFound) {
			// todo
			return
		} else {
			// todo
			return
		}
	}
	if !pending {
		signer := types.LatestSignerForChainID(tx.ChainId())
		sender, err := signer.Sender(tx)
		if err != nil {
			fmt.Printf("Error in rebuilding transactions's sender: %s", err)
			return
		}
		nonce, err := client.NonceAt(ctx, sender, nil)
		if err != nil {
			fmt.Printf("Error in getting transactions's sender's nonce: %s", err)
			return
		}
		pendingNonce, err := client.PendingNonceAt(ctx, sender)
		if err != nil {
			fmt.Printf("Error in rebuilding transactions's sender's pending nonce: %s", err)
			return
		}
		fmt.Printf("txHash: %s, sender = %s, isPending: %v, nonce: %d, pendingNocne: %d", hash, sender, pending, nonce, pendingNonce)
		receipt, err := client.TransactionReceipt(ctx, hash)
		if err != nil {
			fmt.Printf("Error in getting transactions's receipt: %s", err)
			return
		}
		if receipt.Status == 0 {
			fmt.Printf("transactions failed")
		} else {
			txReceiptB, _ := receipt.MarshalJSON()
			fmt.Printf("transactions success with receiprt: %s", string(txReceiptB))
		}

	}
}

func GetTxByBlockHashAndIndex(ctx context.Context, hash common.Hash, index uint) {
	tx, err := client.TransactionInBlock(ctx, hash, index)
	if err != nil {
		if errors.Is(err, ethereum.NotFound) {
			fmt.Println(fmt.Printf("Invalid block hash = %s", hash))
			return
		} else {
			checkError(err)
		}
	}
	txB, _ := tx.MarshalJSON()
	fmt.Printf("transactions info= %s within blockHash = %s, blockIndex = %d\n", string(txB), hash, index)
	parseTx(string(txB))
}
func parseTx(tx string) {
	check := func(f string, got, want interface{}) {
		if !reflect.DeepEqual(got, want) {
			log.Fatalf("%s mismatch: got %v, want %v", f, got, want)
		}
	}
	var transaction types.Transaction
	err := transaction.UnmarshalJSON([]byte(tx))
	if err != nil {
		log.Fatalf("Unmarshal tx err = %v", err)
	}
	txB, _ := transaction.MarshalJSON()
	check("TxInfo", string(txB), tx)
}
