// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Script, console} from "forge-std/Script.sol";
import {TreasuryContract} from "../src/TreasuryContract.sol";
import {FeeRegistryContract} from "../src/FeeRegistryContract.sol";

/// @notice Local (Anvil) deployment script.
///         Usage: forge script script/Deploy.s.sol --rpc-url http://localhost:8545 --broadcast
contract DeployScript is Script {
    function run() external {
        uint256 deployerKey = vm.envUint("DEPLOYER_PRIVATE_KEY");
        address deployer = vm.addr(deployerKey);
        address recorder = vm.envOr("RECORDER_ADDRESS", deployer);

        console.log("Deploying from:", deployer);
        console.log("Recorder address:", recorder);

        // Build approvers list — for local testing, deployer is the only approver
        address[] memory approvers = new address[](3);
        approvers[0] = deployer;
        approvers[1] = vm.envOr("APPROVER_2", deployer);
        approvers[2] = vm.envOr("APPROVER_3", deployer);

        uint256 requiredApprovals = 1; // local: 1-of-3 for easy testing
        uint256 dailyLimit = 1000 ether;

        vm.startBroadcast(deployerKey);

        TreasuryContract treasury = new TreasuryContract(
            deployer,
            approvers,
            requiredApprovals,
            dailyLimit
        );
        console.log("TreasuryContract deployed:", address(treasury));

        FeeRegistryContract feeRegistry = new FeeRegistryContract(deployer, recorder);
        console.log("FeeRegistryContract deployed:", address(feeRegistry));

        vm.stopBroadcast();

        // Write addresses to a file for the Go backend to pick up
        string memory out = string(abi.encodePacked(
            "TREASURY_CONTRACT_ADDRESS=", vm.toString(address(treasury)), "\n",
            "FEE_REGISTRY_CONTRACT_ADDRESS=", vm.toString(address(feeRegistry)), "\n"
        ));
        vm.writeFile("../backend/.contract-addresses.env", out);
        console.log("Addresses written to backend/.contract-addresses.env");
    }
}
