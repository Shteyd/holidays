package usecase

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/Shteyd/holidays/src/service/internal/lib/logger/sl"
	"github.com/Shteyd/holidays/src/service/internal/storage"
)

func (usecase CustomerUsecase) DeleteCustomer(ctx context.Context, customerID int64) error {
	const op = "usecase.CustomerUsecase.DeleteCustomer"

	log := usecase.logger.With(
		slog.String("op", op),
		slog.Int64("customer_id", customerID),
	)

	log.Debug("deleting customer")

	err := usecase.customerStorage.DeleteCustomer(ctx, customerID)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			log.Warn("customer not found")
			return ErrCustomerNotFound
		}

		log.Error("failed to delete customer", sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
