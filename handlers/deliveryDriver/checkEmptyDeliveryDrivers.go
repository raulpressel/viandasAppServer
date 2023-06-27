package handlers

import (
	"net/http"
	deliDriver "viandasApp/db/deliveryDriver"
	dbTanda "viandasApp/db/tanda"
)

func CheckEmptyDeliveryDrivers(rw http.ResponseWriter, r *http.Request) {

	deliveriesModel, err := deliDriver.CheckEmptyDeliveryDrivers()

	if err != nil {
		http.Error(rw, "Error al procesar los Deliviers vacios "+err.Error(), http.StatusInternalServerError)
		return
	}

	if len(deliveriesModel) > 0 {

		for i := range deliveriesModel {

			idTanda, err := dbTanda.CheckExistTandaByAddressId(deliveriesModel[i].AddressID)

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

				deliveriesModel[i].DeliveryDriverID = &tempID
			} else {
				deliveriesModel[i].DeliveryDriverID = nil
			}

		}

		status, err := deliDriver.UpdateDriversInDeliveries(deliveriesModel)

		if err != nil {
			http.Error(rw, "Ocurrio un error al actualizar los deliveries "+err.Error(), http.StatusInternalServerError)
			return
		}

		if !status {
			http.Error(rw, "no se ha logrado actualizar los deliveries ", http.StatusInternalServerError)
			return
		}

	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
