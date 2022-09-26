package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func UploadMenu(dayModel []models.DayMenu, menuModel models.Menu) (bool, error) {

	//var dayModel []models.DayMenu

	var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return false, err
	}

	if err := tx.Save(&menuModel).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	for i, _ := range dayModel {

		dayModel[i].MenuID = menuModel.ID
	}

	if err := tx.CreateInBatches(&dayModel, len(dayModel)).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	return true, tx.Commit().Error

}
