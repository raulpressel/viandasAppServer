package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetClientByIDUser(idkl string) (bool, error) {

	db := db.GetDB()
	var categoryModel models.Category

	err := db.First(&categoryModel, idkl).Error

	return true, err
}
