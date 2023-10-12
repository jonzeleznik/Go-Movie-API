package movies

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jonzeleznik/Go-Movie-API/pkg/requests"
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
	payload := 3
	title := c.Params("title")

	movies, err := t.storage.searchMovies(title)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to search movies",
			"results": nil,
		})
	}

	if len(movies) < payload {
		res, err := requests.NewTmdbRequests().GetTmdbMovieTitle(title)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   "Failed to search movies",
				"results": nil,
			})
		}

		for i, movie := range res.Results {
			if i < 3 {
				movies, err = t.storage.searchMovies(movie.Title)
				if err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"error":   "Failed to search movies",
						"results": nil,
					})
				}

				if movie.Title != movies[0].Title {
					// TODO: create new document in MongoDB
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"error":   "ne dela",
						"results": nil,
					})
				}

			}
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   nil,
		"results": movies,
	})
}
