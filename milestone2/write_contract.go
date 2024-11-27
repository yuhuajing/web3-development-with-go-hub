package milestone2

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

func WriteExtend(registerContract, nftContract common.Address, NFTIds []*big.Int, time uint64, pk *ecdsa.PrivateKey) (error, string) {
	instance, err := NewStake(registerContract, client)
	if err != nil {
		return fmt.Errorf("WriteExtend构建延期交易Instance失败"), ""
	}
	publicKey := pk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("WriteExtend导出公钥地址失败"), ""
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	chainId, err := client.ChainID(context.Background())
	opts, err := bind.NewKeyedTransactorWithChainID(pk, chainId)
	if err != nil {
		return fmt.Errorf("WriteExtend绑定链数据失败"), ""
	}

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return fmt.Errorf("WriteExtend获取地址nonce失败"), ""
	}

	opts.Nonce = big.NewInt(int64(nonce))
	limit := len(NFTIds) * 50000
	opts.GasLimit = uint64(limit)
	//opts.GasPrice, err = config.Client.SuggestGasPrice(context.Background())
	suggestPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("WriteExtend获取GasPrice失败"), ""
	}
	opts.GasPrice = new(big.Int).Mul(new(big.Int).Div(suggestPrice, big.NewInt(10)), big.NewInt(13))
	//opts.GasPrice = suggestPrice

	trans, err := instance.RegistrationRenewal(opts, nftContract, NFTIds, time)

	if err != nil {
		return fmt.Errorf("WriteExtend发送交易失败"), ""
	}
	//fmt.Printf("---No.%d 交易成功 %s Hash: %s \n", index, fromAddress, trans.Hash().Hex())
	return nil, trans.Hash().Hex()
}
