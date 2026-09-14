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

// ScholarshipEscrowContractMetaData contains all meta data concerning the ScholarshipEscrowContract contract.
var ScholarshipEscrowContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"attestor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdcToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ATTESTOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SPONSOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimTranche\",\"inputs\":[{\"name\":\"fundId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"studentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"trancheIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"merkleProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createFund\",\"inputs\":[{\"name\":\"sponsor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"totalAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"trancheCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"trancheAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"funds\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"sponsor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"totalAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"releasedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"trancheCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"trancheAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paused\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasClaimed\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseFund\",\"inputs\":[{\"name\":\"fundId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"publishTrancheRoot\",\"inputs\":[{\"name\":\"fundId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"trancheIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTrustedAttestor\",\"inputs\":[{\"name\":\"newAttestor\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"trancheRoots\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"trustedAttestor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpauseFund\",\"inputs\":[{\"name\":\"fundId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"usdcToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"FundCreated\",\"inputs\":[{\"name\":\"fundId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"sponsor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"totalAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"trancheCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"trancheAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MerkleRootPublished\",\"inputs\":[{\"name\":\"fundId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"trancheIndex\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TrancheClaimed\",\"inputs\":[{\"name\":\"fundId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"studentHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"trancheIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
}

// ScholarshipEscrowContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ScholarshipEscrowContractMetaData.ABI instead.
var ScholarshipEscrowContractABI = ScholarshipEscrowContractMetaData.ABI

// ScholarshipEscrowContract is an auto generated Go binding around an Ethereum contract.
type ScholarshipEscrowContract struct {
	ScholarshipEscrowContractCaller     // Read-only binding to the contract
	ScholarshipEscrowContractTransactor // Write-only binding to the contract
	ScholarshipEscrowContractFilterer   // Log filterer for contract events
}

// ScholarshipEscrowContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ScholarshipEscrowContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScholarshipEscrowContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ScholarshipEscrowContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScholarshipEscrowContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ScholarshipEscrowContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScholarshipEscrowContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ScholarshipEscrowContractSession struct {
	Contract     *ScholarshipEscrowContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ScholarshipEscrowContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ScholarshipEscrowContractCallerSession struct {
	Contract *ScholarshipEscrowContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ScholarshipEscrowContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ScholarshipEscrowContractTransactorSession struct {
	Contract     *ScholarshipEscrowContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ScholarshipEscrowContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ScholarshipEscrowContractRaw struct {
	Contract *ScholarshipEscrowContract // Generic contract binding to access the raw methods on
}

// ScholarshipEscrowContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ScholarshipEscrowContractCallerRaw struct {
	Contract *ScholarshipEscrowContractCaller // Generic read-only contract binding to access the raw methods on
}

// ScholarshipEscrowContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ScholarshipEscrowContractTransactorRaw struct {
	Contract *ScholarshipEscrowContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewScholarshipEscrowContract creates a new instance of ScholarshipEscrowContract, bound to a specific deployed contract.
func NewScholarshipEscrowContract(address common.Address, backend bind.ContractBackend) (*ScholarshipEscrowContract, error) {
	contract, err := bindScholarshipEscrowContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ScholarshipEscrowContract{ScholarshipEscrowContractCaller: ScholarshipEscrowContractCaller{contract: contract}, ScholarshipEscrowContractTransactor: ScholarshipEscrowContractTransactor{contract: contract}, ScholarshipEscrowContractFilterer: ScholarshipEscrowContractFilterer{contract: contract}}, nil
}

// NewScholarshipEscrowContractCaller creates a new read-only instance of ScholarshipEscrowContract, bound to a specific deployed contract.
func NewScholarshipEscrowContractCaller(address common.Address, caller bind.ContractCaller) (*ScholarshipEscrowContractCaller, error) {
	contract, err := bindScholarshipEscrowContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ScholarshipEscrowContractCaller{contract: contract}, nil
}

// NewScholarshipEscrowContractTransactor creates a new write-only instance of ScholarshipEscrowContract, bound to a specific deployed contract.
func NewScholarshipEscrowContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ScholarshipEscrowContractTransactor, error) {
	contract, err := bindScholarshipEscrowContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ScholarshipEscrowContractTransactor{contract: contract}, nil
}

// NewScholarshipEscrowContractFilterer creates a new log filterer instance of ScholarshipEscrowContract, bound to a specific deployed contract.
func NewScholarshipEscrowContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ScholarshipEscrowContractFilterer, error) {
	contract, err := bindScholarshipEscrowContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ScholarshipEscrowContractFilterer{contract: contract}, nil
}

