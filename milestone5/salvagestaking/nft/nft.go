// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nft

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

// NftMetaData contains all meta data concerning the Nft contract.
var NftMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceQueryForZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintERC2309QuantityExceedsLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCompatibleWithSpotMints\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotMintable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnershipNotInitializedForExtraData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequentialMintExceedsLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequentialUpToTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SpotMintTokenIdTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"URIQueryForNonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ConsecutiveTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"caller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentBaseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_tokenURISuffix\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_currentBaseURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initailize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"qty\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"qty\",\"type\":\"uint64\"}],\"name\":\"ownermint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxsupply\",\"type\":\"uint256\"}],\"name\":\"setMaxsupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_mintable\",\"type\":\"bool\"}],\"name\":\"setMintable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"suffix\",\"type\":\"string\"}],\"name\":\"setTokenURISuffix\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"updateCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052600436106101d85760003560e01c8063715018a611610102578063a9852bfb11610095578063d5abeb0111610064578063d5abeb0114610509578063e985e9c51461051f578063f2fde38b14610568578063fc9c8d391461058857600080fd5b8063a9852bfb14610496578063b88d4fde146104b6578063bde17fc0146104c9578063c87b56dd146104e957600080fd5b806397d4ccd2116100d157806397d4ccd21461042c5780639b9df99b14610441578063a22cb46514610461578063a2309ff81461048157600080fd5b8063715018a6146103c45780638da5cb5b146103d957806394b44725146103f757806395d89b411461041757600080fd5b806323b872dd1161017a5780634bf365df116101495780634bf365df1461034a57806355f804b3146103645780636352211e1461038457806370a08231146103a457600080fd5b806323b872dd146102e4578063285d70d4146102f757806342842e0e1461031757806342966c681461032a57600080fd5b8063095ea7b3116101b6578063095ea7b31461026c578063149835a01461028157806317deeba0146102a157806318160ddd146102c157600080fd5b806301ffc9a7146101dd57806306fdde0314610212578063081812fc14610234575b600080fd5b3480156101e957600080fd5b506101fd6101f83660046114d6565b6105a6565b60405190151581526020015b60405180910390f35b34801561021e57600080fd5b506102276105f8565b6040516102099190611543565b34801561024057600080fd5b5061025461024f366004611556565b61068a565b6040516001600160a01b039091168152602001610209565b61027f61027a366004611586565b6106c5565b005b34801561028d57600080fd5b5061027f61029c366004611556565b6106d5565b3480156102ad57600080fd5b506101fd6102bc3660046115c8565b6106e2565b3480156102cd57600080fd5b506102d6610722565b604051908152602001610209565b61027f6102f23660046115fb565b61073b565b34801561030357600080fd5b5061027f610312366004611647565b6108aa565b61027f6103253660046115fb565b6108c5565b34801561033657600080fd5b5061027f610345366004611556565b6108e5565b34801561035657600080fd5b50600a546101fd9060ff1681565b34801561037057600080fd5b5061027f61037f366004611662565b610962565b34801561039057600080fd5b5061025461039f366004611556565b610977565b3480156103b057600080fd5b506102d66103bf3660046116d4565b610982565b3480156103d057600080fd5b5061027f6109c8565b3480156103e557600080fd5b506000546001600160a01b0316610254565b34801561040357600080fd5b5061027f6104123660046116ef565b610a0b565b34801561042357600080fd5b50610227610a1d565b34801561043857600080fd5b50610227610a2c565b34801561044d57600080fd5b5061027f61045c3660046116d4565b610aba565b34801561046d57600080fd5b5061027f61047c366004611719565b610ae4565b34801561048d57600080fd5b506102d6610b50565b3480156104a257600080fd5b5061027f6104b1366004611662565b610b5f565b61027f6104c43660046117cf565b610b74565b3480156104d557600080fd5b5061027f6104e436600461186b565b610bb5565b3480156104f557600080fd5b50610227610504366004611556565b610c69565b34801561051557600080fd5b506102d6600b5481565b34801561052b57600080fd5b506101fd61053a366004611943565b6001600160a01b03918216600090815260086020908152604080832093909416825291909152205460ff1690565b34801561057457600080fd5b5061027f6105833660046116d4565b610d73565b34801561059457600080fd5b50600d546001600160a01b0316610254565b60006301ffc9a760e01b6001600160e01b0319831614806105d757506380ac58cd60e01b6001600160e01b03198316145b806105f25750635b5e139f60e01b6001600160e01b03198316145b92915050565b6060601080546106079061195f565b80601f01602080910402602001604051908101604052809291908181526020018280546106339061195f565b80156106805780601f1061065557610100808354040283529160200191610680565b820191906000526020600020905b81548152906001019060200180831161066357829003601f168201915b5050505050905090565b600061069582610de9565b6106a9576106a96333d1c03960e21b610e2f565b506000908152600760205260409020546001600160a01b031690565b6106d182826001610e39565b5050565b6106dd610edc565b600b55565b60006106ec610f38565b600a5460ff1661070f57604051630952c8a960e11b815260040160405180910390fd5b6107198383610fa3565b50600192915050565b60025460015403600019805b1461073857600954015b90565b60006107468261101e565b6001600160a01b03948516949091508116841461076c5761076c62a1148160e81b610e2f565b600082815260076020526040902080546107988187335b6001600160a01b039081169116811491141790565b6107ba576107a6863361053a565b6107ba576107ba632ce44b5f60e11b610e2f565b80156107c557600082555b6001600160a01b038681166000908152600660205260408082208054600019019055918716808252919020805460010190554260a01b17600160e11b17600085815260056020526040812091909155600160e11b84169003610857576001840160008181526005602052604081205490036108555760015481146108555760008181526005602052604090208490555b505b6001600160a01b0385168481887fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef600080a4806000036108a1576108a1633a954ecd60e21b610e2f565b50505050505050565b6108b2610edc565b600a805460ff1916911515919091179055565b6108e083838360405180602001604052806000815250610b74565b505050565b336108ef82610977565b6001600160a01b0316146109565760405162461bcd60e51b8152602060048201526024808201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865204e4654206f6044820152633bb732b960e11b60648201526084015b60405180910390fd5b61095f816110b4565b50565b61096a610edc565b600e6108e08284836119e7565b60006105f28261101e565b60006001600160a01b0382166109a2576109a26323d3ad8160e21b610e2f565b506001600160a01b031660009081526006602052604090205467ffffffffffffffff1690565b6109d0610edc565b60405162461bcd60e51b815260206004820152601060248201526f434c4f5345445f494e5445524641434560801b604482015260640161094d565b610a13610edc565b6106d18183610fa3565b6060601180546106079061195f565b600e8054610a399061195f565b80601f0160208091040260200160405190810160405280929190818152602001828054610a659061195f565b8015610ab25780601f10610a8757610100808354040283529160200191610ab2565b820191906000526020600020905b815481529060010190602001808311610a9557829003601f168201915b505050505081565b610ac2610edc565b600d80546001600160a01b0319166001600160a01b0392909216919091179055565b3360008181526008602090815260408083206001600160a01b03871680855290835292819020805460ff191686151590811790915590519081529192917f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a35050565b6000610b5a6110bf565b905090565b610b67610edc565b600f6108e08284836119e7565b610b7f84848461073b565b6001600160a01b0383163b15610baf57610b9b848484846110cb565b610baf57610baf6368d2bf6b60e11b610e2f565b50505050565b600c5415610bfb5760405162461bcd60e51b81526020600482015260136024820152721053149150511657d253925512505312569151606a1b604482015260640161094d565b6010610c078882611aa7565b506011610c148782611aa7565b50610c1e816111ad565b600f610c2a8682611aa7565b50600e610c378582611aa7565b5050600b91909155600d80546001600160a01b0319166001600160a01b0390921691909117905550506001600c555050565b6060610c7482610de9565b610c9157604051630a14c4b560e41b815260040160405180910390fd5b6000600e8054610ca09061195f565b80601f0160208091040260200160405190810160405280929190818152602001828054610ccc9061195f565b8015610d195780601f10610cee57610100808354040283529160200191610d19565b820191906000526020600020905b815481529060010190602001808311610cfc57829003601f168201915b505050505090508051600003610d3e5760405180602001604052806000815250610d6c565b80610d48846111fd565b600f604051602001610d5c93929190611b67565b6040516020818303038152906040525b9392505050565b610d7b610edc565b6001600160a01b038116610de05760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161094d565b61095f816111ad565b6000600154821015610e2a5760005b5060008281526005602052604081205490819003610e2057610e1983611c1d565b9250610df8565b600160e01b161590505b919050565b8060005260046000fd5b6000610e4483610977565b9050818015610e5c5750336001600160a01b03821614155b15610e7f57610e6b813361053a565b610e7f57610e7f6367d9dca160e11b610e2f565b60008381526007602052604080822080546001600160a01b0319166001600160a01b0388811691821790925591518693918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a450505050565b6000546001600160a01b03163314610f365760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161094d565b565b600d546001600160a01b03163314610f365760405162461bcd60e51b815260206004820152602860248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206d696e742060448201526718dbdb9d1c9858dd60c21b606482015260840161094d565b600b548267ffffffffffffffff16610fb9610b50565b610fc39190611c34565b111561100a5760405162461bcd60e51b815260206004820152601660248201527545786365656473206d6178696d756d20737570706c7960501b604482015260640161094d565b6106d1818367ffffffffffffffff16611241565b6000818152600560205260409020548060000361109157600154821061104e5761104e636f96cda160e11b610e2f565b5b5060001901600081815260056020526040902054801561104f57600160e01b811660000361107c57919050565b61108c636f96cda160e11b610e2f565b61104f565b600160e01b81166000036110a457919050565b610e2a636f96cda160e11b610e2f565b61095f81600161125b565b6001546000198061072e565b604051630a85bd0160e11b81526000906001600160a01b0385169063150b7a0290611100903390899088908890600401611c47565b6020604051808303816000875af192505050801561113b575060408051601f3d908101601f1916820190925261113891810190611c84565b60015b611190573d808015611169576040519150601f19603f3d011682016040523d82523d6000602084013e61116e565b606091505b508051600003611188576111886368d2bf6b60e11b610e2f565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050949350505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b606060a06040510180604052602081039150506000815280825b600183039250600a81066030018353600a9004806112175750819003601f19909101908152919050565b6106d182826040518060200160405280600081525061139d565b60006112668361101e565b90508060008061128486600090815260076020526040902080549091565b9150915084156112bb57611299818433610783565b6112bb576112a7833361053a565b6112bb576112bb632ce44b5f60e11b610e2f565b80156112c657600082555b6001600160a01b038316600081815260066020526040902080546fffffffffffffffffffffffffffffffff0190554260a01b17600360e01b17600087815260056020526040812091909155600160e11b85169003611354576001860160008181526005602052604081205490036113525760015481146113525760008181526005602052604090208590555b505b60405186906000906001600160a01b038616907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a4505060028054600101905550505050565b6113a78383611401565b6001600160a01b0383163b156108e0576001548281035b6113d160008683806001019450866110cb565b6113e5576113e56368d2bf6b60e11b610e2f565b8181106113be5781600154146113fa57600080fd5b5050505050565b600154600082900361141d5761141d63b562e8dd60e01b610e2f565b60008181526005602090815260408083206001600160a01b0387164260a01b6001881460e11b1781179091558084526006909252822080546801000000000000000186020190559081900361147b5761147b622e076360e81b610e2f565b818301825b808360007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef600080a4818160010191508103611480575060015550505050565b6001600160e01b03198116811461095f57600080fd5b6000602082840312156114e857600080fd5b8135610d6c816114c0565b60005b8381101561150e5781810151838201526020016114f6565b50506000910152565b6000815180845261152f8160208601602086016114f3565b601f01601f19169290920160200192915050565b602081526000610d6c6020830184611517565b60006020828403121561156857600080fd5b5035919050565b80356001600160a01b0381168114610e2a57600080fd5b6000806040838503121561159957600080fd5b6115a28361156f565b946020939093013593505050565b803567ffffffffffffffff81168114610e2a57600080fd5b600080604083850312156115db57600080fd5b6115e4836115b0565b91506115f26020840161156f565b90509250929050565b60008060006060848603121561161057600080fd5b6116198461156f565b92506116276020850161156f565b9150604084013590509250925092565b80358015158114610e2a57600080fd5b60006020828403121561165957600080fd5b610d6c82611637565b6000806020838503121561167557600080fd5b823567ffffffffffffffff8082111561168d57600080fd5b818501915085601f8301126116a157600080fd5b8135818111156116b057600080fd5b8660208285010111156116c257600080fd5b60209290920196919550909350505050565b6000602082840312156116e657600080fd5b610d6c8261156f565b6000806040838503121561170257600080fd5b61170b8361156f565b91506115f2602084016115b0565b6000806040838503121561172c57600080fd5b6117358361156f565b91506115f260208401611637565b634e487b7160e01b600052604160045260246000fd5b600067ffffffffffffffff8084111561177457611774611743565b604051601f8501601f19908116603f0116810190828211818310171561179c5761179c611743565b816040528093508581528686860111156117b557600080fd5b858560208301376000602087830101525050509392505050565b600080600080608085870312156117e557600080fd5b6117ee8561156f565b93506117fc6020860161156f565b925060408501359150606085013567ffffffffffffffff81111561181f57600080fd5b8501601f8101871361183057600080fd5b61183f87823560208401611759565b91505092959194509250565b600082601f83011261185c57600080fd5b610d6c83833560208501611759565b600080600080600080600060e0888a03121561188657600080fd5b873567ffffffffffffffff8082111561189e57600080fd5b6118aa8b838c0161184b565b985060208a01359150808211156118c057600080fd5b6118cc8b838c0161184b565b975060408a01359150808211156118e257600080fd5b6118ee8b838c0161184b565b965060608a013591508082111561190457600080fd5b506119118a828b0161184b565b9450506080880135925061192760a0890161156f565b915061193560c0890161156f565b905092959891949750929550565b6000806040838503121561195657600080fd5b6115e48361156f565b600181811c9082168061197357607f821691505b60208210810361199357634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156108e057600081815260208120601f850160051c810160208610156119c05750805b601f850160051c820191505b818110156119df578281556001016119cc565b505050505050565b67ffffffffffffffff8311156119ff576119ff611743565b611a1383611a0d835461195f565b83611999565b6000601f841160018114611a475760008515611a2f5750838201355b600019600387901b1c1916600186901b1783556113fa565b600083815260209020601f19861690835b82811015611a785786850135825560209485019460019092019101611a58565b5086821015611a955760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b815167ffffffffffffffff811115611ac157611ac1611743565b611ad581611acf845461195f565b84611999565b602080601f831160018114611b0a5760008415611af25750858301515b600019600386901b1c1916600185901b1785556119df565b600085815260208120601f198616915b82811015611b3957888601518255948401946001909101908401611b1a565b5085821015611b575787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b600084516020611b7a8285838a016114f3565b855191840191611b8d8184848a016114f3565b8554920191600090611b9e8161195f565b60018281168015611bb65760018114611bcb57611bf7565b60ff1984168752821515830287019450611bf7565b896000528560002060005b84811015611bef57815489820152908301908701611bd6565b505082870194505b50929a9950505050505050505050565b634e487b7160e01b600052601160045260246000fd5b600081611c2c57611c2c611c07565b506000190190565b808201808211156105f2576105f2611c07565b6001600160a01b0385811682528416602082015260408101839052608060608201819052600090611c7a90830184611517565b9695505050505050565b600060208284031215611c9657600080fd5b8151610d6c816114c056fea26469706673582212206764a9a888e90e7683ac9cccda59fd74a25505231182ad35dce35f0202751a3a64736f6c63430008130033",
}

