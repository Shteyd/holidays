package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Shteyd/holidays/src/service/internal/domain/entity"
	"github.com/Shteyd/holidays/src/service/internal/lib/logger/sl"
)

func (usecase SessionUsecase) GetSessionList(ctx context.Context, customerID int64) ([]entity.Session, error) {
	const op = "usecase.SessionUsecase.GetSessionList"

	log := usecase.logger.With(
		slog.String("op", op),
	)

	log.Debug("get session list")

	sessionList, err := usecase.sessionStorage.SessionByCustomerID(ctx, customerID)
	if err != nil {
		log.Error("failed to get session list", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return sessionList, nil
}
