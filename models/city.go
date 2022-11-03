package models

import "gorm.io/gorm"

type City struct {
	gorm.Model
	ID          int `gorm:"primary_key"`
	Description string
	CP          string
}
