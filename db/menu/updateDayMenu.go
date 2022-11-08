package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func UpdateDayMenu(dayMenu models.DayMenu) (bool, error) {

	/* var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close() */

	db := db.GetDB()

	err := db.Save(&dayMenu)

	if err.Error != nil {
		return false, err.Error
	}
	return true, err.Error

}
