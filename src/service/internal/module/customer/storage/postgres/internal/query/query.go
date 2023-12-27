package query

import (
	"log/slog"

	"github.com/Shteyd/holidays/src/service/internal/module/customer/storage/postgres/internal/model"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

type QueryGenerator struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) QueryGenerator {
	return QueryGenerator{
		logger: logger,
	}
}

func (generator QueryGenerator) CreateCustomer(customer model.Customer) (query string, args []any, err error) {
	query, args, err = goqu.
		Insert(customer.TableName()).
		Rows(customer).
		Returning(goqu.C("id")).
		Prepared(true).
		ToSQL()

	generator.logger.Debug(query, slog.Any("args", args))

	return
}

func (generator QueryGenerator) DeleteCustomer(customerID int64) (query string, args []any, err error) {
	query, args, err = goqu.
		Delete((*model.Customer)(nil).TableName()).
		Where(
			goqu.C("id").Eq(customerID),
		).
		Prepared(true).
		ToSQL()

	generator.logger.Debug(query, slog.Any("args", args))

	return
}

func (generator QueryGenerator) CustomerByID(customerID int64) (query string, args []any, err error) {
	model := (*model.Customer)(nil)

	query, args, err = goqu.
		Select(model).
		From(model.TableName()).
		Where(
			goqu.C("id").Eq(customerID),
		).
		Prepared(true).
		ToSQL()

	generator.logger.Debug(query, slog.Any("args", args))

	return
}

func (generator QueryGenerator) CustomerByCreds(customer model.Customer) (query string, args []any, err error) {
	query, args, err = goqu.
		Select(customer).
		From(customer.TableName()).
		Where(
			goqu.C("email").Eq(customer.Email),
			goqu.C("password_hash").Eq(customer.PasswordHash),
		).
		Prepared(true).
		ToSQL()

	generator.logger.Debug(query, slog.Any("args", args))

	return
}

func (generator QueryGenerator) UpdateCustomer(customer model.Customer) (query string, args []any, err error) {
	// TODO: soft update

	sqlquery := goqu.
		Update(customer.TableName()).
		Set([]exp.UpdateExpression{
			goqu.C("name").Set(customer.Name),
			goqu.C("email").Set(customer.Email),
			goqu.C("password_hash").Set(customer.PasswordHash),
			goqu.C("updated_at").Set(goqu.L("NOW()")),
		}).
		Where(
			goqu.C("id").Eq(customer.ID),
		)

	query, args, err = sqlquery.
		Prepared(true).
		ToSQL()

	generator.logger.Debug(query, slog.Any("args", args))

	return
}
