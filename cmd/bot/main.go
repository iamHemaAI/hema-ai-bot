package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/iamHemaAI/hema-ai-bot/internal/app"
	"github.com/iamHemaAI/hema-ai-bot/internal/app/closer"
	"github.com/iamHemaAI/hema-ai-bot/internal/config"
	applogger "github.com/iamHemaAI/hema-ai-bot/pkg/logger"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		logger.Fatal("failed to initialize logger", zap.Error(err))
	}

	godotenv.Load(".env")

	cfg := config.MustLoad()

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	// configure fast zap logger
	logger = applogger.SetupLogger(env, cfg.Logger)
	defer func() {
		_ = logger.Sync()
	}()

	application := app.New(logger, cfg)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger.Info("bot server starting")

	// Run blocks until ctx is cancelled; graceful shutdown is handled internally by StartConfig.
	serverDone := make(chan struct{})
	go func() {
		defer close(serverDone)
		if err := application.BOTserver.Run(ctx); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			logger.Error("bot server run failed", zap.Error(err))
		}
	}()

	logger.Info("bot server started successfully")

	<-ctx.Done()
	logger.Info("shutting down gracefully")
	stop()

	// Wait for echo's internal graceful shutdown to finish (timeout is set via GracefulTimeout in StartConfig).
	<-serverDone
	logger.Info("bot server shutdown complete")

	closerCtx, closerCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer closerCancel()

	if err := closer.CloseAll(closerCtx, logger); err != nil {
		logger.Error("shutdown failed", zap.Error(err))
	}

	logger.Info("shutdown complete")
}
