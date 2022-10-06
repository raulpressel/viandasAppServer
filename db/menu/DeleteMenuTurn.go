package db

import (
	"viandasApp/db"
)

func DeleteTurnMenu(idMenu int, idTurn int) (bool, error) {

	//var modelTurnMenu models.TurnMenu

	//var modelDayMenu models.DayMenu

	var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	//return err.Error

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return false, err
	}

	id, err := ValidateTurnMenu(idMenu, idTurn)
	if err != nil {
		tx.Rollback()
		return false, err
	}

	if id == 0 {
		tx.Rollback()
		return false, err
	}

	if err := tx.Exec("DELETE FROM day_menus WHERE day_menus.turn_menu_id = ?", id).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	if err := tx.Exec("DELETE FROM turn_menus WHERE turn_menus.menu_id = ? AND turn_menus.turn_id = ?", idMenu, idTurn).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	return true, tx.Commit().Error

}
