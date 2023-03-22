package models

import (
	"gorm.io/gorm"
)

type ClientNotes struct {
	gorm.Model
	ID       int `gorm:"primary_key"`
	Note     string
	ClientID int
	Client   Client `gorm:"foreignKey:ClientID"`
}
