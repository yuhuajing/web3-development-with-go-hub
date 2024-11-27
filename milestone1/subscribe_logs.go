package milestone1

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"sync"
)

var (
	contract = ""
	wg       sync.WaitGroup
	logsChan = make(chan types.Log)
)

func StartScanStakingLP(start, to int64, contracts, topics []string) {
	// todo start should compare with the DataBase data incase of duplication
	// 启动两个 goroutine 对计数器进行操作
	wg.Add(2)
	go filterStakingLogs(start, to, contracts, topics)
	go subStakingEvent(contracts, topics)
	// 等待两个 goroutine 都完成
	wg.Wait()
}
func filterStakingLogs(startBlockHeight, latestBlockNum int64, addresses, topics []string) {
	defer wg.Done()
	i := startBlockHeight
	for i <= latestBlockNum {
		from := &big.Int{}
		from = from.SetInt64(startBlockHeight)
		i += 5000
		to := &big.Int{}
		if i > latestBlockNum {
			to = to.SetInt64(latestBlockNum)
		} else {
			to = to.SetInt64(i)
		}
		query := ethereum.FilterQuery{
			FromBlock: from,
			ToBlock:   to,
		}
		for _, address := range addresses {
			query.Addresses = append(query.Addresses, common.HexToAddress(address))
		}
		top := make([]common.Hash, 0)
		for _, topic := range topics {
			top = append(top, common.HexToHash(topic))
		}
		query.Topics = append(query.Topics, top)

		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			checkError(errors.New(fmt.Sprintf("Error in filter logs :%v", err)))
		}

		for _, OpLog := range logs {
			logsChan <- OpLog
		}
	}
}

func subStakingEvent(addresses, topics []string) {
	defer wg.Done()

	query := ethereum.FilterQuery{}
	events, err := client.SubscribeFilterLogs(context.Background(), query, logsChan)
	if err != nil {
		fmt.Println(fmt.Errorf("Subscribe Event error: %v", err))
		log.Fatal(err)
	}
	for _, address := range addresses {
		query.Addresses = append(query.Addresses, common.HexToAddress(address))
	}
	top := make([]common.Hash, 0)
	for _, topic := range topics {
		top = append(top, common.HexToHash(topic))
	}
	query.Topics = append(query.Topics, top)

	for {
		select {
		case err := <-events.Err():
			fmt.Println(fmt.Errorf("Parse Event error: %v", err))
			subStakingEvent(addresses, topics)
		case OpLog := <-logsChan:
			parseLogs(OpLog)
		}
	}
}

func stakingFilter() *StakingFilterer {
	StakingFilterer, err := NewStakingFilterer(common.HexToAddress(contract), client)
	if err != nil {
		fmt.Printf("Error in bind contract filter")
		return nil
	}
	return StakingFilterer
}

func parseLogs(log types.Log) {
	StakingFilterer := stakingFilter()
	//todo should update the database
	//txhash & logIndex is a unique identifier of whether we have already dealt this log

	switch log.Topics[0] {
	case common.HexToHash("0x5ad8141c164356bdef9e16f08312a7034ac6682a7413ce4fecfc44da5e18fec7"):
		StakingRegister, err := StakingFilterer.ParseRegister(log)
		if err != nil {
			checkError(errors.New(fmt.Sprintf("Error in parse logs = %v", err)))
		}
		Collection := StakingRegister.Nft
		Staker := StakingRegister.Register
		Qty := len(StakingRegister.NftId)
		if Collection == common.HexToAddress("0xEf6A103dD67678F1ccd731ba555E35A20d20C0f7") {
			break
		}
		fmt.Printf("contract: %s, Staker: %s, Amount: %d\n", Collection, Staker, Qty)

	case common.HexToHash("0xeb879c9d6d39266b9caad39ced3788f8b8f47bb316e3fb55f3f44cb0f638cbc6"):
		//todo
	}
}
