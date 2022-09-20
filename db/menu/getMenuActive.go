package db

import (
	"time"
	"viandasApp/db"
	"viandasApp/dtos"
)

func GetMenuActive() ([]dtos.MenuResponse, error) {
	var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	responseModel := []dtos.MenuResponse{}

	var dateTime time.Time = time.Now()

	err := db.Table("day_menus").
		Select("day_menus.date, menus.turn_id, day_menus.food_id, foods.id as food, foods.title as foodtitle, categories.id as categoryid, categories.description as categorydescription").
		Joins("left JOIN foods ON foods.id = day_menus.food_id").
		Joins("left JOIN categories ON foods.category_id = categories.id").
		Joins("left JOIN menus on day_menus.menu_id = menus.id").
		Where("? BETWEEN menus.date_start AND menus.date_end", dateTime).
		Scan(&responseModel).Error

	return responseModel, err

}
