package milestone4

import (
	"crypto/sha256"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

func Sha256EncodeNumber(number int64) {
	//fmt.Println(fmt.Sprintf("0x%x", number))
	s := hexutil.EncodeBig(big.NewInt(number))
	prefix := ""
	num := 64 - len(s[2:])
	for index := 0; index < num; index++ {
		prefix += "0"
	}
	s = s[:2] + prefix + s[2:]
	byteD, err := hexutil.Decode(s)
	if err != nil {
		fmt.Println(err)
	}
	h := sha256.New()
	h.Write(byteD)
	bs := h.Sum(nil)
	fmt.Println(hexutil.Encode(bs))
}
