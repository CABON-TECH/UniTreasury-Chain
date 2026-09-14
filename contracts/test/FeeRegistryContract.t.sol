// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Test, console} from "forge-std/Test.sol";
import {FeeRegistryContract} from "../src/FeeRegistryContract.sol";
import {IFeeRegistry} from "../src/interfaces/IFeeRegistry.sol";

/// @notice Full unit test suite for FeeRegistryContract.
///         Covers: payment recording, deduplication, fee structures, bitmask validation,
///         student semester tracking, checkpoint, reconciliation totals.
contract FeeRegistryContractTest is Test {
    // Redeclare interface events — required by Solidity/Forge to use `emit` in expectEmit
    event ReconciliationCheckpoint(uint256 totalRecords, uint256 totalVolume, uint256 blockNumber);

    FeeRegistryContract public registry;

    address admin = makeAddr("admin");
    address recorder = makeAddr("recorder");
    address stranger = makeAddr("stranger");

    uint256 constant SEMESTER_1 = 20241;
    uint256 constant TUITION = 1000 ether;
    uint256 constant HOSTEL = 300 ether;
    uint256 constant EXAM = 50 ether;

    bytes32 constant STUDENT_HASH = keccak256("student-id-001");
    bytes32 constant RECEIPT_1 = keccak256("receipt-001");
    bytes32 constant RECEIPT_2 = keccak256("receipt-002");

    function setUp() public {
        registry = new FeeRegistryContract(admin, recorder);

        // Set fee structure for semester 1
        vm.prank(admin);
        registry.setFeeStructure(SEMESTER_1, TUITION, HOSTEL, EXAM);
    }

    // ── Deployment ────────────────────────────────────────────────────────────────

    function test_deployment_roles() public view {
        assertTrue(registry.hasRole(registry.FEE_ADMIN_ROLE(), admin));
        assertTrue(registry.hasRole(registry.RECORDER_ROLE(), recorder));
        assertFalse(registry.hasRole(registry.RECORDER_ROLE(), stranger));
    }

    // ── Fee Structures ────────────────────────────────────────────────────────────

    function test_setFeeStructure_happy_path() public view {
        IFeeRegistry.FeeStructure memory fs = registry.getFeeStructure(SEMESTER_1);
        assertEq(fs.tuitionFee, TUITION);
        assertEq(fs.hostingFee, HOSTEL);
        assertEq(fs.examFee, EXAM);
        assertTrue(fs.active);
    }

    function test_setFeeStructure_revert_not_admin() public {
        vm.prank(stranger);
        vm.expectRevert();
        registry.setFeeStructure(SEMESTER_1, TUITION, HOSTEL, EXAM);
    }

    function test_setFeeStructure_revert_zero_semester() public {
        vm.prank(admin);
        vm.expectRevert(abi.encodeWithSelector(IFeeRegistry.FeeRegistry__InvalidSemester.selector, 0));
        registry.setFeeStructure(0, TUITION, HOSTEL, EXAM);
    }

    function test_deactivateFeeStructure() public {
        vm.prank(admin);
        registry.deactivateFeeStructure(SEMESTER_1);
        assertFalse(registry.getFeeStructure(SEMESTER_1).active);
    }

    // ── Payment Recording ─────────────────────────────────────────────────────────

    function test_recordPayment_tuition_only() public {
        // feeType = 1 (tuition only) → amount must equal TUITION
        vm.prank(recorder);
        registry.recordPayment(STUDENT_HASH, RECEIPT_1, TUITION, SEMESTER_1, 1);

        IFeeRegistry.PaymentRecord memory rec = registry.getPaymentRecord(RECEIPT_1);
        assertEq(rec.studentHash, STUDENT_HASH);
        assertEq(rec.amount, TUITION);
        assertEq(rec.semester, SEMESTER_1);
        assertEq(rec.feeType, 1);
        assertGt(rec.recordedAt, 0);
    }

    function test_recordPayment_combined_tuition_and_exam() public {
        // feeType = 5 (1 + 4 = tuition + exam)
        uint256 combined = TUITION + EXAM;
        vm.prank(recorder);
        registry.recordPayment(STUDENT_HASH, RECEIPT_1, combined, SEMESTER_1, 5);

        assertEq(registry.getPaymentRecord(RECEIPT_1).amount, combined);
    }

    function test_recordPayment_marks_student_as_paid() public {
        assertFalse(registry.hasStudentPaidSemester(STUDENT_HASH, SEMESTER_1));

        vm.prank(recorder);
        registry.recordPayment(STUDENT_HASH, RECEIPT_1, TUITION, SEMESTER_1, 1);

        assertTrue(registry.hasStudentPaidSemester(STUDENT_HASH, SEMESTER_1));
    }

    function test_recordPayment_increments_totals() public {
        vm.prank(recorder);
        registry.recordPayment(STUDENT_HASH, RECEIPT_1, TUITION, SEMESTER_1, 1);

        assertEq(registry.getTotalRecords(), 1);
        assertEq(registry.getTotalVolume(), TUITION);

        bytes32 student2 = keccak256("student-002");
        vm.prank(recorder);
        registry.recordPayment(student2, RECEIPT_2, TUITION, SEMESTER_1, 1);

        assertEq(registry.getTotalRecords(), 2);
        assertEq(registry.getTotalVolume(), TUITION * 2);
    }

    function test_recordPayment_revert_not_recorder() public {
        vm.prank(stranger);
        vm.expectRevert();
        registry.recordPayment(STUDENT_HASH, RECEIPT_1, TUITION, SEMESTER_1, 1);
    }

    function test_recordPayment_revert_duplicate_receipt() public {
        vm.prank(recorder);
        registry.recordPayment(STUDENT_HASH, RECEIPT_1, TUITION, SEMESTER_1, 1);

        // Same receiptHash again
        vm.prank(recorder);
        vm.expectRevert(
            abi.encodeWithSelector(IFeeRegistry.FeeRegistry__DuplicateReceipt.selector, RECEIPT_1)
        );
        registry.recordPayment(STUDENT_HASH, RECEIPT_1, TUITION, SEMESTER_1, 1);
    }

    function test_recordPayment_revert_zero_student_hash() public {
        vm.prank(recorder);
        vm.expectRevert(IFeeRegistry.FeeRegistry__ZeroStudentHash.selector);
        registry.recordPayment(bytes32(0), RECEIPT_1, TUITION, SEMESTER_1, 1);
    }

    function test_recordPayment_revert_amount_mismatch() public {
        // feeType = 1 (tuition only) but amount is wrong
        vm.prank(recorder);
        vm.expectRevert(
            abi.encodeWithSelector(IFeeRegistry.FeeRegistry__AmountMismatch.selector, 500 ether, TUITION)
        );
        registry.recordPayment(STUDENT_HASH, RECEIPT_1, 500 ether, SEMESTER_1, 1);
    }

    function test_recordPayment_no_validation_when_fee_structure_inactive() public {
        // Deactivate fee structure — any amount should be accepted
        vm.prank(admin);
        registry.deactivateFeeStructure(SEMESTER_1);

        vm.prank(recorder);
        registry.recordPayment(STUDENT_HASH, RECEIPT_1, 12345 ether, SEMESTER_1, 1);
        assertEq(registry.getPaymentRecord(RECEIPT_1).amount, 12345 ether);
    }

    function test_recordPayment_revert_invalid_semester() public {
        vm.prank(recorder);
        vm.expectRevert(abi.encodeWithSelector(IFeeRegistry.FeeRegistry__InvalidSemester.selector, 0));
        registry.recordPayment(STUDENT_HASH, RECEIPT_1, TUITION, 0, 1);
    }

    // ── Checkpoint ────────────────────────────────────────────────────────────────

    function test_checkpoint_emits_event() public {
        vm.prank(recorder);
        registry.recordPayment(STUDENT_HASH, RECEIPT_1, TUITION, SEMESTER_1, 1);

        vm.expectEmit(false, false, false, true);
        emit ReconciliationCheckpoint(1, TUITION, block.number);

        vm.prank(recorder);
        registry.checkpoint();
    }

    function test_checkpoint_revert_not_recorder() public {
        vm.prank(stranger);
        vm.expectRevert();
        registry.checkpoint();
    }

    // ── Multi-student ─────────────────────────────────────────────────────────────

    function test_multiple_students_independent() public {
        bytes32 s1 = keccak256("s1");
        bytes32 s2 = keccak256("s2");
        bytes32 r1 = keccak256("r1");
        bytes32 r2 = keccak256("r2");

        vm.startPrank(recorder);
        registry.recordPayment(s1, r1, TUITION, SEMESTER_1, 1);
        registry.recordPayment(s2, r2, TUITION, SEMESTER_1, 1);
        vm.stopPrank();

        assertTrue(registry.hasStudentPaidSemester(s1, SEMESTER_1));
        assertTrue(registry.hasStudentPaidSemester(s2, SEMESTER_1));
        assertFalse(registry.hasStudentPaidSemester(keccak256("s3"), SEMESTER_1));
    }

    // ── Fuzz ──────────────────────────────────────────────────────────────────────

    function testFuzz_recordPayment_unique_receipts(uint8 n) public {
        vm.assume(n > 0 && n < 20);

        // Deactivate fee structure to allow arbitrary amounts
        vm.prank(admin);
        registry.deactivateFeeStructure(SEMESTER_1);

        for (uint8 i = 0; i < n; i++) {
            bytes32 sh = keccak256(abi.encode("student", i));
            bytes32 rh = keccak256(abi.encode("receipt", i));
            vm.prank(recorder);
            registry.recordPayment(sh, rh, 100 ether, SEMESTER_1, 0);
        }

        assertEq(registry.getTotalRecords(), n);
    }
}
