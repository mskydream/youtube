package model

import "time"

type Video struct {
	ID        int       `db:"id" json:"id"`
	ChannelID string    `db:"channel_id" json:"channel_id"`
	VideoName string    `db:"video_name" json:"video_name"`
	CreatedAt time.Time `db:"created_at" json:"-"`
}
