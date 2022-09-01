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
	CategoryID  int
	Category    Category `gorm:"foreignKey:CategoryID"`
	Active      bool
	LocationID  int
	LocationImg LocationImg `gorm:"foreignKey:LocationID"`
	//LocationImg LocationImg `gorm:"foreignKey:Location,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
