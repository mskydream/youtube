package usecase

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mskydream/youtube/app/repository"
	"github.com/mskydream/youtube/model"
)

type Auth interface {
	SignUp(userProfile *model.UserProfile) (model.UserProfile, error)
	SignIn(input *model.SignIn) (model.GenerateTokenResponse, error)
}

type Channel interface {
	CreateChannel(userId string, channel *model.Channel) (model.Channel, error)
	GetChannels() ([]model.Channel, error)
	GetChannel(id string) (model.Channel, error)
	UpdateChannel(userId string, channel model.Channel) error
	DeleteChannel(id string) error
}

type TelegramBot interface {
	SendMessageLog(groupChatId int64, log string)
}

type UseCase struct {
	Auth
	Channel
	TelegramBot
}

func NewUseCase(repo *repository.Repository, bot *tgbotapi.BotAPI) *UseCase {
	return &UseCase{
		Auth:        NewAuthUseCase(repo.Auth),
		Channel:     NewChannelUseCase(repo.Channel),
		TelegramBot: NewTelegramBotUseCase(bot),
	}
}
