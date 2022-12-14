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
	CreateChannel(userId string, channel *model.Channel) (model.Channel, error)
	GetChannels() ([]model.Channel, error)
	GetChannel(id string) (model.Channel, error)
	GetMemberChannels(id string) ([]model.Channel, error)
	UpdateChannel(userId string, channel model.Channel) error
	DeleteChannel(id string) error
}

type ChannelSubscriber interface {
	CreateChannelSubscriber(userId string, subscriber *model.ChannelSubscriber) error
	GetChannelSubscribers(userId string) ([]model.ChannelSubscriber, error)
	DeleteChannelSubscriber(userId string, channelId string) error
}

type Video interface {
	CreateVideo(video *model.Video) error
	GetVideos() ([]model.Video, error)
	GetVideo(id string) (model.Video, error)
	UpdateVideo(channelId string, video model.Video) error
	DeleteVideo(id string) error
}
type Repository struct {
	Auth
	Channel
	ChannelSubscriber
	Video
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth:              postgres.NewAuthPostgres(db),
		Channel:           postgres.NewChannelPostgres(db),
		ChannelSubscriber: postgres.NewChannelSubscriberPostgres(db),
		Video:             postgres.NewVideoPostgres(db),
	}
}
