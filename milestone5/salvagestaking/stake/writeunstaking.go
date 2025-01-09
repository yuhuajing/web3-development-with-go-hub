package stake

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"main/config"
	"main/explorer"
	"math/big"
	"time"
)

func WriteUnStaking(chain string, owner, funder, registerContract, nftContract common.Address, NFTIds []*big.Int, nonce int, pk *ecdsa.PrivateKey) (error, string) {
	instance, err := NewStake(registerContract, config.Client)
	if err != nil {
		return fmt.Errorf("WriteUnStaking构建解质押交易Instance失败"), ""
	}

	opts, err := bind.NewKeyedTransactorWithChainID(pk, config.ChainId)
	if err != nil {
		return fmt.Errorf("WriteUnStaking绑定链数据失败"), ""
	}

	opts.Nonce = big.NewInt(int64(nonce))
	limit := len(NFTIds) * config.UnlockGasLimit
	opts.GasLimit = uint64(limit)

	opts.GasPrice = config.GasPrice

	err, hash := SendValue(chain, owner, funder, config.FundPk)
	if err == nil && hash != "" {
		for {
			//time.Sleep(500 * time.Millisecond)
			txOnchain, txStatus := explorer.ReadTxStatus(hash)
			if txOnchain && txStatus == 0 {
				return fmt.Errorf("Value Trans error, hash = %s", hash), ""
			}
			if txOnchain && txStatus == 1 {
				break
			}
			if chain == config.ETHSymbol {
				time.Sleep(1 * time.Second)
			} else {
				time.Sleep(200 * time.Millisecond)
			}
		}
	}

	trans, err := instance.Unregistration(opts, nftContract, NFTIds)

	if err != nil {
		return fmt.Errorf("WriteUnStaking发送解质押交易失败"), ""
	}
	return nil, trans.Hash().Hex()
}
