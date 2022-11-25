package models

import (
	"gorm.io/gorm"
)

type ClientAddress struct {
	gorm.Model
	ID        int `gorm:"primary_key"`
	ClientID  int
	Client    Client `gorm:"foreignKey:ClientID"`
	AddressID int
	Address   Address `gorm:"foreignKey:AddressID"`
}
