package models

import (
	"gorm.io/gorm"
)

/*usuario es el modelo de usuario de la base de mysql*/

type FoodCategory struct {
	gorm.Model
	ID         int `gorm:"primary_key"`
	FoodID     int
	Food       Food `gorm:"foreignKey:FoodID"`
	CategoryID int
	Category   Category `gorm:"foreignKey:CategoryID"`
}
