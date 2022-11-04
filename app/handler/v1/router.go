package v1

import (
	"github.com/gofiber/fiber/v2"
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
		}
	}
}
