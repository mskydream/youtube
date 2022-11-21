package model

import "time"

type Video struct {
	Id        string    `db:"id" json:"id"`
	ChannelId string    `db:"youtube_channel_id" json:"channel_id"`
	VideoName string    `db:"video_name" json:"video_name"`
	CreatedAt time.Time `db:"created_at" json:"-"`
}

type VideoView struct {
	Id               string `db:"id" json:"id"`
	YoutubeAccountId string `db:""`
}
