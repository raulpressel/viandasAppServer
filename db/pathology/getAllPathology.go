package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetAllPathology() ([]models.Pathology, error) {

	db := db.GetDB()

	var responseModel []models.Pathology

	err := db.Find(&responseModel).Error

	return responseModel, err

}
