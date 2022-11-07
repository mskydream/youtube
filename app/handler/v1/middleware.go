package v1

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mskydream/youtube/model"
)

var (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	// userCtx model.UserProfile
)

func (h *Handler) userInfo(ctx *fiber.Ctx) error {
	header := ctx.GetRespHeader(authorizationHeader)
	if header == "" {
		ctx.Status(401)
		return errors.New("not auth")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		ctx.Status(401).JSON(model.ErrorResponse{IsSuccess: false, Message: "Необходима авторизация"})
		return errors.New("not auth")
	}

	if len(headerParts[1]) == 0 {
		ctx.Status(401).JSON(model.ErrorResponse{IsSuccess: false, Message: "Необходима авторизация"})
		return errors.New("not auth")
	}

	user, err := h.usecase.ParseToken(headerParts[1])
	if err != nil {
		ctx.Status(401).JSON(model.ErrorResponse{IsSuccess: false, Message: "Необходима авторизация"})
		return errors.New("not auth")
	}

	ctx.Set(userCtx, user)
	return nil
}

func getUserId(ctx *fiber.Ctx) (model.UserProfile, error) {
	// id, ok := ctx.Get(userCtx)
	id, ok := ctx.Get(userCtx, model.UserProfile)
	if !ok {
		return model.UserProfile{}, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return model.UserProfile{}, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
