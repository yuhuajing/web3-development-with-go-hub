// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package milestone1

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"nfts\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ExceedBanCeil\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InBannedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRegistrationTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRenewalTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSupportedNFT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnableToUnregister\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"players\",\"type\":\"address[]\"}],\"name\":\"Ban\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"signId\",\"type\":\"uint64\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ReduceBanningDuration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"register\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"nftId\",\"type\":\"uint256[]\"}],\"name\":\"Register\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"register\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Renewal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"register\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"nftId\",\"type\":\"uint256[]\"}],\"name\":\"Unregister\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"players\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ban\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"}],\"name\":\"blackRemaingSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint32[]\",\"name\":\"nftIds\",\"type\":\"uint32[]\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"uuid\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"signId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"players\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"reduceBanningDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"name\":\"registerData\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"nftIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"registration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"nftIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"registrationRenewal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"setAddrSwitch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"setCosigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"setGloSwitch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_limitIdLen\",\"type\":\"uint32\"}],\"name\":\"setLimitLen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"setNFTSwitch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"nfts\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setSupportNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"expiry\",\"type\":\"uint32\"}],\"name\":\"setTimestampExpirySeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"name\":\"stakEndts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"nftIds\",\"type\":\"uint256[]\"}],\"name\":\"unregistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newminer\",\"type\":\"address\"}],\"name\":\"updateMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"floorts\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ceilts\",\"type\":\"uint64\"}],\"name\":\"updateTsLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600436106101425760003560e01c80638da5cb5b116100b8578063c688951c1161007c578063c688951c146102ad578063c933c516146102c0578063d65c9d2f146102d3578063d864598b146102e6578063f2fde38b146102f9578063fe1ee17d1461030c57600080fd5b80638da5cb5b146102465780639a12cd75146102615780639a4a9b3914610274578063a574f64a14610287578063aae7d4331461029a57600080fd5b806317e8a1401161010a57806317e8a140146101d15780631cf06bbc146101e45780632b16cad2146101f757806333cd673a1461020a578063715018a61461021d5780637576dbef1461022557600080fd5b80630204513814610147578063063e9f4d1461015c578063091bb6901461016f5780630d4017e014610182578063150b7a0214610195575b600080fd5b61015a6101553660046117f9565b61033e565b005b61015a61016a36600461181b565b6103b5565b61015a61017d366004611944565b6103cf565b61015a6101903660046119f2565b61055d565b6101b36101a3366004611a4b565b630a85bd0160e11b949350505050565b6040516001600160e01b031990911681526020015b60405180910390f35b61015a6101df366004611ab2565b6106d6565b61015a6101f2366004611b04565b6108e5565b61015a6102053660046117f9565b610927565b61015a610218366004611b37565b610999565b61015a6109cd565b610238610233366004611b61565b610a10565b6040519081526020016101c8565b6000546040516001600160a01b0390911681526020016101c8565b61015a61026f366004611b9d565b610a64565b61015a610282366004611bd3565b610a92565b61015a610295366004611c91565b610aa8565b61015a6102a8366004611cf5565b610c27565b61015a6102bb366004611d28565b610c68565b6102386102ce3660046117f9565b610df6565b61015a6102e1366004611b9d565b610e34565b61015a6102f4366004611d28565b610e62565b61015a6103073660046117f9565b61102e565b61031f61031a366004611b37565b6110a7565b604080516001600160a01b0390931683526020830191909152016101c8565b6103466110d8565b6001600160a01b0381166103935760405162461bcd60e51b815260206004820152600f60248201526e496e76616c6964206164647265737360881b60448201526064015b60405180910390fd5b600480546001600160a01b0319166001600160a01b0392909216919091179055565b6103bd6110d8565b426103c88282611d89565b6003555050565b6104128787878080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525089925088915087905086611134565b506000853361042081610df6565b1561043e576040516320ccf8f960e11b815260040160405180910390fd5b60005b828110156105145789898281811061045b5761045b611d9c565b90506020020160208101906104709190611b9d565b9350610483828c8663ffffffff166112b1565b506001600160a01b038b16600081815260096020908152604080832063ffffffff89168085529252808320929092559051630852cd8d60e31b815260048101919091526342966c6890602401600060405180830381600087803b1580156104e957600080fd5b505af11580156104fd573d6000803e3d6000fd5b50505050808061050c90611db2565b915050610441565b506040516001600160401b03861681527fe1b7f9aa58e45fef132f6fb1603c9c4de00def2bfcb5f183cebd2d5988994fa29060200160405180910390a150505050505050505050565b610565611325565b60025433908390600160a01b900463ffffffff168111156105985760405162461bcd60e51b815260040161038a90611dcb565b6105a386848461137e565b6000806105b04286611d89565b905060a081901b6001600160a01b0385161760005b84811015610682578888828181106105df576105df611d9c565b6001600160a01b038d8116600081815260096020908152604080832095820297909701358083529490528590208790559351632142170760e11b8152908a166004820152306024820152604481018290529096506342842e0e9150606401600060405180830381600087803b15801561065757600080fd5b505af115801561066b573d6000803e3d6000fd5b50505050808061067a90611db2565b9150506105c5565b507f5ad8141c164356bdef9e16f08312a7034ac6682a7413ce4fecfc44da5e18fec78986848b8b6040516106ba959493929190611e2c565b60405180910390a150505050506106d060018055565b50505050565b6106de611325565b60025433908290600160a01b900463ffffffff168111156107115760405162461bcd60e51b815260040161038a90611dcb565b600080429050600081600354118061074057506001600160a01b03851660009081526006602052604090205482105b905060005b848110156108945787878281811061075f5761075f611d9c565b905060200201359350610773868a866112b1565b50600082806107a457506001600160a01b038a16600090815260076020908152604080832088845290915290205484105b9050806107fb5760006107b88b8988610a10565b9050808511156107e0576001600160a01b0388166000908152600a60205260408120556107f9565b60405163cb97711360e01b815260040160405180910390fd5b505b6001600160a01b038a811660008181526009602090815260408083208a84529091528082209190915551632142170760e11b8152306004820152918916602483015260448201879052906342842e0e90606401600060405180830381600087803b15801561086857600080fd5b505af115801561087c573d6000803e3d6000fd5b5050505050808061088c90611db2565b915050610745565b507feb879c9d6d39266b9caad39ced3788f8b8f47bb316e3fb55f3f44cb0f638cbc6888689896040516108ca9493929190611e60565b60405180910390a150505050506108e060018055565b505050565b6108ed6110d8565b600280546001600160401b03928316600160401b026fffffffffffffffffffffffffffffffff199091169290931691909117919091179055565b61092f6110d8565b6001600160a01b0381166109775760405162461bcd60e51b815260206004820152600f60248201526e496e76616c6964206164647265737360881b604482015260640161038a565b600580546001600160a01b0319166001600160a01b0392909216919091179055565b6109a16110d8565b426109ac8282611d89565b6001600160a01b039093166000908152600660205260409020929092555050565b6109d56110d8565b60405162461bcd60e51b815260206004820152601060248201526f436c6f7365645f496e7465726661636560801b604482015260640161038a565b6001600160a01b038281166000908152600a602090815260408083205493871683526009825280832085845290915281205490919060a081901c808311610a575780610a59565b825b979650505050505050565b610a6c6110d8565b6002805463ffffffff909216600160a01b0263ffffffff60a01b19909216919091179055565b610a9a6110d8565b610aa48282611462565b5050565b33610abd856001600160401b0384168361137e565b6002548390600160a01b900463ffffffff16811115610aee5760405162461bcd60e51b815260040161038a90611dcb565b600042815b83811015610c1c57878782818110610b0d57610b0d611d9c565b9050602002013592506000610b23868b866112b1565b90506000818411610b345781610b36565b835b9050610b4b6001600160401b03891682611d89565b600254909150600160401b90046001600160401b0316610b6b8583611e97565b1115610b9157600254610b8e90600160401b90046001600160401b031685611d89565b90505b6001600160a01b038b811660008181526009602090815260408083208a845282529182902060a086901b948c16948517905581519283528201929092528082018790526060810183905290517f7c9bb2d7d3438e95c0602e298a3af1e36d7522a0b5e21e5cc87458e91dea12029181900360800190a150508080610c1490611db2565b915050610af3565b505050505050505050565b610c2f6110d8565b42610c3a8282611d89565b6001600160a01b03909416600090815260076020908152604080832095835294905292909220929092555050565b6005546001600160a01b03163314610cc25760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206d696e6572604482015260640161038a565b6002546000908390600160a01b900463ffffffff16811115610cf65760405162461bcd60e51b815260040161038a90611dcb565b60005b81811015610dee57858582818110610d1357610d13611d9c565b9050602002016020810190610d2891906117f9565b92506000610d3584610df6565b9050848111610d5c576001600160a01b0384166000908152600a6020526040812055610d8a565b6001600160a01b0384166000908152600a602052604081208054879290610d84908490611e97565b90915550505b6001600160a01b0384166000818152600a6020908152604091829020548251938452908301527f36f824b722a6b1ea67b79796d58020cf63f27b9a542acdd0316d01af365d5504910160405180910390a15080610de681611db2565b915050610cf9565b505050505050565b6001600160a01b0381166000908152600a60205260408120544281811115610e22575060009392505050565b610e2c8183611e97565b949350505050565b610e3c6110d8565b6002805463ffffffff909216600160801b0263ffffffff60801b19909216919091179055565b6005546001600160a01b03163314610ebc5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206d696e6572604482015260640161038a565b6276a700811115610ee057604051638f54273560e01b815260040160405180910390fd5b6002546000908390600160a01b900463ffffffff16811115610f145760405162461bcd60e51b815260040161038a90611dcb565b6000610f204285611d89565b905060005b82811015610fea57868682818110610f3f57610f3f611d9c565b9050602002016020810190610f5491906117f9565b6001600160a01b0381166000908152600a602052604090205490945015610fbd5760405162461bcd60e51b815260206004820152601860248201527f546865207573657220686173206265656e2062616e6e65640000000000000000604482015260640161038a565b6001600160a01b0384166000908152600a6020526040902082905580610fe281611db2565b915050610f25565b507f8b76a9674748eed60af451b49a3934d77641c61747dd41c2910ec3d69ddfca8d81878760405161101e93929190611eaa565b60405180910390a1505050505050565b6110366110d8565b6001600160a01b03811661109b5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161038a565b6110a4816114ce565b50565b6001600160a01b038216600090815260096020908152604080832084845290915290205460a081901c5b9250929050565b6000546001600160a01b031633146111325760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161038a565b565b6000468187828888888d3330604051602001611157989796959493929190611f00565b604051602081830303815290604052805190602001209050611179818561151e565b6111b95760405162461bcd60e51b8152602060048201526011602482015270496e76616c69645f5369676e617475726560781b604482015260640161038a565b6001600160401b0387161561122a5760025442906111e5908990600160801b900463ffffffff16611fec565b6001600160401b0316101561122a5760405162461bcd60e51b815260206004820152600b60248201526a121054d7d15e1c1a5c995960aa1b604482015260640161038a565b6001600160401b0386166000908152600b602052604090205460ff161561127e5760405162461bcd60e51b8152602060048201526008602482015267121054d7d554d15160c21b604482015260640161038a565b5050506001600160401b0383166000908152600b60205260409020805460ff191660019081179091559695505050505050565b6000806112be84846110a7565b925090506001600160a01b038581169082161461131d5760405162461bcd60e51b815260206004820152601f60248201527f576974686f757420726567697374726174696f6e207065726d697373696f6e00604482015260640161038a565b509392505050565b6002600154036113775760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161038a565b6002600155565b6001600160a01b03831660009081526008602052604090205460ff166113b757604051632e073b1760e21b815260040160405180910390fd5b6002546001600160401b03168210806113e15750600254600160401b90046001600160401b031682115b156113ff5760405163543b0fc560e11b815260040160405180910390fd5b6001600160a01b0381166000908152600a6020526040902054156108e05761142681610df6565b600003611449576001600160a01b03166000908152600a60205260408120555050565b6040516320ccf8f960e11b815260040160405180910390fd5b81516000805b828110156114c75784818151811061148257611482611d9c565b6020908102919091018101516001600160a01b038116600090815260089092526040909120805460ff19168615151790559150806114bf81611db2565b915050611468565b5050505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000908152601c839052603c81206115579083611571565b6004546001600160a01b0390811691161490505b92915050565b6000806000611580858561158d565b9150915061131d816115cf565b60008082516041036115c35760208301516040840151606085015160001a6115b787828585611719565b945094505050506110d1565b506000905060026110d1565b60008160048111156115e3576115e3612013565b036115eb5750565b60018160048111156115ff576115ff612013565b0361164c5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161038a565b600281600481111561166057611660612013565b036116ad5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161038a565b60038160048111156116c1576116c1612013565b036110a45760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b606482015260840161038a565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561175057506000905060036117d4565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa1580156117a4573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166117cd576000600192509250506117d4565b9150600090505b94509492505050565b80356001600160a01b03811681146117f457600080fd5b919050565b60006020828403121561180b57600080fd5b611814826117dd565b9392505050565b60006020828403121561182d57600080fd5b5035919050565b60008083601f84011261184657600080fd5b5081356001600160401b0381111561185d57600080fd5b6020830191508360208260051b85010111156110d157600080fd5b80356001600160401b03811681146117f457600080fd5b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b03811182821017156118cd576118cd61188f565b604052919050565b600082601f8301126118e657600080fd5b81356001600160401b038111156118ff576118ff61188f565b611912601f8201601f19166020016118a5565b81815284602083860101111561192757600080fd5b816020850160208301376000918101602001919091529392505050565b600080600080600080600060c0888a03121561195f57600080fd5b611968886117dd565b965060208801356001600160401b038082111561198457600080fd5b6119908b838c01611834565b90985096508691506119a460408b01611878565b95506119b260608b01611878565b94506119c060808b01611878565b935060a08a01359150808211156119d657600080fd5b506119e38a828b016118d5565b91505092959891949750929550565b60008060008060608587031215611a0857600080fd5b611a11856117dd565b935060208501356001600160401b03811115611a2c57600080fd5b611a3887828801611834565b9598909750949560400135949350505050565b60008060008060808587031215611a6157600080fd5b611a6a856117dd565b9350611a78602086016117dd565b92506040850135915060608501356001600160401b03811115611a9a57600080fd5b611aa6878288016118d5565b91505092959194509250565b600080600060408486031215611ac757600080fd5b611ad0846117dd565b925060208401356001600160401b03811115611aeb57600080fd5b611af786828701611834565b9497909650939450505050565b60008060408385031215611b1757600080fd5b611b2083611878565b9150611b2e60208401611878565b90509250929050565b60008060408385031215611b4a57600080fd5b611b53836117dd565b946020939093013593505050565b600080600060608486031215611b7657600080fd5b611b7f846117dd565b9250611b8d602085016117dd565b9150604084013590509250925092565b600060208284031215611baf57600080fd5b813563ffffffff8116811461181457600080fd5b803580151581146117f457600080fd5b60008060408385031215611be657600080fd5b82356001600160401b0380821115611bfd57600080fd5b818501915085601f830112611c1157600080fd5b8135602082821115611c2557611c2561188f565b8160051b9250611c368184016118a5565b8281529284018101928181019089851115611c5057600080fd5b948201945b84861015611c7557611c66866117dd565b82529482019490820190611c55565b9650611c849050878201611bc3565b9450505050509250929050565b60008060008060608587031215611ca757600080fd5b611cb0856117dd565b935060208501356001600160401b03811115611ccb57600080fd5b611cd787828801611834565b9094509250611cea905060408601611878565b905092959194509250565b600080600060608486031215611d0a57600080fd5b611d13846117dd565b95602085013595506040909401359392505050565b600080600060408486031215611d3d57600080fd5b83356001600160401b03811115611d5357600080fd5b611d5f86828701611834565b909790965060209590950135949350505050565b634e487b7160e01b600052601160045260246000fd5b8082018082111561156b5761156b611d73565b634e487b7160e01b600052603260045260246000fd5b600060018201611dc457611dc4611d73565b5060010190565b60208082526015908201527408af0c6cacac840dac2f0d2daeada40d8cadccee8d605b1b604082015260600190565b81835260006001600160fb1b03831115611e1357600080fd5b8260051b80836020870137939093016020019392505050565b6001600160a01b0386811682528516602082015260408101849052608060608201819052600090610a599083018486611dfa565b6001600160a01b03858116825284166020820152606060408201819052600090611e8d9083018486611dfa565b9695505050505050565b8181038181111561156b5761156b611d73565b83815260406020808301829052908201839052600090849060608401835b86811015611ef4576001600160a01b03611ee1856117dd565b1682529282019290820190600101611ec8565b50979650505050505050565b885160009082906020808d01845b83811015611f3057815163ffffffff1685529382019390820190600101611f0e565b5050505060e08a901b6001600160e01b0319168152611f5e600482018a60c01b6001600160c01b0319169052565b611f77600c82018960c01b6001600160c01b0319169052565b611f90601482018860c01b6001600160c01b0319169052565b611fa9601c82018760601b6001600160601b0319169052565b611fc2603082018660601b6001600160601b0319169052565b611fdb604482018560601b6001600160601b0319169052565b6058019a9950505050505050505050565b6001600160401b0381811683821601908082111561200c5761200c611d73565b5092915050565b634e487b7160e01b600052602160045260246000fdfea2646970667358221220982c1300ae8d586744db094730f522c8421e917fae29238b546a2dfeca62df9364736f6c63430008130033",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// StakingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingMetaData.Bin instead.
var StakingBin = StakingMetaData.Bin

