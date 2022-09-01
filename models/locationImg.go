package models

import "gorm.io/gorm"

type LocationImg struct {
	gorm.Model
	ID       int `gorm:"primary_key"`
	Location string
}
