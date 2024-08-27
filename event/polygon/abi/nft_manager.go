// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// NftManagerMetaData contains all meta data concerning the NftManager contract.
var NftManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_fccTokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenUsdtAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_redemptionPoolAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"_nextTokenId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"basicMineAmt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNFT\",\"inputs\":[{\"name\":\"_businessName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_description\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_imgUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_businessAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_website\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_social\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_type\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"fccTokenAddr\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApproved\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMerchantNTFDeadline\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenBalance\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserNTFDeadline\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isApprovedForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"merchantNftDeadline\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"merchantValue\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minedAmt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nftMintType\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerOf\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proMineAmt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"redemptionPoolAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRedemptionPool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setApprovalForAll\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUriPrefix\",\"inputs\":[{\"name\":\"_uriPrefix\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setValues\",\"inputs\":[{\"name\":\"_merchantValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_userValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenURI\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenUsdtAddr\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalMineAmt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"uri\",\"inputs\":[{\"name\":\"uri_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"uriPrefix\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"userNftDeadline\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"userValue\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawNativeToken\",\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawToken\",\"inputs\":[{\"name\":\"_tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApprovalForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BatchMetadataUpdate\",\"inputs\":[{\"name\":\"_fromTokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_toTokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CreateNFT\",\"inputs\":[{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_businessName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"_description\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"_imgUrl\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"_businessAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"_webSite\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"_social\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"_value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_deadline\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_type\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetadataUpdate\",\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Received\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetValidTime\",\"inputs\":[{\"name\":\"setter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_time\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetValues\",\"inputs\":[{\"name\":\"_setterAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_merchantValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_userValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UriPrefixSet\",\"inputs\":[{\"name\":\"setterAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"urlPrefix\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"withdrawer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawUToken\",\"inputs\":[{\"name\":\"withdrawer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_tokenAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ERC721IncorrectOwner\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InsufficientApproval\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721NonexistentToken\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
}

// NftManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use NftManagerMetaData.ABI instead.
var NftManagerABI = NftManagerMetaData.ABI

// NftManager is an auto generated Go binding around an Ethereum contract.
type NftManager struct {
	NftManagerCaller     // Read-only binding to the contract
	NftManagerTransactor // Write-only binding to the contract
	NftManagerFilterer   // Log filterer for contract events
}

// NftManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type NftManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NftManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NftManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NftManagerSession struct {
	Contract     *NftManager       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NftManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NftManagerCallerSession struct {
	Contract *NftManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// NftManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NftManagerTransactorSession struct {
	Contract     *NftManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// NftManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type NftManagerRaw struct {
	Contract *NftManager // Generic contract binding to access the raw methods on
}

// NftManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NftManagerCallerRaw struct {
	Contract *NftManagerCaller // Generic read-only contract binding to access the raw methods on
}

// NftManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NftManagerTransactorRaw struct {
	Contract *NftManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNftManager creates a new instance of NftManager, bound to a specific deployed contract.
func NewNftManager(address common.Address, backend bind.ContractBackend) (*NftManager, error) {
	contract, err := bindNftManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NftManager{NftManagerCaller: NftManagerCaller{contract: contract}, NftManagerTransactor: NftManagerTransactor{contract: contract}, NftManagerFilterer: NftManagerFilterer{contract: contract}}, nil
}

// NewNftManagerCaller creates a new read-only instance of NftManager, bound to a specific deployed contract.
func NewNftManagerCaller(address common.Address, caller bind.ContractCaller) (*NftManagerCaller, error) {
	contract, err := bindNftManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NftManagerCaller{contract: contract}, nil
}

// NewNftManagerTransactor creates a new write-only instance of NftManager, bound to a specific deployed contract.
func NewNftManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*NftManagerTransactor, error) {
	contract, err := bindNftManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NftManagerTransactor{contract: contract}, nil
}

// NewNftManagerFilterer creates a new log filterer instance of NftManager, bound to a specific deployed contract.
func NewNftManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*NftManagerFilterer, error) {
	contract, err := bindNftManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NftManagerFilterer{contract: contract}, nil
}

// bindNftManager binds a generic wrapper to an already deployed contract.
func bindNftManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NftManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NftManager *NftManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NftManager.Contract.NftManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NftManager *NftManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NftManager.Contract.NftManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NftManager *NftManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NftManager.Contract.NftManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NftManager *NftManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NftManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NftManager *NftManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NftManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NftManager *NftManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NftManager.Contract.contract.Transact(opts, method, params...)
}

