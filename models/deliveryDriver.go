package models

import (
	"time"

	"gorm.io/gorm"
)

type DeliveryDriver struct {
	gorm.Model
	ID        int `gorm:"primary_key"`
	DNI       int
	Name      string
	LastName  string
	Phone     string
	BornDate  time.Time
	Active    bool
	AddressID int
	Address   Address `gorm:"foreignKey:AddressID"`
	VehicleID int
	Vehicle   Vehicle `gorm:"foreignKey:VehicleID"`
}
