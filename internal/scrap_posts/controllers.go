package scrapposts

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonzeleznik/Go-Movie-API/pkg/requests"
	"github.com/jonzeleznik/Go-Movie-API/pkg/scraper"
)

func checkPosts(c *fiber.Ctx) error {
	posts := scraper.HollywoodReporter()

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": nil,
		"results": posts,
	})
}

func postsWithGPT(c *fiber.Ctx) error {
	gpt := requests.NewGPTRequests()
	resp, err := gpt.GetChatGPT()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
			"results": nil,
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": nil,
		"results": resp,
	})
}
