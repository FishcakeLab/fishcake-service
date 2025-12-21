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

// StakingManagerMetaData contains all meta data concerning the StakingManager contract.
var StakingManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aprOffset\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dayTimeStamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositIntoStaking\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"stakingType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isAutoRenew\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"fccAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feManagerAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIFishcakeEventManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakingAprFunding\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"messageNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"halfAprTimeStamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_fccAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_feManagerAddress\",\"type\":\"address\",\"internalType\":\"contractIFishcakeEventManager\"},{\"name\":\"_nftManagerAddress\",\"type\":\"address\",\"internalType\":\"contractINftManager\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockHalfYears\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockNinetyDays\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOneYears\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockSixtyDays\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockThirtyDays\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"messageNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minStakeAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nftManagerAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINftManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setHalfAprTimeStamp\",\"inputs\":[{\"name\":\"t\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalStakingAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdrawETHFromContract\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFromStakingWithAprIncome\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"messageNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Received\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeHolderDepositStaking\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"stakingType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"startStakingTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"endStakingTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"bindingNft\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nftApr\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isAutoRenew\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"messageNonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeHolderWithdrawStaking\",\"inputs\":[{\"name\":\"recipant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"withdrawAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"messageNonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"rewardAprFunding\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawETHFromContract\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FundingUnderStaking\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoFundingForStaking\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// StakingManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingManagerMetaData.ABI instead.
var StakingManagerABI = StakingManagerMetaData.ABI

// StakingManager is an auto generated Go binding around an Ethereum contract.
type StakingManager struct {
	StakingManagerCaller     // Read-only binding to the contract
	StakingManagerTransactor // Write-only binding to the contract
	StakingManagerFilterer   // Log filterer for contract events
}

// StakingManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingManagerSession struct {
	Contract     *StakingManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingManagerCallerSession struct {
	Contract *StakingManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StakingManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingManagerTransactorSession struct {
	Contract     *StakingManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StakingManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingManagerRaw struct {
	Contract *StakingManager // Generic contract binding to access the raw methods on
}

// StakingManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingManagerCallerRaw struct {
	Contract *StakingManagerCaller // Generic read-only contract binding to access the raw methods on
}

// StakingManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingManagerTransactorRaw struct {
	Contract *StakingManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingManager creates a new instance of StakingManager, bound to a specific deployed contract.
func NewStakingManager(address common.Address, backend bind.ContractBackend) (*StakingManager, error) {
	contract, err := bindStakingManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingManager{StakingManagerCaller: StakingManagerCaller{contract: contract}, StakingManagerTransactor: StakingManagerTransactor{contract: contract}, StakingManagerFilterer: StakingManagerFilterer{contract: contract}}, nil
}

// NewStakingManagerCaller creates a new read-only instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerCaller(address common.Address, caller bind.ContractCaller) (*StakingManagerCaller, error) {
	contract, err := bindStakingManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingManagerCaller{contract: contract}, nil
}

// NewStakingManagerTransactor creates a new write-only instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingManagerTransactor, error) {
	contract, err := bindStakingManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingManagerTransactor{contract: contract}, nil
}

// NewStakingManagerFilterer creates a new log filterer instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingManagerFilterer, error) {
	contract, err := bindStakingManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingManagerFilterer{contract: contract}, nil
}

// bindStakingManager binds a generic wrapper to an already deployed contract.
func bindStakingManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingManager *StakingManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingManager.Contract.StakingManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingManager *StakingManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.Contract.StakingManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingManager *StakingManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingManager.Contract.StakingManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingManager *StakingManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingManager *StakingManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingManager *StakingManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingManager.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_StakingManager *StakingManagerCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_StakingManager *StakingManagerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _StakingManager.Contract.UPGRADEINTERFACEVERSION(&_StakingManager.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_StakingManager *StakingManagerCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _StakingManager.Contract.UPGRADEINTERFACEVERSION(&_StakingManager.CallOpts)
}

// AprOffset is a free data retrieval call binding the contract method 0x94e2daba.
//
// Solidity: function aprOffset() view returns(uint256)
func (_StakingManager *StakingManagerCaller) AprOffset(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "aprOffset")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AprOffset is a free data retrieval call binding the contract method 0x94e2daba.
//
// Solidity: function aprOffset() view returns(uint256)
func (_StakingManager *StakingManagerSession) AprOffset() (*big.Int, error) {
	return _StakingManager.Contract.AprOffset(&_StakingManager.CallOpts)
}

// AprOffset is a free data retrieval call binding the contract method 0x94e2daba.
//
// Solidity: function aprOffset() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) AprOffset() (*big.Int, error) {
	return _StakingManager.Contract.AprOffset(&_StakingManager.CallOpts)
}

