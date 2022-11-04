package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mskydream/youtube/app/repository"
	"github.com/mskydream/youtube/model"
)

type AuthUseCase struct {
	repo repository.Auth
}

func NewAuthUseCase(repo repository.Auth) *AuthUseCase {
	return &AuthUseCase{
		repo: repo,
	}
}

func (a *AuthUseCase) SignUp(ctx *fiber.Ctx, userProfile *model.UserProfile) (model.UserProfile, error) {
	return a.repo.SignUp(ctx, userProfile)
}
