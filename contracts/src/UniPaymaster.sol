// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {UserOperation} from "./MockEntryPoint.sol";

contract UniPaymaster {
    address public entryPoint;
    mapping(address => bool) public allowedTargets;

    constructor(address _entryPoint) {
        entryPoint = _entryPoint;
    }

    function setAllowedTarget(address target, bool allowed) external {
        allowedTargets[target] = allowed;
    }

    // Mock validate method
    function validatePaymasterUserOp(UserOperation calldata op, bytes32, uint256) external view returns (bytes memory, uint256) {
        // Normally validates that the target is allowed and there is enough ETH deposited
        require(allowedTargets[op.sender], "Target not allowed");
        return ("", 0);
    }
}
