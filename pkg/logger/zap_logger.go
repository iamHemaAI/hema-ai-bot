// using zap for structured logging
package logger

import (
	"github.com/iamHemaAI/hema-ai-bot/internal/config"
	"go.uber.org/zap"
)

// SetupLogger initializes and returns a zap logger based on the environment and configuration
func SetupLogger(env string, configs map[string]config.LoggerConfig) *zap.Logger {
	cfg, ok := configs[env]
	if !ok {
		cfg = configs["dev"]
	}

	level, err := zap.ParseAtomicLevel(cfg.Level)
	if err != nil {
		panic("invalid logger level: " + err.Error())
	}

	encoderCfg := zap.NewDevelopmentEncoderConfig()
	if !cfg.Development {
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	logger, err := zap.Config{
		Level:            level,
		Development:      cfg.Development,
		Encoding:         cfg.Encoding,
		EncoderConfig:    encoderCfg,
		OutputPaths:      cfg.OutputPaths,
		ErrorOutputPaths: cfg.ErrorOutputPaths,
	}.Build(zap.AddStacktrace(zap.FatalLevel))
	if err != nil {
		panic("failed to init logger: " + err.Error())
	}

	return logger
}
