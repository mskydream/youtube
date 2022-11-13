package model

import "time"

type Channel struct {
	ID          int       `db:"id" json:"id"`
	UserID      string    `db:"user_id" json:"user_id"`
	ChannelName string    `db:"channel_name" json:"channel_name"`
	CreatedAt   time.Time `db:"created_at" json:"-"`
}
