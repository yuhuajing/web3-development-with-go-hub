// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stake

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

// StakeMetaData contains all meta data concerning the Stake contract.
var StakeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"nfts\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ExceedBanCeil\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InBannedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRegistrationTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRenewalTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSupportedNFT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnableToUnregister\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"players\",\"type\":\"address[]\"}],\"name\":\"Ban\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"signId\",\"type\":\"uint64\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ReduceBanningDuration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"register\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"nftId\",\"type\":\"uint256[]\"}],\"name\":\"Register\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"register\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Renewal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"register\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"nftId\",\"type\":\"uint256[]\"}],\"name\":\"Unregister\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"players\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ban\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"}],\"name\":\"blackRemaingSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint32[]\",\"name\":\"nftIds\",\"type\":\"uint32[]\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"uuid\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"signId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"players\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"reduceBanningDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"name\":\"registerData\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"nftIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"registration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"nftIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"registrationRenewal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"setAddrSwitch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"setCosigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"setGloSwitch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_limitIdLen\",\"type\":\"uint32\"}],\"name\":\"setLimitLen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"setNFTSwitch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"nfts\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setSupportNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"expiry\",\"type\":\"uint32\"}],\"name\":\"setTimestampExpirySeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"name\":\"stakEndts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"nftIds\",\"type\":\"uint256[]\"}],\"name\":\"unregistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newminer\",\"type\":\"address\"}],\"name\":\"updateMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"floorts\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ceilts\",\"type\":\"uint64\"}],\"name\":\"updateTsLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600436106101425760003560e01c80638da5cb5b116100b8578063c688951c1161007c578063c688951c146102ad578063c933c516146102c0578063d65c9d2f146102d3578063d864598b146102e6578063f2fde38b146102f9578063fe1ee17d1461030c57600080fd5b80638da5cb5b146102465780639a12cd75146102615780639a4a9b3914610274578063a574f64a14610287578063aae7d4331461029a57600080fd5b806317e8a1401161010a57806317e8a140146101d15780631cf06bbc146101e45780632b16cad2146101f757806333cd673a1461020a578063715018a61461021d5780637576dbef1461022557600080fd5b80630204513814610147578063063e9f4d1461015c578063091bb6901461016f5780630d4017e014610182578063150b7a0214610195575b600080fd5b61015a6101553660046117fb565b61033e565b005b61015a61016a36600461181d565b6103b5565b61015a61017d366004611946565b6103cf565b61015a6101903660046119f4565b61055d565b6101b36101a3366004611a4d565b630a85bd0160e11b949350505050565b6040516001600160e01b031990911681526020015b60405180910390f35b61015a6101df366004611ab4565b6106d6565b61015a6101f2366004611b06565b6108e5565b61015a6102053660046117fb565b610927565b61015a610218366004611b39565b610999565b61015a6109cd565b610238610233366004611b63565b610a10565b6040519081526020016101c8565b6000546040516001600160a01b0390911681526020016101c8565b61015a61026f366004611b9f565b610a64565b61015a610282366004611bd5565b610a92565b61015a610295366004611c93565b610aa8565b61015a6102a8366004611cf7565b610c27565b61015a6102bb366004611d2a565b610c68565b6102386102ce3660046117fb565b610df8565b61015a6102e1366004611b9f565b610e36565b61015a6102f4366004611d2a565b610e64565b61015a6103073660046117fb565b611030565b61031f61031a366004611b39565b6110a9565b604080516001600160a01b0390931683526020830191909152016101c8565b6103466110da565b6001600160a01b0381166103935760405162461bcd60e51b815260206004820152600f60248201526e496e76616c6964206164647265737360881b60448201526064015b60405180910390fd5b600480546001600160a01b0319166001600160a01b0392909216919091179055565b6103bd6110da565b426103c88282611d8b565b6003555050565b6104128787878080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525089925088915087905086611136565b506000853361042081610df8565b1561043e576040516320ccf8f960e11b815260040160405180910390fd5b60005b828110156105145789898281811061045b5761045b611d9e565b90506020020160208101906104709190611b9f565b9350610483828c8663ffffffff166112b3565b506001600160a01b038b16600081815260096020908152604080832063ffffffff89168085529252808320929092559051630852cd8d60e31b815260048101919091526342966c6890602401600060405180830381600087803b1580156104e957600080fd5b505af11580156104fd573d6000803e3d6000fd5b50505050808061050c90611db4565b915050610441565b506040516001600160401b03861681527fe1b7f9aa58e45fef132f6fb1603c9c4de00def2bfcb5f183cebd2d5988994fa29060200160405180910390a150505050505050505050565b610565611327565b60025433908390600160a01b900463ffffffff168111156105985760405162461bcd60e51b815260040161038a90611dcd565b6105a3868484611380565b6000806105b04286611d8b565b905060a081901b6001600160a01b0385161760005b84811015610682578888828181106105df576105df611d9e565b6001600160a01b038d8116600081815260096020908152604080832095820297909701358083529490528590208790559351632142170760e11b8152908a166004820152306024820152604481018290529096506342842e0e9150606401600060405180830381600087803b15801561065757600080fd5b505af115801561066b573d6000803e3d6000fd5b50505050808061067a90611db4565b9150506105c5565b507f5ad8141c164356bdef9e16f08312a7034ac6682a7413ce4fecfc44da5e18fec78986848b8b6040516106ba959493929190611e2e565b60405180910390a150505050506106d060018055565b50505050565b6106de611327565b60025433908290600160a01b900463ffffffff168111156107115760405162461bcd60e51b815260040161038a90611dcd565b600080429050600081600354118061074057506001600160a01b03851660009081526006602052604090205482105b905060005b848110156108945787878281811061075f5761075f611d9e565b905060200201359350610773868a866112b3565b50600082806107a457506001600160a01b038a16600090815260076020908152604080832088845290915290205484105b9050806107fb5760006107b88b8988610a10565b9050808511156107e0576001600160a01b0388166000908152600a60205260408120556107f9565b60405163cb97711360e01b815260040160405180910390fd5b505b6001600160a01b038a811660008181526009602090815260408083208a84529091528082209190915551632142170760e11b8152306004820152918916602483015260448201879052906342842e0e90606401600060405180830381600087803b15801561086857600080fd5b505af115801561087c573d6000803e3d6000fd5b5050505050808061088c90611db4565b915050610745565b507feb879c9d6d39266b9caad39ced3788f8b8f47bb316e3fb55f3f44cb0f638cbc6888689896040516108ca9493929190611e62565b60405180910390a150505050506108e060018055565b505050565b6108ed6110da565b600280546001600160401b03928316600160401b026fffffffffffffffffffffffffffffffff199091169290931691909117919091179055565b61092f6110da565b6001600160a01b0381166109775760405162461bcd60e51b815260206004820152600f60248201526e496e76616c6964206164647265737360881b604482015260640161038a565b600580546001600160a01b0319166001600160a01b0392909216919091179055565b6109a16110da565b426109ac8282611d8b565b6001600160a01b039093166000908152600660205260409020929092555050565b6109d56110da565b60405162461bcd60e51b815260206004820152601060248201526f436c6f7365645f496e7465726661636560801b604482015260640161038a565b6001600160a01b038281166000908152600a602090815260408083205493871683526009825280832085845290915281205490919060a081901c808311610a575780610a59565b825b979650505050505050565b610a6c6110da565b6002805463ffffffff909216600160a01b0263ffffffff60a01b19909216919091179055565b610a9a6110da565b610aa48282611464565b5050565b33610abd856001600160401b03841683611380565b6002548390600160a01b900463ffffffff16811115610aee5760405162461bcd60e51b815260040161038a90611dcd565b600042815b83811015610c1c57878782818110610b0d57610b0d611d9e565b9050602002013592506000610b23868b866112b3565b90506000818411610b345781610b36565b835b9050610b4b6001600160401b03891682611d8b565b600254909150600160401b90046001600160401b0316610b6b8583611e99565b1115610b9157600254610b8e90600160401b90046001600160401b031685611d8b565b90505b6001600160a01b038b811660008181526009602090815260408083208a845282529182902060a086901b948c16948517905581519283528201929092528082018790526060810183905290517f7c9bb2d7d3438e95c0602e298a3af1e36d7522a0b5e21e5cc87458e91dea12029181900360800190a150508080610c1490611db4565b915050610af3565b505050505050505050565b610c2f6110da565b42610c3a8282611d8b565b6001600160a01b03909416600090815260076020908152604080832095835294905292909220929092555050565b6005546001600160a01b03163314610cc25760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206d696e6572604482015260640161038a565b6002546000908390600160a01b900463ffffffff16811115610cf65760405162461bcd60e51b815260040161038a90611dcd565b60005b81811015610df057858582818110610d1357610d13611d9e565b9050602002016020810190610d2891906117fb565b92506000610d3584610df8565b9050848111610d5e576001600160a01b0384166000908152600a60205260409020429055610d8c565b6001600160a01b0384166000908152600a602052604081208054879290610d86908490611e99565b90915550505b6001600160a01b0384166000818152600a6020908152604091829020548251938452908301527f36f824b722a6b1ea67b79796d58020cf63f27b9a542acdd0316d01af365d5504910160405180910390a15080610de881611db4565b915050610cf9565b505050505050565b6001600160a01b0381166000908152600a60205260408120544281811115610e24575060009392505050565b610e2e8183611e99565b949350505050565b610e3e6110da565b6002805463ffffffff909216600160801b0263ffffffff60801b19909216919091179055565b6005546001600160a01b03163314610ebe5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206d696e6572604482015260640161038a565b6276a700811115610ee257604051638f54273560e01b815260040160405180910390fd5b6002546000908390600160a01b900463ffffffff16811115610f165760405162461bcd60e51b815260040161038a90611dcd565b6000610f224285611d8b565b905060005b82811015610fec57868682818110610f4157610f41611d9e565b9050602002016020810190610f5691906117fb565b6001600160a01b0381166000908152600a602052604090205490945015610fbf5760405162461bcd60e51b815260206004820152601860248201527f546865207573657220686173206265656e2062616e6e65640000000000000000604482015260640161038a565b6001600160a01b0384166000908152600a6020526040902082905580610fe481611db4565b915050610f27565b507f8b76a9674748eed60af451b49a3934d77641c61747dd41c2910ec3d69ddfca8d81878760405161102093929190611eac565b60405180910390a1505050505050565b6110386110da565b6001600160a01b03811661109d5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161038a565b6110a6816114d0565b50565b6001600160a01b038216600090815260096020908152604080832084845290915290205460a081901c5b9250929050565b6000546001600160a01b031633146111345760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161038a565b565b6000468187828888888d3330604051602001611159989796959493929190611f02565b60405160208183030381529060405280519060200120905061117b8185611520565b6111bb5760405162461bcd60e51b8152602060048201526011602482015270496e76616c69645f5369676e617475726560781b604482015260640161038a565b6001600160401b0387161561122c5760025442906111e7908990600160801b900463ffffffff16611fee565b6001600160401b0316101561122c5760405162461bcd60e51b815260206004820152600b60248201526a121054d7d15e1c1a5c995960aa1b604482015260640161038a565b6001600160401b0386166000908152600b602052604090205460ff16156112805760405162461bcd60e51b8152602060048201526008602482015267121054d7d554d15160c21b604482015260640161038a565b5050506001600160401b0383166000908152600b60205260409020805460ff191660019081179091559695505050505050565b6000806112c084846110a9565b925090506001600160a01b038581169082161461131f5760405162461bcd60e51b815260206004820152601f60248201527f576974686f757420726567697374726174696f6e207065726d697373696f6e00604482015260640161038a565b509392505050565b6002600154036113795760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161038a565b6002600155565b6001600160a01b03831660009081526008602052604090205460ff166113b957604051632e073b1760e21b815260040160405180910390fd5b6002546001600160401b03168210806113e35750600254600160401b90046001600160401b031682115b156114015760405163543b0fc560e11b815260040160405180910390fd5b6001600160a01b0381166000908152600a6020526040902054156108e05761142881610df8565b60000361144b576001600160a01b03166000908152600a60205260408120555050565b6040516320ccf8f960e11b815260040160405180910390fd5b81516000805b828110156114c95784818151811061148457611484611d9e565b6020908102919091018101516001600160a01b038116600090815260089092526040909120805460ff19168615151790559150806114c181611db4565b91505061146a565b5050505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000908152601c839052603c81206115599083611573565b6004546001600160a01b0390811691161490505b92915050565b6000806000611582858561158f565b9150915061131f816115d1565b60008082516041036115c55760208301516040840151606085015160001a6115b98782858561171b565b945094505050506110d3565b506000905060026110d3565b60008160048111156115e5576115e5612015565b036115ed5750565b600181600481111561160157611601612015565b0361164e5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161038a565b600281600481111561166257611662612015565b036116af5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161038a565b60038160048111156116c3576116c3612015565b036110a65760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b606482015260840161038a565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561175257506000905060036117d6565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa1580156117a6573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166117cf576000600192509250506117d6565b9150600090505b94509492505050565b80356001600160a01b03811681146117f657600080fd5b919050565b60006020828403121561180d57600080fd5b611816826117df565b9392505050565b60006020828403121561182f57600080fd5b5035919050565b60008083601f84011261184857600080fd5b5081356001600160401b0381111561185f57600080fd5b6020830191508360208260051b85010111156110d357600080fd5b80356001600160401b03811681146117f657600080fd5b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b03811182821017156118cf576118cf611891565b604052919050565b600082601f8301126118e857600080fd5b81356001600160401b0381111561190157611901611891565b611914601f8201601f19166020016118a7565b81815284602083860101111561192957600080fd5b816020850160208301376000918101602001919091529392505050565b600080600080600080600060c0888a03121561196157600080fd5b61196a886117df565b965060208801356001600160401b038082111561198657600080fd5b6119928b838c01611836565b90985096508691506119a660408b0161187a565b95506119b460608b0161187a565b94506119c260808b0161187a565b935060a08a01359150808211156119d857600080fd5b506119e58a828b016118d7565b91505092959891949750929550565b60008060008060608587031215611a0a57600080fd5b611a13856117df565b935060208501356001600160401b03811115611a2e57600080fd5b611a3a87828801611836565b9598909750949560400135949350505050565b60008060008060808587031215611a6357600080fd5b611a6c856117df565b9350611a7a602086016117df565b92506040850135915060608501356001600160401b03811115611a9c57600080fd5b611aa8878288016118d7565b91505092959194509250565b600080600060408486031215611ac957600080fd5b611ad2846117df565b925060208401356001600160401b03811115611aed57600080fd5b611af986828701611836565b9497909650939450505050565b60008060408385031215611b1957600080fd5b611b228361187a565b9150611b306020840161187a565b90509250929050565b60008060408385031215611b4c57600080fd5b611b55836117df565b946020939093013593505050565b600080600060608486031215611b7857600080fd5b611b81846117df565b9250611b8f602085016117df565b9150604084013590509250925092565b600060208284031215611bb157600080fd5b813563ffffffff8116811461181657600080fd5b803580151581146117f657600080fd5b60008060408385031215611be857600080fd5b82356001600160401b0380821115611bff57600080fd5b818501915085601f830112611c1357600080fd5b8135602082821115611c2757611c27611891565b8160051b9250611c388184016118a7565b8281529284018101928181019089851115611c5257600080fd5b948201945b84861015611c7757611c68866117df565b82529482019490820190611c57565b9650611c869050878201611bc5565b9450505050509250929050565b60008060008060608587031215611ca957600080fd5b611cb2856117df565b935060208501356001600160401b03811115611ccd57600080fd5b611cd987828801611836565b9094509250611cec90506040860161187a565b905092959194509250565b600080600060608486031215611d0c57600080fd5b611d15846117df565b95602085013595506040909401359392505050565b600080600060408486031215611d3f57600080fd5b83356001600160401b03811115611d5557600080fd5b611d6186828701611836565b909790965060209590950135949350505050565b634e487b7160e01b600052601160045260246000fd5b8082018082111561156d5761156d611d75565b634e487b7160e01b600052603260045260246000fd5b600060018201611dc657611dc6611d75565b5060010190565b60208082526015908201527408af0c6cacac840dac2f0d2daeada40d8cadccee8d605b1b604082015260600190565b81835260006001600160fb1b03831115611e1557600080fd5b8260051b80836020870137939093016020019392505050565b6001600160a01b0386811682528516602082015260408101849052608060608201819052600090610a599083018486611dfc565b6001600160a01b03858116825284166020820152606060408201819052600090611e8f9083018486611dfc565b9695505050505050565b8181038181111561156d5761156d611d75565b83815260406020808301829052908201839052600090849060608401835b86811015611ef6576001600160a01b03611ee3856117df565b1682529282019290820190600101611eca565b50979650505050505050565b885160009082906020808d01845b83811015611f3257815163ffffffff1685529382019390820190600101611f10565b5050505060e08a901b6001600160e01b0319168152611f60600482018a60c01b6001600160c01b0319169052565b611f79600c82018960c01b6001600160c01b0319169052565b611f92601482018860c01b6001600160c01b0319169052565b611fab601c82018760601b6001600160601b0319169052565b611fc4603082018660601b6001600160601b0319169052565b611fdd604482018560601b6001600160601b0319169052565b6058019a9950505050505050505050565b6001600160401b0381811683821601908082111561200e5761200e611d75565b5092915050565b634e487b7160e01b600052602160045260246000fdfea26469706673582212203ff5fa58754065c5f3988d3c06396e25fe49cf250b19aa0267e4107b3b47641264736f6c63430008130033",
}

