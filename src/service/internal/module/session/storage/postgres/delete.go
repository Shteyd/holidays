package postgres

import (
	"context"
	"fmt"

	"github.com/Shteyd/holidays/src/service/internal/module/session/storage"
)

func (postgres PostgresStorage) DeleteByToken(ctx context.Context, token string) error {
	const op = "storage.postgres.DeleteByToken"

	query, args, err := postgres.queries.DeleteByToken(token)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	info, err := postgres.conn.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if info.RowsAffected() == 0 {
		return storage.ErrSessionNotFound
	}

	return nil
}

func (postgres PostgresStorage) DeleteByCustomerID(ctx context.Context, customerID int64) error {
	const op = "storage.postgres.DeleteByCustomerID"

	query, args, err := postgres.queries.DeleteByCustomerID(customerID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = postgres.conn.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
