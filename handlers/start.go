package handlers

import (
	"PowerBook/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

func Greeting(bot *tgbotapi.BotAPI, update tgbotapi.Update) string {
	user := models.User{
		UserID:   update.Message.From.ID,
		Username: update.Message.From.UserName,
		Name:     update.Message.From.FirstName,
	}

	// Отправляем приветствие
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, "+user.Name)
	_, _ = bot.Send(msg)

	// Пауза
	time.Sleep(1 * time.Second)

	// Создаем inline-клавиатуру
	msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Меня зовут Пауэр и я хочу помочь тебе!\n\n/register если хочешь зарегестрироваться!")
	msg.ReplyMarkup = models.GetStartKeyboard()
	_, _ = bot.Send(msg)

	return user.Username
}