// bindScholarshipEscrowContract binds a generic wrapper to an already deployed contract.
func bindScholarshipEscrowContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ScholarshipEscrowContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ScholarshipEscrowContract *ScholarshipEscrowContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ScholarshipEscrowContract.Contract.ScholarshipEscrowContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ScholarshipEscrowContract *ScholarshipEscrowContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.ScholarshipEscrowContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ScholarshipEscrowContract *ScholarshipEscrowContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.ScholarshipEscrowContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ScholarshipEscrowContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ScholarshipEscrowContract.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) ADMINROLE() ([32]byte, error) {
	return _ScholarshipEscrowContract.Contract.ADMINROLE(&_ScholarshipEscrowContract.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCallerSession) ADMINROLE() ([32]byte, error) {
	return _ScholarshipEscrowContract.Contract.ADMINROLE(&_ScholarshipEscrowContract.CallOpts)
}

// ATTESTORROLE is a free data retrieval call binding the contract method 0x62723644.
//
// Solidity: function ATTESTOR_ROLE() view returns(bytes32)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCaller) ATTESTORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ScholarshipEscrowContract.contract.Call(opts, &out, "ATTESTOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ATTESTORROLE is a free data retrieval call binding the contract method 0x62723644.
//
// Solidity: function ATTESTOR_ROLE() view returns(bytes32)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) ATTESTORROLE() ([32]byte, error) {
	return _ScholarshipEscrowContract.Contract.ATTESTORROLE(&_ScholarshipEscrowContract.CallOpts)
}

// ATTESTORROLE is a free data retrieval call binding the contract method 0x62723644.
//
// Solidity: function ATTESTOR_ROLE() view returns(bytes32)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCallerSession) ATTESTORROLE() ([32]byte, error) {
	return _ScholarshipEscrowContract.Contract.ATTESTORROLE(&_ScholarshipEscrowContract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ScholarshipEscrowContract.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ScholarshipEscrowContract.Contract.DEFAULTADMINROLE(&_ScholarshipEscrowContract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ScholarshipEscrowContract.Contract.DEFAULTADMINROLE(&_ScholarshipEscrowContract.CallOpts)
}

// SPONSORROLE is a free data retrieval call binding the contract method 0xc2d79444.
//
// Solidity: function SPONSOR_ROLE() view returns(bytes32)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCaller) SPONSORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ScholarshipEscrowContract.contract.Call(opts, &out, "SPONSOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SPONSORROLE is a free data retrieval call binding the contract method 0xc2d79444.
//
// Solidity: function SPONSOR_ROLE() view returns(bytes32)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) SPONSORROLE() ([32]byte, error) {
	return _ScholarshipEscrowContract.Contract.SPONSORROLE(&_ScholarshipEscrowContract.CallOpts)
}

// SPONSORROLE is a free data retrieval call binding the contract method 0xc2d79444.
//
// Solidity: function SPONSOR_ROLE() view returns(bytes32)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCallerSession) SPONSORROLE() ([32]byte, error) {
	return _ScholarshipEscrowContract.Contract.SPONSORROLE(&_ScholarshipEscrowContract.CallOpts)
}

