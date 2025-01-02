package handlers

import (
	"MuhasabaDiscipline/db"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SetupHandlers(bot *tgbotapi.BotAPI, queries *db.Queries) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil && update.Message.IsCommand() {
			HandleCommand(bot, update, queries)
		} else if update.CallbackQuery != nil {
			HandleInline(bot, update, queries)
		}
	}
}
