package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetTandaById(id int) (models.Tanda, error) {
	db := db.GetDB()

	var tandaModel models.Tanda

	err := db.First(&tandaModel, id).Error

	return tandaModel, err

}