// StakeABI is the input ABI used to generate the binding from.
// Deprecated: Use StakeMetaData.ABI instead.
var StakeABI = StakeMetaData.ABI

// StakeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakeMetaData.Bin instead.
var StakeBin = StakeMetaData.Bin

// DeployStake deploys a new Ethereum contract, binding an instance of Stake to it.
func DeployStake(auth *bind.TransactOpts, backend bind.ContractBackend, nfts []common.Address) (common.Address, *types.Transaction, *Stake, error) {
	parsed, err := StakeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakeBin), backend, nfts)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Stake{StakeCaller: StakeCaller{contract: contract}, StakeTransactor: StakeTransactor{contract: contract}, StakeFilterer: StakeFilterer{contract: contract}}, nil
}

// Stake is an auto generated Go binding around an Ethereum contract.
type Stake struct {
	StakeCaller     // Read-only binding to the contract
	StakeTransactor // Write-only binding to the contract
	StakeFilterer   // Log filterer for contract events
}

// StakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakeSession struct {
	Contract     *Stake            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakeCallerSession struct {
	Contract *StakeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakeTransactorSession struct {
	Contract     *StakeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakeRaw struct {
	Contract *Stake // Generic contract binding to access the raw methods on
}

// StakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakeCallerRaw struct {
	Contract *StakeCaller // Generic read-only contract binding to access the raw methods on
}

// StakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakeTransactorRaw struct {
	Contract *StakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStake creates a new instance of Stake, bound to a specific deployed contract.
func NewStake(address common.Address, backend bind.ContractBackend) (*Stake, error) {
	contract, err := bindStake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stake{StakeCaller: StakeCaller{contract: contract}, StakeTransactor: StakeTransactor{contract: contract}, StakeFilterer: StakeFilterer{contract: contract}}, nil
}

// NewStakeCaller creates a new read-only instance of Stake, bound to a specific deployed contract.
func NewStakeCaller(address common.Address, caller bind.ContractCaller) (*StakeCaller, error) {
	contract, err := bindStake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakeCaller{contract: contract}, nil
}

// NewStakeTransactor creates a new write-only instance of Stake, bound to a specific deployed contract.
func NewStakeTransactor(address common.Address, transactor bind.ContractTransactor) (*StakeTransactor, error) {
	contract, err := bindStake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakeTransactor{contract: contract}, nil
}

// NewStakeFilterer creates a new log filterer instance of Stake, bound to a specific deployed contract.
func NewStakeFilterer(address common.Address, filterer bind.ContractFilterer) (*StakeFilterer, error) {
	contract, err := bindStake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakeFilterer{contract: contract}, nil
}

// bindStake binds a generic wrapper to an already deployed contract.
func bindStake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stake *StakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stake.Contract.StakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stake *StakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.Contract.StakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stake *StakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stake.Contract.StakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stake *StakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stake *StakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stake *StakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stake.Contract.contract.Transact(opts, method, params...)
}

