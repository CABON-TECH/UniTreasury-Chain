// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface ITreasury {
    event Deposited(address indexed depositor, uint256 amount, uint256 newBalance);
    event WithdrawalProposed(uint256 indexed proposalId, address indexed proposer, address indexed recipient, uint256 amount, string purpose);
    event WithdrawalApproved(uint256 indexed proposalId, address indexed approver, uint256 approvalCount);
    event WithdrawalExecuted(uint256 indexed proposalId, address indexed recipient, uint256 amount);
    event WithdrawalCancelled(uint256 indexed proposalId, address indexed canceller);
    event TreasuryFrozen(address indexed by);
    event TreasuryUnfrozen(address indexed by);
    event DailyLimitUpdated(uint256 oldLimit, uint256 newLimit);
    event RequiredApprovalsUpdated(uint256 oldRequired, uint256 newRequired);

    event SignerChangeProposed(uint256 indexed proposalId, address indexed proposer, address target, address replacement, uint8 changeType);
    event SignerChangeApproved(uint256 indexed proposalId, address indexed approver, uint256 approvalCount);
    event SignerChangeExecuted(uint256 indexed proposalId, address target, address replacement, uint8 changeType);


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

    enum ProposalStatus { Pending, Executed, Cancelled }

    enum ChangeType { Add, Remove, Replace }

    struct SignerChangeProposal {
        uint256 id;
        address proposer;
        address targetSigner;
        address newSigner;
        ChangeType changeType;
        ProposalStatus status;
        uint256 approvalCount;
        uint256 createdAt;
    }

    struct WithdrawalProposal {
        uint256 id;
        address proposer;
        address recipient;
        uint256 amount;
        string purpose;
        ProposalStatus status;
        uint256 approvalCount;
        uint256 createdAt;
    }

    function deposit(uint256 amount) external;
    function proposeWithdrawal(address recipient, uint256 amount, string calldata purpose) external returns (uint256 proposalId);
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

    function proposeSignerChange(address target, address replacement, uint8 changeType) external returns (uint256 proposalId);
    function approveSignerChange(uint256 proposalId) external;
    function executeSignerChange(uint256 proposalId) external;
    function cancelSignerChange(uint256 proposalId) external;
    function getSignerProposal(uint256 proposalId) external view returns (SignerChangeProposal memory);
    function hasApprovedSignerChange(uint256 proposalId, address approver) external view returns (bool);

}
