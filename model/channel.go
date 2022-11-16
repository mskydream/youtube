package model

import "time"

type Channel struct {
	Id               string    `db:"id" json:"id"`
	YoutubeAccountId string    `db:"youtube_account_id" json:"youtube_account_id"`
	ChannelName      string    `db:"channel_name" json:"channel_name"`
	CreatedAt        time.Time `db:"created_at" json:"-"`
}

type ChannelSubscriber struct {
	YoutubeAccountId string    `db:"youtube_account_id" json:"youtube_account_id"`
	YoutubeChannelId string    `db:"youtube_channel_id" json:"youtube_channel_id"`
	CreatedAt        time.Time `db:"created_at" json:"-"`
}
