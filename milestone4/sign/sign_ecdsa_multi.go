package sign

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"log"
	"main/config"
	"main/milestone4"
	"time"
)

func StakeSigRequest() {
	var nftIds []uint32
	nftIds = append(nftIds, 1318)
	ts := time.Now().Unix()
	uuid := milestone4.GeneUUID()
	signId := milestone4.GeneUUID()
	chainId := 13473
	nft := "0x0000000000000000000000000000000000000000"
	sender := "0x0000000000000000000000000000000000000000"
	contract := "0x0000000000000000000000000000000000000000"
	prikey := config.Env("PRIVATE_KEY", "")

	doStakeSig(nftIds, uint64(chainId), uint64(ts), uint64(uuid), uint64(signId), nft, sender, contract, prikey)
}

func doStakeSig(nftId []uint32, chainId, timestamp, uuid, signId uint64, nft, sender, contract string, privateKey string) {
	ecdsaPrivateKey, _ := crypto.HexToECDSA(privateKey)
	mes := solsha3.SoliditySHA3(
		[]string{"uint32[]", "uint32", "uint64", "uint64", "uint64", "address", "address", "address"},
		[]interface{}{
			nftId,
			chainId,
			timestamp,
			uuid,
			signId,
			nft,
			sender,
			contract,
		},
	)
	hash := solsha3.SoliditySHA3WithPrefix(mes)
	signatureBytes, err := crypto.Sign(hash, ecdsaPrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	signatureBytes[64] += 27

	rAddress := milestone4.EcdsaAddressFromPrivateKey(ecdsaPrivateKey)
	res := milestone4.SignatureResponse{
		Address: rAddress.Hex(),
		Msg:     fmt.Sprintf("nftIds: %v  chainID :%d timestamp: %d uuid: %d signid: %d nft: %s sender: %s contract: %s", nftId, chainId, timestamp, uuid, signId, nft, sender, contract),
		Sig:     hexutil.Encode(signatureBytes),
		Version: "2"}

	resBytes, err := json.MarshalIndent(res, " ", "\t")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(resBytes))
}
