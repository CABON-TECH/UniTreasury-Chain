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
