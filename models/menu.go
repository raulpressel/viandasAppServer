package models

import (
	"time"

	"gorm.io/gorm"
)

/*usuario es el modelo de usuario de la base de mysql*/

type Menu struct {
	gorm.Model
	ID int `gorm:"primary_key"`

	DateStart time.Time
	DateEnd   time.Time
	/* TurnId    int
	TurnMenu  TurnMenu `gorm:"foreignKey:TurnId"` */
}
