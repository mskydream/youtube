package model

import "time"

type Channel struct {
	ID          int       `db:"id" json:"-"`
	UserID      int       `db:"user_id" json:"user_id"`
	ChannelName string    `db:"channel_name" json:"channel_name"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
