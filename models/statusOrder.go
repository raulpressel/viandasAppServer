package models

import (
	"gorm.io/gorm"
)

/*usuario es el modelo de usuario de la base de mysql*/

type StatusOrder struct {
	gorm.Model
	ID          int `gorm:"primary_key"`
	Description string
}