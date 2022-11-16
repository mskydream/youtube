package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mskydream/youtube/app/handler/v1/middleware"
	"github.com/mskydream/youtube/model"
)

func (h *Handler) createChannelSubscriber(ctx *fiber.Ctx) error {
	var subscriber model.ChannelSubscriber

	claims, err := middleware.ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(401).JSON(
			model.ErrorResponse{
				IsSuccess: false,
				Message:   "unauthorized",
			},
		)
	}

	if err := ctx.BodyParser(&subscriber); err != nil {
		return ctx.Status(400).JSON(
			model.ErrorResponse{
				IsSuccess: false,
				Message:   "incorrect input",
			},
		)
	}

	err = h.usecase.CreateChannelSubscriber(claims.UserId, &subscriber)
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
			Message:   "subscribe success",
		},
	)
}

func (h *Handler) allChannelSubscribers(ctx *fiber.Ctx) error {
	claims, err := middleware.ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(401).JSON(
			model.ErrorResponse{
				IsSuccess: false,
				Message:   "unauthorized",
			},
		)
	}

	res, err := h.usecase.GetChannelSubscribers(claims.UserId)
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
			Message:   "all channel subscribers",
			Data:      res,
		},
	)
}

func (h *Handler) unsubscribeChannel(ctx *fiber.Ctx) error {
	channelId := ctx.Params("channel_id")

	claims, err := middleware.ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(401).JSON(
			model.ErrorResponse{
				IsSuccess: false,
				Message:   "unauthorized",
			},
		)
	}

	err = h.usecase.DeleteChannelSubscriber(claims.UserId, channelId)
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
			Message:   "unsubscribed",
		},
	)
}
