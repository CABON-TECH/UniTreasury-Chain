// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {ITreasury} from "./interfaces/ITreasury.sol";

/// @title TreasuryContract
/// @notice Multi-sig treasury for university funds with daily withdrawal limits.
///         Proposals are created by PROPOSER_ROLE, approved by APPROVER_ROLE (k-of-n),
///         and executed by EXECUTOR_ROLE (typically a contract admin) only after
///         requiredApprovals threshold is met.
///
/// @dev Security considerations:
///      - ReentrancyGuard on executeWithdrawal (checks-effects-interactions pattern)
///      - Approvals stored as mapping(proposalId => mapping(approver => bool)) to prevent
///        array-manipulation attacks and allow O(1) duplicate check
///      - Daily limit resets at UTC midnight (block.timestamp / 86400)
///      - Treasury freeze halts all new proposals and executions
contract TreasuryContract is ITreasury, AccessControl, ReentrancyGuard {
    // ── Roles ─────────────────────────────────────────────────────────────────────

    bytes32 public constant PROPOSER_ROLE = keccak256("PROPOSER_ROLE");
    bytes32 public constant APPROVER_ROLE = keccak256("APPROVER_ROLE");
    bytes32 public constant EXECUTOR_ROLE = keccak256("EXECUTOR_ROLE");
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    // ── State ─────────────────────────────────────────────────────────────────────

    uint256 private _proposalCounter;
    mapping(uint256 => WithdrawalProposal) private _proposals;
    mapping(uint256 => mapping(address => bool)) private _approvals;

    bool private _frozen;
    uint256 public requiredApprovals;
    uint256 public dailyLimit;

    // Daily limit tracking: UTC day => cumulative withdrawn
    mapping(uint256 => uint256) private _dailyWithdrawn;

    // ── Constructor ───────────────────────────────────────────────────────────────

    /// @param admin           Address that gets DEFAULT_ADMIN_ROLE + ADMIN_ROLE
    /// @param initialApprovers Addresses granted APPROVER_ROLE at deployment
    /// @param _requiredApprovals k-of-n threshold (must be <= initialApprovers.length)
    /// @param _dailyLimit     Maximum ETH withdrawable per UTC day (in wei)
    constructor(
        address admin,
        address[] memory initialApprovers,
        uint256 _requiredApprovals,
        uint256 _dailyLimit
    ) {
        if (admin == address(0)) revert Treasury__ZeroAddress();
        require(_requiredApprovals > 0 && _requiredApprovals <= initialApprovers.length,
            "TreasuryContract: invalid approval threshold");
        require(_dailyLimit > 0, "TreasuryContract: daily limit must be > 0");

        _grantRole(DEFAULT_ADMIN_ROLE, admin);
        _grantRole(ADMIN_ROLE, admin);
        _grantRole(PROPOSER_ROLE, admin);
        _grantRole(EXECUTOR_ROLE, admin);

        for (uint256 i = 0; i < initialApprovers.length; i++) {
            if (initialApprovers[i] == address(0)) revert Treasury__ZeroAddress();
            _grantRole(APPROVER_ROLE, initialApprovers[i]);
        }

        requiredApprovals = _requiredApprovals;
        dailyLimit = _dailyLimit;
    }

    // ── Receive ───────────────────────────────────────────────────────────────────

    receive() external payable {
        emit Deposited(msg.sender, msg.value, address(this).balance);
    }

    // ── External: Finance Operations ──────────────────────────────────────────────

    /// @inheritdoc ITreasury
    function deposit() external payable override {
        emit Deposited(msg.sender, msg.value, address(this).balance);
    }

    /// @inheritdoc ITreasury
    function proposeWithdrawal(
        address payable recipient,
        uint256 amount,
        string calldata purpose
    ) external override onlyRole(PROPOSER_ROLE) returns (uint256 proposalId) {
        if (_frozen) revert Treasury__Frozen();
        if (recipient == address(0)) revert Treasury__ZeroAddress();
        if (amount == 0) revert Treasury__ZeroAmount();
        if (address(this).balance < amount)
            revert Treasury__InsufficientBalance(amount, address(this).balance);

        proposalId = ++_proposalCounter;
        _proposals[proposalId] = WithdrawalProposal({
            id: proposalId,
            proposer: msg.sender,
            recipient: recipient,
            amount: amount,
            purpose: purpose,
            status: ProposalStatus.Pending,
            approvalCount: 0,
            createdAt: block.timestamp
        });

        emit WithdrawalProposed(proposalId, msg.sender, recipient, amount, purpose);
    }

    /// @inheritdoc ITreasury
    function approveWithdrawal(uint256 proposalId) external override onlyRole(APPROVER_ROLE) {
        WithdrawalProposal storage proposal = _getActivePendingProposal(proposalId);
        if (_approvals[proposalId][msg.sender])
            revert Treasury__AlreadyApproved(proposalId, msg.sender);

        _approvals[proposalId][msg.sender] = true;
        proposal.approvalCount++;

        emit WithdrawalApproved(proposalId, msg.sender, proposal.approvalCount);
    }

    /// @inheritdoc ITreasury
    /// @dev Checks-effects-interactions: all state updated before external call
    function executeWithdrawal(uint256 proposalId)
        external
        override
        nonReentrant
        onlyRole(EXECUTOR_ROLE)
    {
        if (_frozen) revert Treasury__Frozen();
        WithdrawalProposal storage proposal = _getActivePendingProposal(proposalId);

        if (proposal.approvalCount < requiredApprovals)
            revert Treasury__InsufficientApprovals(proposal.approvalCount, requiredApprovals);
        if (address(this).balance < proposal.amount)
            revert Treasury__InsufficientBalance(proposal.amount, address(this).balance);

        // Daily limit check
        uint256 today = block.timestamp / 1 days;
        uint256 alreadyToday = _dailyWithdrawn[today];
        if (alreadyToday + proposal.amount > dailyLimit)
            revert Treasury__DailyLimitExceeded(proposal.amount, dailyLimit - alreadyToday);

        // Effects
        proposal.status = ProposalStatus.Executed;
        _dailyWithdrawn[today] = alreadyToday + proposal.amount;
        uint256 amount = proposal.amount;
        address payable recipient = proposal.recipient;

        // Interaction
        (bool ok,) = recipient.call{value: amount}("");
        require(ok, "TreasuryContract: transfer failed");

        emit WithdrawalExecuted(proposalId, recipient, amount);
    }

    /// @inheritdoc ITreasury
    function cancelWithdrawal(uint256 proposalId) external override {
        WithdrawalProposal storage proposal = _getActivePendingProposal(proposalId);
        // Only the proposer or an admin can cancel
        if (msg.sender != proposal.proposer && !hasRole(ADMIN_ROLE, msg.sender))
            revert Treasury__NotAuthorized();

        proposal.status = ProposalStatus.Cancelled;
        emit WithdrawalCancelled(proposalId, msg.sender);
    }

    // ── External: Admin Operations ────────────────────────────────────────────────

    /// @inheritdoc ITreasury
    function freeze() external override onlyRole(ADMIN_ROLE) {
        _frozen = true;
        emit TreasuryFrozen(msg.sender);
    }

    /// @inheritdoc ITreasury
    function unfreeze() external override onlyRole(ADMIN_ROLE) {
        _frozen = false;
        emit TreasuryUnfrozen(msg.sender);
    }

    /// @inheritdoc ITreasury
    function setDailyLimit(uint256 newLimit) external override onlyRole(ADMIN_ROLE) {
        require(newLimit > 0, "TreasuryContract: limit must be > 0");
        emit DailyLimitUpdated(dailyLimit, newLimit);
        dailyLimit = newLimit;
    }

    /// @inheritdoc ITreasury
    function setRequiredApprovals(uint256 newRequired) external override onlyRole(ADMIN_ROLE) {
        require(newRequired > 0, "TreasuryContract: must require >= 1 approval");
        emit RequiredApprovalsUpdated(requiredApprovals, newRequired);
        requiredApprovals = newRequired;
    }

    // ── External: View ────────────────────────────────────────────────────────────

    /// @inheritdoc ITreasury
    function getProposal(uint256 proposalId) external view override returns (WithdrawalProposal memory) {
        if (_proposals[proposalId].id == 0) revert Treasury__ProposalNotFound(proposalId);
        return _proposals[proposalId];
    }

    /// @inheritdoc ITreasury
    function hasApproved(uint256 proposalId, address approver) external view override returns (bool) {
        return _approvals[proposalId][approver];
    }

    /// @inheritdoc ITreasury
    function getBalance() external view override returns (uint256) {
        return address(this).balance;
    }

    /// @inheritdoc ITreasury
    function isFrozen() external view override returns (bool) {
        return _frozen;
    }

    /// @inheritdoc ITreasury
    function getDailyWithdrawnAmount() external view override returns (uint256) {
        return _dailyWithdrawn[block.timestamp / 1 days];
    }

    // ── Internal Helpers ──────────────────────────────────────────────────────────

    function _getActivePendingProposal(uint256 proposalId)
        internal
        view
        returns (WithdrawalProposal storage proposal)
    {
        proposal = _proposals[proposalId];
        if (proposal.id == 0) revert Treasury__ProposalNotFound(proposalId);
        if (proposal.status != ProposalStatus.Pending)
            revert Treasury__ProposalNotPending(proposalId);
    }
}
