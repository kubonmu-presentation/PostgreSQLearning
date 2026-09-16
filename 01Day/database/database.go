package database

import (
	"context"
	"errors"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv" 
)

func Connect() (*pgxpool.Pool, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, errors.New("DATABASE_URL이 없습니다")
	}

	db, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(context.Background()); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
