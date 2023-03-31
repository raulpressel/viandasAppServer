package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"viandasApp/db"
	dbAddress "viandasApp/db/address"
	dbClient "viandasApp/db/client"
	dbMenu "viandasApp/db/menu"
	dbOrder "viandasApp/db/order"
	"viandasApp/dtos"
	"viandasApp/models"
)

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

	db.ExistTable(orderModel)

	db.ExistTable(dOrderModel)

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

	orderModel.Status = true

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

			dOrderModel.Active = true

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

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(orderId)

}
