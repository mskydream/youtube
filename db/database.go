package db

import (
	"context"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mskydream/youtube/config"
)

func InitDatabase(cfg config.Config) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), cfg.DB.DatabaseURL)
	if err != nil {
		return nil, err
	}

	defer pool.Close()

	m, err := migrate.New("file://db/migration", cfg.DB.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("error, not connected to schema, %w", err)
	}

	if err = m.Up(); err != nil {
		if err.Error() != "no change" {
			return nil, fmt.Errorf("error, cannot the up schema, %w", err)
		}
	}

	return pool, nil
}
