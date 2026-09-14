// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IScholarshipEscrow} from "./interfaces/IScholarshipEscrow.sol";
import {AttestationLib} from "./libraries/AttestationLib.sol";

contract ScholarshipEscrowContract is IScholarshipEscrow, AccessControl, ReentrancyGuard {
    bytes32 public constant ADMIN_ROLE = DEFAULT_ADMIN_ROLE;
    bytes32 public constant ATTESTOR_ROLE = keccak256("ATTESTOR_ROLE");
    bytes32 public constant SPONSOR_ROLE = keccak256("SPONSOR_ROLE");

    IERC20 public immutable usdcToken;

    uint256 private _nextFundId = 1;
    mapping(uint256 => ScholarshipFund) public funds;
    
    // fundId => studentHash => trancheIndex => released
    mapping(uint256 => mapping(bytes32 => mapping(uint256 => bool))) public hasReleased;

    address public trustedAttestor;
    bytes32 public immutable DOMAIN_SEPARATOR;

    constructor(address admin, address attestor, address _usdcToken) {
        _grantRole(ADMIN_ROLE, admin);
        _grantRole(ATTESTOR_ROLE, attestor);
        trustedAttestor = attestor;
        usdcToken = IERC20(_usdcToken);

        DOMAIN_SEPARATOR = AttestationLib.domainSeparator(address(this));
    }

    function createFund(
        address sponsor,
        uint256 totalAmount,
        uint256 trancheCount,
        uint256 trancheAmount
    ) external onlyRole(ADMIN_ROLE) returns (uint256) {
        require(trancheCount > 0, "Invalid tranche count");
        require(trancheAmount > 0, "Invalid tranche amount");
        require(totalAmount > 0, "Must deposit funds");
        
        require(usdcToken.transferFrom(msg.sender, address(this), totalAmount), "USDC transfer failed");

        uint256 fundId = _nextFundId++;
        
        funds[fundId] = ScholarshipFund({
            sponsor: sponsor,
            totalAmount: totalAmount,
            releasedAmount: 0,
            trancheCount: trancheCount,
            trancheAmount: trancheAmount,
            paused: false
        });

        emit FundCreated(fundId, sponsor, totalAmount, trancheCount, trancheAmount);
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

        uint256 nonce = AttestationLib.computeNonce(fundId, studentHash, trancheIndex);
        
        // Construct the struct inside memory manually to avoid struct definition mismatch with library
        AttestationLib.TrancheAttestation memory att = AttestationLib.TrancheAttestation({
            fundId: fundId,
            studentHash: studentHash,
            trancheIndex: trancheIndex,
            recipient: recipient,
            nonce: nonce
        });

        address signer = AttestationLib.recoverSigner(att, DOMAIN_SEPARATOR, signature);
        
        require(signer == trustedAttestor, "Invalid or unauthorized attestation");
        require(hasRole(ATTESTOR_ROLE, signer), "Signer lacks attestor role");

        hasReleased[fundId][studentHash][trancheIndex] = true;
        fund.releasedAmount += fund.trancheAmount;

        require(usdcToken.transfer(recipient, fund.trancheAmount), "USDC transfer failed");

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
