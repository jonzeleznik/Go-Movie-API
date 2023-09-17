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
	r.POST("/postMovies", controllers.MovieCreate)
	r.GET("/getMovies/title/:title", controllers.MovieGetByTitle)
	r.GET("/getMovies/id/:id", controllers.MovieGetById)

	r.POST("/postWatchList", controllers.AddToWatchList)
	r.GET("/getWatchList", controllers.GetWatchList)
	// r.Run((":" + initializers.EnvVars["PORT"]))
	r.Run()
}
