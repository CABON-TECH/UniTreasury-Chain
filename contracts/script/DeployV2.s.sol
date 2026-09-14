// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;
import "forge-std/Script.sol";
import {ScholarshipEscrowContractV2} from "../src/ScholarshipEscrowContractV2.sol";
contract DeployV2 is Script {
    function run() external {
        uint256 pk = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(pk);
        ScholarshipEscrowContractV2 logicV2 = new ScholarshipEscrowContractV2();
        vm.stopBroadcast();
        console.log("LOGIC_V2=%s", address(logicV2));
    }
}
