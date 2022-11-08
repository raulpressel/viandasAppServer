package models

import "gorm.io/gorm"

type Direction struct {
	gorm.Model
	ID          int `gorm:"primary_key"`
	Street      string
	Number      string
	Dpto        string
	Piso        string
	Observation string
	CityID      int
	City        City `gorm:"foreignKey:CityID"`
}
