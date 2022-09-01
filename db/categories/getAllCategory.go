package db

import (
	"viandasApp/db"
	"viandasApp/dtos"
)

func GetAllCategory() ([]dtos.AllCategoryResponse, error) {
	var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var responseModel []dtos.AllCategoryResponse

	err := db.Table("categories").
		Select("categories.id, categories.description").
		Scan(&responseModel).Error

	return responseModel, err

}
