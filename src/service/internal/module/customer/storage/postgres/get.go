package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Shteyd/holidays/src/service/internal/domain/entity"
	"github.com/Shteyd/holidays/src/service/internal/module/customer/storage"
	"github.com/Shteyd/holidays/src/service/internal/module/customer/storage/postgres/internal/model"
	"github.com/jackc/pgx/v5"
)

func (postgres PostgresStorage) CustomerByID(
	ctx context.Context,
	customerID int64,
) (entity.Customer, error) {
	const op = "storage.PostgresStorage.CustomerByID"

	query, args, err := postgres.queries.CustomerByID(customerID)
	if err != nil {
		return entity.Customer{}, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := postgres.conn.Query(ctx, query, args...)
	if err != nil {
		return entity.Customer{}, fmt.Errorf("%s: %w", op, err)
	}

	model, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[model.Customer])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Customer{}, storage.ErrCustomerNotFound
		}

		return entity.Customer{}, fmt.Errorf("%s: %w", op, err)
	}

	return model.ToEntity(), nil
}

func (postgres PostgresStorage) CustomerByCreds(
	ctx context.Context,
	email string,
	password string,
) (entity.Customer, error) {
	const op = "storage.PostgresStorage.CustomerByCreds"

	customer := model.Customer{
		Email:        email,
		PasswordHash: password,
	}

	query, args, err := postgres.queries.CustomerByCreds(customer)
	if err != nil {
		return entity.Customer{}, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := postgres.conn.Query(ctx, query, args...)
	if err != nil {
		return entity.Customer{}, fmt.Errorf("%s: %w", op, err)
	}

	model, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[model.Customer])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Customer{}, storage.ErrCustomerNotFound
		}

		return entity.Customer{}, fmt.Errorf("%s: %w", op, err)
	}

	return model.ToEntity(), nil
}
