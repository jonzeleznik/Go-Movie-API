package models

import "gorm.io/gorm"

type WatchListMovie struct {
	gorm.Model
	Movie_ID int16
	Rate     int16 `gorm:"default:0"`
	Watched  bool  `gorm:"default:false"`
}
