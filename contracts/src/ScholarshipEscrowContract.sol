// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {MerkleProof} from "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";
import {IScholarshipEscrow} from "./interfaces/IScholarshipEscrow.sol";

contract ScholarshipEscrowContract is IScholarshipEscrow, AccessControl, ReentrancyGuard {
    bytes32 public constant ADMIN_ROLE = DEFAULT_ADMIN_ROLE;
    bytes32 public constant ATTESTOR_ROLE = keccak256("ATTESTOR_ROLE");
    bytes32 public constant SPONSOR_ROLE = keccak256("SPONSOR_ROLE");

    IERC20 public immutable usdcToken;

    uint256 private _nextFundId = 1;
    mapping(uint256 => ScholarshipFund) public funds;
    
    // fundId => trancheIndex => merkleRoot
    mapping(uint256 => mapping(uint256 => bytes32)) public trancheRoots;
    
    // fundId => studentHash => trancheIndex => claimed
    mapping(uint256 => mapping(bytes32 => mapping(uint256 => bool))) public hasClaimed;

    address public trustedAttestor; // The entity allowed to publish roots

    constructor(address admin, address attestor, address _usdcToken) {
        _grantRole(ADMIN_ROLE, admin);
        _grantRole(ATTESTOR_ROLE, attestor);
        trustedAttestor = attestor;
        usdcToken = IERC20(_usdcToken);
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

    function publishTrancheRoot(
        uint256 fundId,
        uint256 trancheIndex,
        bytes32 merkleRoot
    ) external onlyRole(ATTESTOR_ROLE) {
        require(funds[fundId].totalAmount > 0, "Fund does not exist");
        require(!funds[fundId].paused, "Fund is paused");
        require(trancheIndex < funds[fundId].trancheCount, "Invalid tranche index");
        require(trancheRoots[fundId][trancheIndex] == bytes32(0), "Root already published");

        trancheRoots[fundId][trancheIndex] = merkleRoot;
        
        emit MerkleRootPublished(fundId, trancheIndex, merkleRoot);
    }

    function claimTranche(
        uint256 fundId,
        bytes32 studentHash,
        uint256 trancheIndex,
        address recipient,
        bytes32[] calldata merkleProof
    ) external nonReentrant {
        ScholarshipFund storage fund = funds[fundId];
        require(fund.totalAmount > 0, "Fund does not exist");
        require(!fund.paused, "Fund is paused");
        require(trancheIndex < fund.trancheCount, "Invalid tranche index");
        require(!hasClaimed[fundId][studentHash][trancheIndex], "Tranche already claimed");
        require(fund.totalAmount - fund.releasedAmount >= fund.trancheAmount, "Insufficient funds");

        bytes32 root = trancheRoots[fundId][trancheIndex];
        require(root != bytes32(0), "Tranche root not published yet");

        // Verify the merkle proof
        // Leaf = keccak256(abi.encodePacked(fundId, studentHash, trancheIndex, recipient, fund.trancheAmount))
        bytes32 leaf = keccak256(bytes.concat(keccak256(abi.encode(fundId, studentHash, trancheIndex, recipient, fund.trancheAmount))));
        
        require(MerkleProof.verify(merkleProof, root, leaf), "Invalid Merkle proof");

        hasClaimed[fundId][studentHash][trancheIndex] = true;
        fund.releasedAmount += fund.trancheAmount;

        require(usdcToken.transfer(recipient, fund.trancheAmount), "USDC transfer failed");

        emit TrancheClaimed(fundId, studentHash, trancheIndex, fund.trancheAmount, recipient);
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