// DayTimeStamp is a free data retrieval call binding the contract method 0xc837d046.
//
// Solidity: function dayTimeStamp() view returns(uint256)
func (_StakingManager *StakingManagerCaller) DayTimeStamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "dayTimeStamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DayTimeStamp is a free data retrieval call binding the contract method 0xc837d046.
//
// Solidity: function dayTimeStamp() view returns(uint256)
func (_StakingManager *StakingManagerSession) DayTimeStamp() (*big.Int, error) {
	return _StakingManager.Contract.DayTimeStamp(&_StakingManager.CallOpts)
}

// DayTimeStamp is a free data retrieval call binding the contract method 0xc837d046.
//
// Solidity: function dayTimeStamp() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) DayTimeStamp() (*big.Int, error) {
	return _StakingManager.Contract.DayTimeStamp(&_StakingManager.CallOpts)
}

// FccAddress is a free data retrieval call binding the contract method 0x4bc92c86.
//
// Solidity: function fccAddress() view returns(address)
func (_StakingManager *StakingManagerCaller) FccAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "fccAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FccAddress is a free data retrieval call binding the contract method 0x4bc92c86.
//
// Solidity: function fccAddress() view returns(address)
func (_StakingManager *StakingManagerSession) FccAddress() (common.Address, error) {
	return _StakingManager.Contract.FccAddress(&_StakingManager.CallOpts)
}

// FccAddress is a free data retrieval call binding the contract method 0x4bc92c86.
//
// Solidity: function fccAddress() view returns(address)
func (_StakingManager *StakingManagerCallerSession) FccAddress() (common.Address, error) {
	return _StakingManager.Contract.FccAddress(&_StakingManager.CallOpts)
}

// FeManagerAddress is a free data retrieval call binding the contract method 0xee653473.
//
// Solidity: function feManagerAddress() view returns(address)
func (_StakingManager *StakingManagerCaller) FeManagerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "feManagerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeManagerAddress is a free data retrieval call binding the contract method 0xee653473.
//
// Solidity: function feManagerAddress() view returns(address)
func (_StakingManager *StakingManagerSession) FeManagerAddress() (common.Address, error) {
	return _StakingManager.Contract.FeManagerAddress(&_StakingManager.CallOpts)
}

// FeManagerAddress is a free data retrieval call binding the contract method 0xee653473.
//
// Solidity: function feManagerAddress() view returns(address)
func (_StakingManager *StakingManagerCallerSession) FeManagerAddress() (common.Address, error) {
	return _StakingManager.Contract.FeManagerAddress(&_StakingManager.CallOpts)
}

// GetStakingAprFunding is a free data retrieval call binding the contract method 0x95fbba0b.
//
// Solidity: function getStakingAprFunding(uint256 amount, uint256 messageNonce) view returns(uint256)
func (_StakingManager *StakingManagerCaller) GetStakingAprFunding(opts *bind.CallOpts, amount *big.Int, messageNonce *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getStakingAprFunding", amount, messageNonce)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingAprFunding is a free data retrieval call binding the contract method 0x95fbba0b.
//
// Solidity: function getStakingAprFunding(uint256 amount, uint256 messageNonce) view returns(uint256)
func (_StakingManager *StakingManagerSession) GetStakingAprFunding(amount *big.Int, messageNonce *big.Int) (*big.Int, error) {
	return _StakingManager.Contract.GetStakingAprFunding(&_StakingManager.CallOpts, amount, messageNonce)
}

// GetStakingAprFunding is a free data retrieval call binding the contract method 0x95fbba0b.
//
// Solidity: function getStakingAprFunding(uint256 amount, uint256 messageNonce) view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetStakingAprFunding(amount *big.Int, messageNonce *big.Int) (*big.Int, error) {
	return _StakingManager.Contract.GetStakingAprFunding(&_StakingManager.CallOpts, amount, messageNonce)
}

// HalfAprTimeStamp is a free data retrieval call binding the contract method 0xde8ad38f.
//
// Solidity: function halfAprTimeStamp() view returns(uint256)
func (_StakingManager *StakingManagerCaller) HalfAprTimeStamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "halfAprTimeStamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HalfAprTimeStamp is a free data retrieval call binding the contract method 0xde8ad38f.
//
// Solidity: function halfAprTimeStamp() view returns(uint256)
func (_StakingManager *StakingManagerSession) HalfAprTimeStamp() (*big.Int, error) {
	return _StakingManager.Contract.HalfAprTimeStamp(&_StakingManager.CallOpts)
}

// HalfAprTimeStamp is a free data retrieval call binding the contract method 0xde8ad38f.
//
// Solidity: function halfAprTimeStamp() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) HalfAprTimeStamp() (*big.Int, error) {
	return _StakingManager.Contract.HalfAprTimeStamp(&_StakingManager.CallOpts)
}

