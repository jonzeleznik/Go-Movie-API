package watchlist

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type WatchListController struct {
	storage *WatchListStorage
}

func NewWatchListController(storage *WatchListStorage) *WatchListController {
	return &WatchListController{
		storage: storage,
	}
}

func (t *WatchListController) create(c *fiber.Ctx) error {
	// parse the request body
	var req WatchListDB
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// create the movie
	id, err := t.storage.createWatchList(req)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create movie",
			"results": nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   nil,
		"results": id,
	})
}
