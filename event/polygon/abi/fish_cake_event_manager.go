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

// FishcakeEventManagerMetaData contains all meta data concerning the FishcakeEventManager contract.
var FishcakeEventManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"FccTokenAddr\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NTFLastMineTime\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UsdtTokenAddr\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activityAdd\",\"inputs\":[{\"name\":\"_businessName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_activityContent\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_latitudeLongitude\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_activityDeadLine\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_totalDropAmts\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_dropType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_dropNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minDropAmt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_maxDropAmt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_tokenContractAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activityDroppedToAccount\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activityFinish\",\"inputs\":[{\"name\":\"_activityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activityInfoArrs\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"activityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"businessAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"businessName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"activityContent\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"latitudeLongitude\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"activityCreateTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"activityDeadLine\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dropType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"dropNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minDropAmt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxDropAmt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenContractAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activityInfoChangedIdx\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activityInfoExtArrs\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"activityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"alreadyDropAmts\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"alreadyDropNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"businessMinedAmt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"businessMinedWithdrawedAmt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"activityStatus\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deleteMinedFishcakePower\",\"inputs\":[{\"name\":\"_miner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deleteMinerMineAmount\",\"inputs\":[{\"name\":\"_miner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"drop\",\"inputs\":[{\"name\":\"_activityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_userAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_dropAmt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"dropInfoArrs\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"activityId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"userAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dropTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dropAmt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinedFishcakePower\",\"inputs\":[{\"name\":\"_miner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinerMineAmount\",\"inputs\":[{\"name\":\"_miner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"iNFTManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINftManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_fccAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdtTokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_NFTManagerAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isMint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxDeadLine\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"merchantOnceMaxMineAmt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minePercent\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minedAmt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minedFishcakePower\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minerMineAmount\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"oneDay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalMineAmt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userOnceMaxMineAmt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ActivityAdd\",\"inputs\":[{\"name\":\"who\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_activityId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"_totalDropAmts\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_businessName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"_activityContent\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"_latitudeLongitude\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"_activityDeadLine\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_dropType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"_dropNumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_minDropAmt\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_maxDropAmt\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_tokenContractAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ActivityFinish\",\"inputs\":[{\"name\":\"_activityId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"_tokenContractAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_returnAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_minedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AddMineAmt\",\"inputs\":[{\"name\":\"who\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_addMineAmt\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Drop\",\"inputs\":[{\"name\":\"who\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_activityId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"_dropAmt\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetMinePercent\",\"inputs\":[{\"name\":\"minePercent\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetValidTime\",\"inputs\":[{\"name\":\"who\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_time\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
}

// FishcakeEventManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use FishcakeEventManagerMetaData.ABI instead.
var FishcakeEventManagerABI = FishcakeEventManagerMetaData.ABI

// FishcakeEventManager is an auto generated Go binding around an Ethereum contract.
type FishcakeEventManager struct {
	FishcakeEventManagerCaller     // Read-only binding to the contract
	FishcakeEventManagerTransactor // Write-only binding to the contract
	FishcakeEventManagerFilterer   // Log filterer for contract events
}

// FishcakeEventManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type FishcakeEventManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FishcakeEventManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FishcakeEventManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FishcakeEventManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FishcakeEventManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FishcakeEventManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FishcakeEventManagerSession struct {
	Contract     *FishcakeEventManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FishcakeEventManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FishcakeEventManagerCallerSession struct {
	Contract *FishcakeEventManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// FishcakeEventManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FishcakeEventManagerTransactorSession struct {
	Contract     *FishcakeEventManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// FishcakeEventManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type FishcakeEventManagerRaw struct {
	Contract *FishcakeEventManager // Generic contract binding to access the raw methods on
}

// FishcakeEventManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FishcakeEventManagerCallerRaw struct {
	Contract *FishcakeEventManagerCaller // Generic read-only contract binding to access the raw methods on
}

// FishcakeEventManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FishcakeEventManagerTransactorRaw struct {
	Contract *FishcakeEventManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFishcakeEventManager creates a new instance of FishcakeEventManager, bound to a specific deployed contract.
func NewFishcakeEventManager(address common.Address, backend bind.ContractBackend) (*FishcakeEventManager, error) {
	contract, err := bindFishcakeEventManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FishcakeEventManager{FishcakeEventManagerCaller: FishcakeEventManagerCaller{contract: contract}, FishcakeEventManagerTransactor: FishcakeEventManagerTransactor{contract: contract}, FishcakeEventManagerFilterer: FishcakeEventManagerFilterer{contract: contract}}, nil
}

// NewFishcakeEventManagerCaller creates a new read-only instance of FishcakeEventManager, bound to a specific deployed contract.
func NewFishcakeEventManagerCaller(address common.Address, caller bind.ContractCaller) (*FishcakeEventManagerCaller, error) {
	contract, err := bindFishcakeEventManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FishcakeEventManagerCaller{contract: contract}, nil
}

// NewFishcakeEventManagerTransactor creates a new write-only instance of FishcakeEventManager, bound to a specific deployed contract.
func NewFishcakeEventManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*FishcakeEventManagerTransactor, error) {
	contract, err := bindFishcakeEventManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FishcakeEventManagerTransactor{contract: contract}, nil
}

// NewFishcakeEventManagerFilterer creates a new log filterer instance of FishcakeEventManager, bound to a specific deployed contract.
func NewFishcakeEventManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*FishcakeEventManagerFilterer, error) {
	contract, err := bindFishcakeEventManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FishcakeEventManagerFilterer{contract: contract}, nil
}

// bindFishcakeEventManager binds a generic wrapper to an already deployed contract.
func bindFishcakeEventManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FishcakeEventManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FishcakeEventManager *FishcakeEventManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FishcakeEventManager.Contract.FishcakeEventManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FishcakeEventManager *FishcakeEventManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FishcakeEventManager.Contract.FishcakeEventManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FishcakeEventManager *FishcakeEventManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FishcakeEventManager.Contract.FishcakeEventManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FishcakeEventManager *FishcakeEventManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FishcakeEventManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FishcakeEventManager *FishcakeEventManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FishcakeEventManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FishcakeEventManager *FishcakeEventManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FishcakeEventManager.Contract.contract.Transact(opts, method, params...)
}

// FccTokenAddr is a free data retrieval call binding the contract method 0x4ba1332f.
//
// Solidity: function FccTokenAddr() view returns(address)
func (_FishcakeEventManager *FishcakeEventManagerCaller) FccTokenAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FishcakeEventManager.contract.Call(opts, &out, "FccTokenAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FccTokenAddr is a free data retrieval call binding the contract method 0x4ba1332f.
//
// Solidity: function FccTokenAddr() view returns(address)
func (_FishcakeEventManager *FishcakeEventManagerSession) FccTokenAddr() (common.Address, error) {
	return _FishcakeEventManager.Contract.FccTokenAddr(&_FishcakeEventManager.CallOpts)
}

// FccTokenAddr is a free data retrieval call binding the contract method 0x4ba1332f.
//
// Solidity: function FccTokenAddr() view returns(address)
func (_FishcakeEventManager *FishcakeEventManagerCallerSession) FccTokenAddr() (common.Address, error) {
	return _FishcakeEventManager.Contract.FccTokenAddr(&_FishcakeEventManager.CallOpts)
}

// NTFLastMineTime is a free data retrieval call binding the contract method 0x13f34fed.
//
// Solidity: function NTFLastMineTime(address ) view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerCaller) NTFLastMineTime(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FishcakeEventManager.contract.Call(opts, &out, "NTFLastMineTime", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NTFLastMineTime is a free data retrieval call binding the contract method 0x13f34fed.
//
// Solidity: function NTFLastMineTime(address ) view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerSession) NTFLastMineTime(arg0 common.Address) (*big.Int, error) {
	return _FishcakeEventManager.Contract.NTFLastMineTime(&_FishcakeEventManager.CallOpts, arg0)
}

// NTFLastMineTime is a free data retrieval call binding the contract method 0x13f34fed.
//
// Solidity: function NTFLastMineTime(address ) view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerCallerSession) NTFLastMineTime(arg0 common.Address) (*big.Int, error) {
	return _FishcakeEventManager.Contract.NTFLastMineTime(&_FishcakeEventManager.CallOpts, arg0)
}

