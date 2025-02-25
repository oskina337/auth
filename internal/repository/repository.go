package repository

import (
	"auth/internal/model"
	"context"
)

// для работы с БД (репослой)
type AuthRepository interface {
	CreateUser(ctx context.Context, user model.User) (int64, error)
	GetUser(ctx context.Context, id int64) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
}
