package models

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// GetStartKeyboard возвращает inline-клавиатуру для приветствия
func GetStartKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Зарегистрироваться", "register"),
			tgbotapi.NewInlineKeyboardButtonData("Отмена", "cancel"),
		),
	)
}
