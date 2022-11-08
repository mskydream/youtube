package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mskydream/youtube/app/handler"
	"github.com/mskydream/youtube/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Panic(err)
	}
	// go bot.RunTelegramBot(cfg)
	app := fiber.New()

	err = handler.Run(app, cfg)
	if err != nil {
		log.Panic(err)
	}
}
