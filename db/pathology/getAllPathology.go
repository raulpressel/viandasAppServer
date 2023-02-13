package db

import (
	"viandasApp/db"
	"viandasApp/dtos"
)

func GetAllPathology() ([]dtos.PathologyResponse, error) {

	db := db.GetDB()

	var responseModel []dtos.PathologyResponse

	err := db.Table("pathologies").
		Select("pathologies.id, pathologies.description, pathologies.color").
		Where("pathologies.active = 1").
		Scan(&responseModel).Error

	return responseModel, err

}
