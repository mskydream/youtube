package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mskydream/youtube/app/handler/v1/middleware"
	"github.com/mskydream/youtube/app/usecase"
)

type Handler struct {
	usecase *usecase.UseCase
}

func NewHandler(usecase *usecase.UseCase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) InitRouterV1(app *fiber.App) {
	api := app.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.Post("sign_up", h.signUp)
			auth.Post("sign_in", h.signIn)
		}

		channel := api.Group("channel/")
		{
			channel.Post("", middleware.JWTProtected(), h.createChannel)
			channel.Get("", h.getChannels)
			channel.Get(":id", h.getChannel)
			channel.Put(":id", middleware.JWTProtected(), h.changeChannel)
			channel.Delete(":id", middleware.JWTProtected(), h.deleteChannel)

			subscriber := channel.Group("subscriber/")
			{
				subscriber.Post("", middleware.JWTProtected(), h.createChannelSubscriber)
				subscriber.Get("list", middleware.JWTProtected(), h.allChannelSubscribers)
				subscriber.Delete(":channel_id", middleware.JWTProtected(), h.unsubscribeChannel)
			}
			video := channel.Group("video/")
			{
				video.Post("", middleware.JWTProtected(), h.createVideo)
			}
		}
	}
}
