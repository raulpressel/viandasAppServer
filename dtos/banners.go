package dtos

import "time"

type BannersResponse struct {
	Location string `json:"urlImage"`

	//LocationImg LocationImg `gorm:"foreignKey:Location,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type AllBannersResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	DateStart time.Time `json:"dateStart"`
	DateEnd   time.Time `json:"dateEnd"`
	Location  string    `json:"urlImage"`
}
