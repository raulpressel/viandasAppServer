package db

import (
	"viandasApp/db"
)

func UpdateDayOrderAddress(idDayOrder int, idAddress int) (bool, error) {

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

	dayOrderModel, err := GetDayOrderById(idDayOrder)
	if err != nil {
		tx.Rollback()
		return false, err
	}

	if dayOrderModel.ID == 0 {
		tx.Rollback()
		return false, err
	}

	dayOrderModel.AddressID = idAddress

	if err := tx.Save(&dayOrderModel).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	return true, tx.Commit().Error

}
