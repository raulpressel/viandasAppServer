package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetMenuByTurnMenuID(turnMenuID int) (models.Menu, error) {
	/* var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close() */

	db := db.GetDB()

	var turnMenuModel models.TurnMenu

	var menuModel models.Menu

	err := db.First(&turnMenuModel, turnMenuID).Error

	err = db.First(&menuModel, turnMenuModel.MenuID).Error

	return menuModel, err

}
