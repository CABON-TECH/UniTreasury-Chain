# UniTreasury-Chain

UniTreasury-Chain is a decentralized, enterprise-grade Web3 university treasury and scholarship management system. It aims to revolutionize how universities manage their endowments, issue scholarships, and handle financial governance transparently and securely on the blockchain.

This project consists of:
- **Smart Contracts (Solidity/Foundry):** Handles robust on-chain logic for endowments, decentralized governance, treasury management, zero-knowledge proofs, and multi-signature timelocks.
- **Backend API (Go/Gin):** Serves as a transparent portal for users, manages off-chain data (users, roles), and interacts with the smart contracts using `abigen` bindings.
- **Database (PostgreSQL):** Stores relational data for fast querying, indexing events from the blockchain.

## Features Implemented

The system incorporates the following advanced Web3 features:
1. **Cross-Chain Payouts (LayerZero):** Allows students to claim their USDC scholarships across different EVM-compatible networks seamlessly.
2. **Gasless Transactions (Account Abstraction/Paymaster):** Integrates EIP-4337 to sponsor gas fees, enabling students to interact with the platform without holding native ETH.
3. **Yield-Bearing Endowments (Aave V3):** Automatically stakes idle Treasury funds into Aave to earn interest, with just-in-time unstaking during withdrawals.
4. **Zero-Knowledge (ZK) Identity:** Allows students to cryptographically prove their enrollment status on-chain while keeping their real-world identities private using Merkle Trees and nullifiers.
5. **Multi-Signature Timelock:** Enforces a mandatory 48-hour delay on all major Admin/Finance treasury withdrawals to prevent malicious actions and provide time for audits.

## Prerequisites

- [Go](https://golang.org/doc/install) (1.21+)
- [Foundry](https://getfoundry.sh/) (Forge, Anvil, Cast)
- [PostgreSQL](https://www.postgresql.org/download/) (or Docker)
- [Python 3](https://www.python.org/downloads/) (for E2E testing)

## Setup & Installation

### 1. Database Setup
Spin up the local PostgreSQL database using Docker (or point to an existing instance):
```bash
docker compose -f backend/docker-compose.yml up -d
```
Then run the database migrations:
```bash
make migrate-up
```

### 2. Smart Contracts
Install dependencies and build the contracts:
```bash
make contracts-install
make contracts-build
```

### 3. Backend Setup
Install Go dependencies and build the backend API:
```bash
make backend-tidy
make backend-build
```

## Running the Project

You can start the local environment and API manually, or use the provided End-to-End (E2E) script which handles starting the Anvil local chain, deploying the contracts, generating `.env` configurations, and starting the Go API.

```bash
# Starts anvil, deploys contracts, updates env, and starts the API
./run_test_flow.sh
```

## Testing

The project includes Python-based E2E scripts to verify core functionalities (Gasless claims, ZK Identity, Timelocks, etc.).

When running `./run_test_flow.sh`, the following tests in the `e2e/` directory are executed automatically:
- `test_claim_final2.py`: Tests the scholarship claim process.
- `test_gasless_claim.py`: Verifies EIP-4337 gasless transactions.
- `test_treasury_yield.py`: Asserts that Aave yield generation and withdrawal unstaking works.
- `test_zk_identity.py`: Evaluates ZK-proof generation and verification for student enrollment.
- `test_timelock.py`: Validates the 48-hour multi-sig timelock delay on treasury withdrawals.

## Architecture

- `contracts/`: Foundry project containing all Solidity smart contracts (`TreasuryContract.sol`, `ScholarshipEscrowContract.sol`, ZK circuits, Paymasters).
- `backend/`: Go backend containing the REST API (`cmd/api`), database repositories, and blockchain transactors (`internal/blockchain`).
- `e2e/`: Python end-to-end integration tests.

## License
MIT License
