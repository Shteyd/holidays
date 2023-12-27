package usecase

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/Shteyd/holidays/src/service/internal/domain/entity"
	"github.com/Shteyd/holidays/src/service/internal/lib/logger/sl"
	"github.com/Shteyd/holidays/src/service/internal/module/customer/storage"
)

func (usecase CustomerUsecase) GetCustomerByCreds(ctx context.Context, email string, password string) (entity.Customer, error) {
	const op = "usecase.CustomerUsecase.GetCustomerByCreds"

	log := usecase.logger.With(
		slog.String("op", op),
		slog.String("email", email),
	)

	log.Debug("getting customer")

	passwordHash, err := usecase.passwordManager.Hash(password)
	if err != nil {
		log.Error("failed to hash password", sl.Err(err))
		return entity.Customer{}, fmt.Errorf("%s: %w", op, err)
	}

	customer, err := usecase.customerStorage.CustomerByCreds(ctx, email, passwordHash)
	if err != nil {
		if errors.Is(err, storage.ErrCustomerNotFound) {
			log.Warn("customer not found")
			return entity.Customer{}, ErrCustomerNotFound
		}

		log.Error("failed to get customer", sl.Err(err))
		return entity.Customer{}, fmt.Errorf("%s: %w", op, err)
	}

	return customer, nil
}
