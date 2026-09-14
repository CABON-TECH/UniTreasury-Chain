#!/bin/bash
set -e

echo "🧹 Cleaning up old Anvil processes..."
pkill anvil || true
sleep 1

echo "🚀 Starting local Anvil node..."
~/.foundry/bin/anvil > anvil.log 2>&1 &
ANVIL_PID=$!
sleep 2

echo "🏗️  Deploying Merkle-enabled Contracts..."
export PATH="$HOME/.foundry/bin:$PATH"
cd contracts
forge script script/DeployAll.s.sol --rpc-url http://localhost:8545 --broadcast > deploy.log

echo "🔍 Extracting Addresses..."
USDC_ADDR=$(grep "MOCK_USDC_ADDRESS=" deploy.log | cut -d '=' -f 2 | tr -d '[:space:]' | sed 's/\x1b\[[0-9;]*m//g')
TREASURY_ADDR=$(grep "TREASURY_CONTRACT_ADDRESS=" deploy.log | cut -d '=' -f 2 | tr -d '[:space:]' | sed 's/\x1b\[[0-9;]*m//g')
FEE_ADDR=$(grep "FEE_REGISTRY_CONTRACT_ADDRESS=" deploy.log | cut -d '=' -f 2 | tr -d '[:space:]' | sed 's/\x1b\[[0-9;]*m//g')
ESCROW_ADDR=$(grep "SCHOLARSHIP_ESCROW_CONTRACT_ADDRESS=" deploy.log | cut -d '=' -f 2 | tr -d '[:space:]' | sed 's/\x1b\[[0-9;]*m//g')

echo "⚙️  Updating Backend .env..."
cd ../backend
sed -i "s/^TREASURY_CONTRACT_ADDRESS=.*/TREASURY_CONTRACT_ADDRESS=$TREASURY_ADDR/" .env
sed -i "s/^FEE_REGISTRY_CONTRACT_ADDRESS=.*/FEE_REGISTRY_CONTRACT_ADDRESS=$FEE_ADDR/" .env
sed -i "s/^SCHOLARSHIP_ESCROW_CONTRACT_ADDRESS=.*/SCHOLARSHIP_ESCROW_CONTRACT_ADDRESS=$ESCROW_ADDR/" .env

echo "🌱 Seeding Blockchain state to match Database..."
PK=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
cast send $USDC_ADDR "approve(address,uint256)" $ESCROW_ADDR 1000000000000000000 --private-key $PK --rpc-url http://localhost:8545 > /dev/null
cast send $ESCROW_ADDR "createFund(address,uint256,uint256,uint256)" 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 1000000000000000000 2 500000000000000000 --private-key $PK --rpc-url http://localhost:8545 > /dev/null

echo "✅ Ready!"
echo "Now run the worker to generate the Merkle Tree and publish the root:"
echo "cd backend && make worker"

echo "🧹 Clearing old database claims so the worker processes them again..."
psql "postgres://unitreasury:password@localhost:5433/unitreasury?sslmode=disable" -c "DELETE FROM tranche_releases;" > /dev/null
