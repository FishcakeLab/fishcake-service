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

// NftTokenManagerMetaData contains all meta data concerning the NftTokenManager contract.
var NftTokenManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_supplyLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MINTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowlistMint\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowlistMintActive\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchMintRemainingTokens\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnActive\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApproved\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasMinted\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isApprovedForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"merkleRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerOf\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setAllowlistMint\",\"inputs\":[{\"name\":\"status\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setApprovalForAll\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBurn\",\"inputs\":[{\"name\":\"_burnActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxSupply\",\"inputs\":[{\"name\":\"_supplyLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMerkleRoot\",\"inputs\":[{\"name\":\"_merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUriPrefix\",\"inputs\":[{\"name\":\"_uriPrefix\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supplyLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenURI\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalMinted\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"uriPrefix\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AllowlistMintSet\",\"inputs\":[{\"name\":\"changedBy\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"AllowlistMint\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApprovalForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BurnActiveSet\",\"inputs\":[{\"name\":\"changedBy\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"burnActive\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConsecutiveTransfer\",\"inputs\":[{\"name\":\"fromTokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"toTokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MerkleRootSet\",\"inputs\":[{\"name\":\"changedBy\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SupplyLimitSet\",\"inputs\":[{\"name\":\"changedBy\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"supplyLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UriPrefixSet\",\"inputs\":[{\"name\":\"changedBy\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"urlPrefix\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ApprovalQueryForNonexistentToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BalanceQueryForZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BurnNotActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MintERC2309QuantityExceedsLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MintToZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MintZeroQuantity\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnerQueryForNonexistentToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnershipNotInitializedForExtraData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferCallerNotOwnerNorApproved\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferFromIncorrectOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferToNonERC721ReceiverImplementer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferToZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"URIQueryForNonexistentToken\",\"inputs\":[]}]",
}

// NftTokenManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use NftTokenManagerMetaData.ABI instead.
var NftTokenManagerABI = NftTokenManagerMetaData.ABI

// NftTokenManager is an auto generated Go binding around an Ethereum contract.
type NftTokenManager struct {
	NftTokenManagerCaller     // Read-only binding to the contract
	NftTokenManagerTransactor // Write-only binding to the contract
	NftTokenManagerFilterer   // Log filterer for contract events
}

// NftTokenManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type NftTokenManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftTokenManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NftTokenManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftTokenManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NftTokenManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftTokenManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NftTokenManagerSession struct {
	Contract     *NftTokenManager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NftTokenManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NftTokenManagerCallerSession struct {
	Contract *NftTokenManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// NftTokenManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NftTokenManagerTransactorSession struct {
	Contract     *NftTokenManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// NftTokenManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type NftTokenManagerRaw struct {
	Contract *NftTokenManager // Generic contract binding to access the raw methods on
}

// NftTokenManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NftTokenManagerCallerRaw struct {
	Contract *NftTokenManagerCaller // Generic read-only contract binding to access the raw methods on
}

// NftTokenManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NftTokenManagerTransactorRaw struct {
	Contract *NftTokenManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNftTokenManager creates a new instance of NftTokenManager, bound to a specific deployed contract.
func NewNftTokenManager(address common.Address, backend bind.ContractBackend) (*NftTokenManager, error) {
	contract, err := bindNftTokenManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NftTokenManager{NftTokenManagerCaller: NftTokenManagerCaller{contract: contract}, NftTokenManagerTransactor: NftTokenManagerTransactor{contract: contract}, NftTokenManagerFilterer: NftTokenManagerFilterer{contract: contract}}, nil
}

// NewNftTokenManagerCaller creates a new read-only instance of NftTokenManager, bound to a specific deployed contract.
func NewNftTokenManagerCaller(address common.Address, caller bind.ContractCaller) (*NftTokenManagerCaller, error) {
	contract, err := bindNftTokenManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NftTokenManagerCaller{contract: contract}, nil
}

// NewNftTokenManagerTransactor creates a new write-only instance of NftTokenManager, bound to a specific deployed contract.
func NewNftTokenManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*NftTokenManagerTransactor, error) {
	contract, err := bindNftTokenManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NftTokenManagerTransactor{contract: contract}, nil
}

// NewNftTokenManagerFilterer creates a new log filterer instance of NftTokenManager, bound to a specific deployed contract.
func NewNftTokenManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*NftTokenManagerFilterer, error) {
	contract, err := bindNftTokenManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NftTokenManagerFilterer{contract: contract}, nil
}

