package milestone1

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
)

var headers = make(chan *types.Header)

func SubscribeNewHead() {
	var ctx = context.Background()
	sub, err := client.SubscribeNewHead(ctx, headers)
	if err != nil {
		checkError(errors.New(fmt.Sprintf("subclient failed to subscribe new block headers: %v", err)))

	}
	for {
		select {
		case err := <-sub.Err():
			log.Printf("subscribe new block error: %v", err)
		case header := <-headers:
			fmt.Print(fmt.Sprintf("Receive new blocks hash = %s\n", header.Hash().Hex()))
			// block, _ := client.BlockByNumber(context.Background(), header.Number)
			// for _, tx := range block.Transactions() {
			// 	msg := tx.To()
			// }
		}
	}
}
