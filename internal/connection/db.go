package connection

import (
	"context"
	"fmt"

	"employee_attendance/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewConnection(ctx context.Context, cfg *config.Config) (*sqlx.DB, error) {
	DSN := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.Sslmode)
	db, err := sqlx.Open("postgres", DSN)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}
	if err = db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("db.PingContext: %w", err)
	}
	return db, nil
}
