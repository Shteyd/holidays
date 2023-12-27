package postgres

import (
	"context"
	"fmt"

	"github.com/Shteyd/holidays/src/service/internal/module/customer/storage"
	"github.com/Shteyd/holidays/src/service/internal/module/customer/storage/postgres/internal/model"
)

func (postgres PostgresStorage) UpdateCustomer(
	ctx context.Context,
	customerID int64,
	name string,
	email string,
	passwordHash string,
) error {
	const op = "storage.PostgresStorage.UpdateCustomer"

	customer := model.Customer{
		ID:           customerID,
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
	}

	query, args, err := postgres.queries.UpdateCustomer(customer)
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