// BlackRemaingSeconds is a free data retrieval call binding the contract method 0xc933c516.
//
// Solidity: function blackRemaingSeconds(address player) view returns(uint256)
func (_Stake *StakeCaller) BlackRemaingSeconds(opts *bind.CallOpts, player common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "blackRemaingSeconds", player)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlackRemaingSeconds is a free data retrieval call binding the contract method 0xc933c516.
//
// Solidity: function blackRemaingSeconds(address player) view returns(uint256)
func (_Stake *StakeSession) BlackRemaingSeconds(player common.Address) (*big.Int, error) {
	return _Stake.Contract.BlackRemaingSeconds(&_Stake.CallOpts, player)
}

// BlackRemaingSeconds is a free data retrieval call binding the contract method 0xc933c516.
//
// Solidity: function blackRemaingSeconds(address player) view returns(uint256)
func (_Stake *StakeCallerSession) BlackRemaingSeconds(player common.Address) (*big.Int, error) {
	return _Stake.Contract.BlackRemaingSeconds(&_Stake.CallOpts, player)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Stake *StakeCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Stake *StakeSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Stake.Contract.OnERC721Received(&_Stake.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Stake *StakeCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Stake.Contract.OnERC721Received(&_Stake.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stake *StakeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stake *StakeSession) Owner() (common.Address, error) {
	return _Stake.Contract.Owner(&_Stake.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stake *StakeCallerSession) Owner() (common.Address, error) {
	return _Stake.Contract.Owner(&_Stake.CallOpts)
}

// RegisterData is a free data retrieval call binding the contract method 0xfe1ee17d.
//
// Solidity: function registerData(address nft, uint256 nftId) view returns(address, uint256)
func (_Stake *StakeCaller) RegisterData(opts *bind.CallOpts, nft common.Address, nftId *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "registerData", nft, nftId)

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
func (_Stake *StakeSession) RegisterData(nft common.Address, nftId *big.Int) (common.Address, *big.Int, error) {
	return _Stake.Contract.RegisterData(&_Stake.CallOpts, nft, nftId)
}

// RegisterData is a free data retrieval call binding the contract method 0xfe1ee17d.
//
// Solidity: function registerData(address nft, uint256 nftId) view returns(address, uint256)
func (_Stake *StakeCallerSession) RegisterData(nft common.Address, nftId *big.Int) (common.Address, *big.Int, error) {
	return _Stake.Contract.RegisterData(&_Stake.CallOpts, nft, nftId)
}

// RenounceOwnership is a free data retrieval call binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() view returns()
func (_Stake *StakeCaller) RenounceOwnership(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "renounceOwnership")

	if err != nil {
		return err
	}

	return err

}

// RenounceOwnership is a free data retrieval call binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() view returns()
func (_Stake *StakeSession) RenounceOwnership() error {
	return _Stake.Contract.RenounceOwnership(&_Stake.CallOpts)
}

// RenounceOwnership is a free data retrieval call binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() view returns()
func (_Stake *StakeCallerSession) RenounceOwnership() error {
	return _Stake.Contract.RenounceOwnership(&_Stake.CallOpts)
}

// StakEndts is a free data retrieval call binding the contract method 0x7576dbef.
//
// Solidity: function stakEndts(address nft, address player, uint256 nftId) view returns(uint256)
func (_Stake *StakeCaller) StakEndts(opts *bind.CallOpts, nft common.Address, player common.Address, nftId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "stakEndts", nft, player, nftId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakEndts is a free data retrieval call binding the contract method 0x7576dbef.
//
// Solidity: function stakEndts(address nft, address player, uint256 nftId) view returns(uint256)
func (_Stake *StakeSession) StakEndts(nft common.Address, player common.Address, nftId *big.Int) (*big.Int, error) {
	return _Stake.Contract.StakEndts(&_Stake.CallOpts, nft, player, nftId)
}

// StakEndts is a free data retrieval call binding the contract method 0x7576dbef.
//
// Solidity: function stakEndts(address nft, address player, uint256 nftId) view returns(uint256)
func (_Stake *StakeCallerSession) StakEndts(nft common.Address, player common.Address, nftId *big.Int) (*big.Int, error) {
	return _Stake.Contract.StakEndts(&_Stake.CallOpts, nft, player, nftId)
}

// Ban is a paid mutator transaction binding the contract method 0xd864598b.
//
// Solidity: function ban(address[] players, uint256 timestamp) returns()
func (_Stake *StakeTransactor) Ban(opts *bind.TransactOpts, players []common.Address, timestamp *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "ban", players, timestamp)
}

// Ban is a paid mutator transaction binding the contract method 0xd864598b.
//
// Solidity: function ban(address[] players, uint256 timestamp) returns()
func (_Stake *StakeSession) Ban(players []common.Address, timestamp *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.Ban(&_Stake.TransactOpts, players, timestamp)
}

// Ban is a paid mutator transaction binding the contract method 0xd864598b.
//
// Solidity: function ban(address[] players, uint256 timestamp) returns()
func (_Stake *StakeTransactorSession) Ban(players []common.Address, timestamp *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.Ban(&_Stake.TransactOpts, players, timestamp)
}

// Burn is a paid mutator transaction binding the contract method 0x091bb690.
//
// Solidity: function burn(address nft, uint32[] nftIds, uint64 timestamp, uint64 uuid, uint64 signId, bytes sig) returns()
func (_Stake *StakeTransactor) Burn(opts *bind.TransactOpts, nft common.Address, nftIds []uint32, timestamp uint64, uuid uint64, signId uint64, sig []byte) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "burn", nft, nftIds, timestamp, uuid, signId, sig)
}

