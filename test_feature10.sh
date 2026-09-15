#!/bin/bash
export PATH="$HOME/.foundry/bin:$PATH"

echo "==========================================================="
echo " TESTING FEATURE 10: MULTI-SIG KEY ROTATION & RECOVERY "
echo "==========================================================="
echo ""
echo "Running Foundry Simulation: TreasurySignerRecoveryTest..."
echo "Scenario: Approver 3's key is compromised. Approver 1 & Approver 2 vote to replace it."
echo ""

cd contracts && forge test --match-test testSignerRecovery -vvv
