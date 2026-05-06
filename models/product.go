package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title             string
	Description       string
	Price             float32
	Available         bool
	ProductCategoryID int
	ProductCategory   ProductCategory `gorm:"foreignKey:ProductCategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
