package scrapposts

import "github.com/gofiber/fiber/v2"

func AddScrapPostsRoutes(app *fiber.App) {
	rmovies := app.Group("/scrapposts")

	// add middlewares here

	// add routes here
	rmovies.Get("/", checkPosts)
}
