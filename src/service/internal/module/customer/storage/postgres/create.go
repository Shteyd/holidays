package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Shteyd/holidays/src/service/internal/module/customer/storage"
	"github.com/Shteyd/holidays/src/service/internal/module/customer/storage/postgres/internal/model"
	"github.com/jackc/pgx/v5"
)

func (postgres PostgresStorage) CreateCustomer(
	ctx context.Context,
	name string,
	email string,
	passwordHash string,
) (int64, error) {
	const op = "storage.PostgresStorage.CreateCustomer"

	customer := model.Customer{
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
	}

	query, args, err := postgres.queries.CreateCustomer(customer)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	var customerID int64
	err = postgres.conn.QueryRow(ctx, query, args...).Scan(&customerID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, storage.ErrCustomerNotFound
		}

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return customerID, nil
}
