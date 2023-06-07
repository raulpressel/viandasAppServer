package models

import (
	"gorm.io/gorm"
)

type DayOrder struct {
	gorm.Model
	ID          int `gorm:"primary_key"`
	Observation string
	Amount      int
	Status      bool
	//Active      bool
	AddressID   int
	Address     Address `gorm:"foreignKey:AddressID"`
	DayMenuID   int
	DayMenu     DayMenu `gorm:"foreignKey:DayMenuID"`
	OrderID     int
	Order       Order `gorm:"foreignKey:OrderID"`
}
