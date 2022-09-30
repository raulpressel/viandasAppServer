package db

import (
	"time"
	"viandasApp/db"
	"viandasApp/dtos"
)

func GetIdMenuActive() (int, error) {
	var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var modelMenu dtos.AllMenu

	var dateTime time.Time = time.Now()

	err := db.Table("menus").
		Select("menus.id as menuid").
		Where("? BETWEEN menus.date_start and menus.date_end", dateTime.Format("2006-01-02")).
		Scan(&modelMenu).Error

	return modelMenu.Menuid, err

}
