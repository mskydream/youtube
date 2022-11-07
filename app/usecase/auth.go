package usecase

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mskydream/youtube/app/repository"
	"github.com/mskydream/youtube/config"
	"github.com/mskydream/youtube/model"
)

type tokenClaims struct {
	jwt.StandardClaims
	User model.UserProfile
}

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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(cfg.Token.LifeTime)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user,
	})

	signed, err := token.SignedString([]byte(cfg.Token.Key))
	response.Token = signed
	return
}

func (u *AuthUseCase) ParseToken(accessToken string) (model.UserProfile, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return model.UserProfile{}, err
	}

	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(cfg.Token.Key), nil
	})
	if err != nil {
		return model.UserProfile{}, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return model.UserProfile{}, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.User, nil
}