// Funds is a free data retrieval call binding the contract method 0x7b8e8895.
//
// Solidity: function funds(uint256 ) view returns(address sponsor, uint256 totalAmount, uint256 releasedAmount, uint256 trancheCount, uint256 trancheAmount, bool paused)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCaller) Funds(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Sponsor        common.Address
	TotalAmount    *big.Int
	ReleasedAmount *big.Int
	TrancheCount   *big.Int
	TrancheAmount  *big.Int
	Paused         bool
}, error) {
	var out []interface{}
	err := _ScholarshipEscrowContract.contract.Call(opts, &out, "funds", arg0)

	outstruct := new(struct {
		Sponsor        common.Address
		TotalAmount    *big.Int
		ReleasedAmount *big.Int
		TrancheCount   *big.Int
		TrancheAmount  *big.Int
		Paused         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Sponsor = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TotalAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ReleasedAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TrancheCount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TrancheAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Paused = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// Funds is a free data retrieval call binding the contract method 0x7b8e8895.
//
// Solidity: function funds(uint256 ) view returns(address sponsor, uint256 totalAmount, uint256 releasedAmount, uint256 trancheCount, uint256 trancheAmount, bool paused)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) Funds(arg0 *big.Int) (struct {
	Sponsor        common.Address
	TotalAmount    *big.Int
	ReleasedAmount *big.Int
	TrancheCount   *big.Int
	TrancheAmount  *big.Int
	Paused         bool
}, error) {
	return _ScholarshipEscrowContract.Contract.Funds(&_ScholarshipEscrowContract.CallOpts, arg0)
}

// Funds is a free data retrieval call binding the contract method 0x7b8e8895.
//
// Solidity: function funds(uint256 ) view returns(address sponsor, uint256 totalAmount, uint256 releasedAmount, uint256 trancheCount, uint256 trancheAmount, bool paused)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCallerSession) Funds(arg0 *big.Int) (struct {
	Sponsor        common.Address
	TotalAmount    *big.Int
	ReleasedAmount *big.Int
	TrancheCount   *big.Int
	TrancheAmount  *big.Int
	Paused         bool
}, error) {
	return _ScholarshipEscrowContract.Contract.Funds(&_ScholarshipEscrowContract.CallOpts, arg0)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ScholarshipEscrowContract.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ScholarshipEscrowContract.Contract.GetRoleAdmin(&_ScholarshipEscrowContract.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ScholarshipEscrowContract.Contract.GetRoleAdmin(&_ScholarshipEscrowContract.CallOpts, role)
}

// HasClaimed is a free data retrieval call binding the contract method 0xf6cf3dca.
//
// Solidity: function hasClaimed(uint256 , bytes32 , uint256 ) view returns(bool)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCaller) HasClaimed(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte, arg2 *big.Int) (bool, error) {
	var out []interface{}
	err := _ScholarshipEscrowContract.contract.Call(opts, &out, "hasClaimed", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasClaimed is a free data retrieval call binding the contract method 0xf6cf3dca.
//
// Solidity: function hasClaimed(uint256 , bytes32 , uint256 ) view returns(bool)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) HasClaimed(arg0 *big.Int, arg1 [32]byte, arg2 *big.Int) (bool, error) {
	return _ScholarshipEscrowContract.Contract.HasClaimed(&_ScholarshipEscrowContract.CallOpts, arg0, arg1, arg2)
}

// HasClaimed is a free data retrieval call binding the contract method 0xf6cf3dca.
//
// Solidity: function hasClaimed(uint256 , bytes32 , uint256 ) view returns(bool)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCallerSession) HasClaimed(arg0 *big.Int, arg1 [32]byte, arg2 *big.Int) (bool, error) {
	return _ScholarshipEscrowContract.Contract.HasClaimed(&_ScholarshipEscrowContract.CallOpts, arg0, arg1, arg2)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ScholarshipEscrowContract.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ScholarshipEscrowContract.Contract.HasRole(&_ScholarshipEscrowContract.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ScholarshipEscrowContract.Contract.HasRole(&_ScholarshipEscrowContract.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ScholarshipEscrowContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ScholarshipEscrowContract.Contract.SupportsInterface(&_ScholarshipEscrowContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ScholarshipEscrowContract.Contract.SupportsInterface(&_ScholarshipEscrowContract.CallOpts, interfaceId)
}

// TrancheRoots is a free data retrieval call binding the contract method 0x4b23c222.
//
// Solidity: function trancheRoots(uint256 , uint256 ) view returns(bytes32)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCaller) TrancheRoots(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ScholarshipEscrowContract.contract.Call(opts, &out, "trancheRoots", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TrancheRoots is a free data retrieval call binding the contract method 0x4b23c222.
//
// Solidity: function trancheRoots(uint256 , uint256 ) view returns(bytes32)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) TrancheRoots(arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	return _ScholarshipEscrowContract.Contract.TrancheRoots(&_ScholarshipEscrowContract.CallOpts, arg0, arg1)
}

// TrancheRoots is a free data retrieval call binding the contract method 0x4b23c222.
//
// Solidity: function trancheRoots(uint256 , uint256 ) view returns(bytes32)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCallerSession) TrancheRoots(arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	return _ScholarshipEscrowContract.Contract.TrancheRoots(&_ScholarshipEscrowContract.CallOpts, arg0, arg1)
}

// TrustedAttestor is a free data retrieval call binding the contract method 0xac6bffbd.
//
// Solidity: function trustedAttestor() view returns(address)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCaller) TrustedAttestor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScholarshipEscrowContract.contract.Call(opts, &out, "trustedAttestor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedAttestor is a free data retrieval call binding the contract method 0xac6bffbd.
//
// Solidity: function trustedAttestor() view returns(address)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) TrustedAttestor() (common.Address, error) {
	return _ScholarshipEscrowContract.Contract.TrustedAttestor(&_ScholarshipEscrowContract.CallOpts)
}

// TrustedAttestor is a free data retrieval call binding the contract method 0xac6bffbd.
//
// Solidity: function trustedAttestor() view returns(address)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCallerSession) TrustedAttestor() (common.Address, error) {
	return _ScholarshipEscrowContract.Contract.TrustedAttestor(&_ScholarshipEscrowContract.CallOpts)
}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCaller) UsdcToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScholarshipEscrowContract.contract.Call(opts, &out, "usdcToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) UsdcToken() (common.Address, error) {
	return _ScholarshipEscrowContract.Contract.UsdcToken(&_ScholarshipEscrowContract.CallOpts)
}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCallerSession) UsdcToken() (common.Address, error) {
	return _ScholarshipEscrowContract.Contract.UsdcToken(&_ScholarshipEscrowContract.CallOpts)
}

