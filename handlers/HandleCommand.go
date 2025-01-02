package handlers

import (
	"MuhasabaDiscipline/db"
	"MuhasabaDiscipline/keyboards"
	"MuhasabaDiscipline/notifications"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"
)

func HandleCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update, queries *db.Queries) {
	userID := update.Message.Chat.ID
	switch update.Message.Command() {
	case "start":
		text := "<b>Салам Алейкум</b>, я личный бот Ерасыла ,который помогает ему <b>воспитать нафс по воле Аллаха</>"
		msg := tgbotapi.NewMessage(userID, text)
		msg.ParseMode = "HTML"
		_, err := bot.Send(msg)
		if err != nil {
			log.Println("НЕ МОГУ ОТПРАВИТЬ НА СТАРТЕ")
		}
	case "muhasaba":
		sakina_msg(bot, queries, userID)
	case "menu":
		text := "<b>Постоянное Меню</b>\n\n⚖️/muhasaba\n🌅/duha"
		msg := tgbotapi.NewMessage(userID, text)
		msg.ParseMode = "HTML"
		msg.ReplyMarkup = keyboards.InlineMenu()
		_, err := bot.Send(msg)
		if err != nil {
			log.Println("НЕ МОГУ ОТПРАВИТЬ САКИНА")
		}
	case "duha":
		notifications.PrayerLoad(bot, update)
	case "stat_fajr":
		keyboards.UpdateCalendar(bot, queries, update, "Fajr", time.Now())
	case "stat_duha":
		keyboards.UpdateCalendar(bot, queries, update, "Duha", time.Now())
	case "stat_tafsir":
		keyboards.UpdateCalendar(bot, queries, update, "Tafsir", time.Now())

	case "stat":
		text := "<b>Любимые дела перед Аллахом</b> те, которые совершаемы <b>постоянно</b> (на протяжении наиболее продолжительного периода), хотя [будь они] и малочисленны."
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		msg.ParseMode = "HTML"
		msg.ReplyMarkup = keyboards.InlineStat()
		_, err := bot.Send(msg)
		if err != nil {
			log.Println(err)
		}
	}
}

func sakina_msg(bot *tgbotapi.BotAPI, queries *db.Queries, userID int64) {
	text := "Ерасыл, <b>постоянство - ключ к любви от Аллаха.</b> Как бы не было сложно, <b>не ищи оправдания.</b> Ты - адвокат для окружающих и судья для себя!"
	msg := tgbotapi.NewMessage(userID, text)
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = keyboards.InlineSakinah(queries, userID)
	_, err := bot.Send(msg)
	if err != nil {
		log.Println("НЕ МОГУ ОТПРАВИТЬ САКИНА")
	}
}
