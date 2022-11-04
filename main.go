package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mskydream/youtube/app/handler"
)

func main() {
	app := fiber.New()

	err := handler.Run(app)
	if err != nil {
		panic(err)
	}
}
