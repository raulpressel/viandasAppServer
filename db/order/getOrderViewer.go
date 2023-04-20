package db

import (
	"viandasApp/db"
	dbMenu "viandasApp/db/menu"
	"viandasApp/dtos"
	"viandasApp/models"
)

func GetOrderViewer(id int) (*dtos.OrderViewerResponse, error) {

	db := db.GetDB()

	modelOrder := []models.Order{}

	var daysOrderModel models.DayOrder

	var dtoOrder dtos.OrderViewerResponse

	var orders dtos.OrderResponse

	//responseAllMenu := []dtos.AllMenuResponse{}

	err := db.Table("orders").
		Select("orders.id, orders.order_date, orders.observation, orders.total, orders.status_order_id").
		Where("orders.client_id = ?", id).
		Scan(&modelOrder).Error

	for _, ord := range modelOrder {

		err = db.Table("day_orders").
			Select("day_orders.id, day_orders.observation, day_orders.amount, day_orders.status, day_orders.address_id, day_orders.day_menu_id").
			Where("day_orders.order_id = ?", ord.ID).
			Scan(&daysOrderModel).Error

		dayMenuModel, err := dbMenu.GetDayMenuById(daysOrderModel.DayMenuID)
		if err != nil {
			return nil, err
		}

		menuModel, err := dbMenu.GetMenuByTurnMenuID(dayMenuModel.TurnMenuID)
		if err != nil {
			return nil, err
		}

		modelStatusOrder, err := GetStatusOrder(ord.StatusOrderID)
		if err != nil {
			return nil, err
		}

		orders.ID = ord.ID
		orders.OrderDate = ord.OrderDate
		orders.Observation = ord.Observation
		orders.Status.ID = modelStatusOrder.ID
		orders.Status.Description = modelStatusOrder.Description
		orders.Total = ord.Total
		orders.DateStart = menuModel.DateStart
		orders.DateEnd = menuModel.DateEnd

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

	return &dtoOrder, err

}
