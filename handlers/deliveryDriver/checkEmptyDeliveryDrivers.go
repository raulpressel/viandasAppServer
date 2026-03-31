package handlers

import (
	"encoding/json"
	"net/http"
	deliDriver "viandasApp/db/deliveryDriver"
	dbTanda "viandasApp/db/tanda"
)

type responseEmptyDrivers struct {
	Status bool  `json:"status"`
	Error  error `json:"error"`
}

func CheckEmptyDeliveryDrivers(rw http.ResponseWriter, r *http.Request) {

	var response responseEmptyDrivers

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

		response.Status, response.Error = deliDriver.UpdateDriversInDeliveries(deliveriesModel)

	}

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(response)

}
