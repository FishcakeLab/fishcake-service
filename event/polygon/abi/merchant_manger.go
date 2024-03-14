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

// MerchantMangerMetaData contains all meta data concerning the MerchantManger contract.
var MerchantMangerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"FccTokenAddr\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activityAdd\",\"inputs\":[{\"name\":\"_businessName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_activityContent\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_latitudeLongitude\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_activityDeadLine\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_totalDropAmts\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_dropType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_dropNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minDropAmt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_maxDropAmt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_tokenContractAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"_ret\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_activityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activityDropedToAccount\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activityFinish\",\"inputs\":[{\"name\":\"_activityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"_ret\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activityInfoArrs\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"activityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"businessAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"businessName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"activityContent\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"latitudeLongitude\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"activityCreateTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"activityDeadLine\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dropType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"dropNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minDropAmt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxDropAmt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenContractAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activityInfoChangedIdx\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activityInfoExtArrs\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"activityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"alreadyDropAmts\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"alreadyDropNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"businessMinedAmt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"businessMinedWithdrawedAmt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"activityStatus\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addMineAmt\",\"inputs\":[{\"name\":\"_addMineAmt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"_ret\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"drop\",\"inputs\":[{\"name\":\"_activityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_userAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_dropAmt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"_ret\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"dropInfoArrs\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"activityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"userAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dropTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dropAmt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"iNFTManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINFTManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_fcc\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_NFTManagerAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"minePercent\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minedAmt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinePercent\",\"inputs\":[{\"name\":\"_minePercent\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"_ret\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalMineAmt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"_ret\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawUToken\",\"inputs\":[{\"name\":\"_tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"_ret\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ActivityAdd\",\"inputs\":[{\"name\":\"who\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_activityId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"_totalDropAmts\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_businessName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"_activityContent\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"_latitudeLongitude\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"_activityDeadLine\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_dropType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"_dropNumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_minDropAmt\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_maxDropAmt\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_tokenContractAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ActivityFinish\",\"inputs\":[{\"name\":\"_activityId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AddMineAmt\",\"inputs\":[{\"name\":\"who\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_addMineAmt\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Drop\",\"inputs\":[{\"name\":\"who\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_activityId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"_dropAmt\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Received\",\"inputs\":[{\"name\":\"who\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetMinePercent\",\"inputs\":[{\"name\":\"minePercent\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetValidTime\",\"inputs\":[{\"name\":\"who\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_time\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawUToken\",\"inputs\":[{\"name\":\"who\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_tokenAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Wthdraw\",\"inputs\":[{\"name\":\"who\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// MerchantMangerABI is the input ABI used to generate the binding from.
// Deprecated: Use MerchantMangerMetaData.ABI instead.
var MerchantMangerABI = MerchantMangerMetaData.ABI

// MerchantManger is an auto generated Go binding around an Ethereum contract.
type MerchantManger struct {
	MerchantMangerCaller     // Read-only binding to the contract
	MerchantMangerTransactor // Write-only binding to the contract
	MerchantMangerFilterer   // Log filterer for contract events
}

// MerchantMangerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerchantMangerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerchantMangerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerchantMangerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerchantMangerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerchantMangerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerchantMangerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerchantMangerSession struct {
	Contract     *MerchantManger   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerchantMangerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerchantMangerCallerSession struct {
	Contract *MerchantMangerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MerchantMangerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerchantMangerTransactorSession struct {
	Contract     *MerchantMangerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MerchantMangerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerchantMangerRaw struct {
	Contract *MerchantManger // Generic contract binding to access the raw methods on
}

// MerchantMangerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerchantMangerCallerRaw struct {
	Contract *MerchantMangerCaller // Generic read-only contract binding to access the raw methods on
}

// MerchantMangerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerchantMangerTransactorRaw struct {
	Contract *MerchantMangerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerchantManger creates a new instance of MerchantManger, bound to a specific deployed contract.
func NewMerchantManger(address common.Address, backend bind.ContractBackend) (*MerchantManger, error) {
	contract, err := bindMerchantManger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerchantManger{MerchantMangerCaller: MerchantMangerCaller{contract: contract}, MerchantMangerTransactor: MerchantMangerTransactor{contract: contract}, MerchantMangerFilterer: MerchantMangerFilterer{contract: contract}}, nil
}

// NewMerchantMangerCaller creates a new read-only instance of MerchantManger, bound to a specific deployed contract.
func NewMerchantMangerCaller(address common.Address, caller bind.ContractCaller) (*MerchantMangerCaller, error) {
	contract, err := bindMerchantManger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerchantMangerCaller{contract: contract}, nil
}

// NewMerchantMangerTransactor creates a new write-only instance of MerchantManger, bound to a specific deployed contract.
func NewMerchantMangerTransactor(address common.Address, transactor bind.ContractTransactor) (*MerchantMangerTransactor, error) {
	contract, err := bindMerchantManger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerchantMangerTransactor{contract: contract}, nil
}

// NewMerchantMangerFilterer creates a new log filterer instance of MerchantManger, bound to a specific deployed contract.
func NewMerchantMangerFilterer(address common.Address, filterer bind.ContractFilterer) (*MerchantMangerFilterer, error) {
	contract, err := bindMerchantManger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerchantMangerFilterer{contract: contract}, nil
}

// bindMerchantManger binds a generic wrapper to an already deployed contract.
func bindMerchantManger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MerchantMangerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerchantManger *MerchantMangerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerchantManger.Contract.MerchantMangerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerchantManger *MerchantMangerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerchantManger.Contract.MerchantMangerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerchantManger *MerchantMangerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerchantManger.Contract.MerchantMangerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerchantManger *MerchantMangerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerchantManger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerchantManger *MerchantMangerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerchantManger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerchantManger *MerchantMangerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerchantManger.Contract.contract.Transact(opts, method, params...)
}

