package postgres

import (
	"log/slog"

	"github.com/Shteyd/holidays/src/service/internal/module/customer/storage/postgres/internal/query"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStorage struct {
	conn    *pgxpool.Pool
	queries query.QueryGenerator
}

func New(logger *slog.Logger, conn *pgxpool.Pool) PostgresStorage {
	queryGenerator := query.New(logger)

	return PostgresStorage{
		conn:    conn,
		queries: queryGenerator,
	}
}
