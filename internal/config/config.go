package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadConfig загружает переменные из .env
func LoadConfig() string {
	err := godotenv.Load("../.env") // Загружаем .env
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	// Читаем переменные
	dbName := os.Getenv("PG_DATABASE_NAME")
	dbUser := os.Getenv("PG_USER")
	dbPassword := os.Getenv("PG_PASSWORD")
	dbPort := os.Getenv("PG_PORT")
	dbHost := "localhost" // Можно тоже вынести в .env

	// Формируем строку подключения
	dbDSN := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		dbHost, dbPort, dbName, dbUser, dbPassword)

	return dbDSN
}
