package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetCategoryById(id int) (models.Category, error) {
	/* 	var db = db.ConnectDB()
	   	sqlDB, _ := db.DB()
	   	defer sqlDB.Close()
	*/

	db := db.GetDB()
	var categoryModel models.Category

	err := db.First(&categoryModel, id).Error

	return categoryModel, err

}
