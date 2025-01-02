package notifications

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"

	"github.com/hablullah/go-prayer"
)

func PrayerLoad(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	userID := update.Message.From.ID
	asiaAlmaty, _ := time.LoadLocation("Asia/Oral")
	schedules, err := prayer.Calculate(prayer.Config{
		Latitude:           43.262964,  // Latitude for Almaty
		Longitude:          76.919729,  // Longitude for Almaty
		Timezone:           asiaAlmaty, // Almaty timezone
		TwilightConvention: prayer.ISNA(),
		AsrConvention:      prayer.Hanafi,
		PreciseToSeconds:   false, // Include seconds in calculation
	}, time.Now().Year())

	if err != nil {
		log.Fatalf("Error calculating prayer times: %v", err)
	}

	if err != nil {
		log.Fatalf("Invalid TIMEZONE value: %v", err)
	}
	now := time.Now().In(asiaAlmaty)
	for _, schedule := range schedules {
		if schedule.Date == now.Format("2006-01-02") {
			makruhTimeStart := schedule.Sunrise.Add(40 * time.Minute)
			makruhTimeEnd := schedule.Zuhr.Add(-40 * time.Minute)
			text := fmt.Sprintf(
				"<b>Время Духа намаз на сегодня:</b>\n\n<b>От:</b> %s 🌄          <b>До:</b> %s 🌇",
				makruhTimeStart.Format("15:04"),
				makruhTimeEnd.Format("15:04"),
			)
			msg := tgbotapi.NewMessage(userID, text)
			msg.ParseMode = "HTML"
			_, err := bot.Send(msg)
			if err != nil {
				log.Fatalf("Error sending message: %v", err)
			}
		}
	}
}
