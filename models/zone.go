package models

import (
	"gorm.io/gorm"
)

type Zone struct {
	gorm.Model
	ID          int `gorm:"primary_key"`
	Description string
	Price       float32
	Active      bool
}
