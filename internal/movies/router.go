package movies

import "github.com/gin-gonic/gin"

func AddMovieRoutes(superRoute *gin.RouterGroup, controller *MovieController) {
	movies := superRoute.Group("/movies")

	// add middlewares here

	// add routes here
	{
		movies.POST("/", controller.create)
		movies.GET("/", controller.getAll)
	}
}
