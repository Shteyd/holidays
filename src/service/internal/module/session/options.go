package session

import (
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type options struct {
	logger      *slog.Logger
	tokenLength uint16
	database    *pgxpool.Pool
	timeout     time.Duration
}

type Option func(*options)

func WithLogger(logger *slog.Logger) Option {
	return func(opts *options) {
		opts.logger = logger
	}
}

func WithTokenLength(tokenLength uint16) Option {
	return func(opts *options) {
		opts.tokenLength = tokenLength
	}
}

func WithDatabase(database *pgxpool.Pool) Option {
	return func(opts *options) {
		opts.database = database
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(opts *options) {
		opts.timeout = timeout
	}
}
