package postgres

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mskydream/youtube/model"
)

type AuthPostgres struct {
	pool *pgxpool.Pool
}

func NewAuthPostgres(pool *pgxpool.Pool) *AuthPostgres {
	return &AuthPostgres{
		pool: pool,
	}
}

func (a *AuthPostgres) SignUp(ctx *fiber.Ctx, input *model.UserProfile) (res model.UserProfile, err error) {
	query := fmt.Sprintln(`
	INSERT INTO user_profile 
	(fist_name, last_name, gender, email, pass, created_at) values ($1, $2, $3, $4, $5, now()) 
	RETURNING id, first_name, last_name, gender, email, pass, created_at`)
	err = a.pool.QueryRow(context.Background(), query, input.FistName, input.LastName, input.Gender, input.Password).Scan(&res)
	fmt.Printf("res: %v\n", res)
	return res, nil
}
