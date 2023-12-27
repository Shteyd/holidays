package model

import (
	"time"

	"github.com/Shteyd/holidays/src/service/internal/domain/entity"
)

type Customer struct {
	ID           int64     `db:"id" goqu:"skipinsert,skipupdate"`
	Name         string    `db:"name" goqu:"omitempty"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	IsConfirmed  bool      `db:"is_confirmed" goqu:"skipinsert"`
	CreatedAt    time.Time `db:"created_at" goqu:"skipinsert,skipupdate"`
	UpdatedAt    time.Time `db:"updated_at" goqu:"skipinsert"`
}

func NewCustomer(customer entity.Customer) Customer {
	return Customer{
		ID:           customer.ID,
		Name:         customer.Name,
		Email:        customer.Email,
		PasswordHash: customer.Password,
		IsConfirmed:  customer.IsConfirmed,
		CreatedAt:    customer.CreatedAt,
		UpdatedAt:    customer.UpdatedAt,
	}
}

func (model Customer) TableName() string {
	return "customer"
}

func (model Customer) ToEntity() entity.Customer {
	return entity.Customer{
		ID:          model.ID,
		Name:        model.Name,
		Email:       model.Email,
		Password:    model.PasswordHash,
		IsConfirmed: model.IsConfirmed,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}
