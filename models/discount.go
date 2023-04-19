package models

import (
	"gorm.io/gorm"
)

type Discount struct {
	gorm.Model
	ID          int `gorm:"primary_key"`
	Description string
	Cant        int
	Percentage  float32
	Active      bool
}
