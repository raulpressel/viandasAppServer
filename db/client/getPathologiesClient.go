package db

import (
	"viandasApp/db"
	"viandasApp/dtos"
)

func GetPathologiesClient(id int) ([]dtos.PathologyResponse, error) {

	db := db.GetDB()

	/* var pathologiesModel []models.Pathology

	err := db.Find(&pathologiesModel).
		Where("select pathology_id from client_pathologies where client_id =", id).Error */

	var response []dtos.PathologyResponse

	//err := db.Find(&responseModel).Error

	err := db.Table("pathologies").
		Select("pathologies.id, pathologies.description").
		Where("pathologies.active = 1").
		Where("pathologies.id IN (select pathology_id from client_pathologies where client_id = ?)", id).
		Scan(&response).Error

	return response, err

}
