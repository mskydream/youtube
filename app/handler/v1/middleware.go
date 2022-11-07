package v1

// import (
// 	"strings"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/mskydream/youtube/model"
// )

// var (
// 	authorizationHeader = "Authorization"
// 	// userCtx             = "userId"
// 	userCtx model.UserProfile
// )

// func (h *Handler) userInfo(ctx *fiber.Ctx) {
// 	header := ctx.GetRespHeader(authorizationHeader)
// 	if header == "" {
// 		ctx.Status(401)
// 		return
// 	}

// 	headerParts := strings.Split(header, " ")
// 	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
// 		ctx.Status(401).JSON(model.ErrorResponse{IsSuccess: false, Message: "Необходима авторизация"})
// 		return
// 	}

// 	if len(headerParts[1]) == 0 {
// 		ctx.Status(401).JSON(model.ErrorResponse{IsSuccess: false, Message: "Необходима авторизация"})
// 		return
// 	}

// 	_, err := h.usecase.ParseToken(headerParts[1])
// 	if err != nil {
// 		ctx.Status(401).JSON(model.ErrorResponse{IsSuccess: false, Message: "Необходима авторизация"})
// 		return
// 	}

// 	// ctx.Set(userCtx, userId)
// }

// func getUserId(ctx *fiber.Ctx) (model.UserProfile, error) {
// 	// id, ok := ctx.Get(userCtx)
// 	id, ok := ctx.Get(userCtx.Email)
// 	if !ok {
// 		return model.UserProfile{}, errors.New("user id not found")
// 	}

// 	idInt, ok := id.(int)
// 	if !ok {
// 		return model.UserProfile{}, errors.New("user id is of invalid type")
// 	}

// 	return idInt, nil
// }
