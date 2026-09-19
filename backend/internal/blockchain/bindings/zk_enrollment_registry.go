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

// ZKEnrollmentRegistryMetaData contains all meta data concerning the ZKEnrollmentRegistry contract.
var ZKEnrollmentRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_verifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REGISTRAR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"enrollmentMerkleRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isEnrollmentVerified\",\"inputs\":[{\"name\":\"nullifierHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"student\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proveEnrollment\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nullifierHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"claimedRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateMerkleRoot\",\"inputs\":[{\"name\":\"newRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifiedNullifier\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractMockZKVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"EnrollmentProofVerified\",\"inputs\":[{\"name\":\"nullifierHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"student\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MerkleRootUpdated\",\"inputs\":[{\"name\":\"newRoot\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"updatedBy\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZKRegistry__InvalidProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZKRegistry__NullifierAlreadyUsed\",\"inputs\":[]}]",
}

// ZKEnrollmentRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ZKEnrollmentRegistryMetaData.ABI instead.
var ZKEnrollmentRegistryABI = ZKEnrollmentRegistryMetaData.ABI

// ZKEnrollmentRegistry is an auto generated Go binding around an Ethereum contract.
type ZKEnrollmentRegistry struct {
	ZKEnrollmentRegistryCaller     // Read-only binding to the contract
	ZKEnrollmentRegistryTransactor // Write-only binding to the contract
	ZKEnrollmentRegistryFilterer   // Log filterer for contract events
}

// ZKEnrollmentRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZKEnrollmentRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKEnrollmentRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZKEnrollmentRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKEnrollmentRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZKEnrollmentRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKEnrollmentRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZKEnrollmentRegistrySession struct {
	Contract     *ZKEnrollmentRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ZKEnrollmentRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZKEnrollmentRegistryCallerSession struct {
	Contract *ZKEnrollmentRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ZKEnrollmentRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZKEnrollmentRegistryTransactorSession struct {
	Contract     *ZKEnrollmentRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ZKEnrollmentRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZKEnrollmentRegistryRaw struct {
	Contract *ZKEnrollmentRegistry // Generic contract binding to access the raw methods on
}

// ZKEnrollmentRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZKEnrollmentRegistryCallerRaw struct {
	Contract *ZKEnrollmentRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ZKEnrollmentRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZKEnrollmentRegistryTransactorRaw struct {
	Contract *ZKEnrollmentRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZKEnrollmentRegistry creates a new instance of ZKEnrollmentRegistry, bound to a specific deployed contract.
func NewZKEnrollmentRegistry(address common.Address, backend bind.ContractBackend) (*ZKEnrollmentRegistry, error) {
	contract, err := bindZKEnrollmentRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZKEnrollmentRegistry{ZKEnrollmentRegistryCaller: ZKEnrollmentRegistryCaller{contract: contract}, ZKEnrollmentRegistryTransactor: ZKEnrollmentRegistryTransactor{contract: contract}, ZKEnrollmentRegistryFilterer: ZKEnrollmentRegistryFilterer{contract: contract}}, nil
}

// NewZKEnrollmentRegistryCaller creates a new read-only instance of ZKEnrollmentRegistry, bound to a specific deployed contract.
func NewZKEnrollmentRegistryCaller(address common.Address, caller bind.ContractCaller) (*ZKEnrollmentRegistryCaller, error) {
	contract, err := bindZKEnrollmentRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZKEnrollmentRegistryCaller{contract: contract}, nil
}

// NewZKEnrollmentRegistryTransactor creates a new write-only instance of ZKEnrollmentRegistry, bound to a specific deployed contract.
func NewZKEnrollmentRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ZKEnrollmentRegistryTransactor, error) {
	contract, err := bindZKEnrollmentRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZKEnrollmentRegistryTransactor{contract: contract}, nil
}

// NewZKEnrollmentRegistryFilterer creates a new log filterer instance of ZKEnrollmentRegistry, bound to a specific deployed contract.
func NewZKEnrollmentRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ZKEnrollmentRegistryFilterer, error) {
	contract, err := bindZKEnrollmentRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZKEnrollmentRegistryFilterer{contract: contract}, nil
}

// bindZKEnrollmentRegistry binds a generic wrapper to an already deployed contract.
func bindZKEnrollmentRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZKEnrollmentRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKEnrollmentRegistry.Contract.ZKEnrollmentRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.Contract.ZKEnrollmentRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.Contract.ZKEnrollmentRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKEnrollmentRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKEnrollmentRegistry.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistrySession) ADMINROLE() ([32]byte, error) {
	return _ZKEnrollmentRegistry.Contract.ADMINROLE(&_ZKEnrollmentRegistry.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCallerSession) ADMINROLE() ([32]byte, error) {
	return _ZKEnrollmentRegistry.Contract.ADMINROLE(&_ZKEnrollmentRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKEnrollmentRegistry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZKEnrollmentRegistry.Contract.DEFAULTADMINROLE(&_ZKEnrollmentRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZKEnrollmentRegistry.Contract.DEFAULTADMINROLE(&_ZKEnrollmentRegistry.CallOpts)
}

// REGISTRARROLE is a free data retrieval call binding the contract method 0xf68e9553.
//
// Solidity: function REGISTRAR_ROLE() view returns(bytes32)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCaller) REGISTRARROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKEnrollmentRegistry.contract.Call(opts, &out, "REGISTRAR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REGISTRARROLE is a free data retrieval call binding the contract method 0xf68e9553.
//
// Solidity: function REGISTRAR_ROLE() view returns(bytes32)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistrySession) REGISTRARROLE() ([32]byte, error) {
	return _ZKEnrollmentRegistry.Contract.REGISTRARROLE(&_ZKEnrollmentRegistry.CallOpts)
}

// REGISTRARROLE is a free data retrieval call binding the contract method 0xf68e9553.
//
// Solidity: function REGISTRAR_ROLE() view returns(bytes32)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCallerSession) REGISTRARROLE() ([32]byte, error) {
	return _ZKEnrollmentRegistry.Contract.REGISTRARROLE(&_ZKEnrollmentRegistry.CallOpts)
}

// EnrollmentMerkleRoot is a free data retrieval call binding the contract method 0x1d35e0fd.
//
// Solidity: function enrollmentMerkleRoot() view returns(bytes32)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCaller) EnrollmentMerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKEnrollmentRegistry.contract.Call(opts, &out, "enrollmentMerkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EnrollmentMerkleRoot is a free data retrieval call binding the contract method 0x1d35e0fd.
//
// Solidity: function enrollmentMerkleRoot() view returns(bytes32)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistrySession) EnrollmentMerkleRoot() ([32]byte, error) {
	return _ZKEnrollmentRegistry.Contract.EnrollmentMerkleRoot(&_ZKEnrollmentRegistry.CallOpts)
}

