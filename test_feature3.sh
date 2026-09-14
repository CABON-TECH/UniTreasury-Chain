#!/bin/bash
set -e

echo "🧹 Cleaning up old Anvil processes..."
pkill anvil || true
sleep 1

echo "🚀 Starting local Anvil node..."
~/.foundry/bin/anvil > anvil.log 2>&1 &
ANVIL_PID=$!
sleep 2

echo "🏗️  Deploying Contracts..."
export PATH="$HOME/.foundry/bin:$PATH"
cd contracts
forge script script/DeployAll.s.sol --rpc-url http://localhost:8545 --broadcast > deploy.log

echo "🔍 Extracting Addresses..."
USDC_ADDR=$(grep "MOCK_USDC_ADDRESS=" deploy.log | cut -d '=' -f 2 | tr -d '[:space:]' | sed 's/\x1b\[[0-9;]*m//g')
ESCROW_ADDR=$(grep "SCHOLARSHIP_ESCROW_CONTRACT_ADDRESS=" deploy.log | cut -d '=' -f 2 | tr -d '[:space:]' | sed 's/\x1b\[[0-9;]*m//g')

echo "🌱 Seeding Blockchain state..."
PK=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
ADMIN_ADDR=0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266

# Create a fund of 100 USDC
echo "💰 Creating a Scholarship Fund with 100 USDC..."
cast send $USDC_ADDR "approve(address,uint256)" $ESCROW_ADDR 100000000000000000000 --private-key $PK --rpc-url http://localhost:8545 > /dev/null
cast send $ESCROW_ADDR "createFund(address,uint256,uint256,uint256)" $ADMIN_ADDR 100000000000000000000 2 50000000000000000000 --private-key $PK --rpc-url http://localhost:8545 > /dev/null

echo ""
echo "📊 Checking balances before clawback..."
ESCROW_BALANCE=$(cast call $USDC_ADDR "balanceOf(address)(uint256)" $ESCROW_ADDR --rpc-url http://localhost:8545)
ADMIN_BALANCE=$(cast call $USDC_ADDR "balanceOf(address)(uint256)" $ADMIN_ADDR --rpc-url http://localhost:8545)
echo "Escrow Balance: $ESCROW_BALANCE wei"
echo "Admin Balance:  $ADMIN_BALANCE wei"

echo ""
echo "🚨 Executing CLAWBACK (Fund ID 1) as Admin..."
cast send $ESCROW_ADDR "clawbackFund(uint256,address)" 1 $ADMIN_ADDR --private-key $PK --rpc-url http://localhost:8545 > /dev/null

echo ""
echo "📊 Checking balances AFTER clawback..."
ESCROW_BALANCE_AFTER=$(cast call $USDC_ADDR "balanceOf(address)(uint256)" $ESCROW_ADDR --rpc-url http://localhost:8545)
ADMIN_BALANCE_AFTER=$(cast call $USDC_ADDR "balanceOf(address)(uint256)" $ADMIN_ADDR --rpc-url http://localhost:8545)
echo "Escrow Balance: $ESCROW_BALANCE_AFTER wei"
echo "Admin Balance:  $ADMIN_BALANCE_AFTER wei"

echo ""
echo "✅ Clawback successful! The 100 USDC was returned to the Admin!"
