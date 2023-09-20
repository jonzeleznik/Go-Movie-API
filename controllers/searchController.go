package controllers

import (
	"e/initializers"
	"e/models"
	"e/requests"
	"fmt"
	"reflect"

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

		for i, movie := range tmdbMovies.Results {
			if i < 3 {
				// -> top payload check if in DB
				initializers.DB.Where("TMDB_ID = ?", (movie.Id)).Find(&movies)
				if len(movies) == 0 {
					id := movie.Id
					details := requests.GetTmdbID(id)

					// CANT FIGURE OUT THIS PART :(
					// Genre is always null...
					genres := []models.Genre{}
					aGenres := []models.Genre{}
					initializers.DB.Find(&aGenres)

					for _, mGenres := range details.Genres {
						for _, availGenres := range aGenres {
							if reflect.DeepEqual(mGenres.Name, availGenres.Name) {
								g := models.Genre{Id: availGenres.Id, Name: availGenres.Name}
								genres = append(genres, g)
							}
						}
					}

					fmt.Println(genres)
					// -> if not write them to DB
					movie := models.Movie{
						Title:         details.Original_title,
						TMDB_ID:       id,
						IMDB_ID:       details.Imdb_id,
						Overview:      details.Overview,
						Release_date:  details.Release_date,
						Runtime:       details.Runtime,
						Poster_path:   details.Poster_path,
						Backdrop_path: details.Backdrop_path,
						Genre:         genres,
					}
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
