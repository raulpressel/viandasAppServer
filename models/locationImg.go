package models

import "gorm.io/gorm"

type LocationImg struct {
	gorm.Model
	ID       uint `gorm:"primary_key"`
	Location string
}
