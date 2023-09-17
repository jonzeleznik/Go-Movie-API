package main

import (
	"e/initializers"
	"e/models"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Movie{})
	initializers.DB.AutoMigrate(&models.WatchListMovie{})
}
