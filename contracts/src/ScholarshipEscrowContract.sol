// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {IScholarshipEscrow} from "./interfaces/IScholarshipEscrow.sol";
import {AttestationLib} from "./libraries/AttestationLib.sol";

contract ScholarshipEscrowContract is IScholarshipEscrow, AccessControl, ReentrancyGuard {
    bytes32 public constant ADMIN_ROLE = DEFAULT_ADMIN_ROLE;
    bytes32 public constant ATTESTOR_ROLE = keccak256("ATTESTOR_ROLE");
    bytes32 public constant SPONSOR_ROLE = keccak256("SPONSOR_ROLE");

    uint256 private _nextFundId = 1;
    mapping(uint256 => ScholarshipFund) public funds;
    
    // fundId => studentHash => trancheIndex => released
    mapping(uint256 => mapping(bytes32 => mapping(uint256 => bool))) public hasReleased;

    // We store the attestor address directly for EIP-712 recovery verification.
    address public trustedAttestor;

    // EIP-712 Domain Separator
    bytes32 public immutable DOMAIN_SEPARATOR;

    constructor(address admin, address attestor) {
        _grantRole(ADMIN_ROLE, admin);
        _grantRole(ATTESTOR_ROLE, attestor);
        trustedAttestor = attestor;

        DOMAIN_SEPARATOR = AttestationLib.computeDomainSeparator(
            "UniTreasury Scholarship",
            "1",
            block.chainid,
            address(this)
        );
    }

    function createFund(
        address sponsor,
        uint256 trancheCount,
        uint256 trancheAmount
    ) external payable onlyRole(ADMIN_ROLE) returns (uint256) {
        require(trancheCount > 0, "Invalid tranche count");
        require(trancheAmount > 0, "Invalid tranche amount");
        require(msg.value > 0, "Must deposit funds");
        // For simplicity, msg.value should be evenly divisible or at least enough for some students.

        uint256 fundId = _nextFundId++;
        
        funds[fundId] = ScholarshipFund({
            sponsor: sponsor,
            totalAmount: msg.value,
            releasedAmount: 0,
            trancheCount: trancheCount,
            trancheAmount: trancheAmount,
            paused: false
        });

        emit FundCreated(fundId, sponsor, msg.value, trancheCount, trancheAmount);
        return fundId;
    }

    function releaseTranche(
        uint256 fundId,
        bytes32 studentHash,
        uint256 trancheIndex,
        address recipient,
        bytes calldata signature
    ) external nonReentrant {
        ScholarshipFund storage fund = funds[fundId];
        require(fund.totalAmount > 0, "Fund does not exist");
        require(!fund.paused, "Fund is paused");
        require(trancheIndex < fund.trancheCount, "Invalid tranche index");
        require(!hasReleased[fundId][studentHash][trancheIndex], "Tranche already released");
        require(fund.totalAmount - fund.releasedAmount >= fund.trancheAmount, "Insufficient funds");

        // Verify the EIP-712 signature
        uint256 nonce = AttestationLib.computeNonce(fundId, studentHash, trancheIndex);
        bytes32 structHash = AttestationLib.hashTrancheRelease(fundId, studentHash, trancheIndex, recipient, nonce);
        address signer = AttestationLib.recoverSigner(DOMAIN_SEPARATOR, structHash, signature);
        
        require(signer == trustedAttestor, "Invalid or unauthorized attestation");
        require(hasRole(ATTESTOR_ROLE, signer), "Signer lacks attestor role");

        hasReleased[fundId][studentHash][trancheIndex] = true;
        fund.releasedAmount += fund.trancheAmount;

        (bool success, ) = recipient.call{value: fund.trancheAmount}("");
        require(success, "ETH transfer failed");

        emit TrancheReleased(fundId, studentHash, trancheIndex, fund.trancheAmount, recipient);
    }

    function pauseFund(uint256 fundId) external onlyRole(ADMIN_ROLE) {
        require(funds[fundId].totalAmount > 0, "Fund does not exist");
        funds[fundId].paused = true;
    }

    function unpauseFund(uint256 fundId) external onlyRole(ADMIN_ROLE) {
        require(funds[fundId].totalAmount > 0, "Fund does not exist");
        funds[fundId].paused = false;
    }

    function setTrustedAttestor(address newAttestor) external onlyRole(ADMIN_ROLE) {
        require(newAttestor != address(0), "Invalid attestor");
        _revokeRole(ATTESTOR_ROLE, trustedAttestor);
        _grantRole(ATTESTOR_ROLE, newAttestor);
        trustedAttestor = newAttestor;
    }
}
