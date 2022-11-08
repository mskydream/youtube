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
	CreateChannel(userId string, channel *model.Channel) (model.Channel, error)
	GetChannels() ([]model.Channel, error)
	GetChannel(id string) (model.Channel, error)
	UpdateChannel(userId string, channel model.Channel) error
	DeleteChannel(id string) error
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
