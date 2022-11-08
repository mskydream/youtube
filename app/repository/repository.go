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

type Channel interface {
	CreateChannel(channel *model.Channel, userId int) (model.Channel, error)
	GetChannels() ([]model.Channel, error)
	GetChannel(id string) (model.Channel, error)
	UpdateChannel(userId int, channel model.Channel) error
}

type Repository struct {
	Auth
	Channel
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth:    postgres.NewAuthPostgres(db),
		Channel: postgres.NewChannelPostgres(db),
	}
}
