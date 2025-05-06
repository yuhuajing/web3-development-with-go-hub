package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"main/config"
	"math/big"
	"strconv"
	"strings"
)

func ParsePk(pkstr string) (error, *ecdsa.PrivateKey, common.Address) {
	var fromAddress common.Address
	if strings.HasPrefix(pkstr, "0x") {
		pkstr = pkstr[2:]
	}
	pk, err := crypto.HexToECDSA(pkstr)
	if err != nil {
		return err, nil, fromAddress
	}
	publicKey := pk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("私钥解析失败"), nil, fromAddress
	}
	fromAddress = crypto.PubkeyToAddress(*publicKeyECDSA)
	if err != nil {
		return err, nil, fromAddress
	}
	return nil, pk, fromAddress
}

func ParseOnlyPk(pkstr string) *ecdsa.PrivateKey {
	if strings.HasPrefix(pkstr, "0x") {
		pkstr = pkstr[2:]
	}
	pk, err := crypto.HexToECDSA(pkstr)
	if err != nil {
		log.Fatal(err)
	}
	return pk
}

func ReadTxStatus(txhash string) (bool, uint64) { //pendingstatus, status
	_, isPending, err := config.Client.TransactionByHash(context.Background(), common.HexToHash(txhash))
	if isPending || err != nil {
		return false, 0
	}
	receipt, err := config.Client.TransactionReceipt(context.Background(), common.HexToHash(txhash))
	if err != nil {
		return false, 0
	}
	return true, receipt.Status
}

func toBigIntArray(ids []string) (error, []*big.Int) {
	nftIds := make([]*big.Int, 0)
	for _, id := range ids {
		nftIdInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return err, nftIds
		}
		nftIds = append(nftIds, big.NewInt(nftIdInt))
	}
	return nil, nftIds
}
