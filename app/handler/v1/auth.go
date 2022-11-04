package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mskydream/youtube/model"
)

func (h *Handler) signUp(ctx *fiber.Ctx) error {
	var userProfile model.UserProfile

	if err := ctx.BodyParser(&userProfile); err != nil {
		return ctx.Status(400).JSON(
			model.ErrorResponse{
				IsSuccess: false,
				Message:   "incorrect input",
			},
		)
	}

	res, err := h.usecase.SignUp(ctx, &userProfile)
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
