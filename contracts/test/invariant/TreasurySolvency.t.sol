// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Test} from "forge-std/Test.sol";
import {StdInvariant} from "forge-std/StdInvariant.sol";
import {TreasuryContract} from "../../src/TreasuryContract.sol";
import {ITreasury} from "../../src/interfaces/ITreasury.sol";

/// @notice Invariant test: sum of all executed withdrawals must never exceed total deposits.
///         Foundry's invariant fuzzer will call handler functions in random order,
///         then assert the invariant after each sequence.
contract TreasurySolvencyInvariantTest is StdInvariant, Test {
    TreasuryContract public treasury;
    TreasurySolvencyHandler public handler;

    function setUp() public {
        address admin = makeAddr("admin");
        address[] memory approvers = new address[](3);
        approvers[0] = makeAddr("a1");
        approvers[1] = makeAddr("a2");
        approvers[2] = makeAddr("a3");

        vm.prank(admin);
        treasury = new TreasuryContract(admin, approvers, 2, 1000 ether);

        handler = new TreasurySolvencyHandler(treasury, admin, approvers);
        targetContract(address(handler));
    }

    /// @notice Core invariant: the contract balance must equal totalDeposited - totalWithdrawn
    function invariant_solvency() public view {
        assertEq(
            address(treasury).balance,
            handler.totalDeposited() - handler.totalWithdrawn(),
            "Invariant violated: balance != deposits - withdrawals"
        );
    }

    /// @notice Invariant: daily withdrawn must never exceed the daily limit
    function invariant_daily_limit() public view {
        assertLe(
            treasury.getDailyWithdrawnAmount(),
            treasury.dailyLimit(),
            "Invariant violated: daily limit exceeded"
        );
    }
}

/// @notice Handler contract drives the invariant fuzzer with bounded actions
contract TreasurySolvencyHandler is Test {
    TreasuryContract public treasury;
    address public admin;
    address[] public approvers;
    address public proposer;
    address public executor;

    uint256 public totalDeposited;
    uint256 public totalWithdrawn;

    uint256[] public pendingProposalIds;

    constructor(TreasuryContract _treasury, address _admin, address[] memory _approvers) {
        treasury = _treasury;
        admin = _admin;
        approvers = _approvers;

        proposer = makeAddr("handler_proposer");
        executor = makeAddr("handler_executor");

        vm.startPrank(_admin);
        treasury.grantRole(treasury.PROPOSER_ROLE(), proposer);
        treasury.grantRole(treasury.EXECUTOR_ROLE(), executor);
        vm.stopPrank();

        // Seed treasury
        vm.deal(address(treasury), 500 ether);
        totalDeposited = 500 ether;
    }

    function deposit(uint96 amount) public {
        vm.assume(amount > 0 && amount < 100 ether);
        vm.deal(address(this), amount);
        treasury.deposit{value: amount}();
        totalDeposited += amount;
    }

    function proposeWithdrawal(uint96 amount) public {
        uint256 bal = address(treasury).balance;
        vm.assume(amount > 0 && amount <= bal && amount <= treasury.dailyLimit());

        vm.prank(proposer);
        uint256 id = treasury.proposeWithdrawal(payable(makeAddr("recipient")), amount, "invariant");
        pendingProposalIds.push(id);
    }

    function approveAndExecute(uint256 seed) public {
        if (pendingProposalIds.length == 0) return;
        uint256 idx = seed % pendingProposalIds.length;
        uint256 id = pendingProposalIds[idx];

        // Try to get the proposal — skip if already executed/cancelled
        try treasury.getProposal(id) returns (ITreasury.WithdrawalProposal memory p) {
            if (p.status != ITreasury.ProposalStatus.Pending) return;

            vm.prank(approvers[0]);
            try treasury.approveWithdrawal(id) {} catch {}
            vm.prank(approvers[1]);
            try treasury.approveWithdrawal(id) {} catch {}

            uint256 before = address(treasury).balance;
            vm.prank(executor);
            try treasury.executeWithdrawal(id) {
                totalWithdrawn += before - address(treasury).balance;
                // Remove from pending
                pendingProposalIds[idx] = pendingProposalIds[pendingProposalIds.length - 1];
                pendingProposalIds.pop();
            } catch {}
        } catch {}
    }

    receive() external payable {}
}
