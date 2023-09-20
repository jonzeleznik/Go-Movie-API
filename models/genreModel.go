package models

type Genre struct {
	Id   uint `gorm:"primaryKey"`
	Name string
}
