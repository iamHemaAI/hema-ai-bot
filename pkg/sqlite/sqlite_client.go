package sqliteclient

import (
	"database/sql"
	"fmt"

	"github.com/iamHemaAI/hema-ai-bot/internal/config"
	"go.uber.org/zap"
	_ "modernc.org/sqlite"
)

type Sqlite struct {
	db *sql.DB
}

func New(logger *zap.Logger, cfg config.SqliteConfig) (*Sqlite, error) {
	db, err := sql.Open("sqlite", cfg.Path)
	if err != nil {
		logger.Error("open sqlite", zap.Error(err))
		return nil, fmt.Errorf("open sqlite: %w", err)
	}

	if err := db.Ping(); err != nil {
		logger.Error("ping sqlite", zap.Error(err))
		return nil, fmt.Errorf("ping sqlite: %w", err)
	}

	logger.Info("sqlite connected", zap.String("path", cfg.Path))

	return &Sqlite{db: db}, nil
}

func (s *Sqlite) DB() *sql.DB {
	return s.db
}

func (s *Sqlite) Close() error {
	return s.db.Close()
}
