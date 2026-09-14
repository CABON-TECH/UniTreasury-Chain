package bindings

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TreasuryContractABI is the ABI JSON for TreasuryContract.
const TreasuryContractABI = `[
  {"type":"constructor","inputs":[{"name":"admin","type":"address"},{"name":"initialApprovers","type":"address[]"},{"name":"_requiredApprovals","type":"uint256"},{"name":"_dailyLimit","type":"uint256"}],"stateMutability":"nonpayable"},
  {"type":"function","name":"deposit","inputs":[],"outputs":[],"stateMutability":"payable"},
  {"type":"function","name":"proposeWithdrawal","inputs":[{"name":"recipient","type":"address"},{"name":"amount","type":"uint256"},{"name":"purpose","type":"string"}],"outputs":[{"name":"proposalId","type":"uint256"}],"stateMutability":"nonpayable"},
  {"type":"function","name":"approveWithdrawal","inputs":[{"name":"proposalId","type":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},
  {"type":"function","name":"executeWithdrawal","inputs":[{"name":"proposalId","type":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},
  {"type":"function","name":"cancelWithdrawal","inputs":[{"name":"proposalId","type":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},
  {"type":"function","name":"freeze","inputs":[],"outputs":[],"stateMutability":"nonpayable"},
  {"type":"function","name":"unfreeze","inputs":[],"outputs":[],"stateMutability":"nonpayable"},
  {"type":"function","name":"setDailyLimit","inputs":[{"name":"newLimit","type":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},
  {"type":"function","name":"setRequiredApprovals","inputs":[{"name":"newRequired","type":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},
  {"type":"function","name":"getBalance","inputs":[],"outputs":[{"name":"","type":"uint256"}],"stateMutability":"view"},
  {"type":"function","name":"isFrozen","inputs":[],"outputs":[{"name":"","type":"bool"}],"stateMutability":"view"},
  {"type":"function","name":"requiredApprovals","inputs":[],"outputs":[{"name":"","type":"uint256"}],"stateMutability":"view"},
  {"type":"function","name":"dailyLimit","inputs":[],"outputs":[{"name":"","type":"uint256"}],"stateMutability":"view"},
  {"type":"function","name":"getDailyWithdrawnAmount","inputs":[],"outputs":[{"name":"","type":"uint256"}],"stateMutability":"view"},
  {"type":"function","name":"hasApproved","inputs":[{"name":"proposalId","type":"uint256"},{"name":"approver","type":"address"}],"outputs":[{"name":"","type":"bool"}],"stateMutability":"view"},
  {"type":"event","name":"WithdrawalProposed","inputs":[{"name":"proposalId","type":"uint256","indexed":true},{"name":"proposer","type":"address","indexed":true},{"name":"recipient","type":"address","indexed":true},{"name":"amount","type":"uint256","indexed":false},{"name":"purpose","type":"string","indexed":false}],"anonymous":false},
  {"type":"event","name":"WithdrawalApproved","inputs":[{"name":"proposalId","type":"uint256","indexed":true},{"name":"approver","type":"address","indexed":true},{"name":"approvalCount","type":"uint256","indexed":false}],"anonymous":false},
  {"type":"event","name":"WithdrawalExecuted","inputs":[{"name":"proposalId","type":"uint256","indexed":true},{"name":"recipient","type":"address","indexed":true},{"name":"amount","type":"uint256","indexed":false}],"anonymous":false},
  {"type":"event","name":"Deposited","inputs":[{"name":"depositor","type":"address","indexed":true},{"name":"amount","type":"uint256","indexed":false},{"name":"newBalance","type":"uint256","indexed":false}],"anonymous":false}
]`

// TreasuryContract is a Go binding for TreasuryContract.
type TreasuryContract struct {
	TreasuryContractCaller
	TreasuryContractTransactor
}

// TreasuryContractCaller contains read-only methods.
type TreasuryContractCaller struct {
	contract *bind.BoundContract
}

// TreasuryContractTransactor contains state-changing methods.
type TreasuryContractTransactor struct {
	contract *bind.BoundContract
}

// NewTreasuryContract creates a new binding instance.
func NewTreasuryContract(address common.Address, backend bind.ContractBackend) (*TreasuryContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TreasuryContractABI))
	if err != nil {
		return nil, err
	}
	contract := bind.NewBoundContract(address, parsed, backend, backend, backend)
	return &TreasuryContract{
		TreasuryContractCaller:     TreasuryContractCaller{contract: contract},
		TreasuryContractTransactor: TreasuryContractTransactor{contract: contract},
	}, nil
}

// ── Transactor methods ────────────────────────────────────────────────────────

func (t *TreasuryContractTransactor) ProposeWithdrawal(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, purpose string) (*types.Transaction, error) {
	return t.contract.Transact(opts, "proposeWithdrawal", recipient, amount, purpose)
}

func (t *TreasuryContractTransactor) ApproveWithdrawal(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return t.contract.Transact(opts, "approveWithdrawal", proposalId)
}

func (t *TreasuryContractTransactor) ExecuteWithdrawal(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return t.contract.Transact(opts, "executeWithdrawal", proposalId)
}

func (t *TreasuryContractTransactor) CancelWithdrawal(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return t.contract.Transact(opts, "cancelWithdrawal", proposalId)
}

func (t *TreasuryContractTransactor) Freeze(opts *bind.TransactOpts) (*types.Transaction, error) {
	return t.contract.Transact(opts, "freeze")
}

func (t *TreasuryContractTransactor) Unfreeze(opts *bind.TransactOpts) (*types.Transaction, error) {
	return t.contract.Transact(opts, "unfreeze")
}

// ── Caller methods ────────────────────────────────────────────────────────────

func (c *TreasuryContractCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	if err := c.contract.Call(opts, &out, "getBalance"); err != nil {
		return nil, err
	}
	return *abi.ConvertType(out[0], new(*big.Int)).(**big.Int), nil
}

func (c *TreasuryContractCaller) IsFrozen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	if err := c.contract.Call(opts, &out, "isFrozen"); err != nil {
		return false, err
	}
	return *abi.ConvertType(out[0], new(bool)).(*bool), nil
}

func (c *TreasuryContractCaller) GetDailyWithdrawnAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	if err := c.contract.Call(opts, &out, "getDailyWithdrawnAmount"); err != nil {
		return nil, err
	}
	return *abi.ConvertType(out[0], new(*big.Int)).(**big.Int), nil
}
