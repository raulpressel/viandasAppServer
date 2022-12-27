package models

import (
	"gorm.io/gorm"
)

type Tanda struct {
	gorm.Model
	ID               int `gorm:"primary_key"`
	Description      string
	HourStart        string
	HourEnd          string
	Active           bool
	DeliveryDriverID int
	DeliveryDriver   DeliveryDriver `gorm:"foreignKey:DeliveryDriverID"`
}
