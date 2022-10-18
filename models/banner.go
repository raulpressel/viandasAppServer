package models

import (
	"time"

	"gorm.io/gorm"
)

/*usuario es el modelo de usuario de la base de mysql*/

type Banner struct {
	gorm.Model
	ID          int `gorm:"primary_key"`
	Title       string
	DateStart   time.Time
	DateEnd     time.Time
	Active      bool
	LocationID  *int
	LocationImg LocationImg `gorm:"foreignKey:LocationID"`
	//LocationImg LocationImg `gorm:"foreignKey:Location,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
