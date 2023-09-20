package controllers

import (
	"e/initializers"
	"e/models"
	"e/requests"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToDB()
}

func Search(c *gin.Context) {
	payload := 3
	// Read title from url
	title := c.Param("title")
	// Check if already in local DB
	var movies []models.Movie
	initializers.DB.Where("title LIKE ?", ("%" + title + "%")).Find(&movies)
	if len(movies) >= payload {
		// If in DB >= payload hits -> return those
		c.JSON(200, gin.H{
			"len":    len(movies),
			"movies": movies,
		})
	} else {
		// if in DB < payload hits
		// -> get from TMDB
		tmdbMovies := requests.GetTmdbMovieTitle(title)
		var movies []models.Movie
		fmt.Println(tmdbMovies.Total_results)

		for i, movie := range tmdbMovies.Results {
			if i < 3 {
				// -> top payload check if in DB
				initializers.DB.Where("TMDB_ID = ?", (movie.Id)).Find(&movies)
				fmt.Println(movies)
				if len(movies) == 0 {
					// -> if not write them to DB
					movie := models.Movie{Title: movie.Original_title, TMDB_ID: movie.Id}
					initializers.DB.Create(&movie)
				}
			}
		}

		initializers.DB.Where("title LIKE ?", ("%" + title + "%")).Find(&movies)
		c.JSON(200, gin.H{
			"movies": movies,
		})
	}
}
