package movies

import "github.com/gofiber/fiber/v2"

func AddMovieRoutes(app *fiber.App, controller *MovieController) {
	rmovies := app.Group("/movies")

	// add middlewares here

	// add routes here
	rmovies.Post("/", controller.create)
	rmovies.Get("/", controller.getAll)
	rmovies.Get("/search/:title", controller.search)
}
