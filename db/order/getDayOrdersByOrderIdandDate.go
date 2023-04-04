package db

import (
	"time"
	"viandasApp/db"
	"viandasApp/models"
)

func GetDayOrdersByOrderIdandDate(id int) ([]models.DayOrder, error) {
	db := db.GetDB()

	var dayOrdersModel []models.DayOrder

	date := time.Now()

	if date.Hour() > 12 {
		date = date.AddDate(0, 0, 1)
	}

	err := db.Table("day_orders").
		Joins("left join day_menus on day_menus.id = day_orders.day_menu_id").
		Where("day_orders.order_id = ?", id).
		Where("date(day_menus.date) >= ?", date.Format("2006-01-02")).
		Find(&dayOrdersModel).Error

	return dayOrdersModel, err

}
