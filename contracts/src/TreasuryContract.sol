// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {ITreasury} from "./interfaces/ITreasury.sol";

contract TreasuryContract is ITreasury, AccessControl, ReentrancyGuard {
    bytes32 public constant ADMIN_ROLE = DEFAULT_ADMIN_ROLE;
    bytes32 public constant PROPOSER_ROLE = keccak256("PROPOSER_ROLE");
    bytes32 public constant APPROVER_ROLE = keccak256("APPROVER_ROLE");
    bytes32 public constant EXECUTOR_ROLE = keccak256("EXECUTOR_ROLE");

    IERC20 public immutable usdcToken;

    uint256 public requiredApprovals;
    uint256 public dailyLimit;
    
    bool private _frozen;
    uint256 private _proposalCounter;
    
    mapping(uint256 => WithdrawalProposal) private _proposals;
    mapping(uint256 => mapping(address => bool)) private _approvals;
    mapping(uint256 => uint256) private _dailyWithdrawn;

    uint256 private _signerProposalCounter;
    mapping(uint256 => SignerChangeProposal) private _signerProposals;
    mapping(uint256 => mapping(address => bool)) private _signerApprovals;


    constructor(
        address admin,
        address[] memory initialApprovers,
        uint256 _requiredApprovals,
        uint256 _dailyLimit,
        address _usdcToken
    ) {
        if (admin == address(0)) revert Treasury__ZeroAddress();
        require(_requiredApprovals > 0 && _requiredApprovals <= initialApprovers.length, "Invalid approval threshold");
        require(_dailyLimit > 0, "Daily limit must be > 0");

        _grantRole(DEFAULT_ADMIN_ROLE, admin);
        _grantRole(PROPOSER_ROLE, admin);
        _grantRole(EXECUTOR_ROLE, admin);

        for (uint256 i = 0; i < initialApprovers.length; i++) {
            if (initialApprovers[i] == address(0)) revert Treasury__ZeroAddress();
            _grantRole(APPROVER_ROLE, initialApprovers[i]);
        }

        requiredApprovals = _requiredApprovals;
        dailyLimit = _dailyLimit;
        usdcToken = IERC20(_usdcToken);
    }

    function deposit(uint256 amount) external override {
        require(amount > 0, "Amount must be > 0");
        require(usdcToken.transferFrom(msg.sender, address(this), amount), "Transfer failed");
        emit Deposited(msg.sender, amount, getBalance());
    }

    function proposeWithdrawal(
        address recipient,
        uint256 amount,
        string calldata purpose
    ) external override onlyRole(PROPOSER_ROLE) returns (uint256 proposalId) {
        if (_frozen) revert Treasury__Frozen();
        if (recipient == address(0)) revert Treasury__ZeroAddress();
        if (amount == 0) revert Treasury__ZeroAmount();
        if (getBalance() < amount)
            revert Treasury__InsufficientBalance(amount, getBalance());

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

    function approveWithdrawal(uint256 proposalId) external override onlyRole(APPROVER_ROLE) {
        WithdrawalProposal storage proposal = _getActivePendingProposal(proposalId);
        if (_approvals[proposalId][msg.sender])
            revert Treasury__AlreadyApproved(proposalId, msg.sender);

        _approvals[proposalId][msg.sender] = true;
        proposal.approvalCount++;

        emit WithdrawalApproved(proposalId, msg.sender, proposal.approvalCount);
    }

    function executeWithdrawal(uint256 proposalId) external override nonReentrant onlyRole(EXECUTOR_ROLE) {
        if (_frozen) revert Treasury__Frozen();
        WithdrawalProposal storage proposal = _getActivePendingProposal(proposalId);

        if (proposal.approvalCount < requiredApprovals)
            revert Treasury__InsufficientApprovals(proposal.approvalCount, requiredApprovals);
        if (getBalance() < proposal.amount)
            revert Treasury__InsufficientBalance(proposal.amount, getBalance());

        uint256 today = block.timestamp / 1 days;
        uint256 alreadyToday = _dailyWithdrawn[today];
        if (alreadyToday + proposal.amount > dailyLimit)
            revert Treasury__DailyLimitExceeded(proposal.amount, dailyLimit - alreadyToday);

        proposal.status = ProposalStatus.Executed;
        _dailyWithdrawn[today] = alreadyToday + proposal.amount;
        
        require(usdcToken.transfer(proposal.recipient, proposal.amount), "Transfer failed");

        emit WithdrawalExecuted(proposalId, proposal.recipient, proposal.amount);
    }

    function cancelWithdrawal(uint256 proposalId) external override {
        WithdrawalProposal storage proposal = _getActivePendingProposal(proposalId);
        if (msg.sender != proposal.proposer && !hasRole(ADMIN_ROLE, msg.sender))
            revert Treasury__NotAuthorized();

        proposal.status = ProposalStatus.Cancelled;
        emit WithdrawalCancelled(proposalId, msg.sender);
    }

    function freeze() external override onlyRole(ADMIN_ROLE) {
        _frozen = true;
        emit TreasuryFrozen(msg.sender);
    }

    function unfreeze() external override onlyRole(ADMIN_ROLE) {
        _frozen = false;
        emit TreasuryUnfrozen(msg.sender);
    }

    function setDailyLimit(uint256 newLimit) external override onlyRole(ADMIN_ROLE) {
        require(newLimit > 0, "Limit must be > 0");
        emit DailyLimitUpdated(dailyLimit, newLimit);
        dailyLimit = newLimit;
    }

    function setRequiredApprovals(uint256 newRequired) external override onlyRole(ADMIN_ROLE) {
        require(newRequired > 0, "Must require >= 1 approval");
        emit RequiredApprovalsUpdated(requiredApprovals, newRequired);
        requiredApprovals = newRequired;
    }

    function getProposal(uint256 proposalId) external view override returns (WithdrawalProposal memory) {
        if (_proposals[proposalId].id == 0) revert Treasury__ProposalNotFound(proposalId);
        return _proposals[proposalId];
    }

    function hasApproved(uint256 proposalId, address approver) external view override returns (bool) {
        return _approvals[proposalId][approver];
    }

    function getBalance() public view override returns (uint256) {
        return usdcToken.balanceOf(address(this));
    }

    function isFrozen() external view override returns (bool) {
        return _frozen;
    }

    function getDailyWithdrawnAmount() external view override returns (uint256) {
        return _dailyWithdrawn[block.timestamp / 1 days];
    }

    function _getActivePendingProposal(uint256 proposalId) internal view returns (WithdrawalProposal storage proposal) {
        proposal = _proposals[proposalId];
        if (proposal.id == 0) revert Treasury__ProposalNotFound(proposalId);
        if (proposal.status != ProposalStatus.Pending) revert Treasury__ProposalNotPending(proposalId);
    }

    // --- Signer Change Logic ---

    function proposeSignerChange(
        address target,
        address replacement,
        uint8 changeType
    ) external override onlyRole(APPROVER_ROLE) returns (uint256 proposalId) {
        if (_frozen) revert Treasury__Frozen();
        require(changeType <= 2, "Invalid change type");
        if (changeType == uint8(ChangeType.Add) || changeType == uint8(ChangeType.Replace)) {
            require(replacement != address(0), "Invalid replacement address");
        }
        if (changeType == uint8(ChangeType.Remove) || changeType == uint8(ChangeType.Replace)) {
            require(target != address(0), "Invalid target address");
        }

        proposalId = ++_signerProposalCounter;
        _signerProposals[proposalId] = SignerChangeProposal({
            id: proposalId,
            proposer: msg.sender,
            targetSigner: target,
            newSigner: replacement,
            changeType: ChangeType(changeType),
            status: ProposalStatus.Pending,
            approvalCount: 0,
            createdAt: block.timestamp
        });

        emit SignerChangeProposed(proposalId, msg.sender, target, replacement, changeType);
    }

    function approveSignerChange(uint256 proposalId) external override onlyRole(APPROVER_ROLE) {
        SignerChangeProposal storage proposal = _getActivePendingSignerProposal(proposalId);
        if (_signerApprovals[proposalId][msg.sender])
            revert Treasury__AlreadyApproved(proposalId, msg.sender);

        _signerApprovals[proposalId][msg.sender] = true;
        proposal.approvalCount++;

        emit SignerChangeApproved(proposalId, msg.sender, proposal.approvalCount);
    }

    function executeSignerChange(uint256 proposalId) external override nonReentrant onlyRole(EXECUTOR_ROLE) {
        if (_frozen) revert Treasury__Frozen();
        SignerChangeProposal storage proposal = _getActivePendingSignerProposal(proposalId);

        if (proposal.approvalCount < requiredApprovals)
            revert Treasury__InsufficientApprovals(proposal.approvalCount, requiredApprovals);

        proposal.status = ProposalStatus.Executed;

        if (proposal.changeType == ChangeType.Add) {
            _grantRole(APPROVER_ROLE, proposal.newSigner);
        } else if (proposal.changeType == ChangeType.Remove) {
            _revokeRole(APPROVER_ROLE, proposal.targetSigner);
        } else if (proposal.changeType == ChangeType.Replace) {
            _revokeRole(APPROVER_ROLE, proposal.targetSigner);
            _grantRole(APPROVER_ROLE, proposal.newSigner);
        }

        emit SignerChangeExecuted(proposalId, proposal.targetSigner, proposal.newSigner, uint8(proposal.changeType));
    }

    function cancelSignerChange(uint256 proposalId) external override {
        SignerChangeProposal storage proposal = _getActivePendingSignerProposal(proposalId);
        if (msg.sender != proposal.proposer && !hasRole(ADMIN_ROLE, msg.sender))
            revert Treasury__NotAuthorized();

        proposal.status = ProposalStatus.Cancelled;
        // Reusing the same Cancelled event but we should maybe have a specific one, or just update state.
    }

    function getSignerProposal(uint256 proposalId) external view override returns (SignerChangeProposal memory) {
        if (_signerProposals[proposalId].id == 0) revert Treasury__ProposalNotFound(proposalId);
        return _signerProposals[proposalId];
    }

    function hasApprovedSignerChange(uint256 proposalId, address approver) external view override returns (bool) {
        return _signerApprovals[proposalId][approver];
    }

    function _getActivePendingSignerProposal(uint256 proposalId) internal view returns (SignerChangeProposal storage proposal) {
        proposal = _signerProposals[proposalId];
        if (proposal.id == 0) revert Treasury__ProposalNotFound(proposalId);
        if (proposal.status != ProposalStatus.Pending) revert Treasury__ProposalNotPending(proposalId);
    }
}
