package usecase

import (
	"context"
	"errors"
	"log/slog"

	"github.com/Shteyd/holidays/src/service/internal/domain/entity"
)

var (
	ErrCustomerNotFound = errors.New("customer not found")
)

type passwordManager interface {
	Hash(password string) (string, error)
	Compare(hashedPassword string, password string) error
}

type customerStorage interface {
	CreateCustomer(
		ctx context.Context,
		name string,
		email string,
		passwordHash string,
	) (int64, error)
	DeleteCustomer(
		ctx context.Context,
		customerID int64,
	) error
	CustomerByID(
		ctx context.Context,
		customerID int64,
	) (entity.Customer, error)
	CustomerByCreds(
		ctx context.Context,
		email string,
		password string,
	) (entity.Customer, error)
	UpdateCustomer(
		ctx context.Context,
		customerID int64,
		name string,
		email string,
		passwordHash string,
	) error
}

type CustomerUsecase struct {
	logger          *slog.Logger
	passwordManager passwordManager
	customerStorage customerStorage
}

func New(
	logger *slog.Logger,
	passwordManager passwordManager,
	customerStorage customerStorage,
) CustomerUsecase {
	return CustomerUsecase{
		logger:          logger,
		passwordManager: passwordManager,
		customerStorage: customerStorage,
	}
}
