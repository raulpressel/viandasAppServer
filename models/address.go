package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	ID          int `gorm:"primary_key"`
	Street      string
	Number      string
	Departament string
	Floor       string
	Observation string
	IDZone      int
	Lat         string
	Lng         string
	CityID      int
	City        City `gorm:"foreignKey:CityID"`
	Active      bool
	Favourite   bool
}
