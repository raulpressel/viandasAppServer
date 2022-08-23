package db

import (
	"fmt"
	"viandasApp/dtos"
)

func GetAllBanners() []dtos.AllBannersResponse {
	var db = ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	responseModel := []dtos.AllBannersResponse{}

	err := db.Table("location_imgs").
		Select("banners.id, banners.title, banners.date_start, banners.date_end, location_imgs.location").
		Joins("JOIN banners ON banners.location_id = location_imgs.id").
		Scan(&responseModel).Error

	if err != nil {
		fmt.Println(err)
		return responseModel
	}

	return responseModel

}