// UsdtTokenAddr is a free data retrieval call binding the contract method 0xb6d18c8c.
//
// Solidity: function UsdtTokenAddr() view returns(address)
func (_FishcakeEventManager *FishcakeEventManagerCaller) UsdtTokenAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FishcakeEventManager.contract.Call(opts, &out, "UsdtTokenAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdtTokenAddr is a free data retrieval call binding the contract method 0xb6d18c8c.
//
// Solidity: function UsdtTokenAddr() view returns(address)
func (_FishcakeEventManager *FishcakeEventManagerSession) UsdtTokenAddr() (common.Address, error) {
	return _FishcakeEventManager.Contract.UsdtTokenAddr(&_FishcakeEventManager.CallOpts)
}

// UsdtTokenAddr is a free data retrieval call binding the contract method 0xb6d18c8c.
//
// Solidity: function UsdtTokenAddr() view returns(address)
func (_FishcakeEventManager *FishcakeEventManagerCallerSession) UsdtTokenAddr() (common.Address, error) {
	return _FishcakeEventManager.Contract.UsdtTokenAddr(&_FishcakeEventManager.CallOpts)
}

// ActivityDroppedToAccount is a free data retrieval call binding the contract method 0xacae5d25.
//
// Solidity: function activityDroppedToAccount(uint256 , address ) view returns(bool)
func (_FishcakeEventManager *FishcakeEventManagerCaller) ActivityDroppedToAccount(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _FishcakeEventManager.contract.Call(opts, &out, "activityDroppedToAccount", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ActivityDroppedToAccount is a free data retrieval call binding the contract method 0xacae5d25.
//
// Solidity: function activityDroppedToAccount(uint256 , address ) view returns(bool)
func (_FishcakeEventManager *FishcakeEventManagerSession) ActivityDroppedToAccount(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _FishcakeEventManager.Contract.ActivityDroppedToAccount(&_FishcakeEventManager.CallOpts, arg0, arg1)
}

// ActivityDroppedToAccount is a free data retrieval call binding the contract method 0xacae5d25.
//
// Solidity: function activityDroppedToAccount(uint256 , address ) view returns(bool)
func (_FishcakeEventManager *FishcakeEventManagerCallerSession) ActivityDroppedToAccount(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _FishcakeEventManager.Contract.ActivityDroppedToAccount(&_FishcakeEventManager.CallOpts, arg0, arg1)
}

// ActivityInfoArrs is a free data retrieval call binding the contract method 0xa0752825.
//
// Solidity: function activityInfoArrs(uint256 ) view returns(uint256 activityId, address businessAccount, string businessName, string activityContent, string latitudeLongitude, uint256 activityCreateTime, uint256 activityDeadLine, uint8 dropType, uint256 dropNumber, uint256 minDropAmt, uint256 maxDropAmt, address tokenContractAddr)
func (_FishcakeEventManager *FishcakeEventManagerCaller) ActivityInfoArrs(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _FishcakeEventManager.contract.Call(opts, &out, "activityInfoArrs", arg0)

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
func (_FishcakeEventManager *FishcakeEventManagerSession) ActivityInfoArrs(arg0 *big.Int) (struct {
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
	return _FishcakeEventManager.Contract.ActivityInfoArrs(&_FishcakeEventManager.CallOpts, arg0)
}

// ActivityInfoArrs is a free data retrieval call binding the contract method 0xa0752825.
//
// Solidity: function activityInfoArrs(uint256 ) view returns(uint256 activityId, address businessAccount, string businessName, string activityContent, string latitudeLongitude, uint256 activityCreateTime, uint256 activityDeadLine, uint8 dropType, uint256 dropNumber, uint256 minDropAmt, uint256 maxDropAmt, address tokenContractAddr)
func (_FishcakeEventManager *FishcakeEventManagerCallerSession) ActivityInfoArrs(arg0 *big.Int) (struct {
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
	return _FishcakeEventManager.Contract.ActivityInfoArrs(&_FishcakeEventManager.CallOpts, arg0)
}

// ActivityInfoChangedIdx is a free data retrieval call binding the contract method 0xfefee0e1.
//
// Solidity: function activityInfoChangedIdx(uint256 ) view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerCaller) ActivityInfoChangedIdx(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FishcakeEventManager.contract.Call(opts, &out, "activityInfoChangedIdx", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActivityInfoChangedIdx is a free data retrieval call binding the contract method 0xfefee0e1.
//
// Solidity: function activityInfoChangedIdx(uint256 ) view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerSession) ActivityInfoChangedIdx(arg0 *big.Int) (*big.Int, error) {
	return _FishcakeEventManager.Contract.ActivityInfoChangedIdx(&_FishcakeEventManager.CallOpts, arg0)
}

// ActivityInfoChangedIdx is a free data retrieval call binding the contract method 0xfefee0e1.
//
// Solidity: function activityInfoChangedIdx(uint256 ) view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerCallerSession) ActivityInfoChangedIdx(arg0 *big.Int) (*big.Int, error) {
	return _FishcakeEventManager.Contract.ActivityInfoChangedIdx(&_FishcakeEventManager.CallOpts, arg0)
}

// ActivityInfoExtArrs is a free data retrieval call binding the contract method 0x4d43959d.
//
// Solidity: function activityInfoExtArrs(uint256 ) view returns(uint256 activityId, uint256 alreadyDropAmts, uint256 alreadyDropNumber, uint256 businessMinedAmt, uint256 businessMinedWithdrawedAmt, uint8 activityStatus)
func (_FishcakeEventManager *FishcakeEventManagerCaller) ActivityInfoExtArrs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ActivityId                 *big.Int
	AlreadyDropAmts            *big.Int
	AlreadyDropNumber          *big.Int
	BusinessMinedAmt           *big.Int
	BusinessMinedWithdrawedAmt *big.Int
	ActivityStatus             uint8
}, error) {
	var out []interface{}
	err := _FishcakeEventManager.contract.Call(opts, &out, "activityInfoExtArrs", arg0)

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
func (_FishcakeEventManager *FishcakeEventManagerSession) ActivityInfoExtArrs(arg0 *big.Int) (struct {
	ActivityId                 *big.Int
	AlreadyDropAmts            *big.Int
	AlreadyDropNumber          *big.Int
	BusinessMinedAmt           *big.Int
	BusinessMinedWithdrawedAmt *big.Int
	ActivityStatus             uint8
}, error) {
	return _FishcakeEventManager.Contract.ActivityInfoExtArrs(&_FishcakeEventManager.CallOpts, arg0)
}

// ActivityInfoExtArrs is a free data retrieval call binding the contract method 0x4d43959d.
//
// Solidity: function activityInfoExtArrs(uint256 ) view returns(uint256 activityId, uint256 alreadyDropAmts, uint256 alreadyDropNumber, uint256 businessMinedAmt, uint256 businessMinedWithdrawedAmt, uint8 activityStatus)
func (_FishcakeEventManager *FishcakeEventManagerCallerSession) ActivityInfoExtArrs(arg0 *big.Int) (struct {
	ActivityId                 *big.Int
	AlreadyDropAmts            *big.Int
	AlreadyDropNumber          *big.Int
	BusinessMinedAmt           *big.Int
	BusinessMinedWithdrawedAmt *big.Int
	ActivityStatus             uint8
}, error) {
	return _FishcakeEventManager.Contract.ActivityInfoExtArrs(&_FishcakeEventManager.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FishcakeEventManager.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _FishcakeEventManager.Contract.Allowance(&_FishcakeEventManager.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _FishcakeEventManager.Contract.Allowance(&_FishcakeEventManager.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FishcakeEventManager.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _FishcakeEventManager.Contract.BalanceOf(&_FishcakeEventManager.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _FishcakeEventManager.Contract.BalanceOf(&_FishcakeEventManager.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FishcakeEventManager *FishcakeEventManagerCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _FishcakeEventManager.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FishcakeEventManager *FishcakeEventManagerSession) Decimals() (uint8, error) {
	return _FishcakeEventManager.Contract.Decimals(&_FishcakeEventManager.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FishcakeEventManager *FishcakeEventManagerCallerSession) Decimals() (uint8, error) {
	return _FishcakeEventManager.Contract.Decimals(&_FishcakeEventManager.CallOpts)
}

// DropInfoArrs is a free data retrieval call binding the contract method 0x2a916c20.
//
// Solidity: function dropInfoArrs(uint256 ) view returns(uint256 activityId, address userAccount, uint256 dropTime, uint256 dropAmt)
func (_FishcakeEventManager *FishcakeEventManagerCaller) DropInfoArrs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ActivityId  *big.Int
	UserAccount common.Address
	DropTime    *big.Int
	DropAmt     *big.Int
}, error) {
	var out []interface{}
	err := _FishcakeEventManager.contract.Call(opts, &out, "dropInfoArrs", arg0)

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
func (_FishcakeEventManager *FishcakeEventManagerSession) DropInfoArrs(arg0 *big.Int) (struct {
	ActivityId  *big.Int
	UserAccount common.Address
	DropTime    *big.Int
	DropAmt     *big.Int
}, error) {
	return _FishcakeEventManager.Contract.DropInfoArrs(&_FishcakeEventManager.CallOpts, arg0)
}

// DropInfoArrs is a free data retrieval call binding the contract method 0x2a916c20.
//
// Solidity: function dropInfoArrs(uint256 ) view returns(uint256 activityId, address userAccount, uint256 dropTime, uint256 dropAmt)
func (_FishcakeEventManager *FishcakeEventManagerCallerSession) DropInfoArrs(arg0 *big.Int) (struct {
	ActivityId  *big.Int
	UserAccount common.Address
	DropTime    *big.Int
	DropAmt     *big.Int
}, error) {
	return _FishcakeEventManager.Contract.DropInfoArrs(&_FishcakeEventManager.CallOpts, arg0)
}

// GetMinedFishcakePower is a free data retrieval call binding the contract method 0x4c6681e1.
//
// Solidity: function getMinedFishcakePower(address _miner) view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerCaller) GetMinedFishcakePower(opts *bind.CallOpts, _miner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FishcakeEventManager.contract.Call(opts, &out, "getMinedFishcakePower", _miner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinedFishcakePower is a free data retrieval call binding the contract method 0x4c6681e1.
//
// Solidity: function getMinedFishcakePower(address _miner) view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerSession) GetMinedFishcakePower(_miner common.Address) (*big.Int, error) {
	return _FishcakeEventManager.Contract.GetMinedFishcakePower(&_FishcakeEventManager.CallOpts, _miner)
}

// GetMinedFishcakePower is a free data retrieval call binding the contract method 0x4c6681e1.
//
// Solidity: function getMinedFishcakePower(address _miner) view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerCallerSession) GetMinedFishcakePower(_miner common.Address) (*big.Int, error) {
	return _FishcakeEventManager.Contract.GetMinedFishcakePower(&_FishcakeEventManager.CallOpts, _miner)
}

// GetMinerMineAmount is a free data retrieval call binding the contract method 0xa731299b.
//
// Solidity: function getMinerMineAmount(address _miner) view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerCaller) GetMinerMineAmount(opts *bind.CallOpts, _miner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FishcakeEventManager.contract.Call(opts, &out, "getMinerMineAmount", _miner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinerMineAmount is a free data retrieval call binding the contract method 0xa731299b.
//
// Solidity: function getMinerMineAmount(address _miner) view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerSession) GetMinerMineAmount(_miner common.Address) (*big.Int, error) {
	return _FishcakeEventManager.Contract.GetMinerMineAmount(&_FishcakeEventManager.CallOpts, _miner)
}

// GetMinerMineAmount is a free data retrieval call binding the contract method 0xa731299b.
//
// Solidity: function getMinerMineAmount(address _miner) view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerCallerSession) GetMinerMineAmount(_miner common.Address) (*big.Int, error) {
	return _FishcakeEventManager.Contract.GetMinerMineAmount(&_FishcakeEventManager.CallOpts, _miner)
}

// INFTManager is a free data retrieval call binding the contract method 0xbbe68511.
//
// Solidity: function iNFTManager() view returns(address)
func (_FishcakeEventManager *FishcakeEventManagerCaller) INFTManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FishcakeEventManager.contract.Call(opts, &out, "iNFTManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// INFTManager is a free data retrieval call binding the contract method 0xbbe68511.
//
// Solidity: function iNFTManager() view returns(address)
func (_FishcakeEventManager *FishcakeEventManagerSession) INFTManager() (common.Address, error) {
	return _FishcakeEventManager.Contract.INFTManager(&_FishcakeEventManager.CallOpts)
}

// INFTManager is a free data retrieval call binding the contract method 0xbbe68511.
//
// Solidity: function iNFTManager() view returns(address)
func (_FishcakeEventManager *FishcakeEventManagerCallerSession) INFTManager() (common.Address, error) {
	return _FishcakeEventManager.Contract.INFTManager(&_FishcakeEventManager.CallOpts)
}

// IsMint is a free data retrieval call binding the contract method 0x33b3f944.
//
// Solidity: function isMint() view returns(bool)
func (_FishcakeEventManager *FishcakeEventManagerCaller) IsMint(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FishcakeEventManager.contract.Call(opts, &out, "isMint")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMint is a free data retrieval call binding the contract method 0x33b3f944.
//
// Solidity: function isMint() view returns(bool)
func (_FishcakeEventManager *FishcakeEventManagerSession) IsMint() (bool, error) {
	return _FishcakeEventManager.Contract.IsMint(&_FishcakeEventManager.CallOpts)
}

// IsMint is a free data retrieval call binding the contract method 0x33b3f944.
//
// Solidity: function isMint() view returns(bool)
func (_FishcakeEventManager *FishcakeEventManagerCallerSession) IsMint() (bool, error) {
	return _FishcakeEventManager.Contract.IsMint(&_FishcakeEventManager.CallOpts)
}

// MaxDeadLine is a free data retrieval call binding the contract method 0x4dc143d5.
//
// Solidity: function maxDeadLine() view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerCaller) MaxDeadLine(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FishcakeEventManager.contract.Call(opts, &out, "maxDeadLine")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeadLine is a free data retrieval call binding the contract method 0x4dc143d5.
//
// Solidity: function maxDeadLine() view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerSession) MaxDeadLine() (*big.Int, error) {
	return _FishcakeEventManager.Contract.MaxDeadLine(&_FishcakeEventManager.CallOpts)
}

// MaxDeadLine is a free data retrieval call binding the contract method 0x4dc143d5.
//
// Solidity: function maxDeadLine() view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerCallerSession) MaxDeadLine() (*big.Int, error) {
	return _FishcakeEventManager.Contract.MaxDeadLine(&_FishcakeEventManager.CallOpts)
}

// MerchantOnceMaxMineAmt is a free data retrieval call binding the contract method 0xcdd544e3.
//
// Solidity: function merchantOnceMaxMineAmt() view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerCaller) MerchantOnceMaxMineAmt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FishcakeEventManager.contract.Call(opts, &out, "merchantOnceMaxMineAmt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MerchantOnceMaxMineAmt is a free data retrieval call binding the contract method 0xcdd544e3.
//
// Solidity: function merchantOnceMaxMineAmt() view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerSession) MerchantOnceMaxMineAmt() (*big.Int, error) {
	return _FishcakeEventManager.Contract.MerchantOnceMaxMineAmt(&_FishcakeEventManager.CallOpts)
}

// MerchantOnceMaxMineAmt is a free data retrieval call binding the contract method 0xcdd544e3.
//
// Solidity: function merchantOnceMaxMineAmt() view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerCallerSession) MerchantOnceMaxMineAmt() (*big.Int, error) {
	return _FishcakeEventManager.Contract.MerchantOnceMaxMineAmt(&_FishcakeEventManager.CallOpts)
}

// MinePercent is a free data retrieval call binding the contract method 0xa6d93915.
//
// Solidity: function minePercent() view returns(uint8)
func (_FishcakeEventManager *FishcakeEventManagerCaller) MinePercent(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _FishcakeEventManager.contract.Call(opts, &out, "minePercent")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MinePercent is a free data retrieval call binding the contract method 0xa6d93915.
//
// Solidity: function minePercent() view returns(uint8)
func (_FishcakeEventManager *FishcakeEventManagerSession) MinePercent() (uint8, error) {
	return _FishcakeEventManager.Contract.MinePercent(&_FishcakeEventManager.CallOpts)
}

// MinePercent is a free data retrieval call binding the contract method 0xa6d93915.
//
// Solidity: function minePercent() view returns(uint8)
func (_FishcakeEventManager *FishcakeEventManagerCallerSession) MinePercent() (uint8, error) {
	return _FishcakeEventManager.Contract.MinePercent(&_FishcakeEventManager.CallOpts)
}

// MinedAmt is a free data retrieval call binding the contract method 0xf7859c28.
//
// Solidity: function minedAmt() view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerCaller) MinedAmt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FishcakeEventManager.contract.Call(opts, &out, "minedAmt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinedAmt is a free data retrieval call binding the contract method 0xf7859c28.
//
// Solidity: function minedAmt() view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerSession) MinedAmt() (*big.Int, error) {
	return _FishcakeEventManager.Contract.MinedAmt(&_FishcakeEventManager.CallOpts)
}

// MinedAmt is a free data retrieval call binding the contract method 0xf7859c28.
//
// Solidity: function minedAmt() view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerCallerSession) MinedAmt() (*big.Int, error) {
	return _FishcakeEventManager.Contract.MinedAmt(&_FishcakeEventManager.CallOpts)
}

// MinedFishcakePower is a free data retrieval call binding the contract method 0xd4621430.
//
// Solidity: function minedFishcakePower(address ) view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerCaller) MinedFishcakePower(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FishcakeEventManager.contract.Call(opts, &out, "minedFishcakePower", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinedFishcakePower is a free data retrieval call binding the contract method 0xd4621430.
//
// Solidity: function minedFishcakePower(address ) view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerSession) MinedFishcakePower(arg0 common.Address) (*big.Int, error) {
	return _FishcakeEventManager.Contract.MinedFishcakePower(&_FishcakeEventManager.CallOpts, arg0)
}

