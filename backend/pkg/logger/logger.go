package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New creates a zap logger appropriate for the given environment.
// In production, uses JSON format with Info level.
// In development, uses console format with Debug level.
func New(env string) (*zap.Logger, error) {
	var cfg zap.Config
	if env == "production" {
		cfg = zap.NewProductionConfig()
		cfg.EncoderConfig.TimeKey = "ts"
	} else {
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	return cfg.Build()
}

// Must creates a logger and panics if construction fails.
func Must(env string) *zap.Logger {
	l, err := New(env)
	if err != nil {
		panic("logger: failed to create logger: " + err.Error())
	}
	return l
}
