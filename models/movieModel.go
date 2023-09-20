package models

import (
	"gorm.io/gorm"
)

// type Genre struct {
// 	Id   int
// 	Name string
// }

type Movie struct {
	gorm.Model
	Title         string
	TMDB_ID       int
	IMDB_ID       string
	Overview      string
	Genres        []Genre `gorm:"foreignKey:Id"`
	Release_date  string
	Rntime        int
	Poster_path   string
	Backdrop_path string
}
