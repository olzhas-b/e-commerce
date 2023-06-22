package postgresdb

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

var txKey struct{}

type Postgres struct {
	Pool *pgxpool.Pool
}

type config struct {
	connAttempts int
	connTimeout  time.Duration
	pool         pgxpool.Pool
}

const (
	defaultConnAttempts = 3
	defaultConnTimeout  = 50 * time.Millisecond
)

func New(ctx context.Context, url string, opts ...Option) (*Postgres, error) {
	cfg := &config{
		connAttempts: defaultConnAttempts,
		connTimeout:  defaultConnTimeout,
	}
	pg := &Postgres{}

	for _, opt := range opts {
		opt(cfg)
	}

	for cfg.connAttempts > 0 {
		cfg.connAttempts--
		pool, err := pgxpool.New(ctx, url)
		if err != nil {
			time.Sleep(cfg.connTimeout)
			continue
		}
		pg.Pool = pool
		break
	}

	err := pg.Pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("[postgres] ping: %w", err)
	}
	return pg, nil
}
