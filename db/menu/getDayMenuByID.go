package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetDayMenuById(id int) (models.DayMenu, error) {
	var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var dayMenuModel models.DayMenu

	err := db.First(&dayMenuModel, id).Error

	return dayMenuModel, err

}
