package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mskydream/youtube/config"
	"github.com/mskydream/youtube/model"
	"github.com/mskydream/youtube/pkg"
)

func (h *Handler) signUp(ctx *fiber.Ctx) error {
	var userProfile model.UserProfile

	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	if err := ctx.BodyParser(&userProfile); err != nil {
		return ctx.Status(400).JSON(
			model.ErrorResponse{
				IsSuccess: false,
				Message:   "incorrect input",
			},
		)
	}

	userProfile.Password = pkg.GeneratePasswordHash(userProfile.Password, cfg.Salt)

	res, err := h.usecase.SignUp(&userProfile)
	if err != nil {
		h.usecase.TelegramBot.SendMessageLog(cfg.Telegram.ChatId, "Sign up server error")
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

func (h *Handler) signIn(ctx *fiber.Ctx) error {
	var input model.SignIn

	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(400).JSON(
			model.ErrorResponse{
				IsSuccess: false,
				Message:   "incorrect input",
			},
		)
	}

	input.Password = pkg.GeneratePasswordHash(input.Password, cfg.Salt)

	token, err := h.usecase.SignIn(&input)
	if err != nil {
		h.usecase.TelegramBot.SendMessageLog(cfg.Telegram.ChatId, "Sign in server error")
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
			Message:   "sign in success",
			Data:      token,
		},
	)
}
