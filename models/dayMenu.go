package models

import (
	"time"

	"gorm.io/gorm"
)

/*usuario es el modelo de usuario de la base de mysql*/

type DayMenu struct {
	gorm.Model
	ID             int `gorm:"primary_key"`
	Date           time.Time
	FoodCategoryID int
	FoodCategory   FoodCategory `gorm:"foreignKey:FoodCategoryID"`
	TurnMenuID     int
	TurnMenu       TurnMenu `gorm:"foreignKey:TurnMenuID"`
}
