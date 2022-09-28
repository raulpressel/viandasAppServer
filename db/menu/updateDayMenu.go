package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func UpdateDayMenu(dayMenu models.DayMenu) (bool, error) {

	var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	/* 	SELECT * FROM day_menus
	left join foods on food_id = foods.id
	where foods.category_id = (select c.id from foods f, categories c where f.category_id = c.id and f.id = 2)
	and date = "2022-09-01 00:00:00.00"


	update day_menus
	left join foods on food_id = foods.id
	set food_id = 1
	where foods.category_id = (select c.id from foods f, categories c where f.category_id = c.id and f.id = 2)
	and date = "2022-09-01 00:00:00.00" */

	//userModel.Password, _ = EncryptPassword(userModel.Password)

	err := db.Save(&dayMenu)

	if err.Error != nil {
		return false, err.Error
	}
	return true, err.Error

}
