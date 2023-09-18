package main

import (
	"e/controllers"
	"e/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	r := gin.Default()
	r.POST("/movies", controllers.MovieCreate)
	r.GET("/movies/title/:title", controllers.MovieGetByTitle)
	r.GET("/movies/id/:id", controllers.MovieGetById)

	r.POST("/watchList", controllers.AddToWatchList)
	r.GET("/watchList", controllers.GetWatchList)

	r.GET(("/search/:title"), controllers.Search)
	// r.Run((":" + initializers.EnvVars["PORT"]))
	r.Run()
}
