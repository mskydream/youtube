package pkg

import "github.com/mskydream/youtube/model"

func CheckOwnerChannel(id string, channels []model.Channel) bool {
	for _, channel := range channels {
		if id == channel.YoutubeAccountId {
			return true
		}
	}
	return false
}
