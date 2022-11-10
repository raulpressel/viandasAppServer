package models

import (
	"gorm.io/gorm"
)

type ClientPathology struct {
	gorm.Model
	ID          int `gorm:"primary_key"`
	ClientID    int
	Client      Client `gorm:"foreignKey:ClientID"`
	PathologyID int
	Pathology   Pathology `gorm:"foreignKey:PathologyID"`
}
