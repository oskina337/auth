package user

import (
	"auth/internal/repository"
	"auth/internal/service"
)

type serv struct {
	authRepository repository.AuthRepository
}

func NewSerice(authRepository repository.AuthRepository) service.AuthSerther {
	return &serv{
		authRepository: authRepository,
	}
}
