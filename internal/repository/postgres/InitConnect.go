package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"os"
)

func CreateConnection(ctx context.Context) (*pgxpool.Pool, error) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(".env file not found")
	}
	key := os.Getenv("DATABASE_URL")
	pool, err := pgxpool.New(ctx, key)
	if err != nil {
		return nil, err
	}
	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return pool, err
}
