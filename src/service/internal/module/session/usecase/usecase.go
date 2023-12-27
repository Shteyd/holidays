package usecase

import (
	"context"
	"errors"
	"log/slog"

	"github.com/Shteyd/holidays/src/service/internal/domain/entity"
)

var (
	ErrSessionNotFound = errors.New("session not found")
)

type TokenManager interface {
	GenerateToken(ctx context.Context) (string, error)
}

type SessionStorage interface {
	DeleteByCustomerID(ctx context.Context, customerID int64) error
	DeleteByToken(ctx context.Context, token string) error
	RegisterToken(ctx context.Context, session entity.Session) error
	SessionByCustomerID(ctx context.Context, customerID int64) ([]entity.Session, error)
	SessionByToken(ctx context.Context, token string) (entity.Session, error)
}

type SessionUsecase struct {
	logger         *slog.Logger
	tokenManager   TokenManager
	sessionStorage SessionStorage
}

func New(
	logger *slog.Logger,
	tokenManager TokenManager,
	sessionStorage SessionStorage,
) SessionUsecase {
	return SessionUsecase{
		logger:         logger,
		tokenManager:   tokenManager,
		sessionStorage: sessionStorage,
	}
}
