package models

import "gorm.io/gorm"

type ProductCategory struct {
	gorm.Model
	Title       string
	Description string
	Active      bool
}
