package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/mskydream/youtube/model"
)

type ChannelPostgres struct {
	db *sqlx.DB
}

func NewChannelPostgres(db *sqlx.DB) *ChannelPostgres {
	return &ChannelPostgres{
		db: db,
	}
}

func (r *ChannelPostgres) CreateChannel(userId string, channel *model.Channel) (res model.Channel, err error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return
	}

	defer tx.Rollback()

	query := `INSERT INTO youtube_channel (youtube_account_id, channel_name,created_at)
				VALUES ($1, $2, NOW()) RETURNING id, youtube_account_id, channel_name, created_at`

	err = r.db.QueryRow(query, userId, channel.ChannelName).Scan(&res.Id, &res.YoutubeAccountId, &res.ChannelName, &res.CreatedAt)
	if err != nil {
		return
	}
	return
}

func (r *ChannelPostgres) GetChannels() (channels []model.Channel, err error) {
	return channels, r.db.Select(&channels, `SELECT id, youtube_account_id, channel_name, created_at FROM youtube_channel`)
}

func (r *ChannelPostgres) GetChannel(id string) (channel model.Channel, err error) {
	return channel, r.db.Get(&channel, `SELECT id, youtube_account_id, channel_name, created_at FROM youtube_channel WHERE id = $1`, id)
}

func (r *ChannelPostgres) UpdateChannel(userId string, channel model.Channel) error {
	return nil
}

func (r *ChannelPostgres) DeleteChannel(id string) (err error) {
	_, err = r.db.Exec(`DELETE FROM youtube_channel WHERE id = $1`, id)
	return
}

func (r *ChannelPostgres) GetMemberChannels(id string) (channels []model.Channel, err error) {
	return channels, r.db.Select(&channels, `SELECT id, youtube_account_id, channel_name, created_at FROM youtube_channel WHERE  youtube_account_id = $1`, id)
}
