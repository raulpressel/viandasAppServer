package db

import (
	"errors"
	"viandasApp/db"
	"viandasApp/models"

	"gorm.io/gorm"
)

func CheckExistTandaByAddressId(idAddress int) (int, error) {

	var idTanda int

	var tandaAddressModel models.TandaAddress

	db := db.GetDB()

	if err := db.First(&tandaAddressModel, "address_id = ?", idAddress).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			idTanda = 0
			return idTanda, nil
		}
	}

	return tandaAddressModel.TandaID, db.Error

}
