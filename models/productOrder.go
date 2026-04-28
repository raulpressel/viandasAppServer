package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductOrderItem struct {
	gorm.Model
	ProductOrderID       uint
	ProductTitle         string
	ProductCategoryTitle string
	Cant                 int
}

type ProductOrder struct {
	gorm.Model
	ClientName     string
	ClientLastName string
	Date           time.Time
	Status         string
	Products       []ProductOrderItem `gorm:"foreignKey:ProductOrderID"`
}
