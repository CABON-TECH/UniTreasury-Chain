// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title IScholarshipEscrow
/// @notice Interface for the scholarship escrow contract with attestation-gated tranche release
interface IScholarshipEscrow {
    // ── Events ──────────────────────────────────────────────────────────────────

    event FundCreated(
        uint256 indexed fundId,
        address indexed sponsor,
        uint256 totalAmount,
        uint256 trancheCount
    );
    event TrancheFunded(uint256 indexed fundId, uint256 trancheIndex, uint256 amount);
    event TrancheReleased(
        uint256 indexed fundId,
        bytes32 indexed studentHash,
        uint256 trancheIndex,
        uint256 amount,
        address recipient
    );
    event AttestorUpdated(address indexed oldAttestor, address indexed newAttestor);
    event FundPaused(uint256 indexed fundId);
    event FundResumed(uint256 indexed fundId);

    // ── Errors ───────────────────────────────────────────────────────────────────

    error Escrow__FundNotFound(uint256 fundId);
    error Escrow__InvalidTranche(uint256 fundId, uint256 trancheIndex);
    error Escrow__TrancheAlreadyReleased(uint256 fundId, uint256 trancheIndex);
    error Escrow__InvalidAttestation();
    error Escrow__AttestationReused(bytes32 nonce);
    error Escrow__FundPaused(uint256 fundId);
    error Escrow__InsufficientFundBalance(uint256 fundId, uint256 needed, uint256 available);
    error Escrow__ZeroAttestor();
    error Escrow__NotAuthorized();

    // ── Types ────────────────────────────────────────────────────────────────────

    struct ScholarshipFund {
        uint256 id;
        address sponsor;
        uint256 totalAmount;
        uint256 releasedAmount;
        uint256 trancheCount;
        uint256 trancheAmount;  // amount per tranche
        bool paused;
    }

    // ── Functions ────────────────────────────────────────────────────────────────

    function createFund(uint256 trancheCount) external payable returns (uint256 fundId);

    function releaseTranche(
        uint256 fundId,
        bytes32 studentHash,
        uint256 trancheIndex,
        address payable recipient,
        bytes calldata attestation
    ) external;

    function pauseFund(uint256 fundId) external;

    function resumeFund(uint256 fundId) external;

    function setAttestor(address newAttestor) external;

    function getFund(uint256 fundId) external view returns (ScholarshipFund memory);

    function isTrancheReleased(uint256 fundId, uint256 trancheIndex) external view returns (bool);

    function getAttestor() external view returns (address);
}
