package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetDeliveryDriverIdByTandaId(id int) (int, error) {
	db := db.GetDB()

	var tandaModel models.Tanda

	err := db.First(&tandaModel, id).Error

	return tandaModel.DeliveryDriverID, err

}
