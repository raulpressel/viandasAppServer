package db

import (
	"time"
	"viandasApp/db"
	"viandasApp/models"
)

func GetDayOrdersByOrderIdAndDay(idOrder int, date time.Time) ([]models.DayOrder, error) {
	db := db.GetDB()

	var dayOrdersModel []models.DayOrder

	err := db.Table("day_orders").
		Select("day_orders.*").
		Joins("left join day_menus on day_menus.id = day_orders.day_menu_id").
		Where("day_orders.order_id = ?", idOrder).
		Where("date(day_menus.date) = ?", date.Format("2006-01-02")).
		Find(&dayOrdersModel).Error

	return dayOrdersModel, err

}
