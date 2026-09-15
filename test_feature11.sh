#!/bin/bash

echo "========================================================"
echo " TESTING FEATURE 11: WEBSOCKETS FOR REAL-TIME UI "
echo "========================================================"
echo ""

echo "[1/3] Starting API and Worker in the background..."
cd backend
kill -9 $(lsof -t -i :8081) 2>/dev/null || true
pkill -f "./worker" || true
pkill -f "./api" || true

nohup ./api > api.log 2>&1 &
nohup ./worker > worker.log 2>&1 &

sleep 2

echo "[2/3] Open your browser to: http://localhost:8081/dashboard"
echo "      (If prompted to login, click 'Student Login' or sign in as any student)"
echo ""
echo "      Keep the browser window visible, then come back here!"
echo ""
read -p "Press [ENTER] when you are looking at the Student Dashboard..."

echo ""
echo "[3/3] Firing a mock Blockchain Event into the network..."

cat << 'GO_EOF' > trigger_event.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		dbUrl = "postgres://unitreasury:password@localhost:5433/unitreasury?sslmode=disable"
	}
	pool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	payloadMap := map[string]interface{}{
		"type": "BlockchainEvent",
		"contract": "Escrow",
		"tx_hash": "0xMockWebSocketEventHasArrivedSuccessfully1234",
	}
	bMsg, _ := json.Marshal(payloadMap)
	
	_, err = pool.Exec(context.Background(), "NOTIFY ws_events, '" + string(bMsg) + "'")
	if err != nil {
		panic(err)
	}
	fmt.Println("Event successfully emitted via Postgres NOTIFY!")
}
GO_EOF

go run trigger_event.go
rm trigger_event.go

echo ""
echo "Check your browser! You should see a blue Toast notification pop up instantly!"
echo ""
