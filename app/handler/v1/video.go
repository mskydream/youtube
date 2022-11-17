package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mskydream/youtube/app/handler/v1/middleware"
	"github.com/mskydream/youtube/model"
)

func (h *Handler) createVideo(ctx *fiber.Ctx) error {
	var video model.Video

	claims, err := middleware.ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(401).JSON(
			model.ErrorResponse{
				IsSuccess: false,
				Message:   "unauthorized",
			},
		)
	}

	if err := ctx.BodyParser(&video); err != nil {
		return ctx.Status(400).JSON(
			model.ErrorResponse{
				IsSuccess: false,
				Message:   "incorrect input",
			},
		)
	}

	err = h.usecase.CreateVideo(claims.UserId, &video)
	if err != nil {
		return ctx.Status(500).JSON(
			model.ErrorResponse{
				IsSuccess: false,
				Message:   "server error",
			},
		)
	}

	return ctx.Status(200).JSON(
		model.SuccessResponse{
			IsSuccess: true,
			Message:   "vidoe created",
		},
	)
}
