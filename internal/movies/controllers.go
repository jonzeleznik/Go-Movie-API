package movies

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MovieController struct {
	storage *MovieStorage
}

func NewMovieController(storage *MovieStorage) *MovieController {
	return &MovieController{
		storage: storage,
	}
}

func (t *MovieController) create(c *fiber.Ctx) error {
	// parse the request body
	var req MovieDB
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// create the movie
	id, err := t.storage.createMovie(req)
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

func (t *MovieController) getAll(c *fiber.Ctx) error {
	// get all movies
	movies, err := t.storage.getAllMovies()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to get all movies",
			"results": nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   nil,
		"results": movies,
	})
}

func (t *MovieController) search(c *fiber.Ctx) error {
	title := c.Params("title")

	movies, err := t.storage.searchMovies(title)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to get all movies",
			"results": nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   nil,
		"results": movies,
	})
}
