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

func (m *MovieController) create(c *fiber.Ctx) error {
	// parse the request body
	var req MovieDB
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// create the movie
	id, err := m.storage.createMovie(req)
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

func (m *MovieController) getAll(c *fiber.Ctx) error {
	// get all movies
	movies, err := m.storage.getAllMovies()
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

func (m *MovieController) search(c *fiber.Ctx) error {
	payload := 3
	title := c.Params("title")

	movies, err := m.storage.searchMovies(title)
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
				"error":   "Failed to search movies. Database error",
				"results": nil,
			})
		}

		for i, movie := range res.Results {
			if i < payload {
				movies, err = m.storage.searchMovies(movie.Title)
				if err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"error":   "Failed to search movies. Database error",
						"results": nil,
					})
				}
				if len(movies) == 0 || movie.Title != movies[0].Title {
					// TODO: search by ID then create document
					// convert to correct struct
					var genres []Genre
					for _, g := range movie.Genres {
						fmt.Println(g.Id)
						genres = append(genres, Genre{Name: "m", Id: g.Id})
					}

					req := MovieDB{
						Title:         movie.Title,
						TMDB_ID:       movie.Id,
						IMDB_ID:       movie.Imdb_id,
						Overview:      movie.Overview,
						Genre:         genres,
						Release_date:  movie.Release_date,
						Runtime:       movie.Runtime,
						Poster_path:   movie.Poster_path,
						Backdrop_path: movie.Backdrop_path,
					}
					fmt.Println(req)

					// create the movie
					_, err := m.storage.createMovie(req)
					if err != nil {
						fmt.Println(err)
						return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
							"message": "Failed to search movies. Database error",
							"results": nil,
						})
					}

				}

			}
		}
		movies, err = m.storage.searchMovies(title)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   "Failed to search movies. Database error",
				"results": nil,
			})
		}

	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   nil,
		"results": movies,
	})
}
