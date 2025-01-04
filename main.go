package main

import (
	"PowerBook/config"
	"PowerBook/db" // Импорт вашего сгенерированного пакета SQLC
	"PowerBook/handlers"
	"PowerBook/services"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Загрузка конфигурации
	config.LoadConfig()

	// Загрузка Базы данных с параметров конфигурации
	envPost := config.DB_HOST
	envPort := config.DB_PORT
	envUser := config.DB_USER
	envPass := config.DB_PASSWORD
	envDbnm := config.DB_NAME

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", envUser, envPass, envPost, envPort, envDbnm)
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error with DB connection: %v", err)
	}
	defer dbConn.Close()

	// Проверка соединения с базой данных
	if err := dbConn.Ping(); err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Инициализация SQLC Queries
	queries := db.New(dbConn)

	// Инициализация сервисов (если необходимо)
	services.InitDB(dbConn)

	// Создание Telegram-бота
	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Настройка маршрутов команд
	handlers.SetupHandlers(bot, queries)
}
