package handlers

import (
	"errors"
	"net/http"
	"strconv"
	dbAddress "viandasApp/db/address"
	dbClient "viandasApp/db/client"
	dbDelivery "viandasApp/db/deliveryDriver"
	dbMenu "viandasApp/db/menu"
	dbOrder "viandasApp/db/order"
	dbSetting "viandasApp/db/setting"
	dbTanda "viandasApp/db/tanda"
	"viandasApp/handlers"
	"viandasApp/models"

	"gorm.io/gorm"
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

	dayOrderModel, err := dbOrder.GetDayOrderById(idDayOrder)
	if err != nil {
		http.Error(rw, "no existe la dayorder solicitada", http.StatusBadRequest)
		return
	}

	orderModel, err := dbOrder.GetModelOrderById(dayOrderModel.OrderID)
	if err != nil {
		http.Error(rw, "Orden no encontrada", http.StatusBadRequest)
		return
	}

	usr := handlers.GetUser()

	if !usr.Admin {

		client, err := dbClient.GetClientByIDUser(usr.ID)
		if err != nil {
			http.Error(rw, "Cliente no encontrado", http.StatusBadRequest)
			return
		}

		if orderModel.ClientID != client.Client.ID {
			http.Error(rw, "la orden no pertenece al cliente logueado", http.StatusBadRequest)
			return
		}

		// el retiro en local (direccion 100) no figura entre las direcciones del cliente
		bandAdd := idAddress == 100

		for _, valor := range client.Client.Address {
			if valor.ID == idAddress {
				bandAdd = true
			}

		}

		if !bandAdd {
			http.Error(rw, "la direccion enviada no corresponde al cliente logueado", http.StatusBadRequest)
			return
		}
	}

	addressModel, err := dbAddress.GetAddressById(idAddress)

	if err != nil || addressModel.ID == 0 {
		http.Error(rw, "La direccion enviada no existe ", http.StatusBadRequest)
		return
	}

	dayMenuModel, err := dbMenu.GetDayMenuById(dayOrderModel.DayMenuID)
	if err != nil {
		http.Error(rw, "Ocurrio un error al obtener el menu del dia "+err.Error(), http.StatusInternalServerError)
		return
	}

	// hay un solo envio por fecha, asi que todos los platos del dia van a la misma direccion
	dayOrdersModel, err := dbOrder.GetDayOrdersByOrderIdAndDay(orderModel.ID, dayMenuModel.Date)
	if err != nil {
		http.Error(rw, "Ocurrio un error al obtener los platos del dia "+err.Error(), http.StatusInternalServerError)
		return
	}

	if len(dayOrdersModel) == 0 {
		dayOrdersModel = append(dayOrdersModel, dayOrderModel)
	}

	for i := range dayOrdersModel {
		dayOrdersModel[i].AddressID = idAddress
	}

	var deliveryModel *models.Delivery

	modelDelivery, err := dbDelivery.GetDeliveryByOrderIDandDate(orderModel.ID, dayMenuModel.Date)

	if err == nil {

		var deliveryPrice float32 = 0

		if idAddress != 100 {
			zoneModel, err := dbSetting.GetZoneById(addressModel.IDZone)
			if err != nil {
				http.Error(rw, "La direccion enviada no tiene una zona de envio valida", http.StatusBadRequest)
				return
			}
			deliveryPrice = zoneModel.Price
		}

		orderModel.Total = orderModel.Total - modelDelivery.DeliveryPrice + deliveryPrice

		modelDelivery.AddressID = idAddress
		modelDelivery.DeliveryPrice = deliveryPrice

		idTanda, err := dbTanda.CheckExistTandaByAddressId(idAddress)
		if err != nil {
			http.Error(rw, "Ocurrio un error al obtener el ID de la Tanda "+err.Error(), http.StatusInternalServerError)
			return
		}

		if idTanda > 0 {
			deliveryDriverId, err := dbTanda.GetDeliveryDriverIdByTandaId(idTanda)
			if err != nil {
				http.Error(rw, "Ocurrio un error al obtener el ID del Cadete "+err.Error(), http.StatusInternalServerError)
				return
			}

			tempID := uint(deliveryDriverId)

			modelDelivery.DeliveryDriverID = &tempID
		} else {
			modelDelivery.DeliveryDriverID = nil
		}

		deliveryModel = &modelDelivery

	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(rw, "Ocurrio un error al obtener el envio del dia "+err.Error(), http.StatusInternalServerError)
		return
	}

	status, err := dbOrder.UpdateDayOrderAddress(dayOrdersModel, deliveryModel, orderModel)

	if err != nil {
		http.Error(rw, "Orden no encontrada", http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "no se ha logrado modificar la direccion de la orden", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
