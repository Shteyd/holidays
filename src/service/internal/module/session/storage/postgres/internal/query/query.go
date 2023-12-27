package query

import (
	"log/slog"
	"time"

	"github.com/Shteyd/holidays/src/service/internal/module/session/storage/postgres/internal/model"
	"github.com/doug-martin/goqu/v9"
)

type QueryGenerator struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) QueryGenerator {
	return QueryGenerator{
		logger: logger,
	}
}

func (generator QueryGenerator) RegisterToken(session model.Session) (query string, args []any, err error) {
	query, args, err = goqu.
		Insert(session.TableName()).
		Rows(session).
		Prepared(true).
		ToSQL()

	generator.logger.Debug(query, slog.Any("args", args))

	return
}

func (generator QueryGenerator) GetByToken(token string) (query string, args []any, err error) {
	query, args, err = goqu.
		From((*model.Session)(nil).TableName()).
		Where(
			goqu.C("token").Eq(token),
			goqu.C("expires_at").Gt(time.Now()),
		).
		Prepared(true).
		ToSQL()

	generator.logger.Debug(query, slog.Any("args", args))

	return
}

func (generator QueryGenerator) DeleteByToken(token string) (query string, args []any, err error) {
	query, args, err = goqu.
		Delete((*model.Session)(nil).TableName()).
		Where(goqu.C("token").Eq(token)).
		Prepared(true).
		ToSQL()

	generator.logger.Debug(query, slog.Any("args", args))

	return
}

func (generator QueryGenerator) DeleteByCustomerID(customerID int64) (query string, args []any, err error) {
	query, args, err = goqu.
		Delete((*model.Session)(nil).TableName()).
		Where(goqu.C("customer_id").Eq(customerID)).
		Prepared(true).
		ToSQL()

	generator.logger.Debug(query, slog.Any("args", args))

	return
}

func (generator QueryGenerator) GetByCustomerID(customerID int64) (query string, args []any, err error) {
	query, args, err = goqu.
		From((*model.Session)(nil).TableName()).
		Where(
			goqu.C("customer_id").Eq(customerID),
			goqu.C("expires_at").Gt(time.Now()),
		).
		Prepared(true).
		Order(goqu.C("created_at").Desc()).
		ToSQL()

	generator.logger.Debug(query, slog.Any("args", args))

	return
}
