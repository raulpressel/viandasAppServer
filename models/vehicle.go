package models

import (
	"gorm.io/gorm"
)

type Vehicle struct {
	gorm.Model
	ID     int `gorm:"primary_key"`
	Brand  string
	Models string
	Patent string
	Year   int
}
