package config
import (
	"fmt"
	"os"
	"strconv"
	"time"
	"github.com/joho/godotenv"
)
type Config struct {
	Port string
	Env  string
	DatabaseURL string
	RPCURL  string
	ChainID int64
	TreasuryAddress          string
	FeeRegistryAddress       string
	EscrowAddress            string
	AttestorPrivateKey string
	JWTSecret      string
	JWTExpiryHours int
	IndexerStartBlock    uint64
	IndexerPollInterval  time.Duration
	IndexerBatchSize     uint64
}
func Load() (*Config, error) {
	_ = godotenv.Load()
	c := &Config{
		Port:               getEnv("PORT", "8080"),
		Env:                getEnv("ENV", "development"),
		DatabaseURL:        mustGetEnv("DATABASE_URL"),
		RPCURL:             mustGetEnv("SEPOLIA_RPC_URL"),
		TreasuryAddress:    mustGetEnv("TREASURY_CONTRACT_ADDRESS"),
		FeeRegistryAddress: mustGetEnv("FEE_REGISTRY_CONTRACT_ADDRESS"),
		EscrowAddress:      getEnv("SCHOLARSHIP_ESCROW_CONTRACT_ADDRESS", ""),
		AttestorPrivateKey: mustGetEnv("ATTESTOR_PRIVATE_KEY"),
		JWTSecret:          mustGetEnv("JWT_SECRET"),
	}
	chainID, err := parseInt64("CHAIN_ID", 11155111)
	if err != nil {
		return nil, fmt.Errorf("config: CHAIN_ID: %w", err)
	}
	c.ChainID = chainID
	jwtHours, err := parseInt("JWT_EXPIRY_HOURS", 24)
	if err != nil {
		return nil, fmt.Errorf("config: JWT_EXPIRY_HOURS: %w", err)
	}
	c.JWTExpiryHours = jwtHours
	startBlock, err := parseUint64("INDEXER_START_BLOCK", 0)
	if err != nil {
		return nil, fmt.Errorf("config: INDEXER_START_BLOCK: %w", err)
	}
	c.IndexerStartBlock = startBlock
	pollSecs, err := parseInt("INDEXER_POLL_INTERVAL_SECONDS", 15)
	if err != nil {
		return nil, fmt.Errorf("config: INDEXER_POLL_INTERVAL_SECONDS: %w", err)
	}
	c.IndexerPollInterval = time.Duration(pollSecs) * time.Second
	batchSize, err := parseUint64("INDEXER_BATCH_SIZE", 100)
	if err != nil {
		return nil, fmt.Errorf("config: INDEXER_BATCH_SIZE: %w", err)
	}
	c.IndexerBatchSize = batchSize
	return c, nil
}
func (c *Config) IsProduction() bool {
	return c.Env == "production"
}
func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
func mustGetEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(fmt.Sprintf("config: required environment variable %q is not set", key))
	}
	return v
}
func parseInt(key string, fallback int) (int, error) {
	v := os.Getenv(key)
	if v == "" {
		return fallback, nil
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("invalid int for %s=%q: %w", key, v, err)
	}
	return n, nil
}
func parseInt64(key string, fallback int64) (int64, error) {
	v := os.Getenv(key)
	if v == "" {
		return fallback, nil
	}
	n, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid int64 for %s=%q: %w", key, v, err)
	}
	return n, nil
}
func parseUint64(key string, fallback uint64) (uint64, error) {
	v := os.Getenv(key)
	if v == "" {
		return fallback, nil
	}
	n, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid uint64 for %s=%q: %w", key, v, err)
	}
	return n, nil
}
