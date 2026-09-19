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

// FeeRegistryContractMetaData contains all meta data concerning the FeeRegistryContract contract.
var FeeRegistryContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recorder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdcToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RECORDER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"studentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"termId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recordFee\",\"inputs\":[{\"name\":\"studentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"termId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usdcToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"FeeRecorded\",\"inputs\":[{\"name\":\"studentHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"termId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// FeeRegistryContractABI is the input ABI used to generate the binding from.
// Deprecated: Use FeeRegistryContractMetaData.ABI instead.
var FeeRegistryContractABI = FeeRegistryContractMetaData.ABI

// FeeRegistryContract is an auto generated Go binding around an Ethereum contract.
type FeeRegistryContract struct {
	FeeRegistryContractCaller     // Read-only binding to the contract
	FeeRegistryContractTransactor // Write-only binding to the contract
	FeeRegistryContractFilterer   // Log filterer for contract events
}

// FeeRegistryContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeeRegistryContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeRegistryContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeRegistryContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeRegistryContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeRegistryContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeRegistryContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeRegistryContractSession struct {
	Contract     *FeeRegistryContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FeeRegistryContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeRegistryContractCallerSession struct {
	Contract *FeeRegistryContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// FeeRegistryContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeRegistryContractTransactorSession struct {
	Contract     *FeeRegistryContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// FeeRegistryContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeeRegistryContractRaw struct {
	Contract *FeeRegistryContract // Generic contract binding to access the raw methods on
}

// FeeRegistryContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeRegistryContractCallerRaw struct {
	Contract *FeeRegistryContractCaller // Generic read-only contract binding to access the raw methods on
}

// FeeRegistryContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeRegistryContractTransactorRaw struct {
	Contract *FeeRegistryContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeRegistryContract creates a new instance of FeeRegistryContract, bound to a specific deployed contract.
func NewFeeRegistryContract(address common.Address, backend bind.ContractBackend) (*FeeRegistryContract, error) {
	contract, err := bindFeeRegistryContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeRegistryContract{FeeRegistryContractCaller: FeeRegistryContractCaller{contract: contract}, FeeRegistryContractTransactor: FeeRegistryContractTransactor{contract: contract}, FeeRegistryContractFilterer: FeeRegistryContractFilterer{contract: contract}}, nil
}

// NewFeeRegistryContractCaller creates a new read-only instance of FeeRegistryContract, bound to a specific deployed contract.
func NewFeeRegistryContractCaller(address common.Address, caller bind.ContractCaller) (*FeeRegistryContractCaller, error) {
	contract, err := bindFeeRegistryContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeRegistryContractCaller{contract: contract}, nil
}

// NewFeeRegistryContractTransactor creates a new write-only instance of FeeRegistryContract, bound to a specific deployed contract.
func NewFeeRegistryContractTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeRegistryContractTransactor, error) {
	contract, err := bindFeeRegistryContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeRegistryContractTransactor{contract: contract}, nil
}

// NewFeeRegistryContractFilterer creates a new log filterer instance of FeeRegistryContract, bound to a specific deployed contract.
func NewFeeRegistryContractFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeRegistryContractFilterer, error) {
	contract, err := bindFeeRegistryContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeRegistryContractFilterer{contract: contract}, nil
}

