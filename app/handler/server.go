package handler

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	v1 "github.com/mskydream/youtube/app/handler/v1"
	"github.com/mskydream/youtube/app/repository"
	"github.com/mskydream/youtube/app/usecase"
	"github.com/mskydream/youtube/config"
	"github.com/mskydream/youtube/db"
)

func Run(app *fiber.App, cfg config.Config) error {
	db, err := db.InitDatabase(cfg)
	if err != nil {
		return err
	}
	log.Println("Database success connected...")

	repo := repository.NewRepository(db)
	usecase := usecase.NewUseCase(repo)
	handler := v1.NewHandler(usecase)

	app.Use(logger.New())
	handler.InitRouterV1(app)

	log.Fatal(app.Listen(cfg.Port))

	err = app.Listen(cfg.Port)
	if err != nil {
		fmt.Println("Error in the port")
	}
	return err
}
