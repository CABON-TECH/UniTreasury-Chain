// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title ITreasury
/// @notice Interface for the multi-sig treasury contract
interface ITreasury {
    // ── Events ──────────────────────────────────────────────────────────────────

    event Deposited(address indexed depositor, uint256 amount, uint256 newBalance);
    event WithdrawalProposed(
        uint256 indexed proposalId,
        address indexed proposer,
        address indexed recipient,
        uint256 amount,
        string purpose
    );
    event WithdrawalApproved(uint256 indexed proposalId, address indexed approver, uint256 approvalCount);
    event WithdrawalExecuted(uint256 indexed proposalId, address indexed recipient, uint256 amount);
    event WithdrawalCancelled(uint256 indexed proposalId, address indexed canceller);
    event TreasuryFrozen(address indexed by);
    event TreasuryUnfrozen(address indexed by);
    event DailyLimitUpdated(uint256 oldLimit, uint256 newLimit);
    event RequiredApprovalsUpdated(uint256 oldRequired, uint256 newRequired);

    // ── Errors ───────────────────────────────────────────────────────────────────

    error Treasury__Frozen();
    error Treasury__InsufficientBalance(uint256 requested, uint256 available);
    error Treasury__DailyLimitExceeded(uint256 requested, uint256 remainingToday);
    error Treasury__ProposalNotFound(uint256 proposalId);
    error Treasury__ProposalNotPending(uint256 proposalId);
    error Treasury__AlreadyApproved(uint256 proposalId, address approver);
    error Treasury__InsufficientApprovals(uint256 have, uint256 need);
    error Treasury__ZeroAddress();
    error Treasury__ZeroAmount();
    error Treasury__NotAuthorized();

    // ── Types ────────────────────────────────────────────────────────────────────

    enum ProposalStatus {
        Pending,
        Executed,
        Cancelled
    }

    struct WithdrawalProposal {
        uint256 id;
        address proposer;
        address payable recipient;
        uint256 amount;
        string purpose;
        ProposalStatus status;
        uint256 approvalCount;
        uint256 createdAt;
    }

    // ── Functions ────────────────────────────────────────────────────────────────

    function deposit() external payable;

    function proposeWithdrawal(
        address payable recipient,
        uint256 amount,
        string calldata purpose
    ) external returns (uint256 proposalId);

    function approveWithdrawal(uint256 proposalId) external;

    function executeWithdrawal(uint256 proposalId) external;

    function cancelWithdrawal(uint256 proposalId) external;

    function freeze() external;

    function unfreeze() external;

    function setDailyLimit(uint256 newLimit) external;

    function setRequiredApprovals(uint256 newRequired) external;

    function getProposal(uint256 proposalId) external view returns (WithdrawalProposal memory);

    function hasApproved(uint256 proposalId, address approver) external view returns (bool);

    function getBalance() external view returns (uint256);

    function isFrozen() external view returns (bool);

    function getDailyWithdrawnAmount() external view returns (uint256);
}
