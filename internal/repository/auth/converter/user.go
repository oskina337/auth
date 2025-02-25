package converter

import (
	"auth/internal/model"
	modelAuth "auth/internal/repository/auth/model"
)

func ToAuthFromRepo(user *modelAuth.User) *model.User {
	return &model.User{
		ID:              user.ID,
		Name:            user.Name,
		Email:           user.Email,
		Password:        user.Password,
		PasswordConfirm: user.PasswordConfirm,
		Role:            model.Role(user.Role),
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
	}
}