// NextTokenId is a free data retrieval call binding the contract method 0x4a60f620.
//
// Solidity: function _nextTokenId() view returns(uint256)
func (_NftManager *NftManagerCaller) NextTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NftManager.contract.Call(opts, &out, "_nextTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenId is a free data retrieval call binding the contract method 0x4a60f620.
//
// Solidity: function _nextTokenId() view returns(uint256)
func (_NftManager *NftManagerSession) NextTokenId() (*big.Int, error) {
	return _NftManager.Contract.NextTokenId(&_NftManager.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x4a60f620.
//
// Solidity: function _nextTokenId() view returns(uint256)
func (_NftManager *NftManagerCallerSession) NextTokenId() (*big.Int, error) {
	return _NftManager.Contract.NextTokenId(&_NftManager.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NftManager *NftManagerCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NftManager.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NftManager *NftManagerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _NftManager.Contract.BalanceOf(&_NftManager.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NftManager *NftManagerCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _NftManager.Contract.BalanceOf(&_NftManager.CallOpts, owner)
}

// BasicMineAmt is a free data retrieval call binding the contract method 0x94399817.
//
// Solidity: function basicMineAmt() view returns(uint256)
func (_NftManager *NftManagerCaller) BasicMineAmt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NftManager.contract.Call(opts, &out, "basicMineAmt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BasicMineAmt is a free data retrieval call binding the contract method 0x94399817.
//
// Solidity: function basicMineAmt() view returns(uint256)
func (_NftManager *NftManagerSession) BasicMineAmt() (*big.Int, error) {
	return _NftManager.Contract.BasicMineAmt(&_NftManager.CallOpts)
}

// BasicMineAmt is a free data retrieval call binding the contract method 0x94399817.
//
// Solidity: function basicMineAmt() view returns(uint256)
func (_NftManager *NftManagerCallerSession) BasicMineAmt() (*big.Int, error) {
	return _NftManager.Contract.BasicMineAmt(&_NftManager.CallOpts)
}

// FccTokenAddr is a free data retrieval call binding the contract method 0xa73dd74c.
//
// Solidity: function fccTokenAddr() view returns(address)
func (_NftManager *NftManagerCaller) FccTokenAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NftManager.contract.Call(opts, &out, "fccTokenAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FccTokenAddr is a free data retrieval call binding the contract method 0xa73dd74c.
//
// Solidity: function fccTokenAddr() view returns(address)
func (_NftManager *NftManagerSession) FccTokenAddr() (common.Address, error) {
	return _NftManager.Contract.FccTokenAddr(&_NftManager.CallOpts)
}

// FccTokenAddr is a free data retrieval call binding the contract method 0xa73dd74c.
//
// Solidity: function fccTokenAddr() view returns(address)
func (_NftManager *NftManagerCallerSession) FccTokenAddr() (common.Address, error) {
	return _NftManager.Contract.FccTokenAddr(&_NftManager.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NftManager *NftManagerCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NftManager.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NftManager *NftManagerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _NftManager.Contract.GetApproved(&_NftManager.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NftManager *NftManagerCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _NftManager.Contract.GetApproved(&_NftManager.CallOpts, tokenId)
}

// GetMerchantNTFDeadline is a free data retrieval call binding the contract method 0x39795078.
//
// Solidity: function getMerchantNTFDeadline(address _account) view returns(uint256)
func (_NftManager *NftManagerCaller) GetMerchantNTFDeadline(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NftManager.contract.Call(opts, &out, "getMerchantNTFDeadline", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMerchantNTFDeadline is a free data retrieval call binding the contract method 0x39795078.
//
// Solidity: function getMerchantNTFDeadline(address _account) view returns(uint256)
func (_NftManager *NftManagerSession) GetMerchantNTFDeadline(_account common.Address) (*big.Int, error) {
	return _NftManager.Contract.GetMerchantNTFDeadline(&_NftManager.CallOpts, _account)
}

// GetMerchantNTFDeadline is a free data retrieval call binding the contract method 0x39795078.
//
// Solidity: function getMerchantNTFDeadline(address _account) view returns(uint256)
func (_NftManager *NftManagerCallerSession) GetMerchantNTFDeadline(_account common.Address) (*big.Int, error) {
	return _NftManager.Contract.GetMerchantNTFDeadline(&_NftManager.CallOpts, _account)
}

// GetTokenBalance is a free data retrieval call binding the contract method 0x3aecd0e3.
//
// Solidity: function getTokenBalance(address tokenAddress) view returns(uint256)
func (_NftManager *NftManagerCaller) GetTokenBalance(opts *bind.CallOpts, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NftManager.contract.Call(opts, &out, "getTokenBalance", tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenBalance is a free data retrieval call binding the contract method 0x3aecd0e3.
//
// Solidity: function getTokenBalance(address tokenAddress) view returns(uint256)
func (_NftManager *NftManagerSession) GetTokenBalance(tokenAddress common.Address) (*big.Int, error) {
	return _NftManager.Contract.GetTokenBalance(&_NftManager.CallOpts, tokenAddress)
}

// GetTokenBalance is a free data retrieval call binding the contract method 0x3aecd0e3.
//
// Solidity: function getTokenBalance(address tokenAddress) view returns(uint256)
func (_NftManager *NftManagerCallerSession) GetTokenBalance(tokenAddress common.Address) (*big.Int, error) {
	return _NftManager.Contract.GetTokenBalance(&_NftManager.CallOpts, tokenAddress)
}

// GetUserNTFDeadline is a free data retrieval call binding the contract method 0x8dce8b13.
//
// Solidity: function getUserNTFDeadline(address _account) view returns(uint256)
func (_NftManager *NftManagerCaller) GetUserNTFDeadline(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NftManager.contract.Call(opts, &out, "getUserNTFDeadline", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserNTFDeadline is a free data retrieval call binding the contract method 0x8dce8b13.
//
// Solidity: function getUserNTFDeadline(address _account) view returns(uint256)
func (_NftManager *NftManagerSession) GetUserNTFDeadline(_account common.Address) (*big.Int, error) {
	return _NftManager.Contract.GetUserNTFDeadline(&_NftManager.CallOpts, _account)
}

// GetUserNTFDeadline is a free data retrieval call binding the contract method 0x8dce8b13.
//
// Solidity: function getUserNTFDeadline(address _account) view returns(uint256)
func (_NftManager *NftManagerCallerSession) GetUserNTFDeadline(_account common.Address) (*big.Int, error) {
	return _NftManager.Contract.GetUserNTFDeadline(&_NftManager.CallOpts, _account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NftManager *NftManagerCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _NftManager.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NftManager *NftManagerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _NftManager.Contract.IsApprovedForAll(&_NftManager.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NftManager *NftManagerCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _NftManager.Contract.IsApprovedForAll(&_NftManager.CallOpts, owner, operator)
}

// MerchantNftDeadline is a free data retrieval call binding the contract method 0x1c2347a8.
//
// Solidity: function merchantNftDeadline(address ) view returns(uint256)
func (_NftManager *NftManagerCaller) MerchantNftDeadline(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NftManager.contract.Call(opts, &out, "merchantNftDeadline", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MerchantNftDeadline is a free data retrieval call binding the contract method 0x1c2347a8.
//
// Solidity: function merchantNftDeadline(address ) view returns(uint256)
func (_NftManager *NftManagerSession) MerchantNftDeadline(arg0 common.Address) (*big.Int, error) {
	return _NftManager.Contract.MerchantNftDeadline(&_NftManager.CallOpts, arg0)
}

// MerchantNftDeadline is a free data retrieval call binding the contract method 0x1c2347a8.
//
// Solidity: function merchantNftDeadline(address ) view returns(uint256)
func (_NftManager *NftManagerCallerSession) MerchantNftDeadline(arg0 common.Address) (*big.Int, error) {
	return _NftManager.Contract.MerchantNftDeadline(&_NftManager.CallOpts, arg0)
}

// MerchantValue is a free data retrieval call binding the contract method 0x748dada1.
//
// Solidity: function merchantValue() view returns(uint256)
func (_NftManager *NftManagerCaller) MerchantValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NftManager.contract.Call(opts, &out, "merchantValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MerchantValue is a free data retrieval call binding the contract method 0x748dada1.
//
// Solidity: function merchantValue() view returns(uint256)
func (_NftManager *NftManagerSession) MerchantValue() (*big.Int, error) {
	return _NftManager.Contract.MerchantValue(&_NftManager.CallOpts)
}

// MerchantValue is a free data retrieval call binding the contract method 0x748dada1.
//
// Solidity: function merchantValue() view returns(uint256)
func (_NftManager *NftManagerCallerSession) MerchantValue() (*big.Int, error) {
	return _NftManager.Contract.MerchantValue(&_NftManager.CallOpts)
}

// MinedAmt is a free data retrieval call binding the contract method 0xf7859c28.
//
// Solidity: function minedAmt() view returns(uint256)
func (_NftManager *NftManagerCaller) MinedAmt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NftManager.contract.Call(opts, &out, "minedAmt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinedAmt is a free data retrieval call binding the contract method 0xf7859c28.
//
// Solidity: function minedAmt() view returns(uint256)
func (_NftManager *NftManagerSession) MinedAmt() (*big.Int, error) {
	return _NftManager.Contract.MinedAmt(&_NftManager.CallOpts)
}

// MinedAmt is a free data retrieval call binding the contract method 0xf7859c28.
//
// Solidity: function minedAmt() view returns(uint256)
func (_NftManager *NftManagerCallerSession) MinedAmt() (*big.Int, error) {
	return _NftManager.Contract.MinedAmt(&_NftManager.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NftManager *NftManagerCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NftManager.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NftManager *NftManagerSession) Name() (string, error) {
	return _NftManager.Contract.Name(&_NftManager.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NftManager *NftManagerCallerSession) Name() (string, error) {
	return _NftManager.Contract.Name(&_NftManager.CallOpts)
}

// NftMintType is a free data retrieval call binding the contract method 0x2ac8693b.
//
// Solidity: function nftMintType(uint256 ) view returns(uint8)
func (_NftManager *NftManagerCaller) NftMintType(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var out []interface{}
	err := _NftManager.contract.Call(opts, &out, "nftMintType", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// NftMintType is a free data retrieval call binding the contract method 0x2ac8693b.
//
// Solidity: function nftMintType(uint256 ) view returns(uint8)
func (_NftManager *NftManagerSession) NftMintType(arg0 *big.Int) (uint8, error) {
	return _NftManager.Contract.NftMintType(&_NftManager.CallOpts, arg0)
}

// NftMintType is a free data retrieval call binding the contract method 0x2ac8693b.
//
// Solidity: function nftMintType(uint256 ) view returns(uint8)
func (_NftManager *NftManagerCallerSession) NftMintType(arg0 *big.Int) (uint8, error) {
	return _NftManager.Contract.NftMintType(&_NftManager.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NftManager *NftManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NftManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NftManager *NftManagerSession) Owner() (common.Address, error) {
	return _NftManager.Contract.Owner(&_NftManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NftManager *NftManagerCallerSession) Owner() (common.Address, error) {
	return _NftManager.Contract.Owner(&_NftManager.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NftManager *NftManagerCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NftManager.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NftManager *NftManagerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _NftManager.Contract.OwnerOf(&_NftManager.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NftManager *NftManagerCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _NftManager.Contract.OwnerOf(&_NftManager.CallOpts, tokenId)
}

// ProMineAmt is a free data retrieval call binding the contract method 0x80c7fb97.
//
// Solidity: function proMineAmt() view returns(uint256)
func (_NftManager *NftManagerCaller) ProMineAmt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NftManager.contract.Call(opts, &out, "proMineAmt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProMineAmt is a free data retrieval call binding the contract method 0x80c7fb97.
//
// Solidity: function proMineAmt() view returns(uint256)
func (_NftManager *NftManagerSession) ProMineAmt() (*big.Int, error) {
	return _NftManager.Contract.ProMineAmt(&_NftManager.CallOpts)
}

// ProMineAmt is a free data retrieval call binding the contract method 0x80c7fb97.
//
// Solidity: function proMineAmt() view returns(uint256)
func (_NftManager *NftManagerCallerSession) ProMineAmt() (*big.Int, error) {
	return _NftManager.Contract.ProMineAmt(&_NftManager.CallOpts)
}

// RedemptionPoolAddress is a free data retrieval call binding the contract method 0xb6648e47.
//
// Solidity: function redemptionPoolAddress() view returns(address)
func (_NftManager *NftManagerCaller) RedemptionPoolAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NftManager.contract.Call(opts, &out, "redemptionPoolAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RedemptionPoolAddress is a free data retrieval call binding the contract method 0xb6648e47.
//
// Solidity: function redemptionPoolAddress() view returns(address)
func (_NftManager *NftManagerSession) RedemptionPoolAddress() (common.Address, error) {
	return _NftManager.Contract.RedemptionPoolAddress(&_NftManager.CallOpts)
}

// RedemptionPoolAddress is a free data retrieval call binding the contract method 0xb6648e47.
//
// Solidity: function redemptionPoolAddress() view returns(address)
func (_NftManager *NftManagerCallerSession) RedemptionPoolAddress() (common.Address, error) {
	return _NftManager.Contract.RedemptionPoolAddress(&_NftManager.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NftManager *NftManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _NftManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NftManager *NftManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NftManager.Contract.SupportsInterface(&_NftManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NftManager *NftManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NftManager.Contract.SupportsInterface(&_NftManager.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NftManager *NftManagerCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NftManager.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NftManager *NftManagerSession) Symbol() (string, error) {
	return _NftManager.Contract.Symbol(&_NftManager.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NftManager *NftManagerCallerSession) Symbol() (string, error) {
	return _NftManager.Contract.Symbol(&_NftManager.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_NftManager *NftManagerCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _NftManager.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_NftManager *NftManagerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _NftManager.Contract.TokenURI(&_NftManager.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_NftManager *NftManagerCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _NftManager.Contract.TokenURI(&_NftManager.CallOpts, tokenId)
}

// TokenUsdtAddr is a free data retrieval call binding the contract method 0xdd62df73.
//
// Solidity: function tokenUsdtAddr() view returns(address)
func (_NftManager *NftManagerCaller) TokenUsdtAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NftManager.contract.Call(opts, &out, "tokenUsdtAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenUsdtAddr is a free data retrieval call binding the contract method 0xdd62df73.
//
// Solidity: function tokenUsdtAddr() view returns(address)
func (_NftManager *NftManagerSession) TokenUsdtAddr() (common.Address, error) {
	return _NftManager.Contract.TokenUsdtAddr(&_NftManager.CallOpts)
}

// TokenUsdtAddr is a free data retrieval call binding the contract method 0xdd62df73.
//
// Solidity: function tokenUsdtAddr() view returns(address)
func (_NftManager *NftManagerCallerSession) TokenUsdtAddr() (common.Address, error) {
	return _NftManager.Contract.TokenUsdtAddr(&_NftManager.CallOpts)
}

// TotalMineAmt is a free data retrieval call binding the contract method 0x88e40a5a.
//
// Solidity: function totalMineAmt() view returns(uint256)
func (_NftManager *NftManagerCaller) TotalMineAmt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NftManager.contract.Call(opts, &out, "totalMineAmt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMineAmt is a free data retrieval call binding the contract method 0x88e40a5a.
//
// Solidity: function totalMineAmt() view returns(uint256)
func (_NftManager *NftManagerSession) TotalMineAmt() (*big.Int, error) {
	return _NftManager.Contract.TotalMineAmt(&_NftManager.CallOpts)
}

// TotalMineAmt is a free data retrieval call binding the contract method 0x88e40a5a.
//
// Solidity: function totalMineAmt() view returns(uint256)
func (_NftManager *NftManagerCallerSession) TotalMineAmt() (*big.Int, error) {
	return _NftManager.Contract.TotalMineAmt(&_NftManager.CallOpts)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 uri_) view returns(string)
func (_NftManager *NftManagerCaller) Uri(opts *bind.CallOpts, uri_ *big.Int) (string, error) {
	var out []interface{}
	err := _NftManager.contract.Call(opts, &out, "uri", uri_)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 uri_) view returns(string)
func (_NftManager *NftManagerSession) Uri(uri_ *big.Int) (string, error) {
	return _NftManager.Contract.Uri(&_NftManager.CallOpts, uri_)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 uri_) view returns(string)
func (_NftManager *NftManagerCallerSession) Uri(uri_ *big.Int) (string, error) {
	return _NftManager.Contract.Uri(&_NftManager.CallOpts, uri_)
}

// UriPrefix is a free data retrieval call binding the contract method 0x62b99ad4.
//
// Solidity: function uriPrefix() view returns(string)
func (_NftManager *NftManagerCaller) UriPrefix(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NftManager.contract.Call(opts, &out, "uriPrefix")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UriPrefix is a free data retrieval call binding the contract method 0x62b99ad4.
//
// Solidity: function uriPrefix() view returns(string)
func (_NftManager *NftManagerSession) UriPrefix() (string, error) {
	return _NftManager.Contract.UriPrefix(&_NftManager.CallOpts)
}

// UriPrefix is a free data retrieval call binding the contract method 0x62b99ad4.
//
// Solidity: function uriPrefix() view returns(string)
func (_NftManager *NftManagerCallerSession) UriPrefix() (string, error) {
	return _NftManager.Contract.UriPrefix(&_NftManager.CallOpts)
}

// UserNftDeadline is a free data retrieval call binding the contract method 0x0ddf45df.
//
// Solidity: function userNftDeadline(address ) view returns(uint256)
func (_NftManager *NftManagerCaller) UserNftDeadline(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NftManager.contract.Call(opts, &out, "userNftDeadline", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserNftDeadline is a free data retrieval call binding the contract method 0x0ddf45df.
//
// Solidity: function userNftDeadline(address ) view returns(uint256)
func (_NftManager *NftManagerSession) UserNftDeadline(arg0 common.Address) (*big.Int, error) {
	return _NftManager.Contract.UserNftDeadline(&_NftManager.CallOpts, arg0)
}

// UserNftDeadline is a free data retrieval call binding the contract method 0x0ddf45df.
//
// Solidity: function userNftDeadline(address ) view returns(uint256)
func (_NftManager *NftManagerCallerSession) UserNftDeadline(arg0 common.Address) (*big.Int, error) {
	return _NftManager.Contract.UserNftDeadline(&_NftManager.CallOpts, arg0)
}

// UserValue is a free data retrieval call binding the contract method 0x8678d9d4.
//
// Solidity: function userValue() view returns(uint256)
func (_NftManager *NftManagerCaller) UserValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NftManager.contract.Call(opts, &out, "userValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserValue is a free data retrieval call binding the contract method 0x8678d9d4.
//
// Solidity: function userValue() view returns(uint256)
func (_NftManager *NftManagerSession) UserValue() (*big.Int, error) {
	return _NftManager.Contract.UserValue(&_NftManager.CallOpts)
}

// UserValue is a free data retrieval call binding the contract method 0x8678d9d4.
//
// Solidity: function userValue() view returns(uint256)
func (_NftManager *NftManagerCallerSession) UserValue() (*big.Int, error) {
	return _NftManager.Contract.UserValue(&_NftManager.CallOpts)
}

// ValidTime is a free data retrieval call binding the contract method 0xbbe44917.
//
// Solidity: function validTime() view returns(uint256)
func (_NftManager *NftManagerCaller) ValidTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NftManager.contract.Call(opts, &out, "validTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidTime is a free data retrieval call binding the contract method 0xbbe44917.
//
// Solidity: function validTime() view returns(uint256)
func (_NftManager *NftManagerSession) ValidTime() (*big.Int, error) {
	return _NftManager.Contract.ValidTime(&_NftManager.CallOpts)
}

// ValidTime is a free data retrieval call binding the contract method 0xbbe44917.
//
// Solidity: function validTime() view returns(uint256)
func (_NftManager *NftManagerCallerSession) ValidTime() (*big.Int, error) {
	return _NftManager.Contract.ValidTime(&_NftManager.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_NftManager *NftManagerTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftManager.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_NftManager *NftManagerSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftManager.Contract.Approve(&_NftManager.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_NftManager *NftManagerTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftManager.Contract.Approve(&_NftManager.TransactOpts, to, tokenId)
}

// CreateNFT is a paid mutator transaction binding the contract method 0x463d5059.
//
// Solidity: function createNFT(string _businessName, string _description, string _imgUrl, string _businessAddress, string _website, string _social, uint8 _type) returns(bool, uint256)
func (_NftManager *NftManagerTransactor) CreateNFT(opts *bind.TransactOpts, _businessName string, _description string, _imgUrl string, _businessAddress string, _website string, _social string, _type uint8) (*types.Transaction, error) {
	return _NftManager.contract.Transact(opts, "createNFT", _businessName, _description, _imgUrl, _businessAddress, _website, _social, _type)
}

// CreateNFT is a paid mutator transaction binding the contract method 0x463d5059.
//
// Solidity: function createNFT(string _businessName, string _description, string _imgUrl, string _businessAddress, string _website, string _social, uint8 _type) returns(bool, uint256)
func (_NftManager *NftManagerSession) CreateNFT(_businessName string, _description string, _imgUrl string, _businessAddress string, _website string, _social string, _type uint8) (*types.Transaction, error) {
	return _NftManager.Contract.CreateNFT(&_NftManager.TransactOpts, _businessName, _description, _imgUrl, _businessAddress, _website, _social, _type)
}

// CreateNFT is a paid mutator transaction binding the contract method 0x463d5059.
//
// Solidity: function createNFT(string _businessName, string _description, string _imgUrl, string _businessAddress, string _website, string _social, uint8 _type) returns(bool, uint256)
func (_NftManager *NftManagerTransactorSession) CreateNFT(_businessName string, _description string, _imgUrl string, _businessAddress string, _website string, _social string, _type uint8) (*types.Transaction, error) {
	return _NftManager.Contract.CreateNFT(&_NftManager.TransactOpts, _businessName, _description, _imgUrl, _businessAddress, _website, _social, _type)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _initialOwner) returns()
func (_NftManager *NftManagerTransactor) Initialize(opts *bind.TransactOpts, _initialOwner common.Address) (*types.Transaction, error) {
	return _NftManager.contract.Transact(opts, "initialize", _initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _initialOwner) returns()
func (_NftManager *NftManagerSession) Initialize(_initialOwner common.Address) (*types.Transaction, error) {
	return _NftManager.Contract.Initialize(&_NftManager.TransactOpts, _initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _initialOwner) returns()
func (_NftManager *NftManagerTransactorSession) Initialize(_initialOwner common.Address) (*types.Transaction, error) {
	return _NftManager.Contract.Initialize(&_NftManager.TransactOpts, _initialOwner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NftManager *NftManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NftManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NftManager *NftManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _NftManager.Contract.RenounceOwnership(&_NftManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NftManager *NftManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NftManager.Contract.RenounceOwnership(&_NftManager.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_NftManager *NftManagerTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftManager.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_NftManager *NftManagerSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftManager.Contract.SafeTransferFrom(&_NftManager.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_NftManager *NftManagerTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftManager.Contract.SafeTransferFrom(&_NftManager.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_NftManager *NftManagerTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _NftManager.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_NftManager *NftManagerSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _NftManager.Contract.SafeTransferFrom0(&_NftManager.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_NftManager *NftManagerTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _NftManager.Contract.SafeTransferFrom0(&_NftManager.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NftManager *NftManagerTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _NftManager.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NftManager *NftManagerSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _NftManager.Contract.SetApprovalForAll(&_NftManager.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NftManager *NftManagerTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _NftManager.Contract.SetApprovalForAll(&_NftManager.TransactOpts, operator, approved)
}

// SetUriPrefix is a paid mutator transaction binding the contract method 0x7ec4a659.
//
// Solidity: function setUriPrefix(string _uriPrefix) returns()
func (_NftManager *NftManagerTransactor) SetUriPrefix(opts *bind.TransactOpts, _uriPrefix string) (*types.Transaction, error) {
	return _NftManager.contract.Transact(opts, "setUriPrefix", _uriPrefix)
}

// SetUriPrefix is a paid mutator transaction binding the contract method 0x7ec4a659.
//
// Solidity: function setUriPrefix(string _uriPrefix) returns()
func (_NftManager *NftManagerSession) SetUriPrefix(_uriPrefix string) (*types.Transaction, error) {
	return _NftManager.Contract.SetUriPrefix(&_NftManager.TransactOpts, _uriPrefix)
}

// SetUriPrefix is a paid mutator transaction binding the contract method 0x7ec4a659.
//
// Solidity: function setUriPrefix(string _uriPrefix) returns()
func (_NftManager *NftManagerTransactorSession) SetUriPrefix(_uriPrefix string) (*types.Transaction, error) {
	return _NftManager.Contract.SetUriPrefix(&_NftManager.TransactOpts, _uriPrefix)
}

// SetValues is a paid mutator transaction binding the contract method 0x9c3dfaf6.
//
// Solidity: function setValues(uint256 _merchantValue, uint256 _userValue) returns()
func (_NftManager *NftManagerTransactor) SetValues(opts *bind.TransactOpts, _merchantValue *big.Int, _userValue *big.Int) (*types.Transaction, error) {
	return _NftManager.contract.Transact(opts, "setValues", _merchantValue, _userValue)
}

// SetValues is a paid mutator transaction binding the contract method 0x9c3dfaf6.
//
// Solidity: function setValues(uint256 _merchantValue, uint256 _userValue) returns()
func (_NftManager *NftManagerSession) SetValues(_merchantValue *big.Int, _userValue *big.Int) (*types.Transaction, error) {
	return _NftManager.Contract.SetValues(&_NftManager.TransactOpts, _merchantValue, _userValue)
}

// SetValues is a paid mutator transaction binding the contract method 0x9c3dfaf6.
//
// Solidity: function setValues(uint256 _merchantValue, uint256 _userValue) returns()
func (_NftManager *NftManagerTransactorSession) SetValues(_merchantValue *big.Int, _userValue *big.Int) (*types.Transaction, error) {
	return _NftManager.Contract.SetValues(&_NftManager.TransactOpts, _merchantValue, _userValue)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_NftManager *NftManagerTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftManager.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_NftManager *NftManagerSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftManager.Contract.TransferFrom(&_NftManager.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_NftManager *NftManagerTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftManager.Contract.TransferFrom(&_NftManager.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NftManager *NftManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NftManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NftManager *NftManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NftManager.Contract.TransferOwnership(&_NftManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NftManager *NftManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NftManager.Contract.TransferOwnership(&_NftManager.TransactOpts, newOwner)
}

// WithdrawNativeToken is a paid mutator transaction binding the contract method 0x536c6bfa.
//
// Solidity: function withdrawNativeToken(address _recipient, uint256 _amount) returns(bool)
func (_NftManager *NftManagerTransactor) WithdrawNativeToken(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _NftManager.contract.Transact(opts, "withdrawNativeToken", _recipient, _amount)
}

// WithdrawNativeToken is a paid mutator transaction binding the contract method 0x536c6bfa.
//
// Solidity: function withdrawNativeToken(address _recipient, uint256 _amount) returns(bool)
func (_NftManager *NftManagerSession) WithdrawNativeToken(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _NftManager.Contract.WithdrawNativeToken(&_NftManager.TransactOpts, _recipient, _amount)
}

// WithdrawNativeToken is a paid mutator transaction binding the contract method 0x536c6bfa.
//
// Solidity: function withdrawNativeToken(address _recipient, uint256 _amount) returns(bool)
func (_NftManager *NftManagerTransactorSession) WithdrawNativeToken(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _NftManager.Contract.WithdrawNativeToken(&_NftManager.TransactOpts, _recipient, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _tokenAddr, address _account, uint256 _value) returns(bool)
func (_NftManager *NftManagerTransactor) WithdrawToken(opts *bind.TransactOpts, _tokenAddr common.Address, _account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _NftManager.contract.Transact(opts, "withdrawToken", _tokenAddr, _account, _value)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _tokenAddr, address _account, uint256 _value) returns(bool)
func (_NftManager *NftManagerSession) WithdrawToken(_tokenAddr common.Address, _account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _NftManager.Contract.WithdrawToken(&_NftManager.TransactOpts, _tokenAddr, _account, _value)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _tokenAddr, address _account, uint256 _value) returns(bool)
func (_NftManager *NftManagerTransactorSession) WithdrawToken(_tokenAddr common.Address, _account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _NftManager.Contract.WithdrawToken(&_NftManager.TransactOpts, _tokenAddr, _account, _value)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NftManager *NftManagerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NftManager.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NftManager *NftManagerSession) Receive() (*types.Transaction, error) {
	return _NftManager.Contract.Receive(&_NftManager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NftManager *NftManagerTransactorSession) Receive() (*types.Transaction, error) {
	return _NftManager.Contract.Receive(&_NftManager.TransactOpts)
}

// NftManagerApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the NftManager contract.
type NftManagerApprovalIterator struct {
	Event *NftManagerApproval // Event containing the contract specifics and raw log

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
func (it *NftManagerApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftManagerApproval)
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
		it.Event = new(NftManagerApproval)
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
func (it *NftManagerApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftManagerApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftManagerApproval represents a Approval event raised by the NftManager contract.
type NftManagerApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_NftManager *NftManagerFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*NftManagerApprovalIterator, error) {

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

	logs, sub, err := _NftManager.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftManagerApprovalIterator{contract: _NftManager.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_NftManager *NftManagerFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NftManagerApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _NftManager.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftManagerApproval)
				if err := _NftManager.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_NftManager *NftManagerFilterer) ParseApproval(log types.Log) (*NftManagerApproval, error) {
	event := new(NftManagerApproval)
	if err := _NftManager.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftManagerApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the NftManager contract.
type NftManagerApprovalForAllIterator struct {
	Event *NftManagerApprovalForAll // Event containing the contract specifics and raw log

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
func (it *NftManagerApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftManagerApprovalForAll)
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
		it.Event = new(NftManagerApprovalForAll)
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
func (it *NftManagerApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftManagerApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftManagerApprovalForAll represents a ApprovalForAll event raised by the NftManager contract.
type NftManagerApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_NftManager *NftManagerFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*NftManagerApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _NftManager.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &NftManagerApprovalForAllIterator{contract: _NftManager.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_NftManager *NftManagerFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *NftManagerApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _NftManager.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftManagerApprovalForAll)
				if err := _NftManager.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_NftManager *NftManagerFilterer) ParseApprovalForAll(log types.Log) (*NftManagerApprovalForAll, error) {
	event := new(NftManagerApprovalForAll)
	if err := _NftManager.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftManagerBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the NftManager contract.
type NftManagerBatchMetadataUpdateIterator struct {
	Event *NftManagerBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *NftManagerBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftManagerBatchMetadataUpdate)
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
		it.Event = new(NftManagerBatchMetadataUpdate)
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
func (it *NftManagerBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftManagerBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftManagerBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the NftManager contract.
type NftManagerBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_NftManager *NftManagerFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*NftManagerBatchMetadataUpdateIterator, error) {

	logs, sub, err := _NftManager.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &NftManagerBatchMetadataUpdateIterator{contract: _NftManager.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_NftManager *NftManagerFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *NftManagerBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _NftManager.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftManagerBatchMetadataUpdate)
				if err := _NftManager.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_NftManager *NftManagerFilterer) ParseBatchMetadataUpdate(log types.Log) (*NftManagerBatchMetadataUpdate, error) {
	event := new(NftManagerBatchMetadataUpdate)
	if err := _NftManager.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftManagerCreateNFTIterator is returned from FilterCreateNFT and is used to iterate over the raw logs and unpacked data for CreateNFT events raised by the NftManager contract.
type NftManagerCreateNFTIterator struct {
	Event *NftManagerCreateNFT // Event containing the contract specifics and raw log

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
func (it *NftManagerCreateNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftManagerCreateNFT)
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
		it.Event = new(NftManagerCreateNFT)
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
func (it *NftManagerCreateNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftManagerCreateNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftManagerCreateNFT represents a CreateNFT event raised by the NftManager contract.
type NftManagerCreateNFT struct {
	Creator         common.Address
	TokenId         *big.Int
	BusinessName    string
	Description     string
	ImgUrl          string
	BusinessAddress string
	WebSite         string
	Social          string
	Value           *big.Int
	Deadline        *big.Int
	Type            uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCreateNFT is a free log retrieval operation binding the contract event 0x2f07f1d4b070b2c0047f7ddd468bfd61a682b1ad5553c708a1fa6023812e0888.
//
// Solidity: event CreateNFT(address indexed creator, uint256 _tokenId, string _businessName, string _description, string _imgUrl, string _businessAddress, string _webSite, string _social, uint256 _value, uint256 _deadline, uint8 _type)
func (_NftManager *NftManagerFilterer) FilterCreateNFT(opts *bind.FilterOpts, creator []common.Address) (*NftManagerCreateNFTIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _NftManager.contract.FilterLogs(opts, "CreateNFT", creatorRule)
	if err != nil {
		return nil, err
	}
	return &NftManagerCreateNFTIterator{contract: _NftManager.contract, event: "CreateNFT", logs: logs, sub: sub}, nil
}

// WatchCreateNFT is a free log subscription operation binding the contract event 0x2f07f1d4b070b2c0047f7ddd468bfd61a682b1ad5553c708a1fa6023812e0888.
//
// Solidity: event CreateNFT(address indexed creator, uint256 _tokenId, string _businessName, string _description, string _imgUrl, string _businessAddress, string _webSite, string _social, uint256 _value, uint256 _deadline, uint8 _type)
func (_NftManager *NftManagerFilterer) WatchCreateNFT(opts *bind.WatchOpts, sink chan<- *NftManagerCreateNFT, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _NftManager.contract.WatchLogs(opts, "CreateNFT", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftManagerCreateNFT)
				if err := _NftManager.contract.UnpackLog(event, "CreateNFT", log); err != nil {
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

// ParseCreateNFT is a log parse operation binding the contract event 0x2f07f1d4b070b2c0047f7ddd468bfd61a682b1ad5553c708a1fa6023812e0888.
//
// Solidity: event CreateNFT(address indexed creator, uint256 _tokenId, string _businessName, string _description, string _imgUrl, string _businessAddress, string _webSite, string _social, uint256 _value, uint256 _deadline, uint8 _type)
func (_NftManager *NftManagerFilterer) ParseCreateNFT(log types.Log) (*NftManagerCreateNFT, error) {
	event := new(NftManagerCreateNFT)
	if err := _NftManager.contract.UnpackLog(event, "CreateNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NftManager contract.
type NftManagerInitializedIterator struct {
	Event *NftManagerInitialized // Event containing the contract specifics and raw log

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
func (it *NftManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftManagerInitialized)
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
		it.Event = new(NftManagerInitialized)
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
func (it *NftManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftManagerInitialized represents a Initialized event raised by the NftManager contract.
type NftManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NftManager *NftManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*NftManagerInitializedIterator, error) {

	logs, sub, err := _NftManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NftManagerInitializedIterator{contract: _NftManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NftManager *NftManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NftManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _NftManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftManagerInitialized)
				if err := _NftManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NftManager *NftManagerFilterer) ParseInitialized(log types.Log) (*NftManagerInitialized, error) {
	event := new(NftManagerInitialized)
	if err := _NftManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftManagerMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the NftManager contract.
type NftManagerMetadataUpdateIterator struct {
	Event *NftManagerMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *NftManagerMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftManagerMetadataUpdate)
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
		it.Event = new(NftManagerMetadataUpdate)
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
func (it *NftManagerMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftManagerMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftManagerMetadataUpdate represents a MetadataUpdate event raised by the NftManager contract.
type NftManagerMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_NftManager *NftManagerFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*NftManagerMetadataUpdateIterator, error) {

	logs, sub, err := _NftManager.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &NftManagerMetadataUpdateIterator{contract: _NftManager.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_NftManager *NftManagerFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *NftManagerMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _NftManager.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftManagerMetadataUpdate)
				if err := _NftManager.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_NftManager *NftManagerFilterer) ParseMetadataUpdate(log types.Log) (*NftManagerMetadataUpdate, error) {
	event := new(NftManagerMetadataUpdate)
	if err := _NftManager.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NftManager contract.
type NftManagerOwnershipTransferredIterator struct {
	Event *NftManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NftManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftManagerOwnershipTransferred)
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
		it.Event = new(NftManagerOwnershipTransferred)
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
func (it *NftManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftManagerOwnershipTransferred represents a OwnershipTransferred event raised by the NftManager contract.
type NftManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NftManager *NftManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NftManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NftManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NftManagerOwnershipTransferredIterator{contract: _NftManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NftManager *NftManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NftManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NftManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftManagerOwnershipTransferred)
				if err := _NftManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NftManager *NftManagerFilterer) ParseOwnershipTransferred(log types.Log) (*NftManagerOwnershipTransferred, error) {
	event := new(NftManagerOwnershipTransferred)
	if err := _NftManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftManagerReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the NftManager contract.
type NftManagerReceivedIterator struct {
	Event *NftManagerReceived // Event containing the contract specifics and raw log

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
func (it *NftManagerReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftManagerReceived)
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
		it.Event = new(NftManagerReceived)
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
func (it *NftManagerReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftManagerReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftManagerReceived represents a Received event raised by the NftManager contract.
type NftManagerReceived struct {
	Receiver common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed receiver, uint256 _value)
func (_NftManager *NftManagerFilterer) FilterReceived(opts *bind.FilterOpts, receiver []common.Address) (*NftManagerReceivedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _NftManager.contract.FilterLogs(opts, "Received", receiverRule)
	if err != nil {
		return nil, err
	}
	return &NftManagerReceivedIterator{contract: _NftManager.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed receiver, uint256 _value)
func (_NftManager *NftManagerFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *NftManagerReceived, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _NftManager.contract.WatchLogs(opts, "Received", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftManagerReceived)
				if err := _NftManager.contract.UnpackLog(event, "Received", log); err != nil {
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

// ParseReceived is a log parse operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed receiver, uint256 _value)
func (_NftManager *NftManagerFilterer) ParseReceived(log types.Log) (*NftManagerReceived, error) {
	event := new(NftManagerReceived)
	if err := _NftManager.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftManagerSetValidTimeIterator is returned from FilterSetValidTime and is used to iterate over the raw logs and unpacked data for SetValidTime events raised by the NftManager contract.
type NftManagerSetValidTimeIterator struct {
	Event *NftManagerSetValidTime // Event containing the contract specifics and raw log

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
func (it *NftManagerSetValidTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftManagerSetValidTime)
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
		it.Event = new(NftManagerSetValidTime)
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
func (it *NftManagerSetValidTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftManagerSetValidTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftManagerSetValidTime represents a SetValidTime event raised by the NftManager contract.
type NftManagerSetValidTime struct {
	Setter common.Address
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetValidTime is a free log retrieval operation binding the contract event 0x8651cd625889c15ab8587456aba8cc25b781c5fa977a1811f04967c2acad02b8.
//
// Solidity: event SetValidTime(address indexed setter, uint256 _time)
func (_NftManager *NftManagerFilterer) FilterSetValidTime(opts *bind.FilterOpts, setter []common.Address) (*NftManagerSetValidTimeIterator, error) {

	var setterRule []interface{}
	for _, setterItem := range setter {
		setterRule = append(setterRule, setterItem)
	}

	logs, sub, err := _NftManager.contract.FilterLogs(opts, "SetValidTime", setterRule)
	if err != nil {
		return nil, err
	}
	return &NftManagerSetValidTimeIterator{contract: _NftManager.contract, event: "SetValidTime", logs: logs, sub: sub}, nil
}

// WatchSetValidTime is a free log subscription operation binding the contract event 0x8651cd625889c15ab8587456aba8cc25b781c5fa977a1811f04967c2acad02b8.
//
// Solidity: event SetValidTime(address indexed setter, uint256 _time)
func (_NftManager *NftManagerFilterer) WatchSetValidTime(opts *bind.WatchOpts, sink chan<- *NftManagerSetValidTime, setter []common.Address) (event.Subscription, error) {

	var setterRule []interface{}
	for _, setterItem := range setter {
		setterRule = append(setterRule, setterItem)
	}

	logs, sub, err := _NftManager.contract.WatchLogs(opts, "SetValidTime", setterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftManagerSetValidTime)
				if err := _NftManager.contract.UnpackLog(event, "SetValidTime", log); err != nil {
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

// ParseSetValidTime is a log parse operation binding the contract event 0x8651cd625889c15ab8587456aba8cc25b781c5fa977a1811f04967c2acad02b8.
//
// Solidity: event SetValidTime(address indexed setter, uint256 _time)
func (_NftManager *NftManagerFilterer) ParseSetValidTime(log types.Log) (*NftManagerSetValidTime, error) {
	event := new(NftManagerSetValidTime)
	if err := _NftManager.contract.UnpackLog(event, "SetValidTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftManagerSetValuesIterator is returned from FilterSetValues and is used to iterate over the raw logs and unpacked data for SetValues events raised by the NftManager contract.
type NftManagerSetValuesIterator struct {
	Event *NftManagerSetValues // Event containing the contract specifics and raw log

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
func (it *NftManagerSetValuesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftManagerSetValues)
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
		it.Event = new(NftManagerSetValues)
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
func (it *NftManagerSetValuesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftManagerSetValuesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftManagerSetValues represents a SetValues event raised by the NftManager contract.
type NftManagerSetValues struct {
	SetterAddress common.Address
	MerchantValue *big.Int
	UserValue     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetValues is a free log retrieval operation binding the contract event 0x606bc42962399672a7f4089c08b1af4d207c5ff123b0c5e3e35aa03decb13fb8.
//
// Solidity: event SetValues(address indexed _setterAddress, uint256 _merchantValue, uint256 _userValue)
func (_NftManager *NftManagerFilterer) FilterSetValues(opts *bind.FilterOpts, _setterAddress []common.Address) (*NftManagerSetValuesIterator, error) {

	var _setterAddressRule []interface{}
	for _, _setterAddressItem := range _setterAddress {
		_setterAddressRule = append(_setterAddressRule, _setterAddressItem)
	}

	logs, sub, err := _NftManager.contract.FilterLogs(opts, "SetValues", _setterAddressRule)
	if err != nil {
		return nil, err
	}
	return &NftManagerSetValuesIterator{contract: _NftManager.contract, event: "SetValues", logs: logs, sub: sub}, nil
}

// WatchSetValues is a free log subscription operation binding the contract event 0x606bc42962399672a7f4089c08b1af4d207c5ff123b0c5e3e35aa03decb13fb8.
//
// Solidity: event SetValues(address indexed _setterAddress, uint256 _merchantValue, uint256 _userValue)
func (_NftManager *NftManagerFilterer) WatchSetValues(opts *bind.WatchOpts, sink chan<- *NftManagerSetValues, _setterAddress []common.Address) (event.Subscription, error) {

	var _setterAddressRule []interface{}
	for _, _setterAddressItem := range _setterAddress {
		_setterAddressRule = append(_setterAddressRule, _setterAddressItem)
	}

	logs, sub, err := _NftManager.contract.WatchLogs(opts, "SetValues", _setterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftManagerSetValues)
				if err := _NftManager.contract.UnpackLog(event, "SetValues", log); err != nil {
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

// ParseSetValues is a log parse operation binding the contract event 0x606bc42962399672a7f4089c08b1af4d207c5ff123b0c5e3e35aa03decb13fb8.
//
// Solidity: event SetValues(address indexed _setterAddress, uint256 _merchantValue, uint256 _userValue)
func (_NftManager *NftManagerFilterer) ParseSetValues(log types.Log) (*NftManagerSetValues, error) {
	event := new(NftManagerSetValues)
	if err := _NftManager.contract.UnpackLog(event, "SetValues", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftManagerTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the NftManager contract.
type NftManagerTransferIterator struct {
	Event *NftManagerTransfer // Event containing the contract specifics and raw log

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
func (it *NftManagerTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftManagerTransfer)
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
		it.Event = new(NftManagerTransfer)
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
func (it *NftManagerTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftManagerTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftManagerTransfer represents a Transfer event raised by the NftManager contract.
type NftManagerTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NftManager *NftManagerFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*NftManagerTransferIterator, error) {

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

	logs, sub, err := _NftManager.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftManagerTransferIterator{contract: _NftManager.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NftManager *NftManagerFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NftManagerTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _NftManager.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftManagerTransfer)
				if err := _NftManager.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_NftManager *NftManagerFilterer) ParseTransfer(log types.Log) (*NftManagerTransfer, error) {
	event := new(NftManagerTransfer)
	if err := _NftManager.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftManagerUriPrefixSetIterator is returned from FilterUriPrefixSet and is used to iterate over the raw logs and unpacked data for UriPrefixSet events raised by the NftManager contract.
type NftManagerUriPrefixSetIterator struct {
	Event *NftManagerUriPrefixSet // Event containing the contract specifics and raw log

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
func (it *NftManagerUriPrefixSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftManagerUriPrefixSet)
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
		it.Event = new(NftManagerUriPrefixSet)
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
func (it *NftManagerUriPrefixSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftManagerUriPrefixSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftManagerUriPrefixSet represents a UriPrefixSet event raised by the NftManager contract.
type NftManagerUriPrefixSet struct {
	SetterAddress common.Address
	UrlPrefix     string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUriPrefixSet is a free log retrieval operation binding the contract event 0x8aed2472df7eda2618a927a1e73256d01e814a5b03b0fff4931cc2774e39b0c6.
//
// Solidity: event UriPrefixSet(address indexed setterAddress, string urlPrefix)
func (_NftManager *NftManagerFilterer) FilterUriPrefixSet(opts *bind.FilterOpts, setterAddress []common.Address) (*NftManagerUriPrefixSetIterator, error) {

	var setterAddressRule []interface{}
	for _, setterAddressItem := range setterAddress {
		setterAddressRule = append(setterAddressRule, setterAddressItem)
	}

	logs, sub, err := _NftManager.contract.FilterLogs(opts, "UriPrefixSet", setterAddressRule)
	if err != nil {
		return nil, err
	}
	return &NftManagerUriPrefixSetIterator{contract: _NftManager.contract, event: "UriPrefixSet", logs: logs, sub: sub}, nil
}

// WatchUriPrefixSet is a free log subscription operation binding the contract event 0x8aed2472df7eda2618a927a1e73256d01e814a5b03b0fff4931cc2774e39b0c6.
//
// Solidity: event UriPrefixSet(address indexed setterAddress, string urlPrefix)
func (_NftManager *NftManagerFilterer) WatchUriPrefixSet(opts *bind.WatchOpts, sink chan<- *NftManagerUriPrefixSet, setterAddress []common.Address) (event.Subscription, error) {

	var setterAddressRule []interface{}
	for _, setterAddressItem := range setterAddress {
		setterAddressRule = append(setterAddressRule, setterAddressItem)
	}

	logs, sub, err := _NftManager.contract.WatchLogs(opts, "UriPrefixSet", setterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftManagerUriPrefixSet)
				if err := _NftManager.contract.UnpackLog(event, "UriPrefixSet", log); err != nil {
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

// ParseUriPrefixSet is a log parse operation binding the contract event 0x8aed2472df7eda2618a927a1e73256d01e814a5b03b0fff4931cc2774e39b0c6.
//
// Solidity: event UriPrefixSet(address indexed setterAddress, string urlPrefix)
func (_NftManager *NftManagerFilterer) ParseUriPrefixSet(log types.Log) (*NftManagerUriPrefixSet, error) {
	event := new(NftManagerUriPrefixSet)
	if err := _NftManager.contract.UnpackLog(event, "UriPrefixSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftManagerWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the NftManager contract.
type NftManagerWithdrawIterator struct {
	Event *NftManagerWithdraw // Event containing the contract specifics and raw log

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
func (it *NftManagerWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftManagerWithdraw)
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
		it.Event = new(NftManagerWithdraw)
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
func (it *NftManagerWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftManagerWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftManagerWithdraw represents a Withdraw event raised by the NftManager contract.
type NftManagerWithdraw struct {
	Withdrawer common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed withdrawer, uint256 _amount)
func (_NftManager *NftManagerFilterer) FilterWithdraw(opts *bind.FilterOpts, withdrawer []common.Address) (*NftManagerWithdrawIterator, error) {

	var withdrawerRule []interface{}
	for _, withdrawerItem := range withdrawer {
		withdrawerRule = append(withdrawerRule, withdrawerItem)
	}

	logs, sub, err := _NftManager.contract.FilterLogs(opts, "Withdraw", withdrawerRule)
	if err != nil {
		return nil, err
	}
	return &NftManagerWithdrawIterator{contract: _NftManager.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed withdrawer, uint256 _amount)
func (_NftManager *NftManagerFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *NftManagerWithdraw, withdrawer []common.Address) (event.Subscription, error) {

	var withdrawerRule []interface{}
	for _, withdrawerItem := range withdrawer {
		withdrawerRule = append(withdrawerRule, withdrawerItem)
	}

	logs, sub, err := _NftManager.contract.WatchLogs(opts, "Withdraw", withdrawerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftManagerWithdraw)
				if err := _NftManager.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed withdrawer, uint256 _amount)
func (_NftManager *NftManagerFilterer) ParseWithdraw(log types.Log) (*NftManagerWithdraw, error) {
	event := new(NftManagerWithdraw)
	if err := _NftManager.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftManagerWithdrawUTokenIterator is returned from FilterWithdrawUToken and is used to iterate over the raw logs and unpacked data for WithdrawUToken events raised by the NftManager contract.
type NftManagerWithdrawUTokenIterator struct {
	Event *NftManagerWithdrawUToken // Event containing the contract specifics and raw log

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
func (it *NftManagerWithdrawUTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftManagerWithdrawUToken)
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
		it.Event = new(NftManagerWithdrawUToken)
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
func (it *NftManagerWithdrawUTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftManagerWithdrawUTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftManagerWithdrawUToken represents a WithdrawUToken event raised by the NftManager contract.
type NftManagerWithdrawUToken struct {
	Withdrawer common.Address
	TokenAddr  common.Address
	Account    common.Address
	Value      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawUToken is a free log retrieval operation binding the contract event 0xf314360eddb258cb001ab9dd025a30d7be229f08113f1e5617a7f82a00e0e78f.
//
// Solidity: event WithdrawUToken(address indexed withdrawer, address indexed _tokenAddr, address indexed _account, uint256 _value)
func (_NftManager *NftManagerFilterer) FilterWithdrawUToken(opts *bind.FilterOpts, withdrawer []common.Address, _tokenAddr []common.Address, _account []common.Address) (*NftManagerWithdrawUTokenIterator, error) {

	var withdrawerRule []interface{}
	for _, withdrawerItem := range withdrawer {
		withdrawerRule = append(withdrawerRule, withdrawerItem)
	}
	var _tokenAddrRule []interface{}
	for _, _tokenAddrItem := range _tokenAddr {
		_tokenAddrRule = append(_tokenAddrRule, _tokenAddrItem)
	}
	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _NftManager.contract.FilterLogs(opts, "WithdrawUToken", withdrawerRule, _tokenAddrRule, _accountRule)
	if err != nil {
		return nil, err
	}
	return &NftManagerWithdrawUTokenIterator{contract: _NftManager.contract, event: "WithdrawUToken", logs: logs, sub: sub}, nil
}

// WatchWithdrawUToken is a free log subscription operation binding the contract event 0xf314360eddb258cb001ab9dd025a30d7be229f08113f1e5617a7f82a00e0e78f.
//
// Solidity: event WithdrawUToken(address indexed withdrawer, address indexed _tokenAddr, address indexed _account, uint256 _value)
func (_NftManager *NftManagerFilterer) WatchWithdrawUToken(opts *bind.WatchOpts, sink chan<- *NftManagerWithdrawUToken, withdrawer []common.Address, _tokenAddr []common.Address, _account []common.Address) (event.Subscription, error) {

	var withdrawerRule []interface{}
	for _, withdrawerItem := range withdrawer {
		withdrawerRule = append(withdrawerRule, withdrawerItem)
	}
	var _tokenAddrRule []interface{}
	for _, _tokenAddrItem := range _tokenAddr {
		_tokenAddrRule = append(_tokenAddrRule, _tokenAddrItem)
	}
	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _NftManager.contract.WatchLogs(opts, "WithdrawUToken", withdrawerRule, _tokenAddrRule, _accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftManagerWithdrawUToken)
				if err := _NftManager.contract.UnpackLog(event, "WithdrawUToken", log); err != nil {
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

// ParseWithdrawUToken is a log parse operation binding the contract event 0xf314360eddb258cb001ab9dd025a30d7be229f08113f1e5617a7f82a00e0e78f.
//
// Solidity: event WithdrawUToken(address indexed withdrawer, address indexed _tokenAddr, address indexed _account, uint256 _value)
func (_NftManager *NftManagerFilterer) ParseWithdrawUToken(log types.Log) (*NftManagerWithdrawUToken, error) {
	event := new(NftManagerWithdrawUToken)
	if err := _NftManager.contract.UnpackLog(event, "WithdrawUToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
