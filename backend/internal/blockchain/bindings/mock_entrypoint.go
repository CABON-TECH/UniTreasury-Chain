// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"time"

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
	_ = time.Tick
	_ = context.Background
)

// UserOperation is an auto generated low-level Go binding around an user-defined struct.
type UserOperation struct {
	Sender               common.Address
	Nonce                *big.Int
	InitCode             []byte
	CallData             []byte
	CallGasLimit         *big.Int
	VerificationGasLimit *big.Int
	PreVerificationGas   *big.Int
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
	PaymasterAndData     []byte
	Signature            []byte
}

// MockEntryPointMetaData contains all meta data concerning the MockEntryPoint contract.
var MockEntryPointMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"handleOps\",\"inputs\":[{\"name\":\"ops\",\"type\":\"tuple[]\",\"internalType\":\"structUserOperation[]\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verificationGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"preVerificationGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymasterAndData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// MockEntryPointABI is the input ABI used to generate the binding from.
// Deprecated: Use MockEntryPointMetaData.ABI instead.
var MockEntryPointABI = MockEntryPointMetaData.ABI

// MockEntryPoint is an auto generated Go binding around an Ethereum contract.
type MockEntryPoint struct {
	MockEntryPointCaller     // Read-only binding to the contract
	MockEntryPointTransactor // Write-only binding to the contract
	MockEntryPointFilterer   // Log filterer for contract events
}

// MockEntryPointCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockEntryPointCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockEntryPointTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockEntryPointTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockEntryPointFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockEntryPointFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockEntryPointSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockEntryPointSession struct {
	Contract     *MockEntryPoint   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockEntryPointCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockEntryPointCallerSession struct {
	Contract *MockEntryPointCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MockEntryPointTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockEntryPointTransactorSession struct {
	Contract     *MockEntryPointTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MockEntryPointRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockEntryPointRaw struct {
	Contract *MockEntryPoint // Generic contract binding to access the raw methods on
}

// MockEntryPointCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockEntryPointCallerRaw struct {
	Contract *MockEntryPointCaller // Generic read-only contract binding to access the raw methods on
}

// MockEntryPointTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockEntryPointTransactorRaw struct {
	Contract *MockEntryPointTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockEntryPoint creates a new instance of MockEntryPoint, bound to a specific deployed contract.
func NewMockEntryPoint(address common.Address, backend bind.ContractBackend) (*MockEntryPoint, error) {
	contract, err := bindMockEntryPoint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockEntryPoint{MockEntryPointCaller: MockEntryPointCaller{contract: contract}, MockEntryPointTransactor: MockEntryPointTransactor{contract: contract}, MockEntryPointFilterer: MockEntryPointFilterer{contract: contract}}, nil
}

// NewMockEntryPointCaller creates a new read-only instance of MockEntryPoint, bound to a specific deployed contract.
func NewMockEntryPointCaller(address common.Address, caller bind.ContractCaller) (*MockEntryPointCaller, error) {
	contract, err := bindMockEntryPoint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockEntryPointCaller{contract: contract}, nil
}

// NewMockEntryPointTransactor creates a new write-only instance of MockEntryPoint, bound to a specific deployed contract.
func NewMockEntryPointTransactor(address common.Address, transactor bind.ContractTransactor) (*MockEntryPointTransactor, error) {
	contract, err := bindMockEntryPoint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockEntryPointTransactor{contract: contract}, nil
}

// NewMockEntryPointFilterer creates a new log filterer instance of MockEntryPoint, bound to a specific deployed contract.
func NewMockEntryPointFilterer(address common.Address, filterer bind.ContractFilterer) (*MockEntryPointFilterer, error) {
	contract, err := bindMockEntryPoint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockEntryPointFilterer{contract: contract}, nil
}

// bindMockEntryPoint binds a generic wrapper to an already deployed contract.
func bindMockEntryPoint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockEntryPointMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockEntryPoint *MockEntryPointRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockEntryPoint.Contract.MockEntryPointCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockEntryPoint *MockEntryPointRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockEntryPoint.Contract.MockEntryPointTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockEntryPoint *MockEntryPointRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockEntryPoint.Contract.MockEntryPointTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockEntryPoint *MockEntryPointCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockEntryPoint.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockEntryPoint *MockEntryPointTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockEntryPoint.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockEntryPoint *MockEntryPointTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockEntryPoint.Contract.contract.Transact(opts, method, params...)
}

// HandleOps is a paid mutator transaction binding the contract method 0x1fad948c.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] ops, address beneficiary) returns()
func (_MockEntryPoint *MockEntryPointTransactor) HandleOps(opts *bind.TransactOpts, ops []UserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _MockEntryPoint.contract.Transact(opts, "handleOps", ops, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x1fad948c.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] ops, address beneficiary) returns()
func (_MockEntryPoint *MockEntryPointSession) HandleOps(ops []UserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _MockEntryPoint.Contract.HandleOps(&_MockEntryPoint.TransactOpts, ops, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x1fad948c.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] ops, address beneficiary) returns()
func (_MockEntryPoint *MockEntryPointTransactorSession) HandleOps(ops []UserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _MockEntryPoint.Contract.HandleOps(&_MockEntryPoint.TransactOpts, ops, beneficiary)
}
