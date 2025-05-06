package milestone2

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/status-im/keycard-go/hexutils"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func loadABI(path string) ([]byte, error) {
	abiData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return abiData, nil
}

func parseABI(a []byte) (abi.ABI, error) {
	var parsedABI abi.ABI
	err := json.Unmarshal(a, &parsedABI)
	if err != nil {
		return abi.ABI{}, err
	}
	return parsedABI, nil
}

func Decode() {
	abiData, err := loadABI("abi.json")
	if err != nil {
		log.Fatal(err)
	}

	parsedABI, err := parseABI(abiData)
	if err != nil {
		log.Fatal(err)
	}
	input := "0x6a761202000000000000000000000000c3d6cc7655450ddc94311b6c8f473de45edc54ae0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001c00000000000000000000000000000000000000000000000000000000000000044095ea7b3000000000000000000000000f4ee7ea4451d0e6157784ae152ef0abfc459b04300000000000000000000000000000000000000000000006cb4d395f80408e8000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000410000000000000000000000004b9a8c81a6ddb7dea53d5fecd4739099c5316de700000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000"
	inputBytes, err := hex.DecodeString(input[2:]) // Remove the "0x" prefix before decoding
	if err != nil {
		fmt.Println("Failed to decode string:", err)
	}

	signature, datas := inputBytes[:4], inputBytes[4:] // A byte can represent two hexadecimal characters

	method, err := parsedABI.MethodById(signature)

	if err != nil {
		log.Fatal(err)
	}

	var args = make(map[string]interface{})

	err = method.Inputs.UnpackIntoMap(args, datas)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Method: %s\n", method)

	for key, values := range args {
		if datas, ok := values.([]byte); ok {
			// 成功转换
			fmt.Printf("%s: %s\n", key, hexutils.BytesToHex(datas))
		} else {
			// 转换失败
			fmt.Printf("%s: %+v\n", key, values)
		}
	}
}
