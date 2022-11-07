package usecase

import (
	"github.com/mskydream/youtube/app/repository"
	"github.com/mskydream/youtube/model"
)

type ChannelUseCase struct {
	repo repository.Channel
}

func NewChannelUseCase(repo repository.Channel) *ChannelUseCase {
	return &ChannelUseCase{
		repo: repo,
	}
}

func (u *ChannelUseCase) CreateChannel(channel *model.Channel, userId int) (model.Channel, error) {
	return u.repo.CreateChannel(channel, userId)
}
