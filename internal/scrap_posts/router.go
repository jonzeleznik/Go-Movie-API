package scrapposts

import "github.com/gofiber/fiber/v2"

func AddScrapPostsRoutes(app *fiber.App) {
	rcsraping := app.Group("/scrapposts")

	// add middlewares here

	// add routes here
	rcsraping.Get("/", checkPosts)
	rcsraping.Get("/gpt/", postsWithGPT)
}
