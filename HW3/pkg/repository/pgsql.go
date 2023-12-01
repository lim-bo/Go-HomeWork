package repository

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PgRepo struct {
	M    *sync.Mutex
	Pool *pgxpool.Pool
}

func New(ctx context.Context, connStr string) (*PgRepo, error) {
	pool, err := pgxpool.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}
	return &PgRepo{M: &sync.Mutex{}, Pool: pool}, nil
}
