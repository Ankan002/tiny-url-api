package routes

import (
	"github.com/Ankan002/tiny-url-api/controller/url"
	"github.com/gofiber/fiber/v2"
)

func ShortUrlRouter(router fiber.Router) {
	router.Post("/shorten-url", url.ShortenUrl)
}