// EnrollmentMerkleRoot is a free data retrieval call binding the contract method 0x1d35e0fd.
//
// Solidity: function enrollmentMerkleRoot() view returns(bytes32)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCallerSession) EnrollmentMerkleRoot() ([32]byte, error) {
	return _ZKEnrollmentRegistry.Contract.EnrollmentMerkleRoot(&_ZKEnrollmentRegistry.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ZKEnrollmentRegistry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZKEnrollmentRegistry.Contract.GetRoleAdmin(&_ZKEnrollmentRegistry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZKEnrollmentRegistry.Contract.GetRoleAdmin(&_ZKEnrollmentRegistry.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ZKEnrollmentRegistry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZKEnrollmentRegistry.Contract.HasRole(&_ZKEnrollmentRegistry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZKEnrollmentRegistry.Contract.HasRole(&_ZKEnrollmentRegistry.CallOpts, role, account)
}

// IsEnrollmentVerified is a free data retrieval call binding the contract method 0xede0ad1a.
//
// Solidity: function isEnrollmentVerified(bytes32 nullifierHash) view returns(address student)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCaller) IsEnrollmentVerified(opts *bind.CallOpts, nullifierHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ZKEnrollmentRegistry.contract.Call(opts, &out, "isEnrollmentVerified", nullifierHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IsEnrollmentVerified is a free data retrieval call binding the contract method 0xede0ad1a.
//
// Solidity: function isEnrollmentVerified(bytes32 nullifierHash) view returns(address student)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistrySession) IsEnrollmentVerified(nullifierHash [32]byte) (common.Address, error) {
	return _ZKEnrollmentRegistry.Contract.IsEnrollmentVerified(&_ZKEnrollmentRegistry.CallOpts, nullifierHash)
}

// IsEnrollmentVerified is a free data retrieval call binding the contract method 0xede0ad1a.
//
// Solidity: function isEnrollmentVerified(bytes32 nullifierHash) view returns(address student)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCallerSession) IsEnrollmentVerified(nullifierHash [32]byte) (common.Address, error) {
	return _ZKEnrollmentRegistry.Contract.IsEnrollmentVerified(&_ZKEnrollmentRegistry.CallOpts, nullifierHash)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ZKEnrollmentRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZKEnrollmentRegistry.Contract.SupportsInterface(&_ZKEnrollmentRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZKEnrollmentRegistry.Contract.SupportsInterface(&_ZKEnrollmentRegistry.CallOpts, interfaceId)
}

// VerifiedNullifier is a free data retrieval call binding the contract method 0x68b2219b.
//
// Solidity: function verifiedNullifier(bytes32 ) view returns(address)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCaller) VerifiedNullifier(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ZKEnrollmentRegistry.contract.Call(opts, &out, "verifiedNullifier", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VerifiedNullifier is a free data retrieval call binding the contract method 0x68b2219b.
//
// Solidity: function verifiedNullifier(bytes32 ) view returns(address)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistrySession) VerifiedNullifier(arg0 [32]byte) (common.Address, error) {
	return _ZKEnrollmentRegistry.Contract.VerifiedNullifier(&_ZKEnrollmentRegistry.CallOpts, arg0)
}

// VerifiedNullifier is a free data retrieval call binding the contract method 0x68b2219b.
//
// Solidity: function verifiedNullifier(bytes32 ) view returns(address)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCallerSession) VerifiedNullifier(arg0 [32]byte) (common.Address, error) {
	return _ZKEnrollmentRegistry.Contract.VerifiedNullifier(&_ZKEnrollmentRegistry.CallOpts, arg0)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZKEnrollmentRegistry.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistrySession) Verifier() (common.Address, error) {
	return _ZKEnrollmentRegistry.Contract.Verifier(&_ZKEnrollmentRegistry.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryCallerSession) Verifier() (common.Address, error) {
	return _ZKEnrollmentRegistry.Contract.Verifier(&_ZKEnrollmentRegistry.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.Contract.GrantRole(&_ZKEnrollmentRegistry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.Contract.GrantRole(&_ZKEnrollmentRegistry.TransactOpts, role, account)
}

// ProveEnrollment is a paid mutator transaction binding the contract method 0x0980d9d9.
//
// Solidity: function proveEnrollment(bytes proof, bytes32 nullifierHash, bytes32 claimedRoot) returns()
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryTransactor) ProveEnrollment(opts *bind.TransactOpts, proof []byte, nullifierHash [32]byte, claimedRoot [32]byte) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.contract.Transact(opts, "proveEnrollment", proof, nullifierHash, claimedRoot)
}

// ProveEnrollment is a paid mutator transaction binding the contract method 0x0980d9d9.
//
// Solidity: function proveEnrollment(bytes proof, bytes32 nullifierHash, bytes32 claimedRoot) returns()
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistrySession) ProveEnrollment(proof []byte, nullifierHash [32]byte, claimedRoot [32]byte) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.Contract.ProveEnrollment(&_ZKEnrollmentRegistry.TransactOpts, proof, nullifierHash, claimedRoot)
}