// LockHalfYears is a free data retrieval call binding the contract method 0xf8956f49.
//
// Solidity: function lockHalfYears() view returns(uint256)
func (_StakingManager *StakingManagerCaller) LockHalfYears(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "lockHalfYears")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockHalfYears is a free data retrieval call binding the contract method 0xf8956f49.
//
// Solidity: function lockHalfYears() view returns(uint256)
func (_StakingManager *StakingManagerSession) LockHalfYears() (*big.Int, error) {
	return _StakingManager.Contract.LockHalfYears(&_StakingManager.CallOpts)
}

// LockHalfYears is a free data retrieval call binding the contract method 0xf8956f49.
//
// Solidity: function lockHalfYears() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) LockHalfYears() (*big.Int, error) {
	return _StakingManager.Contract.LockHalfYears(&_StakingManager.CallOpts)
}

// LockNinetyDays is a free data retrieval call binding the contract method 0x58e5aca6.
//
// Solidity: function lockNinetyDays() view returns(uint256)
func (_StakingManager *StakingManagerCaller) LockNinetyDays(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "lockNinetyDays")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockNinetyDays is a free data retrieval call binding the contract method 0x58e5aca6.
//
// Solidity: function lockNinetyDays() view returns(uint256)
func (_StakingManager *StakingManagerSession) LockNinetyDays() (*big.Int, error) {
	return _StakingManager.Contract.LockNinetyDays(&_StakingManager.CallOpts)
}

// LockNinetyDays is a free data retrieval call binding the contract method 0x58e5aca6.
//
// Solidity: function lockNinetyDays() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) LockNinetyDays() (*big.Int, error) {
	return _StakingManager.Contract.LockNinetyDays(&_StakingManager.CallOpts)
}

// LockOneYears is a free data retrieval call binding the contract method 0x6a8049f8.
//
// Solidity: function lockOneYears() view returns(uint256)
func (_StakingManager *StakingManagerCaller) LockOneYears(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "lockOneYears")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockOneYears is a free data retrieval call binding the contract method 0x6a8049f8.
//
// Solidity: function lockOneYears() view returns(uint256)
func (_StakingManager *StakingManagerSession) LockOneYears() (*big.Int, error) {
	return _StakingManager.Contract.LockOneYears(&_StakingManager.CallOpts)
}

// LockOneYears is a free data retrieval call binding the contract method 0x6a8049f8.
//
// Solidity: function lockOneYears() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) LockOneYears() (*big.Int, error) {
	return _StakingManager.Contract.LockOneYears(&_StakingManager.CallOpts)
}

// LockSixtyDays is a free data retrieval call binding the contract method 0x1979e37f.
//
// Solidity: function lockSixtyDays() view returns(uint256)
func (_StakingManager *StakingManagerCaller) LockSixtyDays(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "lockSixtyDays")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockSixtyDays is a free data retrieval call binding the contract method 0x1979e37f.
//
// Solidity: function lockSixtyDays() view returns(uint256)
func (_StakingManager *StakingManagerSession) LockSixtyDays() (*big.Int, error) {
	return _StakingManager.Contract.LockSixtyDays(&_StakingManager.CallOpts)
}

// LockSixtyDays is a free data retrieval call binding the contract method 0x1979e37f.
//
// Solidity: function lockSixtyDays() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) LockSixtyDays() (*big.Int, error) {
	return _StakingManager.Contract.LockSixtyDays(&_StakingManager.CallOpts)
}

// LockThirtyDays is a free data retrieval call binding the contract method 0x41971eb2.
//
// Solidity: function lockThirtyDays() view returns(uint256)
func (_StakingManager *StakingManagerCaller) LockThirtyDays(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "lockThirtyDays")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockThirtyDays is a free data retrieval call binding the contract method 0x41971eb2.
//
// Solidity: function lockThirtyDays() view returns(uint256)
func (_StakingManager *StakingManagerSession) LockThirtyDays() (*big.Int, error) {
	return _StakingManager.Contract.LockThirtyDays(&_StakingManager.CallOpts)
}

// LockThirtyDays is a free data retrieval call binding the contract method 0x41971eb2.
//
// Solidity: function lockThirtyDays() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) LockThirtyDays() (*big.Int, error) {
	return _StakingManager.Contract.LockThirtyDays(&_StakingManager.CallOpts)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_StakingManager *StakingManagerCaller) MessageNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "messageNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_StakingManager *StakingManagerSession) MessageNonce() (*big.Int, error) {
	return _StakingManager.Contract.MessageNonce(&_StakingManager.CallOpts)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) MessageNonce() (*big.Int, error) {
	return _StakingManager.Contract.MessageNonce(&_StakingManager.CallOpts)
}

