package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func ValidateTurnId(idTurn int) (int, error) {
	var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var turnModel models.Turn

	err := db.Table("turns").
		Select("turns.id").
		Where("turns.id = ?", idTurn).
		Scan(&turnModel).Error

	return turnModel.ID, err

}
