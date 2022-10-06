package db

import (
	"time"
	"viandasApp/db"
	"viandasApp/dtos"
)

func GetIdMenuActiveByDate(dateStart time.Time, dateEnd time.Time) (int, error) {
	var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var modelMenu dtos.Menu

	//var dateTime time.Time = time.Now()

	err := db.Table("menus").
		Select("menus.id as menuid").
		Where("? BETWEEN menus.date_start and menus.date_end OR ? BETWEEN menus.date_start and menus.date_end AND menus.active = 1", dateStart.Format("2006-01-02"), dateEnd.Format("2006-01-02")).
		Scan(&modelMenu).Error

	return modelMenu.Menuid, err

}
