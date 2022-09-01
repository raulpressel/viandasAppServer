package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetLocationImgById(id int) (models.LocationImg, error) {
	var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var locationModel models.LocationImg

	err := db.First(&locationModel, id).Error

	return locationModel, err

}
