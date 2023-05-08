package db

import (
	"time"
	"viandasApp/db"
	"viandasApp/models"
)

func FinishedOrder() (bool, error) {

	db := db.GetDB()

	var modelOrders []models.Order

	//var modelDayMenu models.DayMenu

	date := time.Now()

	date = date.AddDate(0, 0, 2)

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return false, err
	}

	if err := tx.Table("orders").
		Where("orders.status_order_id = 1").
		Find(&modelOrders).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	if len(modelOrders) > 0 {

		for i := range modelOrders {

			var dateC time.Time

			if err := tx.Table("day_menus").Select("MAX(day_menus.date)").
				Joins("left join day_orders on day_menus.id = day_orders.day_menu_id").
				Where("day_orders.order_id = ?", modelOrders[i].ID).
				Find(&dateC).Error; err != nil {
				tx.Rollback()
				return false, err
			}

			if date.After(dateC) {
				if err := tx.Model(&modelOrders[i]).Update("status_order_id", 2).Error; err != nil {
					return false, err
				}

				/*

					modelOrder, err := GetModelOrderById(modelOrders[i].ID)
					if err != nil {
						return false, err
					} */
				/* modelOrders.StatusOrderID = 2

				modelOrders[i] = modelOrder */

			}

			/* if err := tx.Save(&modelOrders).Error; err != nil {
				tx.Rollback()
				return false, err
			} */

		}
	}

	//select distinct(orders.id) from orders join day_orders on order_id = orders.id join day_menus on day_menu_id = day_menus.id where day_menus.date < "2023-6-4" and orders.status_order_id = 1;

	return true, tx.Commit().Error

}
