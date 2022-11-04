package models

import (
	"gorm.io/gorm"
)

type ClientDir struct {
	gorm.Model
	ID          int `gorm:"primary_key"`
	ClientID    int
	Client      Client `gorm:"foreignKey:ClientID"`
	DirectionID int
	Direction   Direction `gorm:"foreignKey:DirectionID"`
}
