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

	err := db.Table("menus").
		Select("menus.id as menuid").
		Where("? BETWEEN menus.date_start and menus.date_end OR ? BETWEEN menus.date_start and menus.date_end OR menus.date_start BETWEEN ? and ? OR menus.date_end BETWEEN ? and ?", dateStart.Format("2006-01-02"), dateEnd.Format("2006-01-02"), dateStart.Format("2006-01-02"), dateEnd.Format("2006-01-02"), dateStart.Format("2006-01-02"), dateEnd.Format("2006-01-02")).
		Scan(&modelMenu).Error

	return modelMenu.Menuid, err

}
