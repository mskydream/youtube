package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/mskydream/youtube/model"
)

type ChannelSubscriberPostgres struct {
	db *sqlx.DB
}

func NewChannelSubscriberPostgres(db *sqlx.DB) *ChannelSubscriberPostgres {
	return &ChannelSubscriberPostgres{
		db: db,
	}
}

func (r *ChannelSubscriberPostgres) CreateChannelSubscriber(userId string, subscriber *model.ChannelSubscriber) error {
	_, err := r.db.Query(`INSERT INTO channel_subscriber(youtube_account_id, youtube_channel_id, created_at) VALUES ($1, $2, now())`, userId, subscriber.YoutubeChannelId)
	return err
}

func (r *ChannelSubscriberPostgres) GetChannelSubscribers(userId string) (channelSubscriber []model.ChannelSubscriber, err error) {
	return channelSubscriber, r.db.Select(&channelSubscriber, `SELECT youtube_account_id, youtube_channel_id, created_at FROM channel_subscriber WHERE youtube_account_id = $1`, userId)
}

func (r *ChannelSubscriberPostgres) DeleteChannelSubscriber(userId string, channelId string) error {
	_, err := r.db.Exec(`DELETE FROM channel_subscriber WHERE youtube_account_id = $1 and youtube_channel_id = $2`, userId, channelId)
	return err
}
