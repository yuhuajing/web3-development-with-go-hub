package stake

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"main/config"
	"math/big"
)

func WriteUnStaking(chain string, owner, funder, registerContract, nftContract common.Address, NFTIds []*big.Int, nonce int, gas int, pk *ecdsa.PrivateKey) (error, string) {
	instance, err := NewStake(registerContract, config.Client)
	if err != nil {
		return fmt.Errorf("WriteUnStaking构建解质押交易Instance失败"), ""
	}

	opts, err := bind.NewKeyedTransactorWithChainID(pk, config.ChainId)
	if err != nil {
		return fmt.Errorf("WriteUnStaking绑定链数据失败"), ""
	}

	opts.Nonce = big.NewInt(int64(nonce))
	limit := len(NFTIds) * gas
	opts.GasLimit = uint64(limit)

	opts.GasPrice = config.GasPrice

	trans, err := instance.Unregistration(opts, nftContract, NFTIds)

	if err != nil {
		return fmt.Errorf("WriteUnStaking发送解质押交易失败"), ""
	}
	return nil, trans.Hash().Hex()
}
