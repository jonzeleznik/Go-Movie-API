package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title   string
	TMDB_ID int
	Genre   string
}
