package url

import (
	"github.com/Ankan002/tiny-url-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func Redirect(c *fiber.Ctx) error {
	requestedUrl := c.Query("url")

	if requestedUrl == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Please provide us with a valid URL",
		})
	}

	fetchedUrlData := models.MiniUrl{}

	urlFetchError := mgm.Coll(&models.MiniUrl{}).First(bson.M{"short_id": requestedUrl}, &fetchedUrlData)

	if urlFetchError != nil && urlFetchError.Error() == "mongo: no documents in result" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No url found with the given url id",
		})
	}

	if urlFetchError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": urlFetchError.Error(),
		})
	}

	return c.Redirect(fetchedUrlData.Url, 301)
}
