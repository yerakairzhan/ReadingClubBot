package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

// HandleCallbackQuery обрабатывает callback-запросы
func HandleCallbackQuery(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.CallbackQuery == nil {
		return
	}

	// Логируем callback data
	log.Printf("Callback data received: %s", update.CallbackQuery.Data)

	switch update.CallbackQuery.Data {
	case "register":
		RegisterHandler(bot, update)
	default:
		log.Printf("Unknown callback data: %s", update.CallbackQuery.Data)
	}
}

// handleRegister обработка нажатия "Зарегистрироваться"

// handleCancel обработка нажатия "Отмена"
func handleCancel(bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery) {
	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Регистрация отменена.")
	_, _ = bot.Send(msg)
}

// sendReply отправляет сообщение в чат
func sendReply(bot *tgbotapi.BotAPI, chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	_, _ = bot.Send(msg)
}
