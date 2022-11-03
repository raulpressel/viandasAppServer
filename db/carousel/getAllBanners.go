package db

import (
	"viandasApp/db"
	"viandasApp/dtos"
)

func GetAllBanners() ([]dtos.AllBannersResponse, error) {
	/* 	var db = db.ConnectDB()
	   	sqlDB, _ := db.DB()
	   	defer sqlDB.Close()
	*/

	db := db.GetDB()
	responseModel := []dtos.AllBannersResponse{}

	err := db.Table("location_imgs").
		Select("banners.id, banners.title, banners.date_start, banners.date_end, location_imgs.location").
		Joins("JOIN banners ON banners.location_id = location_imgs.id").
		Where("banners.active = 1").
		Scan(&responseModel).Error

	return responseModel, err

}