// MinStakeAmount is a free data retrieval call binding the contract method 0xf1887684.
//
// Solidity: function minStakeAmount() view returns(uint256)
func (_StakingManager *StakingManagerCaller) MinStakeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "minStakeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStakeAmount is a free data retrieval call binding the contract method 0xf1887684.
//
// Solidity: function minStakeAmount() view returns(uint256)
func (_StakingManager *StakingManagerSession) MinStakeAmount() (*big.Int, error) {
	return _StakingManager.Contract.MinStakeAmount(&_StakingManager.CallOpts)
}

// MinStakeAmount is a free data retrieval call binding the contract method 0xf1887684.
//
// Solidity: function minStakeAmount() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) MinStakeAmount() (*big.Int, error) {
	return _StakingManager.Contract.MinStakeAmount(&_StakingManager.CallOpts)
}

// NftManagerAddress is a free data retrieval call binding the contract method 0xd5aa93c8.
//
// Solidity: function nftManagerAddress() view returns(address)
func (_StakingManager *StakingManagerCaller) NftManagerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "nftManagerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftManagerAddress is a free data retrieval call binding the contract method 0xd5aa93c8.
//
// Solidity: function nftManagerAddress() view returns(address)
func (_StakingManager *StakingManagerSession) NftManagerAddress() (common.Address, error) {
	return _StakingManager.Contract.NftManagerAddress(&_StakingManager.CallOpts)
}

