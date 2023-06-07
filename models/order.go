package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID            int `gorm:"primary_key"`
	Observation   string
	OrderDate     time.Time
	Total         float32
	StatusOrderID int
	StatusOrder   StatusOrder
	Paid          bool
	ClientID      int
	Client        Client `gorm:"foreignKey:ClientID"`
}
