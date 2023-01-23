package db

import (
	"sort"
	"viandasApp/db"
	dbAdd "viandasApp/db/address"
	dbCat "viandasApp/db/categories"
	dbCity "viandasApp/db/city"
	dbFood "viandasApp/db/food"
	dbImg "viandasApp/db/img"
	dbMenu "viandasApp/db/menu"
	"viandasApp/dtos"
	"viandasApp/models"
)

func GetOrderById(idOrder int) (dtos.FullOrderResponse, error) {

	db := db.GetDB()

	var orderModel models.Order

	var responseOrder dtos.FullOrderResponse

	daysOrderModel := []models.DayOrder{}

	dayOrderResponse := []dtos.DayOrderResponse{}

	var imgCatModel models.LocationImg

	var imgFoodModel models.LocationImg

	err := db.Table("orders").
		Select("orders.id, orders.order_date, orders.observation, orders.total, orders.status").
		Where("orders.id = ?", idOrder).
		Scan(&orderModel).Error

	err = db.Table("day_orders").
		Select("day_orders.id, day_orders.observation, day_orders.amount, day_orders.status, day_orders.address_id, day_orders.day_menu_id").
		Where("day_orders.order_id = ?", orderModel.ID).
		Scan(&daysOrderModel).Error

	for _, valor := range daysOrderModel {

		dayMenuModel, err := dbMenu.GetDayMenuById(valor.DayMenuID)
		if err != nil {
			return responseOrder, err
		}

		foodModel, err := dbFood.GetFoodById(dayMenuModel.FoodID)
		if err != nil {
			return responseOrder, err
		}

		if foodModel.LocationID != nil {
			imgFoodModel, _ = dbImg.GetLocationImgById(*foodModel.LocationID)
		}

		catModel, err := dbCat.GetCategoryById(dayMenuModel.CategoryID)
		if err != nil {
			return responseOrder, err
		}
		if catModel.LocationID != nil {
			imgCatModel, _ = dbImg.GetLocationImgById(*catModel.LocationID)
		}

		addressModel, err := dbAdd.GetAddressById(valor.AddressID)

		if err != nil {
			return responseOrder, err
		}

		cityModel, err := dbCity.GetCityById(addressModel.CityID)
		if err != nil {
			return responseOrder, err
		}

		dayOrder := dtos.DayOrderResponse{
			ID:          valor.ID,
			Date:        dayMenuModel.Date,
			Amount:      valor.Amount,
			Observation: valor.Observation,
			Status:      valor.Status,
			Food: dtos.FoodResponse{
				ID:          foodModel.ID,
				Title:       foodModel.Title,
				Description: foodModel.Description,
				Location:    imgFoodModel.Location,
			},
			Category: dtos.CategoryResponse{
				ID:          catModel.ID,
				Description: catModel.Description,
				Title:       catModel.Title,
				Price:       catModel.Price,
				Color:       catModel.Color,
				Location:    imgCatModel.Location,
			},
			Address: dtos.AddressRespone{
				ID:          addressModel.ID,
				Street:      addressModel.Street,
				Number:      addressModel.Number,
				Floor:       addressModel.Floor,
				Departament: addressModel.Departament,
				Observation: addressModel.Observation,
				City: dtos.AllCityResponse{
					ID:          cityModel.ID,
					Description: cityModel.Description,
					CP:          cityModel.CP,
				},
			},
		}

		dayOrderResponse = append(dayOrderResponse, dayOrder)

	}

	sort.SliceStable(dayOrderResponse, func(i, j int) bool {

		return dayOrderResponse[i].Date.Before(dayOrderResponse[j].Date)

	})

	responseOrder.ID = orderModel.ID
	responseOrder.Observation = orderModel.Observation
	responseOrder.OrderDate = orderModel.OrderDate
	responseOrder.Status = orderModel.Status
	responseOrder.Total = orderModel.Total
	responseOrder.DayOrder = dayOrderResponse

	return responseOrder, err

}
