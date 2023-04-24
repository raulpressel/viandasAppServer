package db

import (
	"viandasApp/db"
)

func CheckFeriado(dates []string) (bool, error) {

	db := db.GetDB()
	var cant int = 0

	db.Table("day_menus").
		Select("COUNT(day_menus.id)").
		Joins("left JOIN foods ON foods.id = day_menus.food_id").
		Where("date(day_menus.date) IN (?)", dates).
		Where("foods.title LIKE ?", "%"+"Feriado"+"%").
		Scan(&cant)

	return cant > 0, db.Error
}
