package usecase

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/Shteyd/holidays/src/service/internal/lib/logger/sl"
	"github.com/Shteyd/holidays/src/service/internal/storage"
)

func (usecase CustomerUsecase) UpdateCustomer(
	ctx context.Context,
	customerID int64,
	name string,
	email string,
	passwordHash string,
) error {
	const op = "usecase.CustomerUsecase.UpdateCustomer"

	log := usecase.logger.With(
		slog.String("op", op),
		slog.Int64("customer_id", customerID),
	)

	log.Debug("updating customer")

	err := usecase.customerStorage.UpdateCustomer(ctx, customerID, name, email, passwordHash)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			log.Warn("customer not found")
			return ErrCustomerNotFound
		}

		log.Error("failed to update customer", sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
