package auth

import (
	"auth/internal/db"
	"auth/internal/model"
	"auth/internal/repository"
	"auth/internal/repository/auth/converter"
	modelAuth "auth/internal/repository/auth/model"
	"strings"

	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx"
)

// храним в константе имя таблицы и ее колонки
const (
	tableName = "users"

	idColumn            = "id"
	nameColumn          = "name"
	emailColumn         = "email"
	password_hashColumn = "password_hash"
	roleColumn          = "role"
	created_atColumn    = "created_at"
	updated_atColumn    = "updated_at"
)

var _ repository.AuthRepository = (*repo)(nil)

type repo struct {
	db db.Client
}

// NewRepository - создание нового репозитория
func NewRepository(db db.Client) repository.AuthRepository {
	return &repo{db: db}
}

// CreateUser - добавление нового пользователя
func (r *repo) CreateUser(ctx context.Context, user model.User) (int64, error) {

	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)`
	err := r.db.DB().QueryRow(ctx, checkQuery, user.Email).Scan(&exists)
	if err != nil {
		return 0, fmt.Errorf("failed to check user existence: %w", err)
	}
	if exists {
		return 0, fmt.Errorf("user with email %s already exists", user.Email)
	}

	var id int64
	query := fmt.Sprintf(`INSERT INTO %s (%s, %s, %s, %s, %s, %s) 
                          VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING %s`,
		tableName, nameColumn, emailColumn, password_hashColumn, roleColumn, created_atColumn, updated_atColumn, idColumn)

	err = r.db.DB().QueryRow(ctx, query, user.Name, user.Email, user.Password, user.Role).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}
	return id, nil
}

// GetUser - получает пользователя по ID
func (r *repo) GetUser(ctx context.Context, id int64) (*model.User, error) {
	var user modelAuth.User

	query := fmt.Sprintf(`SELECT %s, %s, %s, %s, %s, %s, %s FROM %s WHERE %s = $1`,
		idColumn, nameColumn, emailColumn, password_hashColumn, roleColumn, created_atColumn, updated_atColumn, tableName, idColumn)

	err := r.db.DB().QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &model.User{}, fmt.Errorf("user with id %d not found", id)
		}
		return &model.User{}, fmt.Errorf("failed to get user: %w", err)
	}

	return converter.ToAuthFromRepo(&user), nil
}

// UpdateUser - обновляет данные полльзователя в БД
func (r *repo) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	columns := []string{nameColumn, emailColumn, password_hashColumn, roleColumn}
	setClauses := []string{}
	args := []interface{}{}

	// Формируем список SET-частей запроса и аргументы
	for i, col := range columns {
		setClauses = append(setClauses, fmt.Sprintf("%s = COALESCE($%d, %s)", col, i+1, col))
		args = append(args, userFieldValue(user, i))
	}

	// Добавляем updated_at и ID
	query := fmt.Sprintf(`UPDATE %s SET %s, %s = NOW() WHERE %s = $%d`,
		tableName,
		strings.Join(setClauses, ", "),
		updated_atColumn,
		idColumn, len(columns)+1)

	args = append(args, user.ID) // Добавляем ID в аргументы

	_, err := r.db.DB().Exec(ctx, query, args...)
	if err != nil {
		return &model.User{}, fmt.Errorf("failed to updat user: %w", err)
	}

	return user, nil
}

// Для заполнения обновлений
func userFieldValue(user *model.User, index int) interface{} {
	switch index {
	case 0:
		if user.Name == "" {
			return nil
		}
		return user.Name
	case 1:
		if user.Email == "" {
			return nil
		}
		return user.Email
	case 2:
		if user.Password == "" {
			return nil
		}
		return user.Password
	case 3:
		return user.Role
	}
	return nil
}
