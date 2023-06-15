package db

import (
	"time"
	"viandasApp/db"
)

func DeleteMenu(idMenu int, idTurn int) (bool, error) {

	var dateTime time.Time = time.Now()

	db := db.GetDB()

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

	if err := tx.Exec("UPDATE menus set active = 0 WHERE menus.id  = ? AND menus.date_end < ?", idMenu, dateTime.Format("2006-01-02")).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	return true, tx.Commit().Error

}
