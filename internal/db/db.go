package db

import (
	"auth/internal/config"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Client - интерфейс для работы с БД
type Client interface {
	DB() *pgxpool.Pool
	Close()
}

// client - структура, реализующая интерфейс Client
type client struct {
	pool *pgxpool.Pool
}

// DB - возвращает пул соединений
func (c *client) DB() *pgxpool.Pool {
	return c.pool
}

// Close - закрывает соединение с БД
func (c *client) Close() {
	if c.pool != nil {
		c.pool.Close()
		fmt.Println("🛑 Соединение с БД закрыто")
	}
}

// NewClient - создаёт новый клиент БД
func NewClient() (Client, error) {
	dbDSN := config.LoadConfig()

	pool, err := pgxpool.New(context.Background(), dbDSN)
	if err != nil {
		return nil, fmt.Errorf("❌ ошибка подключения к БД: %w", err)
	}

	fmt.Println("✅ Подключение к БД установлено")
	return &client{pool: pool}, nil
}