// NftABI is the input ABI used to generate the binding from.
// Deprecated: Use NftMetaData.ABI instead.
var NftABI = NftMetaData.ABI

// NftBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NftMetaData.Bin instead.
var NftBin = NftMetaData.Bin

// DeployNft deploys a new Ethereum contract, binding an instance of Nft to it.
func DeployNft(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Nft, error) {
	parsed, err := NftMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NftBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Nft{NftCaller: NftCaller{contract: contract}, NftTransactor: NftTransactor{contract: contract}, NftFilterer: NftFilterer{contract: contract}}, nil
}

// Nft is an auto generated Go binding around an Ethereum contract.
type Nft struct {
	NftCaller     // Read-only binding to the contract
	NftTransactor // Write-only binding to the contract
	NftFilterer   // Log filterer for contract events
}

// NftCaller is an auto generated read-only Go binding around an Ethereum contract.
type NftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NftSession struct {
	Contract     *Nft              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NftCallerSession struct {
	Contract *NftCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NftTransactorSession struct {
	Contract     *NftTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NftRaw is an auto generated low-level Go binding around an Ethereum contract.
type NftRaw struct {
	Contract *Nft // Generic contract binding to access the raw methods on
}

// NftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NftCallerRaw struct {
	Contract *NftCaller // Generic read-only contract binding to access the raw methods on
}

// NftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NftTransactorRaw struct {
	Contract *NftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNft creates a new instance of Nft, bound to a specific deployed contract.
func NewNft(address common.Address, backend bind.ContractBackend) (*Nft, error) {
	contract, err := bindNft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nft{NftCaller: NftCaller{contract: contract}, NftTransactor: NftTransactor{contract: contract}, NftFilterer: NftFilterer{contract: contract}}, nil
}

// NewNftCaller creates a new read-only instance of Nft, bound to a specific deployed contract.
func NewNftCaller(address common.Address, caller bind.ContractCaller) (*NftCaller, error) {
	contract, err := bindNft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NftCaller{contract: contract}, nil
}

// NewNftTransactor creates a new write-only instance of Nft, bound to a specific deployed contract.
func NewNftTransactor(address common.Address, transactor bind.ContractTransactor) (*NftTransactor, error) {
	contract, err := bindNft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NftTransactor{contract: contract}, nil
}

// NewNftFilterer creates a new log filterer instance of Nft, bound to a specific deployed contract.
func NewNftFilterer(address common.Address, filterer bind.ContractFilterer) (*NftFilterer, error) {
	contract, err := bindNft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NftFilterer{contract: contract}, nil
}

// bindNft binds a generic wrapper to an already deployed contract.
func bindNft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NftMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nft *NftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nft.Contract.NftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nft *NftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nft.Contract.NftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nft *NftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nft.Contract.NftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nft *NftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nft *NftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nft *NftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nft.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Nft *NftCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Nft *NftSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Nft.Contract.BalanceOf(&_Nft.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Nft *NftCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Nft.Contract.BalanceOf(&_Nft.CallOpts, owner)
}

// Caller is a free data retrieval call binding the contract method 0xfc9c8d39.
//
// Solidity: function caller() view returns(address)
func (_Nft *NftCaller) Caller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "caller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Caller is a free data retrieval call binding the contract method 0xfc9c8d39.
//
// Solidity: function caller() view returns(address)
func (_Nft *NftSession) Caller() (common.Address, error) {
	return _Nft.Contract.Caller(&_Nft.CallOpts)
}

// Caller is a free data retrieval call binding the contract method 0xfc9c8d39.
//
// Solidity: function caller() view returns(address)
func (_Nft *NftCallerSession) Caller() (common.Address, error) {
	return _Nft.Contract.Caller(&_Nft.CallOpts)
}

// CurrentBaseURI is a free data retrieval call binding the contract method 0x97d4ccd2.
//
// Solidity: function currentBaseURI() view returns(string)
func (_Nft *NftCaller) CurrentBaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "currentBaseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CurrentBaseURI is a free data retrieval call binding the contract method 0x97d4ccd2.
//
// Solidity: function currentBaseURI() view returns(string)
func (_Nft *NftSession) CurrentBaseURI() (string, error) {
	return _Nft.Contract.CurrentBaseURI(&_Nft.CallOpts)
}

// CurrentBaseURI is a free data retrieval call binding the contract method 0x97d4ccd2.
//
// Solidity: function currentBaseURI() view returns(string)
func (_Nft *NftCallerSession) CurrentBaseURI() (string, error) {
	return _Nft.Contract.CurrentBaseURI(&_Nft.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Nft *NftCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Nft *NftSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Nft.Contract.GetApproved(&_Nft.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Nft *NftCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Nft.Contract.GetApproved(&_Nft.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Nft *NftCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Nft *NftSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Nft.Contract.IsApprovedForAll(&_Nft.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Nft *NftCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Nft.Contract.IsApprovedForAll(&_Nft.CallOpts, owner, operator)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Nft *NftCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Nft *NftSession) MaxSupply() (*big.Int, error) {
	return _Nft.Contract.MaxSupply(&_Nft.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Nft *NftCallerSession) MaxSupply() (*big.Int, error) {
	return _Nft.Contract.MaxSupply(&_Nft.CallOpts)
}

// Mintable is a free data retrieval call binding the contract method 0x4bf365df.
//
// Solidity: function mintable() view returns(bool)
func (_Nft *NftCaller) Mintable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "mintable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Mintable is a free data retrieval call binding the contract method 0x4bf365df.
//
// Solidity: function mintable() view returns(bool)
func (_Nft *NftSession) Mintable() (bool, error) {
	return _Nft.Contract.Mintable(&_Nft.CallOpts)
}

// Mintable is a free data retrieval call binding the contract method 0x4bf365df.
//
// Solidity: function mintable() view returns(bool)
func (_Nft *NftCallerSession) Mintable() (bool, error) {
	return _Nft.Contract.Mintable(&_Nft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nft *NftCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nft *NftSession) Name() (string, error) {
	return _Nft.Contract.Name(&_Nft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nft *NftCallerSession) Name() (string, error) {
	return _Nft.Contract.Name(&_Nft.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nft *NftCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nft *NftSession) Owner() (common.Address, error) {
	return _Nft.Contract.Owner(&_Nft.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nft *NftCallerSession) Owner() (common.Address, error) {
	return _Nft.Contract.Owner(&_Nft.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Nft *NftCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Nft *NftSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Nft.Contract.OwnerOf(&_Nft.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Nft *NftCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Nft.Contract.OwnerOf(&_Nft.CallOpts, tokenId)
}

// RenounceOwnership is a free data retrieval call binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() view returns()
func (_Nft *NftCaller) RenounceOwnership(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "renounceOwnership")

	if err != nil {
		return err
	}

	return err

}

// RenounceOwnership is a free data retrieval call binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() view returns()
func (_Nft *NftSession) RenounceOwnership() error {
	return _Nft.Contract.RenounceOwnership(&_Nft.CallOpts)
}

// RenounceOwnership is a free data retrieval call binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() view returns()
func (_Nft *NftCallerSession) RenounceOwnership() error {
	return _Nft.Contract.RenounceOwnership(&_Nft.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Nft *NftCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Nft *NftSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Nft.Contract.SupportsInterface(&_Nft.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Nft *NftCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Nft.Contract.SupportsInterface(&_Nft.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nft *NftCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nft *NftSession) Symbol() (string, error) {
	return _Nft.Contract.Symbol(&_Nft.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nft *NftCallerSession) Symbol() (string, error) {
	return _Nft.Contract.Symbol(&_Nft.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Nft *NftCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Nft *NftSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Nft.Contract.TokenURI(&_Nft.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Nft *NftCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Nft.Contract.TokenURI(&_Nft.CallOpts, tokenId)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_Nft *NftCaller) TotalMinted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "totalMinted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_Nft *NftSession) TotalMinted() (*big.Int, error) {
	return _Nft.Contract.TotalMinted(&_Nft.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_Nft *NftCallerSession) TotalMinted() (*big.Int, error) {
	return _Nft.Contract.TotalMinted(&_Nft.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 result)
func (_Nft *NftCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 result)
func (_Nft *NftSession) TotalSupply() (*big.Int, error) {
	return _Nft.Contract.TotalSupply(&_Nft.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 result)
func (_Nft *NftCallerSession) TotalSupply() (*big.Int, error) {
	return _Nft.Contract.TotalSupply(&_Nft.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Nft *NftTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Nft *NftSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.Approve(&_Nft.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Nft *NftTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.Approve(&_Nft.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Nft *NftTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Nft *NftSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.Burn(&_Nft.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Nft *NftTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.Burn(&_Nft.TransactOpts, tokenId)
}

// Initailize is a paid mutator transaction binding the contract method 0xbde17fc0.
//
// Solidity: function initailize(string _name, string _symbol, string _tokenURISuffix, string _currentBaseURI, uint256 _maxSupply, address _caller, address _owner) returns()
func (_Nft *NftTransactor) Initailize(opts *bind.TransactOpts, _name string, _symbol string, _tokenURISuffix string, _currentBaseURI string, _maxSupply *big.Int, _caller common.Address, _owner common.Address) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "initailize", _name, _symbol, _tokenURISuffix, _currentBaseURI, _maxSupply, _caller, _owner)
}

// Initailize is a paid mutator transaction binding the contract method 0xbde17fc0.
//
// Solidity: function initailize(string _name, string _symbol, string _tokenURISuffix, string _currentBaseURI, uint256 _maxSupply, address _caller, address _owner) returns()
func (_Nft *NftSession) Initailize(_name string, _symbol string, _tokenURISuffix string, _currentBaseURI string, _maxSupply *big.Int, _caller common.Address, _owner common.Address) (*types.Transaction, error) {
	return _Nft.Contract.Initailize(&_Nft.TransactOpts, _name, _symbol, _tokenURISuffix, _currentBaseURI, _maxSupply, _caller, _owner)
}

// Initailize is a paid mutator transaction binding the contract method 0xbde17fc0.
//
// Solidity: function initailize(string _name, string _symbol, string _tokenURISuffix, string _currentBaseURI, uint256 _maxSupply, address _caller, address _owner) returns()
func (_Nft *NftTransactorSession) Initailize(_name string, _symbol string, _tokenURISuffix string, _currentBaseURI string, _maxSupply *big.Int, _caller common.Address, _owner common.Address) (*types.Transaction, error) {
	return _Nft.Contract.Initailize(&_Nft.TransactOpts, _name, _symbol, _tokenURISuffix, _currentBaseURI, _maxSupply, _caller, _owner)
}

// Mint is a paid mutator transaction binding the contract method 0x17deeba0.
//
// Solidity: function mint(uint64 qty, address to) returns(bool)
func (_Nft *NftTransactor) Mint(opts *bind.TransactOpts, qty uint64, to common.Address) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "mint", qty, to)
}

// Mint is a paid mutator transaction binding the contract method 0x17deeba0.
//
// Solidity: function mint(uint64 qty, address to) returns(bool)
func (_Nft *NftSession) Mint(qty uint64, to common.Address) (*types.Transaction, error) {
	return _Nft.Contract.Mint(&_Nft.TransactOpts, qty, to)
}

// Mint is a paid mutator transaction binding the contract method 0x17deeba0.
//
// Solidity: function mint(uint64 qty, address to) returns(bool)
func (_Nft *NftTransactorSession) Mint(qty uint64, to common.Address) (*types.Transaction, error) {
	return _Nft.Contract.Mint(&_Nft.TransactOpts, qty, to)
}

// Ownermint is a paid mutator transaction binding the contract method 0x94b44725.
//
// Solidity: function ownermint(address to, uint64 qty) returns()
func (_Nft *NftTransactor) Ownermint(opts *bind.TransactOpts, to common.Address, qty uint64) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "ownermint", to, qty)
}

// Ownermint is a paid mutator transaction binding the contract method 0x94b44725.
//
// Solidity: function ownermint(address to, uint64 qty) returns()
func (_Nft *NftSession) Ownermint(to common.Address, qty uint64) (*types.Transaction, error) {
	return _Nft.Contract.Ownermint(&_Nft.TransactOpts, to, qty)
}

// Ownermint is a paid mutator transaction binding the contract method 0x94b44725.
//
// Solidity: function ownermint(address to, uint64 qty) returns()
func (_Nft *NftTransactorSession) Ownermint(to common.Address, qty uint64) (*types.Transaction, error) {
	return _Nft.Contract.Ownermint(&_Nft.TransactOpts, to, qty)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Nft *NftTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Nft *NftSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.SafeTransferFrom(&_Nft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Nft *NftTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.SafeTransferFrom(&_Nft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_Nft *NftTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_Nft *NftSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Nft.Contract.SafeTransferFrom0(&_Nft.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_Nft *NftTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Nft.Contract.SafeTransferFrom0(&_Nft.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Nft *NftTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Nft *NftSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Nft.Contract.SetApprovalForAll(&_Nft.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Nft *NftTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Nft.Contract.SetApprovalForAll(&_Nft.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_Nft *NftTransactor) SetBaseURI(opts *bind.TransactOpts, baseURI string) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "setBaseURI", baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_Nft *NftSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _Nft.Contract.SetBaseURI(&_Nft.TransactOpts, baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_Nft *NftTransactorSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _Nft.Contract.SetBaseURI(&_Nft.TransactOpts, baseURI)
}

// SetMaxsupply is a paid mutator transaction binding the contract method 0x149835a0.
//
// Solidity: function setMaxsupply(uint256 _maxsupply) returns()
func (_Nft *NftTransactor) SetMaxsupply(opts *bind.TransactOpts, _maxsupply *big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "setMaxsupply", _maxsupply)
}

// SetMaxsupply is a paid mutator transaction binding the contract method 0x149835a0.
//
// Solidity: function setMaxsupply(uint256 _maxsupply) returns()
func (_Nft *NftSession) SetMaxsupply(_maxsupply *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.SetMaxsupply(&_Nft.TransactOpts, _maxsupply)
}

// SetMaxsupply is a paid mutator transaction binding the contract method 0x149835a0.
//
// Solidity: function setMaxsupply(uint256 _maxsupply) returns()
func (_Nft *NftTransactorSession) SetMaxsupply(_maxsupply *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.SetMaxsupply(&_Nft.TransactOpts, _maxsupply)
}

// SetMintable is a paid mutator transaction binding the contract method 0x285d70d4.
//
// Solidity: function setMintable(bool _mintable) returns()
func (_Nft *NftTransactor) SetMintable(opts *bind.TransactOpts, _mintable bool) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "setMintable", _mintable)
}

// SetMintable is a paid mutator transaction binding the contract method 0x285d70d4.
//
// Solidity: function setMintable(bool _mintable) returns()
func (_Nft *NftSession) SetMintable(_mintable bool) (*types.Transaction, error) {
	return _Nft.Contract.SetMintable(&_Nft.TransactOpts, _mintable)
}

// SetMintable is a paid mutator transaction binding the contract method 0x285d70d4.
//
// Solidity: function setMintable(bool _mintable) returns()
func (_Nft *NftTransactorSession) SetMintable(_mintable bool) (*types.Transaction, error) {
	return _Nft.Contract.SetMintable(&_Nft.TransactOpts, _mintable)
}

// SetTokenURISuffix is a paid mutator transaction binding the contract method 0xa9852bfb.
//
// Solidity: function setTokenURISuffix(string suffix) returns()
func (_Nft *NftTransactor) SetTokenURISuffix(opts *bind.TransactOpts, suffix string) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "setTokenURISuffix", suffix)
}

// SetTokenURISuffix is a paid mutator transaction binding the contract method 0xa9852bfb.
//
// Solidity: function setTokenURISuffix(string suffix) returns()
func (_Nft *NftSession) SetTokenURISuffix(suffix string) (*types.Transaction, error) {
	return _Nft.Contract.SetTokenURISuffix(&_Nft.TransactOpts, suffix)
}

// SetTokenURISuffix is a paid mutator transaction binding the contract method 0xa9852bfb.
//
// Solidity: function setTokenURISuffix(string suffix) returns()
func (_Nft *NftTransactorSession) SetTokenURISuffix(suffix string) (*types.Transaction, error) {
	return _Nft.Contract.SetTokenURISuffix(&_Nft.TransactOpts, suffix)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Nft *NftTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Nft *NftSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.TransferFrom(&_Nft.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Nft *NftTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.TransferFrom(&_Nft.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nft *NftTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nft *NftSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Nft.Contract.TransferOwnership(&_Nft.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nft *NftTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Nft.Contract.TransferOwnership(&_Nft.TransactOpts, newOwner)
}

// UpdateCaller is a paid mutator transaction binding the contract method 0x9b9df99b.
//
// Solidity: function updateCaller(address _caller) returns()
func (_Nft *NftTransactor) UpdateCaller(opts *bind.TransactOpts, _caller common.Address) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "updateCaller", _caller)
}

// UpdateCaller is a paid mutator transaction binding the contract method 0x9b9df99b.
//
// Solidity: function updateCaller(address _caller) returns()
func (_Nft *NftSession) UpdateCaller(_caller common.Address) (*types.Transaction, error) {
	return _Nft.Contract.UpdateCaller(&_Nft.TransactOpts, _caller)
}

// UpdateCaller is a paid mutator transaction binding the contract method 0x9b9df99b.
//
// Solidity: function updateCaller(address _caller) returns()
func (_Nft *NftTransactorSession) UpdateCaller(_caller common.Address) (*types.Transaction, error) {
	return _Nft.Contract.UpdateCaller(&_Nft.TransactOpts, _caller)
}

// NftApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Nft contract.
type NftApprovalIterator struct {
	Event *NftApproval // Event containing the contract specifics and raw log

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
func (it *NftApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftApproval)
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
		it.Event = new(NftApproval)
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
func (it *NftApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftApproval represents a Approval event raised by the Nft contract.
type NftApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Nft *NftFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*NftApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftApprovalIterator{contract: _Nft.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Nft *NftFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NftApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftApproval)
				if err := _Nft.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Nft *NftFilterer) ParseApproval(log types.Log) (*NftApproval, error) {
	event := new(NftApproval)
	if err := _Nft.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Nft contract.
type NftApprovalForAllIterator struct {
	Event *NftApprovalForAll // Event containing the contract specifics and raw log

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
func (it *NftApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftApprovalForAll)
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
		it.Event = new(NftApprovalForAll)
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
func (it *NftApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftApprovalForAll represents a ApprovalForAll event raised by the Nft contract.
type NftApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Nft *NftFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*NftApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &NftApprovalForAllIterator{contract: _Nft.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Nft *NftFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *NftApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftApprovalForAll)
				if err := _Nft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Nft *NftFilterer) ParseApprovalForAll(log types.Log) (*NftApprovalForAll, error) {
	event := new(NftApprovalForAll)
	if err := _Nft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftConsecutiveTransferIterator is returned from FilterConsecutiveTransfer and is used to iterate over the raw logs and unpacked data for ConsecutiveTransfer events raised by the Nft contract.
type NftConsecutiveTransferIterator struct {
	Event *NftConsecutiveTransfer // Event containing the contract specifics and raw log

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
func (it *NftConsecutiveTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftConsecutiveTransfer)
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
		it.Event = new(NftConsecutiveTransfer)
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
func (it *NftConsecutiveTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftConsecutiveTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftConsecutiveTransfer represents a ConsecutiveTransfer event raised by the Nft contract.
type NftConsecutiveTransfer struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	From        common.Address
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsecutiveTransfer is a free log retrieval operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Nft *NftFilterer) FilterConsecutiveTransfer(opts *bind.FilterOpts, fromTokenId []*big.Int, from []common.Address, to []common.Address) (*NftConsecutiveTransferIterator, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NftConsecutiveTransferIterator{contract: _Nft.contract, event: "ConsecutiveTransfer", logs: logs, sub: sub}, nil
}

// WatchConsecutiveTransfer is a free log subscription operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Nft *NftFilterer) WatchConsecutiveTransfer(opts *bind.WatchOpts, sink chan<- *NftConsecutiveTransfer, fromTokenId []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftConsecutiveTransfer)
				if err := _Nft.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
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

// ParseConsecutiveTransfer is a log parse operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Nft *NftFilterer) ParseConsecutiveTransfer(log types.Log) (*NftConsecutiveTransfer, error) {
	event := new(NftConsecutiveTransfer)
	if err := _Nft.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Nft contract.
type NftOwnershipTransferredIterator struct {
	Event *NftOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NftOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftOwnershipTransferred)
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
		it.Event = new(NftOwnershipTransferred)
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
func (it *NftOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftOwnershipTransferred represents a OwnershipTransferred event raised by the Nft contract.
type NftOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Nft *NftFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NftOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NftOwnershipTransferredIterator{contract: _Nft.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Nft *NftFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NftOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftOwnershipTransferred)
				if err := _Nft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Nft *NftFilterer) ParseOwnershipTransferred(log types.Log) (*NftOwnershipTransferred, error) {
	event := new(NftOwnershipTransferred)
	if err := _Nft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Nft contract.
type NftTransferIterator struct {
	Event *NftTransfer // Event containing the contract specifics and raw log

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
func (it *NftTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftTransfer)
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
		it.Event = new(NftTransfer)
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
func (it *NftTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftTransfer represents a Transfer event raised by the Nft contract.
type NftTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Nft *NftFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*NftTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftTransferIterator{contract: _Nft.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Nft *NftFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NftTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftTransfer)
				if err := _Nft.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Nft *NftFilterer) ParseTransfer(log types.Log) (*NftTransfer, error) {
	event := new(NftTransfer)
	if err := _Nft.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
