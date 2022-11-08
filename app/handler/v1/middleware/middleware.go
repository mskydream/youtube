package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	jwtMiddleware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt"
	"github.com/mskydream/youtube/config"
)

type TokenMetadata struct {
	UserId int
}

func JWTProtected() func(*fiber.Ctx) error {
	cfg, _ := config.LoadConfig()

	config := jwtMiddleware.Config{
		SigningKey:   []byte(cfg.Token.Key),
		ContextKey:   "jwt", // used in private routes
		ErrorHandler: jwtError,
	}

	return jwtMiddleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": true,
		"msg":   err.Error(),
	})
}

func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		id := int(claims["id"].(float64))

		return &TokenMetadata{
			UserId: id,
		}, nil
	}

	return nil, err
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	onlyToken := strings.Split(bearToken, " ")

	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)

	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	cfg, _ := config.LoadConfig()
	return []byte(cfg.Token.Key), nil
}
