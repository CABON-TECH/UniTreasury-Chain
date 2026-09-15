// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Script, console} from "forge-std/Script.sol";
import {TreasuryContract} from "../src/TreasuryContract.sol";
import {FeeRegistryContract} from "../src/FeeRegistryContract.sol";
import {ScholarshipEscrowContract} from "../src/ScholarshipEscrowContract.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
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
        usdc.transfer(address(treasury), 500_000 * 10**18);

        FeeRegistryContract feeRegistry = new FeeRegistryContract(deployer, deployer, address(usdc));
        console.log("FEE_REGISTRY_CONTRACT_ADDRESS=", address(feeRegistry));

        ScholarshipEscrowContract logic = new ScholarshipEscrowContract();
        bytes memory data = abi.encodeWithSelector(ScholarshipEscrowContract.initialize.selector, deployer, deployer, address(usdc));
        ERC1967Proxy proxy = new ERC1967Proxy(address(logic), data);
        ScholarshipEscrowContract escrow = ScholarshipEscrowContract(address(proxy));
        console.log("SCHOLARSHIP_ESCROW_CONTRACT_ADDRESS=", address(escrow));

        vm.stopBroadcast();
    }
}
