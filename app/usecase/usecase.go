package usecase

import (
	"github.com/mskydream/youtube/app/repository"
	"github.com/mskydream/youtube/model"
)

type Auth interface {
	SignUp(userProfile *model.UserProfile) (model.UserProfile, error)
	SignIn(input *model.SignIn) (model.GenerateTokenResponse, error)
	ParseToken(token string) (model.UserProfile, error)
}

type Channel interface {
	CreateChannel(channel *model.Channel, userId int) (model.Channel, error)
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
