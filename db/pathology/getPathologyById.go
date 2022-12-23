package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetPathologyById(id int) (models.Pathology, error) {
	db := db.GetDB()

	var pathologyModel models.Pathology

	err := db.First(&pathologyModel, id).Error

	return pathologyModel, err

}
