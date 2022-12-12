package db

import (
	"viandasApp/db"
	"viandasApp/dtos"
	"viandasApp/models"
)

func GetAllOrder(id int) (dtos.OrderViewerResponse, error) {

	db := db.GetDB()

	modelOrder := []models.Order{}

	var dtoOrder dtos.OrderViewerResponse

	var orders dtos.OrderResponse

	//responseAllMenu := []dtos.AllMenuResponse{}

	err := db.Table("orders").
		Select("orders.id, orders.order_date, orders.observation, orders.total, orders.status").
		Where("orders.client_id = ?", id).
		Scan(&modelOrder).Error

	for _, ord := range modelOrder {
		orders.ID = ord.ID
		orders.OrderDate = ord.OrderDate
		orders.Observation = ord.Observation
		orders.Status = ord.Status
		orders.Total = ord.Total

		dtoOrder.Order = append(dtoOrder.Order, orders)
	}

	/* 	for _, valor := range modelAllMenu {
		id, err := GetIdMenuActive(valor.Menuid)
		if err != nil {
			return responseAllMenu, err
		}
		if id > 0 {
			valor.IsCurrent = true
		}
		responseAllMenu = append(responseAllMenu, *valor.ToAllMenuResponse())
	} */

	return dtoOrder, err

}
