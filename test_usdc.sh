#!/bin/bash
set -e

echo "Starting local Anvil node..."
~/.foundry/bin/anvil > anvil.log 2>&1 &
ANVIL_PID=$!
sleep 2

echo "Deploying Contracts & Mock USDC..."
export PATH="$HOME/.foundry/bin:$PATH"
cd contracts
forge script script/DeployAll.s.sol --rpc-url http://localhost:8545 --broadcast > deploy.log

echo "Extracting Addresses..."
USDC_ADDR=$(grep "MOCK_USDC_ADDRESS=" deploy.log | cut -d '=' -f 2 | tr -d '[:space:]' | sed 's/\x1b\[[0-9;]*m//g')
TREASURY_ADDR=$(grep "TREASURY_CONTRACT_ADDRESS=" deploy.log | cut -d '=' -f 2 | tr -d '[:space:]' | sed 's/\x1b\[[0-9;]*m//g')
FEE_ADDR=$(grep "FEE_REGISTRY_CONTRACT_ADDRESS=" deploy.log | cut -d '=' -f 2 | tr -d '[:space:]' | sed 's/\x1b\[[0-9;]*m//g')
ESCROW_ADDR=$(grep "SCHOLARSHIP_ESCROW_CONTRACT_ADDRESS=" deploy.log | cut -d '=' -f 2 | tr -d '[:space:]' | sed 's/\x1b\[[0-9;]*m//g')

echo "Updating Backend .env..."
cd ../backend
sed -i "s/^TREASURY_CONTRACT_ADDRESS=.*/TREASURY_CONTRACT_ADDRESS=$TREASURY_ADDR/" .env
sed -i "s/^FEE_REGISTRY_CONTRACT_ADDRESS=.*/FEE_REGISTRY_CONTRACT_ADDRESS=$FEE_ADDR/" .env
sed -i "s/^SCHOLARSHIP_ESCROW_CONTRACT_ADDRESS=.*/SCHOLARSHIP_ESCROW_CONTRACT_ADDRESS=$ESCROW_ADDR/" .env
sed -i "s|^SEPOLIA_RPC_URL=.*|SEPOLIA_RPC_URL=http://localhost:8545|" .env
sed -i "s/^CHAIN_ID=.*/CHAIN_ID=31337/" .env

echo ""
echo "=========================================="
echo "✅ LOCAL BLOCKCHAIN READY!"
echo "USDC:    $USDC_ADDR"
echo "Escrow:  $ESCROW_ADDR"
echo "Treasury:$TREASURY_ADDR"
echo "=========================================="
echo "To test the system:"
echo "1. In this terminal, run:  cd backend && make run"
echo "2. In a NEW terminal, run: cd backend && make worker"
echo "3. The orchestrator will now successfully broadcast the USDC transaction to your local blockchain!"
echo "=========================================="
echo "(To stop Anvil later, run: kill $ANVIL_PID)"
