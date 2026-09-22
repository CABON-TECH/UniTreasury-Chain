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
    function validatePaymasterUserOp(UserOperation calldata op, bytes32, uint256) external view returns (bytes memory, uint256) {
        require(allowedTargets[op.sender], "Target not allowed");
        return ("", 0);
    }
}
