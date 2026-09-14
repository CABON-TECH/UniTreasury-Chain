// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import {MockUSDC} from "../src/MockUSDC.sol";
import {ScholarshipEscrowContract} from "../src/ScholarshipEscrowContract.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

contract DeployUpgradeable is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);

        vm.startBroadcast(deployerPrivateKey);

        // 1. Deploy Mock USDC
        MockUSDC usdc = new MockUSDC();
        usdc.mint(deployer, 1_000_000 * 10**18);

        // 2. Deploy Logic Contract (Implementation)
        ScholarshipEscrowContract logic = new ScholarshipEscrowContract();

        // 3. Deploy Proxy, pointing to logic and calling initialize()
        bytes memory data = abi.encodeWithSelector(
            ScholarshipEscrowContract.initialize.selector,
            deployer,
            deployer,
            address(usdc)
        );
        ERC1967Proxy proxy = new ERC1967Proxy(address(logic), data);

        // Wrap the proxy address in the interface
        ScholarshipEscrowContract escrow = ScholarshipEscrowContract(address(proxy));

        vm.stopBroadcast();

        console.log("MOCK_USDC_ADDRESS=%s", address(usdc));
        console.log("SCHOLARSHIP_ESCROW_PROXY_ADDRESS=%s", address(proxy));
        console.log("SCHOLARSHIP_ESCROW_LOGIC_ADDRESS=%s", address(logic));
    }
}