// Burn is a paid mutator transaction binding the contract method 0x091bb690.
//
// Solidity: function burn(address nft, uint32[] nftIds, uint64 timestamp, uint64 uuid, uint64 signId, bytes sig) returns()
func (_Stake *StakeSession) Burn(nft common.Address, nftIds []uint32, timestamp uint64, uuid uint64, signId uint64, sig []byte) (*types.Transaction, error) {
	return _Stake.Contract.Burn(&_Stake.TransactOpts, nft, nftIds, timestamp, uuid, signId, sig)
}

// Burn is a paid mutator transaction binding the contract method 0x091bb690.
//
// Solidity: function burn(address nft, uint32[] nftIds, uint64 timestamp, uint64 uuid, uint64 signId, bytes sig) returns()
func (_Stake *StakeTransactorSession) Burn(nft common.Address, nftIds []uint32, timestamp uint64, uuid uint64, signId uint64, sig []byte) (*types.Transaction, error) {
	return _Stake.Contract.Burn(&_Stake.TransactOpts, nft, nftIds, timestamp, uuid, signId, sig)
}

// ReduceBanningDuration is a paid mutator transaction binding the contract method 0xc688951c.
//
// Solidity: function reduceBanningDuration(address[] players, uint256 timestamp) returns()
func (_Stake *StakeTransactor) ReduceBanningDuration(opts *bind.TransactOpts, players []common.Address, timestamp *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "reduceBanningDuration", players, timestamp)
}

