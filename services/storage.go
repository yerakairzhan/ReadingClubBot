package services

import (
	"PowerBook/db"
	"PowerBook/models"
	"context"
	"database/sql"
	"log"
	"strconv"
)

var users []models.User
var queries *db.Queries

func InitDB(dbConn *sql.DB) {
	queries = db.New(dbConn)
}

func SaveUser(user models.User) error {
	users = append(users, user)
	params := db.CreateUserParams{
		UserID:   strconv.FormatInt(int64(user.UserID), 10),
		Username: user.Username,
	}

	ctx := context.Background()
	err := queries.CreateUser(ctx, params)
	if err != nil {
		log.Printf("Error saving user to database: %v", err)
		return err
	}
	return nil
}
