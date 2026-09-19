// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {MockAToken} from "./MockAToken.sol";

contract MockLendingPool {
    mapping(address => address) public reserveToAToken;

    function initReserve(address asset, address aToken) external {
        reserveToAToken[asset] = aToken;
    }

    // Aave V3 supply signature
    function supply(
        address asset,
        uint256 amount,
        address onBehalfOf,
        uint16 /* referralCode */
    ) external {
        address aTokenAddress = reserveToAToken[asset];
        require(aTokenAddress != address(0), "Reserve not initialized");

        // Transfer underlying asset from user to pool
        bool success = IERC20(asset).transferFrom(msg.sender, address(this), amount);
        require(success, "Transfer failed");

        // Mint aTokens to onBehalfOf
        MockAToken(aTokenAddress).mint(onBehalfOf, amount);
    }

    // Aave V3 withdraw signature
    function withdraw(
        address asset,
        uint256 amount,
        address to
    ) external returns (uint256) {
        address aTokenAddress = reserveToAToken[asset];
        require(aTokenAddress != address(0), "Reserve not initialized");

        // Burn aTokens from the caller
        MockAToken(aTokenAddress).burn(msg.sender, amount);

        // Transfer underlying asset to the specified address
        bool success = IERC20(asset).transfer(to, amount);
        require(success, "Transfer failed");

        return amount;
    }
}
