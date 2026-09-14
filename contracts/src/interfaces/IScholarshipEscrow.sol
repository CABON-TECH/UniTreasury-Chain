// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IScholarshipEscrow {
    event FundCreated(uint256 indexed fundId, address indexed sponsor, uint256 totalAmount, uint256 trancheCount, uint256 trancheAmount);
    event TrancheReleased(uint256 indexed fundId, bytes32 indexed studentHash, uint256 trancheIndex, uint256 amount, address recipient);

    struct ScholarshipFund {
        address sponsor;
        uint256 totalAmount;
        uint256 releasedAmount;
        uint256 trancheCount;
        uint256 trancheAmount;
        bool paused;
    }

    function createFund(address sponsor, uint256 totalAmount, uint256 trancheCount, uint256 trancheAmount) external returns (uint256);
    function releaseTranche(uint256 fundId, bytes32 studentHash, uint256 trancheIndex, address recipient, bytes calldata attestation) external;
    function pauseFund(uint256 fundId) external;
    function unpauseFund(uint256 fundId) external;
    function setTrustedAttestor(address newAttestor) external;
}
