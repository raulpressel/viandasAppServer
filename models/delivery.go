package models

import (
	"time"

	"gorm.io/gorm"
)

type Delivery struct {
	gorm.Model
	ID                 int `gorm:"primary_key"`
	OrderID            int
	Order              Order     `gorm:"foreignKey:OrderID"`
	DeliveryDate       time.Time //
	DeliveryPrice      float32   //
	DeliveryDriverID   *uint
	DeliveryDriver     DeliveryDriver `gorm:"foreignKey:DeliveryDriverID"`
	AddressID          int            //
	Address            Address        `gorm:"foreignKey:AddressID"`
	DeliveryMenuPrice  float32
	DeliveryMenuAmount int
	PercentageDiscount float32
	Status             bool //
}
