package db

import (
	"time"
	"viandasApp/db"
	dbMenu "viandasApp/db/menu"
	"viandasApp/dtos"
	"viandasApp/models"
)

func GetOrders(date time.Time) (*dtos.OrderViewerResponse, error) {

	db := db.GetDB()

	var modelDayMenu models.DayMenu

	var modelOrder models.Order

	var modelClient models.Client

	

	modelDayOrder :=  []models.DayOrder{}

	modelTanda := []models.Tanda{}

	modelTandaAddress := []models.TandaAddress{}

	if err := db.Find(&modelTanda).Error; err != nil {
		return nil, err
	}
	
	for _, tanda := range modelTanda {
		if err := db.Find(&modelTandaAddress, "tanda_id = ?", tanda.ID).Error; err != nil {
			return nil, err
		}	

		var idAddresses []int

		for _, 	tandaAddress := range 	modelTandaAddress {
			idAddresses = append(idAddresses, tandaAddress.AddressID)
		}	

		if err := db.Find(&modelDayOrder, "address_id IN (?)", idAddresses).Error; err != nil {
			return nil, err
		}

		for _, dayOrder := range modelDayOrder{
			
			// reemplazar por FIRST o funciones que ya existen

			if err := db.Find(&modelDayMenu, "id = ? AND date = ?", dayOrder.DayMenuID, date).Error; err != nil {
				return nil, err
			}

			if err := db.Find(&modelOrder, "id = ?", dayOrder.OrderID).Error; err != nil {
				return nil, err
			}

			if err := db.Find(&modelClient, "id = ?", modelOrder.ClientID).Error; err != nil {
				return nil, err
			}

			//pathology hacer consulta q traiga 

			

			modelDayMenu.CategoryID

			
			
			


		}



	}

	




	
	if err := db.Find(&modelDayMenu, "date = ?", date.Format("2006-01-02")).Error; err != nil {
		return nil, err
	}

	for _, dayMenu := range modelDayMenu {

		if err := db.Find(&modelDayOrder, "day_menu_id = ?", dayMenu.ID).Error; err != nil {
			return nil, err
		}
	
		for _, dayOrder := range modelDayOrder {
			dayOrder.AddressID
		}
		

	}

		err = db.Table("day_orders").
			Select("day_orders.id, day_orders.observation, day_orders.amount, day_orders.status, day_orders.address_id, day_orders.day_menu_id").
			Where("day_orders.order_id = ?", ord.ID).
			Scan(&daysOrderModel).Error

		dayMenuModel, err := dbMenu.GetDayMenuById(daysOrderModel.DayMenuID)
		if err != nil {
			return dtoOrder, err
		}

		menuModel, err := dbMenu.GetMenuByTurnMenuID(dayMenuModel.TurnMenuID)
		if err != nil {
			return dtoOrder, err
		}

		orders.ID = ord.ID
		orders.OrderDate = ord.OrderDate
		orders.Observation = ord.Observation
		orders.Status = ord.Status
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

	return dtoOrder, err

}
