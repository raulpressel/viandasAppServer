package db

import (
	"time"
	"viandasApp/db"
	"viandasApp/dtos"
)

func GetDayMenuByDate(date time.Time) ([]dtos.DayMenuDateDto, error) {
	var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	dayMenuDto := []dtos.DayMenuDateDto{}

	var dateTime time.Time = time.Now()

	err := db.Table("day_menus").
		Select("day_menus.id as id, day_menus.date as date, day_menus.food_id as foodid, foods.title as foodtitle, foods.description as fooddescription , foods.category_id as categoryid, categories.description as categorydescription, categories.title as categorytitle , categories.price as categoryprice, location_imgs.location as foodlocation").
		Joins("left JOIN turn_menus on turn_menus.id = day_menus.turn_menu_id").
		Joins("left JOIN menus on menus.id = turn_menus.menu_id").
		Joins("left JOIN foods on foods.id = day_menus.food_id").
		Joins("left JOIN categories on foods.category_id = categories.id").
		Joins("left JOIN location_imgs ON foods.location_id = location_imgs.id").
		Where("day_menus.date = ? and ? BETWEEN menus.date_start and menus.date_end", date.Format("2006-01-02"), dateTime.Format("2006-01-02")).
		Scan(&dayMenuDto).Error

	return dayMenuDto, err

}
