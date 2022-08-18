package dtos

import "time"

type BannersRequest struct {
	OnlyActive bool `json:"onlyActive"`

	//LocationImg LocationImg `gorm:"foreignKey:Location,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type BannersResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	DateStart time.Time `json:"dateStart"`
	DateEnd   time.Time `json:"dateEnd"`
	Location  string    `json:"urlImage"`
}
