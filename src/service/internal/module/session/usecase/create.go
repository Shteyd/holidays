package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Shteyd/holidays/src/service/internal/domain/entity"
	"github.com/Shteyd/holidays/src/service/internal/lib/logger/sl"
)

func (usecase SessionUsecase) CreateSession(
	ctx context.Context,
	customerID int64,
	userAgent string,
) error {
	const op = "usecase.SessionUsecase.CreateSession"

	log := usecase.logger.With(
		slog.String("op", op),
	)

	log.Debug("create session")

	token, err := usecase.tokenManager.GenerateToken(ctx)
	if err != nil {
		log.Error("failed to generate token", sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	session := entity.Session{
		CustomerID: customerID,
		Token:      token,
		UserAgent:  userAgent,
	}

	err = usecase.sessionStorage.RegisterToken(ctx, session)
	if err != nil {
		log.Error("failed to register token", sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
