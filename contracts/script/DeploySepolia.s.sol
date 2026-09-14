// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Script, console} from "forge-std/Script.sol";
import {TreasuryContract} from "../src/TreasuryContract.sol";
import {FeeRegistryContract} from "../src/FeeRegistryContract.sol";

/// @notice Sepolia testnet deployment script.
///         Usage: make deploy-sepolia (see root Makefile)
///         Required env vars: SEPOLIA_RPC_URL, DEPLOYER_PRIVATE_KEY, ETHERSCAN_API_KEY
///         Optional env vars: APPROVER_2, APPROVER_3, RECORDER_ADDRESS
contract DeploySepoliaScript is Script {
    function run() external {
        uint256 deployerKey = vm.envUint("DEPLOYER_PRIVATE_KEY");
        address deployer = vm.addr(deployerKey);
        address recorder = vm.envOr("RECORDER_ADDRESS", deployer);

        // Sepolia: use 2-of-3 approvals
        address[] memory approvers = new address[](3);
        approvers[0] = deployer;
        approvers[1] = vm.envOr("APPROVER_2", deployer);
        approvers[2] = vm.envOr("APPROVER_3", deployer);

        uint256 requiredApprovals = 2;
        uint256 dailyLimit = 10 ether; // conservative for testnet

        console.log("=== Deploying to Sepolia ===");
        console.log("Deployer:", deployer);
        console.log("Recorder:", recorder);
        console.log("Required approvals:", requiredApprovals);
        console.log("Daily limit (wei):", dailyLimit);

        vm.startBroadcast(deployerKey);

        TreasuryContract treasury = new TreasuryContract(
            deployer,
            approvers,
            requiredApprovals,
            dailyLimit
        );
        console.log("TreasuryContract:", address(treasury));

        FeeRegistryContract feeRegistry = new FeeRegistryContract(deployer, recorder);
        console.log("FeeRegistryContract:", address(feeRegistry));

        vm.stopBroadcast();

        console.log("=== Deployment complete ===");
        console.log("Verify on Etherscan: https://sepolia.etherscan.io/address/", address(treasury));
    }
}
