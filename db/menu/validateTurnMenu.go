package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func ValidateTurnMenu(idMenu int, idTurn int) (int, error) {
	var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var turnMenuModel models.TurnMenu

	err := db.Table("turn_menus").
		Select("turn_menus.id").
		Where("turn_menus.menu_id = ? and turn_menus.turn_id = ?", idMenu, idTurn).
		Scan(&turnMenuModel).Error

	return turnMenuModel.ID, err

}
