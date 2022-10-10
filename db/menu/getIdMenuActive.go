package db

import (
	"time"
	"viandasApp/db"
	"viandasApp/dtos"
)

func GetIdMenuActive(idMenu int) (int, error) {
	var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var modelMenu dtos.Menu

	var dateTime time.Time = time.Now()

	err := db.Table("menus").
		Select("menus.id as menuid").
		Where("? BETWEEN menus.date_start and menus.date_end AND menus.id = ? ", dateTime.Format("2006-01-02"), idMenu).
		Scan(&modelMenu).Error

	return modelMenu.Menuid, err

}
