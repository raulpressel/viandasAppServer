package handlers

import (
	"encoding/json"
	"net/http"
	"sort"
	"time"

	dbAddress "viandasApp/db/address"
	dbCategory "viandasApp/db/categories"
	dbClient "viandasApp/db/client"
	dbMenu "viandasApp/db/menu"
	dbSetting "viandasApp/db/setting"
	"viandasApp/dtos"
	"viandasApp/models"
)

var priceTable = []struct {
	MinAmount   int
	MaxAmount   int
	PriceFactor float32
}{
	{1, 2, 1.0},
	{3, 4, 1.5},
	{5, 6, 2.0},
	{7, 8, 2.5},
	{9, 10, 3.0},
}

type responseTotal struct {
	SubTotal float32 `json:"subTotal"`
	Discount float32 `json:"discount"`
	Total    float32 `json:"total"`
	Delivery float32 `json:"delivery"`
}

func CalcPriceOrder(rw http.ResponseWriter, r *http.Request) {

	var orderDto dtos.OrderRequest

	var orderModel models.Order

	var dOrderModel models.DayOrder

	var response responseTotal

	err := json.NewDecoder(r.Body).Decode(&orderDto)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	clientModel, err := dbClient.GetClientById(orderDto.IDClient)

	if err != nil {
		http.Error(rw, "Ocurrio un error al obtener el ID del cliente "+err.Error(), http.StatusInternalServerError)
		return
	}

	if clientModel.ID == 0 {
		http.Error(rw, "El Cliente enviado no existe ", http.StatusBadRequest)
		return
	}

	orderModel.ClientID = orderDto.IDClient

	orderModel.Observation = orderDto.Observation

	//orderModel.Total = orderDto.Total

	orderModel.StatusOrderID = 1 //se da de alta orden y queda con estado 1 - Activa

	orderModel.Paid = false

	orderModel.OrderDate, err = time.Parse(time.RFC3339, orderDto.Date)
	if err != nil {
		http.Error(rw, "Error en el formato de fecha recibido "+err.Error(), http.StatusBadRequest)
		return
	}

	var amount int = 0

	var price float32 = 0.0

	//var priceDelivery float32 = 0.0

	for _, day := range orderDto.DaysOrderRequest {

		if day.Amount > 0 {

			//dOrderModel.Amount = day.Amount

			amount = amount + dOrderModel.Amount

			dayMenuModel, err := dbMenu.GetDayMenuById(day.IDDayFood)

			if err != nil {
				http.Error(rw, "Ocurrio un error al obtener el ID del menu "+err.Error(), http.StatusInternalServerError)
				return
			}

			if dayMenuModel.ID == 0 {
				http.Error(rw, "El Day Menu enviado no existe ", http.StatusBadRequest)
				return
			}

			cat, _ := dbCategory.GetCategoryById(dayMenuModel.CategoryID)

			price = price + cat.Price

			dOrderModel.DayMenuID = day.IDDayFood

			dOrderModel.Observation = day.Observation

			//dOrderModel.Active = true

			dOrderModel.Status = true

			addressModel, err := dbAddress.GetAddressById(day.IDAddress)

			if err != nil {
				http.Error(rw, "Ocurrio un error al obtener el ID de la direccion "+err.Error(), http.StatusInternalServerError)
				return
			}

			if addressModel.ID == 0 {
				http.Error(rw, "La direccion enviada no existe ", http.StatusBadRequest)
				return
			}

			zoneModel, _ := dbSetting.GetZoneById(addressModel.IDZone)

			var priceFactor float32 = 1.0
			for _, entry := range priceTable {
				if day.Amount >= entry.MinAmount && day.Amount <= entry.MaxAmount {
					priceFactor = entry.PriceFactor
					break
				}
			}

			zoneModel.Price *= priceFactor

			response.Delivery = response.Delivery + zoneModel.Price

		}
	}

	discounts, _ := dbSetting.GetDiscounts()

	sort.Slice(discounts, func(i, j int) bool {
		return discounts[i].Cant > discounts[j].Cant
	})

	response.SubTotal = price

	response.Discount = 0.0

	for i := range discounts {
		if amount >= discounts[i].Cant {
			response.Discount = price * (discounts[i].Percentage / 100)
			break
		} else if i < len(discounts)-1 && amount >= discounts[i+1].Cant && amount < discounts[i].Cant {
			response.Discount = price * (discounts[i+1].Percentage / 100)
			break
		}
	}

	response.Total = (response.SubTotal - response.Discount) + response.Delivery

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(response)

}
