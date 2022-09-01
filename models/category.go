package models

import (
	"gorm.io/gorm"
)

/*usuario es el modelo de usuario de la base de mysql*/

type Category struct {
	gorm.Model
	ID          int `gorm:"primary_key"`
	Description string
	//LocationImg LocationImg `gorm:"foreignKey:Location,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
