package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Shteyd/holidays/src/service/internal/lib/logger/sl"
)

func (usecase CustomerUsecase) CreateCustomer(ctx context.Context, name string, email string, password string) (int64, error) {
	const op = "usecase.CustomerUsecase.SaveCustomer"

	log := usecase.logger.With(
		slog.String("op", op),
		slog.String("email", email),
	)

	log.Debug("saving customer")

	passwordHash, err := usecase.passwordManager.Hash(password)
	if err != nil {
		log.Error("failed to hash password", sl.Err(err))
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	customerID, err := usecase.customerStorage.CreateCustomer(ctx, name, email, passwordHash)
	if err != nil {
		log.Error("failed to save customer", sl.Err(err))
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return customerID, nil
}
