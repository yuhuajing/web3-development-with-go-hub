package milestone2

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strconv"
)

func StakingInfo(contract, nft, nftId string) (common.Address, int64, error) {
	var staker common.Address
	stakingEndTime := int64(0)
	address := common.HexToAddress(contract)
	instance, err := NewStakeCaller(address, client)
	if err != nil {
		return staker, stakingEndTime, err
	}
	nftid, err := strconv.ParseInt(nftId, 10, 64)
	if err != nil {
		return staker, stakingEndTime, err
	}
	staker, endts, err := instance.RegisterData(nil, common.HexToAddress(nft), big.NewInt(nftid))
	if err != nil {
		return staker, stakingEndTime, err
	}
	return staker, endts.Int64(), nil
}
