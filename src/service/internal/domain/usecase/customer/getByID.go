package customer

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/Shteyd/holidays/src/service/internal/domain/entity"
	"github.com/Shteyd/holidays/src/service/internal/lib/logger/sl"
	"github.com/Shteyd/holidays/src/service/internal/storage"
)

func (usecase CustomerUsecase) GetCustomerByID(ctx context.Context, customerID int64) (entity.Customer, error) {
	const op = "usecase.CustomerUsecase.GetCustomerByID"

	log := usecase.logger.With(
		slog.String("op", op),
		slog.Int64("customer_id", customerID),
	)

	log.Debug("getting customer")

	customer, err := usecase.customerProvider.CustomerByID(ctx, customerID)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			log.Warn("customer not found")
			return entity.Customer{}, ErrCustomerNotFound
		}

		log.Error("failed to get customer", sl.Err(err))
		return entity.Customer{}, fmt.Errorf("%s: %w", op, err)
	}

	return customer, nil
}
