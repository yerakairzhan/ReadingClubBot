package handlers

import (
	"PowerBook/db"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

// SetupHandlers настраивает обработчики команд Telegram
func SetupHandlers(bot *tgbotapi.BotAPI, queries *db.Queries) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	// Получаем последнее обновление
	updatesOffset, err := bot.GetUpdates(u)
	if err != nil {
		log.Fatalf("Failed to get updates: %v", err)
	}
	if len(updatesOffset) > 0 {
		lastUpdate := updatesOffset[len(updatesOffset)-1]
		u.Offset = lastUpdate.UpdateID + 1
	}

	updates := bot.GetUpdatesChan(u)

	// Основной цикл обработки обновлений
	for update := range updates {
		// Обрабатываем callback-запросы
		if update.CallbackQuery != nil {
			HandleCallbackQuery(bot, update)
			continue
		}

		if update.Message != nil && update.Message.IsCommand() {
			switch update.Message.Command() {
			case "register":
				name := RegisterHandler(bot, update)
				log.Printf("User %s registered via message.", name)
			case "start":
				name := Greeting(bot, update)
				log.Printf("%s Started!", name)
			case "admin":
				log.Println("Try to be an Admin")
				AdminCommand(bot, update, queries) // Передаём queries
			default:
				log.Printf("Unknown Command: %s", update.Message.Command())
			}
		}
	}
}
