package keyboards

import (
	"MuhasabaDiscipline/db"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func InlineSakinah(queries *db.Queries, userID int64) tgbotapi.InlineKeyboardMarkup {
	ctx := context.Background()
	var keyboardRows [][]tgbotapi.InlineKeyboardButton
	if exists, _ := queries.CheckFajrExists(ctx, userID); !exists {
		keyboardRows = append(keyboardRows, tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("1️⃣    Фаджр", "fajr_done"),
		))
	}

	if exists, _ := queries.CheckDuhaExists(ctx, userID); !exists {
		keyboardRows = append(keyboardRows, tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("2️⃣     Духа", "duha_done"),
		))
	}

	if exists, _ := queries.CheckTafsirExists(ctx, userID); !exists {
		keyboardRows = append(keyboardRows, tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("3️⃣   Тафсир", "tafsir_done"),
		))
	}

	keyboardRows = append(keyboardRows, tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Минимумы", "minimums"),
	))

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(keyboardRows...)
	return inlineKeyboard
}

func InlineMenu() tgbotapi.InlineKeyboardMarkup {
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("1️⃣    Фаджр", "fajr_menu"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("2️⃣     Духа", "duha_menu"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("3️⃣   Тафсир", "tafsir_menu"),
		),
	)
	return inlineKeyboard
}

func InlineStat() tgbotapi.InlineKeyboardMarkup {
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("1️⃣    Фаджр", "stat_fajr"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("2️⃣     Духа", "stat_duha"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("3️⃣   Тафсир", "stat_tafsir"),
		),
	)
	return inlineKeyboard
}
