package db

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/mskydream/youtube/config"
)

func InitDatabase(cfg config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", cfg.DB.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	db.SetMaxOpenConns(cfg.DB.MaxConnections)

	m, err := migrate.New("file://db/migration", cfg.DB.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("error, not connected to schema, %w", err)
	}

	if err = m.Up(); err != nil {
		if err.Error() != "no change" {
			return nil, fmt.Errorf("error, cannot the up schema, %w", err)
		}
	}

	if err = db.Ping(); err != nil {
		defer db.Close()
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	log.Println("Database success connected...")
	return db, nil
}
