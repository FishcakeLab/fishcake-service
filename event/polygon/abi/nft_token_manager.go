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

// NFTManagerMetaData contains all meta data concerning the NFTManager contract.
var NFTManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_fccAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdtAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_redemptionPoolAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"FccTokenAddr\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UsdtTokenAddr\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"basicMineAmt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNFT\",\"inputs\":[{\"name\":\"_businessName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_description\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_imgUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_businessAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_webSite\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_social\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_type\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"_ret\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getApproved\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMerchantNTFDeadline\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUTokenBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserNTFDeadline\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isApprovedForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"merchantNTFDeadline\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minedAmt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nftTypeMap\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerOf\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proMineAmt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"redemptionPoolAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setApprovalForAll\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUriPrefix\",\"inputs\":[{\"name\":\"_uriPrefix\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"_ret\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setValues\",\"inputs\":[{\"name\":\"_merchantValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_userValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"_ret\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenURI\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalMineAmt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"uriPrefix\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"userNTFDeadline\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"_ret\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawUToken\",\"inputs\":[{\"name\":\"_tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"_ret\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApprovalForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CreateNFT\",\"inputs\":[{\"name\":\"who\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_businessName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"_description\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"_imgUrl\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"_businessAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"_webSite\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"_social\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"_value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_deadline\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_type\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Received\",\"inputs\":[{\"name\":\"who\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetValidTime\",\"inputs\":[{\"name\":\"who\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_time\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetValues\",\"inputs\":[{\"name\":\"who\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_merchantValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_userValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UriPrefixSet\",\"inputs\":[{\"name\":\"who\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"urlPrefix\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawUToken\",\"inputs\":[{\"name\":\"who\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_tokenAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Wthdraw\",\"inputs\":[{\"name\":\"who\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721IncorrectOwner\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InsufficientApproval\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721NonexistentToken\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// NFTManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use NFTManagerMetaData.ABI instead.
var NFTManagerABI = NFTManagerMetaData.ABI

// NFTManager is an auto generated Go binding around an Ethereum contract.
type NFTManager struct {
	NFTManagerCaller     // Read-only binding to the contract
	NFTManagerTransactor // Write-only binding to the contract
	NFTManagerFilterer   // Log filterer for contract events
}

// NFTManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type NFTManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NFTManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NFTManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NFTManagerSession struct {
	Contract     *NFTManager       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NFTManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NFTManagerCallerSession struct {
	Contract *NFTManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// NFTManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NFTManagerTransactorSession struct {
	Contract     *NFTManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// NFTManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type NFTManagerRaw struct {
	Contract *NFTManager // Generic contract binding to access the raw methods on
}

// NFTManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NFTManagerCallerRaw struct {
	Contract *NFTManagerCaller // Generic read-only contract binding to access the raw methods on
}

// NFTManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NFTManagerTransactorRaw struct {
	Contract *NFTManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNFTManager creates a new instance of NFTManager, bound to a specific deployed contract.
func NewNFTManager(address common.Address, backend bind.ContractBackend) (*NFTManager, error) {
	contract, err := bindNFTManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NFTManager{NFTManagerCaller: NFTManagerCaller{contract: contract}, NFTManagerTransactor: NFTManagerTransactor{contract: contract}, NFTManagerFilterer: NFTManagerFilterer{contract: contract}}, nil
}

// NewNFTManagerCaller creates a new read-only instance of NFTManager, bound to a specific deployed contract.
func NewNFTManagerCaller(address common.Address, caller bind.ContractCaller) (*NFTManagerCaller, error) {
	contract, err := bindNFTManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NFTManagerCaller{contract: contract}, nil
}

// NewNFTManagerTransactor creates a new write-only instance of NFTManager, bound to a specific deployed contract.
func NewNFTManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*NFTManagerTransactor, error) {
	contract, err := bindNFTManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NFTManagerTransactor{contract: contract}, nil
}

// NewNFTManagerFilterer creates a new log filterer instance of NFTManager, bound to a specific deployed contract.
func NewNFTManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*NFTManagerFilterer, error) {
	contract, err := bindNFTManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NFTManagerFilterer{contract: contract}, nil
}

// bindNFTManager binds a generic wrapper to an already deployed contract.
func bindNFTManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NFTManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFTManager *NFTManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFTManager.Contract.NFTManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFTManager *NFTManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTManager.Contract.NFTManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFTManager *NFTManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFTManager.Contract.NFTManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFTManager *NFTManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFTManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFTManager *NFTManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFTManager *NFTManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFTManager.Contract.contract.Transact(opts, method, params...)
}

