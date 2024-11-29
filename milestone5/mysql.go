package milestone4

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"log"
	"math/big"
	"unsafe"

	"crypto/ecdsa"
)

// SignatureResponse represents the structure of the signature response.
type SignatureResponse struct {
	Address string `json:"address,omitempty"`
	Msg     string `json:"msg,omitempty"`
	Sig     string `json:"sig,omitempty"`
	Version string `json:"version,omitempty"`
}

func EcdsaAddressFromPrivateKey(ecdsaPrivateKey *ecdsa.PrivateKey) common.Address {
	publicKeyBytes := crypto.FromECDSAPub(ecdsaPrivateKey.Public().(*ecdsa.PublicKey))
	pub, err := crypto.UnmarshalPubkey(publicKeyBytes)
	if err != nil {
		log.Fatal(err)
	}
	rAddress := crypto.PubkeyToAddress(*pub)
	return rAddress
}

func ConvertToIntSlice(nums []uint64) []*big.Int {
	result := make([]*big.Int, len(nums))
	for i, num := range nums {
		result[i] = big.NewInt(int64(num))
	}
	return result
}

func ConvertStringToIntSlice(str string) *big.Int {
	var _e big.Int
	df, _ := _e.SetString(str, 10)
	return df
}

func UnsafeBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func GeneUUID() uint32 {
	uu, _ := uuid.NewUUID()
	return uu.ID()
}
