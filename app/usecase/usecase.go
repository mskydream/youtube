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

type UseCase struct {
	Auth
}

func NewUseCase(repo *repository.Repository) *UseCase {
	return &UseCase{
		Auth: NewAuthUseCase(repo.Auth),
	}
}