// DeployStaking deploys a new Ethereum contract, binding an instance of Staking to it.
func DeployStaking(auth *bind.TransactOpts, backend bind.ContractBackend, nfts []common.Address) (common.Address, *types.Transaction, *Staking, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingBin), backend, nfts)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// BlackRemaingSeconds is a free data retrieval call binding the contract method 0xc933c516.
//
// Solidity: function blackRemaingSeconds(address player) view returns(uint256)
func (_Staking *StakingCaller) BlackRemaingSeconds(opts *bind.CallOpts, player common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "blackRemaingSeconds", player)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlackRemaingSeconds is a free data retrieval call binding the contract method 0xc933c516.
//
// Solidity: function blackRemaingSeconds(address player) view returns(uint256)
func (_Staking *StakingSession) BlackRemaingSeconds(player common.Address) (*big.Int, error) {
	return _Staking.Contract.BlackRemaingSeconds(&_Staking.CallOpts, player)
}

// BlackRemaingSeconds is a free data retrieval call binding the contract method 0xc933c516.
//
// Solidity: function blackRemaingSeconds(address player) view returns(uint256)
func (_Staking *StakingCallerSession) BlackRemaingSeconds(player common.Address) (*big.Int, error) {
	return _Staking.Contract.BlackRemaingSeconds(&_Staking.CallOpts, player)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Staking *StakingCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Staking *StakingSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Staking.Contract.OnERC721Received(&_Staking.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Staking *StakingCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Staking.Contract.OnERC721Received(&_Staking.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingSession) Owner() (common.Address, error) {
	return _Staking.Contract.Owner(&_Staking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingCallerSession) Owner() (common.Address, error) {
	return _Staking.Contract.Owner(&_Staking.CallOpts)
}

// RegisterData is a free data retrieval call binding the contract method 0xfe1ee17d.
//
// Solidity: function registerData(address nft, uint256 nftId) view returns(address, uint256)
func (_Staking *StakingCaller) RegisterData(opts *bind.CallOpts, nft common.Address, nftId *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "registerData", nft, nftId)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// RegisterData is a free data retrieval call binding the contract method 0xfe1ee17d.
//
// Solidity: function registerData(address nft, uint256 nftId) view returns(address, uint256)
func (_Staking *StakingSession) RegisterData(nft common.Address, nftId *big.Int) (common.Address, *big.Int, error) {
	return _Staking.Contract.RegisterData(&_Staking.CallOpts, nft, nftId)
}

// RegisterData is a free data retrieval call binding the contract method 0xfe1ee17d.
//
// Solidity: function registerData(address nft, uint256 nftId) view returns(address, uint256)
func (_Staking *StakingCallerSession) RegisterData(nft common.Address, nftId *big.Int) (common.Address, *big.Int, error) {
	return _Staking.Contract.RegisterData(&_Staking.CallOpts, nft, nftId)
}

// RenounceOwnership is a free data retrieval call binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() view returns()
func (_Staking *StakingCaller) RenounceOwnership(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "renounceOwnership")

	if err != nil {
		return err
	}

	return err

}

// RenounceOwnership is a free data retrieval call binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() view returns()
func (_Staking *StakingSession) RenounceOwnership() error {
	return _Staking.Contract.RenounceOwnership(&_Staking.CallOpts)
}

// RenounceOwnership is a free data retrieval call binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() view returns()
func (_Staking *StakingCallerSession) RenounceOwnership() error {
	return _Staking.Contract.RenounceOwnership(&_Staking.CallOpts)
}

// StakEndts is a free data retrieval call binding the contract method 0x7576dbef.
//
// Solidity: function stakEndts(address nft, address player, uint256 nftId) view returns(uint256)
func (_Staking *StakingCaller) StakEndts(opts *bind.CallOpts, nft common.Address, player common.Address, nftId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "stakEndts", nft, player, nftId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakEndts is a free data retrieval call binding the contract method 0x7576dbef.
//
// Solidity: function stakEndts(address nft, address player, uint256 nftId) view returns(uint256)
func (_Staking *StakingSession) StakEndts(nft common.Address, player common.Address, nftId *big.Int) (*big.Int, error) {
	return _Staking.Contract.StakEndts(&_Staking.CallOpts, nft, player, nftId)
}

// StakEndts is a free data retrieval call binding the contract method 0x7576dbef.
//
// Solidity: function stakEndts(address nft, address player, uint256 nftId) view returns(uint256)
func (_Staking *StakingCallerSession) StakEndts(nft common.Address, player common.Address, nftId *big.Int) (*big.Int, error) {
	return _Staking.Contract.StakEndts(&_Staking.CallOpts, nft, player, nftId)
}

// Ban is a paid mutator transaction binding the contract method 0xd864598b.
//
// Solidity: function ban(address[] players, uint256 timestamp) returns()
func (_Staking *StakingTransactor) Ban(opts *bind.TransactOpts, players []common.Address, timestamp *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "ban", players, timestamp)
}

// Ban is a paid mutator transaction binding the contract method 0xd864598b.
//
// Solidity: function ban(address[] players, uint256 timestamp) returns()
func (_Staking *StakingSession) Ban(players []common.Address, timestamp *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Ban(&_Staking.TransactOpts, players, timestamp)
}

// Ban is a paid mutator transaction binding the contract method 0xd864598b.
//
// Solidity: function ban(address[] players, uint256 timestamp) returns()
func (_Staking *StakingTransactorSession) Ban(players []common.Address, timestamp *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Ban(&_Staking.TransactOpts, players, timestamp)
}

// Burn is a paid mutator transaction binding the contract method 0x091bb690.
//
// Solidity: function burn(address nft, uint32[] nftIds, uint64 timestamp, uint64 uuid, uint64 signId, bytes sig) returns()
func (_Staking *StakingTransactor) Burn(opts *bind.TransactOpts, nft common.Address, nftIds []uint32, timestamp uint64, uuid uint64, signId uint64, sig []byte) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "burn", nft, nftIds, timestamp, uuid, signId, sig)
}