// MinedFishcakePower is a free data retrieval call binding the contract method 0xd4621430.
//
// Solidity: function minedFishcakePower(address ) view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerCallerSession) MinedFishcakePower(arg0 common.Address) (*big.Int, error) {
	return _FishcakeEventManager.Contract.MinedFishcakePower(&_FishcakeEventManager.CallOpts, arg0)
}

// MinerMineAmount is a free data retrieval call binding the contract method 0x265d3051.
//
// Solidity: function minerMineAmount(address ) view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerCaller) MinerMineAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FishcakeEventManager.contract.Call(opts, &out, "minerMineAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinerMineAmount is a free data retrieval call binding the contract method 0x265d3051.
//
// Solidity: function minerMineAmount(address ) view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerSession) MinerMineAmount(arg0 common.Address) (*big.Int, error) {
	return _FishcakeEventManager.Contract.MinerMineAmount(&_FishcakeEventManager.CallOpts, arg0)
}

// MinerMineAmount is a free data retrieval call binding the contract method 0x265d3051.
//
// Solidity: function minerMineAmount(address ) view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerCallerSession) MinerMineAmount(arg0 common.Address) (*big.Int, error) {
	return _FishcakeEventManager.Contract.MinerMineAmount(&_FishcakeEventManager.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FishcakeEventManager *FishcakeEventManagerCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FishcakeEventManager.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FishcakeEventManager *FishcakeEventManagerSession) Name() (string, error) {
	return _FishcakeEventManager.Contract.Name(&_FishcakeEventManager.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FishcakeEventManager *FishcakeEventManagerCallerSession) Name() (string, error) {
	return _FishcakeEventManager.Contract.Name(&_FishcakeEventManager.CallOpts)
}

// OneDay is a free data retrieval call binding the contract method 0x4bd9351b.
//
// Solidity: function oneDay() view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerCaller) OneDay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FishcakeEventManager.contract.Call(opts, &out, "oneDay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OneDay is a free data retrieval call binding the contract method 0x4bd9351b.
//
// Solidity: function oneDay() view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerSession) OneDay() (*big.Int, error) {
	return _FishcakeEventManager.Contract.OneDay(&_FishcakeEventManager.CallOpts)
}

