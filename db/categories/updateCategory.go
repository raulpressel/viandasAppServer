package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func UpdateCategory(categoryModel models.Category) (bool, error) {

	var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	//userModel.Password, _ = EncryptPassword(userModel.Password)

	err := db.Save(&categoryModel)

	if err.Error != nil {
		return false, err.Error
	}
	return true, err.Error

}
