package usecase

import (
	"github.com/mskydream/youtube/app/repository"
	"github.com/mskydream/youtube/model"
)

type Auth interface {
	SignUp(userProfile *model.UserProfile) (model.UserProfile, error)
	SignIn(input *model.SignIn) (model.GenerateTokenResponse, error)
}

type Channel interface {
	CreateChannel(channel *model.Channel, userId int) (model.Channel, error)
	GetChannels() ([]model.Channel, error)
	GetChannel(id string) (model.Channel, error)
	UpdateChannel(userId int, channel model.Channel) error
}

type UseCase struct {
	Auth
	Channel
}

func NewUseCase(repo *repository.Repository) *UseCase {
	return &UseCase{
		Auth:    NewAuthUseCase(repo.Auth),
		Channel: NewChannelUseCase(repo.Channel),
	}
}