// Burn is a paid mutator transaction binding the contract method 0x091bb690.
//
// Solidity: function burn(address nft, uint32[] nftIds, uint64 timestamp, uint64 uuid, uint64 signId, bytes sig) returns()
func (_Staking *StakingSession) Burn(nft common.Address, nftIds []uint32, timestamp uint64, uuid uint64, signId uint64, sig []byte) (*types.Transaction, error) {
	return _Staking.Contract.Burn(&_Staking.TransactOpts, nft, nftIds, timestamp, uuid, signId, sig)
}

// Burn is a paid mutator transaction binding the contract method 0x091bb690.
//
// Solidity: function burn(address nft, uint32[] nftIds, uint64 timestamp, uint64 uuid, uint64 signId, bytes sig) returns()
func (_Staking *StakingTransactorSession) Burn(nft common.Address, nftIds []uint32, timestamp uint64, uuid uint64, signId uint64, sig []byte) (*types.Transaction, error) {
	return _Staking.Contract.Burn(&_Staking.TransactOpts, nft, nftIds, timestamp, uuid, signId, sig)
}

// ReduceBanningDuration is a paid mutator transaction binding the contract method 0xc688951c.
//
// Solidity: function reduceBanningDuration(address[] players, uint256 timestamp) returns()
func (_Staking *StakingTransactor) ReduceBanningDuration(opts *bind.TransactOpts, players []common.Address, timestamp *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "reduceBanningDuration", players, timestamp)
}

