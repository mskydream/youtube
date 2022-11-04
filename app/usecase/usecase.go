package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mskydream/youtube/app/repository"
	"github.com/mskydream/youtube/model"
)

type Auth interface {
	SignUp(ctx *fiber.Ctx, userProfile *model.UserProfile) (model.UserProfile, error)
}

type UseCase struct {
	Auth
}

func NewUseCase(repo *repository.Repository) *UseCase {
	return &UseCase{
		Auth: NewAuthUseCase(repo.Auth),
	}
}
