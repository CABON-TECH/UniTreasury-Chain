// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title IFeeRegistry
/// @notice Interface for the student fee registry contract
interface IFeeRegistry {
    // ── Events ──────────────────────────────────────────────────────────────────

    event PaymentRecorded(
        bytes32 indexed studentHash,
        bytes32 indexed receiptHash,
        uint256 amount,
        uint256 semester,
        uint256 timestamp
    );
    event FeeStructureSet(uint256 indexed semester, uint256 tuitionFee, uint256 hostingFee, uint256 examFee);
    event FeeStructureDeactivated(uint256 indexed semester);
    event ReconciliationCheckpoint(uint256 totalRecords, uint256 totalVolume, uint256 blockNumber);

    // ── Errors ───────────────────────────────────────────────────────────────────

    error FeeRegistry__DuplicateReceipt(bytes32 receiptHash);
    error FeeRegistry__InvalidSemester(uint256 semester);
    error FeeRegistry__FeeStructureNotActive(uint256 semester);
    error FeeRegistry__AmountMismatch(uint256 paid, uint256 expected);
    error FeeRegistry__ZeroStudentHash();
    error FeeRegistry__NotAuthorized();

    // ── Types ────────────────────────────────────────────────────────────────────

    struct FeeStructure {
        uint256 tuitionFee;
        uint256 hostingFee;
        uint256 examFee;
        bool active;
    }

    struct PaymentRecord {
        bytes32 studentHash;   // keccak256 of student ID — pseudonymized
        bytes32 receiptHash;   // unique receipt identifier
        uint256 amount;
        uint256 semester;
        uint256 feeType;       // bitmask: 1=tuition, 2=hostel, 4=exam
        uint256 recordedAt;
    }

    // ── Functions ────────────────────────────────────────────────────────────────

    function recordPayment(
        bytes32 studentHash,
        bytes32 receiptHash,
        uint256 amount,
        uint256 semester,
        uint256 feeType
    ) external;

    function setFeeStructure(
        uint256 semester,
        uint256 tuitionFee,
        uint256 hostingFee,
        uint256 examFee
    ) external;

    function deactivateFeeStructure(uint256 semester) external;

    function checkpoint() external;

    function getPaymentRecord(bytes32 receiptHash) external view returns (PaymentRecord memory);

    function hasStudentPaidSemester(bytes32 studentHash, uint256 semester) external view returns (bool);

    function getFeeStructure(uint256 semester) external view returns (FeeStructure memory);

    function getTotalVolume() external view returns (uint256);

    function getTotalRecords() external view returns (uint256);
}
