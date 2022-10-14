package models

import (
	"gorm.io/gorm"
)

/*usuario es el modelo de usuario de la base de mysql*/

type Category struct {
	gorm.Model
	ID          int `gorm:"primary_key"`
	Description string
	Title       string
	Price       float32
	Active      bool
	LocationID  int
	LocationImg LocationImg `gorm:"foreignKey:LocationID"`
}
