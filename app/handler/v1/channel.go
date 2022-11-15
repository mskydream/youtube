package v1

import (
	"github.com/gofiber/fiber/v2"

	"github.com/mskydream/youtube/app/handler/v1/middleware"
	"github.com/mskydream/youtube/config"
	"github.com/mskydream/youtube/model"
)

func (h *Handler) createChannel(ctx *fiber.Ctx) error {
	var channel model.Channel

	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	claims, err := middleware.ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(401).JSON(
			model.ErrorResponse{
				IsSuccess: false,
				Message:   "unauthorized",
			},
		)
	}

	if err := ctx.BodyParser(&channel); err != nil {
		return ctx.Status(400).JSON(
			model.ErrorResponse{
				IsSuccess: false,
				Message:   "incorrect input",
			},
		)
	}

	res, err := h.usecase.CreateChannel(claims.UserId, &channel)
	if err != nil {
		if err.Error() == `ERROR: duplicate key value violates unique constraint "youtube_channel_channel_name_key" (SQLSTATE 23505)` {
			return ctx.Status(400).JSON(
				model.ErrorResponse{
					IsSuccess: false,
					Message:   "channel exist",
				},
			)
		}

		h.usecase.TelegramBot.SendMessageLog(cfg.Telegram.ChatId, "Create channel server error user id: "+claims.UserId+"\nChannel name: "+channel.ChannelName)
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

func (h *Handler) getChannels(ctx *fiber.Ctx) error {
	channels, err := h.usecase.GetChannels()
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
			Data:      channels,
		},
	)
}

func (h *Handler) getChannel(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	channel, err := h.usecase.GetChannel(id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return ctx.Status(400).JSON(
				model.ErrorResponse{
					IsSuccess: false,
					Message:   "no rows in result set",
				},
			)
		}

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
			Message:   "channel",
			Data:      channel,
		},
	)
}

func (h *Handler) changeChannel(ctx *fiber.Ctx) error {
	channelId := ctx.Params("id")
	var channel model.Channel

	claims, err := middleware.ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(401).JSON(
			model.ErrorResponse{
				IsSuccess: false,
				Message:   "unauthorized",
			},
		)
	}

	if err := ctx.BodyParser(&channel); err != nil {
		return ctx.Status(400).JSON(
			model.ErrorResponse{
				IsSuccess: false,
				Message:   "incorrect input",
			},
		)
	}

	checkResult, err := h.usecase.GetChannel(channelId)
	if err != nil {
		return ctx.Status(500).JSON(
			model.ErrorResponse{
				IsSuccess: false,
				Message:   "server error",
			},
		)
	}

	if checkResult.YoutubeAccountId != claims.UserId {
		return ctx.Status(400).JSON(
			model.ErrorResponse{
				IsSuccess: false,
				Message:   "not enough rights",
			},
		)
	}

	err = h.usecase.UpdateChannel(claims.UserId, channel)
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
			Message:   "success",
		},
	)
}

func (h *Handler) DeleteChannel(ctx *fiber.Ctx) error {
	channelId := ctx.Params("id")

	claims, err := middleware.ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(401).JSON(
			model.ErrorResponse{
				IsSuccess: false,
				Message:   "unauthorized",
			},
		)
	}

	checkResult, err := h.usecase.GetChannel(channelId)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return ctx.Status(400).JSON(
				model.ErrorResponse{
					IsSuccess: false,
					Message:   "no rows in result set",
				},
			)
		}

		return ctx.Status(500).JSON(
			model.ErrorResponse{
				IsSuccess: false,
				Message:   "server error",
			},
		)
	}

	if checkResult.YoutubeAccountId != claims.UserId {
		return ctx.Status(400).JSON(
			model.ErrorResponse{
				IsSuccess: false,
				Message:   "not enough rights",
			},
		)
	}

	err = h.usecase.DeleteChannel(channelId)
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
			Message:   "deleted channel",
		},
	)
}