// ReduceBanningDuration is a paid mutator transaction binding the contract method 0xc688951c.
//
// Solidity: function reduceBanningDuration(address[] players, uint256 timestamp) returns()
func (_Stake *StakeSession) ReduceBanningDuration(players []common.Address, timestamp *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.ReduceBanningDuration(&_Stake.TransactOpts, players, timestamp)
}

// ReduceBanningDuration is a paid mutator transaction binding the contract method 0xc688951c.
//
// Solidity: function reduceBanningDuration(address[] players, uint256 timestamp) returns()
func (_Stake *StakeTransactorSession) ReduceBanningDuration(players []common.Address, timestamp *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.ReduceBanningDuration(&_Stake.TransactOpts, players, timestamp)
}

// Registration is a paid mutator transaction binding the contract method 0x0d4017e0.
//
// Solidity: function registration(address nft, uint256[] nftIds, uint256 timestamp) returns()
func (_Stake *StakeTransactor) Registration(opts *bind.TransactOpts, nft common.Address, nftIds []*big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "registration", nft, nftIds, timestamp)
}

// Registration is a paid mutator transaction binding the contract method 0x0d4017e0.
//
// Solidity: function registration(address nft, uint256[] nftIds, uint256 timestamp) returns()
func (_Stake *StakeSession) Registration(nft common.Address, nftIds []*big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.Registration(&_Stake.TransactOpts, nft, nftIds, timestamp)
}

// Registration is a paid mutator transaction binding the contract method 0x0d4017e0.
//
// Solidity: function registration(address nft, uint256[] nftIds, uint256 timestamp) returns()
func (_Stake *StakeTransactorSession) Registration(nft common.Address, nftIds []*big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.Registration(&_Stake.TransactOpts, nft, nftIds, timestamp)
}

// RegistrationRenewal is a paid mutator transaction binding the contract method 0xa574f64a.
//
// Solidity: function registrationRenewal(address nft, uint256[] nftIds, uint64 timestamp) returns()
func (_Stake *StakeTransactor) RegistrationRenewal(opts *bind.TransactOpts, nft common.Address, nftIds []*big.Int, timestamp uint64) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "registrationRenewal", nft, nftIds, timestamp)
}

// RegistrationRenewal is a paid mutator transaction binding the contract method 0xa574f64a.
//
// Solidity: function registrationRenewal(address nft, uint256[] nftIds, uint64 timestamp) returns()
func (_Stake *StakeSession) RegistrationRenewal(nft common.Address, nftIds []*big.Int, timestamp uint64) (*types.Transaction, error) {
	return _Stake.Contract.RegistrationRenewal(&_Stake.TransactOpts, nft, nftIds, timestamp)
}

// RegistrationRenewal is a paid mutator transaction binding the contract method 0xa574f64a.
//
// Solidity: function registrationRenewal(address nft, uint256[] nftIds, uint64 timestamp) returns()
func (_Stake *StakeTransactorSession) RegistrationRenewal(nft common.Address, nftIds []*big.Int, timestamp uint64) (*types.Transaction, error) {
	return _Stake.Contract.RegistrationRenewal(&_Stake.TransactOpts, nft, nftIds, timestamp)
}

// SetAddrSwitch is a paid mutator transaction binding the contract method 0x33cd673a.
//
// Solidity: function setAddrSwitch(address addr, uint256 duration) returns()
func (_Stake *StakeTransactor) SetAddrSwitch(opts *bind.TransactOpts, addr common.Address, duration *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "setAddrSwitch", addr, duration)
}

// SetAddrSwitch is a paid mutator transaction binding the contract method 0x33cd673a.
//
// Solidity: function setAddrSwitch(address addr, uint256 duration) returns()
func (_Stake *StakeSession) SetAddrSwitch(addr common.Address, duration *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.SetAddrSwitch(&_Stake.TransactOpts, addr, duration)
}

// SetAddrSwitch is a paid mutator transaction binding the contract method 0x33cd673a.
//
// Solidity: function setAddrSwitch(address addr, uint256 duration) returns()
func (_Stake *StakeTransactorSession) SetAddrSwitch(addr common.Address, duration *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.SetAddrSwitch(&_Stake.TransactOpts, addr, duration)
}

// SetCosigner is a paid mutator transaction binding the contract method 0x02045138.
//
// Solidity: function setCosigner(address cosigner) returns()
func (_Stake *StakeTransactor) SetCosigner(opts *bind.TransactOpts, cosigner common.Address) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "setCosigner", cosigner)
}

// SetCosigner is a paid mutator transaction binding the contract method 0x02045138.
//
// Solidity: function setCosigner(address cosigner) returns()
func (_Stake *StakeSession) SetCosigner(cosigner common.Address) (*types.Transaction, error) {
	return _Stake.Contract.SetCosigner(&_Stake.TransactOpts, cosigner)
}

// SetCosigner is a paid mutator transaction binding the contract method 0x02045138.
//
// Solidity: function setCosigner(address cosigner) returns()
func (_Stake *StakeTransactorSession) SetCosigner(cosigner common.Address) (*types.Transaction, error) {
	return _Stake.Contract.SetCosigner(&_Stake.TransactOpts, cosigner)
}

// SetGloSwitch is a paid mutator transaction binding the contract method 0x063e9f4d.
//
// Solidity: function setGloSwitch(uint256 duration) returns()
func (_Stake *StakeTransactor) SetGloSwitch(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "setGloSwitch", duration)
}

// SetGloSwitch is a paid mutator transaction binding the contract method 0x063e9f4d.
//
// Solidity: function setGloSwitch(uint256 duration) returns()
func (_Stake *StakeSession) SetGloSwitch(duration *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.SetGloSwitch(&_Stake.TransactOpts, duration)
}

// SetGloSwitch is a paid mutator transaction binding the contract method 0x063e9f4d.
//
// Solidity: function setGloSwitch(uint256 duration) returns()
func (_Stake *StakeTransactorSession) SetGloSwitch(duration *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.SetGloSwitch(&_Stake.TransactOpts, duration)
}

