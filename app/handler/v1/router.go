package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mskydream/youtube/app/handler/v1/middleware"
	"github.com/mskydream/youtube/app/usecase"
)

type Handler struct {
	usecase *usecase.UseCase
	// bot     *tgbotapi.BotAPI
}

func NewHandler(usecase *usecase.UseCase) *Handler {
	return &Handler{
		usecase: usecase,
		// bot:     bot,
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
			channel.Delete(":id", middleware.JWTProtected(), h.DeleteChannel)
		}
	}
}
