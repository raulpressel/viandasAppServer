package models

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	ID             int    `gorm:"primary_key"`
	Name           string //given_name
	LastName       string //family_name
	Email          string //email
	IDUserKL       string //sub
	PhonePrimary   string
	PhoneSecondary string
	Observation    string
	BornDate       time.Time
	Active         bool
}
