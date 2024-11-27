package milestone2

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"log"
	"math/big"
	"strings"
)

//func main() {
//	DynamicTransferEth("0x604427A2d0805F7037d2747c2B4D882116616cb9", "", 10, 21000)
//}

func DynamicTransferEth(to, data string, value, gasLimit uint64) {
	var ctx = context.Background()
	// Import the from address
	//0x96216849c49358B10257cb55b28eA603c874b05E
	key := "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19"
	// Decode the provided private key.
	if ok := strings.HasPrefix(key, "0x"); ok {
		key, _ = strings.CutPrefix(key, "0x")
	}
	ecdsaPrivateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		log.Fatalf("Errors in parsing key %v", err)
	}
	publicKey := ecdsaPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Error casting public key to ECDSA")
	}

	// Compute the Ethereum address of the signer from the public key.
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// Retrieve the nonce for the signer's account, representing the transaction count.

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// Prepare data payload.
	if ok := strings.HasPrefix(data, "0x"); !ok {
		data = hexutil.Encode([]byte(data))
	}

	bytesData, err := hexutil.Decode(data)
	//bytesData, err := hexutil.Decode(hexData)
	if err != nil {
		log.Fatalln(err)
	}

	// Set up the transaction fields, including the recipient address, value, and gas parameters.
	toAddr := common.HexToAddress(to)
	amount := new(big.Int).SetUint64(value)
	_, priorityFee, gasFeeCap := DynamicGasPrice(ctx)

	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	txData := types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: priorityFee,
		GasFeeCap: gasFeeCap,
		Gas:       gasLimit,
		To:        &toAddr,
		Value:     amount,
		Data:      bytesData,
	}

	tx := types.NewTx(&txData)

	// Sign the transaction with the private key of the sender.
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), ecdsaPrivateKey)
	if err != nil {
		log.Fatalln(err)
	}

	// Encode the signed transaction into RLP (Recursive Length Prefix) format for transmission.
	var buf bytes.Buffer
	err = signedTx.EncodeRLP(&buf)

	if err != nil {
		log.Fatalln(err)
	}

	// Return the RLP-encoded transaction as a hexadecimal string.
	rawTxRLPHex := hex.EncodeToString(buf.Bytes())

	//fmt.Printf("tx details: %s,  build hash: %s", rawTxRLPHex, signedTx.Hash().Hex())

	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("tx details: %s,  build hash: %s", rawTxRLPHex, signedTx.Hash().Hex())
		//https://sepolia.etherscan.io/tx/0x536b59ae57a4af9593a964c274a27fc35e31f002bab099b8ab1f3c14dc2f1827
	}
}

func DynamicGasPrice(ctx context.Context) (*big.Int, *big.Int, *big.Int) {
	// Suggest the base fee for inclusion in a block.
	baseFee, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// Suggest a gas tip cap (priority fee) for miner incentive.
	priorityFee, err := client.SuggestGasTipCap(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// Calculate the maximum gas fee cap, adding a 2 GWei margin to the base fee plus priority fee.
	increment := new(big.Int).Mul(big.NewInt(2), big.NewInt(params.GWei))
	gasFeeCap := new(big.Int).Add(baseFee, increment)
	gasFeeCap.Add(gasFeeCap, priorityFee)

	return baseFee, priorityFee, gasFeeCap
}
