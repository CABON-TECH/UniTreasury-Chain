// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IFeeRegistry {
    event FeeRecorded(bytes32 indexed studentHash, uint256 indexed termId, uint256 amount);
    function recordFee(bytes32 studentHash, uint256 termId, uint256 amount) external;
    function getFee(bytes32 studentHash, uint256 termId) external view returns (uint256);
}
