// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Script, console} from "forge-std/Script.sol";
import {TreasuryContract} from "../src/TreasuryContract.sol";
import {FeeRegistryContract} from "../src/FeeRegistryContract.sol";
import {ScholarshipEscrowContract} from "../src/ScholarshipEscrowContract.sol";
import {MockUSDC} from "../src/MockUSDC.sol";

contract DeployAllScript is Script {
    function run() external {
        uint256 deployerKey = 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80; // default anvil account 0
        address deployer = vm.addr(deployerKey);
        
        vm.startBroadcast(deployerKey);

        MockUSDC usdc = new MockUSDC();
        usdc.mint(deployer, 1_000_000 * 10**18);
        console.log("MOCK_USDC_ADDRESS=", address(usdc));

        address[] memory approvers = new address[](1);
        approvers[0] = deployer;

        TreasuryContract treasury = new TreasuryContract(
            deployer,
            approvers,
            1,
            1000 * 10**18,
            address(usdc)
        );
        console.log("TREASURY_CONTRACT_ADDRESS=", address(treasury));

        FeeRegistryContract feeRegistry = new FeeRegistryContract(deployer, deployer, address(usdc));
        console.log("FEE_REGISTRY_CONTRACT_ADDRESS=", address(feeRegistry));

        ScholarshipEscrowContract escrow = new ScholarshipEscrowContract(deployer, deployer, address(usdc));
        console.log("SCHOLARSHIP_ESCROW_CONTRACT_ADDRESS=", address(escrow));

        vm.stopBroadcast();
    }
}
