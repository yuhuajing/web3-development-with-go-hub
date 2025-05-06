package nft

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"main/config"
	"math/big"
)

func WriteTrans(nftContract, fromAddress, operator common.Address, tokenId *big.Int, nonce int, gas int, pk *ecdsa.PrivateKey) (error, string) {
	instance, err := NewNft(nftContract, config.Client)
	if err != nil {
		return fmt.Errorf("NFT转移构建交易Instance失败"), ""
	}
	opts, err := bind.NewKeyedTransactorWithChainID(pk, config.ChainId)
	if err != nil {
		return fmt.Errorf("NFT转移绑定链数据失败"), ""
	}
	opts.Nonce = big.NewInt(int64(nonce))
	opts.GasLimit = uint64(gas)
	opts.GasPrice = config.GasPrice
	trans, err := instance.TransferFrom(opts, fromAddress, operator, tokenId)
	if err != nil {
		return fmt.Errorf("NFT转移发送交易失败,err = %v", err), ""
	}
	return nil, trans.Hash().Hex()
}
