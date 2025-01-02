package keyboards

import (
	"MuhasabaDiscipline/db"
	"context"
	"database/sql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
	"time"
)

// Fetch habit data from the database
func getDataFromDB(queries *db.Queries, habit string, userID int64, targetTime time.Time) ([]time.Time, error) {
	ctx := context.Background()
	var data []sql.NullTime
	var err error

	switch habit {
	case "Fajr":
		data, err = queries.GetFajrData(ctx, userID)
	case "Duha":
		data, err = queries.GetDuhaData(ctx, userID)
	case "Tafsir":
		data, err = queries.GetTafsirData(ctx, userID)
	default:
		return nil, fmt.Errorf("invalid habit type: %s", habit)
	}

	if err != nil {
		return nil, err
	}

	var validDates []time.Time
	for _, nt := range data {
		if nt.Valid && nt.Time.Year() == targetTime.Year() && nt.Time.Month() == targetTime.Month() {
			validDates = append(validDates, nt.Time)
		}
	}

	return validDates, nil
}

// Generate the calendar
func generateCalendar(habit string, currentTime time.Time, data []time.Time) tgbotapi.InlineKeyboardMarkup {
	currentYear, currentMonth := currentTime.Year(), currentTime.Month()
	monthName := currentMonth.String()
	firstDay := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, time.UTC)
	firstWeekday := firstDay.Weekday()
	if firstWeekday == 0 {
		firstWeekday = 7 // Sunday as the last day
	}
	daysInMonth := time.Date(currentYear, currentMonth+1, 0, 0, 0, 0, 0, time.UTC).Day()

	completedDays := make(map[int]bool)
	for _, date := range data {
		completedDays[date.Day()] = true
	}

	var rows [][]tgbotapi.InlineKeyboardButton

	// Weekday headers
	weekdayRow := []tgbotapi.InlineKeyboardButton{
		tgbotapi.NewInlineKeyboardButtonData("Пн", "empty"),
		tgbotapi.NewInlineKeyboardButtonData("Вт", "empty"),
		tgbotapi.NewInlineKeyboardButtonData("Ср", "empty"),
		tgbotapi.NewInlineKeyboardButtonData("Чт", "empty"),
		tgbotapi.NewInlineKeyboardButtonData("Пт", "empty"),
		tgbotapi.NewInlineKeyboardButtonData("Сб", "empty"),
		tgbotapi.NewInlineKeyboardButtonData("Вс", "empty"),
	}
	rows = append(rows, weekdayRow)

	// Empty buttons for the first week
	var currentRow []tgbotapi.InlineKeyboardButton
	for i := 1; i < int(firstWeekday); i++ {
		currentRow = append(currentRow, tgbotapi.NewInlineKeyboardButtonData(" ", "empty"))
	}

	// Fill days of the month
	now := time.Now()
	for day := 1; day <= daysInMonth; day++ {
		dayStr := strconv.Itoa(day)
		icon := " " // Empty button for days beyond current day or month

		dayDate := time.Date(currentYear, currentMonth, day, 0, 0, 0, 0, time.UTC)
		if dayDate.Before(now) || dayDate.Equal(now) {
			if completedDays[day] {
				icon = "✅" // Mark completed days
			} else {
				icon = "🚫" // Not completed
			}
		}

		currentRow = append(currentRow, tgbotapi.NewInlineKeyboardButtonData(dayStr+icon, habit+"_"+dayStr))

		// Break rows after 7 days
		if len(currentRow) == 7 {
			rows = append(rows, currentRow)
			currentRow = []tgbotapi.InlineKeyboardButton{}
		}
	}

	// Ensure the last row ends with Sunday
	if len(currentRow) > 0 {
		for len(currentRow) < 7 {
			currentRow = append(currentRow, tgbotapi.NewInlineKeyboardButtonData(" ", "empty"))
		}
		rows = append(rows, currentRow)
	}

	// Add navigation buttons
	navigationRow := []tgbotapi.InlineKeyboardButton{
		tgbotapi.NewInlineKeyboardButtonData("<<", fmt.Sprintf("prev_month:%s:%d:%d", habit, currentYear, currentMonth)),
		tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%s %d", monthName, currentYear), "current_month"),
		tgbotapi.NewInlineKeyboardButtonData(">>", fmt.Sprintf("next_month:%s:%d:%d", habit, currentYear, currentMonth)),
	}
	rows = append(rows, navigationRow)

	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

func UpdateCalendar(bot *tgbotapi.BotAPI, queries *db.Queries, update tgbotapi.Update, habit string, currentTime time.Time) {
	var chatID int64
	var messageID int

	if update.CallbackQuery != nil {
		chatID = update.CallbackQuery.Message.Chat.ID
		messageID = update.CallbackQuery.Message.MessageID
	} else if update.Message != nil {
		chatID = update.Message.Chat.ID
	} else {
		log.Println("Neither CallbackQuery nor Message is present in update")
		return
	}

	// Fetch data for the selected habit
	data, err := getDataFromDB(queries, habit, chatID, currentTime)
	if err != nil {
		log.Printf("Error fetching habit data: %v", err)
		if update.CallbackQuery != nil {
			sendCallbackMessage(bot, update.CallbackQuery.ID, "Ошибка загрузки данных.")
		}
		return
	}

	// Generate the calendar
	calendar := generateCalendar(habit, currentTime, data)

	if update.CallbackQuery != nil {
		// Edit the existing message for CallbackQuery
		edit := tgbotapi.NewEditMessageReplyMarkup(chatID, messageID, calendar)
		_, err = bot.Request(edit)
		if err != nil {
			log.Printf("Error updating calendar message: %v", err)
		}
	} else {
		// Send a new message for regular commands
		msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Ваш календарь для %s:", habit))
		msg.ReplyMarkup = calendar
		_, err = bot.Send(msg)
		if err != nil {
			log.Printf("Error sending calendar message: %v", err)
		}
	}
}

// Handle callback data for navigation and updates
func HandleCallback(bot *tgbotapi.BotAPI, queries *db.Queries, update tgbotapi.Update) {
	callback := update.CallbackQuery
	data := callback.Data

	// Parse callback data (format: action:habit:year:month)
	parts := strings.Split(data, ":")
	if len(parts) < 4 {
		log.Printf("Invalid callback data format: %s", data)
		sendCallbackMessage(bot, callback.ID, "Ошибка обработки данных.")
		return
	}

	action := parts[0]
	habit := parts[1]
	year, _ := strconv.Atoi(parts[2])
	month, _ := strconv.Atoi(parts[3])
	currentTime := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)

	// Adjust the date based on action
	switch action {
	case "prev_month":
		currentTime = currentTime.AddDate(0, -1, 0)
	case "next_month":
		currentTime = currentTime.AddDate(0, 1, 0)
	case "current_month":
		currentTime = time.Now()
	default:
		log.Printf("Unknown callback action: %s", action)
		sendCallbackMessage(bot, callback.ID, "Неизвестное действие.")
		return
	}

	// Update the calendar
	UpdateCalendar(bot, queries, update, habit, currentTime)
}

// Helper function to send callback response
func sendCallbackMessage(bot *tgbotapi.BotAPI, callbackID, text string) {
	callback := tgbotapi.NewCallback(callbackID, text)
	if _, err := bot.Request(callback); err != nil {
		log.Printf("Failed to send callback message: %v", err)
	}
}
