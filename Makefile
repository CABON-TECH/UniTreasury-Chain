.PHONY: all contracts-install contracts-test contracts-build generate-bindings backend-tidy backend-run migrate-up migrate-down

# Contracts
contracts-install:
	cd contracts && forge install

contracts-build:
	cd contracts && forge build

contracts-test:
	cd contracts && forge test -vvv

contracts-coverage:
	cd contracts && forge coverage --report summary

contracts-gas:
	cd contracts && forge test --gas-report

deploy-local:
	cd contracts && forge script script/Deploy.s.sol --rpc-url http://localhost:8545 --broadcast

deploy-sepolia:
	cd contracts && forge script script/DeploySepolia.s.sol \
		--rpc-url $(SEPOLIA_RPC_URL) \
		--private-key $(DEPLOYER_PRIVATE_KEY) \
		--broadcast \
		--verify \
		--etherscan-api-key $(ETHERSCAN_API_KEY)

# Code generation
generate-bindings: contracts-build
	@echo "Generating Go bindings from Foundry ABIs..."
	@mkdir -p backend/internal/blockchain/bindings
	abigen --abi contracts/out/TreasuryContract.sol/TreasuryContract.json \
		--pkg bindings --type TreasuryContract \
		--out backend/internal/blockchain/bindings/treasury.go
	abigen --abi contracts/out/FeeRegistryContract.sol/FeeRegistryContract.json \
		--pkg bindings --type FeeRegistryContract \
		--out backend/internal/blockchain/bindings/feeregistry.go
	abigen --abi contracts/out/ScholarshipEscrowContract.sol/ScholarshipEscrowContract.json \
		--pkg bindings --type ScholarshipEscrowContract \
		--out backend/internal/blockchain/bindings/escrow.go

# Backend
backend-tidy:
	cd backend && go mod tidy

backend-build:
	cd backend && go build ./cmd/api && go build ./cmd/worker

backend-run:
	cd backend && go run ./cmd/api

worker-run:
	cd backend && go run ./cmd/worker

backend-test:
	cd backend && go test ./... -v

# Database
migrate-up:
	cd backend && migrate -path migrations -database "$(DATABASE_URL)" up

migrate-down:
	cd backend && migrate -path migrations -database "$(DATABASE_URL)" down

# Docker
up:
	docker compose -f backend/docker-compose.yml up -d

down:
	docker compose -f backend/docker-compose.yml down
