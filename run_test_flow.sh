#!/bin/bash
ng existing services..."
killall anvil || true
killall api || true
echo "Starting anvil..."
export PATH=$PATH:/home/cabon-tech/.foundry/bin
nohup anvil > anvil.log 2>&1 &
sleep 2
echo "Deploying contracts..."
cd contracts
forge script script/DeployAll.s.sol:DeployAllScript --rpc-url http://127.0.0.1:8545 --broadcast --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 > forge_out.txt
cat forge_out.txt
cd ..
echo "Updating environment..."
python3 update_env.py
source backend/.env
cast send $MOCK_USDC_ADDRESS "approve(address,uint256)" $SCHOLARSHIP_ESCROW_CONTRACT_ADDRESS 50000 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --rpc-url http://127.0.0.1:8545 > /dev/null
cast send $SCHOLARSHIP_ESCROW_CONTRACT_ADDRESS "createFund(address,uint256,uint256,uint256)" 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 50000 5 10000 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --rpc-url http://127.0.0.1:8545 > /dev/null
PGPASSWORD=notmysim_secret psql -h localhost -p 5433 -U notmysim -d unitreasury -c "UPDATE scholarship_funds SET on_chain_id=0 WHERE id=1; TRUNCATE withdrawal_proposals CASCADE;" > /dev/null
echo "Starting API..."
cd backend
go build ./cmd/api
nohup ./api > api.log 2>&1 &
cd ..
sleep 2
echo "Running Python test scripts..."
python3 e2e/test_claim_final2.py
python3 e2e/test_gasless_claim.py
python3 e2e/test_treasury_yield.py
python3 e2e/test_zk_identity.py
python3 e2e/test_timelock.py
echo "Done!"
