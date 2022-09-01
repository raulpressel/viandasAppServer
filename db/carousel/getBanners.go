package db

import (
	"time"
	"viandasApp/db"
	"viandasApp/dtos"
)

func GetBanners() ([]dtos.BannersResponse, error) {
	var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	responseModel := []dtos.BannersResponse{}

	var dateTime time.Time = time.Now()

	err := db.Table("location_imgs").
		Select("location_imgs.location").
		Joins("JOIN banners ON banners.location_id = location_imgs.id").
		Where("? BETWEEN banners.date_start AND banners.date_end", dateTime).
		Scan(&responseModel).Error

	return responseModel, err

}