// FccTokenAddr is a free data retrieval call binding the contract method 0x4ba1332f.
//
// Solidity: function FccTokenAddr() view returns(address)
func (_NFTManager *NFTManagerCaller) FccTokenAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NFTManager.contract.Call(opts, &out, "FccTokenAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FccTokenAddr is a free data retrieval call binding the contract method 0x4ba1332f.
//
// Solidity: function FccTokenAddr() view returns(address)
func (_NFTManager *NFTManagerSession) FccTokenAddr() (common.Address, error) {
	return _NFTManager.Contract.FccTokenAddr(&_NFTManager.CallOpts)
}

// FccTokenAddr is a free data retrieval call binding the contract method 0x4ba1332f.
//
// Solidity: function FccTokenAddr() view returns(address)
func (_NFTManager *NFTManagerCallerSession) FccTokenAddr() (common.Address, error) {
	return _NFTManager.Contract.FccTokenAddr(&_NFTManager.CallOpts)
}

// UsdtTokenAddr is a free data retrieval call binding the contract method 0xb6d18c8c.
//
// Solidity: function UsdtTokenAddr() view returns(address)
func (_NFTManager *NFTManagerCaller) UsdtTokenAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NFTManager.contract.Call(opts, &out, "UsdtTokenAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdtTokenAddr is a free data retrieval call binding the contract method 0xb6d18c8c.
//
// Solidity: function UsdtTokenAddr() view returns(address)
func (_NFTManager *NFTManagerSession) UsdtTokenAddr() (common.Address, error) {
	return _NFTManager.Contract.UsdtTokenAddr(&_NFTManager.CallOpts)
}

// UsdtTokenAddr is a free data retrieval call binding the contract method 0xb6d18c8c.
//
// Solidity: function UsdtTokenAddr() view returns(address)
func (_NFTManager *NFTManagerCallerSession) UsdtTokenAddr() (common.Address, error) {
	return _NFTManager.Contract.UsdtTokenAddr(&_NFTManager.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NFTManager *NFTManagerCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NFTManager.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NFTManager *NFTManagerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _NFTManager.Contract.BalanceOf(&_NFTManager.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NFTManager *NFTManagerCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _NFTManager.Contract.BalanceOf(&_NFTManager.CallOpts, owner)
}

// BasicMineAmt is a free data retrieval call binding the contract method 0x94399817.
//
// Solidity: function basicMineAmt() view returns(uint256)
func (_NFTManager *NFTManagerCaller) BasicMineAmt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTManager.contract.Call(opts, &out, "basicMineAmt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BasicMineAmt is a free data retrieval call binding the contract method 0x94399817.
//
// Solidity: function basicMineAmt() view returns(uint256)
func (_NFTManager *NFTManagerSession) BasicMineAmt() (*big.Int, error) {
	return _NFTManager.Contract.BasicMineAmt(&_NFTManager.CallOpts)
}

// BasicMineAmt is a free data retrieval call binding the contract method 0x94399817.
//
// Solidity: function basicMineAmt() view returns(uint256)
func (_NFTManager *NFTManagerCallerSession) BasicMineAmt() (*big.Int, error) {
	return _NFTManager.Contract.BasicMineAmt(&_NFTManager.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NFTManager *NFTManagerCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NFTManager.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NFTManager *NFTManagerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _NFTManager.Contract.GetApproved(&_NFTManager.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NFTManager *NFTManagerCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _NFTManager.Contract.GetApproved(&_NFTManager.CallOpts, tokenId)
}

// GetMerchantNTFDeadline is a free data retrieval call binding the contract method 0x39795078.
//
// Solidity: function getMerchantNTFDeadline(address _account) view returns(uint256)
func (_NFTManager *NFTManagerCaller) GetMerchantNTFDeadline(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NFTManager.contract.Call(opts, &out, "getMerchantNTFDeadline", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMerchantNTFDeadline is a free data retrieval call binding the contract method 0x39795078.
//
// Solidity: function getMerchantNTFDeadline(address _account) view returns(uint256)
func (_NFTManager *NFTManagerSession) GetMerchantNTFDeadline(_account common.Address) (*big.Int, error) {
	return _NFTManager.Contract.GetMerchantNTFDeadline(&_NFTManager.CallOpts, _account)
}

// GetMerchantNTFDeadline is a free data retrieval call binding the contract method 0x39795078.
//
// Solidity: function getMerchantNTFDeadline(address _account) view returns(uint256)
func (_NFTManager *NFTManagerCallerSession) GetMerchantNTFDeadline(_account common.Address) (*big.Int, error) {
	return _NFTManager.Contract.GetMerchantNTFDeadline(&_NFTManager.CallOpts, _account)
}

// GetUTokenBalance is a free data retrieval call binding the contract method 0x96e6c5ee.
//
// Solidity: function getUTokenBalance() view returns(uint256)
func (_NFTManager *NFTManagerCaller) GetUTokenBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTManager.contract.Call(opts, &out, "getUTokenBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUTokenBalance is a free data retrieval call binding the contract method 0x96e6c5ee.
//
// Solidity: function getUTokenBalance() view returns(uint256)
func (_NFTManager *NFTManagerSession) GetUTokenBalance() (*big.Int, error) {
	return _NFTManager.Contract.GetUTokenBalance(&_NFTManager.CallOpts)
}

// GetUTokenBalance is a free data retrieval call binding the contract method 0x96e6c5ee.
//
// Solidity: function getUTokenBalance() view returns(uint256)
func (_NFTManager *NFTManagerCallerSession) GetUTokenBalance() (*big.Int, error) {
	return _NFTManager.Contract.GetUTokenBalance(&_NFTManager.CallOpts)
}

// GetUserNTFDeadline is a free data retrieval call binding the contract method 0x8dce8b13.
//
// Solidity: function getUserNTFDeadline(address _account) view returns(uint256)
func (_NFTManager *NFTManagerCaller) GetUserNTFDeadline(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NFTManager.contract.Call(opts, &out, "getUserNTFDeadline", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserNTFDeadline is a free data retrieval call binding the contract method 0x8dce8b13.
//
// Solidity: function getUserNTFDeadline(address _account) view returns(uint256)
func (_NFTManager *NFTManagerSession) GetUserNTFDeadline(_account common.Address) (*big.Int, error) {
	return _NFTManager.Contract.GetUserNTFDeadline(&_NFTManager.CallOpts, _account)
}

// GetUserNTFDeadline is a free data retrieval call binding the contract method 0x8dce8b13.
//
// Solidity: function getUserNTFDeadline(address _account) view returns(uint256)
func (_NFTManager *NFTManagerCallerSession) GetUserNTFDeadline(_account common.Address) (*big.Int, error) {
	return _NFTManager.Contract.GetUserNTFDeadline(&_NFTManager.CallOpts, _account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NFTManager *NFTManagerCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _NFTManager.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NFTManager *NFTManagerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _NFTManager.Contract.IsApprovedForAll(&_NFTManager.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NFTManager *NFTManagerCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _NFTManager.Contract.IsApprovedForAll(&_NFTManager.CallOpts, owner, operator)
}

// MerchantNTFDeadline is a free data retrieval call binding the contract method 0xc498d1b0.
//
// Solidity: function merchantNTFDeadline(address ) view returns(uint256)
func (_NFTManager *NFTManagerCaller) MerchantNTFDeadline(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NFTManager.contract.Call(opts, &out, "merchantNTFDeadline", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MerchantNTFDeadline is a free data retrieval call binding the contract method 0xc498d1b0.
//
// Solidity: function merchantNTFDeadline(address ) view returns(uint256)
func (_NFTManager *NFTManagerSession) MerchantNTFDeadline(arg0 common.Address) (*big.Int, error) {
	return _NFTManager.Contract.MerchantNTFDeadline(&_NFTManager.CallOpts, arg0)
}

// MerchantNTFDeadline is a free data retrieval call binding the contract method 0xc498d1b0.
//
// Solidity: function merchantNTFDeadline(address ) view returns(uint256)
func (_NFTManager *NFTManagerCallerSession) MerchantNTFDeadline(arg0 common.Address) (*big.Int, error) {
	return _NFTManager.Contract.MerchantNTFDeadline(&_NFTManager.CallOpts, arg0)
}

// MinedAmt is a free data retrieval call binding the contract method 0xf7859c28.
//
// Solidity: function minedAmt() view returns(uint256)
func (_NFTManager *NFTManagerCaller) MinedAmt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTManager.contract.Call(opts, &out, "minedAmt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinedAmt is a free data retrieval call binding the contract method 0xf7859c28.
//
// Solidity: function minedAmt() view returns(uint256)
func (_NFTManager *NFTManagerSession) MinedAmt() (*big.Int, error) {
	return _NFTManager.Contract.MinedAmt(&_NFTManager.CallOpts)
}

// MinedAmt is a free data retrieval call binding the contract method 0xf7859c28.
//
// Solidity: function minedAmt() view returns(uint256)
func (_NFTManager *NFTManagerCallerSession) MinedAmt() (*big.Int, error) {
	return _NFTManager.Contract.MinedAmt(&_NFTManager.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NFTManager *NFTManagerCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NFTManager.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NFTManager *NFTManagerSession) Name() (string, error) {
	return _NFTManager.Contract.Name(&_NFTManager.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NFTManager *NFTManagerCallerSession) Name() (string, error) {
	return _NFTManager.Contract.Name(&_NFTManager.CallOpts)
}

// NftTypeMap is a free data retrieval call binding the contract method 0x78b0fbbd.
//
// Solidity: function nftTypeMap(uint256 ) view returns(uint8)
func (_NFTManager *NFTManagerCaller) NftTypeMap(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var out []interface{}
	err := _NFTManager.contract.Call(opts, &out, "nftTypeMap", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// NftTypeMap is a free data retrieval call binding the contract method 0x78b0fbbd.
//
// Solidity: function nftTypeMap(uint256 ) view returns(uint8)
func (_NFTManager *NFTManagerSession) NftTypeMap(arg0 *big.Int) (uint8, error) {
	return _NFTManager.Contract.NftTypeMap(&_NFTManager.CallOpts, arg0)
}

// NftTypeMap is a free data retrieval call binding the contract method 0x78b0fbbd.
//
// Solidity: function nftTypeMap(uint256 ) view returns(uint8)
func (_NFTManager *NFTManagerCallerSession) NftTypeMap(arg0 *big.Int) (uint8, error) {
	return _NFTManager.Contract.NftTypeMap(&_NFTManager.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFTManager *NFTManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NFTManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFTManager *NFTManagerSession) Owner() (common.Address, error) {
	return _NFTManager.Contract.Owner(&_NFTManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFTManager *NFTManagerCallerSession) Owner() (common.Address, error) {
	return _NFTManager.Contract.Owner(&_NFTManager.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NFTManager *NFTManagerCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NFTManager.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NFTManager *NFTManagerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _NFTManager.Contract.OwnerOf(&_NFTManager.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NFTManager *NFTManagerCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _NFTManager.Contract.OwnerOf(&_NFTManager.CallOpts, tokenId)
}

// ProMineAmt is a free data retrieval call binding the contract method 0x80c7fb97.
//
// Solidity: function proMineAmt() view returns(uint256)
func (_NFTManager *NFTManagerCaller) ProMineAmt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTManager.contract.Call(opts, &out, "proMineAmt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProMineAmt is a free data retrieval call binding the contract method 0x80c7fb97.
//
// Solidity: function proMineAmt() view returns(uint256)
func (_NFTManager *NFTManagerSession) ProMineAmt() (*big.Int, error) {
	return _NFTManager.Contract.ProMineAmt(&_NFTManager.CallOpts)
}

// ProMineAmt is a free data retrieval call binding the contract method 0x80c7fb97.
//
// Solidity: function proMineAmt() view returns(uint256)
func (_NFTManager *NFTManagerCallerSession) ProMineAmt() (*big.Int, error) {
	return _NFTManager.Contract.ProMineAmt(&_NFTManager.CallOpts)
}

// RedemptionPoolAddress is a free data retrieval call binding the contract method 0xb6648e47.
//
// Solidity: function redemptionPoolAddress() view returns(address)
func (_NFTManager *NFTManagerCaller) RedemptionPoolAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NFTManager.contract.Call(opts, &out, "redemptionPoolAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RedemptionPoolAddress is a free data retrieval call binding the contract method 0xb6648e47.
//
// Solidity: function redemptionPoolAddress() view returns(address)
func (_NFTManager *NFTManagerSession) RedemptionPoolAddress() (common.Address, error) {
	return _NFTManager.Contract.RedemptionPoolAddress(&_NFTManager.CallOpts)
}

// RedemptionPoolAddress is a free data retrieval call binding the contract method 0xb6648e47.
//
// Solidity: function redemptionPoolAddress() view returns(address)
func (_NFTManager *NFTManagerCallerSession) RedemptionPoolAddress() (common.Address, error) {
	return _NFTManager.Contract.RedemptionPoolAddress(&_NFTManager.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NFTManager *NFTManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _NFTManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NFTManager *NFTManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NFTManager.Contract.SupportsInterface(&_NFTManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NFTManager *NFTManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NFTManager.Contract.SupportsInterface(&_NFTManager.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NFTManager *NFTManagerCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NFTManager.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NFTManager *NFTManagerSession) Symbol() (string, error) {
	return _NFTManager.Contract.Symbol(&_NFTManager.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NFTManager *NFTManagerCallerSession) Symbol() (string, error) {
	return _NFTManager.Contract.Symbol(&_NFTManager.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_NFTManager *NFTManagerCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _NFTManager.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_NFTManager *NFTManagerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _NFTManager.Contract.TokenURI(&_NFTManager.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_NFTManager *NFTManagerCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _NFTManager.Contract.TokenURI(&_NFTManager.CallOpts, tokenId)
}

// TotalMineAmt is a free data retrieval call binding the contract method 0x88e40a5a.
//
// Solidity: function totalMineAmt() view returns(uint256)
func (_NFTManager *NFTManagerCaller) TotalMineAmt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTManager.contract.Call(opts, &out, "totalMineAmt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMineAmt is a free data retrieval call binding the contract method 0x88e40a5a.
//
// Solidity: function totalMineAmt() view returns(uint256)
func (_NFTManager *NFTManagerSession) TotalMineAmt() (*big.Int, error) {
	return _NFTManager.Contract.TotalMineAmt(&_NFTManager.CallOpts)
}

// TotalMineAmt is a free data retrieval call binding the contract method 0x88e40a5a.
//
// Solidity: function totalMineAmt() view returns(uint256)
func (_NFTManager *NFTManagerCallerSession) TotalMineAmt() (*big.Int, error) {
	return _NFTManager.Contract.TotalMineAmt(&_NFTManager.CallOpts)
}

// UriPrefix is a free data retrieval call binding the contract method 0x62b99ad4.
//
// Solidity: function uriPrefix() view returns(string)
func (_NFTManager *NFTManagerCaller) UriPrefix(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NFTManager.contract.Call(opts, &out, "uriPrefix")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UriPrefix is a free data retrieval call binding the contract method 0x62b99ad4.
//
// Solidity: function uriPrefix() view returns(string)
func (_NFTManager *NFTManagerSession) UriPrefix() (string, error) {
	return _NFTManager.Contract.UriPrefix(&_NFTManager.CallOpts)
}

// UriPrefix is a free data retrieval call binding the contract method 0x62b99ad4.
//
// Solidity: function uriPrefix() view returns(string)
func (_NFTManager *NFTManagerCallerSession) UriPrefix() (string, error) {
	return _NFTManager.Contract.UriPrefix(&_NFTManager.CallOpts)
}

// UserNTFDeadline is a free data retrieval call binding the contract method 0x121e3012.
//
// Solidity: function userNTFDeadline(address ) view returns(uint256)
func (_NFTManager *NFTManagerCaller) UserNTFDeadline(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NFTManager.contract.Call(opts, &out, "userNTFDeadline", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserNTFDeadline is a free data retrieval call binding the contract method 0x121e3012.
//
// Solidity: function userNTFDeadline(address ) view returns(uint256)
func (_NFTManager *NFTManagerSession) UserNTFDeadline(arg0 common.Address) (*big.Int, error) {
	return _NFTManager.Contract.UserNTFDeadline(&_NFTManager.CallOpts, arg0)
}

// UserNTFDeadline is a free data retrieval call binding the contract method 0x121e3012.
//
// Solidity: function userNTFDeadline(address ) view returns(uint256)
func (_NFTManager *NFTManagerCallerSession) UserNTFDeadline(arg0 common.Address) (*big.Int, error) {
	return _NFTManager.Contract.UserNTFDeadline(&_NFTManager.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_NFTManager *NFTManagerTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTManager.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_NFTManager *NFTManagerSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTManager.Contract.Approve(&_NFTManager.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_NFTManager *NFTManagerTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTManager.Contract.Approve(&_NFTManager.TransactOpts, to, tokenId)
}

// CreateNFT is a paid mutator transaction binding the contract method 0x463d5059.
//
// Solidity: function createNFT(string _businessName, string _description, string _imgUrl, string _businessAddress, string _webSite, string _social, uint8 _type) returns(bool _ret, uint256 _tokenId)
func (_NFTManager *NFTManagerTransactor) CreateNFT(opts *bind.TransactOpts, _businessName string, _description string, _imgUrl string, _businessAddress string, _webSite string, _social string, _type uint8) (*types.Transaction, error) {
	return _NFTManager.contract.Transact(opts, "createNFT", _businessName, _description, _imgUrl, _businessAddress, _webSite, _social, _type)
}

// CreateNFT is a paid mutator transaction binding the contract method 0x463d5059.
//
// Solidity: function createNFT(string _businessName, string _description, string _imgUrl, string _businessAddress, string _webSite, string _social, uint8 _type) returns(bool _ret, uint256 _tokenId)
func (_NFTManager *NFTManagerSession) CreateNFT(_businessName string, _description string, _imgUrl string, _businessAddress string, _webSite string, _social string, _type uint8) (*types.Transaction, error) {
	return _NFTManager.Contract.CreateNFT(&_NFTManager.TransactOpts, _businessName, _description, _imgUrl, _businessAddress, _webSite, _social, _type)
}

// CreateNFT is a paid mutator transaction binding the contract method 0x463d5059.
//
// Solidity: function createNFT(string _businessName, string _description, string _imgUrl, string _businessAddress, string _webSite, string _social, uint8 _type) returns(bool _ret, uint256 _tokenId)
func (_NFTManager *NFTManagerTransactorSession) CreateNFT(_businessName string, _description string, _imgUrl string, _businessAddress string, _webSite string, _social string, _type uint8) (*types.Transaction, error) {
	return _NFTManager.Contract.CreateNFT(&_NFTManager.TransactOpts, _businessName, _description, _imgUrl, _businessAddress, _webSite, _social, _type)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NFTManager *NFTManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NFTManager *NFTManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _NFTManager.Contract.RenounceOwnership(&_NFTManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NFTManager *NFTManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NFTManager.Contract.RenounceOwnership(&_NFTManager.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_NFTManager *NFTManagerTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTManager.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_NFTManager *NFTManagerSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTManager.Contract.SafeTransferFrom(&_NFTManager.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_NFTManager *NFTManagerTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTManager.Contract.SafeTransferFrom(&_NFTManager.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_NFTManager *NFTManagerTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _NFTManager.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_NFTManager *NFTManagerSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _NFTManager.Contract.SafeTransferFrom0(&_NFTManager.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_NFTManager *NFTManagerTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _NFTManager.Contract.SafeTransferFrom0(&_NFTManager.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NFTManager *NFTManagerTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _NFTManager.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NFTManager *NFTManagerSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _NFTManager.Contract.SetApprovalForAll(&_NFTManager.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NFTManager *NFTManagerTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _NFTManager.Contract.SetApprovalForAll(&_NFTManager.TransactOpts, operator, approved)
}

// SetUriPrefix is a paid mutator transaction binding the contract method 0x7ec4a659.
//
// Solidity: function setUriPrefix(string _uriPrefix) returns(bool _ret)
func (_NFTManager *NFTManagerTransactor) SetUriPrefix(opts *bind.TransactOpts, _uriPrefix string) (*types.Transaction, error) {
	return _NFTManager.contract.Transact(opts, "setUriPrefix", _uriPrefix)
}

// SetUriPrefix is a paid mutator transaction binding the contract method 0x7ec4a659.
//
// Solidity: function setUriPrefix(string _uriPrefix) returns(bool _ret)
func (_NFTManager *NFTManagerSession) SetUriPrefix(_uriPrefix string) (*types.Transaction, error) {
	return _NFTManager.Contract.SetUriPrefix(&_NFTManager.TransactOpts, _uriPrefix)
}

// SetUriPrefix is a paid mutator transaction binding the contract method 0x7ec4a659.
//
// Solidity: function setUriPrefix(string _uriPrefix) returns(bool _ret)
func (_NFTManager *NFTManagerTransactorSession) SetUriPrefix(_uriPrefix string) (*types.Transaction, error) {
	return _NFTManager.Contract.SetUriPrefix(&_NFTManager.TransactOpts, _uriPrefix)
}

// SetValues is a paid mutator transaction binding the contract method 0x9c3dfaf6.
//
// Solidity: function setValues(uint256 _merchantValue, uint256 _userValue) returns(bool _ret)
func (_NFTManager *NFTManagerTransactor) SetValues(opts *bind.TransactOpts, _merchantValue *big.Int, _userValue *big.Int) (*types.Transaction, error) {
	return _NFTManager.contract.Transact(opts, "setValues", _merchantValue, _userValue)
}

// SetValues is a paid mutator transaction binding the contract method 0x9c3dfaf6.
//
// Solidity: function setValues(uint256 _merchantValue, uint256 _userValue) returns(bool _ret)
func (_NFTManager *NFTManagerSession) SetValues(_merchantValue *big.Int, _userValue *big.Int) (*types.Transaction, error) {
	return _NFTManager.Contract.SetValues(&_NFTManager.TransactOpts, _merchantValue, _userValue)
}

// SetValues is a paid mutator transaction binding the contract method 0x9c3dfaf6.
//
// Solidity: function setValues(uint256 _merchantValue, uint256 _userValue) returns(bool _ret)
func (_NFTManager *NFTManagerTransactorSession) SetValues(_merchantValue *big.Int, _userValue *big.Int) (*types.Transaction, error) {
	return _NFTManager.Contract.SetValues(&_NFTManager.TransactOpts, _merchantValue, _userValue)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_NFTManager *NFTManagerTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTManager.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_NFTManager *NFTManagerSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTManager.Contract.TransferFrom(&_NFTManager.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_NFTManager *NFTManagerTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTManager.Contract.TransferFrom(&_NFTManager.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFTManager *NFTManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NFTManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFTManager *NFTManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NFTManager.Contract.TransferOwnership(&_NFTManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFTManager *NFTManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NFTManager.Contract.TransferOwnership(&_NFTManager.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _recipient, uint256 _amount) returns(bool _ret)
func (_NFTManager *NFTManagerTransactor) Withdraw(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _NFTManager.contract.Transact(opts, "withdraw", _recipient, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _recipient, uint256 _amount) returns(bool _ret)
func (_NFTManager *NFTManagerSession) Withdraw(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _NFTManager.Contract.Withdraw(&_NFTManager.TransactOpts, _recipient, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _recipient, uint256 _amount) returns(bool _ret)
func (_NFTManager *NFTManagerTransactorSession) Withdraw(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _NFTManager.Contract.Withdraw(&_NFTManager.TransactOpts, _recipient, _amount)
}

// WithdrawUToken is a paid mutator transaction binding the contract method 0xde3f2b22.
//
// Solidity: function withdrawUToken(address _tokenAddr, address _account, uint256 _value) returns(bool _ret)
func (_NFTManager *NFTManagerTransactor) WithdrawUToken(opts *bind.TransactOpts, _tokenAddr common.Address, _account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _NFTManager.contract.Transact(opts, "withdrawUToken", _tokenAddr, _account, _value)
}

// WithdrawUToken is a paid mutator transaction binding the contract method 0xde3f2b22.
//
// Solidity: function withdrawUToken(address _tokenAddr, address _account, uint256 _value) returns(bool _ret)
func (_NFTManager *NFTManagerSession) WithdrawUToken(_tokenAddr common.Address, _account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _NFTManager.Contract.WithdrawUToken(&_NFTManager.TransactOpts, _tokenAddr, _account, _value)
}

// WithdrawUToken is a paid mutator transaction binding the contract method 0xde3f2b22.
//
// Solidity: function withdrawUToken(address _tokenAddr, address _account, uint256 _value) returns(bool _ret)
func (_NFTManager *NFTManagerTransactorSession) WithdrawUToken(_tokenAddr common.Address, _account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _NFTManager.Contract.WithdrawUToken(&_NFTManager.TransactOpts, _tokenAddr, _account, _value)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NFTManager *NFTManagerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTManager.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NFTManager *NFTManagerSession) Receive() (*types.Transaction, error) {
	return _NFTManager.Contract.Receive(&_NFTManager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NFTManager *NFTManagerTransactorSession) Receive() (*types.Transaction, error) {
	return _NFTManager.Contract.Receive(&_NFTManager.TransactOpts)
}

// NFTManagerApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the NFTManager contract.
type NFTManagerApprovalIterator struct {
	Event *NFTManagerApproval // Event containing the contract specifics and raw log

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
func (it *NFTManagerApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTManagerApproval)
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
		it.Event = new(NFTManagerApproval)
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
func (it *NFTManagerApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTManagerApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTManagerApproval represents a Approval event raised by the NFTManager contract.
type NFTManagerApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_NFTManager *NFTManagerFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*NFTManagerApprovalIterator, error) {

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

	logs, sub, err := _NFTManager.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NFTManagerApprovalIterator{contract: _NFTManager.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_NFTManager *NFTManagerFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NFTManagerApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _NFTManager.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTManagerApproval)
				if err := _NFTManager.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_NFTManager *NFTManagerFilterer) ParseApproval(log types.Log) (*NFTManagerApproval, error) {
	event := new(NFTManagerApproval)
	if err := _NFTManager.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTManagerApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the NFTManager contract.
type NFTManagerApprovalForAllIterator struct {
	Event *NFTManagerApprovalForAll // Event containing the contract specifics and raw log

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
func (it *NFTManagerApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTManagerApprovalForAll)
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
		it.Event = new(NFTManagerApprovalForAll)
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
func (it *NFTManagerApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTManagerApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTManagerApprovalForAll represents a ApprovalForAll event raised by the NFTManager contract.
type NFTManagerApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_NFTManager *NFTManagerFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*NFTManagerApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _NFTManager.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &NFTManagerApprovalForAllIterator{contract: _NFTManager.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_NFTManager *NFTManagerFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *NFTManagerApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _NFTManager.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTManagerApprovalForAll)
				if err := _NFTManager.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_NFTManager *NFTManagerFilterer) ParseApprovalForAll(log types.Log) (*NFTManagerApprovalForAll, error) {
	event := new(NFTManagerApprovalForAll)
	if err := _NFTManager.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTManagerCreateNFTIterator is returned from FilterCreateNFT and is used to iterate over the raw logs and unpacked data for CreateNFT events raised by the NFTManager contract.
type NFTManagerCreateNFTIterator struct {
	Event *NFTManagerCreateNFT // Event containing the contract specifics and raw log

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
func (it *NFTManagerCreateNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTManagerCreateNFT)
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
		it.Event = new(NFTManagerCreateNFT)
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
func (it *NFTManagerCreateNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTManagerCreateNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTManagerCreateNFT represents a CreateNFT event raised by the NFTManager contract.
type NFTManagerCreateNFT struct {
	Who             common.Address
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
// Solidity: event CreateNFT(address indexed who, uint256 _tokenId, string _businessName, string _description, string _imgUrl, string _businessAddress, string _webSite, string _social, uint256 _value, uint256 _deadline, uint8 _type)
func (_NFTManager *NFTManagerFilterer) FilterCreateNFT(opts *bind.FilterOpts, who []common.Address) (*NFTManagerCreateNFTIterator, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _NFTManager.contract.FilterLogs(opts, "CreateNFT", whoRule)
	if err != nil {
		return nil, err
	}
	return &NFTManagerCreateNFTIterator{contract: _NFTManager.contract, event: "CreateNFT", logs: logs, sub: sub}, nil
}

// WatchCreateNFT is a free log subscription operation binding the contract event 0x2f07f1d4b070b2c0047f7ddd468bfd61a682b1ad5553c708a1fa6023812e0888.
//
// Solidity: event CreateNFT(address indexed who, uint256 _tokenId, string _businessName, string _description, string _imgUrl, string _businessAddress, string _webSite, string _social, uint256 _value, uint256 _deadline, uint8 _type)
func (_NFTManager *NFTManagerFilterer) WatchCreateNFT(opts *bind.WatchOpts, sink chan<- *NFTManagerCreateNFT, who []common.Address) (event.Subscription, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _NFTManager.contract.WatchLogs(opts, "CreateNFT", whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTManagerCreateNFT)
				if err := _NFTManager.contract.UnpackLog(event, "CreateNFT", log); err != nil {
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
// Solidity: event CreateNFT(address indexed who, uint256 _tokenId, string _businessName, string _description, string _imgUrl, string _businessAddress, string _webSite, string _social, uint256 _value, uint256 _deadline, uint8 _type)
func (_NFTManager *NFTManagerFilterer) ParseCreateNFT(log types.Log) (*NFTManagerCreateNFT, error) {
	event := new(NFTManagerCreateNFT)
	if err := _NFTManager.contract.UnpackLog(event, "CreateNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NFTManager contract.
type NFTManagerOwnershipTransferredIterator struct {
	Event *NFTManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NFTManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTManagerOwnershipTransferred)
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
		it.Event = new(NFTManagerOwnershipTransferred)
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
func (it *NFTManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTManagerOwnershipTransferred represents a OwnershipTransferred event raised by the NFTManager contract.
type NFTManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NFTManager *NFTManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NFTManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NFTManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NFTManagerOwnershipTransferredIterator{contract: _NFTManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NFTManager *NFTManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NFTManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NFTManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTManagerOwnershipTransferred)
				if err := _NFTManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NFTManager *NFTManagerFilterer) ParseOwnershipTransferred(log types.Log) (*NFTManagerOwnershipTransferred, error) {
	event := new(NFTManagerOwnershipTransferred)
	if err := _NFTManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTManagerReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the NFTManager contract.
type NFTManagerReceivedIterator struct {
	Event *NFTManagerReceived // Event containing the contract specifics and raw log

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
func (it *NFTManagerReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTManagerReceived)
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
		it.Event = new(NFTManagerReceived)
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
func (it *NFTManagerReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTManagerReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTManagerReceived represents a Received event raised by the NFTManager contract.
type NFTManagerReceived struct {
	Who   common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed who, uint256 _value)
func (_NFTManager *NFTManagerFilterer) FilterReceived(opts *bind.FilterOpts, who []common.Address) (*NFTManagerReceivedIterator, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _NFTManager.contract.FilterLogs(opts, "Received", whoRule)
	if err != nil {
		return nil, err
	}
	return &NFTManagerReceivedIterator{contract: _NFTManager.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed who, uint256 _value)
func (_NFTManager *NFTManagerFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *NFTManagerReceived, who []common.Address) (event.Subscription, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _NFTManager.contract.WatchLogs(opts, "Received", whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTManagerReceived)
				if err := _NFTManager.contract.UnpackLog(event, "Received", log); err != nil {
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
// Solidity: event Received(address indexed who, uint256 _value)
func (_NFTManager *NFTManagerFilterer) ParseReceived(log types.Log) (*NFTManagerReceived, error) {
	event := new(NFTManagerReceived)
	if err := _NFTManager.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTManagerSetValidTimeIterator is returned from FilterSetValidTime and is used to iterate over the raw logs and unpacked data for SetValidTime events raised by the NFTManager contract.
type NFTManagerSetValidTimeIterator struct {
	Event *NFTManagerSetValidTime // Event containing the contract specifics and raw log

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
func (it *NFTManagerSetValidTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTManagerSetValidTime)
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
		it.Event = new(NFTManagerSetValidTime)
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
func (it *NFTManagerSetValidTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTManagerSetValidTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTManagerSetValidTime represents a SetValidTime event raised by the NFTManager contract.
type NFTManagerSetValidTime struct {
	Who  common.Address
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetValidTime is a free log retrieval operation binding the contract event 0x8651cd625889c15ab8587456aba8cc25b781c5fa977a1811f04967c2acad02b8.
//
// Solidity: event SetValidTime(address indexed who, uint256 _time)
func (_NFTManager *NFTManagerFilterer) FilterSetValidTime(opts *bind.FilterOpts, who []common.Address) (*NFTManagerSetValidTimeIterator, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _NFTManager.contract.FilterLogs(opts, "SetValidTime", whoRule)
	if err != nil {
		return nil, err
	}
	return &NFTManagerSetValidTimeIterator{contract: _NFTManager.contract, event: "SetValidTime", logs: logs, sub: sub}, nil
}

// WatchSetValidTime is a free log subscription operation binding the contract event 0x8651cd625889c15ab8587456aba8cc25b781c5fa977a1811f04967c2acad02b8.
//
// Solidity: event SetValidTime(address indexed who, uint256 _time)
func (_NFTManager *NFTManagerFilterer) WatchSetValidTime(opts *bind.WatchOpts, sink chan<- *NFTManagerSetValidTime, who []common.Address) (event.Subscription, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _NFTManager.contract.WatchLogs(opts, "SetValidTime", whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTManagerSetValidTime)
				if err := _NFTManager.contract.UnpackLog(event, "SetValidTime", log); err != nil {
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
// Solidity: event SetValidTime(address indexed who, uint256 _time)
func (_NFTManager *NFTManagerFilterer) ParseSetValidTime(log types.Log) (*NFTManagerSetValidTime, error) {
	event := new(NFTManagerSetValidTime)
	if err := _NFTManager.contract.UnpackLog(event, "SetValidTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTManagerSetValuesIterator is returned from FilterSetValues and is used to iterate over the raw logs and unpacked data for SetValues events raised by the NFTManager contract.
type NFTManagerSetValuesIterator struct {
	Event *NFTManagerSetValues // Event containing the contract specifics and raw log

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
func (it *NFTManagerSetValuesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTManagerSetValues)
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
		it.Event = new(NFTManagerSetValues)
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
func (it *NFTManagerSetValuesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTManagerSetValuesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTManagerSetValues represents a SetValues event raised by the NFTManager contract.
type NFTManagerSetValues struct {
	Who           common.Address
	MerchantValue *big.Int
	UserValue     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetValues is a free log retrieval operation binding the contract event 0x606bc42962399672a7f4089c08b1af4d207c5ff123b0c5e3e35aa03decb13fb8.
//
// Solidity: event SetValues(address indexed who, uint256 _merchantValue, uint256 _userValue)
func (_NFTManager *NFTManagerFilterer) FilterSetValues(opts *bind.FilterOpts, who []common.Address) (*NFTManagerSetValuesIterator, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _NFTManager.contract.FilterLogs(opts, "SetValues", whoRule)
	if err != nil {
		return nil, err
	}
	return &NFTManagerSetValuesIterator{contract: _NFTManager.contract, event: "SetValues", logs: logs, sub: sub}, nil
}

// WatchSetValues is a free log subscription operation binding the contract event 0x606bc42962399672a7f4089c08b1af4d207c5ff123b0c5e3e35aa03decb13fb8.
//
// Solidity: event SetValues(address indexed who, uint256 _merchantValue, uint256 _userValue)
func (_NFTManager *NFTManagerFilterer) WatchSetValues(opts *bind.WatchOpts, sink chan<- *NFTManagerSetValues, who []common.Address) (event.Subscription, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _NFTManager.contract.WatchLogs(opts, "SetValues", whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTManagerSetValues)
				if err := _NFTManager.contract.UnpackLog(event, "SetValues", log); err != nil {
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
// Solidity: event SetValues(address indexed who, uint256 _merchantValue, uint256 _userValue)
func (_NFTManager *NFTManagerFilterer) ParseSetValues(log types.Log) (*NFTManagerSetValues, error) {
	event := new(NFTManagerSetValues)
	if err := _NFTManager.contract.UnpackLog(event, "SetValues", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTManagerTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the NFTManager contract.
type NFTManagerTransferIterator struct {
	Event *NFTManagerTransfer // Event containing the contract specifics and raw log

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
func (it *NFTManagerTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTManagerTransfer)
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
		it.Event = new(NFTManagerTransfer)
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
func (it *NFTManagerTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTManagerTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTManagerTransfer represents a Transfer event raised by the NFTManager contract.
type NFTManagerTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NFTManager *NFTManagerFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*NFTManagerTransferIterator, error) {

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

	logs, sub, err := _NFTManager.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NFTManagerTransferIterator{contract: _NFTManager.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NFTManager *NFTManagerFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NFTManagerTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _NFTManager.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTManagerTransfer)
				if err := _NFTManager.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_NFTManager *NFTManagerFilterer) ParseTransfer(log types.Log) (*NFTManagerTransfer, error) {
	event := new(NFTManagerTransfer)
	if err := _NFTManager.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTManagerUriPrefixSetIterator is returned from FilterUriPrefixSet and is used to iterate over the raw logs and unpacked data for UriPrefixSet events raised by the NFTManager contract.
type NFTManagerUriPrefixSetIterator struct {
	Event *NFTManagerUriPrefixSet // Event containing the contract specifics and raw log

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
func (it *NFTManagerUriPrefixSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTManagerUriPrefixSet)
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
		it.Event = new(NFTManagerUriPrefixSet)
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
func (it *NFTManagerUriPrefixSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTManagerUriPrefixSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTManagerUriPrefixSet represents a UriPrefixSet event raised by the NFTManager contract.
type NFTManagerUriPrefixSet struct {
	Who       common.Address
	UrlPrefix string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUriPrefixSet is a free log retrieval operation binding the contract event 0x8aed2472df7eda2618a927a1e73256d01e814a5b03b0fff4931cc2774e39b0c6.
//
// Solidity: event UriPrefixSet(address indexed who, string urlPrefix)
func (_NFTManager *NFTManagerFilterer) FilterUriPrefixSet(opts *bind.FilterOpts, who []common.Address) (*NFTManagerUriPrefixSetIterator, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _NFTManager.contract.FilterLogs(opts, "UriPrefixSet", whoRule)
	if err != nil {
		return nil, err
	}
	return &NFTManagerUriPrefixSetIterator{contract: _NFTManager.contract, event: "UriPrefixSet", logs: logs, sub: sub}, nil
}

// WatchUriPrefixSet is a free log subscription operation binding the contract event 0x8aed2472df7eda2618a927a1e73256d01e814a5b03b0fff4931cc2774e39b0c6.
//
// Solidity: event UriPrefixSet(address indexed who, string urlPrefix)
func (_NFTManager *NFTManagerFilterer) WatchUriPrefixSet(opts *bind.WatchOpts, sink chan<- *NFTManagerUriPrefixSet, who []common.Address) (event.Subscription, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _NFTManager.contract.WatchLogs(opts, "UriPrefixSet", whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTManagerUriPrefixSet)
				if err := _NFTManager.contract.UnpackLog(event, "UriPrefixSet", log); err != nil {
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
// Solidity: event UriPrefixSet(address indexed who, string urlPrefix)
func (_NFTManager *NFTManagerFilterer) ParseUriPrefixSet(log types.Log) (*NFTManagerUriPrefixSet, error) {
	event := new(NFTManagerUriPrefixSet)
	if err := _NFTManager.contract.UnpackLog(event, "UriPrefixSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTManagerWithdrawUTokenIterator is returned from FilterWithdrawUToken and is used to iterate over the raw logs and unpacked data for WithdrawUToken events raised by the NFTManager contract.
type NFTManagerWithdrawUTokenIterator struct {
	Event *NFTManagerWithdrawUToken // Event containing the contract specifics and raw log

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
func (it *NFTManagerWithdrawUTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTManagerWithdrawUToken)
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
		it.Event = new(NFTManagerWithdrawUToken)
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
func (it *NFTManagerWithdrawUTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTManagerWithdrawUTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTManagerWithdrawUToken represents a WithdrawUToken event raised by the NFTManager contract.
type NFTManagerWithdrawUToken struct {
	Who       common.Address
	TokenAddr common.Address
	Account   common.Address
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawUToken is a free log retrieval operation binding the contract event 0xf314360eddb258cb001ab9dd025a30d7be229f08113f1e5617a7f82a00e0e78f.
//
// Solidity: event WithdrawUToken(address indexed who, address indexed _tokenAddr, address indexed _account, uint256 _value)
func (_NFTManager *NFTManagerFilterer) FilterWithdrawUToken(opts *bind.FilterOpts, who []common.Address, _tokenAddr []common.Address, _account []common.Address) (*NFTManagerWithdrawUTokenIterator, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}
	var _tokenAddrRule []interface{}
	for _, _tokenAddrItem := range _tokenAddr {
		_tokenAddrRule = append(_tokenAddrRule, _tokenAddrItem)
	}
	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _NFTManager.contract.FilterLogs(opts, "WithdrawUToken", whoRule, _tokenAddrRule, _accountRule)
	if err != nil {
		return nil, err
	}
	return &NFTManagerWithdrawUTokenIterator{contract: _NFTManager.contract, event: "WithdrawUToken", logs: logs, sub: sub}, nil
}

// WatchWithdrawUToken is a free log subscription operation binding the contract event 0xf314360eddb258cb001ab9dd025a30d7be229f08113f1e5617a7f82a00e0e78f.
//
// Solidity: event WithdrawUToken(address indexed who, address indexed _tokenAddr, address indexed _account, uint256 _value)
func (_NFTManager *NFTManagerFilterer) WatchWithdrawUToken(opts *bind.WatchOpts, sink chan<- *NFTManagerWithdrawUToken, who []common.Address, _tokenAddr []common.Address, _account []common.Address) (event.Subscription, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}
	var _tokenAddrRule []interface{}
	for _, _tokenAddrItem := range _tokenAddr {
		_tokenAddrRule = append(_tokenAddrRule, _tokenAddrItem)
	}
	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _NFTManager.contract.WatchLogs(opts, "WithdrawUToken", whoRule, _tokenAddrRule, _accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTManagerWithdrawUToken)
				if err := _NFTManager.contract.UnpackLog(event, "WithdrawUToken", log); err != nil {
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
// Solidity: event WithdrawUToken(address indexed who, address indexed _tokenAddr, address indexed _account, uint256 _value)
func (_NFTManager *NFTManagerFilterer) ParseWithdrawUToken(log types.Log) (*NFTManagerWithdrawUToken, error) {
	event := new(NFTManagerWithdrawUToken)
	if err := _NFTManager.contract.UnpackLog(event, "WithdrawUToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTManagerWthdrawIterator is returned from FilterWthdraw and is used to iterate over the raw logs and unpacked data for Wthdraw events raised by the NFTManager contract.
type NFTManagerWthdrawIterator struct {
	Event *NFTManagerWthdraw // Event containing the contract specifics and raw log

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
func (it *NFTManagerWthdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTManagerWthdraw)
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
		it.Event = new(NFTManagerWthdraw)
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
func (it *NFTManagerWthdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTManagerWthdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTManagerWthdraw represents a Wthdraw event raised by the NFTManager contract.
type NFTManagerWthdraw struct {
	Who    common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWthdraw is a free log retrieval operation binding the contract event 0x622ece8fe6c9adcdeb3aacfcf426facb5e94734cb9eb4cca7b66595853b87340.
//
// Solidity: event Wthdraw(address indexed who, uint256 _amount)
func (_NFTManager *NFTManagerFilterer) FilterWthdraw(opts *bind.FilterOpts, who []common.Address) (*NFTManagerWthdrawIterator, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _NFTManager.contract.FilterLogs(opts, "Wthdraw", whoRule)
	if err != nil {
		return nil, err
	}
	return &NFTManagerWthdrawIterator{contract: _NFTManager.contract, event: "Wthdraw", logs: logs, sub: sub}, nil
}

// WatchWthdraw is a free log subscription operation binding the contract event 0x622ece8fe6c9adcdeb3aacfcf426facb5e94734cb9eb4cca7b66595853b87340.
//
// Solidity: event Wthdraw(address indexed who, uint256 _amount)
func (_NFTManager *NFTManagerFilterer) WatchWthdraw(opts *bind.WatchOpts, sink chan<- *NFTManagerWthdraw, who []common.Address) (event.Subscription, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _NFTManager.contract.WatchLogs(opts, "Wthdraw", whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTManagerWthdraw)
				if err := _NFTManager.contract.UnpackLog(event, "Wthdraw", log); err != nil {
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

// ParseWthdraw is a log parse operation binding the contract event 0x622ece8fe6c9adcdeb3aacfcf426facb5e94734cb9eb4cca7b66595853b87340.
//
// Solidity: event Wthdraw(address indexed who, uint256 _amount)
func (_NFTManager *NFTManagerFilterer) ParseWthdraw(log types.Log) (*NFTManagerWthdraw, error) {
	event := new(NFTManagerWthdraw)
	if err := _NFTManager.contract.UnpackLog(event, "Wthdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