// bindNftTokenManager binds a generic wrapper to an already deployed contract.
func bindNftTokenManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NftTokenManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NftTokenManager *NftTokenManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NftTokenManager.Contract.NftTokenManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NftTokenManager *NftTokenManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NftTokenManager.Contract.NftTokenManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NftTokenManager *NftTokenManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NftTokenManager.Contract.NftTokenManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NftTokenManager *NftTokenManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NftTokenManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NftTokenManager *NftTokenManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NftTokenManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NftTokenManager *NftTokenManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NftTokenManager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NftTokenManager *NftTokenManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NftTokenManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NftTokenManager *NftTokenManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _NftTokenManager.Contract.DEFAULTADMINROLE(&_NftTokenManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NftTokenManager *NftTokenManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _NftTokenManager.Contract.DEFAULTADMINROLE(&_NftTokenManager.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_NftTokenManager *NftTokenManagerCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NftTokenManager.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_NftTokenManager *NftTokenManagerSession) MINTERROLE() ([32]byte, error) {
	return _NftTokenManager.Contract.MINTERROLE(&_NftTokenManager.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_NftTokenManager *NftTokenManagerCallerSession) MINTERROLE() ([32]byte, error) {
	return _NftTokenManager.Contract.MINTERROLE(&_NftTokenManager.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_NftTokenManager *NftTokenManagerCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NftTokenManager.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_NftTokenManager *NftTokenManagerSession) PAUSERROLE() ([32]byte, error) {
	return _NftTokenManager.Contract.PAUSERROLE(&_NftTokenManager.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_NftTokenManager *NftTokenManagerCallerSession) PAUSERROLE() ([32]byte, error) {
	return _NftTokenManager.Contract.PAUSERROLE(&_NftTokenManager.CallOpts)
}

// AllowlistMintActive is a free data retrieval call binding the contract method 0x7417d6cc.
//
// Solidity: function allowlistMintActive() view returns(bool)
func (_NftTokenManager *NftTokenManagerCaller) AllowlistMintActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NftTokenManager.contract.Call(opts, &out, "allowlistMintActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowlistMintActive is a free data retrieval call binding the contract method 0x7417d6cc.
//
// Solidity: function allowlistMintActive() view returns(bool)
func (_NftTokenManager *NftTokenManagerSession) AllowlistMintActive() (bool, error) {
	return _NftTokenManager.Contract.AllowlistMintActive(&_NftTokenManager.CallOpts)
}

// AllowlistMintActive is a free data retrieval call binding the contract method 0x7417d6cc.
//
// Solidity: function allowlistMintActive() view returns(bool)
func (_NftTokenManager *NftTokenManagerCallerSession) AllowlistMintActive() (bool, error) {
	return _NftTokenManager.Contract.AllowlistMintActive(&_NftTokenManager.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NftTokenManager *NftTokenManagerCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NftTokenManager.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NftTokenManager *NftTokenManagerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _NftTokenManager.Contract.BalanceOf(&_NftTokenManager.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NftTokenManager *NftTokenManagerCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _NftTokenManager.Contract.BalanceOf(&_NftTokenManager.CallOpts, owner)
}

// BurnActive is a free data retrieval call binding the contract method 0x864ef3e5.
//
// Solidity: function burnActive() view returns(bool)
func (_NftTokenManager *NftTokenManagerCaller) BurnActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NftTokenManager.contract.Call(opts, &out, "burnActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BurnActive is a free data retrieval call binding the contract method 0x864ef3e5.
//
// Solidity: function burnActive() view returns(bool)
func (_NftTokenManager *NftTokenManagerSession) BurnActive() (bool, error) {
	return _NftTokenManager.Contract.BurnActive(&_NftTokenManager.CallOpts)
}

// BurnActive is a free data retrieval call binding the contract method 0x864ef3e5.
//
// Solidity: function burnActive() view returns(bool)
func (_NftTokenManager *NftTokenManagerCallerSession) BurnActive() (bool, error) {
	return _NftTokenManager.Contract.BurnActive(&_NftTokenManager.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NftTokenManager *NftTokenManagerCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NftTokenManager.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NftTokenManager *NftTokenManagerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _NftTokenManager.Contract.GetApproved(&_NftTokenManager.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NftTokenManager *NftTokenManagerCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _NftTokenManager.Contract.GetApproved(&_NftTokenManager.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NftTokenManager *NftTokenManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _NftTokenManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NftTokenManager *NftTokenManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _NftTokenManager.Contract.GetRoleAdmin(&_NftTokenManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NftTokenManager *NftTokenManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _NftTokenManager.Contract.GetRoleAdmin(&_NftTokenManager.CallOpts, role)
}

// HasMinted is a free data retrieval call binding the contract method 0x38e21cce.
//
// Solidity: function hasMinted(address owner) view returns(bool)
func (_NftTokenManager *NftTokenManagerCaller) HasMinted(opts *bind.CallOpts, owner common.Address) (bool, error) {
	var out []interface{}
	err := _NftTokenManager.contract.Call(opts, &out, "hasMinted", owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasMinted is a free data retrieval call binding the contract method 0x38e21cce.
//
// Solidity: function hasMinted(address owner) view returns(bool)
func (_NftTokenManager *NftTokenManagerSession) HasMinted(owner common.Address) (bool, error) {
	return _NftTokenManager.Contract.HasMinted(&_NftTokenManager.CallOpts, owner)
}

// HasMinted is a free data retrieval call binding the contract method 0x38e21cce.
//
// Solidity: function hasMinted(address owner) view returns(bool)
func (_NftTokenManager *NftTokenManagerCallerSession) HasMinted(owner common.Address) (bool, error) {
	return _NftTokenManager.Contract.HasMinted(&_NftTokenManager.CallOpts, owner)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NftTokenManager *NftTokenManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _NftTokenManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NftTokenManager *NftTokenManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _NftTokenManager.Contract.HasRole(&_NftTokenManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NftTokenManager *NftTokenManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _NftTokenManager.Contract.HasRole(&_NftTokenManager.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NftTokenManager *NftTokenManagerCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _NftTokenManager.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NftTokenManager *NftTokenManagerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _NftTokenManager.Contract.IsApprovedForAll(&_NftTokenManager.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NftTokenManager *NftTokenManagerCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _NftTokenManager.Contract.IsApprovedForAll(&_NftTokenManager.CallOpts, owner, operator)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_NftTokenManager *NftTokenManagerCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NftTokenManager.contract.Call(opts, &out, "merkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_NftTokenManager *NftTokenManagerSession) MerkleRoot() ([32]byte, error) {
	return _NftTokenManager.Contract.MerkleRoot(&_NftTokenManager.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_NftTokenManager *NftTokenManagerCallerSession) MerkleRoot() ([32]byte, error) {
	return _NftTokenManager.Contract.MerkleRoot(&_NftTokenManager.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NftTokenManager *NftTokenManagerCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NftTokenManager.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NftTokenManager *NftTokenManagerSession) Name() (string, error) {
	return _NftTokenManager.Contract.Name(&_NftTokenManager.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NftTokenManager *NftTokenManagerCallerSession) Name() (string, error) {
	return _NftTokenManager.Contract.Name(&_NftTokenManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NftTokenManager *NftTokenManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NftTokenManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NftTokenManager *NftTokenManagerSession) Owner() (common.Address, error) {
	return _NftTokenManager.Contract.Owner(&_NftTokenManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NftTokenManager *NftTokenManagerCallerSession) Owner() (common.Address, error) {
	return _NftTokenManager.Contract.Owner(&_NftTokenManager.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NftTokenManager *NftTokenManagerCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NftTokenManager.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NftTokenManager *NftTokenManagerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _NftTokenManager.Contract.OwnerOf(&_NftTokenManager.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NftTokenManager *NftTokenManagerCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _NftTokenManager.Contract.OwnerOf(&_NftTokenManager.CallOpts, tokenId)
}

// SupplyLimit is a free data retrieval call binding the contract method 0x19d1997a.
//
// Solidity: function supplyLimit() view returns(uint256)
func (_NftTokenManager *NftTokenManagerCaller) SupplyLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NftTokenManager.contract.Call(opts, &out, "supplyLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyLimit is a free data retrieval call binding the contract method 0x19d1997a.
//
// Solidity: function supplyLimit() view returns(uint256)
func (_NftTokenManager *NftTokenManagerSession) SupplyLimit() (*big.Int, error) {
	return _NftTokenManager.Contract.SupplyLimit(&_NftTokenManager.CallOpts)
}

// SupplyLimit is a free data retrieval call binding the contract method 0x19d1997a.
//
// Solidity: function supplyLimit() view returns(uint256)
func (_NftTokenManager *NftTokenManagerCallerSession) SupplyLimit() (*big.Int, error) {
	return _NftTokenManager.Contract.SupplyLimit(&_NftTokenManager.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NftTokenManager *NftTokenManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _NftTokenManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NftTokenManager *NftTokenManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NftTokenManager.Contract.SupportsInterface(&_NftTokenManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NftTokenManager *NftTokenManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NftTokenManager.Contract.SupportsInterface(&_NftTokenManager.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NftTokenManager *NftTokenManagerCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NftTokenManager.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NftTokenManager *NftTokenManagerSession) Symbol() (string, error) {
	return _NftTokenManager.Contract.Symbol(&_NftTokenManager.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NftTokenManager *NftTokenManagerCallerSession) Symbol() (string, error) {
	return _NftTokenManager.Contract.Symbol(&_NftTokenManager.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_NftTokenManager *NftTokenManagerCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _NftTokenManager.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_NftTokenManager *NftTokenManagerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _NftTokenManager.Contract.TokenURI(&_NftTokenManager.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_NftTokenManager *NftTokenManagerCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _NftTokenManager.Contract.TokenURI(&_NftTokenManager.CallOpts, tokenId)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_NftTokenManager *NftTokenManagerCaller) TotalMinted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NftTokenManager.contract.Call(opts, &out, "totalMinted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_NftTokenManager *NftTokenManagerSession) TotalMinted() (*big.Int, error) {
	return _NftTokenManager.Contract.TotalMinted(&_NftTokenManager.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_NftTokenManager *NftTokenManagerCallerSession) TotalMinted() (*big.Int, error) {
	return _NftTokenManager.Contract.TotalMinted(&_NftTokenManager.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NftTokenManager *NftTokenManagerCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NftTokenManager.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NftTokenManager *NftTokenManagerSession) TotalSupply() (*big.Int, error) {
	return _NftTokenManager.Contract.TotalSupply(&_NftTokenManager.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NftTokenManager *NftTokenManagerCallerSession) TotalSupply() (*big.Int, error) {
	return _NftTokenManager.Contract.TotalSupply(&_NftTokenManager.CallOpts)
}

// UriPrefix is a free data retrieval call binding the contract method 0x62b99ad4.
//
// Solidity: function uriPrefix() view returns(string)
func (_NftTokenManager *NftTokenManagerCaller) UriPrefix(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NftTokenManager.contract.Call(opts, &out, "uriPrefix")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UriPrefix is a free data retrieval call binding the contract method 0x62b99ad4.
//
// Solidity: function uriPrefix() view returns(string)
func (_NftTokenManager *NftTokenManagerSession) UriPrefix() (string, error) {
	return _NftTokenManager.Contract.UriPrefix(&_NftTokenManager.CallOpts)
}

// UriPrefix is a free data retrieval call binding the contract method 0x62b99ad4.
//
// Solidity: function uriPrefix() view returns(string)
func (_NftTokenManager *NftTokenManagerCallerSession) UriPrefix() (string, error) {
	return _NftTokenManager.Contract.UriPrefix(&_NftTokenManager.CallOpts)
}

// AllowlistMint is a paid mutator transaction binding the contract method 0x537924ef.
//
// Solidity: function allowlistMint(bytes32[] proof) returns()
func (_NftTokenManager *NftTokenManagerTransactor) AllowlistMint(opts *bind.TransactOpts, proof [][32]byte) (*types.Transaction, error) {
	return _NftTokenManager.contract.Transact(opts, "allowlistMint", proof)
}

// AllowlistMint is a paid mutator transaction binding the contract method 0x537924ef.
//
// Solidity: function allowlistMint(bytes32[] proof) returns()
func (_NftTokenManager *NftTokenManagerSession) AllowlistMint(proof [][32]byte) (*types.Transaction, error) {
	return _NftTokenManager.Contract.AllowlistMint(&_NftTokenManager.TransactOpts, proof)
}

// AllowlistMint is a paid mutator transaction binding the contract method 0x537924ef.
//
// Solidity: function allowlistMint(bytes32[] proof) returns()
func (_NftTokenManager *NftTokenManagerTransactorSession) AllowlistMint(proof [][32]byte) (*types.Transaction, error) {
	return _NftTokenManager.Contract.AllowlistMint(&_NftTokenManager.TransactOpts, proof)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_NftTokenManager *NftTokenManagerTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftTokenManager.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_NftTokenManager *NftTokenManagerSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftTokenManager.Contract.Approve(&_NftTokenManager.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_NftTokenManager *NftTokenManagerTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftTokenManager.Contract.Approve(&_NftTokenManager.TransactOpts, to, tokenId)
}

// BatchMintRemainingTokens is a paid mutator transaction binding the contract method 0x4e7aa674.
//
// Solidity: function batchMintRemainingTokens(address receiver, uint256 amount) returns()
func (_NftTokenManager *NftTokenManagerTransactor) BatchMintRemainingTokens(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NftTokenManager.contract.Transact(opts, "batchMintRemainingTokens", receiver, amount)
}

// BatchMintRemainingTokens is a paid mutator transaction binding the contract method 0x4e7aa674.
//
// Solidity: function batchMintRemainingTokens(address receiver, uint256 amount) returns()
func (_NftTokenManager *NftTokenManagerSession) BatchMintRemainingTokens(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NftTokenManager.Contract.BatchMintRemainingTokens(&_NftTokenManager.TransactOpts, receiver, amount)
}

// BatchMintRemainingTokens is a paid mutator transaction binding the contract method 0x4e7aa674.
//
// Solidity: function batchMintRemainingTokens(address receiver, uint256 amount) returns()
func (_NftTokenManager *NftTokenManagerTransactorSession) BatchMintRemainingTokens(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NftTokenManager.Contract.BatchMintRemainingTokens(&_NftTokenManager.TransactOpts, receiver, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_NftTokenManager *NftTokenManagerTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _NftTokenManager.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_NftTokenManager *NftTokenManagerSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _NftTokenManager.Contract.Burn(&_NftTokenManager.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_NftTokenManager *NftTokenManagerTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _NftTokenManager.Contract.Burn(&_NftTokenManager.TransactOpts, tokenId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NftTokenManager *NftTokenManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NftTokenManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NftTokenManager *NftTokenManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NftTokenManager.Contract.GrantRole(&_NftTokenManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NftTokenManager *NftTokenManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NftTokenManager.Contract.GrantRole(&_NftTokenManager.TransactOpts, role, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NftTokenManager *NftTokenManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NftTokenManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NftTokenManager *NftTokenManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _NftTokenManager.Contract.RenounceOwnership(&_NftTokenManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NftTokenManager *NftTokenManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NftTokenManager.Contract.RenounceOwnership(&_NftTokenManager.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_NftTokenManager *NftTokenManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _NftTokenManager.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_NftTokenManager *NftTokenManagerSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _NftTokenManager.Contract.RenounceRole(&_NftTokenManager.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_NftTokenManager *NftTokenManagerTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _NftTokenManager.Contract.RenounceRole(&_NftTokenManager.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NftTokenManager *NftTokenManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NftTokenManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NftTokenManager *NftTokenManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NftTokenManager.Contract.RevokeRole(&_NftTokenManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NftTokenManager *NftTokenManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NftTokenManager.Contract.RevokeRole(&_NftTokenManager.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_NftTokenManager *NftTokenManagerTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftTokenManager.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_NftTokenManager *NftTokenManagerSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftTokenManager.Contract.SafeTransferFrom(&_NftTokenManager.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_NftTokenManager *NftTokenManagerTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftTokenManager.Contract.SafeTransferFrom(&_NftTokenManager.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_NftTokenManager *NftTokenManagerTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _NftTokenManager.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_NftTokenManager *NftTokenManagerSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _NftTokenManager.Contract.SafeTransferFrom0(&_NftTokenManager.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_NftTokenManager *NftTokenManagerTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _NftTokenManager.Contract.SafeTransferFrom0(&_NftTokenManager.TransactOpts, from, to, tokenId, _data)
}

// SetAllowlistMint is a paid mutator transaction binding the contract method 0x19f09583.
//
// Solidity: function setAllowlistMint(bool status) returns()
func (_NftTokenManager *NftTokenManagerTransactor) SetAllowlistMint(opts *bind.TransactOpts, status bool) (*types.Transaction, error) {
	return _NftTokenManager.contract.Transact(opts, "setAllowlistMint", status)
}

// SetAllowlistMint is a paid mutator transaction binding the contract method 0x19f09583.
//
// Solidity: function setAllowlistMint(bool status) returns()
func (_NftTokenManager *NftTokenManagerSession) SetAllowlistMint(status bool) (*types.Transaction, error) {
	return _NftTokenManager.Contract.SetAllowlistMint(&_NftTokenManager.TransactOpts, status)
}

// SetAllowlistMint is a paid mutator transaction binding the contract method 0x19f09583.
//
// Solidity: function setAllowlistMint(bool status) returns()
func (_NftTokenManager *NftTokenManagerTransactorSession) SetAllowlistMint(status bool) (*types.Transaction, error) {
	return _NftTokenManager.Contract.SetAllowlistMint(&_NftTokenManager.TransactOpts, status)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NftTokenManager *NftTokenManagerTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _NftTokenManager.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NftTokenManager *NftTokenManagerSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _NftTokenManager.Contract.SetApprovalForAll(&_NftTokenManager.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NftTokenManager *NftTokenManagerTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _NftTokenManager.Contract.SetApprovalForAll(&_NftTokenManager.TransactOpts, operator, approved)
}

// SetBurn is a paid mutator transaction binding the contract method 0xbd55cf0d.
//
// Solidity: function setBurn(bool _burnActive) returns()
func (_NftTokenManager *NftTokenManagerTransactor) SetBurn(opts *bind.TransactOpts, _burnActive bool) (*types.Transaction, error) {
	return _NftTokenManager.contract.Transact(opts, "setBurn", _burnActive)
}

// SetBurn is a paid mutator transaction binding the contract method 0xbd55cf0d.
//
// Solidity: function setBurn(bool _burnActive) returns()
func (_NftTokenManager *NftTokenManagerSession) SetBurn(_burnActive bool) (*types.Transaction, error) {
	return _NftTokenManager.Contract.SetBurn(&_NftTokenManager.TransactOpts, _burnActive)
}

// SetBurn is a paid mutator transaction binding the contract method 0xbd55cf0d.
//
// Solidity: function setBurn(bool _burnActive) returns()
func (_NftTokenManager *NftTokenManagerTransactorSession) SetBurn(_burnActive bool) (*types.Transaction, error) {
	return _NftTokenManager.Contract.SetBurn(&_NftTokenManager.TransactOpts, _burnActive)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 _supplyLimit) returns()
func (_NftTokenManager *NftTokenManagerTransactor) SetMaxSupply(opts *bind.TransactOpts, _supplyLimit *big.Int) (*types.Transaction, error) {
	return _NftTokenManager.contract.Transact(opts, "setMaxSupply", _supplyLimit)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 _supplyLimit) returns()
func (_NftTokenManager *NftTokenManagerSession) SetMaxSupply(_supplyLimit *big.Int) (*types.Transaction, error) {
	return _NftTokenManager.Contract.SetMaxSupply(&_NftTokenManager.TransactOpts, _supplyLimit)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 _supplyLimit) returns()
func (_NftTokenManager *NftTokenManagerTransactorSession) SetMaxSupply(_supplyLimit *big.Int) (*types.Transaction, error) {
	return _NftTokenManager.Contract.SetMaxSupply(&_NftTokenManager.TransactOpts, _supplyLimit)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x7cb64759.
//
// Solidity: function setMerkleRoot(bytes32 _merkleRoot) returns()
func (_NftTokenManager *NftTokenManagerTransactor) SetMerkleRoot(opts *bind.TransactOpts, _merkleRoot [32]byte) (*types.Transaction, error) {
	return _NftTokenManager.contract.Transact(opts, "setMerkleRoot", _merkleRoot)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x7cb64759.
//
// Solidity: function setMerkleRoot(bytes32 _merkleRoot) returns()
func (_NftTokenManager *NftTokenManagerSession) SetMerkleRoot(_merkleRoot [32]byte) (*types.Transaction, error) {
	return _NftTokenManager.Contract.SetMerkleRoot(&_NftTokenManager.TransactOpts, _merkleRoot)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x7cb64759.
//
// Solidity: function setMerkleRoot(bytes32 _merkleRoot) returns()
func (_NftTokenManager *NftTokenManagerTransactorSession) SetMerkleRoot(_merkleRoot [32]byte) (*types.Transaction, error) {
	return _NftTokenManager.Contract.SetMerkleRoot(&_NftTokenManager.TransactOpts, _merkleRoot)
}

// SetUriPrefix is a paid mutator transaction binding the contract method 0x7ec4a659.
//
// Solidity: function setUriPrefix(string _uriPrefix) returns()
func (_NftTokenManager *NftTokenManagerTransactor) SetUriPrefix(opts *bind.TransactOpts, _uriPrefix string) (*types.Transaction, error) {
	return _NftTokenManager.contract.Transact(opts, "setUriPrefix", _uriPrefix)
}

// SetUriPrefix is a paid mutator transaction binding the contract method 0x7ec4a659.
//
// Solidity: function setUriPrefix(string _uriPrefix) returns()
func (_NftTokenManager *NftTokenManagerSession) SetUriPrefix(_uriPrefix string) (*types.Transaction, error) {
	return _NftTokenManager.Contract.SetUriPrefix(&_NftTokenManager.TransactOpts, _uriPrefix)
}

// SetUriPrefix is a paid mutator transaction binding the contract method 0x7ec4a659.
//
// Solidity: function setUriPrefix(string _uriPrefix) returns()
func (_NftTokenManager *NftTokenManagerTransactorSession) SetUriPrefix(_uriPrefix string) (*types.Transaction, error) {
	return _NftTokenManager.Contract.SetUriPrefix(&_NftTokenManager.TransactOpts, _uriPrefix)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_NftTokenManager *NftTokenManagerTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftTokenManager.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_NftTokenManager *NftTokenManagerSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftTokenManager.Contract.TransferFrom(&_NftTokenManager.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_NftTokenManager *NftTokenManagerTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftTokenManager.Contract.TransferFrom(&_NftTokenManager.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NftTokenManager *NftTokenManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NftTokenManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NftTokenManager *NftTokenManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NftTokenManager.Contract.TransferOwnership(&_NftTokenManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NftTokenManager *NftTokenManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NftTokenManager.Contract.TransferOwnership(&_NftTokenManager.TransactOpts, newOwner)
}

// NftTokenManagerAllowlistMintSetIterator is returned from FilterAllowlistMintSet and is used to iterate over the raw logs and unpacked data for AllowlistMintSet events raised by the NftTokenManager contract.
type NftTokenManagerAllowlistMintSetIterator struct {
	Event *NftTokenManagerAllowlistMintSet // Event containing the contract specifics and raw log

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
func (it *NftTokenManagerAllowlistMintSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftTokenManagerAllowlistMintSet)
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
		it.Event = new(NftTokenManagerAllowlistMintSet)
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
func (it *NftTokenManagerAllowlistMintSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftTokenManagerAllowlistMintSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftTokenManagerAllowlistMintSet represents a AllowlistMintSet event raised by the NftTokenManager contract.
type NftTokenManagerAllowlistMintSet struct {
	ChangedBy     common.Address
	AllowlistMint bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAllowlistMintSet is a free log retrieval operation binding the contract event 0x0b4b103e6f809edb3e5b2edd0998db0367384fa4113392fe42d4ce76fc7d2678.
//
// Solidity: event AllowlistMintSet(address indexed changedBy, bool AllowlistMint)
func (_NftTokenManager *NftTokenManagerFilterer) FilterAllowlistMintSet(opts *bind.FilterOpts, changedBy []common.Address) (*NftTokenManagerAllowlistMintSetIterator, error) {

	var changedByRule []interface{}
	for _, changedByItem := range changedBy {
		changedByRule = append(changedByRule, changedByItem)
	}

	logs, sub, err := _NftTokenManager.contract.FilterLogs(opts, "AllowlistMintSet", changedByRule)
	if err != nil {
		return nil, err
	}
	return &NftTokenManagerAllowlistMintSetIterator{contract: _NftTokenManager.contract, event: "AllowlistMintSet", logs: logs, sub: sub}, nil
}

// WatchAllowlistMintSet is a free log subscription operation binding the contract event 0x0b4b103e6f809edb3e5b2edd0998db0367384fa4113392fe42d4ce76fc7d2678.
//
// Solidity: event AllowlistMintSet(address indexed changedBy, bool AllowlistMint)
func (_NftTokenManager *NftTokenManagerFilterer) WatchAllowlistMintSet(opts *bind.WatchOpts, sink chan<- *NftTokenManagerAllowlistMintSet, changedBy []common.Address) (event.Subscription, error) {

	var changedByRule []interface{}
	for _, changedByItem := range changedBy {
		changedByRule = append(changedByRule, changedByItem)
	}

	logs, sub, err := _NftTokenManager.contract.WatchLogs(opts, "AllowlistMintSet", changedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftTokenManagerAllowlistMintSet)
				if err := _NftTokenManager.contract.UnpackLog(event, "AllowlistMintSet", log); err != nil {
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

// ParseAllowlistMintSet is a log parse operation binding the contract event 0x0b4b103e6f809edb3e5b2edd0998db0367384fa4113392fe42d4ce76fc7d2678.
//
// Solidity: event AllowlistMintSet(address indexed changedBy, bool AllowlistMint)
func (_NftTokenManager *NftTokenManagerFilterer) ParseAllowlistMintSet(log types.Log) (*NftTokenManagerAllowlistMintSet, error) {
	event := new(NftTokenManagerAllowlistMintSet)
	if err := _NftTokenManager.contract.UnpackLog(event, "AllowlistMintSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftTokenManagerApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the NftTokenManager contract.
type NftTokenManagerApprovalIterator struct {
	Event *NftTokenManagerApproval // Event containing the contract specifics and raw log

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
func (it *NftTokenManagerApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftTokenManagerApproval)
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
		it.Event = new(NftTokenManagerApproval)
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
func (it *NftTokenManagerApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftTokenManagerApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftTokenManagerApproval represents a Approval event raised by the NftTokenManager contract.
type NftTokenManagerApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_NftTokenManager *NftTokenManagerFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*NftTokenManagerApprovalIterator, error) {

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

	logs, sub, err := _NftTokenManager.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftTokenManagerApprovalIterator{contract: _NftTokenManager.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_NftTokenManager *NftTokenManagerFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NftTokenManagerApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _NftTokenManager.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftTokenManagerApproval)
				if err := _NftTokenManager.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_NftTokenManager *NftTokenManagerFilterer) ParseApproval(log types.Log) (*NftTokenManagerApproval, error) {
	event := new(NftTokenManagerApproval)
	if err := _NftTokenManager.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftTokenManagerApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the NftTokenManager contract.
type NftTokenManagerApprovalForAllIterator struct {
	Event *NftTokenManagerApprovalForAll // Event containing the contract specifics and raw log

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
func (it *NftTokenManagerApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftTokenManagerApprovalForAll)
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
		it.Event = new(NftTokenManagerApprovalForAll)
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
func (it *NftTokenManagerApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftTokenManagerApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftTokenManagerApprovalForAll represents a ApprovalForAll event raised by the NftTokenManager contract.
type NftTokenManagerApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_NftTokenManager *NftTokenManagerFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*NftTokenManagerApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _NftTokenManager.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &NftTokenManagerApprovalForAllIterator{contract: _NftTokenManager.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_NftTokenManager *NftTokenManagerFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *NftTokenManagerApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _NftTokenManager.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftTokenManagerApprovalForAll)
				if err := _NftTokenManager.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_NftTokenManager *NftTokenManagerFilterer) ParseApprovalForAll(log types.Log) (*NftTokenManagerApprovalForAll, error) {
	event := new(NftTokenManagerApprovalForAll)
	if err := _NftTokenManager.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftTokenManagerBurnActiveSetIterator is returned from FilterBurnActiveSet and is used to iterate over the raw logs and unpacked data for BurnActiveSet events raised by the NftTokenManager contract.
type NftTokenManagerBurnActiveSetIterator struct {
	Event *NftTokenManagerBurnActiveSet // Event containing the contract specifics and raw log

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
func (it *NftTokenManagerBurnActiveSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftTokenManagerBurnActiveSet)
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
		it.Event = new(NftTokenManagerBurnActiveSet)
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
func (it *NftTokenManagerBurnActiveSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftTokenManagerBurnActiveSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftTokenManagerBurnActiveSet represents a BurnActiveSet event raised by the NftTokenManager contract.
type NftTokenManagerBurnActiveSet struct {
	ChangedBy  common.Address
	BurnActive bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBurnActiveSet is a free log retrieval operation binding the contract event 0xda8236fd9b752f65ed87e7f6aef7d936153f70f09ed9a40cbb68630faff48322.
//
// Solidity: event BurnActiveSet(address indexed changedBy, bool burnActive)
func (_NftTokenManager *NftTokenManagerFilterer) FilterBurnActiveSet(opts *bind.FilterOpts, changedBy []common.Address) (*NftTokenManagerBurnActiveSetIterator, error) {

	var changedByRule []interface{}
	for _, changedByItem := range changedBy {
		changedByRule = append(changedByRule, changedByItem)
	}

	logs, sub, err := _NftTokenManager.contract.FilterLogs(opts, "BurnActiveSet", changedByRule)
	if err != nil {
		return nil, err
	}
	return &NftTokenManagerBurnActiveSetIterator{contract: _NftTokenManager.contract, event: "BurnActiveSet", logs: logs, sub: sub}, nil
}

// WatchBurnActiveSet is a free log subscription operation binding the contract event 0xda8236fd9b752f65ed87e7f6aef7d936153f70f09ed9a40cbb68630faff48322.
//
// Solidity: event BurnActiveSet(address indexed changedBy, bool burnActive)
func (_NftTokenManager *NftTokenManagerFilterer) WatchBurnActiveSet(opts *bind.WatchOpts, sink chan<- *NftTokenManagerBurnActiveSet, changedBy []common.Address) (event.Subscription, error) {

	var changedByRule []interface{}
	for _, changedByItem := range changedBy {
		changedByRule = append(changedByRule, changedByItem)
	}

	logs, sub, err := _NftTokenManager.contract.WatchLogs(opts, "BurnActiveSet", changedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftTokenManagerBurnActiveSet)
				if err := _NftTokenManager.contract.UnpackLog(event, "BurnActiveSet", log); err != nil {
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

// ParseBurnActiveSet is a log parse operation binding the contract event 0xda8236fd9b752f65ed87e7f6aef7d936153f70f09ed9a40cbb68630faff48322.
//
// Solidity: event BurnActiveSet(address indexed changedBy, bool burnActive)
func (_NftTokenManager *NftTokenManagerFilterer) ParseBurnActiveSet(log types.Log) (*NftTokenManagerBurnActiveSet, error) {
	event := new(NftTokenManagerBurnActiveSet)
	if err := _NftTokenManager.contract.UnpackLog(event, "BurnActiveSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftTokenManagerConsecutiveTransferIterator is returned from FilterConsecutiveTransfer and is used to iterate over the raw logs and unpacked data for ConsecutiveTransfer events raised by the NftTokenManager contract.
type NftTokenManagerConsecutiveTransferIterator struct {
	Event *NftTokenManagerConsecutiveTransfer // Event containing the contract specifics and raw log

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
func (it *NftTokenManagerConsecutiveTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftTokenManagerConsecutiveTransfer)
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
		it.Event = new(NftTokenManagerConsecutiveTransfer)
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
func (it *NftTokenManagerConsecutiveTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftTokenManagerConsecutiveTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftTokenManagerConsecutiveTransfer represents a ConsecutiveTransfer event raised by the NftTokenManager contract.
type NftTokenManagerConsecutiveTransfer struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	From        common.Address
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsecutiveTransfer is a free log retrieval operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_NftTokenManager *NftTokenManagerFilterer) FilterConsecutiveTransfer(opts *bind.FilterOpts, fromTokenId []*big.Int, from []common.Address, to []common.Address) (*NftTokenManagerConsecutiveTransferIterator, error) {

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

	logs, sub, err := _NftTokenManager.contract.FilterLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NftTokenManagerConsecutiveTransferIterator{contract: _NftTokenManager.contract, event: "ConsecutiveTransfer", logs: logs, sub: sub}, nil
}

// WatchConsecutiveTransfer is a free log subscription operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_NftTokenManager *NftTokenManagerFilterer) WatchConsecutiveTransfer(opts *bind.WatchOpts, sink chan<- *NftTokenManagerConsecutiveTransfer, fromTokenId []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _NftTokenManager.contract.WatchLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftTokenManagerConsecutiveTransfer)
				if err := _NftTokenManager.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
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
func (_NftTokenManager *NftTokenManagerFilterer) ParseConsecutiveTransfer(log types.Log) (*NftTokenManagerConsecutiveTransfer, error) {
	event := new(NftTokenManagerConsecutiveTransfer)
	if err := _NftTokenManager.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftTokenManagerMerkleRootSetIterator is returned from FilterMerkleRootSet and is used to iterate over the raw logs and unpacked data for MerkleRootSet events raised by the NftTokenManager contract.
type NftTokenManagerMerkleRootSetIterator struct {
	Event *NftTokenManagerMerkleRootSet // Event containing the contract specifics and raw log

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
func (it *NftTokenManagerMerkleRootSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftTokenManagerMerkleRootSet)
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
		it.Event = new(NftTokenManagerMerkleRootSet)
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
func (it *NftTokenManagerMerkleRootSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftTokenManagerMerkleRootSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftTokenManagerMerkleRootSet represents a MerkleRootSet event raised by the NftTokenManager contract.
type NftTokenManagerMerkleRootSet struct {
	ChangedBy common.Address
	Root      [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMerkleRootSet is a free log retrieval operation binding the contract event 0x4ef92811960c6617dd84b673329cb6c4d277e6356640f9025b52373798fd62a4.
//
// Solidity: event MerkleRootSet(address indexed changedBy, bytes32 root)
func (_NftTokenManager *NftTokenManagerFilterer) FilterMerkleRootSet(opts *bind.FilterOpts, changedBy []common.Address) (*NftTokenManagerMerkleRootSetIterator, error) {

	var changedByRule []interface{}
	for _, changedByItem := range changedBy {
		changedByRule = append(changedByRule, changedByItem)
	}

	logs, sub, err := _NftTokenManager.contract.FilterLogs(opts, "MerkleRootSet", changedByRule)
	if err != nil {
		return nil, err
	}
	return &NftTokenManagerMerkleRootSetIterator{contract: _NftTokenManager.contract, event: "MerkleRootSet", logs: logs, sub: sub}, nil
}

// WatchMerkleRootSet is a free log subscription operation binding the contract event 0x4ef92811960c6617dd84b673329cb6c4d277e6356640f9025b52373798fd62a4.
//
// Solidity: event MerkleRootSet(address indexed changedBy, bytes32 root)
func (_NftTokenManager *NftTokenManagerFilterer) WatchMerkleRootSet(opts *bind.WatchOpts, sink chan<- *NftTokenManagerMerkleRootSet, changedBy []common.Address) (event.Subscription, error) {

	var changedByRule []interface{}
	for _, changedByItem := range changedBy {
		changedByRule = append(changedByRule, changedByItem)
	}

	logs, sub, err := _NftTokenManager.contract.WatchLogs(opts, "MerkleRootSet", changedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftTokenManagerMerkleRootSet)
				if err := _NftTokenManager.contract.UnpackLog(event, "MerkleRootSet", log); err != nil {
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

// ParseMerkleRootSet is a log parse operation binding the contract event 0x4ef92811960c6617dd84b673329cb6c4d277e6356640f9025b52373798fd62a4.
//
// Solidity: event MerkleRootSet(address indexed changedBy, bytes32 root)
func (_NftTokenManager *NftTokenManagerFilterer) ParseMerkleRootSet(log types.Log) (*NftTokenManagerMerkleRootSet, error) {
	event := new(NftTokenManagerMerkleRootSet)
	if err := _NftTokenManager.contract.UnpackLog(event, "MerkleRootSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftTokenManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NftTokenManager contract.
type NftTokenManagerOwnershipTransferredIterator struct {
	Event *NftTokenManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NftTokenManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftTokenManagerOwnershipTransferred)
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
		it.Event = new(NftTokenManagerOwnershipTransferred)
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
func (it *NftTokenManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftTokenManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftTokenManagerOwnershipTransferred represents a OwnershipTransferred event raised by the NftTokenManager contract.
type NftTokenManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NftTokenManager *NftTokenManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NftTokenManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NftTokenManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NftTokenManagerOwnershipTransferredIterator{contract: _NftTokenManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NftTokenManager *NftTokenManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NftTokenManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NftTokenManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftTokenManagerOwnershipTransferred)
				if err := _NftTokenManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NftTokenManager *NftTokenManagerFilterer) ParseOwnershipTransferred(log types.Log) (*NftTokenManagerOwnershipTransferred, error) {
	event := new(NftTokenManagerOwnershipTransferred)
	if err := _NftTokenManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftTokenManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the NftTokenManager contract.
type NftTokenManagerRoleAdminChangedIterator struct {
	Event *NftTokenManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *NftTokenManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftTokenManagerRoleAdminChanged)
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
		it.Event = new(NftTokenManagerRoleAdminChanged)
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
func (it *NftTokenManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftTokenManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftTokenManagerRoleAdminChanged represents a RoleAdminChanged event raised by the NftTokenManager contract.
type NftTokenManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NftTokenManager *NftTokenManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*NftTokenManagerRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _NftTokenManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &NftTokenManagerRoleAdminChangedIterator{contract: _NftTokenManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NftTokenManager *NftTokenManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *NftTokenManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _NftTokenManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftTokenManagerRoleAdminChanged)
				if err := _NftTokenManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NftTokenManager *NftTokenManagerFilterer) ParseRoleAdminChanged(log types.Log) (*NftTokenManagerRoleAdminChanged, error) {
	event := new(NftTokenManagerRoleAdminChanged)
	if err := _NftTokenManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftTokenManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the NftTokenManager contract.
type NftTokenManagerRoleGrantedIterator struct {
	Event *NftTokenManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *NftTokenManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftTokenManagerRoleGranted)
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
		it.Event = new(NftTokenManagerRoleGranted)
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
func (it *NftTokenManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftTokenManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftTokenManagerRoleGranted represents a RoleGranted event raised by the NftTokenManager contract.
type NftTokenManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NftTokenManager *NftTokenManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NftTokenManagerRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NftTokenManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NftTokenManagerRoleGrantedIterator{contract: _NftTokenManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NftTokenManager *NftTokenManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *NftTokenManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NftTokenManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftTokenManagerRoleGranted)
				if err := _NftTokenManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NftTokenManager *NftTokenManagerFilterer) ParseRoleGranted(log types.Log) (*NftTokenManagerRoleGranted, error) {
	event := new(NftTokenManagerRoleGranted)
	if err := _NftTokenManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftTokenManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the NftTokenManager contract.
type NftTokenManagerRoleRevokedIterator struct {
	Event *NftTokenManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *NftTokenManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftTokenManagerRoleRevoked)
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
		it.Event = new(NftTokenManagerRoleRevoked)
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
func (it *NftTokenManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftTokenManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftTokenManagerRoleRevoked represents a RoleRevoked event raised by the NftTokenManager contract.
type NftTokenManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NftTokenManager *NftTokenManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NftTokenManagerRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NftTokenManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NftTokenManagerRoleRevokedIterator{contract: _NftTokenManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NftTokenManager *NftTokenManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *NftTokenManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NftTokenManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftTokenManagerRoleRevoked)
				if err := _NftTokenManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NftTokenManager *NftTokenManagerFilterer) ParseRoleRevoked(log types.Log) (*NftTokenManagerRoleRevoked, error) {
	event := new(NftTokenManagerRoleRevoked)
	if err := _NftTokenManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftTokenManagerSupplyLimitSetIterator is returned from FilterSupplyLimitSet and is used to iterate over the raw logs and unpacked data for SupplyLimitSet events raised by the NftTokenManager contract.
type NftTokenManagerSupplyLimitSetIterator struct {
	Event *NftTokenManagerSupplyLimitSet // Event containing the contract specifics and raw log

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
func (it *NftTokenManagerSupplyLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftTokenManagerSupplyLimitSet)
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
		it.Event = new(NftTokenManagerSupplyLimitSet)
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
func (it *NftTokenManagerSupplyLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftTokenManagerSupplyLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftTokenManagerSupplyLimitSet represents a SupplyLimitSet event raised by the NftTokenManager contract.
type NftTokenManagerSupplyLimitSet struct {
	ChangedBy   common.Address
	SupplyLimit *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSupplyLimitSet is a free log retrieval operation binding the contract event 0x2f1ac7412e0c36ff7b73f7c4f474022dfa5d3f1c09e3262f63c9eb4627a9f505.
//
// Solidity: event SupplyLimitSet(address indexed changedBy, uint256 supplyLimit)
func (_NftTokenManager *NftTokenManagerFilterer) FilterSupplyLimitSet(opts *bind.FilterOpts, changedBy []common.Address) (*NftTokenManagerSupplyLimitSetIterator, error) {

	var changedByRule []interface{}
	for _, changedByItem := range changedBy {
		changedByRule = append(changedByRule, changedByItem)
	}

	logs, sub, err := _NftTokenManager.contract.FilterLogs(opts, "SupplyLimitSet", changedByRule)
	if err != nil {
		return nil, err
	}
	return &NftTokenManagerSupplyLimitSetIterator{contract: _NftTokenManager.contract, event: "SupplyLimitSet", logs: logs, sub: sub}, nil
}

// WatchSupplyLimitSet is a free log subscription operation binding the contract event 0x2f1ac7412e0c36ff7b73f7c4f474022dfa5d3f1c09e3262f63c9eb4627a9f505.
//
// Solidity: event SupplyLimitSet(address indexed changedBy, uint256 supplyLimit)
func (_NftTokenManager *NftTokenManagerFilterer) WatchSupplyLimitSet(opts *bind.WatchOpts, sink chan<- *NftTokenManagerSupplyLimitSet, changedBy []common.Address) (event.Subscription, error) {

	var changedByRule []interface{}
	for _, changedByItem := range changedBy {
		changedByRule = append(changedByRule, changedByItem)
	}

	logs, sub, err := _NftTokenManager.contract.WatchLogs(opts, "SupplyLimitSet", changedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftTokenManagerSupplyLimitSet)
				if err := _NftTokenManager.contract.UnpackLog(event, "SupplyLimitSet", log); err != nil {
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

// ParseSupplyLimitSet is a log parse operation binding the contract event 0x2f1ac7412e0c36ff7b73f7c4f474022dfa5d3f1c09e3262f63c9eb4627a9f505.
//
// Solidity: event SupplyLimitSet(address indexed changedBy, uint256 supplyLimit)
func (_NftTokenManager *NftTokenManagerFilterer) ParseSupplyLimitSet(log types.Log) (*NftTokenManagerSupplyLimitSet, error) {
	event := new(NftTokenManagerSupplyLimitSet)
	if err := _NftTokenManager.contract.UnpackLog(event, "SupplyLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftTokenManagerTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the NftTokenManager contract.
type NftTokenManagerTransferIterator struct {
	Event *NftTokenManagerTransfer // Event containing the contract specifics and raw log

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
func (it *NftTokenManagerTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftTokenManagerTransfer)
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
		it.Event = new(NftTokenManagerTransfer)
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
func (it *NftTokenManagerTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftTokenManagerTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftTokenManagerTransfer represents a Transfer event raised by the NftTokenManager contract.
type NftTokenManagerTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NftTokenManager *NftTokenManagerFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*NftTokenManagerTransferIterator, error) {

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

	logs, sub, err := _NftTokenManager.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftTokenManagerTransferIterator{contract: _NftTokenManager.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NftTokenManager *NftTokenManagerFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NftTokenManagerTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _NftTokenManager.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftTokenManagerTransfer)
				if err := _NftTokenManager.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_NftTokenManager *NftTokenManagerFilterer) ParseTransfer(log types.Log) (*NftTokenManagerTransfer, error) {
	event := new(NftTokenManagerTransfer)
	if err := _NftTokenManager.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftTokenManagerUriPrefixSetIterator is returned from FilterUriPrefixSet and is used to iterate over the raw logs and unpacked data for UriPrefixSet events raised by the NftTokenManager contract.
type NftTokenManagerUriPrefixSetIterator struct {
	Event *NftTokenManagerUriPrefixSet // Event containing the contract specifics and raw log

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
func (it *NftTokenManagerUriPrefixSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftTokenManagerUriPrefixSet)
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
		it.Event = new(NftTokenManagerUriPrefixSet)
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
func (it *NftTokenManagerUriPrefixSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftTokenManagerUriPrefixSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftTokenManagerUriPrefixSet represents a UriPrefixSet event raised by the NftTokenManager contract.
type NftTokenManagerUriPrefixSet struct {
	ChangedBy common.Address
	UrlPrefix string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUriPrefixSet is a free log retrieval operation binding the contract event 0x8aed2472df7eda2618a927a1e73256d01e814a5b03b0fff4931cc2774e39b0c6.
//
// Solidity: event UriPrefixSet(address indexed changedBy, string urlPrefix)
func (_NftTokenManager *NftTokenManagerFilterer) FilterUriPrefixSet(opts *bind.FilterOpts, changedBy []common.Address) (*NftTokenManagerUriPrefixSetIterator, error) {

	var changedByRule []interface{}
	for _, changedByItem := range changedBy {
		changedByRule = append(changedByRule, changedByItem)
	}

	logs, sub, err := _NftTokenManager.contract.FilterLogs(opts, "UriPrefixSet", changedByRule)
	if err != nil {
		return nil, err
	}
	return &NftTokenManagerUriPrefixSetIterator{contract: _NftTokenManager.contract, event: "UriPrefixSet", logs: logs, sub: sub}, nil
}

// WatchUriPrefixSet is a free log subscription operation binding the contract event 0x8aed2472df7eda2618a927a1e73256d01e814a5b03b0fff4931cc2774e39b0c6.
//
// Solidity: event UriPrefixSet(address indexed changedBy, string urlPrefix)
func (_NftTokenManager *NftTokenManagerFilterer) WatchUriPrefixSet(opts *bind.WatchOpts, sink chan<- *NftTokenManagerUriPrefixSet, changedBy []common.Address) (event.Subscription, error) {

	var changedByRule []interface{}
	for _, changedByItem := range changedBy {
		changedByRule = append(changedByRule, changedByItem)
	}

	logs, sub, err := _NftTokenManager.contract.WatchLogs(opts, "UriPrefixSet", changedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftTokenManagerUriPrefixSet)
				if err := _NftTokenManager.contract.UnpackLog(event, "UriPrefixSet", log); err != nil {
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
// Solidity: event UriPrefixSet(address indexed changedBy, string urlPrefix)
func (_NftTokenManager *NftTokenManagerFilterer) ParseUriPrefixSet(log types.Log) (*NftTokenManagerUriPrefixSet, error) {
	event := new(NftTokenManagerUriPrefixSet)
	if err := _NftTokenManager.contract.UnpackLog(event, "UriPrefixSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
