package postgresdb

import (
	"time"
)

type Option func(cfg *config)

func WithAttempt(connAttempts int) Option {
	return func(cfg *config) {
		cfg.connAttempts = connAttempts
	}
}

func WithAttemptTimeout(connTimeout time.Duration) Option {
	return func(cfg *config) {
		cfg.connTimeout = connTimeout
	}
}
