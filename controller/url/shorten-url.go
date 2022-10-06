package url

import (
	"github.com/Ankan002/tiny-url-api/helpers"
	"github.com/Ankan002/tiny-url-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"strings"
)

type RequestBody struct {
	Url              string `json:"url"`
	UrlSuggestedName string `json:"urlSuggestedName"`
}

func ShortenUrl(c *fiber.Ctx) error {
	requestBody := RequestBody{}

	bodyDecodingError := c.BodyParser(&requestBody)

	if bodyDecodingError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   bodyDecodingError.Error(),
		})
	}

	if requestBody.Url == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Please provide us with an URL.",
		})
	}

	requestBody.Url = strings.TrimSpace(requestBody.Url)

	isValidUrl := helpers.ValidateUrl(requestBody.Url)

	if !isValidUrl {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "You are not allowed to shorten the url or any subdomain url containing, " + os.Getenv("DOMAIN"),
		})
	}

	requestBody.Url = helpers.EnforceHTTPS(requestBody.Url)

	//TODO: Decide if we want to invoke check reachable URL function

	//_, hostError := http.Get(requestBody.Url)
	//
	//if hostError != nil {
	//	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	//		"success": false,
	//		"error":   "Unable to get the host of the given URL. Please provide us with a valid URL",
	//	})
	//}

	requestBody.UrlSuggestedName = strings.TrimSpace(requestBody.UrlSuggestedName)

	if requestBody.UrlSuggestedName == "" {
		requestBody.UrlSuggestedName = utils.UUIDv4()
	}

	oldFoundUrl := models.MiniUrl{}

	_ = mgm.Coll(&models.MiniUrl{}).First(bson.M{"short_id": requestBody.UrlSuggestedName}, &oldFoundUrl)

	if oldFoundUrl != (models.MiniUrl{}) {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "An URL with same short already exists...",
		})
	}

	newUrl := models.MiniUrl{
		ShortId: requestBody.UrlSuggestedName,
		Url:     requestBody.Url,
	}

	creationError := mgm.Coll(&models.MiniUrl{}).Create(&newUrl)

	if creationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   creationError.Error(),
		})
	}

	var url string

	if os.Getenv("GO_ENV") == "development" {
		url = "http://" + os.Getenv("DOMAIN") + "/api/redirect?url=" + newUrl.ShortId
	} else {
		url = "https://" + os.Getenv("DOMAIN") + "/api/redirect?url=" + newUrl.ShortId
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"url":     url,
	})
}