// OneDay is a free data retrieval call binding the contract method 0x4bd9351b.
//
// Solidity: function oneDay() view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerCallerSession) OneDay() (*big.Int, error) {
	return _FishcakeEventManager.Contract.OneDay(&_FishcakeEventManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FishcakeEventManager *FishcakeEventManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FishcakeEventManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FishcakeEventManager *FishcakeEventManagerSession) Owner() (common.Address, error) {
	return _FishcakeEventManager.Contract.Owner(&_FishcakeEventManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FishcakeEventManager *FishcakeEventManagerCallerSession) Owner() (common.Address, error) {
	return _FishcakeEventManager.Contract.Owner(&_FishcakeEventManager.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FishcakeEventManager *FishcakeEventManagerCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FishcakeEventManager.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FishcakeEventManager *FishcakeEventManagerSession) Symbol() (string, error) {
	return _FishcakeEventManager.Contract.Symbol(&_FishcakeEventManager.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FishcakeEventManager *FishcakeEventManagerCallerSession) Symbol() (string, error) {
	return _FishcakeEventManager.Contract.Symbol(&_FishcakeEventManager.CallOpts)
}

// TotalMineAmt is a free data retrieval call binding the contract method 0x88e40a5a.
//
// Solidity: function totalMineAmt() view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerCaller) TotalMineAmt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FishcakeEventManager.contract.Call(opts, &out, "totalMineAmt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMineAmt is a free data retrieval call binding the contract method 0x88e40a5a.
//
// Solidity: function totalMineAmt() view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerSession) TotalMineAmt() (*big.Int, error) {
	return _FishcakeEventManager.Contract.TotalMineAmt(&_FishcakeEventManager.CallOpts)
}

// TotalMineAmt is a free data retrieval call binding the contract method 0x88e40a5a.
//
// Solidity: function totalMineAmt() view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerCallerSession) TotalMineAmt() (*big.Int, error) {
	return _FishcakeEventManager.Contract.TotalMineAmt(&_FishcakeEventManager.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FishcakeEventManager.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerSession) TotalSupply() (*big.Int, error) {
	return _FishcakeEventManager.Contract.TotalSupply(&_FishcakeEventManager.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerCallerSession) TotalSupply() (*big.Int, error) {
	return _FishcakeEventManager.Contract.TotalSupply(&_FishcakeEventManager.CallOpts)
}

// UserOnceMaxMineAmt is a free data retrieval call binding the contract method 0x7f17bb68.
//
// Solidity: function userOnceMaxMineAmt() view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerCaller) UserOnceMaxMineAmt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FishcakeEventManager.contract.Call(opts, &out, "userOnceMaxMineAmt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserOnceMaxMineAmt is a free data retrieval call binding the contract method 0x7f17bb68.
//
// Solidity: function userOnceMaxMineAmt() view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerSession) UserOnceMaxMineAmt() (*big.Int, error) {
	return _FishcakeEventManager.Contract.UserOnceMaxMineAmt(&_FishcakeEventManager.CallOpts)
}

// UserOnceMaxMineAmt is a free data retrieval call binding the contract method 0x7f17bb68.
//
// Solidity: function userOnceMaxMineAmt() view returns(uint256)
func (_FishcakeEventManager *FishcakeEventManagerCallerSession) UserOnceMaxMineAmt() (*big.Int, error) {
	return _FishcakeEventManager.Contract.UserOnceMaxMineAmt(&_FishcakeEventManager.CallOpts)
}

// ActivityAdd is a paid mutator transaction binding the contract method 0x43b22efc.
//
// Solidity: function activityAdd(string _businessName, string _activityContent, string _latitudeLongitude, uint256 _activityDeadLine, uint256 _totalDropAmts, uint8 _dropType, uint256 _dropNumber, uint256 _minDropAmt, uint256 _maxDropAmt, address _tokenContractAddr) returns(bool, uint256)
func (_FishcakeEventManager *FishcakeEventManagerTransactor) ActivityAdd(opts *bind.TransactOpts, _businessName string, _activityContent string, _latitudeLongitude string, _activityDeadLine *big.Int, _totalDropAmts *big.Int, _dropType uint8, _dropNumber *big.Int, _minDropAmt *big.Int, _maxDropAmt *big.Int, _tokenContractAddr common.Address) (*types.Transaction, error) {
	return _FishcakeEventManager.contract.Transact(opts, "activityAdd", _businessName, _activityContent, _latitudeLongitude, _activityDeadLine, _totalDropAmts, _dropType, _dropNumber, _minDropAmt, _maxDropAmt, _tokenContractAddr)
}

// ActivityAdd is a paid mutator transaction binding the contract method 0x43b22efc.
//
// Solidity: function activityAdd(string _businessName, string _activityContent, string _latitudeLongitude, uint256 _activityDeadLine, uint256 _totalDropAmts, uint8 _dropType, uint256 _dropNumber, uint256 _minDropAmt, uint256 _maxDropAmt, address _tokenContractAddr) returns(bool, uint256)
func (_FishcakeEventManager *FishcakeEventManagerSession) ActivityAdd(_businessName string, _activityContent string, _latitudeLongitude string, _activityDeadLine *big.Int, _totalDropAmts *big.Int, _dropType uint8, _dropNumber *big.Int, _minDropAmt *big.Int, _maxDropAmt *big.Int, _tokenContractAddr common.Address) (*types.Transaction, error) {
	return _FishcakeEventManager.Contract.ActivityAdd(&_FishcakeEventManager.TransactOpts, _businessName, _activityContent, _latitudeLongitude, _activityDeadLine, _totalDropAmts, _dropType, _dropNumber, _minDropAmt, _maxDropAmt, _tokenContractAddr)
}

// ActivityAdd is a paid mutator transaction binding the contract method 0x43b22efc.
//
// Solidity: function activityAdd(string _businessName, string _activityContent, string _latitudeLongitude, uint256 _activityDeadLine, uint256 _totalDropAmts, uint8 _dropType, uint256 _dropNumber, uint256 _minDropAmt, uint256 _maxDropAmt, address _tokenContractAddr) returns(bool, uint256)
func (_FishcakeEventManager *FishcakeEventManagerTransactorSession) ActivityAdd(_businessName string, _activityContent string, _latitudeLongitude string, _activityDeadLine *big.Int, _totalDropAmts *big.Int, _dropType uint8, _dropNumber *big.Int, _minDropAmt *big.Int, _maxDropAmt *big.Int, _tokenContractAddr common.Address) (*types.Transaction, error) {
	return _FishcakeEventManager.Contract.ActivityAdd(&_FishcakeEventManager.TransactOpts, _businessName, _activityContent, _latitudeLongitude, _activityDeadLine, _totalDropAmts, _dropType, _dropNumber, _minDropAmt, _maxDropAmt, _tokenContractAddr)
}

// ActivityFinish is a paid mutator transaction binding the contract method 0xcbd596c4.
//
// Solidity: function activityFinish(uint256 _activityId) returns(bool)
func (_FishcakeEventManager *FishcakeEventManagerTransactor) ActivityFinish(opts *bind.TransactOpts, _activityId *big.Int) (*types.Transaction, error) {
	return _FishcakeEventManager.contract.Transact(opts, "activityFinish", _activityId)
}

// ActivityFinish is a paid mutator transaction binding the contract method 0xcbd596c4.
//
// Solidity: function activityFinish(uint256 _activityId) returns(bool)
func (_FishcakeEventManager *FishcakeEventManagerSession) ActivityFinish(_activityId *big.Int) (*types.Transaction, error) {
	return _FishcakeEventManager.Contract.ActivityFinish(&_FishcakeEventManager.TransactOpts, _activityId)
}

// ActivityFinish is a paid mutator transaction binding the contract method 0xcbd596c4.
//
// Solidity: function activityFinish(uint256 _activityId) returns(bool)
func (_FishcakeEventManager *FishcakeEventManagerTransactorSession) ActivityFinish(_activityId *big.Int) (*types.Transaction, error) {
	return _FishcakeEventManager.Contract.ActivityFinish(&_FishcakeEventManager.TransactOpts, _activityId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_FishcakeEventManager *FishcakeEventManagerTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _FishcakeEventManager.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_FishcakeEventManager *FishcakeEventManagerSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _FishcakeEventManager.Contract.Approve(&_FishcakeEventManager.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_FishcakeEventManager *FishcakeEventManagerTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _FishcakeEventManager.Contract.Approve(&_FishcakeEventManager.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_FishcakeEventManager *FishcakeEventManagerTransactor) Burn(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _FishcakeEventManager.contract.Transact(opts, "burn", value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_FishcakeEventManager *FishcakeEventManagerSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _FishcakeEventManager.Contract.Burn(&_FishcakeEventManager.TransactOpts, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_FishcakeEventManager *FishcakeEventManagerTransactorSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _FishcakeEventManager.Contract.Burn(&_FishcakeEventManager.TransactOpts, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_FishcakeEventManager *FishcakeEventManagerTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, value *big.Int) (*types.Transaction, error) {
	return _FishcakeEventManager.contract.Transact(opts, "burnFrom", account, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_FishcakeEventManager *FishcakeEventManagerSession) BurnFrom(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _FishcakeEventManager.Contract.BurnFrom(&_FishcakeEventManager.TransactOpts, account, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_FishcakeEventManager *FishcakeEventManagerTransactorSession) BurnFrom(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _FishcakeEventManager.Contract.BurnFrom(&_FishcakeEventManager.TransactOpts, account, value)
}

// DeleteMinedFishcakePower is a paid mutator transaction binding the contract method 0xf74b4709.
//
// Solidity: function deleteMinedFishcakePower(address _miner) returns()
func (_FishcakeEventManager *FishcakeEventManagerTransactor) DeleteMinedFishcakePower(opts *bind.TransactOpts, _miner common.Address) (*types.Transaction, error) {
	return _FishcakeEventManager.contract.Transact(opts, "deleteMinedFishcakePower", _miner)
}

// DeleteMinedFishcakePower is a paid mutator transaction binding the contract method 0xf74b4709.
//
// Solidity: function deleteMinedFishcakePower(address _miner) returns()
func (_FishcakeEventManager *FishcakeEventManagerSession) DeleteMinedFishcakePower(_miner common.Address) (*types.Transaction, error) {
	return _FishcakeEventManager.Contract.DeleteMinedFishcakePower(&_FishcakeEventManager.TransactOpts, _miner)
}

// DeleteMinedFishcakePower is a paid mutator transaction binding the contract method 0xf74b4709.
//
// Solidity: function deleteMinedFishcakePower(address _miner) returns()
func (_FishcakeEventManager *FishcakeEventManagerTransactorSession) DeleteMinedFishcakePower(_miner common.Address) (*types.Transaction, error) {
	return _FishcakeEventManager.Contract.DeleteMinedFishcakePower(&_FishcakeEventManager.TransactOpts, _miner)
}

// DeleteMinerMineAmount is a paid mutator transaction binding the contract method 0x97088e49.
//
// Solidity: function deleteMinerMineAmount(address _miner) returns()
func (_FishcakeEventManager *FishcakeEventManagerTransactor) DeleteMinerMineAmount(opts *bind.TransactOpts, _miner common.Address) (*types.Transaction, error) {
	return _FishcakeEventManager.contract.Transact(opts, "deleteMinerMineAmount", _miner)
}

// DeleteMinerMineAmount is a paid mutator transaction binding the contract method 0x97088e49.
//
// Solidity: function deleteMinerMineAmount(address _miner) returns()
func (_FishcakeEventManager *FishcakeEventManagerSession) DeleteMinerMineAmount(_miner common.Address) (*types.Transaction, error) {
	return _FishcakeEventManager.Contract.DeleteMinerMineAmount(&_FishcakeEventManager.TransactOpts, _miner)
}

// DeleteMinerMineAmount is a paid mutator transaction binding the contract method 0x97088e49.
//
// Solidity: function deleteMinerMineAmount(address _miner) returns()
func (_FishcakeEventManager *FishcakeEventManagerTransactorSession) DeleteMinerMineAmount(_miner common.Address) (*types.Transaction, error) {
	return _FishcakeEventManager.Contract.DeleteMinerMineAmount(&_FishcakeEventManager.TransactOpts, _miner)
}

// Drop is a paid mutator transaction binding the contract method 0xc8c1d78d.
//
// Solidity: function drop(uint256 _activityId, address _userAccount, uint256 _dropAmt) returns(bool)
func (_FishcakeEventManager *FishcakeEventManagerTransactor) Drop(opts *bind.TransactOpts, _activityId *big.Int, _userAccount common.Address, _dropAmt *big.Int) (*types.Transaction, error) {
	return _FishcakeEventManager.contract.Transact(opts, "drop", _activityId, _userAccount, _dropAmt)
}

// Drop is a paid mutator transaction binding the contract method 0xc8c1d78d.
//
// Solidity: function drop(uint256 _activityId, address _userAccount, uint256 _dropAmt) returns(bool)
func (_FishcakeEventManager *FishcakeEventManagerSession) Drop(_activityId *big.Int, _userAccount common.Address, _dropAmt *big.Int) (*types.Transaction, error) {
	return _FishcakeEventManager.Contract.Drop(&_FishcakeEventManager.TransactOpts, _activityId, _userAccount, _dropAmt)
}

// Drop is a paid mutator transaction binding the contract method 0xc8c1d78d.
//
// Solidity: function drop(uint256 _activityId, address _userAccount, uint256 _dropAmt) returns(bool)
func (_FishcakeEventManager *FishcakeEventManagerTransactorSession) Drop(_activityId *big.Int, _userAccount common.Address, _dropAmt *big.Int) (*types.Transaction, error) {
	return _FishcakeEventManager.Contract.Drop(&_FishcakeEventManager.TransactOpts, _activityId, _userAccount, _dropAmt)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _initialOwner, address _fccAddress, address _usdtTokenAddr, address _NFTManagerAddr) returns()
func (_FishcakeEventManager *FishcakeEventManagerTransactor) Initialize(opts *bind.TransactOpts, _initialOwner common.Address, _fccAddress common.Address, _usdtTokenAddr common.Address, _NFTManagerAddr common.Address) (*types.Transaction, error) {
	return _FishcakeEventManager.contract.Transact(opts, "initialize", _initialOwner, _fccAddress, _usdtTokenAddr, _NFTManagerAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _initialOwner, address _fccAddress, address _usdtTokenAddr, address _NFTManagerAddr) returns()
func (_FishcakeEventManager *FishcakeEventManagerSession) Initialize(_initialOwner common.Address, _fccAddress common.Address, _usdtTokenAddr common.Address, _NFTManagerAddr common.Address) (*types.Transaction, error) {
	return _FishcakeEventManager.Contract.Initialize(&_FishcakeEventManager.TransactOpts, _initialOwner, _fccAddress, _usdtTokenAddr, _NFTManagerAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _initialOwner, address _fccAddress, address _usdtTokenAddr, address _NFTManagerAddr) returns()
func (_FishcakeEventManager *FishcakeEventManagerTransactorSession) Initialize(_initialOwner common.Address, _fccAddress common.Address, _usdtTokenAddr common.Address, _NFTManagerAddr common.Address) (*types.Transaction, error) {
	return _FishcakeEventManager.Contract.Initialize(&_FishcakeEventManager.TransactOpts, _initialOwner, _fccAddress, _usdtTokenAddr, _NFTManagerAddr)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FishcakeEventManager *FishcakeEventManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FishcakeEventManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FishcakeEventManager *FishcakeEventManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _FishcakeEventManager.Contract.RenounceOwnership(&_FishcakeEventManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FishcakeEventManager *FishcakeEventManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FishcakeEventManager.Contract.RenounceOwnership(&_FishcakeEventManager.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_FishcakeEventManager *FishcakeEventManagerTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _FishcakeEventManager.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_FishcakeEventManager *FishcakeEventManagerSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _FishcakeEventManager.Contract.Transfer(&_FishcakeEventManager.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_FishcakeEventManager *FishcakeEventManagerTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _FishcakeEventManager.Contract.Transfer(&_FishcakeEventManager.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_FishcakeEventManager *FishcakeEventManagerTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _FishcakeEventManager.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_FishcakeEventManager *FishcakeEventManagerSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _FishcakeEventManager.Contract.TransferFrom(&_FishcakeEventManager.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_FishcakeEventManager *FishcakeEventManagerTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _FishcakeEventManager.Contract.TransferFrom(&_FishcakeEventManager.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FishcakeEventManager *FishcakeEventManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FishcakeEventManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FishcakeEventManager *FishcakeEventManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FishcakeEventManager.Contract.TransferOwnership(&_FishcakeEventManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FishcakeEventManager *FishcakeEventManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FishcakeEventManager.Contract.TransferOwnership(&_FishcakeEventManager.TransactOpts, newOwner)
}

// FishcakeEventManagerActivityAddIterator is returned from FilterActivityAdd and is used to iterate over the raw logs and unpacked data for ActivityAdd events raised by the FishcakeEventManager contract.
type FishcakeEventManagerActivityAddIterator struct {
	Event *FishcakeEventManagerActivityAdd // Event containing the contract specifics and raw log

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
func (it *FishcakeEventManagerActivityAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FishcakeEventManagerActivityAdd)
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
		it.Event = new(FishcakeEventManagerActivityAdd)
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
func (it *FishcakeEventManagerActivityAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FishcakeEventManagerActivityAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FishcakeEventManagerActivityAdd represents a ActivityAdd event raised by the FishcakeEventManager contract.
type FishcakeEventManagerActivityAdd struct {
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
func (_FishcakeEventManager *FishcakeEventManagerFilterer) FilterActivityAdd(opts *bind.FilterOpts, who []common.Address, _activityId []*big.Int) (*FishcakeEventManagerActivityAddIterator, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}
	var _activityIdRule []interface{}
	for _, _activityIdItem := range _activityId {
		_activityIdRule = append(_activityIdRule, _activityIdItem)
	}

	logs, sub, err := _FishcakeEventManager.contract.FilterLogs(opts, "ActivityAdd", whoRule, _activityIdRule)
	if err != nil {
		return nil, err
	}
	return &FishcakeEventManagerActivityAddIterator{contract: _FishcakeEventManager.contract, event: "ActivityAdd", logs: logs, sub: sub}, nil
}

// WatchActivityAdd is a free log subscription operation binding the contract event 0xcc07277f6a9ae6659df5a850b317ba47d741264062fc7654d69b9fe11f4fabef.
//
// Solidity: event ActivityAdd(address indexed who, uint256 indexed _activityId, uint256 _totalDropAmts, string _businessName, string _activityContent, string _latitudeLongitude, uint256 _activityDeadLine, uint8 _dropType, uint256 _dropNumber, uint256 _minDropAmt, uint256 _maxDropAmt, address _tokenContractAddr)
func (_FishcakeEventManager *FishcakeEventManagerFilterer) WatchActivityAdd(opts *bind.WatchOpts, sink chan<- *FishcakeEventManagerActivityAdd, who []common.Address, _activityId []*big.Int) (event.Subscription, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}
	var _activityIdRule []interface{}
	for _, _activityIdItem := range _activityId {
		_activityIdRule = append(_activityIdRule, _activityIdItem)
	}

	logs, sub, err := _FishcakeEventManager.contract.WatchLogs(opts, "ActivityAdd", whoRule, _activityIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FishcakeEventManagerActivityAdd)
				if err := _FishcakeEventManager.contract.UnpackLog(event, "ActivityAdd", log); err != nil {
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
func (_FishcakeEventManager *FishcakeEventManagerFilterer) ParseActivityAdd(log types.Log) (*FishcakeEventManagerActivityAdd, error) {
	event := new(FishcakeEventManagerActivityAdd)
	if err := _FishcakeEventManager.contract.UnpackLog(event, "ActivityAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FishcakeEventManagerActivityFinishIterator is returned from FilterActivityFinish and is used to iterate over the raw logs and unpacked data for ActivityFinish events raised by the FishcakeEventManager contract.
type FishcakeEventManagerActivityFinishIterator struct {
	Event *FishcakeEventManagerActivityFinish // Event containing the contract specifics and raw log

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
func (it *FishcakeEventManagerActivityFinishIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FishcakeEventManagerActivityFinish)
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
		it.Event = new(FishcakeEventManagerActivityFinish)
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
func (it *FishcakeEventManagerActivityFinishIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FishcakeEventManagerActivityFinishIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FishcakeEventManagerActivityFinish represents a ActivityFinish event raised by the FishcakeEventManager contract.
type FishcakeEventManagerActivityFinish struct {
	ActivityId        *big.Int
	TokenContractAddr common.Address
	ReturnAmount      *big.Int
	MinedAmount       *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterActivityFinish is a free log retrieval operation binding the contract event 0x32b4fb78ab9b6fcf0181a2a3906e5473b095a3e71595ccc87bd039bbacc6acaf.
//
// Solidity: event ActivityFinish(uint256 indexed _activityId, address _tokenContractAddr, uint256 _returnAmount, uint256 _minedAmount)
func (_FishcakeEventManager *FishcakeEventManagerFilterer) FilterActivityFinish(opts *bind.FilterOpts, _activityId []*big.Int) (*FishcakeEventManagerActivityFinishIterator, error) {

	var _activityIdRule []interface{}
	for _, _activityIdItem := range _activityId {
		_activityIdRule = append(_activityIdRule, _activityIdItem)
	}

	logs, sub, err := _FishcakeEventManager.contract.FilterLogs(opts, "ActivityFinish", _activityIdRule)
	if err != nil {
		return nil, err
	}
	return &FishcakeEventManagerActivityFinishIterator{contract: _FishcakeEventManager.contract, event: "ActivityFinish", logs: logs, sub: sub}, nil
}

// WatchActivityFinish is a free log subscription operation binding the contract event 0x32b4fb78ab9b6fcf0181a2a3906e5473b095a3e71595ccc87bd039bbacc6acaf.
//
// Solidity: event ActivityFinish(uint256 indexed _activityId, address _tokenContractAddr, uint256 _returnAmount, uint256 _minedAmount)
func (_FishcakeEventManager *FishcakeEventManagerFilterer) WatchActivityFinish(opts *bind.WatchOpts, sink chan<- *FishcakeEventManagerActivityFinish, _activityId []*big.Int) (event.Subscription, error) {

	var _activityIdRule []interface{}
	for _, _activityIdItem := range _activityId {
		_activityIdRule = append(_activityIdRule, _activityIdItem)
	}

	logs, sub, err := _FishcakeEventManager.contract.WatchLogs(opts, "ActivityFinish", _activityIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FishcakeEventManagerActivityFinish)
				if err := _FishcakeEventManager.contract.UnpackLog(event, "ActivityFinish", log); err != nil {
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

// ParseActivityFinish is a log parse operation binding the contract event 0x32b4fb78ab9b6fcf0181a2a3906e5473b095a3e71595ccc87bd039bbacc6acaf.
//
// Solidity: event ActivityFinish(uint256 indexed _activityId, address _tokenContractAddr, uint256 _returnAmount, uint256 _minedAmount)
func (_FishcakeEventManager *FishcakeEventManagerFilterer) ParseActivityFinish(log types.Log) (*FishcakeEventManagerActivityFinish, error) {
	event := new(FishcakeEventManagerActivityFinish)
	if err := _FishcakeEventManager.contract.UnpackLog(event, "ActivityFinish", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FishcakeEventManagerAddMineAmtIterator is returned from FilterAddMineAmt and is used to iterate over the raw logs and unpacked data for AddMineAmt events raised by the FishcakeEventManager contract.
type FishcakeEventManagerAddMineAmtIterator struct {
	Event *FishcakeEventManagerAddMineAmt // Event containing the contract specifics and raw log

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
func (it *FishcakeEventManagerAddMineAmtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FishcakeEventManagerAddMineAmt)
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
		it.Event = new(FishcakeEventManagerAddMineAmt)
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
func (it *FishcakeEventManagerAddMineAmtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FishcakeEventManagerAddMineAmtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FishcakeEventManagerAddMineAmt represents a AddMineAmt event raised by the FishcakeEventManager contract.
type FishcakeEventManagerAddMineAmt struct {
	Who        common.Address
	AddMineAmt *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddMineAmt is a free log retrieval operation binding the contract event 0x36a36f6d06ab1b164d3fbb904d2e7657b13fc5444f9ac18bc9a7af5aa6868376.
//
// Solidity: event AddMineAmt(address indexed who, uint256 _addMineAmt)
func (_FishcakeEventManager *FishcakeEventManagerFilterer) FilterAddMineAmt(opts *bind.FilterOpts, who []common.Address) (*FishcakeEventManagerAddMineAmtIterator, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _FishcakeEventManager.contract.FilterLogs(opts, "AddMineAmt", whoRule)
	if err != nil {
		return nil, err
	}
	return &FishcakeEventManagerAddMineAmtIterator{contract: _FishcakeEventManager.contract, event: "AddMineAmt", logs: logs, sub: sub}, nil
}

// WatchAddMineAmt is a free log subscription operation binding the contract event 0x36a36f6d06ab1b164d3fbb904d2e7657b13fc5444f9ac18bc9a7af5aa6868376.
//
// Solidity: event AddMineAmt(address indexed who, uint256 _addMineAmt)
func (_FishcakeEventManager *FishcakeEventManagerFilterer) WatchAddMineAmt(opts *bind.WatchOpts, sink chan<- *FishcakeEventManagerAddMineAmt, who []common.Address) (event.Subscription, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _FishcakeEventManager.contract.WatchLogs(opts, "AddMineAmt", whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FishcakeEventManagerAddMineAmt)
				if err := _FishcakeEventManager.contract.UnpackLog(event, "AddMineAmt", log); err != nil {
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
func (_FishcakeEventManager *FishcakeEventManagerFilterer) ParseAddMineAmt(log types.Log) (*FishcakeEventManagerAddMineAmt, error) {
	event := new(FishcakeEventManagerAddMineAmt)
	if err := _FishcakeEventManager.contract.UnpackLog(event, "AddMineAmt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FishcakeEventManagerApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the FishcakeEventManager contract.
type FishcakeEventManagerApprovalIterator struct {
	Event *FishcakeEventManagerApproval // Event containing the contract specifics and raw log

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
func (it *FishcakeEventManagerApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FishcakeEventManagerApproval)
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
		it.Event = new(FishcakeEventManagerApproval)
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
func (it *FishcakeEventManagerApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FishcakeEventManagerApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FishcakeEventManagerApproval represents a Approval event raised by the FishcakeEventManager contract.
type FishcakeEventManagerApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_FishcakeEventManager *FishcakeEventManagerFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*FishcakeEventManagerApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FishcakeEventManager.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &FishcakeEventManagerApprovalIterator{contract: _FishcakeEventManager.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_FishcakeEventManager *FishcakeEventManagerFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FishcakeEventManagerApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FishcakeEventManager.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FishcakeEventManagerApproval)
				if err := _FishcakeEventManager.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_FishcakeEventManager *FishcakeEventManagerFilterer) ParseApproval(log types.Log) (*FishcakeEventManagerApproval, error) {
	event := new(FishcakeEventManagerApproval)
	if err := _FishcakeEventManager.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FishcakeEventManagerDropIterator is returned from FilterDrop and is used to iterate over the raw logs and unpacked data for Drop events raised by the FishcakeEventManager contract.
type FishcakeEventManagerDropIterator struct {
	Event *FishcakeEventManagerDrop // Event containing the contract specifics and raw log

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
func (it *FishcakeEventManagerDropIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FishcakeEventManagerDrop)
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
		it.Event = new(FishcakeEventManagerDrop)
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
func (it *FishcakeEventManagerDropIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FishcakeEventManagerDropIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FishcakeEventManagerDrop represents a Drop event raised by the FishcakeEventManager contract.
type FishcakeEventManagerDrop struct {
	Who        common.Address
	ActivityId *big.Int
	DropAmt    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDrop is a free log retrieval operation binding the contract event 0xd034d7fdd827eb4f41b668e02da1342cb8eced53d850fef72e1b2cd2c8387bac.
//
// Solidity: event Drop(address indexed who, uint256 indexed _activityId, uint256 _dropAmt)
func (_FishcakeEventManager *FishcakeEventManagerFilterer) FilterDrop(opts *bind.FilterOpts, who []common.Address, _activityId []*big.Int) (*FishcakeEventManagerDropIterator, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}
	var _activityIdRule []interface{}
	for _, _activityIdItem := range _activityId {
		_activityIdRule = append(_activityIdRule, _activityIdItem)
	}

	logs, sub, err := _FishcakeEventManager.contract.FilterLogs(opts, "Drop", whoRule, _activityIdRule)
	if err != nil {
		return nil, err
	}
	return &FishcakeEventManagerDropIterator{contract: _FishcakeEventManager.contract, event: "Drop", logs: logs, sub: sub}, nil
}

// WatchDrop is a free log subscription operation binding the contract event 0xd034d7fdd827eb4f41b668e02da1342cb8eced53d850fef72e1b2cd2c8387bac.
//
// Solidity: event Drop(address indexed who, uint256 indexed _activityId, uint256 _dropAmt)
func (_FishcakeEventManager *FishcakeEventManagerFilterer) WatchDrop(opts *bind.WatchOpts, sink chan<- *FishcakeEventManagerDrop, who []common.Address, _activityId []*big.Int) (event.Subscription, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}
	var _activityIdRule []interface{}
	for _, _activityIdItem := range _activityId {
		_activityIdRule = append(_activityIdRule, _activityIdItem)
	}

	logs, sub, err := _FishcakeEventManager.contract.WatchLogs(opts, "Drop", whoRule, _activityIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FishcakeEventManagerDrop)
				if err := _FishcakeEventManager.contract.UnpackLog(event, "Drop", log); err != nil {
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
func (_FishcakeEventManager *FishcakeEventManagerFilterer) ParseDrop(log types.Log) (*FishcakeEventManagerDrop, error) {
	event := new(FishcakeEventManagerDrop)
	if err := _FishcakeEventManager.contract.UnpackLog(event, "Drop", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FishcakeEventManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the FishcakeEventManager contract.
type FishcakeEventManagerInitializedIterator struct {
	Event *FishcakeEventManagerInitialized // Event containing the contract specifics and raw log

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
func (it *FishcakeEventManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FishcakeEventManagerInitialized)
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
		it.Event = new(FishcakeEventManagerInitialized)
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
func (it *FishcakeEventManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FishcakeEventManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FishcakeEventManagerInitialized represents a Initialized event raised by the FishcakeEventManager contract.
type FishcakeEventManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FishcakeEventManager *FishcakeEventManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*FishcakeEventManagerInitializedIterator, error) {

	logs, sub, err := _FishcakeEventManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &FishcakeEventManagerInitializedIterator{contract: _FishcakeEventManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FishcakeEventManager *FishcakeEventManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *FishcakeEventManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _FishcakeEventManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FishcakeEventManagerInitialized)
				if err := _FishcakeEventManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_FishcakeEventManager *FishcakeEventManagerFilterer) ParseInitialized(log types.Log) (*FishcakeEventManagerInitialized, error) {
	event := new(FishcakeEventManagerInitialized)
	if err := _FishcakeEventManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FishcakeEventManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FishcakeEventManager contract.
type FishcakeEventManagerOwnershipTransferredIterator struct {
	Event *FishcakeEventManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FishcakeEventManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FishcakeEventManagerOwnershipTransferred)
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
		it.Event = new(FishcakeEventManagerOwnershipTransferred)
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
func (it *FishcakeEventManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FishcakeEventManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FishcakeEventManagerOwnershipTransferred represents a OwnershipTransferred event raised by the FishcakeEventManager contract.
type FishcakeEventManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FishcakeEventManager *FishcakeEventManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FishcakeEventManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FishcakeEventManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FishcakeEventManagerOwnershipTransferredIterator{contract: _FishcakeEventManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FishcakeEventManager *FishcakeEventManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FishcakeEventManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FishcakeEventManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FishcakeEventManagerOwnershipTransferred)
				if err := _FishcakeEventManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FishcakeEventManager *FishcakeEventManagerFilterer) ParseOwnershipTransferred(log types.Log) (*FishcakeEventManagerOwnershipTransferred, error) {
	event := new(FishcakeEventManagerOwnershipTransferred)
	if err := _FishcakeEventManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FishcakeEventManagerSetMinePercentIterator is returned from FilterSetMinePercent and is used to iterate over the raw logs and unpacked data for SetMinePercent events raised by the FishcakeEventManager contract.
type FishcakeEventManagerSetMinePercentIterator struct {
	Event *FishcakeEventManagerSetMinePercent // Event containing the contract specifics and raw log

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
func (it *FishcakeEventManagerSetMinePercentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FishcakeEventManagerSetMinePercent)
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
		it.Event = new(FishcakeEventManagerSetMinePercent)
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
func (it *FishcakeEventManagerSetMinePercentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FishcakeEventManagerSetMinePercentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FishcakeEventManagerSetMinePercent represents a SetMinePercent event raised by the FishcakeEventManager contract.
type FishcakeEventManagerSetMinePercent struct {
	MinePercent uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetMinePercent is a free log retrieval operation binding the contract event 0x3b2d1b3b0c360549745cd56590950679dbfca066aea0899c4a5a15cca3adfdf8.
//
// Solidity: event SetMinePercent(uint8 minePercent)
func (_FishcakeEventManager *FishcakeEventManagerFilterer) FilterSetMinePercent(opts *bind.FilterOpts) (*FishcakeEventManagerSetMinePercentIterator, error) {

	logs, sub, err := _FishcakeEventManager.contract.FilterLogs(opts, "SetMinePercent")
	if err != nil {
		return nil, err
	}
	return &FishcakeEventManagerSetMinePercentIterator{contract: _FishcakeEventManager.contract, event: "SetMinePercent", logs: logs, sub: sub}, nil
}

// WatchSetMinePercent is a free log subscription operation binding the contract event 0x3b2d1b3b0c360549745cd56590950679dbfca066aea0899c4a5a15cca3adfdf8.
//
// Solidity: event SetMinePercent(uint8 minePercent)
func (_FishcakeEventManager *FishcakeEventManagerFilterer) WatchSetMinePercent(opts *bind.WatchOpts, sink chan<- *FishcakeEventManagerSetMinePercent) (event.Subscription, error) {

	logs, sub, err := _FishcakeEventManager.contract.WatchLogs(opts, "SetMinePercent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FishcakeEventManagerSetMinePercent)
				if err := _FishcakeEventManager.contract.UnpackLog(event, "SetMinePercent", log); err != nil {
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
func (_FishcakeEventManager *FishcakeEventManagerFilterer) ParseSetMinePercent(log types.Log) (*FishcakeEventManagerSetMinePercent, error) {
	event := new(FishcakeEventManagerSetMinePercent)
	if err := _FishcakeEventManager.contract.UnpackLog(event, "SetMinePercent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FishcakeEventManagerSetValidTimeIterator is returned from FilterSetValidTime and is used to iterate over the raw logs and unpacked data for SetValidTime events raised by the FishcakeEventManager contract.
type FishcakeEventManagerSetValidTimeIterator struct {
	Event *FishcakeEventManagerSetValidTime // Event containing the contract specifics and raw log

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
func (it *FishcakeEventManagerSetValidTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FishcakeEventManagerSetValidTime)
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
		it.Event = new(FishcakeEventManagerSetValidTime)
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
func (it *FishcakeEventManagerSetValidTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FishcakeEventManagerSetValidTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FishcakeEventManagerSetValidTime represents a SetValidTime event raised by the FishcakeEventManager contract.
type FishcakeEventManagerSetValidTime struct {
	Who  common.Address
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetValidTime is a free log retrieval operation binding the contract event 0x8651cd625889c15ab8587456aba8cc25b781c5fa977a1811f04967c2acad02b8.
//
// Solidity: event SetValidTime(address indexed who, uint256 _time)
func (_FishcakeEventManager *FishcakeEventManagerFilterer) FilterSetValidTime(opts *bind.FilterOpts, who []common.Address) (*FishcakeEventManagerSetValidTimeIterator, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _FishcakeEventManager.contract.FilterLogs(opts, "SetValidTime", whoRule)
	if err != nil {
		return nil, err
	}
	return &FishcakeEventManagerSetValidTimeIterator{contract: _FishcakeEventManager.contract, event: "SetValidTime", logs: logs, sub: sub}, nil
}

// WatchSetValidTime is a free log subscription operation binding the contract event 0x8651cd625889c15ab8587456aba8cc25b781c5fa977a1811f04967c2acad02b8.
//
// Solidity: event SetValidTime(address indexed who, uint256 _time)
func (_FishcakeEventManager *FishcakeEventManagerFilterer) WatchSetValidTime(opts *bind.WatchOpts, sink chan<- *FishcakeEventManagerSetValidTime, who []common.Address) (event.Subscription, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _FishcakeEventManager.contract.WatchLogs(opts, "SetValidTime", whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FishcakeEventManagerSetValidTime)
				if err := _FishcakeEventManager.contract.UnpackLog(event, "SetValidTime", log); err != nil {
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
func (_FishcakeEventManager *FishcakeEventManagerFilterer) ParseSetValidTime(log types.Log) (*FishcakeEventManagerSetValidTime, error) {
	event := new(FishcakeEventManagerSetValidTime)
	if err := _FishcakeEventManager.contract.UnpackLog(event, "SetValidTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FishcakeEventManagerTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the FishcakeEventManager contract.
type FishcakeEventManagerTransferIterator struct {
	Event *FishcakeEventManagerTransfer // Event containing the contract specifics and raw log

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
func (it *FishcakeEventManagerTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FishcakeEventManagerTransfer)
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
		it.Event = new(FishcakeEventManagerTransfer)
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
func (it *FishcakeEventManagerTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FishcakeEventManagerTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FishcakeEventManagerTransfer represents a Transfer event raised by the FishcakeEventManager contract.
type FishcakeEventManagerTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FishcakeEventManager *FishcakeEventManagerFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FishcakeEventManagerTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FishcakeEventManager.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FishcakeEventManagerTransferIterator{contract: _FishcakeEventManager.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FishcakeEventManager *FishcakeEventManagerFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FishcakeEventManagerTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FishcakeEventManager.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FishcakeEventManagerTransfer)
				if err := _FishcakeEventManager.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FishcakeEventManager *FishcakeEventManagerFilterer) ParseTransfer(log types.Log) (*FishcakeEventManagerTransfer, error) {
	event := new(FishcakeEventManagerTransfer)
	if err := _FishcakeEventManager.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
