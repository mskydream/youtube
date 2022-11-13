package usecase

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBotUseCase struct {
	bot *tgbotapi.BotAPI
}

func NewTelegramBotUseCase(bot *tgbotapi.BotAPI) *TelegramBotUseCase {
	return &TelegramBotUseCase{
		bot: bot,
	}
}

func (u *TelegramBotUseCase) SendMessageLog(groupChatId int64, log string) {
	msg := tgbotapi.NewMessage(groupChatId, log)
	u.bot.Send(msg)
}
