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
    address approver3 = address(4); 
    address newApprover = address(5); 
    function setUp() public {
        usdc = new MockUSDC();
        address[] memory approvers = new address[](3);
        approvers[0] = approver1;
        approvers[1] = approver2;
        approvers[2] = approver3;
        treasury = new TreasuryContract(admin, approvers, 2, 10000 * 10**6, address(usdc), address(0), address(0), 48 hours);
    }
    function testSignerRecovery() public {
        vm.prank(approver1);
        uint256 proposalId = treasury.proposeSignerChange(approver3, newApprover, 2); 
        vm.prank(approver1);
        treasury.approveSignerChange(proposalId);
        vm.prank(approver2);
        treasury.approveSignerChange(proposalId);
        vm.prank(admin);
        treasury.executeSignerChange(proposalId);
        assertFalse(treasury.hasRole(treasury.APPROVER_ROLE(), approver3));
        assertTrue(treasury.hasRole(treasury.APPROVER_ROLE(), newApprover));
    }
}
