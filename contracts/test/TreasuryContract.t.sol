// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Test, console} from "forge-std/Test.sol";
import {TreasuryContract} from "../src/TreasuryContract.sol";
import {ITreasury} from "../src/interfaces/ITreasury.sol";

/// @notice Full unit test suite for TreasuryContract.
///         Covers: deposit, propose, approve, execute, cancel, freeze/unfreeze,
///         daily limit, double-approval guard, insufficient-approval guard.
contract TreasuryContractTest is Test {
    TreasuryContract public treasury;

    address admin = makeAddr("admin");
    address proposer = makeAddr("proposer");
    address executor = makeAddr("executor");
    address approver1 = makeAddr("approver1");
    address approver2 = makeAddr("approver2");
    address approver3 = makeAddr("approver3");
    address recipient = makeAddr("recipient");
    address stranger = makeAddr("stranger");

    uint256 constant DAILY_LIMIT = 100 ether;
    uint256 constant REQUIRED_APPROVALS = 2;

    address[] approvers;

    function setUp() public {
        approvers.push(approver1);
        approvers.push(approver2);
        approvers.push(approver3);

        vm.startPrank(admin);
        treasury = new TreasuryContract(admin, approvers, REQUIRED_APPROVALS, DAILY_LIMIT);

        // Grant proposer and executor roles separately
        treasury.grantRole(treasury.PROPOSER_ROLE(), proposer);
        treasury.grantRole(treasury.EXECUTOR_ROLE(), executor);
        vm.stopPrank();

        // Fund the treasury
        vm.deal(address(treasury), 200 ether);
    }

    // ── Deployment ────────────────────────────────────────────────────────────────

    function test_deployment_roles() public view {
        assertTrue(treasury.hasRole(treasury.ADMIN_ROLE(), admin));
        assertTrue(treasury.hasRole(treasury.APPROVER_ROLE(), approver1));
        assertTrue(treasury.hasRole(treasury.APPROVER_ROLE(), approver2));
        assertTrue(treasury.hasRole(treasury.APPROVER_ROLE(), approver3));
        assertEq(treasury.requiredApprovals(), REQUIRED_APPROVALS);
        assertEq(treasury.dailyLimit(), DAILY_LIMIT);
    }

    function test_deployment_revert_zero_admin() public {
        vm.expectRevert(ITreasury.Treasury__ZeroAddress.selector);
        new TreasuryContract(address(0), approvers, 2, DAILY_LIMIT);
    }

    function test_deployment_revert_invalid_threshold() public {
        vm.expectRevert("TreasuryContract: invalid approval threshold");
        new TreasuryContract(admin, approvers, 0, DAILY_LIMIT);

        vm.expectRevert("TreasuryContract: invalid approval threshold");
        new TreasuryContract(admin, approvers, 4, DAILY_LIMIT); // > approvers.length
    }

    // ── Deposit ───────────────────────────────────────────────────────────────────

    function test_deposit_via_receive() public {
        uint256 before = treasury.getBalance();
        vm.deal(stranger, 1 ether);
        vm.prank(stranger);
        (bool ok,) = address(treasury).call{value: 1 ether}("");
        assertTrue(ok);
        assertEq(treasury.getBalance(), before + 1 ether);
    }

    function test_deposit_explicit() public {
        uint256 before = treasury.getBalance();
        vm.deal(stranger, 5 ether);
        vm.prank(stranger);
        treasury.deposit{value: 5 ether}();
        assertEq(treasury.getBalance(), before + 5 ether);
    }

    // ── Propose ───────────────────────────────────────────────────────────────────

    function test_proposeWithdrawal_happy_path() public {
        vm.prank(proposer);
        uint256 id = treasury.proposeWithdrawal(payable(recipient), 10 ether, "Q1 budget");

        assertEq(id, 1);
        ITreasury.WithdrawalProposal memory p = treasury.getProposal(id);
        assertEq(p.amount, 10 ether);
        assertEq(p.recipient, recipient);
        assertEq(p.proposer, proposer);
        assertEq(uint8(p.status), uint8(ITreasury.ProposalStatus.Pending));
        assertEq(p.approvalCount, 0);
    }

    function test_proposeWithdrawal_revert_not_proposer() public {
        vm.prank(stranger);
        vm.expectRevert();
        treasury.proposeWithdrawal(payable(recipient), 1 ether, "test");
    }

    function test_proposeWithdrawal_revert_zero_address() public {
        vm.prank(proposer);
        vm.expectRevert(ITreasury.Treasury__ZeroAddress.selector);
        treasury.proposeWithdrawal(payable(address(0)), 1 ether, "test");
    }

    function test_proposeWithdrawal_revert_zero_amount() public {
        vm.prank(proposer);
        vm.expectRevert(ITreasury.Treasury__ZeroAmount.selector);
        treasury.proposeWithdrawal(payable(recipient), 0, "test");
    }

    function test_proposeWithdrawal_revert_insufficient_balance() public {
        vm.prank(proposer);
        vm.expectRevert();
        treasury.proposeWithdrawal(payable(recipient), 9999 ether, "test");
    }

    function test_proposeWithdrawal_revert_frozen() public {
        vm.prank(admin);
        treasury.freeze();

        vm.prank(proposer);
        vm.expectRevert(ITreasury.Treasury__Frozen.selector);
        treasury.proposeWithdrawal(payable(recipient), 1 ether, "test");
    }

    // ── Approve ───────────────────────────────────────────────────────────────────

    function test_approveWithdrawal_happy_path() public {
        uint256 id = _proposeWithdrawal(10 ether);

        vm.prank(approver1);
        treasury.approveWithdrawal(id);
        assertEq(treasury.getProposal(id).approvalCount, 1);
        assertTrue(treasury.hasApproved(id, approver1));

        vm.prank(approver2);
        treasury.approveWithdrawal(id);
        assertEq(treasury.getProposal(id).approvalCount, 2);
    }

    function test_approveWithdrawal_revert_not_approver() public {
        uint256 id = _proposeWithdrawal(10 ether);
        vm.prank(stranger);
        vm.expectRevert();
        treasury.approveWithdrawal(id);
    }

    function test_approveWithdrawal_revert_double_approval() public {
        uint256 id = _proposeWithdrawal(10 ether);

        vm.prank(approver1);
        treasury.approveWithdrawal(id);

        vm.prank(approver1);
        vm.expectRevert(abi.encodeWithSelector(ITreasury.Treasury__AlreadyApproved.selector, id, approver1));
        treasury.approveWithdrawal(id);
    }

    function test_approveWithdrawal_revert_not_found() public {
        vm.prank(approver1);
        vm.expectRevert(abi.encodeWithSelector(ITreasury.Treasury__ProposalNotFound.selector, 999));
        treasury.approveWithdrawal(999);
    }

    function test_approveWithdrawal_revert_already_executed() public {
        uint256 id = _proposeWithdrawal(10 ether);
        _approve(id, approver1, approver2);
        vm.prank(executor);
        treasury.executeWithdrawal(id);

        vm.prank(approver3);
        vm.expectRevert(abi.encodeWithSelector(ITreasury.Treasury__ProposalNotPending.selector, id));
        treasury.approveWithdrawal(id);
    }

    // ── Execute ───────────────────────────────────────────────────────────────────

    function test_executeWithdrawal_happy_path() public {
        uint256 id = _proposeWithdrawal(10 ether);
        _approve(id, approver1, approver2);

        uint256 recipientBefore = recipient.balance;
        vm.prank(executor);
        treasury.executeWithdrawal(id);

        assertEq(recipient.balance, recipientBefore + 10 ether);
        assertEq(uint8(treasury.getProposal(id).status), uint8(ITreasury.ProposalStatus.Executed));
    }

    function test_executeWithdrawal_revert_frozen() public {
        uint256 id = _proposeWithdrawal(10 ether);
        _approve(id, approver1, approver2);

        vm.prank(admin);
        treasury.freeze();

        vm.prank(executor);
        vm.expectRevert(ITreasury.Treasury__Frozen.selector);
        treasury.executeWithdrawal(id);
    }

    function test_executeWithdrawal_revert_insufficient_approvals() public {
        uint256 id = _proposeWithdrawal(10 ether);

        // Only 1 approval, need 2
        vm.prank(approver1);
        treasury.approveWithdrawal(id);

        vm.prank(executor);
        vm.expectRevert(
            abi.encodeWithSelector(ITreasury.Treasury__InsufficientApprovals.selector, 1, REQUIRED_APPROVALS)
        );
        treasury.executeWithdrawal(id);
    }

    function test_executeWithdrawal_revert_daily_limit_exceeded() public {
        // Propose and execute up to daily limit (100 ETH)
        uint256 id1 = _proposeWithdrawal(60 ether);
        _approve(id1, approver1, approver2);
        vm.prank(executor);
        treasury.executeWithdrawal(id1);

        uint256 id2 = _proposeWithdrawal(50 ether);
        _approve(id2, approver1, approver2);

        vm.prank(executor);
        vm.expectRevert();
        treasury.executeWithdrawal(id2); // Would total 110 ETH > 100 ETH limit
    }

    function test_executeWithdrawal_revert_not_executor() public {
        uint256 id = _proposeWithdrawal(10 ether);
        _approve(id, approver1, approver2);

        vm.prank(stranger);
        vm.expectRevert();
        treasury.executeWithdrawal(id);
    }

    function test_executeWithdrawal_daily_limit_resets_next_day() public {
        // Execute up to near limit today
        uint256 id1 = _proposeWithdrawal(90 ether);
        _approve(id1, approver1, approver2);
        vm.prank(executor);
        treasury.executeWithdrawal(id1);

        assertEq(treasury.getDailyWithdrawnAmount(), 90 ether);

        // Advance to next UTC day
        vm.warp(block.timestamp + 1 days);

        // Re-fund treasury since balance dropped
        vm.deal(address(treasury), 200 ether);

        uint256 id2 = _proposeWithdrawal(90 ether);
        _approve(id2, approver1, approver2);
        vm.prank(executor);
        treasury.executeWithdrawal(id2); // Should succeed — fresh day
    }

    // ── Cancel ────────────────────────────────────────────────────────────────────

    function test_cancelWithdrawal_by_proposer() public {
        uint256 id = _proposeWithdrawal(10 ether);

        vm.prank(proposer);
        treasury.cancelWithdrawal(id);

        assertEq(uint8(treasury.getProposal(id).status), uint8(ITreasury.ProposalStatus.Cancelled));
    }

    function test_cancelWithdrawal_by_admin() public {
        uint256 id = _proposeWithdrawal(10 ether);

        vm.prank(admin);
        treasury.cancelWithdrawal(id);

        assertEq(uint8(treasury.getProposal(id).status), uint8(ITreasury.ProposalStatus.Cancelled));
    }

    function test_cancelWithdrawal_revert_by_stranger() public {
        uint256 id = _proposeWithdrawal(10 ether);

        vm.prank(stranger);
        vm.expectRevert(ITreasury.Treasury__NotAuthorized.selector);
        treasury.cancelWithdrawal(id);
    }

    function test_cancelWithdrawal_revert_already_executed() public {
        uint256 id = _proposeWithdrawal(10 ether);
        _approve(id, approver1, approver2);
        vm.prank(executor);
        treasury.executeWithdrawal(id);

        vm.prank(admin);
        vm.expectRevert(abi.encodeWithSelector(ITreasury.Treasury__ProposalNotPending.selector, id));
        treasury.cancelWithdrawal(id);
    }

    // ── Freeze / Unfreeze ─────────────────────────────────────────────────────────

    function test_freeze_and_unfreeze() public {
        assertFalse(treasury.isFrozen());

        vm.prank(admin);
        treasury.freeze();
        assertTrue(treasury.isFrozen());

        vm.prank(admin);
        treasury.unfreeze();
        assertFalse(treasury.isFrozen());
    }

    function test_freeze_revert_non_admin() public {
        vm.prank(stranger);
        vm.expectRevert();
        treasury.freeze();
    }

    // ── Admin Settings ────────────────────────────────────────────────────────────

    function test_setDailyLimit() public {
        vm.prank(admin);
        treasury.setDailyLimit(500 ether);
        assertEq(treasury.dailyLimit(), 500 ether);
    }

    function test_setRequiredApprovals() public {
        vm.prank(admin);
        treasury.setRequiredApprovals(3);
        assertEq(treasury.requiredApprovals(), 3);
    }

    // ── Fuzz ──────────────────────────────────────────────────────────────────────

    function testFuzz_deposit(uint96 amount) public {
        vm.assume(amount > 0);
        vm.deal(stranger, amount);
        vm.prank(stranger);
        treasury.deposit{value: amount}();
        assertTrue(treasury.getBalance() >= amount);
    }

    function testFuzz_proposeWithdrawal_amount_within_balance(uint96 amount) public {
        uint256 balance = treasury.getBalance();
        vm.assume(amount > 0 && amount <= balance);

        vm.prank(proposer);
        uint256 id = treasury.proposeWithdrawal(payable(recipient), amount, "fuzz");
        assertEq(treasury.getProposal(id).amount, amount);
    }

    // ── Helpers ───────────────────────────────────────────────────────────────────

    function _proposeWithdrawal(uint256 amount) internal returns (uint256 id) {
        vm.prank(proposer);
        id = treasury.proposeWithdrawal(payable(recipient), amount, "test purpose");
    }

    function _approve(uint256 id, address a, address b) internal {
        vm.prank(a);
        treasury.approveWithdrawal(id);
        vm.prank(b);
        treasury.approveWithdrawal(id);
    }
}
