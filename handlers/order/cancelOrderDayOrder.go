package handlers

import (
	"net/http"
	"strconv"
	dbAddress "viandasApp/db/address"
	dbCategory "viandasApp/db/categories"
	dbDelivery "viandasApp/db/deliveryDriver"
	dbMenu "viandasApp/db/menu"
	dbSetting "viandasApp/db/setting"

	db "viandasApp/db/order"
)

func CancelOrderDayOrder(rw http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("idDayOrder")

	if len(ID) < 1 {
		http.Error(rw, "debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	idOrder, err := strconv.Atoi(ID)

	if err != nil {
		http.Error(rw, "Error al convertir el ID", http.StatusInternalServerError)
		return
	}

	ca := r.URL.Query().Get("cant")

	cant, err := strconv.Atoi(ca)

	if err != nil {
		http.Error(rw, "Error al convertir la cantidad", http.StatusInternalServerError)
		return
	}

	modelDayOrder, err := db.GetDayOrderById(idOrder)

	if err != nil {
		http.Error(rw, "Day Orden no encontrado con el ID solicitado", http.StatusBadRequest)
		return
	}

	modelDayMenu, err := dbMenu.GetDayMenuById(modelDayOrder.DayMenuID)

	if err != nil {
		http.Error(rw, "Day Menu no encontrado con el ID solicitado", http.StatusBadRequest)
		return
	}

	modelDayOrder.Status = false

	modelDelivery, err := dbDelivery.GetDeliveryByOrderIDandDate(modelDayOrder.OrderID, modelDayMenu.Date)

	delete := true

	if modelDelivery.DeliveryMenuAmount > cant {

		modelDelivery.DeliveryMenuAmount = modelDelivery.DeliveryMenuAmount - cant

		priceFactor := PriceFactor(modelDelivery.DeliveryMenuAmount)

		modelCategory, err := dbCategory.GetCategoryById(modelDayMenu.CategoryID)
		if err != nil {
			http.Error(rw, "Category no encontrada con el ID solicitado", http.StatusBadRequest)
			return
		}

		modelDelivery.DeliveryMenuPrice = modelDelivery.DeliveryMenuPrice - (modelCategory.Price * float32(cant))

		addressModel, err := dbAddress.GetAddressById(modelDelivery.AddressID)

		if err != nil {
			http.Error(rw, "Adress no encontrada con el ID solicitado", http.StatusBadRequest)
			return
		}

		zoneModel, err := dbSetting.GetZoneById(addressModel.IDZone)
		if err != nil {
			http.Error(rw, "Zone no encontrada con el ID solicitado", http.StatusBadRequest)
			return
		}

		modelDelivery.DeliveryPrice = zoneModel.Price * priceFactor

		delete = false

	}

	status, err := db.CancelDayOrder(modelDayOrder, modelDelivery, delete)

	if err != nil {
		http.Error(rw, "error al actualizar el estado del DayOrder", http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "no se pudo cancelar la orden", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
