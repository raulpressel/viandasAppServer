package handlers

import (
	"net/http"
	"strconv"
	dbClient "viandasApp/db/client"
	dbOrder "viandasApp/db/order"
	"viandasApp/handlers"
)

func UpdateDayOrderAddress(rw http.ResponseWriter, r *http.Request) {

	idDO := r.URL.Query().Get("idDayOrder")

	if len(idDO) < 1 {
		http.Error(rw, "El parametro idDayOrder es obligatorio", http.StatusBadRequest)
		return
	}

	idA := r.URL.Query().Get("idAddress")

	if len(idA) < 1 {
		http.Error(rw, "El parametro idAddress es obligatorio", http.StatusBadRequest)
		return
	}

	idDayOrder, err := strconv.Atoi(idDO)

	if err != nil {
		http.Error(rw, "Error al convertir el ID", http.StatusInternalServerError)
		return
	}

	idAddress, err := strconv.Atoi(idA)

	if err != nil {
		http.Error(rw, "Error al convertir el ID", http.StatusInternalServerError)
		return
	}

	usr := handlers.GetUser()

	if !usr.Admin {

		client, err := dbClient.GetClientByIDUser(usr.ID)
		if err != nil {
			http.Error(rw, "Cliente no encontrado", http.StatusBadRequest)
			return
		}

		dayOrderModel, err := dbOrder.GetDayOrderById(idDayOrder)
		if err != nil {
			http.Error(rw, "no existe la dayorder solicitada", http.StatusBadRequest)
			return
		}

		orderModel, err := dbOrder.GetModelOrderById(dayOrderModel.OrderID)
		if err != nil {
			http.Error(rw, "la direccion enviada no corresponde al cliente logueado", http.StatusBadRequest)
			return
		}

		if orderModel.ClientID != client.Client.ID {
			http.Error(rw, "la orden no pertenece al cliente logueado", http.StatusBadRequest)
			return
		}

		var bandAdd bool

		for _, valor := range client.Client.Address {
			if valor.ID == idAddress {
				bandAdd = true
			}

		}

		if !bandAdd {
			http.Error(rw, "la direccion enviada no corresponde al cliente logueado", http.StatusBadRequest)
			return
		}

		dayOrderModel.AddressID = idAddress

		status, err := dbOrder.UpdateDayOrderAddress(dayOrderModel)

		if err != nil {
			http.Error(rw, "Orden no encontrada", http.StatusInternalServerError)
			return
		}

		if !status {
			http.Error(rw, "no se ha logrado modificar la direccion de la orden", http.StatusInternalServerError)
			return
		}

	} else {
		dayOrderModel, err := dbOrder.GetDayOrderById(idDayOrder)
		if err != nil {
			http.Error(rw, "no existe la dayorder solicitada", http.StatusBadRequest)
			return
		}
		dayOrderModel.AddressID = idAddress

		status, err := dbOrder.UpdateDayOrderAddress(dayOrderModel)

		if err != nil {
			http.Error(rw, "Orden no encontrada", http.StatusInternalServerError)
			return
		}

		if !status {
			http.Error(rw, "no se ha logrado modificar la direccion de la orden", http.StatusInternalServerError)
			return
		}
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
