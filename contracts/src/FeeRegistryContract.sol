// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {IFeeRegistry} from "./interfaces/IFeeRegistry.sol";

/// @title FeeRegistryContract
/// @notice Immutable on-chain ledger of student fee payments.
///         Student identities are stored as keccak256 hashes (pseudonymized) per the
///         system's privacy layer — raw student IDs never appear on-chain.
///
/// @dev Design decisions:
///      - receiptHash is the primary deduplication key (unique per payment)
///      - feeType is a bitmask (1=tuition, 2=hostel, 4=exam) allowing combined payments
///      - Fee structures are version-controlled per semester to support tuition changes
///      - checkpoint() emits a reconciliation anchor for the off-chain indexer
contract FeeRegistryContract is IFeeRegistry, AccessControl {
    // ── Roles ─────────────────────────────────────────────────────────────────────

    bytes32 public constant RECORDER_ROLE = keccak256("RECORDER_ROLE");
    bytes32 public constant FEE_ADMIN_ROLE = keccak256("FEE_ADMIN_ROLE");

    // ── State ─────────────────────────────────────────────────────────────────────

    // receiptHash => PaymentRecord
    mapping(bytes32 => PaymentRecord) private _payments;

    // studentHash => semester => paid
    mapping(bytes32 => mapping(uint256 => bool)) private _studentSemesterPaid;

    // semester => FeeStructure
    mapping(uint256 => FeeStructure) private _feeStructures;

    uint256 private _totalRecords;
    uint256 private _totalVolume;

    // ── Constructor ───────────────────────────────────────────────────────────────

    /// @param admin    Address granted DEFAULT_ADMIN_ROLE + FEE_ADMIN_ROLE
    /// @param recorder Address granted RECORDER_ROLE (the Go backend service wallet)
    constructor(address admin, address recorder) {
        if (admin == address(0) || recorder == address(0)) revert FeeRegistry__NotAuthorized();

        _grantRole(DEFAULT_ADMIN_ROLE, admin);
        _grantRole(FEE_ADMIN_ROLE, admin);
        _grantRole(RECORDER_ROLE, recorder);
    }

    // ── External: Payment Recording ────────────────────────────────────────────────

    /// @inheritdoc IFeeRegistry
    /// @dev Only RECORDER_ROLE (Go backend) can call this. The backend validates CSV data,
    ///      computes studentHash = keccak256(studentId), and submits the batch.
    function recordPayment(
        bytes32 studentHash,
        bytes32 receiptHash,
        uint256 amount,
        uint256 semester,
        uint256 feeType
    ) external override onlyRole(RECORDER_ROLE) {
        if (studentHash == bytes32(0)) revert FeeRegistry__ZeroStudentHash();
        if (semester == 0) revert FeeRegistry__InvalidSemester(semester);
        if (_payments[receiptHash].recordedAt != 0)
            revert FeeRegistry__DuplicateReceipt(receiptHash);

        // Validate against active fee structure if one exists for this semester
        FeeStructure storage structure = _feeStructures[semester];
        if (structure.active) {
            uint256 expectedAmount = _computeExpectedAmount(structure, feeType);
            if (expectedAmount > 0 && amount != expectedAmount)
                revert FeeRegistry__AmountMismatch(amount, expectedAmount);
        }

        _payments[receiptHash] = PaymentRecord({
            studentHash: studentHash,
            receiptHash: receiptHash,
            amount: amount,
            semester: semester,
            feeType: feeType,
            recordedAt: block.timestamp
        });

        _studentSemesterPaid[studentHash][semester] = true;

        unchecked {
            _totalRecords++;
            _totalVolume += amount;
        }

        emit PaymentRecorded(studentHash, receiptHash, amount, semester, block.timestamp);
    }

    // ── External: Fee Structure Management ────────────────────────────────────────

    /// @inheritdoc IFeeRegistry
    function setFeeStructure(
        uint256 semester,
        uint256 tuitionFee,
        uint256 hostingFee,
        uint256 examFee
    ) external override onlyRole(FEE_ADMIN_ROLE) {
        if (semester == 0) revert FeeRegistry__InvalidSemester(semester);

        _feeStructures[semester] = FeeStructure({
            tuitionFee: tuitionFee,
            hostingFee: hostingFee,
            examFee: examFee,
            active: true
        });

        emit FeeStructureSet(semester, tuitionFee, hostingFee, examFee);
    }

    /// @inheritdoc IFeeRegistry
    function deactivateFeeStructure(uint256 semester) external override onlyRole(FEE_ADMIN_ROLE) {
        _feeStructures[semester].active = false;
        emit FeeStructureDeactivated(semester);
    }

    // ── External: Reconciliation ───────────────────────────────────────────────────

    /// @inheritdoc IFeeRegistry
    /// @notice Emits a reconciliation checkpoint — off-chain indexer uses this as an
    ///         anchor to verify its running totals match on-chain state.
    function checkpoint() external override onlyRole(RECORDER_ROLE) {
        emit ReconciliationCheckpoint(_totalRecords, _totalVolume, block.number);
    }

    // ── External: View ────────────────────────────────────────────────────────────

    /// @inheritdoc IFeeRegistry
    function getPaymentRecord(bytes32 receiptHash)
        external
        view
        override
        returns (PaymentRecord memory)
    {
        return _payments[receiptHash];
    }

    /// @inheritdoc IFeeRegistry
    function hasStudentPaidSemester(bytes32 studentHash, uint256 semester)
        external
        view
        override
        returns (bool)
    {
        return _studentSemesterPaid[studentHash][semester];
    }

    /// @inheritdoc IFeeRegistry
    function getFeeStructure(uint256 semester)
        external
        view
        override
        returns (FeeStructure memory)
    {
        return _feeStructures[semester];
    }

    /// @inheritdoc IFeeRegistry
    function getTotalVolume() external view override returns (uint256) {
        return _totalVolume;
    }

    /// @inheritdoc IFeeRegistry
    function getTotalRecords() external view override returns (uint256) {
        return _totalRecords;
    }

    // ── Internal Helpers ──────────────────────────────────────────────────────────

    /// @dev Compute the expected payment amount given a fee structure and feeType bitmask.
    ///      Returns 0 if the feeType is not covered by the structure (unvalidated).
    function _computeExpectedAmount(FeeStructure storage structure, uint256 feeType)
        internal
        view
        returns (uint256 expected)
    {
        if (feeType & 1 != 0) expected += structure.tuitionFee;
        if (feeType & 2 != 0) expected += structure.hostingFee;
        if (feeType & 4 != 0) expected += structure.examFee;
    }
}
