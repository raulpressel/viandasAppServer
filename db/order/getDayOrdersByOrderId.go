package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetDayOrdersByOrderId(id int) ([]models.DayOrder, error) {
	db := db.GetDB()

	var dayOrdersModel []models.DayOrder

	err := db.Table("day_orders").
		Joins("left join day_menus on day_menus.id = day_orders.day_menu_id").
		Where("day_orders.order_id = ?", id).
		Find(&dayOrdersModel).Error

	return dayOrdersModel, err

}
