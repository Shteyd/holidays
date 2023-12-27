package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Shteyd/holidays/src/service/internal/domain/entity"
	"github.com/Shteyd/holidays/src/service/internal/module/session/storage"
	"github.com/Shteyd/holidays/src/service/internal/module/session/storage/postgres/internal/model"
	"github.com/jackc/pgx/v5"
)

func (postgres PostgresStorage) SessionByCustomerID(ctx context.Context, customerID int64) ([]entity.Session, error) {
	const op = "storage.postgres.GetByCustomerID"

	query, args, err := postgres.queries.GetByCustomerID(customerID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := postgres.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	models, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[model.Session])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return model.ListToEntity(models), nil
}

func (postgres PostgresStorage) SessionByToken(ctx context.Context, token string) (entity.Session, error) {
	const op = "storage.postgres.GetByToken"

	query, args, err := postgres.queries.GetByToken(token)
	if err != nil {
		return entity.Session{}, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := postgres.conn.Query(ctx, query, args...)
	if err != nil {
		return entity.Session{}, fmt.Errorf("%s: %w", op, err)
	}

	model, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[model.Session])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Session{}, storage.ErrSessionNotFound
		}

		return entity.Session{}, fmt.Errorf("%s: %w", op, err)
	}

	return model.ToEntity(), nil
}
