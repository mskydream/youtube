package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/mskydream/youtube/model"
)

type VideoPostgres struct {
	db *sqlx.DB
}

func NewVideoPostgres(db *sqlx.DB) *VideoPostgres {
	return &VideoPostgres{
		db: db,
	}
}

// func (r *ChannelPostgres) CreateVideo(userId string, channel *model.Channel) (res model.Channel, err error) {
// 	tx, err := r.db.Beginx()
// 	if err != nil {
// 		return
// 	}

// 	defer tx.Rollback()

// 	query := `INSERT INTO youtube_channel (youtube_account_id, channel_name,created_at)
// 				VALUES ($1, $2, NOW()) RETURNING id, youtube_account_id, channel_name, created_at`

//		err = r.db.QueryRow(query, userId, channel.ChannelName).Scan(&res.Id, &res.YoutubeAccountId, &res.ChannelName, &res.CreatedAt)
//		if err != nil {
//			return
//		}
//		return
//	}
func (r *VideoPostgres) CreateVideo(channelId string, video *model.Video) error {
	return nil
}
func (r *VideoPostgres) GetVideos() ([]model.Video, error) {
	return []model.Video{}, nil
}
func (r *VideoPostgres) GetVideo(id string) (model.Video, error) {
	return model.Video{}, nil
}
func (r *VideoPostgres) UpdateVideo(channelId string, video model.Video) error {
	return nil
}
func (r *VideoPostgres) DeleteVideo(id string) error {
	return nil
}
