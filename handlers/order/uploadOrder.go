package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	dbAddress "viandasApp/db/address"
	dbCategories "viandasApp/db/categories"
	dbClient "viandasApp/db/client"
	dbMenu "viandasApp/db/menu"
	dbOrder "viandasApp/db/order"
	"viandasApp/dtos"
	"viandasApp/models"
)

type response struct {
	OrderId      int                  `json:"idOrder"`
	Total        float32              `json:"total"`
	CantDelivery int                  `json:"cantDelivery"`
	Categories   []dtos.CategoryTable `json:"categories"`
}
type categoriesCant struct {
	Cant        int    `json:"cant"`
	Description string `json:"description"`
}

func UploadOrder(rw http.ResponseWriter, r *http.Request) {

	var orderDto dtos.OrderRequest

	var orderModel models.Order

	var dOrderModel models.DayOrder

	var dayOrderModel []models.DayOrder

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

	orderModel.Total = orderDto.Total

	orderModel.StatusOrderID = 1 //se da de alta orden y queda con estado 1 - Activa

	orderModel.Paid = false

	orderModel.OrderDate, err = time.Parse(time.RFC3339, orderDto.Date)
	if err != nil {
		http.Error(rw, "Error en el formato de fecha recibido "+err.Error(), http.StatusBadRequest)
		return
	}

	for _, day := range orderDto.DaysOrderRequest {

		if day.Amount > 0 {

			dOrderModel.Amount = day.Amount

			dayMenuModel, err := dbMenu.GetDayMenuById(day.IDDayFood)

			if err != nil {
				http.Error(rw, "Ocurrio un error al obtener el ID del menu "+err.Error(), http.StatusInternalServerError)
				return
			}

			if dayMenuModel.ID == 0 {
				http.Error(rw, "El Day Menu enviado no existe ", http.StatusBadRequest)
				return
			}

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

			dOrderModel.AddressID = day.IDAddress

			dayOrderModel = append(dayOrderModel, dOrderModel)

		}
	}

	status, err, orderId := dbOrder.UploadOrder(orderModel, dayOrderModel)

	if err != nil {
		http.Error(rw, "Ocurrio un error al intentar ingresar el pedido "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "no se ha logrado cargar el pedido en la BD", http.StatusInternalServerError)
		return
	}

	categoriesCant, cantDelivery := calcAmounts(dayOrderModel)

	res := response{
		OrderId:      orderId.IDOrder,
		Total:        orderModel.Total,
		CantDelivery: cantDelivery,
		Categories:   *categoriesCant,
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(res)

}

func calcAmounts(dayOrderModel []models.DayOrder) (*[]dtos.CategoryTable, int) {

	var arr []int
	var cantEnvios int

	var categories []dtos.CategoryTable

	for _, day := range dayOrderModel {
		if day.AddressID != 100 {
			cantEnvios++
		}

		dayMenu, err := dbMenu.GetDayMenuById(day.DayMenuID)

		if err != nil {
			return nil, cantEnvios
		}

		arr = append(arr, dayMenu.CategoryID)

	}

	counts := countOccurrences(arr)

	for num, count := range counts {
		categoryModel, err := dbCategories.GetCategoryById(num)
		if err != nil {
			return nil, cantEnvios
		}
		category := dtos.CategoryTable{
			Cant: count,
			Category: dtos.CategoryResponse{
				ID:    categoryModel.ID,
				Title: categoryModel.Title,
			},
		}
		categories = append(categories, category)
	}

	return &categories, cantEnvios

}

func countOccurrences(arr []int) map[int]int {
	counts := make(map[int]int)
	for _, num := range arr {
		counts[num]++
	}
	return counts
}
