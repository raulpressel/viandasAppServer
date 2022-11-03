package models

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	ID              int `gorm:"primary_key"`
	Name            string
	LastName        string
	Email           string
	IDUserKL        string
	Tel1            string
	Tel2            string
	Observation     string
	FechaNacimiento time.Time
}
