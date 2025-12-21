package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(connString string) (*pgxpool.Pool, error) {
	retries := 10

	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("failed to create pool: %w", err)
	}

	for i := 1; i <= retries; i++ {
		err := pool.Ping(context.Background())
		if err == nil {
			log.Printf("✅ connected to postgres")
			return pool, nil
		}

		log.Printf("DB not ready, retrying...", err)
		time.Sleep(time.Duration(i+1) * time.Second)
	}

	return nil, fmt.Errorf("could not connect to postgres after %d attempts: %w", retries, err)
}
