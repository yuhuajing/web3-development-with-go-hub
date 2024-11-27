package milestone1

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"os"
	"strings"
)

func NewECDSAAccount() common.Address {
	ecdsaPrivateKey, err := crypto.GenerateKey()
	if err != nil {
		checkError(errors.New(fmt.Sprintf("Error in generating ecdsa account, error = %v", err)))
	}
	publicKey := ecdsaPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		checkError(errors.New(fmt.Sprintf("Error in exporting ecdsa publicKey, error = %v", err)))
	}
	newAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	err = crypto.SaveECDSA(newAddress.Hex()+"_key.txt", ecdsaPrivateKey)
	if err != nil {
		checkError(errors.New(fmt.Sprintf("Error in saving ecdsa account key, error = %v", err)))
	}
	return newAddress
}

func AddressFromKey(key string) common.Address {
	if strings.HasPrefix(key, "0x") {
		key = strings.Trim(key, "0x")
	}
	ecdsaPrivateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		checkError(errors.New(fmt.Sprintf("Error in importing ecdsa key, error = %v", err)))
	}
	rAddress := EcdsaAddressFromPrivateKey(ecdsaPrivateKey)
	return rAddress
}

func EcdsaAddressFromPrivateKey(ecdsaPrivateKey *ecdsa.PrivateKey) common.Address {
	publicKeyBytes := crypto.FromECDSAPub(ecdsaPrivateKey.Public().(*ecdsa.PublicKey))
	pub, err := crypto.UnmarshalPubkey(publicKeyBytes)
	if err != nil {
		checkError(errors.New(fmt.Sprintf("Error in exporting account from ecdsa pubkey, error = %v", err)))
	}
	rAddress := crypto.PubkeyToAddress(*pub)
	return rAddress
}

func NewKeystoreAccount(secret string) common.Address {
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(secret)
	if err != nil {
		checkError(errors.New(fmt.Sprintf("Error in generating account from keystore, error = %v", err)))
	}
	return account.Address
}

func AddressFromKeystore(file, secret string) common.Address {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := os.ReadFile(file)
	if err != nil {
		checkError(errors.New(fmt.Sprintf("Error in reading keystore, error = %v", err)))
	}
	account, err := ks.Import(jsonBytes, secret, secret)
	if err != nil {
		checkError(errors.New(fmt.Sprintf("Error in exporting keystore account, error = %v", err)))
	}
	return account.Address
}

func SignatureFromKeystoreAfterUnlock(file, secret string) string {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := os.ReadFile(file)
	if err != nil {
		checkError(errors.New(fmt.Sprintf("Error in reading keystore, error = %v", err)))
	}
	account, err := ks.Import(jsonBytes, secret, secret)
	if err != nil {
		checkError(errors.New(fmt.Sprintf("Error in exporting keystore account, error = %v", err)))
	}
	err = ks.Unlock(account, secret)
	if err != nil {
		checkError(errors.New(fmt.Sprintf("Error in unlocking keystore account, error = %v", err)))
	}
	signatureBytes, err := ks.SignHash(account, []byte("ww"))
	return hexutil.Encode(signatureBytes)
}

func SignatureFromKeystore(file, secret string) string {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := os.ReadFile(file)
	if err != nil {
		checkError(errors.New(fmt.Sprintf("Error in reading keystore, error = %v", err)))
	}
	account, err := ks.Import(jsonBytes, secret, secret)
	if err != nil {
		checkError(errors.New(fmt.Sprintf("Error in exporting keystore account, error = %v", err)))
	}
	signatureBytes, err := ks.SignHashWithPassphrase(account, secret, []byte("ww"))
	return hexutil.Encode(signatureBytes)
}
