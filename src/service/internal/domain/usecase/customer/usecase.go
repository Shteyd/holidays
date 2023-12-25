package customer

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

type customerProvider interface {
	CustomerByID(
		ctx context.Context,
		customerID int64,
	) (entity.Customer, error)
	CustomerByCreds(
		ctx context.Context,
		email string,
		password string,
	) (entity.Customer, error)
}

type customerManager interface {
	CreateCustomer(
		ctx context.Context,
		name string,
		email string,
		passwordHash string,
	) (int64, error)
	UpdateCustomer(
		ctx context.Context,
		customerID int64,
		name string,
		email string,
		passwordHash string,
	) error
	DeleteCustomer(
		ctx context.Context,
		customerID int64,
	) error
}

type CustomerUsecase struct {
	logger           *slog.Logger
	customerProvider customerProvider
	customerManager  customerManager
	passwordManager  passwordManager
}

func New(
	logger *slog.Logger,
	customerProvider customerProvider,
	customerManager customerManager,
	passwordManager passwordManager,
) CustomerUsecase {
	return CustomerUsecase{
		logger:           logger,
		customerProvider: customerProvider,
		customerManager:  customerManager,
		passwordManager:  passwordManager,
	}
}
