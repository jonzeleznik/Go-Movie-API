package movies

import (
	"strconv"

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
		"error":        nil,
		"result_count": len(movies),
		"results":      movies,
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
		req := requests.NewTmdbRequests()
		res, err := req.GetTmdbMovieTitle(title)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   "Failed to search movies. Database error",
				"results": nil,
			})
		}

		for i, movie := range res.Results {
			if i < payload {
				id := strconv.Itoa(movie.Id)

				// get details
				details, err := req.GetTmdbMovieId(id)
				if err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"error":   "Failed to search movies. Database error",
						"results": nil,
					})
				}

				movies, err = m.storage.searchMovies(details.Title)
				if err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"error":   "Failed to search movies. Database error",
						"results": nil,
					})
				}
				if len(movies) == 0 || details.Title != movies[0].Title {
					// TODO: search by ID then create document
					// convert to correct struct
					var genres []Genre
					for _, g := range details.Genres {
						genres = append(genres, Genre(g))
					}

					doc := MovieDB{
						Title:         details.Title,
						TMDB_ID:       details.Id,
						IMDB_ID:       details.Imdb_id,
						Overview:      details.Overview,
						Genre:         genres,
						Release_date:  details.Release_date,
						Runtime:       details.Runtime,
						Poster_path:   details.Poster_path,
						Backdrop_path: details.Backdrop_path,
					}

					// create the movie
					_, err := m.storage.createMovie(doc)
					if err != nil {
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
		"error":        nil,
		"result_count": len(movies),
		"results":      movies,
	})
}
