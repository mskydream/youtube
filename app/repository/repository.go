package repository

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mskydream/youtube/app/repository/postgres"
	"github.com/mskydream/youtube/model"
)

type Auth interface {
	SignUp(ctx *fiber.Ctx, userProfile *model.UserProfile) (model.UserProfile, error)
}

type Repository struct {
	Auth
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		Auth: postgres.NewAuthPostgres(pool),
	}
}