// ClaimTranche is a paid mutator transaction binding the contract method 0xa1595236.
//
// Solidity: function claimTranche(uint256 fundId, bytes32 studentHash, uint256 trancheIndex, address recipient, bytes32[] merkleProof) returns()
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactor) ClaimTranche(opts *bind.TransactOpts, fundId *big.Int, studentHash [32]byte, trancheIndex *big.Int, recipient common.Address, merkleProof [][32]byte) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.contract.Transact(opts, "claimTranche", fundId, studentHash, trancheIndex, recipient, merkleProof)
}

// ClaimTranche is a paid mutator transaction binding the contract method 0xa1595236.
//
// Solidity: function claimTranche(uint256 fundId, bytes32 studentHash, uint256 trancheIndex, address recipient, bytes32[] merkleProof) returns()
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) ClaimTranche(fundId *big.Int, studentHash [32]byte, trancheIndex *big.Int, recipient common.Address, merkleProof [][32]byte) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.ClaimTranche(&_ScholarshipEscrowContract.TransactOpts, fundId, studentHash, trancheIndex, recipient, merkleProof)
}

// ClaimTranche is a paid mutator transaction binding the contract method 0xa1595236.
//
// Solidity: function claimTranche(uint256 fundId, bytes32 studentHash, uint256 trancheIndex, address recipient, bytes32[] merkleProof) returns()
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactorSession) ClaimTranche(fundId *big.Int, studentHash [32]byte, trancheIndex *big.Int, recipient common.Address, merkleProof [][32]byte) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.ClaimTranche(&_ScholarshipEscrowContract.TransactOpts, fundId, studentHash, trancheIndex, recipient, merkleProof)
}

// CreateFund is a paid mutator transaction binding the contract method 0x1c02efc3.
//
// Solidity: function createFund(address sponsor, uint256 totalAmount, uint256 trancheCount, uint256 trancheAmount) returns(uint256)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactor) CreateFund(opts *bind.TransactOpts, sponsor common.Address, totalAmount *big.Int, trancheCount *big.Int, trancheAmount *big.Int) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.contract.Transact(opts, "createFund", sponsor, totalAmount, trancheCount, trancheAmount)
}

// CreateFund is a paid mutator transaction binding the contract method 0x1c02efc3.
//
// Solidity: function createFund(address sponsor, uint256 totalAmount, uint256 trancheCount, uint256 trancheAmount) returns(uint256)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) CreateFund(sponsor common.Address, totalAmount *big.Int, trancheCount *big.Int, trancheAmount *big.Int) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.CreateFund(&_ScholarshipEscrowContract.TransactOpts, sponsor, totalAmount, trancheCount, trancheAmount)
}

// CreateFund is a paid mutator transaction binding the contract method 0x1c02efc3.
//
// Solidity: function createFund(address sponsor, uint256 totalAmount, uint256 trancheCount, uint256 trancheAmount) returns(uint256)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactorSession) CreateFund(sponsor common.Address, totalAmount *big.Int, trancheCount *big.Int, trancheAmount *big.Int) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.CreateFund(&_ScholarshipEscrowContract.TransactOpts, sponsor, totalAmount, trancheCount, trancheAmount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.GrantRole(&_ScholarshipEscrowContract.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.GrantRole(&_ScholarshipEscrowContract.TransactOpts, role, account)
}

// PauseFund is a paid mutator transaction binding the contract method 0x02bc5803.
//
// Solidity: function pauseFund(uint256 fundId) returns()
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactor) PauseFund(opts *bind.TransactOpts, fundId *big.Int) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.contract.Transact(opts, "pauseFund", fundId)
}

// PauseFund is a paid mutator transaction binding the contract method 0x02bc5803.
//
// Solidity: function pauseFund(uint256 fundId) returns()
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) PauseFund(fundId *big.Int) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.PauseFund(&_ScholarshipEscrowContract.TransactOpts, fundId)
}

// PauseFund is a paid mutator transaction binding the contract method 0x02bc5803.
//
// Solidity: function pauseFund(uint256 fundId) returns()
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactorSession) PauseFund(fundId *big.Int) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.PauseFund(&_ScholarshipEscrowContract.TransactOpts, fundId)
}

// PublishTrancheRoot is a paid mutator transaction binding the contract method 0x5e3e05fa.
//
// Solidity: function publishTrancheRoot(uint256 fundId, uint256 trancheIndex, bytes32 merkleRoot) returns()
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactor) PublishTrancheRoot(opts *bind.TransactOpts, fundId *big.Int, trancheIndex *big.Int, merkleRoot [32]byte) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.contract.Transact(opts, "publishTrancheRoot", fundId, trancheIndex, merkleRoot)
}

