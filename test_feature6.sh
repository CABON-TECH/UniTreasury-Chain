#!/bin/bash
export PATH="$HOME/.foundry/bin:$PATH"
cd backend
go run test_gas_bump.go