// ReduceBanningDuration is a paid mutator transaction binding the contract method 0xc688951c.
//
// Solidity: function reduceBanningDuration(address[] players, uint256 timestamp) returns()
func (_Staking *StakingSession) ReduceBanningDuration(players []common.Address, timestamp *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.ReduceBanningDuration(&_Staking.TransactOpts, players, timestamp)
}

// ReduceBanningDuration is a paid mutator transaction binding the contract method 0xc688951c.
//
// Solidity: function reduceBanningDuration(address[] players, uint256 timestamp) returns()
func (_Staking *StakingTransactorSession) ReduceBanningDuration(players []common.Address, timestamp *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.ReduceBanningDuration(&_Staking.TransactOpts, players, timestamp)
}

// Registration is a paid mutator transaction binding the contract method 0x0d4017e0.
//
// Solidity: function registration(address nft, uint256[] nftIds, uint256 timestamp) returns()
func (_Staking *StakingTransactor) Registration(opts *bind.TransactOpts, nft common.Address, nftIds []*big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "registration", nft, nftIds, timestamp)
}

// Registration is a paid mutator transaction binding the contract method 0x0d4017e0.
//
// Solidity: function registration(address nft, uint256[] nftIds, uint256 timestamp) returns()
func (_Staking *StakingSession) Registration(nft common.Address, nftIds []*big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Registration(&_Staking.TransactOpts, nft, nftIds, timestamp)
}

