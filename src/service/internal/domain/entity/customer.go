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

type Session struct {
	ID         int64
	CustomerID int64
	Token      string
	UserAgent  string
	ExpiresAt  time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
