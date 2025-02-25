package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID              int64
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            Role
	CreatedAt       time.Time
	UpdatedAt       sql.NullTime
}

type Role int32

const (
	Role_USER  Role = 0
	Role_ADMIN Role = 1
)
