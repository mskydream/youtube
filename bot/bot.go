package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mskydream/youtube/config"
)

func RunTelegramBot(cfg config.Config) {
	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	telegramBotUpdate(bot)
}

func telegramBotUpdate(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "start":
			msg.Text = "Hello!"
		default:
			msg.Text = "What is this?"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
