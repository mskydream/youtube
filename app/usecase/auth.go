package usecase

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mskydream/youtube/app/repository"
	"github.com/mskydream/youtube/config"
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

func (u *AuthUseCase) SignUp(userProfile *model.UserProfile) (model.UserProfile, error) {
	return u.repo.SignUp(userProfile)
}

func (u *AuthUseCase) SignIn(input *model.SignIn) (response model.GenerateTokenResponse, err error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return
	}

	user, err := u.repo.GetUser(input)
	if err != nil {
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	signed, err := token.SignedString([]byte(cfg.Token.Key))
	response.Token = signed
	return
}
