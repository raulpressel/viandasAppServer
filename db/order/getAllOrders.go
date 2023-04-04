package db

import (
	"time"
	"viandasApp/db"
	dbClient "viandasApp/db/client"
	"viandasApp/dtos"
	"viandasApp/models"
)

func GetAllOrders(date bool, dateStart time.Time, dateEnd time.Time, active bool, inactive bool, paid bool, notPaid bool) (*[]dtos.OrdersRes, error) {

	db := db.GetDB()

	modelOrder := []models.Order{}

	query := db.Model(&modelOrder)

	if active != inactive {

		if active {
			query = query.Where("orders.status = 1")
		}
		if inactive {
			query = query.Where("orders.status = 0")
		}
	}

	if paid != notPaid {

		if paid {
			query = query.Where("orders.paid = 1")
		}
		if notPaid {
			query = query.Where("orders.paid = 0")
		}
	}

	if date {
		query = query.Where("orders.order_date BETWEEN ? AND ?", dateStart, dateEnd)
	}

	if err := query.Find(&modelOrder).Error; err != nil {
		return nil, db.Error
	}

	var ordersRes []dtos.OrdersRes

	for _, valor := range modelOrder {

		modelClient, err := dbClient.GetClientById(valor.ClientID)
		if err != nil {
			return nil, err
		}

		orderRes := dtos.OrdersRes{
			ID:          valor.ID,
			OrderDate:   valor.OrderDate,
			Observation: valor.Observation,
			Total:       valor.Total,
			Status:      valor.Status,
			Paid:        valor.Paid,
			Client: dtos.Client{
				ID:             modelClient.ID,
				Name:           modelClient.Name,
				LastName:       modelClient.LastName,
				Email:          modelClient.Email,
				PhonePrimary:   modelClient.PhonePrimary,
				PhoneSecondary: modelClient.PhoneSecondary,
				ObsClient:      modelClient.Observation,
				BornDate:       modelClient.BornDate,
			},
		}

		ordersRes = append(ordersRes, orderRes)

	}

	return &ordersRes, db.Error

}