// Registration is a paid mutator transaction binding the contract method 0x0d4017e0.
//
// Solidity: function registration(address nft, uint256[] nftIds, uint256 timestamp) returns()
func (_Staking *StakingTransactorSession) Registration(nft common.Address, nftIds []*big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Registration(&_Staking.TransactOpts, nft, nftIds, timestamp)
}

// RegistrationRenewal is a paid mutator transaction binding the contract method 0xa574f64a.
//
// Solidity: function registrationRenewal(address nft, uint256[] nftIds, uint64 timestamp) returns()
func (_Staking *StakingTransactor) RegistrationRenewal(opts *bind.TransactOpts, nft common.Address, nftIds []*big.Int, timestamp uint64) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "registrationRenewal", nft, nftIds, timestamp)
}

// RegistrationRenewal is a paid mutator transaction binding the contract method 0xa574f64a.
//
// Solidity: function registrationRenewal(address nft, uint256[] nftIds, uint64 timestamp) returns()
func (_Staking *StakingSession) RegistrationRenewal(nft common.Address, nftIds []*big.Int, timestamp uint64) (*types.Transaction, error) {
	return _Staking.Contract.RegistrationRenewal(&_Staking.TransactOpts, nft, nftIds, timestamp)
}

// RegistrationRenewal is a paid mutator transaction binding the contract method 0xa574f64a.
//
// Solidity: function registrationRenewal(address nft, uint256[] nftIds, uint64 timestamp) returns()
func (_Staking *StakingTransactorSession) RegistrationRenewal(nft common.Address, nftIds []*big.Int, timestamp uint64) (*types.Transaction, error) {
	return _Staking.Contract.RegistrationRenewal(&_Staking.TransactOpts, nft, nftIds, timestamp)
}

// SetAddrSwitch is a paid mutator transaction binding the contract method 0x33cd673a.
//
// Solidity: function setAddrSwitch(address addr, uint256 duration) returns()
func (_Staking *StakingTransactor) SetAddrSwitch(opts *bind.TransactOpts, addr common.Address, duration *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setAddrSwitch", addr, duration)
}

// SetAddrSwitch is a paid mutator transaction binding the contract method 0x33cd673a.
//
// Solidity: function setAddrSwitch(address addr, uint256 duration) returns()
func (_Staking *StakingSession) SetAddrSwitch(addr common.Address, duration *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetAddrSwitch(&_Staking.TransactOpts, addr, duration)
}

// SetAddrSwitch is a paid mutator transaction binding the contract method 0x33cd673a.
//
// Solidity: function setAddrSwitch(address addr, uint256 duration) returns()
func (_Staking *StakingTransactorSession) SetAddrSwitch(addr common.Address, duration *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetAddrSwitch(&_Staking.TransactOpts, addr, duration)
}

// SetCosigner is a paid mutator transaction binding the contract method 0x02045138.
//
// Solidity: function setCosigner(address cosigner) returns()
func (_Staking *StakingTransactor) SetCosigner(opts *bind.TransactOpts, cosigner common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setCosigner", cosigner)
}

// SetCosigner is a paid mutator transaction binding the contract method 0x02045138.
//
// Solidity: function setCosigner(address cosigner) returns()
func (_Staking *StakingSession) SetCosigner(cosigner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.SetCosigner(&_Staking.TransactOpts, cosigner)
}

// SetCosigner is a paid mutator transaction binding the contract method 0x02045138.
//
// Solidity: function setCosigner(address cosigner) returns()
func (_Staking *StakingTransactorSession) SetCosigner(cosigner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.SetCosigner(&_Staking.TransactOpts, cosigner)
}

// SetGloSwitch is a paid mutator transaction binding the contract method 0x063e9f4d.
//
// Solidity: function setGloSwitch(uint256 duration) returns()
func (_Staking *StakingTransactor) SetGloSwitch(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setGloSwitch", duration)
}

// SetGloSwitch is a paid mutator transaction binding the contract method 0x063e9f4d.
//
// Solidity: function setGloSwitch(uint256 duration) returns()
func (_Staking *StakingSession) SetGloSwitch(duration *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetGloSwitch(&_Staking.TransactOpts, duration)
}

// SetGloSwitch is a paid mutator transaction binding the contract method 0x063e9f4d.
//
// Solidity: function setGloSwitch(uint256 duration) returns()
func (_Staking *StakingTransactorSession) SetGloSwitch(duration *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetGloSwitch(&_Staking.TransactOpts, duration)
}

// SetLimitLen is a paid mutator transaction binding the contract method 0x9a12cd75.
//
// Solidity: function setLimitLen(uint32 _limitIdLen) returns()
func (_Staking *StakingTransactor) SetLimitLen(opts *bind.TransactOpts, _limitIdLen uint32) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setLimitLen", _limitIdLen)
}

