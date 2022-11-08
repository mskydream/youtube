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

func (u *ChannelUseCase) CreateChannel(userId string, channel *model.Channel) (model.Channel, error) {
	return u.repo.CreateChannel(userId, channel)
}

func (u *ChannelUseCase) GetChannels() ([]model.Channel, error) {
	return u.repo.GetChannels()
}

func (u *ChannelUseCase) GetChannel(id string) (model.Channel, error) {
	return u.repo.GetChannel(id)
}

func (u *ChannelUseCase) UpdateChannel(userId string, channel model.Channel) error {
	return u.repo.UpdateChannel(userId, channel)
}

func (u *ChannelUseCase) DeleteChannel(id string) error {
	return u.repo.DeleteChannel(id)
}
