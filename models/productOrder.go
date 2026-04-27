package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductOrder struct {
	gorm.Model
	ClientName           string
	ClientLastName       string
	ProductCategoryTitle string
	ProductTitle         string
	Cant                 int
	Date                 time.Time
	Status               string
}
