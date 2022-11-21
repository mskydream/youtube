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

func (r *VideoPostgres) CreateVideo(video *model.Video) error {
	_, err := r.db.Exec(`INSERT INTO video(youtube_channel_id, video_name, created_at) VALUES ($1, $2, NOW())`, video.ChannelId, video.VideoName)
	return err
}
func (r *VideoPostgres) GetVideos() (videos []model.Video, err error) {
	return videos, r.db.Select(&videos, `SELECT id, youtube_channel_id, video_name, created_at FROM video`)
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
