package usecase

import (
	"github.com/mskydream/youtube/app/repository"
	"github.com/mskydream/youtube/model"
)

type ChannelSubscriberUseCase struct {
	repo repository.ChannelSubscriber
}

func NewChannelSubscriberUseCase(repo repository.ChannelSubscriber) *ChannelSubscriberUseCase {
	return &ChannelSubscriberUseCase{
		repo: repo,
	}
}

func (u *ChannelSubscriberUseCase) CreateChannelSubscriber(userId string, subscriber *model.ChannelSubscriber) error {
	return u.repo.CreateChannelSubscriber(userId, subscriber)
}

func (u *ChannelSubscriberUseCase) GetChannelSubscribers(userId string) ([]model.ChannelSubscriber, error) {
	return u.repo.GetChannelSubscribers(userId)
}
func (u *ChannelSubscriberUseCase) DeleteChannelSubscriber(userId string, channelId string) error {
	return u.repo.DeleteChannelSubscriber(userId, channelId)
}
