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
	api := app.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.Post("sign_up", h.signUp)
			auth.Post("sign_in", h.signIn)
		}

		v1 := api.Group("v1")
		{
			v1.Use(middleware.JWTProtected())
			channel := v1.Group("channel/")
			{
				channel.Post("", h.createChannel)
				channel.Get("", h.getChannels)
				channel.Get(":id", h.getChannel)
				channel.Put(":id", h.changeChannel)
			}
		}
	}
}
