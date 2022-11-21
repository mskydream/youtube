package usecase

import (
	"github.com/mskydream/youtube/app/repository"
	"github.com/mskydream/youtube/model"
)

type VideoUseCase struct {
	repo repository.Video
}

func NewVideoUseCase(repo repository.Video) *VideoUseCase {
	return &VideoUseCase{
		repo: repo,
	}
}

func (u *VideoUseCase) CreateVideo(video *model.Video) error {
	return u.repo.CreateVideo(video)
}
func (u *VideoUseCase) GetVideos() ([]model.Video, error) {
	return u.repo.GetVideos()
}
func (u *VideoUseCase) GetVideo(id string) (model.Video, error) {
	return u.repo.GetVideo(id)
}
func (u *VideoUseCase) UpdateVideo(channelId string, video model.Video) error {
	return u.repo.UpdateVideo(channelId, video)
}
func (u *VideoUseCase) DeleteVideo(id string) error {
	return u.repo.DeleteVideo(id)
}
