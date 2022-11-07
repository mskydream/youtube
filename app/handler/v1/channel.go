package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mskydream/youtube/model"
)

func (h *Handler) CreateChannel(ctx *fiber.Ctx) error {
	var channel model.Channel

	if err := ctx.BodyParser(&channel); err != nil {
		return ctx.Status(400).JSON(
			model.ErrorResponse{
				IsSuccess: false,
				Message:   "incorrect input",
			},
		)
	}

	userId := 5

	res, err := h.usecase.CreateChannel(&channel, userId)
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
			Message:   "register success",
			Data:      res,
		},
	)
}