// FccTokenAddr is a free data retrieval call binding the contract method 0x4ba1332f.
//
// Solidity: function FccTokenAddr() view returns(address)
func (_MerchantManger *MerchantMangerCaller) FccTokenAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MerchantManger.contract.Call(opts, &out, "FccTokenAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FccTokenAddr is a free data retrieval call binding the contract method 0x4ba1332f.
//
// Solidity: function FccTokenAddr() view returns(address)
func (_MerchantManger *MerchantMangerSession) FccTokenAddr() (common.Address, error) {
	return _MerchantManger.Contract.FccTokenAddr(&_MerchantManger.CallOpts)
}

// FccTokenAddr is a free data retrieval call binding the contract method 0x4ba1332f.
//
// Solidity: function FccTokenAddr() view returns(address)
func (_MerchantManger *MerchantMangerCallerSession) FccTokenAddr() (common.Address, error) {
	return _MerchantManger.Contract.FccTokenAddr(&_MerchantManger.CallOpts)
}

// ActivityDropedToAccount is a free data retrieval call binding the contract method 0x85a455f5.
//
// Solidity: function activityDropedToAccount(uint256 , address ) view returns(bool)
func (_MerchantManger *MerchantMangerCaller) ActivityDropedToAccount(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _MerchantManger.contract.Call(opts, &out, "activityDropedToAccount", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ActivityDropedToAccount is a free data retrieval call binding the contract method 0x85a455f5.
//
// Solidity: function activityDropedToAccount(uint256 , address ) view returns(bool)
func (_MerchantManger *MerchantMangerSession) ActivityDropedToAccount(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _MerchantManger.Contract.ActivityDropedToAccount(&_MerchantManger.CallOpts, arg0, arg1)
}

// ActivityDropedToAccount is a free data retrieval call binding the contract method 0x85a455f5.
//
// Solidity: function activityDropedToAccount(uint256 , address ) view returns(bool)
func (_MerchantManger *MerchantMangerCallerSession) ActivityDropedToAccount(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _MerchantManger.Contract.ActivityDropedToAccount(&_MerchantManger.CallOpts, arg0, arg1)
}

// ActivityInfoArrs is a free data retrieval call binding the contract method 0xa0752825.
//
// Solidity: function activityInfoArrs(uint256 ) view returns(uint256 activityId, address businessAccount, string businessName, string activityContent, string latitudeLongitude, uint256 activityCreateTime, uint256 activityDeadLine, uint8 dropType, uint256 dropNumber, uint256 minDropAmt, uint256 maxDropAmt, address tokenContractAddr)
func (_MerchantManger *MerchantMangerCaller) ActivityInfoArrs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ActivityId         *big.Int
	BusinessAccount    common.Address
	BusinessName       string
	ActivityContent    string
	LatitudeLongitude  string
	ActivityCreateTime *big.Int
	ActivityDeadLine   *big.Int
	DropType           uint8
	DropNumber         *big.Int
	MinDropAmt         *big.Int
	MaxDropAmt         *big.Int
	TokenContractAddr  common.Address
}, error) {
	var out []interface{}
	err := _MerchantManger.contract.Call(opts, &out, "activityInfoArrs", arg0)

	outstruct := new(struct {
		ActivityId         *big.Int
		BusinessAccount    common.Address
		BusinessName       string
		ActivityContent    string
		LatitudeLongitude  string
		ActivityCreateTime *big.Int
		ActivityDeadLine   *big.Int
		DropType           uint8
		DropNumber         *big.Int
		MinDropAmt         *big.Int
		MaxDropAmt         *big.Int
		TokenContractAddr  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ActivityId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BusinessAccount = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.BusinessName = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ActivityContent = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.LatitudeLongitude = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.ActivityCreateTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.ActivityDeadLine = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.DropType = *abi.ConvertType(out[7], new(uint8)).(*uint8)
	outstruct.DropNumber = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.MinDropAmt = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.MaxDropAmt = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.TokenContractAddr = *abi.ConvertType(out[11], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// ActivityInfoArrs is a free data retrieval call binding the contract method 0xa0752825.
//
// Solidity: function activityInfoArrs(uint256 ) view returns(uint256 activityId, address businessAccount, string businessName, string activityContent, string latitudeLongitude, uint256 activityCreateTime, uint256 activityDeadLine, uint8 dropType, uint256 dropNumber, uint256 minDropAmt, uint256 maxDropAmt, address tokenContractAddr)
func (_MerchantManger *MerchantMangerSession) ActivityInfoArrs(arg0 *big.Int) (struct {
	ActivityId         *big.Int
	BusinessAccount    common.Address
	BusinessName       string
	ActivityContent    string
	LatitudeLongitude  string
	ActivityCreateTime *big.Int
	ActivityDeadLine   *big.Int
	DropType           uint8
	DropNumber         *big.Int
	MinDropAmt         *big.Int
	MaxDropAmt         *big.Int
	TokenContractAddr  common.Address
}, error) {
	return _MerchantManger.Contract.ActivityInfoArrs(&_MerchantManger.CallOpts, arg0)
}

// ActivityInfoArrs is a free data retrieval call binding the contract method 0xa0752825.
//
// Solidity: function activityInfoArrs(uint256 ) view returns(uint256 activityId, address businessAccount, string businessName, string activityContent, string latitudeLongitude, uint256 activityCreateTime, uint256 activityDeadLine, uint8 dropType, uint256 dropNumber, uint256 minDropAmt, uint256 maxDropAmt, address tokenContractAddr)
func (_MerchantManger *MerchantMangerCallerSession) ActivityInfoArrs(arg0 *big.Int) (struct {
	ActivityId         *big.Int
	BusinessAccount    common.Address
	BusinessName       string
	ActivityContent    string
	LatitudeLongitude  string
	ActivityCreateTime *big.Int
	ActivityDeadLine   *big.Int
	DropType           uint8
	DropNumber         *big.Int
	MinDropAmt         *big.Int
	MaxDropAmt         *big.Int
	TokenContractAddr  common.Address
}, error) {
	return _MerchantManger.Contract.ActivityInfoArrs(&_MerchantManger.CallOpts, arg0)
}

// ActivityInfoChangedIdx is a free data retrieval call binding the contract method 0xfefee0e1.
//
// Solidity: function activityInfoChangedIdx(uint256 ) view returns(uint256)
func (_MerchantManger *MerchantMangerCaller) ActivityInfoChangedIdx(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MerchantManger.contract.Call(opts, &out, "activityInfoChangedIdx", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActivityInfoChangedIdx is a free data retrieval call binding the contract method 0xfefee0e1.
//
// Solidity: function activityInfoChangedIdx(uint256 ) view returns(uint256)
func (_MerchantManger *MerchantMangerSession) ActivityInfoChangedIdx(arg0 *big.Int) (*big.Int, error) {
	return _MerchantManger.Contract.ActivityInfoChangedIdx(&_MerchantManger.CallOpts, arg0)
}

// ActivityInfoChangedIdx is a free data retrieval call binding the contract method 0xfefee0e1.
//
// Solidity: function activityInfoChangedIdx(uint256 ) view returns(uint256)
func (_MerchantManger *MerchantMangerCallerSession) ActivityInfoChangedIdx(arg0 *big.Int) (*big.Int, error) {
	return _MerchantManger.Contract.ActivityInfoChangedIdx(&_MerchantManger.CallOpts, arg0)
}

// ActivityInfoExtArrs is a free data retrieval call binding the contract method 0x4d43959d.
//
// Solidity: function activityInfoExtArrs(uint256 ) view returns(uint256 activityId, uint256 alreadyDropAmts, uint256 alreadyDropNumber, uint256 businessMinedAmt, uint256 businessMinedWithdrawedAmt, uint8 activityStatus)
func (_MerchantManger *MerchantMangerCaller) ActivityInfoExtArrs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ActivityId                 *big.Int
	AlreadyDropAmts            *big.Int
	AlreadyDropNumber          *big.Int
	BusinessMinedAmt           *big.Int
	BusinessMinedWithdrawedAmt *big.Int
	ActivityStatus             uint8
}, error) {
	var out []interface{}
	err := _MerchantManger.contract.Call(opts, &out, "activityInfoExtArrs", arg0)

	outstruct := new(struct {
		ActivityId                 *big.Int
		AlreadyDropAmts            *big.Int
		AlreadyDropNumber          *big.Int
		BusinessMinedAmt           *big.Int
		BusinessMinedWithdrawedAmt *big.Int
		ActivityStatus             uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ActivityId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AlreadyDropAmts = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AlreadyDropNumber = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.BusinessMinedAmt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.BusinessMinedWithdrawedAmt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.ActivityStatus = *abi.ConvertType(out[5], new(uint8)).(*uint8)

	return *outstruct, err

}

// ActivityInfoExtArrs is a free data retrieval call binding the contract method 0x4d43959d.
//
// Solidity: function activityInfoExtArrs(uint256 ) view returns(uint256 activityId, uint256 alreadyDropAmts, uint256 alreadyDropNumber, uint256 businessMinedAmt, uint256 businessMinedWithdrawedAmt, uint8 activityStatus)
func (_MerchantManger *MerchantMangerSession) ActivityInfoExtArrs(arg0 *big.Int) (struct {
	ActivityId                 *big.Int
	AlreadyDropAmts            *big.Int
	AlreadyDropNumber          *big.Int
	BusinessMinedAmt           *big.Int
	BusinessMinedWithdrawedAmt *big.Int
	ActivityStatus             uint8
}, error) {
	return _MerchantManger.Contract.ActivityInfoExtArrs(&_MerchantManger.CallOpts, arg0)
}

// ActivityInfoExtArrs is a free data retrieval call binding the contract method 0x4d43959d.
//
// Solidity: function activityInfoExtArrs(uint256 ) view returns(uint256 activityId, uint256 alreadyDropAmts, uint256 alreadyDropNumber, uint256 businessMinedAmt, uint256 businessMinedWithdrawedAmt, uint8 activityStatus)
func (_MerchantManger *MerchantMangerCallerSession) ActivityInfoExtArrs(arg0 *big.Int) (struct {
	ActivityId                 *big.Int
	AlreadyDropAmts            *big.Int
	AlreadyDropNumber          *big.Int
	BusinessMinedAmt           *big.Int
	BusinessMinedWithdrawedAmt *big.Int
	ActivityStatus             uint8
}, error) {
	return _MerchantManger.Contract.ActivityInfoExtArrs(&_MerchantManger.CallOpts, arg0)
}

// DropInfoArrs is a free data retrieval call binding the contract method 0x2a916c20.
//
// Solidity: function dropInfoArrs(uint256 ) view returns(uint256 activityId, address userAccount, uint256 dropTime, uint256 dropAmt)
func (_MerchantManger *MerchantMangerCaller) DropInfoArrs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ActivityId  *big.Int
	UserAccount common.Address
	DropTime    *big.Int
	DropAmt     *big.Int
}, error) {
	var out []interface{}
	err := _MerchantManger.contract.Call(opts, &out, "dropInfoArrs", arg0)

	outstruct := new(struct {
		ActivityId  *big.Int
		UserAccount common.Address
		DropTime    *big.Int
		DropAmt     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ActivityId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UserAccount = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.DropTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DropAmt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DropInfoArrs is a free data retrieval call binding the contract method 0x2a916c20.
//
// Solidity: function dropInfoArrs(uint256 ) view returns(uint256 activityId, address userAccount, uint256 dropTime, uint256 dropAmt)
func (_MerchantManger *MerchantMangerSession) DropInfoArrs(arg0 *big.Int) (struct {
	ActivityId  *big.Int
	UserAccount common.Address
	DropTime    *big.Int
	DropAmt     *big.Int
}, error) {
	return _MerchantManger.Contract.DropInfoArrs(&_MerchantManger.CallOpts, arg0)
}

// DropInfoArrs is a free data retrieval call binding the contract method 0x2a916c20.
//
// Solidity: function dropInfoArrs(uint256 ) view returns(uint256 activityId, address userAccount, uint256 dropTime, uint256 dropAmt)
func (_MerchantManger *MerchantMangerCallerSession) DropInfoArrs(arg0 *big.Int) (struct {
	ActivityId  *big.Int
	UserAccount common.Address
	DropTime    *big.Int
	DropAmt     *big.Int
}, error) {
	return _MerchantManger.Contract.DropInfoArrs(&_MerchantManger.CallOpts, arg0)
}

// INFTManager is a free data retrieval call binding the contract method 0xbbe68511.
//
// Solidity: function iNFTManager() view returns(address)
func (_MerchantManger *MerchantMangerCaller) INFTManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MerchantManger.contract.Call(opts, &out, "iNFTManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// INFTManager is a free data retrieval call binding the contract method 0xbbe68511.
//
// Solidity: function iNFTManager() view returns(address)
func (_MerchantManger *MerchantMangerSession) INFTManager() (common.Address, error) {
	return _MerchantManger.Contract.INFTManager(&_MerchantManger.CallOpts)
}

// INFTManager is a free data retrieval call binding the contract method 0xbbe68511.
//
// Solidity: function iNFTManager() view returns(address)
func (_MerchantManger *MerchantMangerCallerSession) INFTManager() (common.Address, error) {
	return _MerchantManger.Contract.INFTManager(&_MerchantManger.CallOpts)
}

// MinePercent is a free data retrieval call binding the contract method 0xa6d93915.
//
// Solidity: function minePercent() view returns(uint8)
func (_MerchantManger *MerchantMangerCaller) MinePercent(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MerchantManger.contract.Call(opts, &out, "minePercent")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MinePercent is a free data retrieval call binding the contract method 0xa6d93915.
//
// Solidity: function minePercent() view returns(uint8)
func (_MerchantManger *MerchantMangerSession) MinePercent() (uint8, error) {
	return _MerchantManger.Contract.MinePercent(&_MerchantManger.CallOpts)
}

// MinePercent is a free data retrieval call binding the contract method 0xa6d93915.
//
// Solidity: function minePercent() view returns(uint8)
func (_MerchantManger *MerchantMangerCallerSession) MinePercent() (uint8, error) {
	return _MerchantManger.Contract.MinePercent(&_MerchantManger.CallOpts)
}

// MinedAmt is a free data retrieval call binding the contract method 0xf7859c28.
//
// Solidity: function minedAmt() view returns(uint256)
func (_MerchantManger *MerchantMangerCaller) MinedAmt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MerchantManger.contract.Call(opts, &out, "minedAmt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinedAmt is a free data retrieval call binding the contract method 0xf7859c28.
//
// Solidity: function minedAmt() view returns(uint256)
func (_MerchantManger *MerchantMangerSession) MinedAmt() (*big.Int, error) {
	return _MerchantManger.Contract.MinedAmt(&_MerchantManger.CallOpts)
}

// MinedAmt is a free data retrieval call binding the contract method 0xf7859c28.
//
// Solidity: function minedAmt() view returns(uint256)
func (_MerchantManger *MerchantMangerCallerSession) MinedAmt() (*big.Int, error) {
	return _MerchantManger.Contract.MinedAmt(&_MerchantManger.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MerchantManger *MerchantMangerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MerchantManger.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MerchantManger *MerchantMangerSession) Owner() (common.Address, error) {
	return _MerchantManger.Contract.Owner(&_MerchantManger.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MerchantManger *MerchantMangerCallerSession) Owner() (common.Address, error) {
	return _MerchantManger.Contract.Owner(&_MerchantManger.CallOpts)
}

// TotalMineAmt is a free data retrieval call binding the contract method 0x88e40a5a.
//
// Solidity: function totalMineAmt() view returns(uint256)
func (_MerchantManger *MerchantMangerCaller) TotalMineAmt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MerchantManger.contract.Call(opts, &out, "totalMineAmt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMineAmt is a free data retrieval call binding the contract method 0x88e40a5a.
//
// Solidity: function totalMineAmt() view returns(uint256)
func (_MerchantManger *MerchantMangerSession) TotalMineAmt() (*big.Int, error) {
	return _MerchantManger.Contract.TotalMineAmt(&_MerchantManger.CallOpts)
}

// TotalMineAmt is a free data retrieval call binding the contract method 0x88e40a5a.
//
// Solidity: function totalMineAmt() view returns(uint256)
func (_MerchantManger *MerchantMangerCallerSession) TotalMineAmt() (*big.Int, error) {
	return _MerchantManger.Contract.TotalMineAmt(&_MerchantManger.CallOpts)
}

// ActivityAdd is a paid mutator transaction binding the contract method 0x43b22efc.
//
// Solidity: function activityAdd(string _businessName, string _activityContent, string _latitudeLongitude, uint256 _activityDeadLine, uint256 _totalDropAmts, uint8 _dropType, uint256 _dropNumber, uint256 _minDropAmt, uint256 _maxDropAmt, address _tokenContractAddr) returns(bool _ret, uint256 _activityId)
func (_MerchantManger *MerchantMangerTransactor) ActivityAdd(opts *bind.TransactOpts, _businessName string, _activityContent string, _latitudeLongitude string, _activityDeadLine *big.Int, _totalDropAmts *big.Int, _dropType uint8, _dropNumber *big.Int, _minDropAmt *big.Int, _maxDropAmt *big.Int, _tokenContractAddr common.Address) (*types.Transaction, error) {
	return _MerchantManger.contract.Transact(opts, "activityAdd", _businessName, _activityContent, _latitudeLongitude, _activityDeadLine, _totalDropAmts, _dropType, _dropNumber, _minDropAmt, _maxDropAmt, _tokenContractAddr)
}

// ActivityAdd is a paid mutator transaction binding the contract method 0x43b22efc.
//
// Solidity: function activityAdd(string _businessName, string _activityContent, string _latitudeLongitude, uint256 _activityDeadLine, uint256 _totalDropAmts, uint8 _dropType, uint256 _dropNumber, uint256 _minDropAmt, uint256 _maxDropAmt, address _tokenContractAddr) returns(bool _ret, uint256 _activityId)
func (_MerchantManger *MerchantMangerSession) ActivityAdd(_businessName string, _activityContent string, _latitudeLongitude string, _activityDeadLine *big.Int, _totalDropAmts *big.Int, _dropType uint8, _dropNumber *big.Int, _minDropAmt *big.Int, _maxDropAmt *big.Int, _tokenContractAddr common.Address) (*types.Transaction, error) {
	return _MerchantManger.Contract.ActivityAdd(&_MerchantManger.TransactOpts, _businessName, _activityContent, _latitudeLongitude, _activityDeadLine, _totalDropAmts, _dropType, _dropNumber, _minDropAmt, _maxDropAmt, _tokenContractAddr)
}

// ActivityAdd is a paid mutator transaction binding the contract method 0x43b22efc.
//
// Solidity: function activityAdd(string _businessName, string _activityContent, string _latitudeLongitude, uint256 _activityDeadLine, uint256 _totalDropAmts, uint8 _dropType, uint256 _dropNumber, uint256 _minDropAmt, uint256 _maxDropAmt, address _tokenContractAddr) returns(bool _ret, uint256 _activityId)
func (_MerchantManger *MerchantMangerTransactorSession) ActivityAdd(_businessName string, _activityContent string, _latitudeLongitude string, _activityDeadLine *big.Int, _totalDropAmts *big.Int, _dropType uint8, _dropNumber *big.Int, _minDropAmt *big.Int, _maxDropAmt *big.Int, _tokenContractAddr common.Address) (*types.Transaction, error) {
	return _MerchantManger.Contract.ActivityAdd(&_MerchantManger.TransactOpts, _businessName, _activityContent, _latitudeLongitude, _activityDeadLine, _totalDropAmts, _dropType, _dropNumber, _minDropAmt, _maxDropAmt, _tokenContractAddr)
}

// ActivityFinish is a paid mutator transaction binding the contract method 0xcbd596c4.
//
// Solidity: function activityFinish(uint256 _activityId) returns(bool _ret)
func (_MerchantManger *MerchantMangerTransactor) ActivityFinish(opts *bind.TransactOpts, _activityId *big.Int) (*types.Transaction, error) {
	return _MerchantManger.contract.Transact(opts, "activityFinish", _activityId)
}

// ActivityFinish is a paid mutator transaction binding the contract method 0xcbd596c4.
//
// Solidity: function activityFinish(uint256 _activityId) returns(bool _ret)
func (_MerchantManger *MerchantMangerSession) ActivityFinish(_activityId *big.Int) (*types.Transaction, error) {
	return _MerchantManger.Contract.ActivityFinish(&_MerchantManger.TransactOpts, _activityId)
}

// ActivityFinish is a paid mutator transaction binding the contract method 0xcbd596c4.
//
// Solidity: function activityFinish(uint256 _activityId) returns(bool _ret)
func (_MerchantManger *MerchantMangerTransactorSession) ActivityFinish(_activityId *big.Int) (*types.Transaction, error) {
	return _MerchantManger.Contract.ActivityFinish(&_MerchantManger.TransactOpts, _activityId)
}

// AddMineAmt is a paid mutator transaction binding the contract method 0x3c915255.
//
// Solidity: function addMineAmt(uint256 _addMineAmt) returns(bool _ret)
func (_MerchantManger *MerchantMangerTransactor) AddMineAmt(opts *bind.TransactOpts, _addMineAmt *big.Int) (*types.Transaction, error) {
	return _MerchantManger.contract.Transact(opts, "addMineAmt", _addMineAmt)
}

// AddMineAmt is a paid mutator transaction binding the contract method 0x3c915255.
//
// Solidity: function addMineAmt(uint256 _addMineAmt) returns(bool _ret)
func (_MerchantManger *MerchantMangerSession) AddMineAmt(_addMineAmt *big.Int) (*types.Transaction, error) {
	return _MerchantManger.Contract.AddMineAmt(&_MerchantManger.TransactOpts, _addMineAmt)
}

// AddMineAmt is a paid mutator transaction binding the contract method 0x3c915255.
//
// Solidity: function addMineAmt(uint256 _addMineAmt) returns(bool _ret)
func (_MerchantManger *MerchantMangerTransactorSession) AddMineAmt(_addMineAmt *big.Int) (*types.Transaction, error) {
	return _MerchantManger.Contract.AddMineAmt(&_MerchantManger.TransactOpts, _addMineAmt)
}

// Drop is a paid mutator transaction binding the contract method 0xc8c1d78d.
//
// Solidity: function drop(uint256 _activityId, address _userAccount, uint256 _dropAmt) returns(bool _ret)
func (_MerchantManger *MerchantMangerTransactor) Drop(opts *bind.TransactOpts, _activityId *big.Int, _userAccount common.Address, _dropAmt *big.Int) (*types.Transaction, error) {
	return _MerchantManger.contract.Transact(opts, "drop", _activityId, _userAccount, _dropAmt)
}

// Drop is a paid mutator transaction binding the contract method 0xc8c1d78d.
//
// Solidity: function drop(uint256 _activityId, address _userAccount, uint256 _dropAmt) returns(bool _ret)
func (_MerchantManger *MerchantMangerSession) Drop(_activityId *big.Int, _userAccount common.Address, _dropAmt *big.Int) (*types.Transaction, error) {
	return _MerchantManger.Contract.Drop(&_MerchantManger.TransactOpts, _activityId, _userAccount, _dropAmt)
}

// Drop is a paid mutator transaction binding the contract method 0xc8c1d78d.
//
// Solidity: function drop(uint256 _activityId, address _userAccount, uint256 _dropAmt) returns(bool _ret)
func (_MerchantManger *MerchantMangerTransactorSession) Drop(_activityId *big.Int, _userAccount common.Address, _dropAmt *big.Int) (*types.Transaction, error) {
	return _MerchantManger.Contract.Drop(&_MerchantManger.TransactOpts, _activityId, _userAccount, _dropAmt)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address initialOwner, address _fcc, address _NFTManagerAddr) returns()
func (_MerchantManger *MerchantMangerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _fcc common.Address, _NFTManagerAddr common.Address) (*types.Transaction, error) {
	return _MerchantManger.contract.Transact(opts, "initialize", initialOwner, _fcc, _NFTManagerAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address initialOwner, address _fcc, address _NFTManagerAddr) returns()
func (_MerchantManger *MerchantMangerSession) Initialize(initialOwner common.Address, _fcc common.Address, _NFTManagerAddr common.Address) (*types.Transaction, error) {
	return _MerchantManger.Contract.Initialize(&_MerchantManger.TransactOpts, initialOwner, _fcc, _NFTManagerAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address initialOwner, address _fcc, address _NFTManagerAddr) returns()
func (_MerchantManger *MerchantMangerTransactorSession) Initialize(initialOwner common.Address, _fcc common.Address, _NFTManagerAddr common.Address) (*types.Transaction, error) {
	return _MerchantManger.Contract.Initialize(&_MerchantManger.TransactOpts, initialOwner, _fcc, _NFTManagerAddr)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MerchantManger *MerchantMangerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerchantManger.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MerchantManger *MerchantMangerSession) RenounceOwnership() (*types.Transaction, error) {
	return _MerchantManger.Contract.RenounceOwnership(&_MerchantManger.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MerchantManger *MerchantMangerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MerchantManger.Contract.RenounceOwnership(&_MerchantManger.TransactOpts)
}

// SetMinePercent is a paid mutator transaction binding the contract method 0xac67fbb6.
//
// Solidity: function setMinePercent(uint8 _minePercent) returns(bool _ret)
func (_MerchantManger *MerchantMangerTransactor) SetMinePercent(opts *bind.TransactOpts, _minePercent uint8) (*types.Transaction, error) {
	return _MerchantManger.contract.Transact(opts, "setMinePercent", _minePercent)
}

// SetMinePercent is a paid mutator transaction binding the contract method 0xac67fbb6.
//
// Solidity: function setMinePercent(uint8 _minePercent) returns(bool _ret)
func (_MerchantManger *MerchantMangerSession) SetMinePercent(_minePercent uint8) (*types.Transaction, error) {
	return _MerchantManger.Contract.SetMinePercent(&_MerchantManger.TransactOpts, _minePercent)
}

// SetMinePercent is a paid mutator transaction binding the contract method 0xac67fbb6.
//
// Solidity: function setMinePercent(uint8 _minePercent) returns(bool _ret)
func (_MerchantManger *MerchantMangerTransactorSession) SetMinePercent(_minePercent uint8) (*types.Transaction, error) {
	return _MerchantManger.Contract.SetMinePercent(&_MerchantManger.TransactOpts, _minePercent)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MerchantManger *MerchantMangerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MerchantManger.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MerchantManger *MerchantMangerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MerchantManger.Contract.TransferOwnership(&_MerchantManger.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MerchantManger *MerchantMangerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MerchantManger.Contract.TransferOwnership(&_MerchantManger.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _recipient, uint256 _amount) returns(bool _ret)
func (_MerchantManger *MerchantMangerTransactor) Withdraw(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MerchantManger.contract.Transact(opts, "withdraw", _recipient, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _recipient, uint256 _amount) returns(bool _ret)
func (_MerchantManger *MerchantMangerSession) Withdraw(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MerchantManger.Contract.Withdraw(&_MerchantManger.TransactOpts, _recipient, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _recipient, uint256 _amount) returns(bool _ret)
func (_MerchantManger *MerchantMangerTransactorSession) Withdraw(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MerchantManger.Contract.Withdraw(&_MerchantManger.TransactOpts, _recipient, _amount)
}

// WithdrawUToken is a paid mutator transaction binding the contract method 0xde3f2b22.
//
// Solidity: function withdrawUToken(address _tokenAddr, address _account, uint256 _value) returns(bool _ret)
func (_MerchantManger *MerchantMangerTransactor) WithdrawUToken(opts *bind.TransactOpts, _tokenAddr common.Address, _account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MerchantManger.contract.Transact(opts, "withdrawUToken", _tokenAddr, _account, _value)
}

// WithdrawUToken is a paid mutator transaction binding the contract method 0xde3f2b22.
//
// Solidity: function withdrawUToken(address _tokenAddr, address _account, uint256 _value) returns(bool _ret)
func (_MerchantManger *MerchantMangerSession) WithdrawUToken(_tokenAddr common.Address, _account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MerchantManger.Contract.WithdrawUToken(&_MerchantManger.TransactOpts, _tokenAddr, _account, _value)
}

// WithdrawUToken is a paid mutator transaction binding the contract method 0xde3f2b22.
//
// Solidity: function withdrawUToken(address _tokenAddr, address _account, uint256 _value) returns(bool _ret)
func (_MerchantManger *MerchantMangerTransactorSession) WithdrawUToken(_tokenAddr common.Address, _account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MerchantManger.Contract.WithdrawUToken(&_MerchantManger.TransactOpts, _tokenAddr, _account, _value)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MerchantManger *MerchantMangerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerchantManger.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MerchantManger *MerchantMangerSession) Receive() (*types.Transaction, error) {
	return _MerchantManger.Contract.Receive(&_MerchantManger.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MerchantManger *MerchantMangerTransactorSession) Receive() (*types.Transaction, error) {
	return _MerchantManger.Contract.Receive(&_MerchantManger.TransactOpts)
}

// MerchantMangerActivityAddIterator is returned from FilterActivityAdd and is used to iterate over the raw logs and unpacked data for ActivityAdd events raised by the MerchantManger contract.
type MerchantMangerActivityAddIterator struct {
	Event *MerchantMangerActivityAdd // Event containing the contract specifics and raw log

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
func (it *MerchantMangerActivityAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantMangerActivityAdd)
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
		it.Event = new(MerchantMangerActivityAdd)
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
func (it *MerchantMangerActivityAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantMangerActivityAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantMangerActivityAdd represents a ActivityAdd event raised by the MerchantManger contract.
type MerchantMangerActivityAdd struct {
	Who               common.Address
	ActivityId        *big.Int
	TotalDropAmts     *big.Int
	BusinessName      string
	ActivityContent   string
	LatitudeLongitude string
	ActivityDeadLine  *big.Int
	DropType          uint8
	DropNumber        *big.Int
	MinDropAmt        *big.Int
	MaxDropAmt        *big.Int
	TokenContractAddr common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterActivityAdd is a free log retrieval operation binding the contract event 0xcc07277f6a9ae6659df5a850b317ba47d741264062fc7654d69b9fe11f4fabef.
//
// Solidity: event ActivityAdd(address indexed who, uint256 indexed _activityId, uint256 _totalDropAmts, string _businessName, string _activityContent, string _latitudeLongitude, uint256 _activityDeadLine, uint8 _dropType, uint256 _dropNumber, uint256 _minDropAmt, uint256 _maxDropAmt, address _tokenContractAddr)
func (_MerchantManger *MerchantMangerFilterer) FilterActivityAdd(opts *bind.FilterOpts, who []common.Address, _activityId []*big.Int) (*MerchantMangerActivityAddIterator, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}
	var _activityIdRule []interface{}
	for _, _activityIdItem := range _activityId {
		_activityIdRule = append(_activityIdRule, _activityIdItem)
	}

	logs, sub, err := _MerchantManger.contract.FilterLogs(opts, "ActivityAdd", whoRule, _activityIdRule)
	if err != nil {
		return nil, err
	}
	return &MerchantMangerActivityAddIterator{contract: _MerchantManger.contract, event: "ActivityAdd", logs: logs, sub: sub}, nil
}

// WatchActivityAdd is a free log subscription operation binding the contract event 0xcc07277f6a9ae6659df5a850b317ba47d741264062fc7654d69b9fe11f4fabef.
//
// Solidity: event ActivityAdd(address indexed who, uint256 indexed _activityId, uint256 _totalDropAmts, string _businessName, string _activityContent, string _latitudeLongitude, uint256 _activityDeadLine, uint8 _dropType, uint256 _dropNumber, uint256 _minDropAmt, uint256 _maxDropAmt, address _tokenContractAddr)
func (_MerchantManger *MerchantMangerFilterer) WatchActivityAdd(opts *bind.WatchOpts, sink chan<- *MerchantMangerActivityAdd, who []common.Address, _activityId []*big.Int) (event.Subscription, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}
	var _activityIdRule []interface{}
	for _, _activityIdItem := range _activityId {
		_activityIdRule = append(_activityIdRule, _activityIdItem)
	}

	logs, sub, err := _MerchantManger.contract.WatchLogs(opts, "ActivityAdd", whoRule, _activityIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantMangerActivityAdd)
				if err := _MerchantManger.contract.UnpackLog(event, "ActivityAdd", log); err != nil {
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

// ParseActivityAdd is a log parse operation binding the contract event 0xcc07277f6a9ae6659df5a850b317ba47d741264062fc7654d69b9fe11f4fabef.
//
// Solidity: event ActivityAdd(address indexed who, uint256 indexed _activityId, uint256 _totalDropAmts, string _businessName, string _activityContent, string _latitudeLongitude, uint256 _activityDeadLine, uint8 _dropType, uint256 _dropNumber, uint256 _minDropAmt, uint256 _maxDropAmt, address _tokenContractAddr)
func (_MerchantManger *MerchantMangerFilterer) ParseActivityAdd(log types.Log) (*MerchantMangerActivityAdd, error) {
	event := new(MerchantMangerActivityAdd)
	if err := _MerchantManger.contract.UnpackLog(event, "ActivityAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerchantMangerActivityFinishIterator is returned from FilterActivityFinish and is used to iterate over the raw logs and unpacked data for ActivityFinish events raised by the MerchantManger contract.
type MerchantMangerActivityFinishIterator struct {
	Event *MerchantMangerActivityFinish // Event containing the contract specifics and raw log

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
func (it *MerchantMangerActivityFinishIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantMangerActivityFinish)
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
		it.Event = new(MerchantMangerActivityFinish)
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
func (it *MerchantMangerActivityFinishIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantMangerActivityFinishIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantMangerActivityFinish represents a ActivityFinish event raised by the MerchantManger contract.
type MerchantMangerActivityFinish struct {
	ActivityId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterActivityFinish is a free log retrieval operation binding the contract event 0x38f456a137dd5646304c56c89456946ca67dce6874da9aa268339cb049f81771.
//
// Solidity: event ActivityFinish(uint256 indexed _activityId)
func (_MerchantManger *MerchantMangerFilterer) FilterActivityFinish(opts *bind.FilterOpts, _activityId []*big.Int) (*MerchantMangerActivityFinishIterator, error) {

	var _activityIdRule []interface{}
	for _, _activityIdItem := range _activityId {
		_activityIdRule = append(_activityIdRule, _activityIdItem)
	}

	logs, sub, err := _MerchantManger.contract.FilterLogs(opts, "ActivityFinish", _activityIdRule)
	if err != nil {
		return nil, err
	}
	return &MerchantMangerActivityFinishIterator{contract: _MerchantManger.contract, event: "ActivityFinish", logs: logs, sub: sub}, nil
}

// WatchActivityFinish is a free log subscription operation binding the contract event 0x38f456a137dd5646304c56c89456946ca67dce6874da9aa268339cb049f81771.
//
// Solidity: event ActivityFinish(uint256 indexed _activityId)
func (_MerchantManger *MerchantMangerFilterer) WatchActivityFinish(opts *bind.WatchOpts, sink chan<- *MerchantMangerActivityFinish, _activityId []*big.Int) (event.Subscription, error) {

	var _activityIdRule []interface{}
	for _, _activityIdItem := range _activityId {
		_activityIdRule = append(_activityIdRule, _activityIdItem)
	}

	logs, sub, err := _MerchantManger.contract.WatchLogs(opts, "ActivityFinish", _activityIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantMangerActivityFinish)
				if err := _MerchantManger.contract.UnpackLog(event, "ActivityFinish", log); err != nil {
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

// ParseActivityFinish is a log parse operation binding the contract event 0x38f456a137dd5646304c56c89456946ca67dce6874da9aa268339cb049f81771.
//
// Solidity: event ActivityFinish(uint256 indexed _activityId)
func (_MerchantManger *MerchantMangerFilterer) ParseActivityFinish(log types.Log) (*MerchantMangerActivityFinish, error) {
	event := new(MerchantMangerActivityFinish)
	if err := _MerchantManger.contract.UnpackLog(event, "ActivityFinish", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerchantMangerAddMineAmtIterator is returned from FilterAddMineAmt and is used to iterate over the raw logs and unpacked data for AddMineAmt events raised by the MerchantManger contract.
type MerchantMangerAddMineAmtIterator struct {
	Event *MerchantMangerAddMineAmt // Event containing the contract specifics and raw log

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
func (it *MerchantMangerAddMineAmtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantMangerAddMineAmt)
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
		it.Event = new(MerchantMangerAddMineAmt)
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
func (it *MerchantMangerAddMineAmtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantMangerAddMineAmtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantMangerAddMineAmt represents a AddMineAmt event raised by the MerchantManger contract.
type MerchantMangerAddMineAmt struct {
	Who        common.Address
	AddMineAmt *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddMineAmt is a free log retrieval operation binding the contract event 0x36a36f6d06ab1b164d3fbb904d2e7657b13fc5444f9ac18bc9a7af5aa6868376.
//
// Solidity: event AddMineAmt(address indexed who, uint256 _addMineAmt)
func (_MerchantManger *MerchantMangerFilterer) FilterAddMineAmt(opts *bind.FilterOpts, who []common.Address) (*MerchantMangerAddMineAmtIterator, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _MerchantManger.contract.FilterLogs(opts, "AddMineAmt", whoRule)
	if err != nil {
		return nil, err
	}
	return &MerchantMangerAddMineAmtIterator{contract: _MerchantManger.contract, event: "AddMineAmt", logs: logs, sub: sub}, nil
}

// WatchAddMineAmt is a free log subscription operation binding the contract event 0x36a36f6d06ab1b164d3fbb904d2e7657b13fc5444f9ac18bc9a7af5aa6868376.
//
// Solidity: event AddMineAmt(address indexed who, uint256 _addMineAmt)
func (_MerchantManger *MerchantMangerFilterer) WatchAddMineAmt(opts *bind.WatchOpts, sink chan<- *MerchantMangerAddMineAmt, who []common.Address) (event.Subscription, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _MerchantManger.contract.WatchLogs(opts, "AddMineAmt", whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantMangerAddMineAmt)
				if err := _MerchantManger.contract.UnpackLog(event, "AddMineAmt", log); err != nil {
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

// ParseAddMineAmt is a log parse operation binding the contract event 0x36a36f6d06ab1b164d3fbb904d2e7657b13fc5444f9ac18bc9a7af5aa6868376.
//
// Solidity: event AddMineAmt(address indexed who, uint256 _addMineAmt)
func (_MerchantManger *MerchantMangerFilterer) ParseAddMineAmt(log types.Log) (*MerchantMangerAddMineAmt, error) {
	event := new(MerchantMangerAddMineAmt)
	if err := _MerchantManger.contract.UnpackLog(event, "AddMineAmt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerchantMangerDropIterator is returned from FilterDrop and is used to iterate over the raw logs and unpacked data for Drop events raised by the MerchantManger contract.
type MerchantMangerDropIterator struct {
	Event *MerchantMangerDrop // Event containing the contract specifics and raw log

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
func (it *MerchantMangerDropIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantMangerDrop)
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
		it.Event = new(MerchantMangerDrop)
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
func (it *MerchantMangerDropIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantMangerDropIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantMangerDrop represents a Drop event raised by the MerchantManger contract.
type MerchantMangerDrop struct {
	Who        common.Address
	ActivityId *big.Int
	DropAmt    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDrop is a free log retrieval operation binding the contract event 0xd034d7fdd827eb4f41b668e02da1342cb8eced53d850fef72e1b2cd2c8387bac.
//
// Solidity: event Drop(address indexed who, uint256 indexed _activityId, uint256 _dropAmt)
func (_MerchantManger *MerchantMangerFilterer) FilterDrop(opts *bind.FilterOpts, who []common.Address, _activityId []*big.Int) (*MerchantMangerDropIterator, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}
	var _activityIdRule []interface{}
	for _, _activityIdItem := range _activityId {
		_activityIdRule = append(_activityIdRule, _activityIdItem)
	}

	logs, sub, err := _MerchantManger.contract.FilterLogs(opts, "Drop", whoRule, _activityIdRule)
	if err != nil {
		return nil, err
	}
	return &MerchantMangerDropIterator{contract: _MerchantManger.contract, event: "Drop", logs: logs, sub: sub}, nil
}

// WatchDrop is a free log subscription operation binding the contract event 0xd034d7fdd827eb4f41b668e02da1342cb8eced53d850fef72e1b2cd2c8387bac.
//
// Solidity: event Drop(address indexed who, uint256 indexed _activityId, uint256 _dropAmt)
func (_MerchantManger *MerchantMangerFilterer) WatchDrop(opts *bind.WatchOpts, sink chan<- *MerchantMangerDrop, who []common.Address, _activityId []*big.Int) (event.Subscription, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}
	var _activityIdRule []interface{}
	for _, _activityIdItem := range _activityId {
		_activityIdRule = append(_activityIdRule, _activityIdItem)
	}

	logs, sub, err := _MerchantManger.contract.WatchLogs(opts, "Drop", whoRule, _activityIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantMangerDrop)
				if err := _MerchantManger.contract.UnpackLog(event, "Drop", log); err != nil {
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

// ParseDrop is a log parse operation binding the contract event 0xd034d7fdd827eb4f41b668e02da1342cb8eced53d850fef72e1b2cd2c8387bac.
//
// Solidity: event Drop(address indexed who, uint256 indexed _activityId, uint256 _dropAmt)
func (_MerchantManger *MerchantMangerFilterer) ParseDrop(log types.Log) (*MerchantMangerDrop, error) {
	event := new(MerchantMangerDrop)
	if err := _MerchantManger.contract.UnpackLog(event, "Drop", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerchantMangerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MerchantManger contract.
type MerchantMangerInitializedIterator struct {
	Event *MerchantMangerInitialized // Event containing the contract specifics and raw log

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
func (it *MerchantMangerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantMangerInitialized)
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
		it.Event = new(MerchantMangerInitialized)
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
func (it *MerchantMangerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantMangerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantMangerInitialized represents a Initialized event raised by the MerchantManger contract.
type MerchantMangerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MerchantManger *MerchantMangerFilterer) FilterInitialized(opts *bind.FilterOpts) (*MerchantMangerInitializedIterator, error) {

	logs, sub, err := _MerchantManger.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MerchantMangerInitializedIterator{contract: _MerchantManger.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MerchantManger *MerchantMangerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MerchantMangerInitialized) (event.Subscription, error) {

	logs, sub, err := _MerchantManger.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantMangerInitialized)
				if err := _MerchantManger.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_MerchantManger *MerchantMangerFilterer) ParseInitialized(log types.Log) (*MerchantMangerInitialized, error) {
	event := new(MerchantMangerInitialized)
	if err := _MerchantManger.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerchantMangerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MerchantManger contract.
type MerchantMangerOwnershipTransferredIterator struct {
	Event *MerchantMangerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MerchantMangerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantMangerOwnershipTransferred)
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
		it.Event = new(MerchantMangerOwnershipTransferred)
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
func (it *MerchantMangerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantMangerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantMangerOwnershipTransferred represents a OwnershipTransferred event raised by the MerchantManger contract.
type MerchantMangerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MerchantManger *MerchantMangerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MerchantMangerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MerchantManger.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MerchantMangerOwnershipTransferredIterator{contract: _MerchantManger.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MerchantManger *MerchantMangerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MerchantMangerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MerchantManger.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantMangerOwnershipTransferred)
				if err := _MerchantManger.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MerchantManger *MerchantMangerFilterer) ParseOwnershipTransferred(log types.Log) (*MerchantMangerOwnershipTransferred, error) {
	event := new(MerchantMangerOwnershipTransferred)
	if err := _MerchantManger.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerchantMangerReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the MerchantManger contract.
type MerchantMangerReceivedIterator struct {
	Event *MerchantMangerReceived // Event containing the contract specifics and raw log

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
func (it *MerchantMangerReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantMangerReceived)
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
		it.Event = new(MerchantMangerReceived)
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
func (it *MerchantMangerReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantMangerReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantMangerReceived represents a Received event raised by the MerchantManger contract.
type MerchantMangerReceived struct {
	Who   common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed who, uint256 _value)
func (_MerchantManger *MerchantMangerFilterer) FilterReceived(opts *bind.FilterOpts, who []common.Address) (*MerchantMangerReceivedIterator, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _MerchantManger.contract.FilterLogs(opts, "Received", whoRule)
	if err != nil {
		return nil, err
	}
	return &MerchantMangerReceivedIterator{contract: _MerchantManger.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed who, uint256 _value)
func (_MerchantManger *MerchantMangerFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *MerchantMangerReceived, who []common.Address) (event.Subscription, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _MerchantManger.contract.WatchLogs(opts, "Received", whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantMangerReceived)
				if err := _MerchantManger.contract.UnpackLog(event, "Received", log); err != nil {
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
func (_MerchantManger *MerchantMangerFilterer) ParseReceived(log types.Log) (*MerchantMangerReceived, error) {
	event := new(MerchantMangerReceived)
	if err := _MerchantManger.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerchantMangerSetMinePercentIterator is returned from FilterSetMinePercent and is used to iterate over the raw logs and unpacked data for SetMinePercent events raised by the MerchantManger contract.
type MerchantMangerSetMinePercentIterator struct {
	Event *MerchantMangerSetMinePercent // Event containing the contract specifics and raw log

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
func (it *MerchantMangerSetMinePercentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantMangerSetMinePercent)
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
		it.Event = new(MerchantMangerSetMinePercent)
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
func (it *MerchantMangerSetMinePercentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantMangerSetMinePercentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantMangerSetMinePercent represents a SetMinePercent event raised by the MerchantManger contract.
type MerchantMangerSetMinePercent struct {
	MinePercent uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetMinePercent is a free log retrieval operation binding the contract event 0x3b2d1b3b0c360549745cd56590950679dbfca066aea0899c4a5a15cca3adfdf8.
//
// Solidity: event SetMinePercent(uint8 minePercent)
func (_MerchantManger *MerchantMangerFilterer) FilterSetMinePercent(opts *bind.FilterOpts) (*MerchantMangerSetMinePercentIterator, error) {

	logs, sub, err := _MerchantManger.contract.FilterLogs(opts, "SetMinePercent")
	if err != nil {
		return nil, err
	}
	return &MerchantMangerSetMinePercentIterator{contract: _MerchantManger.contract, event: "SetMinePercent", logs: logs, sub: sub}, nil
}

// WatchSetMinePercent is a free log subscription operation binding the contract event 0x3b2d1b3b0c360549745cd56590950679dbfca066aea0899c4a5a15cca3adfdf8.
//
// Solidity: event SetMinePercent(uint8 minePercent)
func (_MerchantManger *MerchantMangerFilterer) WatchSetMinePercent(opts *bind.WatchOpts, sink chan<- *MerchantMangerSetMinePercent) (event.Subscription, error) {

	logs, sub, err := _MerchantManger.contract.WatchLogs(opts, "SetMinePercent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantMangerSetMinePercent)
				if err := _MerchantManger.contract.UnpackLog(event, "SetMinePercent", log); err != nil {
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

// ParseSetMinePercent is a log parse operation binding the contract event 0x3b2d1b3b0c360549745cd56590950679dbfca066aea0899c4a5a15cca3adfdf8.
//
// Solidity: event SetMinePercent(uint8 minePercent)
func (_MerchantManger *MerchantMangerFilterer) ParseSetMinePercent(log types.Log) (*MerchantMangerSetMinePercent, error) {
	event := new(MerchantMangerSetMinePercent)
	if err := _MerchantManger.contract.UnpackLog(event, "SetMinePercent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerchantMangerSetValidTimeIterator is returned from FilterSetValidTime and is used to iterate over the raw logs and unpacked data for SetValidTime events raised by the MerchantManger contract.
type MerchantMangerSetValidTimeIterator struct {
	Event *MerchantMangerSetValidTime // Event containing the contract specifics and raw log

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
func (it *MerchantMangerSetValidTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantMangerSetValidTime)
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
		it.Event = new(MerchantMangerSetValidTime)
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
func (it *MerchantMangerSetValidTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantMangerSetValidTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantMangerSetValidTime represents a SetValidTime event raised by the MerchantManger contract.
type MerchantMangerSetValidTime struct {
	Who  common.Address
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetValidTime is a free log retrieval operation binding the contract event 0x8651cd625889c15ab8587456aba8cc25b781c5fa977a1811f04967c2acad02b8.
//
// Solidity: event SetValidTime(address indexed who, uint256 _time)
func (_MerchantManger *MerchantMangerFilterer) FilterSetValidTime(opts *bind.FilterOpts, who []common.Address) (*MerchantMangerSetValidTimeIterator, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _MerchantManger.contract.FilterLogs(opts, "SetValidTime", whoRule)
	if err != nil {
		return nil, err
	}
	return &MerchantMangerSetValidTimeIterator{contract: _MerchantManger.contract, event: "SetValidTime", logs: logs, sub: sub}, nil
}

// WatchSetValidTime is a free log subscription operation binding the contract event 0x8651cd625889c15ab8587456aba8cc25b781c5fa977a1811f04967c2acad02b8.
//
// Solidity: event SetValidTime(address indexed who, uint256 _time)
func (_MerchantManger *MerchantMangerFilterer) WatchSetValidTime(opts *bind.WatchOpts, sink chan<- *MerchantMangerSetValidTime, who []common.Address) (event.Subscription, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _MerchantManger.contract.WatchLogs(opts, "SetValidTime", whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantMangerSetValidTime)
				if err := _MerchantManger.contract.UnpackLog(event, "SetValidTime", log); err != nil {
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
func (_MerchantManger *MerchantMangerFilterer) ParseSetValidTime(log types.Log) (*MerchantMangerSetValidTime, error) {
	event := new(MerchantMangerSetValidTime)
	if err := _MerchantManger.contract.UnpackLog(event, "SetValidTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerchantMangerWithdrawUTokenIterator is returned from FilterWithdrawUToken and is used to iterate over the raw logs and unpacked data for WithdrawUToken events raised by the MerchantManger contract.
type MerchantMangerWithdrawUTokenIterator struct {
	Event *MerchantMangerWithdrawUToken // Event containing the contract specifics and raw log

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
func (it *MerchantMangerWithdrawUTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantMangerWithdrawUToken)
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
		it.Event = new(MerchantMangerWithdrawUToken)
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
func (it *MerchantMangerWithdrawUTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantMangerWithdrawUTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantMangerWithdrawUToken represents a WithdrawUToken event raised by the MerchantManger contract.
type MerchantMangerWithdrawUToken struct {
	Who       common.Address
	TokenAddr common.Address
	Account   common.Address
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawUToken is a free log retrieval operation binding the contract event 0xf314360eddb258cb001ab9dd025a30d7be229f08113f1e5617a7f82a00e0e78f.
//
// Solidity: event WithdrawUToken(address indexed who, address indexed _tokenAddr, address indexed _account, uint256 _value)
func (_MerchantManger *MerchantMangerFilterer) FilterWithdrawUToken(opts *bind.FilterOpts, who []common.Address, _tokenAddr []common.Address, _account []common.Address) (*MerchantMangerWithdrawUTokenIterator, error) {

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

	logs, sub, err := _MerchantManger.contract.FilterLogs(opts, "WithdrawUToken", whoRule, _tokenAddrRule, _accountRule)
	if err != nil {
		return nil, err
	}
	return &MerchantMangerWithdrawUTokenIterator{contract: _MerchantManger.contract, event: "WithdrawUToken", logs: logs, sub: sub}, nil
}

// WatchWithdrawUToken is a free log subscription operation binding the contract event 0xf314360eddb258cb001ab9dd025a30d7be229f08113f1e5617a7f82a00e0e78f.
//
// Solidity: event WithdrawUToken(address indexed who, address indexed _tokenAddr, address indexed _account, uint256 _value)
func (_MerchantManger *MerchantMangerFilterer) WatchWithdrawUToken(opts *bind.WatchOpts, sink chan<- *MerchantMangerWithdrawUToken, who []common.Address, _tokenAddr []common.Address, _account []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _MerchantManger.contract.WatchLogs(opts, "WithdrawUToken", whoRule, _tokenAddrRule, _accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantMangerWithdrawUToken)
				if err := _MerchantManger.contract.UnpackLog(event, "WithdrawUToken", log); err != nil {
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
func (_MerchantManger *MerchantMangerFilterer) ParseWithdrawUToken(log types.Log) (*MerchantMangerWithdrawUToken, error) {
	event := new(MerchantMangerWithdrawUToken)
	if err := _MerchantManger.contract.UnpackLog(event, "WithdrawUToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerchantMangerWthdrawIterator is returned from FilterWthdraw and is used to iterate over the raw logs and unpacked data for Wthdraw events raised by the MerchantManger contract.
type MerchantMangerWthdrawIterator struct {
	Event *MerchantMangerWthdraw // Event containing the contract specifics and raw log

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
func (it *MerchantMangerWthdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantMangerWthdraw)
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
		it.Event = new(MerchantMangerWthdraw)
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
func (it *MerchantMangerWthdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantMangerWthdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantMangerWthdraw represents a Wthdraw event raised by the MerchantManger contract.
type MerchantMangerWthdraw struct {
	Who    common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWthdraw is a free log retrieval operation binding the contract event 0x622ece8fe6c9adcdeb3aacfcf426facb5e94734cb9eb4cca7b66595853b87340.
//
// Solidity: event Wthdraw(address indexed who, uint256 _amount)
func (_MerchantManger *MerchantMangerFilterer) FilterWthdraw(opts *bind.FilterOpts, who []common.Address) (*MerchantMangerWthdrawIterator, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _MerchantManger.contract.FilterLogs(opts, "Wthdraw", whoRule)
	if err != nil {
		return nil, err
	}
	return &MerchantMangerWthdrawIterator{contract: _MerchantManger.contract, event: "Wthdraw", logs: logs, sub: sub}, nil
}

// WatchWthdraw is a free log subscription operation binding the contract event 0x622ece8fe6c9adcdeb3aacfcf426facb5e94734cb9eb4cca7b66595853b87340.
//
// Solidity: event Wthdraw(address indexed who, uint256 _amount)
func (_MerchantManger *MerchantMangerFilterer) WatchWthdraw(opts *bind.WatchOpts, sink chan<- *MerchantMangerWthdraw, who []common.Address) (event.Subscription, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _MerchantManger.contract.WatchLogs(opts, "Wthdraw", whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantMangerWthdraw)
				if err := _MerchantManger.contract.UnpackLog(event, "Wthdraw", log); err != nil {
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
func (_MerchantManger *MerchantMangerFilterer) ParseWthdraw(log types.Log) (*MerchantMangerWthdraw, error) {
	event := new(MerchantMangerWthdraw)
	if err := _MerchantManger.contract.UnpackLog(event, "Wthdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