// PublishTrancheRoot is a paid mutator transaction binding the contract method 0x5e3e05fa.
//
// Solidity: function publishTrancheRoot(uint256 fundId, uint256 trancheIndex, bytes32 merkleRoot) returns()
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) PublishTrancheRoot(fundId *big.Int, trancheIndex *big.Int, merkleRoot [32]byte) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.PublishTrancheRoot(&_ScholarshipEscrowContract.TransactOpts, fundId, trancheIndex, merkleRoot)
}

// PublishTrancheRoot is a paid mutator transaction binding the contract method 0x5e3e05fa.
//
// Solidity: function publishTrancheRoot(uint256 fundId, uint256 trancheIndex, bytes32 merkleRoot) returns()
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactorSession) PublishTrancheRoot(fundId *big.Int, trancheIndex *big.Int, merkleRoot [32]byte) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.PublishTrancheRoot(&_ScholarshipEscrowContract.TransactOpts, fundId, trancheIndex, merkleRoot)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.RenounceRole(&_ScholarshipEscrowContract.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.RenounceRole(&_ScholarshipEscrowContract.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.RevokeRole(&_ScholarshipEscrowContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.RevokeRole(&_ScholarshipEscrowContract.TransactOpts, role, account)
}

// SetTrustedAttestor is a paid mutator transaction binding the contract method 0xe3d477d9.
//
// Solidity: function setTrustedAttestor(address newAttestor) returns()
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactor) SetTrustedAttestor(opts *bind.TransactOpts, newAttestor common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.contract.Transact(opts, "setTrustedAttestor", newAttestor)
}

// SetTrustedAttestor is a paid mutator transaction binding the contract method 0xe3d477d9.
//
// Solidity: function setTrustedAttestor(address newAttestor) returns()
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) SetTrustedAttestor(newAttestor common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.SetTrustedAttestor(&_ScholarshipEscrowContract.TransactOpts, newAttestor)
}

// SetTrustedAttestor is a paid mutator transaction binding the contract method 0xe3d477d9.
//
// Solidity: function setTrustedAttestor(address newAttestor) returns()
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactorSession) SetTrustedAttestor(newAttestor common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.SetTrustedAttestor(&_ScholarshipEscrowContract.TransactOpts, newAttestor)
}

// UnpauseFund is a paid mutator transaction binding the contract method 0x22263434.
//
// Solidity: function unpauseFund(uint256 fundId) returns()
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactor) UnpauseFund(opts *bind.TransactOpts, fundId *big.Int) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.contract.Transact(opts, "unpauseFund", fundId)
}

// UnpauseFund is a paid mutator transaction binding the contract method 0x22263434.
//
// Solidity: function unpauseFund(uint256 fundId) returns()
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) UnpauseFund(fundId *big.Int) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.UnpauseFund(&_ScholarshipEscrowContract.TransactOpts, fundId)
}

// UnpauseFund is a paid mutator transaction binding the contract method 0x22263434.
//
// Solidity: function unpauseFund(uint256 fundId) returns()
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactorSession) UnpauseFund(fundId *big.Int) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.UnpauseFund(&_ScholarshipEscrowContract.TransactOpts, fundId)
}

