package handler

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	v1 "github.com/mskydream/youtube/app/handler/v1"
	"github.com/mskydream/youtube/app/repository"
	"github.com/mskydream/youtube/app/usecase"
	"github.com/mskydream/youtube/bot"
	"github.com/mskydream/youtube/config"
	"github.com/mskydream/youtube/db"
)

func Run(app *fiber.App) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	db, err := db.InitDatabase(cfg)
	if err != nil {
		return err
	}

	ch := make(chan *tgbotapi.BotAPI)
	go bot.RunTelegramBot(cfg, ch)

	repo := repository.NewRepository(db)
	usecase := usecase.NewUseCase(repo, <-ch)
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
