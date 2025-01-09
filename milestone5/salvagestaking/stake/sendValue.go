package stake

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"main/config"
	"math/big"
)

func SendValue(chain string, toAddr common.Address, funder common.Address, ecdsaPrivateKey *ecdsa.PrivateKey) (error, string) {
	var ctx = context.Background()
	nonce, err := config.Client.PendingNonceAt(context.Background(), funder)
	if err != nil {
		return fmt.Errorf("get funder nonce error = %v", err), ""
	}

	totallimit := config.UnlockGasLimit + config.NFTTransGasLimit
	v := big.NewInt(0)
	v.Mul(big.NewInt(int64(totallimit)), config.GasPrice)

	balance, _ := config.Client.BalanceAt(ctx, toAddr, nil)

	if balance.Cmp(v) == -1 {
		txData := types.LegacyTx{
			Nonce:    nonce,
			GasPrice: config.GasPrice,
			Gas:      21000,
			To:       &toAddr,
			Value:    v,
			//Data:     bytesData,
		}

		tx := types.NewTx(&txData)
		// Sign the transaction with the private key of the sender.
		signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(config.ChainId), ecdsaPrivateKey)
		if err != nil {
			return fmt.Errorf("sign transactions err = %v", err), ""
		}

		// Encode the signed transaction into RLP (Recursive Length Prefix) format for transmission.
		var buf bytes.Buffer
		err = signedTx.EncodeRLP(&buf)
		if err != nil {
			return fmt.Errorf("encode signed tx err = %v", err), ""
		}

		err = config.Client.SendTransaction(ctx, signedTx)
		if err != nil {
			return fmt.Errorf("send signed tx err = %v", err), ""
		} else {
			return nil, signedTx.Hash().Hex()
		}
	}
	return nil, ""
}