// ProveEnrollment is a paid mutator transaction binding the contract method 0x0980d9d9.
//
// Solidity: function proveEnrollment(bytes proof, bytes32 nullifierHash, bytes32 claimedRoot) returns()
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryTransactorSession) ProveEnrollment(proof []byte, nullifierHash [32]byte, claimedRoot [32]byte) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.Contract.ProveEnrollment(&_ZKEnrollmentRegistry.TransactOpts, proof, nullifierHash, claimedRoot)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistrySession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.Contract.RenounceRole(&_ZKEnrollmentRegistry.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.Contract.RenounceRole(&_ZKEnrollmentRegistry.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.Contract.RevokeRole(&_ZKEnrollmentRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.Contract.RevokeRole(&_ZKEnrollmentRegistry.TransactOpts, role, account)
}

// UpdateMerkleRoot is a paid mutator transaction binding the contract method 0x4783f0ef.
//
// Solidity: function updateMerkleRoot(bytes32 newRoot) returns()
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryTransactor) UpdateMerkleRoot(opts *bind.TransactOpts, newRoot [32]byte) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.contract.Transact(opts, "updateMerkleRoot", newRoot)
}

// UpdateMerkleRoot is a paid mutator transaction binding the contract method 0x4783f0ef.
//
// Solidity: function updateMerkleRoot(bytes32 newRoot) returns()
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistrySession) UpdateMerkleRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.Contract.UpdateMerkleRoot(&_ZKEnrollmentRegistry.TransactOpts, newRoot)
}

