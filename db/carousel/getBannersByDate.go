package db

import (
	"time"
	"viandasApp/db"
	"viandasApp/dtos"
)

func GetBanners() (*[]dtos.BannersResponse, error) {

	db := db.GetDB()

	responseModel := []dtos.BannersResponse{}

	var dateTime time.Time = time.Now()

	if err := db.Table("location_imgs").
		Select("location_imgs.location").
		Joins("JOIN banners ON banners.location_id = location_imgs.id").
		Where("banners.active = 1 and ? BETWEEN banners.date_start AND banners.date_end", dateTime).
		Scan(&responseModel).Error; err.Error() != "Error 1146: Table 'viandas_db.banners' doesn't exist" {
		return &responseModel, err
	}

	return &responseModel, nil

}
