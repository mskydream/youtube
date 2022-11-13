package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mskydream/youtube/config"
)

func RunTelegramBot(cfg config.Config, ch chan *tgbotapi.BotAPI) {
	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Telegram bot success connected...")
	bot.Debug = true

	ch <- bot
}
