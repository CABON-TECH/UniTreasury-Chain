package logger
import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)
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
func Must(env string) *zap.Logger {
	l, err := New(env)
	if err != nil {
		panic("logger: failed to create logger: " + err.Error())
	}
	return l
}
