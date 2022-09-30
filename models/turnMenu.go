package models

import (
	"gorm.io/gorm"
)

/*usuario es el modelo de usuario de la base de mysql*/

type TurnMenu struct {
	gorm.Model
	ID     int `gorm:"primary_key"`
	MenuID int
	Menu   Menu `gorm:"foreignKey:MenuID"`
	TurnId int
	Turn   Turn `gorm:"foreignKey:TurnId"`
}
