package usecase

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/Shteyd/holidays/src/service/internal/domain/entity"
	"github.com/Shteyd/holidays/src/service/internal/lib/logger/sl"
	"github.com/Shteyd/holidays/src/service/internal/module/session/storage"
)

func (usecase SessionUsecase) GetSessionByToken(ctx context.Context, token string) (entity.Session, error) {
	const op = "usecase.SessionUsecase.GetSessionByToken"

	log := usecase.logger.With(
		slog.String("op", op),
	)

	log.Debug("get session by token")

	session, err := usecase.sessionStorage.SessionByToken(ctx, token)
	if err != nil {
		if errors.Is(err, storage.ErrSessionNotFound) {
			log.Warn("session not found")
			return entity.Session{}, ErrSessionNotFound
		}

		log.Error("failed to get session by token", sl.Err(err))
		return entity.Session{}, fmt.Errorf("%s: %w", op, err)
	}

	return session, nil
}
