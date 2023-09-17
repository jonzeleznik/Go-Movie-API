package controllers

import (
	"e/initializers"
	"e/models"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToDB()
}

func AddToWatchList(c *gin.Context) {
	// get data from body
	var body struct {
		Movie_ID int16
		// TMDB_ID  string
		// Genre    string
	}

	c.Bind(&body)

	watchList := models.WatchListMovie{Movie_ID: body.Movie_ID}
	result := initializers.DB.Create(&watchList)

	// resp
	if result.Error != nil {
		c.JSON(200, gin.H{
			"error": result.Error,
		})
	} else {
		c.JSON(200, gin.H{
			"movie": watchList,
		})
	}
}

func GetWatchList(c *gin.Context) {
	var watchList []models.WatchListMovie
	result := initializers.DB.Find(&watchList)

	// resp
	if result.Error != nil {
		c.JSON(200, gin.H{
			"error": result.Error,
		})
	} else {
		c.JSON(200, gin.H{
			"movie": watchList,
		})
	}
}

func RateMovie(c *gin.Context) {
	// TODO: update rate col
}

func ChangeStatus(c *gin.Context) {
	// TODO: update status col
}