// SetLimitLen is a paid mutator transaction binding the contract method 0x9a12cd75.
//
// Solidity: function setLimitLen(uint32 _limitIdLen) returns()
func (_Stake *StakeTransactor) SetLimitLen(opts *bind.TransactOpts, _limitIdLen uint32) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "setLimitLen", _limitIdLen)
}

// SetLimitLen is a paid mutator transaction binding the contract method 0x9a12cd75.
//
// Solidity: function setLimitLen(uint32 _limitIdLen) returns()
func (_Stake *StakeSession) SetLimitLen(_limitIdLen uint32) (*types.Transaction, error) {
	return _Stake.Contract.SetLimitLen(&_Stake.TransactOpts, _limitIdLen)
}

// SetLimitLen is a paid mutator transaction binding the contract method 0x9a12cd75.
//
// Solidity: function setLimitLen(uint32 _limitIdLen) returns()
func (_Stake *StakeTransactorSession) SetLimitLen(_limitIdLen uint32) (*types.Transaction, error) {
	return _Stake.Contract.SetLimitLen(&_Stake.TransactOpts, _limitIdLen)
}

// SetNFTSwitch is a paid mutator transaction binding the contract method 0xaae7d433.
//
// Solidity: function setNFTSwitch(address nft, uint256 nftId, uint256 duration) returns()
func (_Stake *StakeTransactor) SetNFTSwitch(opts *bind.TransactOpts, nft common.Address, nftId *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "setNFTSwitch", nft, nftId, duration)
}

// SetNFTSwitch is a paid mutator transaction binding the contract method 0xaae7d433.
//
// Solidity: function setNFTSwitch(address nft, uint256 nftId, uint256 duration) returns()
func (_Stake *StakeSession) SetNFTSwitch(nft common.Address, nftId *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.SetNFTSwitch(&_Stake.TransactOpts, nft, nftId, duration)
}

// SetNFTSwitch is a paid mutator transaction binding the contract method 0xaae7d433.
//
// Solidity: function setNFTSwitch(address nft, uint256 nftId, uint256 duration) returns()
func (_Stake *StakeTransactorSession) SetNFTSwitch(nft common.Address, nftId *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.SetNFTSwitch(&_Stake.TransactOpts, nft, nftId, duration)
}

// SetSupportNFT is a paid mutator transaction binding the contract method 0x9a4a9b39.
//
// Solidity: function setSupportNFT(address[] nfts, bool status) returns()
func (_Stake *StakeTransactor) SetSupportNFT(opts *bind.TransactOpts, nfts []common.Address, status bool) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "setSupportNFT", nfts, status)
}

// SetSupportNFT is a paid mutator transaction binding the contract method 0x9a4a9b39.
//
// Solidity: function setSupportNFT(address[] nfts, bool status) returns()
func (_Stake *StakeSession) SetSupportNFT(nfts []common.Address, status bool) (*types.Transaction, error) {
	return _Stake.Contract.SetSupportNFT(&_Stake.TransactOpts, nfts, status)
}

// SetSupportNFT is a paid mutator transaction binding the contract method 0x9a4a9b39.
//
// Solidity: function setSupportNFT(address[] nfts, bool status) returns()
func (_Stake *StakeTransactorSession) SetSupportNFT(nfts []common.Address, status bool) (*types.Transaction, error) {
	return _Stake.Contract.SetSupportNFT(&_Stake.TransactOpts, nfts, status)
}

// SetTimestampExpirySeconds is a paid mutator transaction binding the contract method 0xd65c9d2f.
//
// Solidity: function setTimestampExpirySeconds(uint32 expiry) returns()
func (_Stake *StakeTransactor) SetTimestampExpirySeconds(opts *bind.TransactOpts, expiry uint32) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "setTimestampExpirySeconds", expiry)
}

// SetTimestampExpirySeconds is a paid mutator transaction binding the contract method 0xd65c9d2f.
//
// Solidity: function setTimestampExpirySeconds(uint32 expiry) returns()
func (_Stake *StakeSession) SetTimestampExpirySeconds(expiry uint32) (*types.Transaction, error) {
	return _Stake.Contract.SetTimestampExpirySeconds(&_Stake.TransactOpts, expiry)
}

// SetTimestampExpirySeconds is a paid mutator transaction binding the contract method 0xd65c9d2f.
//
// Solidity: function setTimestampExpirySeconds(uint32 expiry) returns()
func (_Stake *StakeTransactorSession) SetTimestampExpirySeconds(expiry uint32) (*types.Transaction, error) {
	return _Stake.Contract.SetTimestampExpirySeconds(&_Stake.TransactOpts, expiry)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stake *StakeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stake *StakeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Stake.Contract.TransferOwnership(&_Stake.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stake *StakeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Stake.Contract.TransferOwnership(&_Stake.TransactOpts, newOwner)
}

// Unregistration is a paid mutator transaction binding the contract method 0x17e8a140.
//
// Solidity: function unregistration(address nft, uint256[] nftIds) returns()
func (_Stake *StakeTransactor) Unregistration(opts *bind.TransactOpts, nft common.Address, nftIds []*big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "unregistration", nft, nftIds)
}

// Unregistration is a paid mutator transaction binding the contract method 0x17e8a140.
//
// Solidity: function unregistration(address nft, uint256[] nftIds) returns()
func (_Stake *StakeSession) Unregistration(nft common.Address, nftIds []*big.Int) (*types.Transaction, error) {
	return _Stake.Contract.Unregistration(&_Stake.TransactOpts, nft, nftIds)
}

// Unregistration is a paid mutator transaction binding the contract method 0x17e8a140.
//
// Solidity: function unregistration(address nft, uint256[] nftIds) returns()
func (_Stake *StakeTransactorSession) Unregistration(nft common.Address, nftIds []*big.Int) (*types.Transaction, error) {
	return _Stake.Contract.Unregistration(&_Stake.TransactOpts, nft, nftIds)
}

// UpdateMiner is a paid mutator transaction binding the contract method 0x2b16cad2.
//
// Solidity: function updateMiner(address newminer) returns()
func (_Stake *StakeTransactor) UpdateMiner(opts *bind.TransactOpts, newminer common.Address) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "updateMiner", newminer)
}

// UpdateMiner is a paid mutator transaction binding the contract method 0x2b16cad2.
//
// Solidity: function updateMiner(address newminer) returns()
func (_Stake *StakeSession) UpdateMiner(newminer common.Address) (*types.Transaction, error) {
	return _Stake.Contract.UpdateMiner(&_Stake.TransactOpts, newminer)
}

// UpdateMiner is a paid mutator transaction binding the contract method 0x2b16cad2.
//
// Solidity: function updateMiner(address newminer) returns()
func (_Stake *StakeTransactorSession) UpdateMiner(newminer common.Address) (*types.Transaction, error) {
	return _Stake.Contract.UpdateMiner(&_Stake.TransactOpts, newminer)
}

