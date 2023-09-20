package controllers

import (
	"e/initializers"
	"e/models"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToDB()
}

func MovieCreate(c *gin.Context) {
	// get data from body
	var body models.Movie
	c.Bind(&body)

	// convert to int
	tmdb_id := body.TMDB_ID

	// store movie
	movie := models.Movie{Title: body.Title, TMDB_ID: tmdb_id, Genre: body.Genre}
	result := initializers.DB.Create(&movie)

	// resp
	if result.Error != nil {
		c.JSON(200, gin.H{
			"error": result.Error,
		})
	} else {
		c.JSON(200, gin.H{
			"movie": movie,
		})
	}
}

func MovieGetByTitle(c *gin.Context) {
	title := c.Param("title")

	var movies []models.Movie
	initializers.DB.Where("title LIKE ?", ("%" + title + "%")).Find(&movies)

	c.JSON(200, gin.H{
		"movies": movies,
	})
}

func MovieGetById(c *gin.Context) {
	id := c.Param("id")

	var movie models.Movie
	initializers.DB.First(&movie, id)

	c.JSON(200, gin.H{
		"movie": movie,
	})
}

func MovieDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Movie{}, id)

	c.JSON(200, gin.H{
		"deleted": true,
	})
}
