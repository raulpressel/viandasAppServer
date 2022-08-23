package db

import (
	"fmt"
	"time"
	"viandasApp/dtos"
)

func GetBanners() []dtos.BannersResponse {
	var db = ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	responseModel := []dtos.BannersResponse{}

	var dateTime time.Time = time.Now()

	err := db.Table("location_imgs").
		Select("location_imgs.location").
		Joins("JOIN banners ON banners.location_id = location_imgs.id").
		Where("? BETWEEN banners.date_start AND banners.date_end", dateTime).
		Scan(&responseModel).Error

	if err != nil {
		fmt.Println(err)
		return responseModel
	}

	return responseModel

}
