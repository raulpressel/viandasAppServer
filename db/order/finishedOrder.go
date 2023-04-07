package db

import (
	"time"
	"viandasApp/db"
	"viandasApp/models"
)

func FinishedOrder() (bool, error) {

	db := db.GetDB()

	var modelOrders []models.Order

	date := time.Now()

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return false, err
	}

	if err := tx.Distinct("orders.id").
		Joins("join day_orders on day_orders.order_id = orders.id").
		Joins("left join day_menus on day_menus.id = day_orders.day_menu_id").
		Where("date(day_menus.date) < ?", date.Format("2006-01-02")).
		Where("orders.status_order_id = 1").
		Find(&modelOrders).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	if len(modelOrders) > 0 {

		for i := range modelOrders {

			modelOrder, err := GetModelOrderById(modelOrders[i].ID)
			if err != nil {
				return false, err
			}
			modelOrder.StatusOrderID = 2

			modelOrders[i] = modelOrder

		}

		if err := tx.Save(&modelOrders).Error; err != nil {
			tx.Rollback()
			return false, err
		}
	}

	//select distinct(orders.id) from orders join day_orders on order_id = orders.id join day_menus on day_menu_id = day_menus.id where day_menus.date < "2023-6-4" and orders.status_order_id = 1;

	return true, tx.Commit().Error

}
