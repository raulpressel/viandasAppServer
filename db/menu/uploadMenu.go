package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func UploadMenu(dayModel []models.DayMenu, menuModel models.Menu, turnMenuModel models.TurnMenu) (bool, error) {

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

	id, err := GetIdMenuActiveByDate(menuModel.DateStart, menuModel.DateEnd)
	if err != nil {
		tx.Rollback()
		return false, err
	}

	menuModel.ID = id

	if menuModel.ID == 0 {

		if err := tx.Save(&menuModel).Error; err != nil {
			tx.Rollback()
			return false, err
		}
	}

	turnMenuModel.MenuID = menuModel.ID

	idTurnMenu, err := ValidateTurnMenu(turnMenuModel.MenuID, turnMenuModel.TurnId)
	if err != nil {
		tx.Rollback()
		return false, err
	}

	if idTurnMenu == 0 {
		if err := tx.Save(&turnMenuModel).Error; err != nil {
			tx.Rollback()
			return false, err
		}
	} else {
		tx.Rollback()
		return false, err
	}

	for i := range dayModel {

		dayModel[i].TurnMenuID = turnMenuModel.ID
	}

	if err := tx.CreateInBatches(&dayModel, len(dayModel)).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	return true, tx.Commit().Error

}
