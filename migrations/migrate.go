package main

import (
	"e/initializers"
	"e/models"
)

var genres = []models.Genre{
	{Name: "Action"},
	{Name: "Adventure"},
	{Name: "Animation"},
	{Name: "Comedy"},
	{Name: "Crime"},
	{Name: "Documentary"},
	{Name: "Drama"},
	{Name: "Family"},
	{Name: "Fantasy"},
	{Name: "History"},
	{Name: "Horror"},
	{Name: "Music"},
	{Name: "Mystery"},
	{Name: "Romance"},
	{Name: "Science Fiction"},
	{Name: "TV Movie"},
	{Name: "Thriller"},
	{Name: "War"},
	{Name: "Western"},
}

func init() {
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Movie{})
	initializers.DB.AutoMigrate(&models.WatchListMovie{})
	initializers.DB.AutoMigrate(&models.Genre{})
	initializers.DB.Create(genres)
}
