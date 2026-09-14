#!/bin/bash
set -e

echo "🧹 Cleaning up old Anvil processes..."
pkill anvil || true
sleep 1

echo "🚀 Starting local Anvil node..."
~/.foundry/bin/anvil > anvil.log 2>&1 &
ANVIL_PID=$!
sleep 2

echo "🏗️  Deploying Proxy & Implementation V1..."
export PATH="$HOME/.foundry/bin:$PATH"
cd contracts
forge script script/DeployAll.s.sol --rpc-url http://localhost:8545 --broadcast > deploy.log

ESCROW_PROXY=$(grep "SCHOLARSHIP_ESCROW_CONTRACT_ADDRESS=" deploy.log | cut -d '=' -f 2 | tr -d '[:space:]' | sed 's/\x1b\[[0-9;]*m//g')
echo "✅ V1 Proxy deployed at: $ESCROW_PROXY"

echo ""
echo "🚀 Deploying V2 Implementation Contract..."
cat << 'DEPLOY_V2' > script/DeployV2.s.sol
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
DEPLOY_V2

PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 forge script script/DeployV2.s.sol --rpc-url http://localhost:8545 --broadcast > deploy_v2.log
LOGIC_V2=$(grep "LOGIC_V2=" deploy_v2.log | cut -d '=' -f 2 | tr -d '[:space:]' | sed 's/\x1b\[[0-9;]*m//g')
echo "✅ V2 Logic deployed at: $LOGIC_V2"

echo ""
echo "🔄 Upgrading Proxy from V1 -> V2..."
PK=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
cast send $ESCROW_PROXY "upgradeToAndCall(address,bytes)" $LOGIC_V2 0x --private-key $PK --rpc-url http://localhost:8545 > /dev/null

echo "✨ Setting the new 'version' string on V2..."
cast send $ESCROW_PROXY "setVersion(string)" "v2.0.0 - The Upgraded Escrow!" --private-key $PK --rpc-url http://localhost:8545 > /dev/null

VERSION=$(cast call $ESCROW_PROXY "version()(string)" --rpc-url http://localhost:8545)
echo ""
echo "🎉 SUCCESS! The proxy is now pointing to V2. Read version from proxy: $VERSION"
echo ""
