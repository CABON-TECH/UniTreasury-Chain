#!/bin/bash

# Create a fund
curl -s -X POST http://localhost:8081/api/v1/demo/fund -H "Content-Type: application/json" -d '{"amount": 100000, "role": "admin"}'

# Create a student
curl -s -X POST http://localhost:8081/api/v1/students -H "Content-Type: application/json" -d '{"student_id": "student1", "name": "Alice Smith", "major": "Computer Science", "wallet_address": "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"}'

# Get student internal ID
ID=$(curl -s http://localhost:8081/api/v1/students | jq -r '.[0].id')

# Verify KYC
curl -s -X POST http://localhost:8081/api/v1/students/$ID/kyc

# Add credits to reach tranche index 0 (15 credits needed)
curl -s -X POST http://localhost:8081/api/v1/students/$ID/credits -H "Content-Type: application/json" -d '{"credits_to_add": 15, "new_gpa": 3.8}'

# Let the worker index and process the attestations
sleep 5

# Trigger cross chain claim
echo "Claiming L2..."
curl -s -X POST http://localhost:8081/api/v1/students/student1/claim-l2 -H "Content-Type: application/json" -d '{"fund_id": 1, "tranche_index": 0, "recipient": "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC", "dst_chain_id": 42161}'
