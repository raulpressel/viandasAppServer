package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetBannerById(id int) (models.Banner, error) {
	var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var bannerModel models.Banner

	err := db.First(&bannerModel, id).Error

	return bannerModel, err

}