// NftManagerAddress is a free data retrieval call binding the contract method 0xd5aa93c8.
//
// Solidity: function nftManagerAddress() view returns(address)
func (_StakingManager *StakingManagerCallerSession) NftManagerAddress() (common.Address, error) {
	return _StakingManager.Contract.NftManagerAddress(&_StakingManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingManager *StakingManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingManager *StakingManagerSession) Owner() (common.Address, error) {
	return _StakingManager.Contract.Owner(&_StakingManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingManager *StakingManagerCallerSession) Owner() (common.Address, error) {
	return _StakingManager.Contract.Owner(&_StakingManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StakingManager *StakingManagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StakingManager *StakingManagerSession) ProxiableUUID() ([32]byte, error) {
	return _StakingManager.Contract.ProxiableUUID(&_StakingManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _StakingManager.Contract.ProxiableUUID(&_StakingManager.CallOpts)
}

// TotalStakingAmount is a free data retrieval call binding the contract method 0xd201114a.
//
// Solidity: function totalStakingAmount() view returns(uint256)
func (_StakingManager *StakingManagerCaller) TotalStakingAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "totalStakingAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStakingAmount is a free data retrieval call binding the contract method 0xd201114a.
//
// Solidity: function totalStakingAmount() view returns(uint256)
func (_StakingManager *StakingManagerSession) TotalStakingAmount() (*big.Int, error) {
	return _StakingManager.Contract.TotalStakingAmount(&_StakingManager.CallOpts)
}

// TotalStakingAmount is a free data retrieval call binding the contract method 0xd201114a.
//
// Solidity: function totalStakingAmount() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) TotalStakingAmount() (*big.Int, error) {
	return _StakingManager.Contract.TotalStakingAmount(&_StakingManager.CallOpts)
}

// DepositIntoStaking is a paid mutator transaction binding the contract method 0xeb44d775.
//
// Solidity: function depositIntoStaking(uint256 amount, uint8 stakingType, bool isAutoRenew, uint256 tokenId) returns()
func (_StakingManager *StakingManagerTransactor) DepositIntoStaking(opts *bind.TransactOpts, amount *big.Int, stakingType uint8, isAutoRenew bool, tokenId *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "depositIntoStaking", amount, stakingType, isAutoRenew, tokenId)
}

// DepositIntoStaking is a paid mutator transaction binding the contract method 0xeb44d775.
//
// Solidity: function depositIntoStaking(uint256 amount, uint8 stakingType, bool isAutoRenew, uint256 tokenId) returns()
func (_StakingManager *StakingManagerSession) DepositIntoStaking(amount *big.Int, stakingType uint8, isAutoRenew bool, tokenId *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.DepositIntoStaking(&_StakingManager.TransactOpts, amount, stakingType, isAutoRenew, tokenId)
}

// DepositIntoStaking is a paid mutator transaction binding the contract method 0xeb44d775.
//
// Solidity: function depositIntoStaking(uint256 amount, uint8 stakingType, bool isAutoRenew, uint256 tokenId) returns()
func (_StakingManager *StakingManagerTransactorSession) DepositIntoStaking(amount *big.Int, stakingType uint8, isAutoRenew bool, tokenId *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.DepositIntoStaking(&_StakingManager.TransactOpts, amount, stakingType, isAutoRenew, tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _initialOwner, address _fccAddress, address _feManagerAddress, address _nftManagerAddress) returns()
func (_StakingManager *StakingManagerTransactor) Initialize(opts *bind.TransactOpts, _initialOwner common.Address, _fccAddress common.Address, _feManagerAddress common.Address, _nftManagerAddress common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "initialize", _initialOwner, _fccAddress, _feManagerAddress, _nftManagerAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _initialOwner, address _fccAddress, address _feManagerAddress, address _nftManagerAddress) returns()
func (_StakingManager *StakingManagerSession) Initialize(_initialOwner common.Address, _fccAddress common.Address, _feManagerAddress common.Address, _nftManagerAddress common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.Initialize(&_StakingManager.TransactOpts, _initialOwner, _fccAddress, _feManagerAddress, _nftManagerAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _initialOwner, address _fccAddress, address _feManagerAddress, address _nftManagerAddress) returns()
func (_StakingManager *StakingManagerTransactorSession) Initialize(_initialOwner common.Address, _fccAddress common.Address, _feManagerAddress common.Address, _nftManagerAddress common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.Initialize(&_StakingManager.TransactOpts, _initialOwner, _fccAddress, _feManagerAddress, _nftManagerAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingManager *StakingManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingManager *StakingManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingManager.Contract.RenounceOwnership(&_StakingManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingManager *StakingManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingManager.Contract.RenounceOwnership(&_StakingManager.TransactOpts)
}

// SetHalfAprTimeStamp is a paid mutator transaction binding the contract method 0xf75e3a96.
//
// Solidity: function setHalfAprTimeStamp(uint256 t) returns()
func (_StakingManager *StakingManagerTransactor) SetHalfAprTimeStamp(opts *bind.TransactOpts, t *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setHalfAprTimeStamp", t)
}

// SetHalfAprTimeStamp is a paid mutator transaction binding the contract method 0xf75e3a96.
//
// Solidity: function setHalfAprTimeStamp(uint256 t) returns()
func (_StakingManager *StakingManagerSession) SetHalfAprTimeStamp(t *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetHalfAprTimeStamp(&_StakingManager.TransactOpts, t)
}

// SetHalfAprTimeStamp is a paid mutator transaction binding the contract method 0xf75e3a96.
//
// Solidity: function setHalfAprTimeStamp(uint256 t) returns()
func (_StakingManager *StakingManagerTransactorSession) SetHalfAprTimeStamp(t *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetHalfAprTimeStamp(&_StakingManager.TransactOpts, t)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingManager *StakingManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingManager *StakingManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.TransferOwnership(&_StakingManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingManager *StakingManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.TransferOwnership(&_StakingManager.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StakingManager *StakingManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StakingManager *StakingManagerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StakingManager.Contract.UpgradeToAndCall(&_StakingManager.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StakingManager *StakingManagerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StakingManager.Contract.UpgradeToAndCall(&_StakingManager.TransactOpts, newImplementation, data)
}

// WithdrawETHFromContract is a paid mutator transaction binding the contract method 0xa678129d.
//
// Solidity: function withdrawETHFromContract(address to, uint256 amount) returns()
func (_StakingManager *StakingManagerTransactor) WithdrawETHFromContract(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "withdrawETHFromContract", to, amount)
}

// WithdrawETHFromContract is a paid mutator transaction binding the contract method 0xa678129d.
//
// Solidity: function withdrawETHFromContract(address to, uint256 amount) returns()
func (_StakingManager *StakingManagerSession) WithdrawETHFromContract(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.WithdrawETHFromContract(&_StakingManager.TransactOpts, to, amount)
}

// WithdrawETHFromContract is a paid mutator transaction binding the contract method 0xa678129d.
//
// Solidity: function withdrawETHFromContract(address to, uint256 amount) returns()
func (_StakingManager *StakingManagerTransactorSession) WithdrawETHFromContract(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.WithdrawETHFromContract(&_StakingManager.TransactOpts, to, amount)
}

// WithdrawFromStakingWithAprIncome is a paid mutator transaction binding the contract method 0xed07f565.
//
// Solidity: function withdrawFromStakingWithAprIncome(uint256 amount, uint256 messageNonce) returns()
func (_StakingManager *StakingManagerTransactor) WithdrawFromStakingWithAprIncome(opts *bind.TransactOpts, amount *big.Int, messageNonce *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "withdrawFromStakingWithAprIncome", amount, messageNonce)
}

// WithdrawFromStakingWithAprIncome is a paid mutator transaction binding the contract method 0xed07f565.
//
// Solidity: function withdrawFromStakingWithAprIncome(uint256 amount, uint256 messageNonce) returns()
func (_StakingManager *StakingManagerSession) WithdrawFromStakingWithAprIncome(amount *big.Int, messageNonce *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.WithdrawFromStakingWithAprIncome(&_StakingManager.TransactOpts, amount, messageNonce)
}

// WithdrawFromStakingWithAprIncome is a paid mutator transaction binding the contract method 0xed07f565.
//
// Solidity: function withdrawFromStakingWithAprIncome(uint256 amount, uint256 messageNonce) returns()
func (_StakingManager *StakingManagerTransactorSession) WithdrawFromStakingWithAprIncome(amount *big.Int, messageNonce *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.WithdrawFromStakingWithAprIncome(&_StakingManager.TransactOpts, amount, messageNonce)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingManager *StakingManagerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingManager *StakingManagerSession) Receive() (*types.Transaction, error) {
	return _StakingManager.Contract.Receive(&_StakingManager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingManager *StakingManagerTransactorSession) Receive() (*types.Transaction, error) {
	return _StakingManager.Contract.Receive(&_StakingManager.TransactOpts)
}

// StakingManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StakingManager contract.
type StakingManagerInitializedIterator struct {
	Event *StakingManagerInitialized // Event containing the contract specifics and raw log

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
func (it *StakingManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerInitialized)
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
		it.Event = new(StakingManagerInitialized)
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
func (it *StakingManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerInitialized represents a Initialized event raised by the StakingManager contract.
type StakingManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StakingManager *StakingManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakingManagerInitializedIterator, error) {

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakingManagerInitializedIterator{contract: _StakingManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StakingManager *StakingManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakingManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerInitialized)
				if err := _StakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StakingManager *StakingManagerFilterer) ParseInitialized(log types.Log) (*StakingManagerInitialized, error) {
	event := new(StakingManagerInitialized)
	if err := _StakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakingManager contract.
type StakingManagerOwnershipTransferredIterator struct {
	Event *StakingManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerOwnershipTransferred)
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
		it.Event = new(StakingManagerOwnershipTransferred)
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
func (it *StakingManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerOwnershipTransferred represents a OwnershipTransferred event raised by the StakingManager contract.
type StakingManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingManager *StakingManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerOwnershipTransferredIterator{contract: _StakingManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingManager *StakingManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerOwnershipTransferred)
				if err := _StakingManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StakingManager *StakingManagerFilterer) ParseOwnershipTransferred(log types.Log) (*StakingManagerOwnershipTransferred, error) {
	event := new(StakingManagerOwnershipTransferred)
	if err := _StakingManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the StakingManager contract.
type StakingManagerReceivedIterator struct {
	Event *StakingManagerReceived // Event containing the contract specifics and raw log

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
func (it *StakingManagerReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerReceived)
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
		it.Event = new(StakingManagerReceived)
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
func (it *StakingManagerReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerReceived represents a Received event raised by the StakingManager contract.
type StakingManagerReceived struct {
	Receiver common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed receiver, uint256 _value)
func (_StakingManager *StakingManagerFilterer) FilterReceived(opts *bind.FilterOpts, receiver []common.Address) (*StakingManagerReceivedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Received", receiverRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerReceivedIterator{contract: _StakingManager.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed receiver, uint256 _value)
func (_StakingManager *StakingManagerFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *StakingManagerReceived, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Received", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerReceived)
				if err := _StakingManager.contract.UnpackLog(event, "Received", log); err != nil {
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
func (_StakingManager *StakingManagerFilterer) ParseReceived(log types.Log) (*StakingManagerReceived, error) {
	event := new(StakingManagerReceived)
	if err := _StakingManager.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerStakeHolderDepositStakingIterator is returned from FilterStakeHolderDepositStaking and is used to iterate over the raw logs and unpacked data for StakeHolderDepositStaking events raised by the StakingManager contract.
type StakingManagerStakeHolderDepositStakingIterator struct {
	Event *StakingManagerStakeHolderDepositStaking // Event containing the contract specifics and raw log

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
func (it *StakingManagerStakeHolderDepositStakingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerStakeHolderDepositStaking)
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
		it.Event = new(StakingManagerStakeHolderDepositStaking)
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
func (it *StakingManagerStakeHolderDepositStakingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerStakeHolderDepositStakingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerStakeHolderDepositStaking represents a StakeHolderDepositStaking event raised by the StakingManager contract.
type StakingManagerStakeHolderDepositStaking struct {
	Staker           common.Address
	Amount           *big.Int
	StakingType      uint8
	StartStakingTime *big.Int
	EndStakingTime   *big.Int
	BindingNft       *big.Int
	NftApr           *big.Int
	IsAutoRenew      bool
	MessageNonce     *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterStakeHolderDepositStaking is a free log retrieval operation binding the contract event 0x0c1c260f1025eef714f221f5eaf7f127b645e853b7bc9e9e6c57e06d2d0cf608.
//
// Solidity: event StakeHolderDepositStaking(address indexed staker, uint256 amount, uint8 stakingType, uint256 startStakingTime, uint256 endStakingTime, uint256 bindingNft, uint256 nftApr, bool isAutoRenew, uint256 messageNonce)
func (_StakingManager *StakingManagerFilterer) FilterStakeHolderDepositStaking(opts *bind.FilterOpts, staker []common.Address) (*StakingManagerStakeHolderDepositStakingIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "StakeHolderDepositStaking", stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerStakeHolderDepositStakingIterator{contract: _StakingManager.contract, event: "StakeHolderDepositStaking", logs: logs, sub: sub}, nil
}

// WatchStakeHolderDepositStaking is a free log subscription operation binding the contract event 0x0c1c260f1025eef714f221f5eaf7f127b645e853b7bc9e9e6c57e06d2d0cf608.
//
// Solidity: event StakeHolderDepositStaking(address indexed staker, uint256 amount, uint8 stakingType, uint256 startStakingTime, uint256 endStakingTime, uint256 bindingNft, uint256 nftApr, bool isAutoRenew, uint256 messageNonce)
func (_StakingManager *StakingManagerFilterer) WatchStakeHolderDepositStaking(opts *bind.WatchOpts, sink chan<- *StakingManagerStakeHolderDepositStaking, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "StakeHolderDepositStaking", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerStakeHolderDepositStaking)
				if err := _StakingManager.contract.UnpackLog(event, "StakeHolderDepositStaking", log); err != nil {
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

// ParseStakeHolderDepositStaking is a log parse operation binding the contract event 0x0c1c260f1025eef714f221f5eaf7f127b645e853b7bc9e9e6c57e06d2d0cf608.
//
// Solidity: event StakeHolderDepositStaking(address indexed staker, uint256 amount, uint8 stakingType, uint256 startStakingTime, uint256 endStakingTime, uint256 bindingNft, uint256 nftApr, bool isAutoRenew, uint256 messageNonce)
func (_StakingManager *StakingManagerFilterer) ParseStakeHolderDepositStaking(log types.Log) (*StakingManagerStakeHolderDepositStaking, error) {
	event := new(StakingManagerStakeHolderDepositStaking)
	if err := _StakingManager.contract.UnpackLog(event, "StakeHolderDepositStaking", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerStakeHolderWithdrawStakingIterator is returned from FilterStakeHolderWithdrawStaking and is used to iterate over the raw logs and unpacked data for StakeHolderWithdrawStaking events raised by the StakingManager contract.
type StakingManagerStakeHolderWithdrawStakingIterator struct {
	Event *StakingManagerStakeHolderWithdrawStaking // Event containing the contract specifics and raw log

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
func (it *StakingManagerStakeHolderWithdrawStakingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerStakeHolderWithdrawStaking)
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
		it.Event = new(StakingManagerStakeHolderWithdrawStaking)
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
func (it *StakingManagerStakeHolderWithdrawStakingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerStakeHolderWithdrawStakingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerStakeHolderWithdrawStaking represents a StakeHolderWithdrawStaking event raised by the StakingManager contract.
type StakingManagerStakeHolderWithdrawStaking struct {
	Recipant         common.Address
	WithdrawAmount   *big.Int
	MessageNonce     *big.Int
	MessageHash      [32]byte
	RewardAprFunding *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterStakeHolderWithdrawStaking is a free log retrieval operation binding the contract event 0x14040216a07a7cbe47a104adec603100a42f8d5ad1dab9b72982a67f8ac4ebda.
//
// Solidity: event StakeHolderWithdrawStaking(address indexed recipant, uint256 withdrawAmount, uint256 messageNonce, bytes32 messageHash, uint256 rewardAprFunding)
func (_StakingManager *StakingManagerFilterer) FilterStakeHolderWithdrawStaking(opts *bind.FilterOpts, recipant []common.Address) (*StakingManagerStakeHolderWithdrawStakingIterator, error) {

	var recipantRule []interface{}
	for _, recipantItem := range recipant {
		recipantRule = append(recipantRule, recipantItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "StakeHolderWithdrawStaking", recipantRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerStakeHolderWithdrawStakingIterator{contract: _StakingManager.contract, event: "StakeHolderWithdrawStaking", logs: logs, sub: sub}, nil
}

// WatchStakeHolderWithdrawStaking is a free log subscription operation binding the contract event 0x14040216a07a7cbe47a104adec603100a42f8d5ad1dab9b72982a67f8ac4ebda.
//
// Solidity: event StakeHolderWithdrawStaking(address indexed recipant, uint256 withdrawAmount, uint256 messageNonce, bytes32 messageHash, uint256 rewardAprFunding)
func (_StakingManager *StakingManagerFilterer) WatchStakeHolderWithdrawStaking(opts *bind.WatchOpts, sink chan<- *StakingManagerStakeHolderWithdrawStaking, recipant []common.Address) (event.Subscription, error) {

	var recipantRule []interface{}
	for _, recipantItem := range recipant {
		recipantRule = append(recipantRule, recipantItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "StakeHolderWithdrawStaking", recipantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerStakeHolderWithdrawStaking)
				if err := _StakingManager.contract.UnpackLog(event, "StakeHolderWithdrawStaking", log); err != nil {
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

// ParseStakeHolderWithdrawStaking is a log parse operation binding the contract event 0x14040216a07a7cbe47a104adec603100a42f8d5ad1dab9b72982a67f8ac4ebda.
//
// Solidity: event StakeHolderWithdrawStaking(address indexed recipant, uint256 withdrawAmount, uint256 messageNonce, bytes32 messageHash, uint256 rewardAprFunding)
func (_StakingManager *StakingManagerFilterer) ParseStakeHolderWithdrawStaking(log types.Log) (*StakingManagerStakeHolderWithdrawStaking, error) {
	event := new(StakingManagerStakeHolderWithdrawStaking)
	if err := _StakingManager.contract.UnpackLog(event, "StakeHolderWithdrawStaking", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the StakingManager contract.
type StakingManagerUpgradedIterator struct {
	Event *StakingManagerUpgraded // Event containing the contract specifics and raw log

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
func (it *StakingManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerUpgraded)
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
		it.Event = new(StakingManagerUpgraded)
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
func (it *StakingManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerUpgraded represents a Upgraded event raised by the StakingManager contract.
type StakingManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StakingManager *StakingManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*StakingManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerUpgradedIterator{contract: _StakingManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StakingManager *StakingManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *StakingManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerUpgraded)
				if err := _StakingManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StakingManager *StakingManagerFilterer) ParseUpgraded(log types.Log) (*StakingManagerUpgraded, error) {
	event := new(StakingManagerUpgraded)
	if err := _StakingManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerWithdrawETHFromContractIterator is returned from FilterWithdrawETHFromContract and is used to iterate over the raw logs and unpacked data for WithdrawETHFromContract events raised by the StakingManager contract.
type StakingManagerWithdrawETHFromContractIterator struct {
	Event *StakingManagerWithdrawETHFromContract // Event containing the contract specifics and raw log

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
func (it *StakingManagerWithdrawETHFromContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerWithdrawETHFromContract)
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
		it.Event = new(StakingManagerWithdrawETHFromContract)
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
func (it *StakingManagerWithdrawETHFromContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerWithdrawETHFromContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerWithdrawETHFromContract represents a WithdrawETHFromContract event raised by the StakingManager contract.
type StakingManagerWithdrawETHFromContract struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawETHFromContract is a free log retrieval operation binding the contract event 0x10bf05c0c841bf1cc64c2e5bb3f7beceba9587003356ef84b4670ea3fc982cce.
//
// Solidity: event WithdrawETHFromContract(address indexed to, uint256 amount)
func (_StakingManager *StakingManagerFilterer) FilterWithdrawETHFromContract(opts *bind.FilterOpts, to []common.Address) (*StakingManagerWithdrawETHFromContractIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "WithdrawETHFromContract", toRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerWithdrawETHFromContractIterator{contract: _StakingManager.contract, event: "WithdrawETHFromContract", logs: logs, sub: sub}, nil
}

// WatchWithdrawETHFromContract is a free log subscription operation binding the contract event 0x10bf05c0c841bf1cc64c2e5bb3f7beceba9587003356ef84b4670ea3fc982cce.
//
// Solidity: event WithdrawETHFromContract(address indexed to, uint256 amount)
func (_StakingManager *StakingManagerFilterer) WatchWithdrawETHFromContract(opts *bind.WatchOpts, sink chan<- *StakingManagerWithdrawETHFromContract, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "WithdrawETHFromContract", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerWithdrawETHFromContract)
				if err := _StakingManager.contract.UnpackLog(event, "WithdrawETHFromContract", log); err != nil {
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

// ParseWithdrawETHFromContract is a log parse operation binding the contract event 0x10bf05c0c841bf1cc64c2e5bb3f7beceba9587003356ef84b4670ea3fc982cce.
//
// Solidity: event WithdrawETHFromContract(address indexed to, uint256 amount)
func (_StakingManager *StakingManagerFilterer) ParseWithdrawETHFromContract(log types.Log) (*StakingManagerWithdrawETHFromContract, error) {
	event := new(StakingManagerWithdrawETHFromContract)
	if err := _StakingManager.contract.UnpackLog(event, "WithdrawETHFromContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
