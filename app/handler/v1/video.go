package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mskydream/youtube/app/handler/v1/middleware"
	"github.com/mskydream/youtube/model"
	"github.com/mskydream/youtube/pkg"
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

	channels, err := h.usecase.GetMemberChannels(claims.UserId)
	if err != nil {
		return ctx.Status(500).JSON(
			model.ErrorResponse{
				IsSuccess: false,
				Message:   "server error",
			},
		)
	}

	if !pkg.CheckOwnerChannel(video.ChannelId, channels) {
		return ctx.Status(400).JSON(
			model.ErrorResponse{
				IsSuccess: false,
				Message:   "incorrect input",
			},
		)
	}

	err = h.usecase.CreateVideo(&video)
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

func (h *Handler) GetVideos(ctx *fiber.Ctx) error {
	videos, err := h.usecase.GetVideos()

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
			Message:   "all channels",
			Data:      videos,
		},
	)
}
