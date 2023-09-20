package models

import (
	"gorm.io/gorm"
)

type Genre struct {
	Id   int
	Name string
}

type Movie struct {
	gorm.Model
	Title         string
	TMDB_ID       int
	IMDB_ID       string
	Overview      string
	Genre         []Genre `gorm:"foreignKey:Id"`
	Release_date  string
	Runtime       int
	Poster_path   string
	Backdrop_path string
}
