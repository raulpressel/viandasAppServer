package models

import (
	"time"

	"gorm.io/gorm"
)

/*usuario es el modelo de usuario de la base de mysql*/

type Banner struct {
	gorm.Model
	ID          int64
	Title       string
	DateStart   time.Time
	DateEnd     time.Time
	Status      bool
	LocationID  int
	LocationImg LocationImg `gorm:"foreignKey:LocationID"`
	//LocationImg LocationImg `gorm:"foreignKey:Location,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

/* type User2 struct {
	gorm.Model
	Name         string
	CompanyRefer int
	Company      Company
}

type Company struct {
	ID   int
	Name string
} */

/* type User struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   Company ``
} */

/* type User struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   Company
  }

  type Company struct {
	ID   int
	Name string
  } */
