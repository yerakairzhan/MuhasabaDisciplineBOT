package handlers

import (
	"MuhasabaDiscipline/db"
	"MuhasabaDiscipline/keyboards"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
	"time"
)

func HandleInline(bot *tgbotapi.BotAPI, update tgbotapi.Update, queries *db.Queries) {
	callbackData := update.CallbackQuery.Data
	userID := int64(update.CallbackQuery.From.ID) // Telegram user ID
	username := update.CallbackQuery.From.UserName
	ctx := context.Background()
	log.Println(callbackData)
	log.Println(strings.HasPrefix(callbackData, "next_month_"))
	// Ensure the user exists in the users table
	err := queries.InsertUser(ctx, db.InsertUserParams{
		UserID:   userID,
		Username: username,
	})
	if err != nil {
		log.Printf("Failed to insert user %d: %v", userID, err)
		sendCallbackMessage(bot, update.CallbackQuery.ID, "Ошибка при добавлении пользователя.")
		return
	}

	if strings.HasPrefix(callbackData, "next_month") || strings.HasPrefix(callbackData, "prev_month") {
		keyboards.HandleCallback(bot, queries, update)
		return
	} else {

		switch callbackData {
		case "fajr_done":
			err := queries.UpsertFajrDone(ctx, userID)
			if err != nil {
				log.Printf("Failed to upsert fajr as done for user %d: %v", userID, err)
				sendCallbackMessage(bot, update.CallbackQuery.ID, "Ошибка при обновлении Fajr.")
				return
			}
			if update.CallbackQuery.Message != nil {
				err = deleteMessage(bot, update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
				if err != nil {
					log.Printf("Failed to delete message for user %d: %v", userID, err)
				}
			} else {
				log.Println("No message found to delete.")
			}
			sakina_msg(bot, queries, userID)
			sendCallbackMessage(bot, update.CallbackQuery.ID, "Fajr отмечен как выполненный!")

		case "duha_done":
			err := queries.UpsertDuhaDone(ctx, userID)
			if err != nil {
				log.Printf("Failed to upsert duha as done for user %d: %v", userID, err)
				sendCallbackMessage(bot, update.CallbackQuery.ID, "Ошибка при обновлении duha.")
				return
			}
			if update.CallbackQuery.Message != nil {
				err = deleteMessage(bot, update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
				if err != nil {
					log.Printf("Failed to delete message for user %d: %v", userID, err)
				}
			} else {
				log.Println("No message found to delete.")
			}
			sakina_msg(bot, queries, userID)
			sendCallbackMessage(bot, update.CallbackQuery.ID, "Duha отмечен как выполненный!")
		case "tafsir_done":
			err := queries.UpsertTafsirDone(ctx, userID)
			if err != nil {
				log.Printf("Failed to upsert tafsir as done for user %d: %v", userID, err)
				sendCallbackMessage(bot, update.CallbackQuery.ID, "Ошибка при обновлении tafsir.")
				return
			}
			if update.CallbackQuery.Message != nil {
				err = deleteMessage(bot, update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
				if err != nil {
					log.Printf("Failed to delete message for user %d: %v", userID, err)
				}
			} else {
				log.Println("No message found to delete.")
			}
			sakina_msg(bot, queries, userID)
			sendCallbackMessage(bot, update.CallbackQuery.ID, "Tafsir отмечен как выполненный!")
		case "minimums":
			text := `<b>Фаджр</b> : Совершить своевременно` + "\n" +
				`<b>Духа</b> : Совершить до приближения солнца к зениту` + "\n" +
				`<b>Тафсир</b> : Прочесть 10 страниц Azan.ru`
			msg := tgbotapi.NewMessage(userID, text)
			msg.ParseMode = "HTML"
			_, err := bot.Send(msg)
			if err != nil {
				log.Println("НЕ СМОГ ОТПРАВИТЬ МИНИМУМЫ")
			}
			sendCallbackMessage(bot, update.CallbackQuery.ID, "ТОПИ")
		case "fajr_menu":
			err := queries.UpsertFajrDone(ctx, userID)
			if err != nil {
				log.Printf("Failed to upsert fajr as done for user %d: %v", userID, err)
				sendCallbackMessage(bot, update.CallbackQuery.ID, "Ошибка при обновлении duha.")
				return
			}
		case "duha_menu":
			err := queries.UpsertDuhaDone(ctx, userID)
			if err != nil {
				log.Printf("Failed to upsert duha as done for user %d: %v", userID, err)
				sendCallbackMessage(bot, update.CallbackQuery.ID, "Ошибка при обновлении duha.")
				return
			}
		case "tafsir_menu":
			err := queries.UpsertTafsirDone(ctx, userID)
			if err != nil {
				log.Printf("Failed to upsert tafsir as done for user %d: %v", userID, err)
				sendCallbackMessage(bot, update.CallbackQuery.ID, "Ошибка при обновлении duha.")
				return
			}
		case "stat_fajr":
			keyboards.UpdateCalendar(bot, queries, update, "Fajr", time.Now())
		case "stat_duha":
			keyboards.UpdateCalendar(bot, queries, update, "Duha", time.Now())
		case "stat_tafsir":
			keyboards.UpdateCalendar(bot, queries, update, "Tafsir", time.Now())
		default:
			sendCallbackMessage(bot, update.CallbackQuery.ID, "Неизвестная команда.")
		}
	}
}

// Helper function to send callback response
func sendCallbackMessage(bot *tgbotapi.BotAPI, callbackID, text string) {
	callback := tgbotapi.NewCallback(callbackID, text)
	if _, err := bot.Request(callback); err != nil {
		log.Printf("Failed to send callback message: %v", err)
	}
}

func deleteMessage(bot *tgbotapi.BotAPI, chatID int64, messageID int) error {
	deleteConfig := tgbotapi.NewDeleteMessage(chatID, messageID)
	_, err := bot.Request(deleteConfig)
	if err != nil {
		return fmt.Errorf("failed to delete message: %w", err)
	}
	return nil
}
