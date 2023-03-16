package models

import (
	"gorm.io/gorm"
)

type TandaAddress struct {
	gorm.Model
	ID        int `gorm:"primary_key"`
	TandaID   int
	Tanda     Tanda `gorm:"foreignKey:TandaID"`
	AddressID int
	Address   Address `gorm:"foreignKey:AddressID"`
}
