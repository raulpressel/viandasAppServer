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
				modelOrder, err := GetModelOrderById(modelOrders[i].ID)
				if err != nil {
					return false, err
				}
				modelOrder.StatusOrderID = 2

				if err := tx.Save(&modelOrder).Error; err != nil {
					tx.Rollback()
					return false, err
				}

			}

		}
	}

	return true, tx.Commit().Error

}
