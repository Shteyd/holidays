package entity

import (
	"time"
)

type Customer struct {
	ID          int64
	Name        string
	Email       string
	Password    string
	IsConfirmed bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
