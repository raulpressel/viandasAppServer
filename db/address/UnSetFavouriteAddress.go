package db

import (
	"viandasApp/db"
)

func UnSetFavouriteAddress(idClient int, idAddress int) (bool, error) {

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

	return true, tx.Commit().Error
}
