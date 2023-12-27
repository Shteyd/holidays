package postgres

import (
	"context"
	"fmt"

	"github.com/Shteyd/holidays/src/service/internal/domain/entity"
	"github.com/Shteyd/holidays/src/service/internal/module/session/storage/postgres/internal/model"
)

func (postgres PostgresStorage) RegisterToken(ctx context.Context, session entity.Session) error {
	const op = "storage.postgres.RegisterToken"

	model := model.NewSession(session)

	query, args, err := postgres.queries.RegisterToken(model)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = postgres.conn.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