// bindFeeRegistryContract binds a generic wrapper to an already deployed contract.
func bindFeeRegistryContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FeeRegistryContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeRegistryContract *FeeRegistryContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeRegistryContract.Contract.FeeRegistryContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeRegistryContract *FeeRegistryContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeRegistryContract.Contract.FeeRegistryContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeRegistryContract *FeeRegistryContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeRegistryContract.Contract.FeeRegistryContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeRegistryContract *FeeRegistryContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeRegistryContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeRegistryContract *FeeRegistryContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeRegistryContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeRegistryContract *FeeRegistryContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeRegistryContract.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_FeeRegistryContract *FeeRegistryContractCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FeeRegistryContract.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_FeeRegistryContract *FeeRegistryContractSession) ADMINROLE() ([32]byte, error) {
	return _FeeRegistryContract.Contract.ADMINROLE(&_FeeRegistryContract.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_FeeRegistryContract *FeeRegistryContractCallerSession) ADMINROLE() ([32]byte, error) {
	return _FeeRegistryContract.Contract.ADMINROLE(&_FeeRegistryContract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FeeRegistryContract *FeeRegistryContractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FeeRegistryContract.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FeeRegistryContract *FeeRegistryContractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _FeeRegistryContract.Contract.DEFAULTADMINROLE(&_FeeRegistryContract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FeeRegistryContract *FeeRegistryContractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _FeeRegistryContract.Contract.DEFAULTADMINROLE(&_FeeRegistryContract.CallOpts)
}

// RECORDERROLE is a free data retrieval call binding the contract method 0xbea86581.
//
// Solidity: function RECORDER_ROLE() view returns(bytes32)
func (_FeeRegistryContract *FeeRegistryContractCaller) RECORDERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FeeRegistryContract.contract.Call(opts, &out, "RECORDER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RECORDERROLE is a free data retrieval call binding the contract method 0xbea86581.
//
// Solidity: function RECORDER_ROLE() view returns(bytes32)
func (_FeeRegistryContract *FeeRegistryContractSession) RECORDERROLE() ([32]byte, error) {
	return _FeeRegistryContract.Contract.RECORDERROLE(&_FeeRegistryContract.CallOpts)
}

// RECORDERROLE is a free data retrieval call binding the contract method 0xbea86581.
//
// Solidity: function RECORDER_ROLE() view returns(bytes32)
func (_FeeRegistryContract *FeeRegistryContractCallerSession) RECORDERROLE() ([32]byte, error) {
	return _FeeRegistryContract.Contract.RECORDERROLE(&_FeeRegistryContract.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0x399781b8.
//
// Solidity: function getFee(bytes32 studentHash, uint256 termId) view returns(uint256)
func (_FeeRegistryContract *FeeRegistryContractCaller) GetFee(opts *bind.CallOpts, studentHash [32]byte, termId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeeRegistryContract.contract.Call(opts, &out, "getFee", studentHash, termId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFee is a free data retrieval call binding the contract method 0x399781b8.
//
// Solidity: function getFee(bytes32 studentHash, uint256 termId) view returns(uint256)
func (_FeeRegistryContract *FeeRegistryContractSession) GetFee(studentHash [32]byte, termId *big.Int) (*big.Int, error) {
	return _FeeRegistryContract.Contract.GetFee(&_FeeRegistryContract.CallOpts, studentHash, termId)
}

// GetFee is a free data retrieval call binding the contract method 0x399781b8.
//
// Solidity: function getFee(bytes32 studentHash, uint256 termId) view returns(uint256)
func (_FeeRegistryContract *FeeRegistryContractCallerSession) GetFee(studentHash [32]byte, termId *big.Int) (*big.Int, error) {
	return _FeeRegistryContract.Contract.GetFee(&_FeeRegistryContract.CallOpts, studentHash, termId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FeeRegistryContract *FeeRegistryContractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _FeeRegistryContract.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FeeRegistryContract *FeeRegistryContractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _FeeRegistryContract.Contract.GetRoleAdmin(&_FeeRegistryContract.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FeeRegistryContract *FeeRegistryContractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _FeeRegistryContract.Contract.GetRoleAdmin(&_FeeRegistryContract.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FeeRegistryContract *FeeRegistryContractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _FeeRegistryContract.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FeeRegistryContract *FeeRegistryContractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _FeeRegistryContract.Contract.HasRole(&_FeeRegistryContract.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FeeRegistryContract *FeeRegistryContractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _FeeRegistryContract.Contract.HasRole(&_FeeRegistryContract.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FeeRegistryContract *FeeRegistryContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _FeeRegistryContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FeeRegistryContract *FeeRegistryContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FeeRegistryContract.Contract.SupportsInterface(&_FeeRegistryContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FeeRegistryContract *FeeRegistryContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FeeRegistryContract.Contract.SupportsInterface(&_FeeRegistryContract.CallOpts, interfaceId)
}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_FeeRegistryContract *FeeRegistryContractCaller) UsdcToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeRegistryContract.contract.Call(opts, &out, "usdcToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_FeeRegistryContract *FeeRegistryContractSession) UsdcToken() (common.Address, error) {
	return _FeeRegistryContract.Contract.UsdcToken(&_FeeRegistryContract.CallOpts)
}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_FeeRegistryContract *FeeRegistryContractCallerSession) UsdcToken() (common.Address, error) {
	return _FeeRegistryContract.Contract.UsdcToken(&_FeeRegistryContract.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FeeRegistryContract *FeeRegistryContractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FeeRegistryContract.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FeeRegistryContract *FeeRegistryContractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FeeRegistryContract.Contract.GrantRole(&_FeeRegistryContract.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FeeRegistryContract *FeeRegistryContractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FeeRegistryContract.Contract.GrantRole(&_FeeRegistryContract.TransactOpts, role, account)
}

// RecordFee is a paid mutator transaction binding the contract method 0x06d6e715.
//
// Solidity: function recordFee(bytes32 studentHash, uint256 termId, uint256 amount) returns()
func (_FeeRegistryContract *FeeRegistryContractTransactor) RecordFee(opts *bind.TransactOpts, studentHash [32]byte, termId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _FeeRegistryContract.contract.Transact(opts, "recordFee", studentHash, termId, amount)
}

// RecordFee is a paid mutator transaction binding the contract method 0x06d6e715.
//
// Solidity: function recordFee(bytes32 studentHash, uint256 termId, uint256 amount) returns()
func (_FeeRegistryContract *FeeRegistryContractSession) RecordFee(studentHash [32]byte, termId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _FeeRegistryContract.Contract.RecordFee(&_FeeRegistryContract.TransactOpts, studentHash, termId, amount)
}

// RecordFee is a paid mutator transaction binding the contract method 0x06d6e715.
//
// Solidity: function recordFee(bytes32 studentHash, uint256 termId, uint256 amount) returns()
func (_FeeRegistryContract *FeeRegistryContractTransactorSession) RecordFee(studentHash [32]byte, termId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _FeeRegistryContract.Contract.RecordFee(&_FeeRegistryContract.TransactOpts, studentHash, termId, amount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_FeeRegistryContract *FeeRegistryContractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _FeeRegistryContract.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_FeeRegistryContract *FeeRegistryContractSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _FeeRegistryContract.Contract.RenounceRole(&_FeeRegistryContract.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_FeeRegistryContract *FeeRegistryContractTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _FeeRegistryContract.Contract.RenounceRole(&_FeeRegistryContract.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FeeRegistryContract *FeeRegistryContractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FeeRegistryContract.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FeeRegistryContract *FeeRegistryContractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FeeRegistryContract.Contract.RevokeRole(&_FeeRegistryContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FeeRegistryContract *FeeRegistryContractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FeeRegistryContract.Contract.RevokeRole(&_FeeRegistryContract.TransactOpts, role, account)
}

// FeeRegistryContractFeeRecordedIterator is returned from FilterFeeRecorded and is used to iterate over the raw logs and unpacked data for FeeRecorded events raised by the FeeRegistryContract contract.
type FeeRegistryContractFeeRecordedIterator struct {
	Event *FeeRegistryContractFeeRecorded // Event containing the contract specifics and raw log

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
func (it *FeeRegistryContractFeeRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeRegistryContractFeeRecorded)
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
		it.Event = new(FeeRegistryContractFeeRecorded)
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
func (it *FeeRegistryContractFeeRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeRegistryContractFeeRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeRegistryContractFeeRecorded represents a FeeRecorded event raised by the FeeRegistryContract contract.
type FeeRegistryContractFeeRecorded struct {
	StudentHash [32]byte
	TermId      *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeeRecorded is a free log retrieval operation binding the contract event 0xec266d25bda89cd951b7884c2d508e9631b450a18eee737e0e71f6214bd53d6d.
//
// Solidity: event FeeRecorded(bytes32 indexed studentHash, uint256 indexed termId, uint256 amount)
func (_FeeRegistryContract *FeeRegistryContractFilterer) FilterFeeRecorded(opts *bind.FilterOpts, studentHash [][32]byte, termId []*big.Int) (*FeeRegistryContractFeeRecordedIterator, error) {

	var studentHashRule []interface{}
	for _, studentHashItem := range studentHash {
		studentHashRule = append(studentHashRule, studentHashItem)
	}
	var termIdRule []interface{}
	for _, termIdItem := range termId {
		termIdRule = append(termIdRule, termIdItem)
	}

	logs, sub, err := _FeeRegistryContract.contract.FilterLogs(opts, "FeeRecorded", studentHashRule, termIdRule)
	if err != nil {
		return nil, err
	}
	return &FeeRegistryContractFeeRecordedIterator{contract: _FeeRegistryContract.contract, event: "FeeRecorded", logs: logs, sub: sub}, nil
}

// WatchFeeRecorded is a free log subscription operation binding the contract event 0xec266d25bda89cd951b7884c2d508e9631b450a18eee737e0e71f6214bd53d6d.
//
// Solidity: event FeeRecorded(bytes32 indexed studentHash, uint256 indexed termId, uint256 amount)
func (_FeeRegistryContract *FeeRegistryContractFilterer) WatchFeeRecorded(opts *bind.WatchOpts, sink chan<- *FeeRegistryContractFeeRecorded, studentHash [][32]byte, termId []*big.Int) (event.Subscription, error) {

	var studentHashRule []interface{}
	for _, studentHashItem := range studentHash {
		studentHashRule = append(studentHashRule, studentHashItem)
	}
	var termIdRule []interface{}
	for _, termIdItem := range termId {
		termIdRule = append(termIdRule, termIdItem)
	}

	logs, sub, err := _FeeRegistryContract.contract.WatchLogs(opts, "FeeRecorded", studentHashRule, termIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeRegistryContractFeeRecorded)
				if err := _FeeRegistryContract.contract.UnpackLog(event, "FeeRecorded", log); err != nil {
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

// ParseFeeRecorded is a log parse operation binding the contract event 0xec266d25bda89cd951b7884c2d508e9631b450a18eee737e0e71f6214bd53d6d.
//
// Solidity: event FeeRecorded(bytes32 indexed studentHash, uint256 indexed termId, uint256 amount)
func (_FeeRegistryContract *FeeRegistryContractFilterer) ParseFeeRecorded(log types.Log) (*FeeRegistryContractFeeRecorded, error) {
	event := new(FeeRegistryContractFeeRecorded)
	if err := _FeeRegistryContract.contract.UnpackLog(event, "FeeRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeRegistryContractRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the FeeRegistryContract contract.
type FeeRegistryContractRoleAdminChangedIterator struct {
	Event *FeeRegistryContractRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *FeeRegistryContractRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeRegistryContractRoleAdminChanged)
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
		it.Event = new(FeeRegistryContractRoleAdminChanged)
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
func (it *FeeRegistryContractRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeRegistryContractRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeRegistryContractRoleAdminChanged represents a RoleAdminChanged event raised by the FeeRegistryContract contract.
type FeeRegistryContractRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FeeRegistryContract *FeeRegistryContractFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*FeeRegistryContractRoleAdminChangedIterator, error) {

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

	logs, sub, err := _FeeRegistryContract.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &FeeRegistryContractRoleAdminChangedIterator{contract: _FeeRegistryContract.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FeeRegistryContract *FeeRegistryContractFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *FeeRegistryContractRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _FeeRegistryContract.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeRegistryContractRoleAdminChanged)
				if err := _FeeRegistryContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_FeeRegistryContract *FeeRegistryContractFilterer) ParseRoleAdminChanged(log types.Log) (*FeeRegistryContractRoleAdminChanged, error) {
	event := new(FeeRegistryContractRoleAdminChanged)
	if err := _FeeRegistryContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeRegistryContractRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the FeeRegistryContract contract.
type FeeRegistryContractRoleGrantedIterator struct {
	Event *FeeRegistryContractRoleGranted // Event containing the contract specifics and raw log

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
func (it *FeeRegistryContractRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeRegistryContractRoleGranted)
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
		it.Event = new(FeeRegistryContractRoleGranted)
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
func (it *FeeRegistryContractRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeRegistryContractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeRegistryContractRoleGranted represents a RoleGranted event raised by the FeeRegistryContract contract.
type FeeRegistryContractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FeeRegistryContract *FeeRegistryContractFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FeeRegistryContractRoleGrantedIterator, error) {

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

	logs, sub, err := _FeeRegistryContract.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FeeRegistryContractRoleGrantedIterator{contract: _FeeRegistryContract.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FeeRegistryContract *FeeRegistryContractFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *FeeRegistryContractRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _FeeRegistryContract.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeRegistryContractRoleGranted)
				if err := _FeeRegistryContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_FeeRegistryContract *FeeRegistryContractFilterer) ParseRoleGranted(log types.Log) (*FeeRegistryContractRoleGranted, error) {
	event := new(FeeRegistryContractRoleGranted)
	if err := _FeeRegistryContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeRegistryContractRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the FeeRegistryContract contract.
type FeeRegistryContractRoleRevokedIterator struct {
	Event *FeeRegistryContractRoleRevoked // Event containing the contract specifics and raw log

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
func (it *FeeRegistryContractRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeRegistryContractRoleRevoked)
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
		it.Event = new(FeeRegistryContractRoleRevoked)
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
func (it *FeeRegistryContractRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeRegistryContractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeRegistryContractRoleRevoked represents a RoleRevoked event raised by the FeeRegistryContract contract.
type FeeRegistryContractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FeeRegistryContract *FeeRegistryContractFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FeeRegistryContractRoleRevokedIterator, error) {

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

	logs, sub, err := _FeeRegistryContract.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FeeRegistryContractRoleRevokedIterator{contract: _FeeRegistryContract.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FeeRegistryContract *FeeRegistryContractFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *FeeRegistryContractRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _FeeRegistryContract.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeRegistryContractRoleRevoked)
				if err := _FeeRegistryContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_FeeRegistryContract *FeeRegistryContractFilterer) ParseRoleRevoked(log types.Log) (*FeeRegistryContractRoleRevoked, error) {
	event := new(FeeRegistryContractRoleRevoked)
	if err := _FeeRegistryContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