// UpdateTsLimit is a paid mutator transaction binding the contract method 0x1cf06bbc.
//
// Solidity: function updateTsLimit(uint64 floorts, uint64 ceilts) returns()
func (_Stake *StakeTransactor) UpdateTsLimit(opts *bind.TransactOpts, floorts uint64, ceilts uint64) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "updateTsLimit", floorts, ceilts)
}

// UpdateTsLimit is a paid mutator transaction binding the contract method 0x1cf06bbc.
//
// Solidity: function updateTsLimit(uint64 floorts, uint64 ceilts) returns()
func (_Stake *StakeSession) UpdateTsLimit(floorts uint64, ceilts uint64) (*types.Transaction, error) {
	return _Stake.Contract.UpdateTsLimit(&_Stake.TransactOpts, floorts, ceilts)
}

// UpdateTsLimit is a paid mutator transaction binding the contract method 0x1cf06bbc.
//
// Solidity: function updateTsLimit(uint64 floorts, uint64 ceilts) returns()
func (_Stake *StakeTransactorSession) UpdateTsLimit(floorts uint64, ceilts uint64) (*types.Transaction, error) {
	return _Stake.Contract.UpdateTsLimit(&_Stake.TransactOpts, floorts, ceilts)
}

// StakeBanIterator is returned from FilterBan and is used to iterate over the raw logs and unpacked data for Ban events raised by the Stake contract.
type StakeBanIterator struct {
	Event *StakeBan // Event containing the contract specifics and raw log

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
func (it *StakeBanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeBan)
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
		it.Event = new(StakeBan)
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
func (it *StakeBanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeBanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeBan represents a Ban event raised by the Stake contract.
type StakeBan struct {
	Timestamp *big.Int
	Players   []common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBan is a free log retrieval operation binding the contract event 0x8b76a9674748eed60af451b49a3934d77641c61747dd41c2910ec3d69ddfca8d.
//
// Solidity: event Ban(uint256 timestamp, address[] players)
func (_Stake *StakeFilterer) FilterBan(opts *bind.FilterOpts) (*StakeBanIterator, error) {

	logs, sub, err := _Stake.contract.FilterLogs(opts, "Ban")
	if err != nil {
		return nil, err
	}
	return &StakeBanIterator{contract: _Stake.contract, event: "Ban", logs: logs, sub: sub}, nil
}

// WatchBan is a free log subscription operation binding the contract event 0x8b76a9674748eed60af451b49a3934d77641c61747dd41c2910ec3d69ddfca8d.
//
// Solidity: event Ban(uint256 timestamp, address[] players)
func (_Stake *StakeFilterer) WatchBan(opts *bind.WatchOpts, sink chan<- *StakeBan) (event.Subscription, error) {

	logs, sub, err := _Stake.contract.WatchLogs(opts, "Ban")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeBan)
				if err := _Stake.contract.UnpackLog(event, "Ban", log); err != nil {
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
func (_Stake *StakeFilterer) ParseBan(log types.Log) (*StakeBan, error) {
	event := new(StakeBan)
	if err := _Stake.contract.UnpackLog(event, "Ban", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Stake contract.
type StakeBurnIterator struct {
	Event *StakeBurn // Event containing the contract specifics and raw log

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
func (it *StakeBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeBurn)
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
		it.Event = new(StakeBurn)
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
func (it *StakeBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeBurn represents a Burn event raised by the Stake contract.
type StakeBurn struct {
	SignId uint64
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xe1b7f9aa58e45fef132f6fb1603c9c4de00def2bfcb5f183cebd2d5988994fa2.
//
// Solidity: event Burn(uint64 signId)
func (_Stake *StakeFilterer) FilterBurn(opts *bind.FilterOpts) (*StakeBurnIterator, error) {

	logs, sub, err := _Stake.contract.FilterLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return &StakeBurnIterator{contract: _Stake.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xe1b7f9aa58e45fef132f6fb1603c9c4de00def2bfcb5f183cebd2d5988994fa2.
//
// Solidity: event Burn(uint64 signId)
func (_Stake *StakeFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *StakeBurn) (event.Subscription, error) {

	logs, sub, err := _Stake.contract.WatchLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeBurn)
				if err := _Stake.contract.UnpackLog(event, "Burn", log); err != nil {
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
func (_Stake *StakeFilterer) ParseBurn(log types.Log) (*StakeBurn, error) {
	event := new(StakeBurn)
	if err := _Stake.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Stake contract.
type StakeOwnershipTransferredIterator struct {
	Event *StakeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeOwnershipTransferred)
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
		it.Event = new(StakeOwnershipTransferred)
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
func (it *StakeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeOwnershipTransferred represents a OwnershipTransferred event raised by the Stake contract.
type StakeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Stake *StakeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakeOwnershipTransferredIterator{contract: _Stake.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Stake *StakeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeOwnershipTransferred)
				if err := _Stake.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Stake *StakeFilterer) ParseOwnershipTransferred(log types.Log) (*StakeOwnershipTransferred, error) {
	event := new(StakeOwnershipTransferred)
	if err := _Stake.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeReduceBanningDurationIterator is returned from FilterReduceBanningDuration and is used to iterate over the raw logs and unpacked data for ReduceBanningDuration events raised by the Stake contract.
type StakeReduceBanningDurationIterator struct {
	Event *StakeReduceBanningDuration // Event containing the contract specifics and raw log

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
func (it *StakeReduceBanningDurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeReduceBanningDuration)
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
		it.Event = new(StakeReduceBanningDuration)
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
func (it *StakeReduceBanningDurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeReduceBanningDurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeReduceBanningDuration represents a ReduceBanningDuration event raised by the Stake contract.
type StakeReduceBanningDuration struct {
	Player    common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReduceBanningDuration is a free log retrieval operation binding the contract event 0x36f824b722a6b1ea67b79796d58020cf63f27b9a542acdd0316d01af365d5504.
//
// Solidity: event ReduceBanningDuration(address player, uint256 timestamp)
func (_Stake *StakeFilterer) FilterReduceBanningDuration(opts *bind.FilterOpts) (*StakeReduceBanningDurationIterator, error) {

	logs, sub, err := _Stake.contract.FilterLogs(opts, "ReduceBanningDuration")
	if err != nil {
		return nil, err
	}
	return &StakeReduceBanningDurationIterator{contract: _Stake.contract, event: "ReduceBanningDuration", logs: logs, sub: sub}, nil
}

// WatchReduceBanningDuration is a free log subscription operation binding the contract event 0x36f824b722a6b1ea67b79796d58020cf63f27b9a542acdd0316d01af365d5504.
//
// Solidity: event ReduceBanningDuration(address player, uint256 timestamp)
func (_Stake *StakeFilterer) WatchReduceBanningDuration(opts *bind.WatchOpts, sink chan<- *StakeReduceBanningDuration) (event.Subscription, error) {

	logs, sub, err := _Stake.contract.WatchLogs(opts, "ReduceBanningDuration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeReduceBanningDuration)
				if err := _Stake.contract.UnpackLog(event, "ReduceBanningDuration", log); err != nil {
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
func (_Stake *StakeFilterer) ParseReduceBanningDuration(log types.Log) (*StakeReduceBanningDuration, error) {
	event := new(StakeReduceBanningDuration)
	if err := _Stake.contract.UnpackLog(event, "ReduceBanningDuration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeRegisterIterator is returned from FilterRegister and is used to iterate over the raw logs and unpacked data for Register events raised by the Stake contract.
type StakeRegisterIterator struct {
	Event *StakeRegister // Event containing the contract specifics and raw log

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
func (it *StakeRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeRegister)
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
		it.Event = new(StakeRegister)
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
func (it *StakeRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeRegister represents a Register event raised by the Stake contract.
type StakeRegister struct {
	Nft       common.Address
	Register  common.Address
	Timestamp *big.Int
	NftId     []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRegister is a free log retrieval operation binding the contract event 0x5ad8141c164356bdef9e16f08312a7034ac6682a7413ce4fecfc44da5e18fec7.
//
// Solidity: event Register(address nft, address register, uint256 timestamp, uint256[] nftId)
func (_Stake *StakeFilterer) FilterRegister(opts *bind.FilterOpts) (*StakeRegisterIterator, error) {

	logs, sub, err := _Stake.contract.FilterLogs(opts, "Register")
	if err != nil {
		return nil, err
	}
	return &StakeRegisterIterator{contract: _Stake.contract, event: "Register", logs: logs, sub: sub}, nil
}

// WatchRegister is a free log subscription operation binding the contract event 0x5ad8141c164356bdef9e16f08312a7034ac6682a7413ce4fecfc44da5e18fec7.
//
// Solidity: event Register(address nft, address register, uint256 timestamp, uint256[] nftId)
func (_Stake *StakeFilterer) WatchRegister(opts *bind.WatchOpts, sink chan<- *StakeRegister) (event.Subscription, error) {

	logs, sub, err := _Stake.contract.WatchLogs(opts, "Register")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeRegister)
				if err := _Stake.contract.UnpackLog(event, "Register", log); err != nil {
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
func (_Stake *StakeFilterer) ParseRegister(log types.Log) (*StakeRegister, error) {
	event := new(StakeRegister)
	if err := _Stake.contract.UnpackLog(event, "Register", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeRenewalIterator is returned from FilterRenewal and is used to iterate over the raw logs and unpacked data for Renewal events raised by the Stake contract.
type StakeRenewalIterator struct {
	Event *StakeRenewal // Event containing the contract specifics and raw log

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
func (it *StakeRenewalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeRenewal)
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
		it.Event = new(StakeRenewal)
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
func (it *StakeRenewalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeRenewalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeRenewal represents a Renewal event raised by the Stake contract.
type StakeRenewal struct {
	Nft       common.Address
	Register  common.Address
	NftId     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRenewal is a free log retrieval operation binding the contract event 0x7c9bb2d7d3438e95c0602e298a3af1e36d7522a0b5e21e5cc87458e91dea1202.
//
// Solidity: event Renewal(address nft, address register, uint256 nftId, uint256 timestamp)
func (_Stake *StakeFilterer) FilterRenewal(opts *bind.FilterOpts) (*StakeRenewalIterator, error) {

	logs, sub, err := _Stake.contract.FilterLogs(opts, "Renewal")
	if err != nil {
		return nil, err
	}
	return &StakeRenewalIterator{contract: _Stake.contract, event: "Renewal", logs: logs, sub: sub}, nil
}

// WatchRenewal is a free log subscription operation binding the contract event 0x7c9bb2d7d3438e95c0602e298a3af1e36d7522a0b5e21e5cc87458e91dea1202.
//
// Solidity: event Renewal(address nft, address register, uint256 nftId, uint256 timestamp)
func (_Stake *StakeFilterer) WatchRenewal(opts *bind.WatchOpts, sink chan<- *StakeRenewal) (event.Subscription, error) {

	logs, sub, err := _Stake.contract.WatchLogs(opts, "Renewal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeRenewal)
				if err := _Stake.contract.UnpackLog(event, "Renewal", log); err != nil {
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
func (_Stake *StakeFilterer) ParseRenewal(log types.Log) (*StakeRenewal, error) {
	event := new(StakeRenewal)
	if err := _Stake.contract.UnpackLog(event, "Renewal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeUnregisterIterator is returned from FilterUnregister and is used to iterate over the raw logs and unpacked data for Unregister events raised by the Stake contract.
type StakeUnregisterIterator struct {
	Event *StakeUnregister // Event containing the contract specifics and raw log

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
func (it *StakeUnregisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeUnregister)
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
		it.Event = new(StakeUnregister)
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
func (it *StakeUnregisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeUnregisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeUnregister represents a Unregister event raised by the Stake contract.
type StakeUnregister struct {
	Nft      common.Address
	Register common.Address
	NftId    []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUnregister is a free log retrieval operation binding the contract event 0xeb879c9d6d39266b9caad39ced3788f8b8f47bb316e3fb55f3f44cb0f638cbc6.
//
// Solidity: event Unregister(address nft, address register, uint256[] nftId)
func (_Stake *StakeFilterer) FilterUnregister(opts *bind.FilterOpts) (*StakeUnregisterIterator, error) {

	logs, sub, err := _Stake.contract.FilterLogs(opts, "Unregister")
	if err != nil {
		return nil, err
	}
	return &StakeUnregisterIterator{contract: _Stake.contract, event: "Unregister", logs: logs, sub: sub}, nil
}

// WatchUnregister is a free log subscription operation binding the contract event 0xeb879c9d6d39266b9caad39ced3788f8b8f47bb316e3fb55f3f44cb0f638cbc6.
//
// Solidity: event Unregister(address nft, address register, uint256[] nftId)
func (_Stake *StakeFilterer) WatchUnregister(opts *bind.WatchOpts, sink chan<- *StakeUnregister) (event.Subscription, error) {

	logs, sub, err := _Stake.contract.WatchLogs(opts, "Unregister")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeUnregister)
				if err := _Stake.contract.UnpackLog(event, "Unregister", log); err != nil {
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
func (_Stake *StakeFilterer) ParseUnregister(log types.Log) (*StakeUnregister, error) {
	event := new(StakeUnregister)
	if err := _Stake.contract.UnpackLog(event, "Unregister", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
