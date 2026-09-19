// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Script, console} from "forge-std/Script.sol";
import {TreasuryContract} from "../src/TreasuryContract.sol";
import {FeeRegistryContract} from "../src/FeeRegistryContract.sol";
import {ScholarshipEscrowContract} from "../src/ScholarshipEscrowContract.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {MockLendingPool} from "../src/MockLendingPool.sol";
import {MockAToken} from "../src/MockAToken.sol";
import {MockUSDC} from "../src/MockUSDC.sol";
import {MockLayerZeroEndpoint} from "../src/MockLayerZeroEndpoint.sol";
import {MockEntryPoint} from "../src/MockEntryPoint.sol";
import {UniPaymaster} from "../src/UniPaymaster.sol";

import {UniPaymaster} from "../src/UniPaymaster.sol";

contract DeployAllScript is Script {
    function run() external {
        uint256 deployerKey = 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80; // default anvil account 0
        address deployer = vm.addr(deployerKey);
        
        vm.startBroadcast(deployerKey);

        MockUSDC usdc = new MockUSDC();
        usdc.mint(deployer, 1_000_000 * 10**18);
        console.log("MOCK_USDC_ADDRESS=", address(usdc));

        MockLendingPool lendingPool = new MockLendingPool();
        MockAToken aUsdc = new MockAToken("Aave interest bearing USDC", "aUSDC", address(usdc));
        lendingPool.initReserve(address(usdc), address(aUsdc));
        // Transfer ownership of aUsdc to lending pool so it can mint/burn
        aUsdc.transferOwnership(address(lendingPool));
        address[] memory approvers = new address[](1);
        approvers[0] = deployer;

        TreasuryContract treasury = new TreasuryContract(
            deployer,
            approvers,
            1,
            1000 * 10**18,
            address(usdc),
            address(lendingPool),
            address(aUsdc)
        );
        console.log("TREASURY_CONTRACT_ADDRESS=", address(treasury));
        usdc.approve(address(treasury), 500_000 * 10**18);
        treasury.deposit(500_000 * 10**18);

        FeeRegistryContract feeRegistry = new FeeRegistryContract(deployer, deployer, address(usdc));
        console.log("FEE_REGISTRY_CONTRACT_ADDRESS=", address(feeRegistry));

        ScholarshipEscrowContract logic = new ScholarshipEscrowContract();
        bytes memory data = abi.encodeWithSelector(ScholarshipEscrowContract.initialize.selector, deployer, deployer, address(usdc));
        ERC1967Proxy proxy = new ERC1967Proxy(address(logic), data);
        ScholarshipEscrowContract escrow = ScholarshipEscrowContract(address(proxy));
        console.log("SCHOLARSHIP_ESCROW_CONTRACT_ADDRESS=", address(escrow));

        MockLayerZeroEndpoint lzEndpoint = new MockLayerZeroEndpoint();
        escrow.setLzEndpoint(address(lzEndpoint));
        console.log("MOCK_LZ_ENDPOINT_ADDRESS=", address(lzEndpoint));

        MockEntryPoint entryPoint = new MockEntryPoint();
        console.log("ENTRYPOINT_ADDRESS=", address(entryPoint));

        UniPaymaster paymaster = new UniPaymaster(address(entryPoint));
        paymaster.setAllowedTarget(address(escrow), true);
        console.log("PAYMASTER_ADDRESS=", address(paymaster));











        
        vm.stopBroadcast();
        
        

    }
}
