package handlers

import (
	"PowerBook/db"
	"context"
	"database/sql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func AdminCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update, queries *db.Queries) {
	userID := strconv.Itoa(int(update.Message.From.ID)) // Используем имя пользователя как идентификатор
	log.Printf("thats how i see an ID : %s", userID)
	// Проверяем права администратора
	if !isAdmin(queries, userID) {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "У вас нет прав доступа к этой команде.")
		bot.Send(msg)
		return
	}

	// Если администратор, выполняем действие
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Добро пожаловать, администратор!")
	bot.Send(msg)
}

func isAdmin(queries *db.Queries, userID string) bool {
	ctx := context.Background()
	adminID, err := queries.GetAdminByID(ctx, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			// Пользователь не найден в таблице администраторов
			return false
		}
		// Логируем ошибку, если она отличается от "нет строк"
		log.Printf("Error checking admin status: %v", err)
		return false
	}
	return adminID == userID
}
