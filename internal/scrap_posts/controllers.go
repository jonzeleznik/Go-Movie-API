package scrapposts

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonzeleznik/Go-Movie-API/pkg/scraper"
)

func checkPosts(c *fiber.Ctx) error {
	posts := scraper.HollywoodReporter()

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": nil,
		"results": posts,
	})
}