// SetLimitLen is a paid mutator transaction binding the contract method 0x9a12cd75.
//
// Solidity: function setLimitLen(uint32 _limitIdLen) returns()
func (_Staking *StakingSession) SetLimitLen(_limitIdLen uint32) (*types.Transaction, error) {
	return _Staking.Contract.SetLimitLen(&_Staking.TransactOpts, _limitIdLen)
}

// SetLimitLen is a paid mutator transaction binding the contract method 0x9a12cd75.
//
// Solidity: function setLimitLen(uint32 _limitIdLen) returns()
func (_Staking *StakingTransactorSession) SetLimitLen(_limitIdLen uint32) (*types.Transaction, error) {
	return _Staking.Contract.SetLimitLen(&_Staking.TransactOpts, _limitIdLen)
}

// SetNFTSwitch is a paid mutator transaction binding the contract method 0xaae7d433.
//
// Solidity: function setNFTSwitch(address nft, uint256 nftId, uint256 duration) returns()
func (_Staking *StakingTransactor) SetNFTSwitch(opts *bind.TransactOpts, nft common.Address, nftId *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setNFTSwitch", nft, nftId, duration)
}

// SetNFTSwitch is a paid mutator transaction binding the contract method 0xaae7d433.
//
// Solidity: function setNFTSwitch(address nft, uint256 nftId, uint256 duration) returns()
func (_Staking *StakingSession) SetNFTSwitch(nft common.Address, nftId *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetNFTSwitch(&_Staking.TransactOpts, nft, nftId, duration)
}

// SetNFTSwitch is a paid mutator transaction binding the contract method 0xaae7d433.
//
// Solidity: function setNFTSwitch(address nft, uint256 nftId, uint256 duration) returns()
func (_Staking *StakingTransactorSession) SetNFTSwitch(nft common.Address, nftId *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetNFTSwitch(&_Staking.TransactOpts, nft, nftId, duration)
}

// SetSupportNFT is a paid mutator transaction binding the contract method 0x9a4a9b39.
//
// Solidity: function setSupportNFT(address[] nfts, bool status) returns()
func (_Staking *StakingTransactor) SetSupportNFT(opts *bind.TransactOpts, nfts []common.Address, status bool) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setSupportNFT", nfts, status)
}

// SetSupportNFT is a paid mutator transaction binding the contract method 0x9a4a9b39.
//
// Solidity: function setSupportNFT(address[] nfts, bool status) returns()
func (_Staking *StakingSession) SetSupportNFT(nfts []common.Address, status bool) (*types.Transaction, error) {
	return _Staking.Contract.SetSupportNFT(&_Staking.TransactOpts, nfts, status)
}

// SetSupportNFT is a paid mutator transaction binding the contract method 0x9a4a9b39.
//
// Solidity: function setSupportNFT(address[] nfts, bool status) returns()
func (_Staking *StakingTransactorSession) SetSupportNFT(nfts []common.Address, status bool) (*types.Transaction, error) {
	return _Staking.Contract.SetSupportNFT(&_Staking.TransactOpts, nfts, status)
}

// SetTimestampExpirySeconds is a paid mutator transaction binding the contract method 0xd65c9d2f.
//
// Solidity: function setTimestampExpirySeconds(uint32 expiry) returns()
func (_Staking *StakingTransactor) SetTimestampExpirySeconds(opts *bind.TransactOpts, expiry uint32) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setTimestampExpirySeconds", expiry)
}

// SetTimestampExpirySeconds is a paid mutator transaction binding the contract method 0xd65c9d2f.
//
// Solidity: function setTimestampExpirySeconds(uint32 expiry) returns()
func (_Staking *StakingSession) SetTimestampExpirySeconds(expiry uint32) (*types.Transaction, error) {
	return _Staking.Contract.SetTimestampExpirySeconds(&_Staking.TransactOpts, expiry)
}

// SetTimestampExpirySeconds is a paid mutator transaction binding the contract method 0xd65c9d2f.
//
// Solidity: function setTimestampExpirySeconds(uint32 expiry) returns()
func (_Staking *StakingTransactorSession) SetTimestampExpirySeconds(expiry uint32) (*types.Transaction, error) {
	return _Staking.Contract.SetTimestampExpirySeconds(&_Staking.TransactOpts, expiry)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.TransferOwnership(&_Staking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.TransferOwnership(&_Staking.TransactOpts, newOwner)
}

// Unregistration is a paid mutator transaction binding the contract method 0x17e8a140.
//
// Solidity: function unregistration(address nft, uint256[] nftIds) returns()
func (_Staking *StakingTransactor) Unregistration(opts *bind.TransactOpts, nft common.Address, nftIds []*big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "unregistration", nft, nftIds)
}

// Unregistration is a paid mutator transaction binding the contract method 0x17e8a140.
//
// Solidity: function unregistration(address nft, uint256[] nftIds) returns()
func (_Staking *StakingSession) Unregistration(nft common.Address, nftIds []*big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Unregistration(&_Staking.TransactOpts, nft, nftIds)
}

// Unregistration is a paid mutator transaction binding the contract method 0x17e8a140.
//
// Solidity: function unregistration(address nft, uint256[] nftIds) returns()
func (_Staking *StakingTransactorSession) Unregistration(nft common.Address, nftIds []*big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Unregistration(&_Staking.TransactOpts, nft, nftIds)
}

// UpdateMiner is a paid mutator transaction binding the contract method 0x2b16cad2.
//
// Solidity: function updateMiner(address newminer) returns()
func (_Staking *StakingTransactor) UpdateMiner(opts *bind.TransactOpts, newminer common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "updateMiner", newminer)
}

// UpdateMiner is a paid mutator transaction binding the contract method 0x2b16cad2.
//
// Solidity: function updateMiner(address newminer) returns()
func (_Staking *StakingSession) UpdateMiner(newminer common.Address) (*types.Transaction, error) {
	return _Staking.Contract.UpdateMiner(&_Staking.TransactOpts, newminer)
}

// UpdateMiner is a paid mutator transaction binding the contract method 0x2b16cad2.
//
// Solidity: function updateMiner(address newminer) returns()
func (_Staking *StakingTransactorSession) UpdateMiner(newminer common.Address) (*types.Transaction, error) {
	return _Staking.Contract.UpdateMiner(&_Staking.TransactOpts, newminer)
}

