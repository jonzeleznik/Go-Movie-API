package watchlist

import "github.com/gofiber/fiber/v2"

func AddWatchListRoutes(app *fiber.App, controller *WatchListController) {
	rwatchlist := app.Group("/watchlist")

	// add middlewares here

	// add routes here
	rwatchlist.Post("/", controller.create)
	rwatchlist.Get("/", controller.getAll)
}