// UpdateMerkleRoot is a paid mutator transaction binding the contract method 0x4783f0ef.
//
// Solidity: function updateMerkleRoot(bytes32 newRoot) returns()
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryTransactorSession) UpdateMerkleRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _ZKEnrollmentRegistry.Contract.UpdateMerkleRoot(&_ZKEnrollmentRegistry.TransactOpts, newRoot)
}

// ZKEnrollmentRegistryEnrollmentProofVerifiedIterator is returned from FilterEnrollmentProofVerified and is used to iterate over the raw logs and unpacked data for EnrollmentProofVerified events raised by the ZKEnrollmentRegistry contract.
type ZKEnrollmentRegistryEnrollmentProofVerifiedIterator struct {
	Event *ZKEnrollmentRegistryEnrollmentProofVerified // Event containing the contract specifics and raw log

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
func (it *ZKEnrollmentRegistryEnrollmentProofVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEnrollmentRegistryEnrollmentProofVerified)
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
		it.Event = new(ZKEnrollmentRegistryEnrollmentProofVerified)
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
func (it *ZKEnrollmentRegistryEnrollmentProofVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEnrollmentRegistryEnrollmentProofVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEnrollmentRegistryEnrollmentProofVerified represents a EnrollmentProofVerified event raised by the ZKEnrollmentRegistry contract.
type ZKEnrollmentRegistryEnrollmentProofVerified struct {
	NullifierHash [32]byte
	Student       common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEnrollmentProofVerified is a free log retrieval operation binding the contract event 0x152ac2c5828ee9c389053e2211ef80b6e2383ecb5db6000837597583971a3ee6.
//
// Solidity: event EnrollmentProofVerified(bytes32 indexed nullifierHash, address indexed student)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryFilterer) FilterEnrollmentProofVerified(opts *bind.FilterOpts, nullifierHash [][32]byte, student []common.Address) (*ZKEnrollmentRegistryEnrollmentProofVerifiedIterator, error) {

	var nullifierHashRule []interface{}
	for _, nullifierHashItem := range nullifierHash {
		nullifierHashRule = append(nullifierHashRule, nullifierHashItem)
	}
	var studentRule []interface{}
	for _, studentItem := range student {
		studentRule = append(studentRule, studentItem)
	}

	logs, sub, err := _ZKEnrollmentRegistry.contract.FilterLogs(opts, "EnrollmentProofVerified", nullifierHashRule, studentRule)
	if err != nil {
		return nil, err
	}
	return &ZKEnrollmentRegistryEnrollmentProofVerifiedIterator{contract: _ZKEnrollmentRegistry.contract, event: "EnrollmentProofVerified", logs: logs, sub: sub}, nil
}

// WatchEnrollmentProofVerified is a free log subscription operation binding the contract event 0x152ac2c5828ee9c389053e2211ef80b6e2383ecb5db6000837597583971a3ee6.
//
// Solidity: event EnrollmentProofVerified(bytes32 indexed nullifierHash, address indexed student)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryFilterer) WatchEnrollmentProofVerified(opts *bind.WatchOpts, sink chan<- *ZKEnrollmentRegistryEnrollmentProofVerified, nullifierHash [][32]byte, student []common.Address) (event.Subscription, error) {

	var nullifierHashRule []interface{}
	for _, nullifierHashItem := range nullifierHash {
		nullifierHashRule = append(nullifierHashRule, nullifierHashItem)
	}
	var studentRule []interface{}
	for _, studentItem := range student {
		studentRule = append(studentRule, studentItem)
	}

	logs, sub, err := _ZKEnrollmentRegistry.contract.WatchLogs(opts, "EnrollmentProofVerified", nullifierHashRule, studentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEnrollmentRegistryEnrollmentProofVerified)
				if err := _ZKEnrollmentRegistry.contract.UnpackLog(event, "EnrollmentProofVerified", log); err != nil {
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

// ParseEnrollmentProofVerified is a log parse operation binding the contract event 0x152ac2c5828ee9c389053e2211ef80b6e2383ecb5db6000837597583971a3ee6.
//
// Solidity: event EnrollmentProofVerified(bytes32 indexed nullifierHash, address indexed student)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryFilterer) ParseEnrollmentProofVerified(log types.Log) (*ZKEnrollmentRegistryEnrollmentProofVerified, error) {
	event := new(ZKEnrollmentRegistryEnrollmentProofVerified)
	if err := _ZKEnrollmentRegistry.contract.UnpackLog(event, "EnrollmentProofVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKEnrollmentRegistryMerkleRootUpdatedIterator is returned from FilterMerkleRootUpdated and is used to iterate over the raw logs and unpacked data for MerkleRootUpdated events raised by the ZKEnrollmentRegistry contract.
type ZKEnrollmentRegistryMerkleRootUpdatedIterator struct {
	Event *ZKEnrollmentRegistryMerkleRootUpdated // Event containing the contract specifics and raw log

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
func (it *ZKEnrollmentRegistryMerkleRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEnrollmentRegistryMerkleRootUpdated)
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
		it.Event = new(ZKEnrollmentRegistryMerkleRootUpdated)
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
func (it *ZKEnrollmentRegistryMerkleRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEnrollmentRegistryMerkleRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEnrollmentRegistryMerkleRootUpdated represents a MerkleRootUpdated event raised by the ZKEnrollmentRegistry contract.
type ZKEnrollmentRegistryMerkleRootUpdated struct {
	NewRoot   [32]byte
	UpdatedBy common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMerkleRootUpdated is a free log retrieval operation binding the contract event 0xe2c8b1f4bcb8086dc9805a9868ee616914948de41a09786213a2284a8ffba27b.
//
// Solidity: event MerkleRootUpdated(bytes32 indexed newRoot, address indexed updatedBy)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryFilterer) FilterMerkleRootUpdated(opts *bind.FilterOpts, newRoot [][32]byte, updatedBy []common.Address) (*ZKEnrollmentRegistryMerkleRootUpdatedIterator, error) {

	var newRootRule []interface{}
	for _, newRootItem := range newRoot {
		newRootRule = append(newRootRule, newRootItem)
	}
	var updatedByRule []interface{}
	for _, updatedByItem := range updatedBy {
		updatedByRule = append(updatedByRule, updatedByItem)
	}

	logs, sub, err := _ZKEnrollmentRegistry.contract.FilterLogs(opts, "MerkleRootUpdated", newRootRule, updatedByRule)
	if err != nil {
		return nil, err
	}
	return &ZKEnrollmentRegistryMerkleRootUpdatedIterator{contract: _ZKEnrollmentRegistry.contract, event: "MerkleRootUpdated", logs: logs, sub: sub}, nil
}

// WatchMerkleRootUpdated is a free log subscription operation binding the contract event 0xe2c8b1f4bcb8086dc9805a9868ee616914948de41a09786213a2284a8ffba27b.
//
// Solidity: event MerkleRootUpdated(bytes32 indexed newRoot, address indexed updatedBy)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryFilterer) WatchMerkleRootUpdated(opts *bind.WatchOpts, sink chan<- *ZKEnrollmentRegistryMerkleRootUpdated, newRoot [][32]byte, updatedBy []common.Address) (event.Subscription, error) {

	var newRootRule []interface{}
	for _, newRootItem := range newRoot {
		newRootRule = append(newRootRule, newRootItem)
	}
	var updatedByRule []interface{}
	for _, updatedByItem := range updatedBy {
		updatedByRule = append(updatedByRule, updatedByItem)
	}

	logs, sub, err := _ZKEnrollmentRegistry.contract.WatchLogs(opts, "MerkleRootUpdated", newRootRule, updatedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEnrollmentRegistryMerkleRootUpdated)
				if err := _ZKEnrollmentRegistry.contract.UnpackLog(event, "MerkleRootUpdated", log); err != nil {
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

// ParseMerkleRootUpdated is a log parse operation binding the contract event 0xe2c8b1f4bcb8086dc9805a9868ee616914948de41a09786213a2284a8ffba27b.
//
// Solidity: event MerkleRootUpdated(bytes32 indexed newRoot, address indexed updatedBy)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryFilterer) ParseMerkleRootUpdated(log types.Log) (*ZKEnrollmentRegistryMerkleRootUpdated, error) {
	event := new(ZKEnrollmentRegistryMerkleRootUpdated)
	if err := _ZKEnrollmentRegistry.contract.UnpackLog(event, "MerkleRootUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKEnrollmentRegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ZKEnrollmentRegistry contract.
type ZKEnrollmentRegistryRoleAdminChangedIterator struct {
	Event *ZKEnrollmentRegistryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ZKEnrollmentRegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEnrollmentRegistryRoleAdminChanged)
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
		it.Event = new(ZKEnrollmentRegistryRoleAdminChanged)
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
func (it *ZKEnrollmentRegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEnrollmentRegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEnrollmentRegistryRoleAdminChanged represents a RoleAdminChanged event raised by the ZKEnrollmentRegistry contract.
type ZKEnrollmentRegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ZKEnrollmentRegistryRoleAdminChangedIterator, error) {

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

	logs, sub, err := _ZKEnrollmentRegistry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ZKEnrollmentRegistryRoleAdminChangedIterator{contract: _ZKEnrollmentRegistry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ZKEnrollmentRegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ZKEnrollmentRegistry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEnrollmentRegistryRoleAdminChanged)
				if err := _ZKEnrollmentRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryFilterer) ParseRoleAdminChanged(log types.Log) (*ZKEnrollmentRegistryRoleAdminChanged, error) {
	event := new(ZKEnrollmentRegistryRoleAdminChanged)
	if err := _ZKEnrollmentRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKEnrollmentRegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ZKEnrollmentRegistry contract.
type ZKEnrollmentRegistryRoleGrantedIterator struct {
	Event *ZKEnrollmentRegistryRoleGranted // Event containing the contract specifics and raw log

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
func (it *ZKEnrollmentRegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEnrollmentRegistryRoleGranted)
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
		it.Event = new(ZKEnrollmentRegistryRoleGranted)
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
func (it *ZKEnrollmentRegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEnrollmentRegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEnrollmentRegistryRoleGranted represents a RoleGranted event raised by the ZKEnrollmentRegistry contract.
type ZKEnrollmentRegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZKEnrollmentRegistryRoleGrantedIterator, error) {

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

	logs, sub, err := _ZKEnrollmentRegistry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZKEnrollmentRegistryRoleGrantedIterator{contract: _ZKEnrollmentRegistry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ZKEnrollmentRegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ZKEnrollmentRegistry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEnrollmentRegistryRoleGranted)
				if err := _ZKEnrollmentRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryFilterer) ParseRoleGranted(log types.Log) (*ZKEnrollmentRegistryRoleGranted, error) {
	event := new(ZKEnrollmentRegistryRoleGranted)
	if err := _ZKEnrollmentRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKEnrollmentRegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ZKEnrollmentRegistry contract.
type ZKEnrollmentRegistryRoleRevokedIterator struct {
	Event *ZKEnrollmentRegistryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ZKEnrollmentRegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEnrollmentRegistryRoleRevoked)
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
		it.Event = new(ZKEnrollmentRegistryRoleRevoked)
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
func (it *ZKEnrollmentRegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEnrollmentRegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEnrollmentRegistryRoleRevoked represents a RoleRevoked event raised by the ZKEnrollmentRegistry contract.
type ZKEnrollmentRegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZKEnrollmentRegistryRoleRevokedIterator, error) {

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

	logs, sub, err := _ZKEnrollmentRegistry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZKEnrollmentRegistryRoleRevokedIterator{contract: _ZKEnrollmentRegistry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ZKEnrollmentRegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ZKEnrollmentRegistry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEnrollmentRegistryRoleRevoked)
				if err := _ZKEnrollmentRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ZKEnrollmentRegistry *ZKEnrollmentRegistryFilterer) ParseRoleRevoked(log types.Log) (*ZKEnrollmentRegistryRoleRevoked, error) {
	event := new(ZKEnrollmentRegistryRoleRevoked)
	if err := _ZKEnrollmentRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
