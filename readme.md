# UniTreasury Chain

A blockchain-based university treasury management system built for final-year research.

## Monorepo Layout

```
UniTreasury-Chain/
├── contracts/        # Foundry smart contracts
├── backend/          # Go API + worker
├── frontend/         # Next.js dashboard (Sprint 4+)
└── Makefile          # Root-level orchestration
```

## Sprints

| Sprint | Focus | Status |
|--------|-------|--------|
| 1 | Contracts foundation (Treasury + FeeRegistry) | In Progress |
| 2 | Go skeleton + payment ingestion flow | Pending |
| 3 | Treasury multi-sig flow + scholarship escrow | Pending |
| 4+ | Dashboards, reconciliation, evaluation harness | Pending |

## Quick Start

```bash
# Install Foundry
curl -L https://foundry.paradigm.xyz | bash && foundryup

# Contracts
cd contracts && forge install && forge test

# Backend
cd backend && go mod tidy && go run ./cmd/api
```

## Architecture

- **Contracts**: Foundry, Solidity 0.8.20+, OpenZeppelin
- **Backend**: Go 1.22+, PostgreSQL, Alchemy (Sepolia)
- **Frontend**: Next.js 14, role-based access (Admin / Finance / Student)
- **Oracle**: Single trusted attestor (EIP-712), explicit centralization trade-off
