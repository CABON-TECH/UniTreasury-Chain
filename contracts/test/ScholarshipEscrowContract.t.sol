// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Test, console} from "forge-std/Test.sol";
import {ScholarshipEscrowContract} from "../src/ScholarshipEscrowContract.sol";
import {AttestationLib} from "../src/libraries/AttestationLib.sol";

contract ScholarshipEscrowContractTest is Test {
    ScholarshipEscrowContract public escrow;

    address admin = makeAddr("admin");
    address sponsor = makeAddr("sponsor");
    address recipient = makeAddr("recipient");
    
    uint256 attestorPk;
    address attestor;

    bytes32 constant STUDENT_HASH = keccak256("student-001");

    event FundCreated(uint256 indexed fundId, address indexed sponsor, uint256 totalAmount, uint256 trancheCount, uint256 trancheAmount);
    event TrancheReleased(uint256 indexed fundId, bytes32 indexed studentHash, uint256 trancheIndex, uint256 amount, address recipient);

    function setUp() public {
        (attestor, attestorPk) = makeAddrAndKey("attestor");
        
        vm.prank(admin);
        escrow = new ScholarshipEscrowContract(admin, attestor);
    }

    function test_createFund() public {
        vm.deal(admin, 10 ether);
        
        vm.prank(admin);
        vm.expectEmit(true, true, false, true);
        emit FundCreated(1, sponsor, 10 ether, 3, 1 ether);
        
        uint256 fundId = escrow.createFund{value: 10 ether}(sponsor, 3, 1 ether);
        assertEq(fundId, 1);
        
        (address fSponsor, uint256 totalAmount, uint256 releasedAmount, uint256 trancheCount, uint256 trancheAmount, bool paused) = escrow.funds(fundId);
        assertEq(fSponsor, sponsor);
        assertEq(totalAmount, 10 ether);
        assertEq(releasedAmount, 0);
        assertEq(trancheCount, 3);
        assertEq(trancheAmount, 1 ether);
        assertFalse(paused);
    }

    function _signRelease(uint256 fundId, bytes32 studentHash, uint256 trancheIndex, address rec) internal view returns (bytes memory) {
        uint256 nonce = AttestationLib.computeNonce(fundId, studentHash, trancheIndex);
        bytes32 structHash = AttestationLib.hashTrancheRelease(fundId, studentHash, trancheIndex, rec, nonce);
        
        bytes32 digest = keccak256(
            abi.encodePacked(
                "\x19\x01",
                escrow.DOMAIN_SEPARATOR(),
                structHash
            )
        );

        (uint8 v, bytes32 r, bytes32 s) = vm.sign(attestorPk, digest);
        return abi.encodePacked(r, s, v);
    }

    function test_releaseTranche_success() public {
        vm.deal(admin, 10 ether);
        vm.prank(admin);
        uint256 fundId = escrow.createFund{value: 10 ether}(sponsor, 3, 1 ether);

        bytes memory sig = _signRelease(fundId, STUDENT_HASH, 0, recipient);

        vm.expectEmit(true, true, false, true);
        emit TrancheReleased(fundId, STUDENT_HASH, 0, 1 ether, recipient);

        uint256 balBefore = recipient.balance;
        
        // Anyone can submit the release transaction
        escrow.releaseTranche(fundId, STUDENT_HASH, 0, recipient, sig);

        assertEq(recipient.balance, balBefore + 1 ether);
        assertTrue(escrow.hasReleased(fundId, STUDENT_HASH, 0));
        
        (, , uint256 releasedAmount, , , ) = escrow.funds(fundId);
        assertEq(releasedAmount, 1 ether);
    }

    function test_releaseTranche_revert_invalidSignature() public {
        vm.deal(admin, 10 ether);
        vm.prank(admin);
        uint256 fundId = escrow.createFund{value: 10 ether}(sponsor, 3, 1 ether);

        // Sign with a bad key
        (address badAttestor, uint256 badPk) = makeAddrAndKey("badAttestor");
        
        uint256 nonce = AttestationLib.computeNonce(fundId, STUDENT_HASH, 0);
        bytes32 structHash = AttestationLib.hashTrancheRelease(fundId, STUDENT_HASH, 0, recipient, nonce);
        bytes32 digest = keccak256(abi.encodePacked("\x19\x01", escrow.DOMAIN_SEPARATOR(), structHash));
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(badPk, digest);
        bytes memory badSig = abi.encodePacked(r, s, v);

        vm.expectRevert("Invalid or unauthorized attestation");
        escrow.releaseTranche(fundId, STUDENT_HASH, 0, recipient, badSig);
    }

    function test_releaseTranche_revert_alreadyReleased() public {
        vm.deal(admin, 10 ether);
        vm.prank(admin);
        uint256 fundId = escrow.createFund{value: 10 ether}(sponsor, 3, 1 ether);

        bytes memory sig = _signRelease(fundId, STUDENT_HASH, 0, recipient);
        escrow.releaseTranche(fundId, STUDENT_HASH, 0, recipient, sig);

        vm.expectRevert("Tranche already released");
        escrow.releaseTranche(fundId, STUDENT_HASH, 0, recipient, sig);
    }
}
