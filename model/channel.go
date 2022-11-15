package model

import "time"

type Channel struct {
	Id               string    `db:"id" json:"id"`
	YoutubeAccountId string    `db:"youtube_account_id" json:"youtube_account_id"`
	ChannelName      string    `db:"channel_name" json:"channel_name"`
	CreatedAt        time.Time `db:"created_at" json:"-"`
}
