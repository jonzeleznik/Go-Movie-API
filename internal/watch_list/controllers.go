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

func (m *WatchListController) create(c *fiber.Ctx) error {
	// parse the request body
	var req WatchListDB
	if err := c.BodyParser(&req); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// create the movie
	id, err := m.storage.createWatchList(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create watch list movie",
			"results": nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   nil,
		"results": id,
	})
}

func (m *WatchListController) getAll(c *fiber.Ctx) error {
	// get all movies
	movies, err := m.storage.getAllWatchList()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to get all movies",
			"results": nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":        nil,
		"result_count": len(movies),
		"results":      movies,
	})
}
