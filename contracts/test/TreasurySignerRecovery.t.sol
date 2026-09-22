// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/TreasuryContract.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract MockUSDC is ERC20 {
    constructor() ERC20("Mock USDC", "USDC") {
        _mint(msg.sender, 1000000 * 10**6);
    }
}

contract TreasurySignerRecoveryTest is Test {
    TreasuryContract treasury;
    MockUSDC usdc;

    address admin = address(1);
    address approver1 = address(2);
    address approver2 = address(3);
    address approver3 = address(4); // compromised
    address newApprover = address(5); // replacement

    function setUp() public {
        usdc = new MockUSDC();
        
        address[] memory approvers = new address[](3);
        approvers[0] = approver1;
        approvers[1] = approver2;
        approvers[2] = approver3;

        // 3 signers, 2 required approvals
        treasury = new TreasuryContract(admin, approvers, 2, 10000 * 10**6, address(usdc), address(0), address(0), 48 hours);
    }

    function testSignerRecovery() public {
        // Assume approver3 lost their key. 
        // Approver1 and Approver2 vote to replace approver3 with newApprover.

        // 1. Approver1 proposes replacement
        vm.prank(approver1);
        uint256 proposalId = treasury.proposeSignerChange(approver3, newApprover, 2); // 2 = Replace

        // Approver1 also approves it
        vm.prank(approver1);
        treasury.approveSignerChange(proposalId);

        // 2. Approver2 approves replacement (since required approvals = 2)
        vm.prank(approver2);
        treasury.approveSignerChange(proposalId);

        // 3. Anyone with EXECUTOR_ROLE can execute. Wait, in constructor, admin gets EXECUTOR_ROLE.
        // Let's assume admin executes it.
        vm.prank(admin);
        treasury.executeSignerChange(proposalId);

        // Assert approver3 no longer has APPROVER_ROLE
        assertFalse(treasury.hasRole(treasury.APPROVER_ROLE(), approver3));
        
        // Assert newApprover HAS the APPROVER_ROLE
        assertTrue(treasury.hasRole(treasury.APPROVER_ROLE(), newApprover));
    }
}
