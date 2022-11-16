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

type ChannelSubscriber interface {
	CreateChannelSubscriber(userId string, subscriber *model.ChannelSubscriber) error
	GetChannelSubscribers(userId string) ([]model.ChannelSubscriber, error)
	DeleteChannelSubscriber(userId string, channelId string) error
}

type TelegramBot interface {
	SendMessageLog(groupChatId int64, log string)
}

type UseCase struct {
	Auth
	Channel
	ChannelSubscriber
	TelegramBot
}

func NewUseCase(repo *repository.Repository, bot *tgbotapi.BotAPI) *UseCase {
	return &UseCase{
		TelegramBot:       NewTelegramBotUseCase(bot),
		Auth:              NewAuthUseCase(repo.Auth),
		Channel:           NewChannelUseCase(repo.Channel),
		ChannelSubscriber: NewChannelSubscriberUseCase(repo.ChannelSubscriber),
	}
}