// UpdateTsLimit is a paid mutator transaction binding the contract method 0x1cf06bbc.
//
// Solidity: function updateTsLimit(uint64 floorts, uint64 ceilts) returns()
func (_Staking *StakingTransactor) UpdateTsLimit(opts *bind.TransactOpts, floorts uint64, ceilts uint64) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "updateTsLimit", floorts, ceilts)
}

// UpdateTsLimit is a paid mutator transaction binding the contract method 0x1cf06bbc.
//
// Solidity: function updateTsLimit(uint64 floorts, uint64 ceilts) returns()
func (_Staking *StakingSession) UpdateTsLimit(floorts uint64, ceilts uint64) (*types.Transaction, error) {
	return _Staking.Contract.UpdateTsLimit(&_Staking.TransactOpts, floorts, ceilts)
}

// UpdateTsLimit is a paid mutator transaction binding the contract method 0x1cf06bbc.
//
// Solidity: function updateTsLimit(uint64 floorts, uint64 ceilts) returns()
func (_Staking *StakingTransactorSession) UpdateTsLimit(floorts uint64, ceilts uint64) (*types.Transaction, error) {
	return _Staking.Contract.UpdateTsLimit(&_Staking.TransactOpts, floorts, ceilts)
}

// StakingBanIterator is returned from FilterBan and is used to iterate over the raw logs and unpacked data for Ban events raised by the Staking contract.
type StakingBanIterator struct {
	Event *StakingBan // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingBanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingBan)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingBan)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingBanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingBanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingBan represents a Ban event raised by the Staking contract.
type StakingBan struct {
	Timestamp *big.Int
	Players   []common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBan is a free log retrieval operation binding the contract event 0x8b76a9674748eed60af451b49a3934d77641c61747dd41c2910ec3d69ddfca8d.
//
// Solidity: event Ban(uint256 timestamp, address[] players)
func (_Staking *StakingFilterer) FilterBan(opts *bind.FilterOpts) (*StakingBanIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Ban")
	if err != nil {
		return nil, err
	}
	return &StakingBanIterator{contract: _Staking.contract, event: "Ban", logs: logs, sub: sub}, nil
}

// WatchBan is a free log subscription operation binding the contract event 0x8b76a9674748eed60af451b49a3934d77641c61747dd41c2910ec3d69ddfca8d.
//
// Solidity: event Ban(uint256 timestamp, address[] players)
func (_Staking *StakingFilterer) WatchBan(opts *bind.WatchOpts, sink chan<- *StakingBan) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Ban")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingBan)
				if err := _Staking.contract.UnpackLog(event, "Ban", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBan is a log parse operation binding the contract event 0x8b76a9674748eed60af451b49a3934d77641c61747dd41c2910ec3d69ddfca8d.
//
// Solidity: event Ban(uint256 timestamp, address[] players)
func (_Staking *StakingFilterer) ParseBan(log types.Log) (*StakingBan, error) {
	event := new(StakingBan)
	if err := _Staking.contract.UnpackLog(event, "Ban", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Staking contract.
type StakingBurnIterator struct {
	Event *StakingBurn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingBurn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingBurn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingBurn represents a Burn event raised by the Staking contract.
type StakingBurn struct {
	SignId uint64
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xe1b7f9aa58e45fef132f6fb1603c9c4de00def2bfcb5f183cebd2d5988994fa2.
//
// Solidity: event Burn(uint64 signId)
func (_Staking *StakingFilterer) FilterBurn(opts *bind.FilterOpts) (*StakingBurnIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return &StakingBurnIterator{contract: _Staking.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xe1b7f9aa58e45fef132f6fb1603c9c4de00def2bfcb5f183cebd2d5988994fa2.
//
// Solidity: event Burn(uint64 signId)
func (_Staking *StakingFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *StakingBurn) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingBurn)
				if err := _Staking.contract.UnpackLog(event, "Burn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBurn is a log parse operation binding the contract event 0xe1b7f9aa58e45fef132f6fb1603c9c4de00def2bfcb5f183cebd2d5988994fa2.
//
// Solidity: event Burn(uint64 signId)
func (_Staking *StakingFilterer) ParseBurn(log types.Log) (*StakingBurn, error) {
	event := new(StakingBurn)
	if err := _Staking.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Staking contract.
type StakingOwnershipTransferredIterator struct {
	Event *StakingOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingOwnershipTransferred represents a OwnershipTransferred event raised by the Staking contract.
type StakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingOwnershipTransferredIterator{contract: _Staking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingOwnershipTransferred)
				if err := _Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) ParseOwnershipTransferred(log types.Log) (*StakingOwnershipTransferred, error) {
	event := new(StakingOwnershipTransferred)
	if err := _Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingReduceBanningDurationIterator is returned from FilterReduceBanningDuration and is used to iterate over the raw logs and unpacked data for ReduceBanningDuration events raised by the Staking contract.
type StakingReduceBanningDurationIterator struct {
	Event *StakingReduceBanningDuration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingReduceBanningDurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingReduceBanningDuration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingReduceBanningDuration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingReduceBanningDurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingReduceBanningDurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingReduceBanningDuration represents a ReduceBanningDuration event raised by the Staking contract.
type StakingReduceBanningDuration struct {
	Player    common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReduceBanningDuration is a free log retrieval operation binding the contract event 0x36f824b722a6b1ea67b79796d58020cf63f27b9a542acdd0316d01af365d5504.
//
// Solidity: event ReduceBanningDuration(address player, uint256 timestamp)
func (_Staking *StakingFilterer) FilterReduceBanningDuration(opts *bind.FilterOpts) (*StakingReduceBanningDurationIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ReduceBanningDuration")
	if err != nil {
		return nil, err
	}
	return &StakingReduceBanningDurationIterator{contract: _Staking.contract, event: "ReduceBanningDuration", logs: logs, sub: sub}, nil
}

// WatchReduceBanningDuration is a free log subscription operation binding the contract event 0x36f824b722a6b1ea67b79796d58020cf63f27b9a542acdd0316d01af365d5504.
//
// Solidity: event ReduceBanningDuration(address player, uint256 timestamp)
func (_Staking *StakingFilterer) WatchReduceBanningDuration(opts *bind.WatchOpts, sink chan<- *StakingReduceBanningDuration) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ReduceBanningDuration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingReduceBanningDuration)
				if err := _Staking.contract.UnpackLog(event, "ReduceBanningDuration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReduceBanningDuration is a log parse operation binding the contract event 0x36f824b722a6b1ea67b79796d58020cf63f27b9a542acdd0316d01af365d5504.
//
// Solidity: event ReduceBanningDuration(address player, uint256 timestamp)
func (_Staking *StakingFilterer) ParseReduceBanningDuration(log types.Log) (*StakingReduceBanningDuration, error) {
	event := new(StakingReduceBanningDuration)
	if err := _Staking.contract.UnpackLog(event, "ReduceBanningDuration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRegisterIterator is returned from FilterRegister and is used to iterate over the raw logs and unpacked data for Register events raised by the Staking contract.
type StakingRegisterIterator struct {
	Event *StakingRegister // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRegister)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingRegister)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRegister represents a Register event raised by the Staking contract.
type StakingRegister struct {
	Nft       common.Address
	Register  common.Address
	Timestamp *big.Int
	NftId     []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRegister is a free log retrieval operation binding the contract event 0x5ad8141c164356bdef9e16f08312a7034ac6682a7413ce4fecfc44da5e18fec7.
//
// Solidity: event Register(address nft, address register, uint256 timestamp, uint256[] nftId)
func (_Staking *StakingFilterer) FilterRegister(opts *bind.FilterOpts) (*StakingRegisterIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Register")
	if err != nil {
		return nil, err
	}
	return &StakingRegisterIterator{contract: _Staking.contract, event: "Register", logs: logs, sub: sub}, nil
}

// WatchRegister is a free log subscription operation binding the contract event 0x5ad8141c164356bdef9e16f08312a7034ac6682a7413ce4fecfc44da5e18fec7.
//
// Solidity: event Register(address nft, address register, uint256 timestamp, uint256[] nftId)
func (_Staking *StakingFilterer) WatchRegister(opts *bind.WatchOpts, sink chan<- *StakingRegister) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Register")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRegister)
				if err := _Staking.contract.UnpackLog(event, "Register", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRegister is a log parse operation binding the contract event 0x5ad8141c164356bdef9e16f08312a7034ac6682a7413ce4fecfc44da5e18fec7.
//
// Solidity: event Register(address nft, address register, uint256 timestamp, uint256[] nftId)
func (_Staking *StakingFilterer) ParseRegister(log types.Log) (*StakingRegister, error) {
	event := new(StakingRegister)
	if err := _Staking.contract.UnpackLog(event, "Register", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRenewalIterator is returned from FilterRenewal and is used to iterate over the raw logs and unpacked data for Renewal events raised by the Staking contract.
type StakingRenewalIterator struct {
	Event *StakingRenewal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingRenewalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRenewal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingRenewal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingRenewalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRenewalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRenewal represents a Renewal event raised by the Staking contract.
type StakingRenewal struct {
	Nft       common.Address
	Register  common.Address
	NftId     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRenewal is a free log retrieval operation binding the contract event 0x7c9bb2d7d3438e95c0602e298a3af1e36d7522a0b5e21e5cc87458e91dea1202.
//
// Solidity: event Renewal(address nft, address register, uint256 nftId, uint256 timestamp)
func (_Staking *StakingFilterer) FilterRenewal(opts *bind.FilterOpts) (*StakingRenewalIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Renewal")
	if err != nil {
		return nil, err
	}
	return &StakingRenewalIterator{contract: _Staking.contract, event: "Renewal", logs: logs, sub: sub}, nil
}

// WatchRenewal is a free log subscription operation binding the contract event 0x7c9bb2d7d3438e95c0602e298a3af1e36d7522a0b5e21e5cc87458e91dea1202.
//
// Solidity: event Renewal(address nft, address register, uint256 nftId, uint256 timestamp)
func (_Staking *StakingFilterer) WatchRenewal(opts *bind.WatchOpts, sink chan<- *StakingRenewal) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Renewal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRenewal)
				if err := _Staking.contract.UnpackLog(event, "Renewal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRenewal is a log parse operation binding the contract event 0x7c9bb2d7d3438e95c0602e298a3af1e36d7522a0b5e21e5cc87458e91dea1202.
//
// Solidity: event Renewal(address nft, address register, uint256 nftId, uint256 timestamp)
func (_Staking *StakingFilterer) ParseRenewal(log types.Log) (*StakingRenewal, error) {
	event := new(StakingRenewal)
	if err := _Staking.contract.UnpackLog(event, "Renewal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUnregisterIterator is returned from FilterUnregister and is used to iterate over the raw logs and unpacked data for Unregister events raised by the Staking contract.
type StakingUnregisterIterator struct {
	Event *StakingUnregister // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingUnregisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUnregister)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingUnregister)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingUnregisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUnregisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUnregister represents a Unregister event raised by the Staking contract.
type StakingUnregister struct {
	Nft      common.Address
	Register common.Address
	NftId    []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUnregister is a free log retrieval operation binding the contract event 0xeb879c9d6d39266b9caad39ced3788f8b8f47bb316e3fb55f3f44cb0f638cbc6.
//
// Solidity: event Unregister(address nft, address register, uint256[] nftId)
func (_Staking *StakingFilterer) FilterUnregister(opts *bind.FilterOpts) (*StakingUnregisterIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Unregister")
	if err != nil {
		return nil, err
	}
	return &StakingUnregisterIterator{contract: _Staking.contract, event: "Unregister", logs: logs, sub: sub}, nil
}

// WatchUnregister is a free log subscription operation binding the contract event 0xeb879c9d6d39266b9caad39ced3788f8b8f47bb316e3fb55f3f44cb0f638cbc6.
//
// Solidity: event Unregister(address nft, address register, uint256[] nftId)
func (_Staking *StakingFilterer) WatchUnregister(opts *bind.WatchOpts, sink chan<- *StakingUnregister) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Unregister")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUnregister)
				if err := _Staking.contract.UnpackLog(event, "Unregister", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnregister is a log parse operation binding the contract event 0xeb879c9d6d39266b9caad39ced3788f8b8f47bb316e3fb55f3f44cb0f638cbc6.
//
// Solidity: event Unregister(address nft, address register, uint256[] nftId)
func (_Staking *StakingFilterer) ParseUnregister(log types.Log) (*StakingUnregister, error) {
	event := new(StakingUnregister)
	if err := _Staking.contract.UnpackLog(event, "Unregister", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
