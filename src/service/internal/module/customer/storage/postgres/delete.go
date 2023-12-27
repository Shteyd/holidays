package postgres

import (
	"context"
	"fmt"

	"github.com/Shteyd/holidays/src/service/internal/module/customer/storage"
)

func (postgres PostgresStorage) DeleteCustomer(
	ctx context.Context,
	customerID int64,
) error {
	const op = "storage.PostgresStorage.DeleteCustomer"

	query, args, err := postgres.queries.DeleteCustomer(customerID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	info, err := postgres.conn.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if info.RowsAffected() == 0 {
		return storage.ErrCustomerNotFound
	}

	return nil
}
