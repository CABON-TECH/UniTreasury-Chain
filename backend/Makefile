# Convenience Makefile — delegates to the root Makefile.
# Run from backend/ or from the project root; both work.
ROOT := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))/..

.PHONY: run worker test tidy migrate-up migrate-down

run:
	go run ./cmd/api

worker:
	go run ./cmd/worker

test:
	go test ./... -v

tidy:
	go mod tidy

migrate-up:
	migrate -path migrations -database "$$(grep DATABASE_URL .env | cut -d= -f2-)" up

migrate-down:
	migrate -path migrations -database "$$(grep DATABASE_URL .env | cut -d= -f2-)" down
