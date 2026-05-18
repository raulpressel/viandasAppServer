package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func DeleteProductOrder(id int) (bool, error) {
	dbc := db.GetDB()
	err := dbc.Model(&models.ProductOrder{}).Where("id = ?", id).Update("status", "deleted").Error
	return err == nil, err
}
