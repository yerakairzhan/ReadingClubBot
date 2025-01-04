package handlers

import (
	"PowerBook/models"
	"PowerBook/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func RegisterHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) string {
	var chatID int64
	var userID int
	var username string

	// Определяем источник вызова
	if update.Message != nil {
		chatID = update.Message.Chat.ID
		userID = int(update.Message.From.ID)
		username = update.Message.From.UserName
	} else if update.CallbackQuery != nil {
		chatID = update.CallbackQuery.Message.Chat.ID
		userID = int(update.CallbackQuery.From.ID)
		username = update.CallbackQuery.From.UserName
	}

	// Создаем пользователя
	user := models.User{
		UserID:   int64(userID),
		Username: username,
	}

	// Сохраняем пользователя
	err := services.SaveUser(user)
	if err != nil {
		msg := tgbotapi.NewMessage(chatID, "Ошибка при регистрации. Возможно, вы уже зарегистрированы.")
		bot.Send(msg)
		if update.CallbackQuery != nil {
			callbackResponse := tgbotapi.NewCallback(update.CallbackQuery.ID, "Ошибка при регистрации.")
			bot.Request(callbackResponse)
		}
		return "Error with Registration"
	}

	// Успешная регистрация
	msg := tgbotapi.NewMessage(chatID, "Вы успешно зарегистрировались!")
	bot.Send(msg)
	if update.CallbackQuery != nil {
		callbackResponse := tgbotapi.NewCallback(update.CallbackQuery.ID, "Успешная регистрация.")
		bot.Request(callbackResponse)
	}

	return user.Username
}

// Помощник для ответа на callback-запрос
func respondToCallback(bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery, text string) {
	callbackResponse := tgbotapi.NewCallback(callbackQuery.ID, text)
	_, _ = bot.Request(callbackResponse)
}
