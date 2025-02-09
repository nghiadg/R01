package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectPostgresDB() (*pgxpool.Pool, error) {
	// connect to database
	dbUrl := "postgres://postgres:postgres@localhost:5432/postgres"
	config, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to database")
	return pool, nil
}
