// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IFeeRegistry} from "./interfaces/IFeeRegistry.sol";

contract FeeRegistryContract is IFeeRegistry, AccessControl {
    bytes32 public constant ADMIN_ROLE = DEFAULT_ADMIN_ROLE;
    bytes32 public constant RECORDER_ROLE = keccak256("RECORDER_ROLE");

    IERC20 public immutable usdcToken;
    mapping(bytes32 => mapping(uint256 => uint256)) private _fees;

    constructor(address admin, address recorder, address _usdcToken) {
        _grantRole(ADMIN_ROLE, admin);
        _grantRole(RECORDER_ROLE, recorder);
        usdcToken = IERC20(_usdcToken);
    }

    function recordFee(bytes32 studentHash, uint256 termId, uint256 amount) external override onlyRole(RECORDER_ROLE) {
        require(amount > 0, "Amount must be > 0");
        require(usdcToken.transferFrom(msg.sender, address(this), amount), "USDC transfer failed");
        
        _fees[studentHash][termId] += amount;
        emit FeeRecorded(studentHash, termId, amount);
    }

    function getFee(bytes32 studentHash, uint256 termId) external view override returns (uint256) {
        return _fees[studentHash][termId];
    }
}
