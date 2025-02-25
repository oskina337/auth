package user

import (
	"auth/internal/model"
	"context"
)

// Create создаёт новый объект на основе запроса
func (s *serv) Create(ctx context.Context, user *model.User) (int64, error) {
	id, err := s.authRepository.CreateUser(ctx, *user)
	if err != nil {
		return 0, err
	}
	return id, nil
}