// ScholarshipEscrowContractFundCreatedIterator is returned from FilterFundCreated and is used to iterate over the raw logs and unpacked data for FundCreated events raised by the ScholarshipEscrowContract contract.
type ScholarshipEscrowContractFundCreatedIterator struct {
	Event *ScholarshipEscrowContractFundCreated // Event containing the contract specifics and raw log

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
func (it *ScholarshipEscrowContractFundCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScholarshipEscrowContractFundCreated)
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
		it.Event = new(ScholarshipEscrowContractFundCreated)
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
func (it *ScholarshipEscrowContractFundCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScholarshipEscrowContractFundCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScholarshipEscrowContractFundCreated represents a FundCreated event raised by the ScholarshipEscrowContract contract.
type ScholarshipEscrowContractFundCreated struct {
	FundId        *big.Int
	Sponsor       common.Address
	TotalAmount   *big.Int
	TrancheCount  *big.Int
	TrancheAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFundCreated is a free log retrieval operation binding the contract event 0x77f65105cac199e98107aec085f9f5785ec5a22994dc9ed45224d96c9cecce0a.
//
// Solidity: event FundCreated(uint256 indexed fundId, address indexed sponsor, uint256 totalAmount, uint256 trancheCount, uint256 trancheAmount)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) FilterFundCreated(opts *bind.FilterOpts, fundId []*big.Int, sponsor []common.Address) (*ScholarshipEscrowContractFundCreatedIterator, error) {

	var fundIdRule []interface{}
	for _, fundIdItem := range fundId {
		fundIdRule = append(fundIdRule, fundIdItem)
	}
	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}

	logs, sub, err := _ScholarshipEscrowContract.contract.FilterLogs(opts, "FundCreated", fundIdRule, sponsorRule)
	if err != nil {
		return nil, err
	}
	return &ScholarshipEscrowContractFundCreatedIterator{contract: _ScholarshipEscrowContract.contract, event: "FundCreated", logs: logs, sub: sub}, nil
}

// WatchFundCreated is a free log subscription operation binding the contract event 0x77f65105cac199e98107aec085f9f5785ec5a22994dc9ed45224d96c9cecce0a.
//
// Solidity: event FundCreated(uint256 indexed fundId, address indexed sponsor, uint256 totalAmount, uint256 trancheCount, uint256 trancheAmount)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) WatchFundCreated(opts *bind.WatchOpts, sink chan<- *ScholarshipEscrowContractFundCreated, fundId []*big.Int, sponsor []common.Address) (event.Subscription, error) {

	var fundIdRule []interface{}
	for _, fundIdItem := range fundId {
		fundIdRule = append(fundIdRule, fundIdItem)
	}
	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}

	logs, sub, err := _ScholarshipEscrowContract.contract.WatchLogs(opts, "FundCreated", fundIdRule, sponsorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScholarshipEscrowContractFundCreated)
				if err := _ScholarshipEscrowContract.contract.UnpackLog(event, "FundCreated", log); err != nil {
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

// ParseFundCreated is a log parse operation binding the contract event 0x77f65105cac199e98107aec085f9f5785ec5a22994dc9ed45224d96c9cecce0a.
//
// Solidity: event FundCreated(uint256 indexed fundId, address indexed sponsor, uint256 totalAmount, uint256 trancheCount, uint256 trancheAmount)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) ParseFundCreated(log types.Log) (*ScholarshipEscrowContractFundCreated, error) {
	event := new(ScholarshipEscrowContractFundCreated)
	if err := _ScholarshipEscrowContract.contract.UnpackLog(event, "FundCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScholarshipEscrowContractMerkleRootPublishedIterator is returned from FilterMerkleRootPublished and is used to iterate over the raw logs and unpacked data for MerkleRootPublished events raised by the ScholarshipEscrowContract contract.
type ScholarshipEscrowContractMerkleRootPublishedIterator struct {
	Event *ScholarshipEscrowContractMerkleRootPublished // Event containing the contract specifics and raw log

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
func (it *ScholarshipEscrowContractMerkleRootPublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScholarshipEscrowContractMerkleRootPublished)
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
		it.Event = new(ScholarshipEscrowContractMerkleRootPublished)
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
func (it *ScholarshipEscrowContractMerkleRootPublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScholarshipEscrowContractMerkleRootPublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScholarshipEscrowContractMerkleRootPublished represents a MerkleRootPublished event raised by the ScholarshipEscrowContract contract.
type ScholarshipEscrowContractMerkleRootPublished struct {
	FundId       *big.Int
	TrancheIndex *big.Int
	MerkleRoot   [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMerkleRootPublished is a free log retrieval operation binding the contract event 0xe649294352ae0b15ac227a22fb744ab148b2c5e053d26fd38e97890ad47dbd8f.
//
// Solidity: event MerkleRootPublished(uint256 indexed fundId, uint256 indexed trancheIndex, bytes32 merkleRoot)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) FilterMerkleRootPublished(opts *bind.FilterOpts, fundId []*big.Int, trancheIndex []*big.Int) (*ScholarshipEscrowContractMerkleRootPublishedIterator, error) {

	var fundIdRule []interface{}
	for _, fundIdItem := range fundId {
		fundIdRule = append(fundIdRule, fundIdItem)
	}
	var trancheIndexRule []interface{}
	for _, trancheIndexItem := range trancheIndex {
		trancheIndexRule = append(trancheIndexRule, trancheIndexItem)
	}

	logs, sub, err := _ScholarshipEscrowContract.contract.FilterLogs(opts, "MerkleRootPublished", fundIdRule, trancheIndexRule)
	if err != nil {
		return nil, err
	}
	return &ScholarshipEscrowContractMerkleRootPublishedIterator{contract: _ScholarshipEscrowContract.contract, event: "MerkleRootPublished", logs: logs, sub: sub}, nil
}

// WatchMerkleRootPublished is a free log subscription operation binding the contract event 0xe649294352ae0b15ac227a22fb744ab148b2c5e053d26fd38e97890ad47dbd8f.
//
// Solidity: event MerkleRootPublished(uint256 indexed fundId, uint256 indexed trancheIndex, bytes32 merkleRoot)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) WatchMerkleRootPublished(opts *bind.WatchOpts, sink chan<- *ScholarshipEscrowContractMerkleRootPublished, fundId []*big.Int, trancheIndex []*big.Int) (event.Subscription, error) {

	var fundIdRule []interface{}
	for _, fundIdItem := range fundId {
		fundIdRule = append(fundIdRule, fundIdItem)
	}
	var trancheIndexRule []interface{}
	for _, trancheIndexItem := range trancheIndex {
		trancheIndexRule = append(trancheIndexRule, trancheIndexItem)
	}

	logs, sub, err := _ScholarshipEscrowContract.contract.WatchLogs(opts, "MerkleRootPublished", fundIdRule, trancheIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScholarshipEscrowContractMerkleRootPublished)
				if err := _ScholarshipEscrowContract.contract.UnpackLog(event, "MerkleRootPublished", log); err != nil {
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

// ParseMerkleRootPublished is a log parse operation binding the contract event 0xe649294352ae0b15ac227a22fb744ab148b2c5e053d26fd38e97890ad47dbd8f.
//
// Solidity: event MerkleRootPublished(uint256 indexed fundId, uint256 indexed trancheIndex, bytes32 merkleRoot)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) ParseMerkleRootPublished(log types.Log) (*ScholarshipEscrowContractMerkleRootPublished, error) {
	event := new(ScholarshipEscrowContractMerkleRootPublished)
	if err := _ScholarshipEscrowContract.contract.UnpackLog(event, "MerkleRootPublished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScholarshipEscrowContractRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ScholarshipEscrowContract contract.
type ScholarshipEscrowContractRoleAdminChangedIterator struct {
	Event *ScholarshipEscrowContractRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ScholarshipEscrowContractRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScholarshipEscrowContractRoleAdminChanged)
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
		it.Event = new(ScholarshipEscrowContractRoleAdminChanged)
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
func (it *ScholarshipEscrowContractRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScholarshipEscrowContractRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScholarshipEscrowContractRoleAdminChanged represents a RoleAdminChanged event raised by the ScholarshipEscrowContract contract.
type ScholarshipEscrowContractRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ScholarshipEscrowContractRoleAdminChangedIterator, error) {

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

	logs, sub, err := _ScholarshipEscrowContract.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ScholarshipEscrowContractRoleAdminChangedIterator{contract: _ScholarshipEscrowContract.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ScholarshipEscrowContractRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ScholarshipEscrowContract.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScholarshipEscrowContractRoleAdminChanged)
				if err := _ScholarshipEscrowContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) ParseRoleAdminChanged(log types.Log) (*ScholarshipEscrowContractRoleAdminChanged, error) {
	event := new(ScholarshipEscrowContractRoleAdminChanged)
	if err := _ScholarshipEscrowContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScholarshipEscrowContractRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ScholarshipEscrowContract contract.
type ScholarshipEscrowContractRoleGrantedIterator struct {
	Event *ScholarshipEscrowContractRoleGranted // Event containing the contract specifics and raw log

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
func (it *ScholarshipEscrowContractRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScholarshipEscrowContractRoleGranted)
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
		it.Event = new(ScholarshipEscrowContractRoleGranted)
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
func (it *ScholarshipEscrowContractRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScholarshipEscrowContractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScholarshipEscrowContractRoleGranted represents a RoleGranted event raised by the ScholarshipEscrowContract contract.
type ScholarshipEscrowContractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ScholarshipEscrowContractRoleGrantedIterator, error) {

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

	logs, sub, err := _ScholarshipEscrowContract.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ScholarshipEscrowContractRoleGrantedIterator{contract: _ScholarshipEscrowContract.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ScholarshipEscrowContractRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ScholarshipEscrowContract.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScholarshipEscrowContractRoleGranted)
				if err := _ScholarshipEscrowContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) ParseRoleGranted(log types.Log) (*ScholarshipEscrowContractRoleGranted, error) {
	event := new(ScholarshipEscrowContractRoleGranted)
	if err := _ScholarshipEscrowContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScholarshipEscrowContractRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ScholarshipEscrowContract contract.
type ScholarshipEscrowContractRoleRevokedIterator struct {
	Event *ScholarshipEscrowContractRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ScholarshipEscrowContractRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScholarshipEscrowContractRoleRevoked)
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
		it.Event = new(ScholarshipEscrowContractRoleRevoked)
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
func (it *ScholarshipEscrowContractRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScholarshipEscrowContractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScholarshipEscrowContractRoleRevoked represents a RoleRevoked event raised by the ScholarshipEscrowContract contract.
type ScholarshipEscrowContractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ScholarshipEscrowContractRoleRevokedIterator, error) {

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

	logs, sub, err := _ScholarshipEscrowContract.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ScholarshipEscrowContractRoleRevokedIterator{contract: _ScholarshipEscrowContract.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ScholarshipEscrowContractRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ScholarshipEscrowContract.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScholarshipEscrowContractRoleRevoked)
				if err := _ScholarshipEscrowContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) ParseRoleRevoked(log types.Log) (*ScholarshipEscrowContractRoleRevoked, error) {
	event := new(ScholarshipEscrowContractRoleRevoked)
	if err := _ScholarshipEscrowContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScholarshipEscrowContractTrancheClaimedIterator is returned from FilterTrancheClaimed and is used to iterate over the raw logs and unpacked data for TrancheClaimed events raised by the ScholarshipEscrowContract contract.
type ScholarshipEscrowContractTrancheClaimedIterator struct {
	Event *ScholarshipEscrowContractTrancheClaimed // Event containing the contract specifics and raw log

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
func (it *ScholarshipEscrowContractTrancheClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScholarshipEscrowContractTrancheClaimed)
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
		it.Event = new(ScholarshipEscrowContractTrancheClaimed)
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
func (it *ScholarshipEscrowContractTrancheClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScholarshipEscrowContractTrancheClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScholarshipEscrowContractTrancheClaimed represents a TrancheClaimed event raised by the ScholarshipEscrowContract contract.
type ScholarshipEscrowContractTrancheClaimed struct {
	FundId       *big.Int
	StudentHash  [32]byte
	TrancheIndex *big.Int
	Amount       *big.Int
	Recipient    common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTrancheClaimed is a free log retrieval operation binding the contract event 0x7beb97e8a642fc5efdb2a6eafecdd3f6dc96464d6d2098ae9d32061c79487455.
//
// Solidity: event TrancheClaimed(uint256 indexed fundId, bytes32 indexed studentHash, uint256 trancheIndex, uint256 amount, address recipient)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) FilterTrancheClaimed(opts *bind.FilterOpts, fundId []*big.Int, studentHash [][32]byte) (*ScholarshipEscrowContractTrancheClaimedIterator, error) {

	var fundIdRule []interface{}
	for _, fundIdItem := range fundId {
		fundIdRule = append(fundIdRule, fundIdItem)
	}
	var studentHashRule []interface{}
	for _, studentHashItem := range studentHash {
		studentHashRule = append(studentHashRule, studentHashItem)
	}

	logs, sub, err := _ScholarshipEscrowContract.contract.FilterLogs(opts, "TrancheClaimed", fundIdRule, studentHashRule)
	if err != nil {
		return nil, err
	}
	return &ScholarshipEscrowContractTrancheClaimedIterator{contract: _ScholarshipEscrowContract.contract, event: "TrancheClaimed", logs: logs, sub: sub}, nil
}

// WatchTrancheClaimed is a free log subscription operation binding the contract event 0x7beb97e8a642fc5efdb2a6eafecdd3f6dc96464d6d2098ae9d32061c79487455.
//
// Solidity: event TrancheClaimed(uint256 indexed fundId, bytes32 indexed studentHash, uint256 trancheIndex, uint256 amount, address recipient)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) WatchTrancheClaimed(opts *bind.WatchOpts, sink chan<- *ScholarshipEscrowContractTrancheClaimed, fundId []*big.Int, studentHash [][32]byte) (event.Subscription, error) {

	var fundIdRule []interface{}
	for _, fundIdItem := range fundId {
		fundIdRule = append(fundIdRule, fundIdItem)
	}
	var studentHashRule []interface{}
	for _, studentHashItem := range studentHash {
		studentHashRule = append(studentHashRule, studentHashItem)
	}

	logs, sub, err := _ScholarshipEscrowContract.contract.WatchLogs(opts, "TrancheClaimed", fundIdRule, studentHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScholarshipEscrowContractTrancheClaimed)
				if err := _ScholarshipEscrowContract.contract.UnpackLog(event, "TrancheClaimed", log); err != nil {
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

// ParseTrancheClaimed is a log parse operation binding the contract event 0x7beb97e8a642fc5efdb2a6eafecdd3f6dc96464d6d2098ae9d32061c79487455.
//
// Solidity: event TrancheClaimed(uint256 indexed fundId, bytes32 indexed studentHash, uint256 trancheIndex, uint256 amount, address recipient)
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) ParseTrancheClaimed(log types.Log) (*ScholarshipEscrowContractTrancheClaimed, error) {
	event := new(ScholarshipEscrowContractTrancheClaimed)
	if err := _ScholarshipEscrowContract.contract.UnpackLog(event, "TrancheClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
