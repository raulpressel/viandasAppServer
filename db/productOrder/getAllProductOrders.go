package db

import (
	"viandasApp/db"
	"viandasApp/dtos"
	"viandasApp/models"
)

func GetAllProductOrders() ([]dtos.ProductOrderResponse, error) {
	dbc := db.GetDB()
	var orders []models.ProductOrder
	err := dbc.Preload("Products").Find(&orders).Error
	if err != nil {
		return nil, err
	}

	var result []dtos.ProductOrderResponse
	for _, o := range orders {
		var items []dtos.ProductOrderItemResponse
		for _, p := range o.Products {
			items = append(items, dtos.ProductOrderItemResponse{
				ID:                   p.ID,
				ProductTitle:         p.ProductTitle,
				ProductCategoryTitle: p.ProductCategoryTitle,
				Cant:                 p.Cant,
				Status:               p.Status,
			})
		}
		result = append(result, dtos.ProductOrderResponse{
			ID:             o.ID,
			ClientName:     o.ClientName,
			ClientLastName: o.ClientLastName,
			Date:           o.Date,
			Status:         o.Status,
			Products:       items,
		})
	}
	return result, nil
}
