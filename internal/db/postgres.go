package db

import (
	"context"
	"fmt"
	"time"

	"api-golang/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(env config.Envs) (*pgxpool.Pool, error) {

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		env.DBUser,
		env.DBPassword,
		env.DBHost,
		env.DBPort,
		env.DBName,
		"disable",
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}
