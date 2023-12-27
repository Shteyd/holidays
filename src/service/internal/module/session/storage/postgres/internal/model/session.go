package model

import (
	"time"

	"github.com/Shteyd/holidays/src/service/internal/domain/entity"
)

type Session struct {
	ID         int64     `db:"id" goqu:"skipinsert,skipupdate"`
	CustomerID int64     `db:"customer_id"`
	Token      string    `db:"token"`
	UserAgent  string    `db:"user_agent"`
	ExpiresAt  time.Time `db:"expires_at" goqu:"skipinsert"`
	CreatedAt  time.Time `db:"created_at" goqu:"skipinsert,skipupdate"`
	UpdatedAt  time.Time `db:"updated_at" goqu:"skipinsert"`
}

func NewSession(session entity.Session) Session {
	return Session{
		ID:         session.ID,
		CustomerID: session.CustomerID,
		Token:      session.Token,
		UserAgent:  session.UserAgent,
		ExpiresAt:  session.ExpiresAt,
		CreatedAt:  session.CreatedAt,
		UpdatedAt:  session.UpdatedAt,
	}
}

func (model Session) TableName() string {
	return "customer_session"
}

func (model Session) ToEntity() entity.Session {
	return entity.Session{
		ID:         model.ID,
		CustomerID: model.CustomerID,
		Token:      model.Token,
		UserAgent:  model.UserAgent,
		ExpiresAt:  model.ExpiresAt,
		CreatedAt:  model.CreatedAt,
		UpdatedAt:  model.UpdatedAt,
	}
}
