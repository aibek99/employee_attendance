package app

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"employee_attendance/internal/config"
	"employee_attendance/internal/connection"

	"github.com/jmoiron/sqlx"
)

type App struct {
	config *config.Config
	db     *sqlx.DB
	logger *slog.Logger
}

func NewApp(ctx context.Context) (*App, error) {
	cfg, err := config.NewConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("config.NewConfig: %w", err)
	}
	db, err := connection.NewConnection(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("connection.NewConnection: %w", err)
	}
	handler := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(handler)

	return &App{
		config: cfg,
		db:     db,
		logger: logger,
	}, nil
}
