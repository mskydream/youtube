package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/mskydream/youtube/app/repository/postgres"
	"github.com/mskydream/youtube/model"
)

type Auth interface {
	SignUp(userProfile *model.UserProfile) (model.UserProfile, error)
	GetUser(input *model.SignIn) (model.UserProfile, error)
}

type Repository struct {
	Auth
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: postgres.NewAuthPostgres(db),
	}
}
