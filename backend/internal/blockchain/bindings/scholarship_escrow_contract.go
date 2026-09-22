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
var ScholarshipEscrowContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ATTESTOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SPONSOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimCrossChain\",\"inputs\":[{\"name\":\"fundId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"trancheIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"studentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"merkleProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"dstChainId\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"dstAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"claimTranche\",\"inputs\":[{\"name\":\"fundId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"studentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"trancheIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"merkleProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"clawbackFund\",\"inputs\":[{\"name\":\"fundId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createFund\",\"inputs\":[{\"name\":\"sponsor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"totalAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"trancheCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"trancheAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"funds\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"sponsor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"totalAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"releasedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"trancheCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"trancheAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paused\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasClaimed\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"attestor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdcToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lzEndpoint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseFund\",\"inputs\":[{\"name\":\"fundId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"publishTrancheRoot\",\"inputs\":[{\"name\":\"fundId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"trancheIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLzEndpoint\",\"inputs\":[{\"name\":\"_lzEndpoint\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTrustedAttestor\",\"inputs\":[{\"name\":\"newAttestor\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"trancheRoots\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"trustedAttestor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpauseFund\",\"inputs\":[{\"name\":\"fundId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"usdcToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"FundClawedBack\",\"inputs\":[{\"name\":\"fundId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FundCreated\",\"inputs\":[{\"name\":\"fundId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"sponsor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"totalAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"trancheCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"trancheAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MerkleRootPublished\",\"inputs\":[{\"name\":\"fundId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"trancheIndex\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TrancheClaimed\",\"inputs\":[{\"name\":\"fundId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"studentHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"trancheIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}
var ScholarshipEscrowContractABI = ScholarshipEscrowContractMetaData.ABI
type ScholarshipEscrowContract struct {
	ScholarshipEscrowContractCaller     
	ScholarshipEscrowContractTransactor 
	ScholarshipEscrowContractFilterer   
}
type ScholarshipEscrowContractCaller struct {
	contract *bind.BoundContract 
}
type ScholarshipEscrowContractTransactor struct {
	contract *bind.BoundContract 
}
type ScholarshipEscrowContractFilterer struct {
	contract *bind.BoundContract 
}
type ScholarshipEscrowContractSession struct {
	Contract     *ScholarshipEscrowContract 
	CallOpts     bind.CallOpts              
	TransactOpts bind.TransactOpts          
}
type ScholarshipEscrowContractCallerSession struct {
	Contract *ScholarshipEscrowContractCaller 
	CallOpts bind.CallOpts                    
}
type ScholarshipEscrowContractTransactorSession struct {
	Contract     *ScholarshipEscrowContractTransactor 
	TransactOpts bind.TransactOpts                    
}
type ScholarshipEscrowContractRaw struct {
	Contract *ScholarshipEscrowContract 
}
type ScholarshipEscrowContractCallerRaw struct {
	Contract *ScholarshipEscrowContractCaller 
}
type ScholarshipEscrowContractTransactorRaw struct {
	Contract *ScholarshipEscrowContractTransactor 
}
func NewScholarshipEscrowContract(address common.Address, backend bind.ContractBackend) (*ScholarshipEscrowContract, error) {
	contract, err := bindScholarshipEscrowContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ScholarshipEscrowContract{ScholarshipEscrowContractCaller: ScholarshipEscrowContractCaller{contract: contract}, ScholarshipEscrowContractTransactor: ScholarshipEscrowContractTransactor{contract: contract}, ScholarshipEscrowContractFilterer: ScholarshipEscrowContractFilterer{contract: contract}}, nil
}
func NewScholarshipEscrowContractCaller(address common.Address, caller bind.ContractCaller) (*ScholarshipEscrowContractCaller, error) {
	contract, err := bindScholarshipEscrowContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ScholarshipEscrowContractCaller{contract: contract}, nil
}
func NewScholarshipEscrowContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ScholarshipEscrowContractTransactor, error) {
	contract, err := bindScholarshipEscrowContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ScholarshipEscrowContractTransactor{contract: contract}, nil
}
func NewScholarshipEscrowContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ScholarshipEscrowContractFilterer, error) {
	contract, err := bindScholarshipEscrowContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ScholarshipEscrowContractFilterer{contract: contract}, nil
}
func bindScholarshipEscrowContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ScholarshipEscrowContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ScholarshipEscrowContract.Contract.ScholarshipEscrowContractCaller.contract.Call(opts, result, method, params...)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.ScholarshipEscrowContractTransactor.contract.Transfer(opts)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.ScholarshipEscrowContractTransactor.contract.Transact(opts, method, params...)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ScholarshipEscrowContract.Contract.contract.Call(opts, result, method, params...)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.contract.Transfer(opts)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.contract.Transact(opts, method, params...)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ScholarshipEscrowContract.contract.Call(opts, &out, "ADMIN_ROLE")
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) ADMINROLE() ([32]byte, error) {
	return _ScholarshipEscrowContract.Contract.ADMINROLE(&_ScholarshipEscrowContract.CallOpts)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCallerSession) ADMINROLE() ([32]byte, error) {
	return _ScholarshipEscrowContract.Contract.ADMINROLE(&_ScholarshipEscrowContract.CallOpts)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCaller) ATTESTORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ScholarshipEscrowContract.contract.Call(opts, &out, "ATTESTOR_ROLE")
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) ATTESTORROLE() ([32]byte, error) {
	return _ScholarshipEscrowContract.Contract.ATTESTORROLE(&_ScholarshipEscrowContract.CallOpts)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCallerSession) ATTESTORROLE() ([32]byte, error) {
	return _ScholarshipEscrowContract.Contract.ATTESTORROLE(&_ScholarshipEscrowContract.CallOpts)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ScholarshipEscrowContract.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ScholarshipEscrowContract.Contract.DEFAULTADMINROLE(&_ScholarshipEscrowContract.CallOpts)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ScholarshipEscrowContract.Contract.DEFAULTADMINROLE(&_ScholarshipEscrowContract.CallOpts)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCaller) SPONSORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ScholarshipEscrowContract.contract.Call(opts, &out, "SPONSOR_ROLE")
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) SPONSORROLE() ([32]byte, error) {
	return _ScholarshipEscrowContract.Contract.SPONSORROLE(&_ScholarshipEscrowContract.CallOpts)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCallerSession) SPONSORROLE() ([32]byte, error) {
	return _ScholarshipEscrowContract.Contract.SPONSORROLE(&_ScholarshipEscrowContract.CallOpts)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ScholarshipEscrowContract.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ScholarshipEscrowContract.Contract.UPGRADEINTERFACEVERSION(&_ScholarshipEscrowContract.CallOpts)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ScholarshipEscrowContract.Contract.UPGRADEINTERFACEVERSION(&_ScholarshipEscrowContract.CallOpts)
}
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
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ScholarshipEscrowContract.contract.Call(opts, &out, "getRoleAdmin", role)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ScholarshipEscrowContract.Contract.GetRoleAdmin(&_ScholarshipEscrowContract.CallOpts, role)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ScholarshipEscrowContract.Contract.GetRoleAdmin(&_ScholarshipEscrowContract.CallOpts, role)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCaller) HasClaimed(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte, arg2 *big.Int) (bool, error) {
	var out []interface{}
	err := _ScholarshipEscrowContract.contract.Call(opts, &out, "hasClaimed", arg0, arg1, arg2)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) HasClaimed(arg0 *big.Int, arg1 [32]byte, arg2 *big.Int) (bool, error) {
	return _ScholarshipEscrowContract.Contract.HasClaimed(&_ScholarshipEscrowContract.CallOpts, arg0, arg1, arg2)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCallerSession) HasClaimed(arg0 *big.Int, arg1 [32]byte, arg2 *big.Int) (bool, error) {
	return _ScholarshipEscrowContract.Contract.HasClaimed(&_ScholarshipEscrowContract.CallOpts, arg0, arg1, arg2)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ScholarshipEscrowContract.contract.Call(opts, &out, "hasRole", role, account)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ScholarshipEscrowContract.Contract.HasRole(&_ScholarshipEscrowContract.CallOpts, role, account)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ScholarshipEscrowContract.Contract.HasRole(&_ScholarshipEscrowContract.CallOpts, role, account)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCaller) LzEndpoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScholarshipEscrowContract.contract.Call(opts, &out, "lzEndpoint")
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) LzEndpoint() (common.Address, error) {
	return _ScholarshipEscrowContract.Contract.LzEndpoint(&_ScholarshipEscrowContract.CallOpts)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCallerSession) LzEndpoint() (common.Address, error) {
	return _ScholarshipEscrowContract.Contract.LzEndpoint(&_ScholarshipEscrowContract.CallOpts)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ScholarshipEscrowContract.contract.Call(opts, &out, "proxiableUUID")
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) ProxiableUUID() ([32]byte, error) {
	return _ScholarshipEscrowContract.Contract.ProxiableUUID(&_ScholarshipEscrowContract.CallOpts)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ScholarshipEscrowContract.Contract.ProxiableUUID(&_ScholarshipEscrowContract.CallOpts)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ScholarshipEscrowContract.contract.Call(opts, &out, "supportsInterface", interfaceId)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ScholarshipEscrowContract.Contract.SupportsInterface(&_ScholarshipEscrowContract.CallOpts, interfaceId)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ScholarshipEscrowContract.Contract.SupportsInterface(&_ScholarshipEscrowContract.CallOpts, interfaceId)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCaller) TrancheRoots(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ScholarshipEscrowContract.contract.Call(opts, &out, "trancheRoots", arg0, arg1)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) TrancheRoots(arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	return _ScholarshipEscrowContract.Contract.TrancheRoots(&_ScholarshipEscrowContract.CallOpts, arg0, arg1)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCallerSession) TrancheRoots(arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	return _ScholarshipEscrowContract.Contract.TrancheRoots(&_ScholarshipEscrowContract.CallOpts, arg0, arg1)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCaller) TrustedAttestor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScholarshipEscrowContract.contract.Call(opts, &out, "trustedAttestor")
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) TrustedAttestor() (common.Address, error) {
	return _ScholarshipEscrowContract.Contract.TrustedAttestor(&_ScholarshipEscrowContract.CallOpts)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCallerSession) TrustedAttestor() (common.Address, error) {
	return _ScholarshipEscrowContract.Contract.TrustedAttestor(&_ScholarshipEscrowContract.CallOpts)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCaller) UsdcToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScholarshipEscrowContract.contract.Call(opts, &out, "usdcToken")
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) UsdcToken() (common.Address, error) {
	return _ScholarshipEscrowContract.Contract.UsdcToken(&_ScholarshipEscrowContract.CallOpts)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractCallerSession) UsdcToken() (common.Address, error) {
	return _ScholarshipEscrowContract.Contract.UsdcToken(&_ScholarshipEscrowContract.CallOpts)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactor) ClaimCrossChain(opts *bind.TransactOpts, fundId *big.Int, trancheIndex *big.Int, studentHash [32]byte, amount *big.Int, merkleProof [][32]byte, dstChainId uint16, dstAddress []byte) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.contract.Transact(opts, "claimCrossChain", fundId, trancheIndex, studentHash, amount, merkleProof, dstChainId, dstAddress)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) ClaimCrossChain(fundId *big.Int, trancheIndex *big.Int, studentHash [32]byte, amount *big.Int, merkleProof [][32]byte, dstChainId uint16, dstAddress []byte) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.ClaimCrossChain(&_ScholarshipEscrowContract.TransactOpts, fundId, trancheIndex, studentHash, amount, merkleProof, dstChainId, dstAddress)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactorSession) ClaimCrossChain(fundId *big.Int, trancheIndex *big.Int, studentHash [32]byte, amount *big.Int, merkleProof [][32]byte, dstChainId uint16, dstAddress []byte) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.ClaimCrossChain(&_ScholarshipEscrowContract.TransactOpts, fundId, trancheIndex, studentHash, amount, merkleProof, dstChainId, dstAddress)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactor) ClaimTranche(opts *bind.TransactOpts, fundId *big.Int, studentHash [32]byte, trancheIndex *big.Int, recipient common.Address, amount *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.contract.Transact(opts, "claimTranche", fundId, studentHash, trancheIndex, recipient, amount, merkleProof)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) ClaimTranche(fundId *big.Int, studentHash [32]byte, trancheIndex *big.Int, recipient common.Address, amount *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.ClaimTranche(&_ScholarshipEscrowContract.TransactOpts, fundId, studentHash, trancheIndex, recipient, amount, merkleProof)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactorSession) ClaimTranche(fundId *big.Int, studentHash [32]byte, trancheIndex *big.Int, recipient common.Address, amount *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.ClaimTranche(&_ScholarshipEscrowContract.TransactOpts, fundId, studentHash, trancheIndex, recipient, amount, merkleProof)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactor) ClawbackFund(opts *bind.TransactOpts, fundId *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.contract.Transact(opts, "clawbackFund", fundId, recipient)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) ClawbackFund(fundId *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.ClawbackFund(&_ScholarshipEscrowContract.TransactOpts, fundId, recipient)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactorSession) ClawbackFund(fundId *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.ClawbackFund(&_ScholarshipEscrowContract.TransactOpts, fundId, recipient)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactor) CreateFund(opts *bind.TransactOpts, sponsor common.Address, totalAmount *big.Int, trancheCount *big.Int, trancheAmount *big.Int) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.contract.Transact(opts, "createFund", sponsor, totalAmount, trancheCount, trancheAmount)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) CreateFund(sponsor common.Address, totalAmount *big.Int, trancheCount *big.Int, trancheAmount *big.Int) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.CreateFund(&_ScholarshipEscrowContract.TransactOpts, sponsor, totalAmount, trancheCount, trancheAmount)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactorSession) CreateFund(sponsor common.Address, totalAmount *big.Int, trancheCount *big.Int, trancheAmount *big.Int) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.CreateFund(&_ScholarshipEscrowContract.TransactOpts, sponsor, totalAmount, trancheCount, trancheAmount)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.contract.Transact(opts, "grantRole", role, account)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.GrantRole(&_ScholarshipEscrowContract.TransactOpts, role, account)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.GrantRole(&_ScholarshipEscrowContract.TransactOpts, role, account)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactor) Initialize(opts *bind.TransactOpts, admin common.Address, attestor common.Address, _usdcToken common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.contract.Transact(opts, "initialize", admin, attestor, _usdcToken)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) Initialize(admin common.Address, attestor common.Address, _usdcToken common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.Initialize(&_ScholarshipEscrowContract.TransactOpts, admin, attestor, _usdcToken)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactorSession) Initialize(admin common.Address, attestor common.Address, _usdcToken common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.Initialize(&_ScholarshipEscrowContract.TransactOpts, admin, attestor, _usdcToken)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactor) PauseFund(opts *bind.TransactOpts, fundId *big.Int) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.contract.Transact(opts, "pauseFund", fundId)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) PauseFund(fundId *big.Int) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.PauseFund(&_ScholarshipEscrowContract.TransactOpts, fundId)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactorSession) PauseFund(fundId *big.Int) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.PauseFund(&_ScholarshipEscrowContract.TransactOpts, fundId)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactor) PublishTrancheRoot(opts *bind.TransactOpts, fundId *big.Int, trancheIndex *big.Int, merkleRoot [32]byte) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.contract.Transact(opts, "publishTrancheRoot", fundId, trancheIndex, merkleRoot)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) PublishTrancheRoot(fundId *big.Int, trancheIndex *big.Int, merkleRoot [32]byte) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.PublishTrancheRoot(&_ScholarshipEscrowContract.TransactOpts, fundId, trancheIndex, merkleRoot)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactorSession) PublishTrancheRoot(fundId *big.Int, trancheIndex *big.Int, merkleRoot [32]byte) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.PublishTrancheRoot(&_ScholarshipEscrowContract.TransactOpts, fundId, trancheIndex, merkleRoot)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.RenounceRole(&_ScholarshipEscrowContract.TransactOpts, role, callerConfirmation)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.RenounceRole(&_ScholarshipEscrowContract.TransactOpts, role, callerConfirmation)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.contract.Transact(opts, "revokeRole", role, account)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.RevokeRole(&_ScholarshipEscrowContract.TransactOpts, role, account)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.RevokeRole(&_ScholarshipEscrowContract.TransactOpts, role, account)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactor) SetLzEndpoint(opts *bind.TransactOpts, _lzEndpoint common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.contract.Transact(opts, "setLzEndpoint", _lzEndpoint)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) SetLzEndpoint(_lzEndpoint common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.SetLzEndpoint(&_ScholarshipEscrowContract.TransactOpts, _lzEndpoint)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactorSession) SetLzEndpoint(_lzEndpoint common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.SetLzEndpoint(&_ScholarshipEscrowContract.TransactOpts, _lzEndpoint)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactor) SetTrustedAttestor(opts *bind.TransactOpts, newAttestor common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.contract.Transact(opts, "setTrustedAttestor", newAttestor)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) SetTrustedAttestor(newAttestor common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.SetTrustedAttestor(&_ScholarshipEscrowContract.TransactOpts, newAttestor)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactorSession) SetTrustedAttestor(newAttestor common.Address) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.SetTrustedAttestor(&_ScholarshipEscrowContract.TransactOpts, newAttestor)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactor) UnpauseFund(opts *bind.TransactOpts, fundId *big.Int) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.contract.Transact(opts, "unpauseFund", fundId)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) UnpauseFund(fundId *big.Int) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.UnpauseFund(&_ScholarshipEscrowContract.TransactOpts, fundId)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactorSession) UnpauseFund(fundId *big.Int) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.UnpauseFund(&_ScholarshipEscrowContract.TransactOpts, fundId)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.UpgradeToAndCall(&_ScholarshipEscrowContract.TransactOpts, newImplementation, data)
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ScholarshipEscrowContract.Contract.UpgradeToAndCall(&_ScholarshipEscrowContract.TransactOpts, newImplementation, data)
}
type ScholarshipEscrowContractFundClawedBackIterator struct {
	Event *ScholarshipEscrowContractFundClawedBack 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *ScholarshipEscrowContractFundClawedBackIterator) Next() bool {
	if it.fail != nil {
		return false
	}
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScholarshipEscrowContractFundClawedBack)
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
	select {
	case log := <-it.logs:
		it.Event = new(ScholarshipEscrowContractFundClawedBack)
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
func (it *ScholarshipEscrowContractFundClawedBackIterator) Error() error {
	return it.fail
}
func (it *ScholarshipEscrowContractFundClawedBackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type ScholarshipEscrowContractFundClawedBack struct {
	FundId    *big.Int
	Amount    *big.Int
	Recipient common.Address
	Raw       types.Log 
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) FilterFundClawedBack(opts *bind.FilterOpts, fundId []*big.Int) (*ScholarshipEscrowContractFundClawedBackIterator, error) {
	var fundIdRule []interface{}
	for _, fundIdItem := range fundId {
		fundIdRule = append(fundIdRule, fundIdItem)
	}
	logs, sub, err := _ScholarshipEscrowContract.contract.FilterLogs(opts, "FundClawedBack", fundIdRule)
	if err != nil {
		return nil, err
	}
	return &ScholarshipEscrowContractFundClawedBackIterator{contract: _ScholarshipEscrowContract.contract, event: "FundClawedBack", logs: logs, sub: sub}, nil
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) WatchFundClawedBack(opts *bind.WatchOpts, sink chan<- *ScholarshipEscrowContractFundClawedBack, fundId []*big.Int) (event.Subscription, error) {
	var fundIdRule []interface{}
	for _, fundIdItem := range fundId {
		fundIdRule = append(fundIdRule, fundIdItem)
	}
	logs, sub, err := _ScholarshipEscrowContract.contract.WatchLogs(opts, "FundClawedBack", fundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				event := new(ScholarshipEscrowContractFundClawedBack)
				if err := _ScholarshipEscrowContract.contract.UnpackLog(event, "FundClawedBack", log); err != nil {
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
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) ParseFundClawedBack(log types.Log) (*ScholarshipEscrowContractFundClawedBack, error) {
	event := new(ScholarshipEscrowContractFundClawedBack)
	if err := _ScholarshipEscrowContract.contract.UnpackLog(event, "FundClawedBack", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
type ScholarshipEscrowContractFundCreatedIterator struct {
	Event *ScholarshipEscrowContractFundCreated 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *ScholarshipEscrowContractFundCreatedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
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
func (it *ScholarshipEscrowContractFundCreatedIterator) Error() error {
	return it.fail
}
func (it *ScholarshipEscrowContractFundCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type ScholarshipEscrowContractFundCreated struct {
	FundId        *big.Int
	Sponsor       common.Address
	TotalAmount   *big.Int
	TrancheCount  *big.Int
	TrancheAmount *big.Int
	Raw           types.Log 
}
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
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) ParseFundCreated(log types.Log) (*ScholarshipEscrowContractFundCreated, error) {
	event := new(ScholarshipEscrowContractFundCreated)
	if err := _ScholarshipEscrowContract.contract.UnpackLog(event, "FundCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
type ScholarshipEscrowContractInitializedIterator struct {
	Event *ScholarshipEscrowContractInitialized 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *ScholarshipEscrowContractInitializedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScholarshipEscrowContractInitialized)
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
	select {
	case log := <-it.logs:
		it.Event = new(ScholarshipEscrowContractInitialized)
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
func (it *ScholarshipEscrowContractInitializedIterator) Error() error {
	return it.fail
}
func (it *ScholarshipEscrowContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type ScholarshipEscrowContractInitialized struct {
	Version uint64
	Raw     types.Log 
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*ScholarshipEscrowContractInitializedIterator, error) {
	logs, sub, err := _ScholarshipEscrowContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ScholarshipEscrowContractInitializedIterator{contract: _ScholarshipEscrowContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ScholarshipEscrowContractInitialized) (event.Subscription, error) {
	logs, sub, err := _ScholarshipEscrowContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				event := new(ScholarshipEscrowContractInitialized)
				if err := _ScholarshipEscrowContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) ParseInitialized(log types.Log) (*ScholarshipEscrowContractInitialized, error) {
	event := new(ScholarshipEscrowContractInitialized)
	if err := _ScholarshipEscrowContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
type ScholarshipEscrowContractMerkleRootPublishedIterator struct {
	Event *ScholarshipEscrowContractMerkleRootPublished 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *ScholarshipEscrowContractMerkleRootPublishedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
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
func (it *ScholarshipEscrowContractMerkleRootPublishedIterator) Error() error {
	return it.fail
}
func (it *ScholarshipEscrowContractMerkleRootPublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type ScholarshipEscrowContractMerkleRootPublished struct {
	FundId       *big.Int
	TrancheIndex *big.Int
	MerkleRoot   [32]byte
	Raw          types.Log 
}
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
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) ParseMerkleRootPublished(log types.Log) (*ScholarshipEscrowContractMerkleRootPublished, error) {
	event := new(ScholarshipEscrowContractMerkleRootPublished)
	if err := _ScholarshipEscrowContract.contract.UnpackLog(event, "MerkleRootPublished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
type ScholarshipEscrowContractRoleAdminChangedIterator struct {
	Event *ScholarshipEscrowContractRoleAdminChanged 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *ScholarshipEscrowContractRoleAdminChangedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
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
func (it *ScholarshipEscrowContractRoleAdminChangedIterator) Error() error {
	return it.fail
}
func (it *ScholarshipEscrowContractRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type ScholarshipEscrowContractRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log 
}
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
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) ParseRoleAdminChanged(log types.Log) (*ScholarshipEscrowContractRoleAdminChanged, error) {
	event := new(ScholarshipEscrowContractRoleAdminChanged)
	if err := _ScholarshipEscrowContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
type ScholarshipEscrowContractRoleGrantedIterator struct {
	Event *ScholarshipEscrowContractRoleGranted 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *ScholarshipEscrowContractRoleGrantedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
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
func (it *ScholarshipEscrowContractRoleGrantedIterator) Error() error {
	return it.fail
}
func (it *ScholarshipEscrowContractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type ScholarshipEscrowContractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log 
}
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
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) ParseRoleGranted(log types.Log) (*ScholarshipEscrowContractRoleGranted, error) {
	event := new(ScholarshipEscrowContractRoleGranted)
	if err := _ScholarshipEscrowContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
type ScholarshipEscrowContractRoleRevokedIterator struct {
	Event *ScholarshipEscrowContractRoleRevoked 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *ScholarshipEscrowContractRoleRevokedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
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
func (it *ScholarshipEscrowContractRoleRevokedIterator) Error() error {
	return it.fail
}
func (it *ScholarshipEscrowContractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type ScholarshipEscrowContractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log 
}
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
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) ParseRoleRevoked(log types.Log) (*ScholarshipEscrowContractRoleRevoked, error) {
	event := new(ScholarshipEscrowContractRoleRevoked)
	if err := _ScholarshipEscrowContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
type ScholarshipEscrowContractTrancheClaimedIterator struct {
	Event *ScholarshipEscrowContractTrancheClaimed 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *ScholarshipEscrowContractTrancheClaimedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
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
func (it *ScholarshipEscrowContractTrancheClaimedIterator) Error() error {
	return it.fail
}
func (it *ScholarshipEscrowContractTrancheClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type ScholarshipEscrowContractTrancheClaimed struct {
	FundId       *big.Int
	StudentHash  [32]byte
	TrancheIndex *big.Int
	Amount       *big.Int
	Recipient    common.Address
	Raw          types.Log 
}
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
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) ParseTrancheClaimed(log types.Log) (*ScholarshipEscrowContractTrancheClaimed, error) {
	event := new(ScholarshipEscrowContractTrancheClaimed)
	if err := _ScholarshipEscrowContract.contract.UnpackLog(event, "TrancheClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
type ScholarshipEscrowContractUpgradedIterator struct {
	Event *ScholarshipEscrowContractUpgraded 
	contract *bind.BoundContract 
	event    string              
	logs chan types.Log        
	sub  ethereum.Subscription 
	done bool                  
	fail error                 
}
func (it *ScholarshipEscrowContractUpgradedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScholarshipEscrowContractUpgraded)
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
	select {
	case log := <-it.logs:
		it.Event = new(ScholarshipEscrowContractUpgraded)
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
func (it *ScholarshipEscrowContractUpgradedIterator) Error() error {
	return it.fail
}
func (it *ScholarshipEscrowContractUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
type ScholarshipEscrowContractUpgraded struct {
	Implementation common.Address
	Raw            types.Log 
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ScholarshipEscrowContractUpgradedIterator, error) {
	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}
	logs, sub, err := _ScholarshipEscrowContract.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ScholarshipEscrowContractUpgradedIterator{contract: _ScholarshipEscrowContract.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ScholarshipEscrowContractUpgraded, implementation []common.Address) (event.Subscription, error) {
	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}
	logs, sub, err := _ScholarshipEscrowContract.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				event := new(ScholarshipEscrowContractUpgraded)
				if err := _ScholarshipEscrowContract.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_ScholarshipEscrowContract *ScholarshipEscrowContractFilterer) ParseUpgraded(log types.Log) (*ScholarshipEscrowContractUpgraded, error) {
	event := new(ScholarshipEscrowContractUpgraded)
	if err := _ScholarshipEscrowContract.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
