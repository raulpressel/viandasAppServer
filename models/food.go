package models

import (
	"gorm.io/gorm"
)

/*usuario es el modelo de usuario de la base de mysql*/

type Food struct {
	gorm.Model
	ID          int `gorm:"primary_key"`
	Title       string
	Description string
	Active      bool
	LocationID  *int
	LocationImg LocationImg `gorm:"foreignKey:LocationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
