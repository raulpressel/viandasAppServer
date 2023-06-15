package db

import (
	"viandasApp/db"
	"viandasApp/dtos"
)

func GetAllMenu() ([]dtos.AllMenuResponse, error) {
	/* var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	*/
	db := db.GetDB()
	modelAllMenu := []dtos.AllMenu{}

	responseAllMenu := []dtos.AllMenuResponse{}

	err := db.Table("menus").
		Select("menus.id as menuid, menus.date_start as menudatestart, menus.date_end menudateend, turns.id as turnid, turns.description as turndescription").
		Joins("left JOIN turn_menus ON turn_menus.menu_id = menus.id").
		Joins("left JOIN turns ON turn_menus.turn_id = turns.id").
		Where("menus.active = 1").
		Order("menus.date_start desc").
		Order("turn_menus.turn_id asc").
		Scan(&modelAllMenu).Error

	for _, valor := range modelAllMenu {
		id, err := GetIdMenuActive(valor.Menuid)
		if err != nil {
			return responseAllMenu, err
		}
		if id > 0 {
			valor.IsCurrent = true
		}
		responseAllMenu = append(responseAllMenu, *valor.ToAllMenuResponse())
	}

	return responseAllMenu, err

}
