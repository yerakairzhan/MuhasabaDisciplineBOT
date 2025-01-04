package main

import (
	"MuhasabaDiscipline/config"
	"MuhasabaDiscipline/db"
	"MuhasabaDiscipline/handlers"
	"database/sql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	config.LoadConfig()
	BotToken := config.BOTAPI

	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	envPost := config.DB_HOST
	envPort := config.DB_PORT
	envUser := config.DB_USER
	envPass := config.DB_PASSWORD
	envDbnm := config.DB_NAME

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", envUser, envPass, envPost, envPort, envDbnm)
	dbConn, err := sql.Open("postgres", connStr)
	log.Println(connStr)
	if err != nil {
		log.Fatalf("Error with DB connection: %v", err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	handlers.SetupHandlers(bot, queries)
}
