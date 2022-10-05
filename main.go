package main

import (
	"github.com/Ankan002/tiny-url-api/config"
	"github.com/Ankan002/tiny-url-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
)

func main() {
	if os.Getenv("GO_ENV") != "production" {
		config.LoadEnv()
	}
	app := fiber.New()

	config.ConnectToDB()
	app.Use(cors.New())
	app.Use(logger.New())

	router := app.Group("/api")

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(fiber.Map{
			"success": true,
			"message": "Welcome to Tiny URL API",
		})
	})

	routes.ShortUrlRouter(router)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
