package dtos

import (
	"viandasApp/models"
)

type UploadBanner struct {
	Banner models.Banner `json:"banner"`

	//LocationImg LocationImg `gorm:"foreignKey:Location,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
