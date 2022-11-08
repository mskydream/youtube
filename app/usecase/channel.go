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

func (u *ChannelUseCase) GetChannels() ([]model.Channel, error) {
	return u.repo.GetChannels()
}

func (u *ChannelUseCase) GetChannel(id string) (model.Channel, error) {
	return u.repo.GetChannel(id)
}

func (u *ChannelUseCase) UpdateChannel(userId int, channel model.Channel) error {
	return u.repo.UpdateChannel(userId, channel)
}
