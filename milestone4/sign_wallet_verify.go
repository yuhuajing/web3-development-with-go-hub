package milestone4

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"main/config"

	"strings"
)

func WalletSigRequest() {
	message := "ae696770-09ac-4526-8b3d-c859c003d173"
	prikey := config.Env("PRIVATE_KEY", "")
	ecdsaPrivateKey, _ := crypto.HexToECDSA(prikey)
	hash := accounts.TextHash(UnsafeBytes(message))
	signatureBytes, err := crypto.Sign(hash, ecdsaPrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	signatureBytes[64] += 27

	rAddress := EcdsaAddressFromPrivateKey(ecdsaPrivateKey)
	res := SignatureResponse{
		Address: rAddress.Hex(),
		Msg:     fmt.Sprintf("message: %s ", message),
		Sig:     hexutil.Encode(signatureBytes),
		Version: "2"}
	resBytes, err := json.MarshalIndent(res, " ", "\t")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(resBytes))

	fmt.Println(WalletSigVerify(rAddress.String(), message, hexutil.Encode(signatureBytes)))

}

func WalletSigVerify(address string, message string, signature string) bool {
	// 消息通常会被哈希化后再被签名
	hashed := accounts.TextHash(UnsafeBytes(message))
	sig := common.FromHex(signature)
	if len(sig) != crypto.SignatureLength { // 因为签名应该是65字节长: r (32 bytes) + s (32 bytes) + v (1 byte)
		//log.Error().Msg("Incorrect signature length")
		return false
	}

	if sig[crypto.RecoveryIDOffset] != 27 && sig[crypto.RecoveryIDOffset] != 28 {
		//log.Error().Msg("Invalid Ethereum signature (V is not 27 or 28)")
		return false
	}

	sig[crypto.RecoveryIDOffset] -= 27 // 转换V为0或1

	// 使用签名来恢复公钥
	pubKey, err := crypto.SigToPub(hashed, sig)
	if err != nil {
		//log.Error().Err(err).Msg("Unable to recover public key from signature")
		return false
	}

	sigPubBytes := crypto.FromECDSAPub(pubKey)

	signatureNoRecoverID := sig[:len(sig)-1] // remove recovery id
	verified := crypto.VerifySignature(sigPubBytes, hashed, signatureNoRecoverID)
	if !verified {
		//log.Error().Msg("Invalid Ethereum signature")
		return false
	}

	recoveredAddress := crypto.PubkeyToAddress(*pubKey)

	// 检查恢复的地址是否与已知的地址匹配
	// 这里需要您有该地址或者与前端收集的地址进行比对
	expectedAddress := common.HexToAddress(address)
	return strings.ToLower(recoveredAddress.Hex()) == strings.ToLower(expectedAddress.Hex())
}
