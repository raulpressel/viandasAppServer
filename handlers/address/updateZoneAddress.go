package handlers

import (
	"encoding/json"
	"net/http"
	dbAddress "viandasApp/db/address"
	"viandasApp/dtos"
	"viandasApp/handlers"
	"viandasApp/models"
)

// UpdateZoneAddress recalcula la zona de todas las direcciones activas a partir de
// su lat/lng con los poligonos vigentes. Con dryRun=true solo informa los cambios.
func UpdateZoneAddress(rw http.ResponseWriter, r *http.Request) {

	dryRun := r.URL.Query().Get("dryRun") == "true"

	addressesModel, err := dbAddress.GetAllActiveAddresses()

	if err != nil {
		http.Error(rw, "Ocurrio un error al obtener las direcciones "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := dtos.UpdateZoneAddressResponse{
		DryRun:      dryRun,
		Total:       len(addressesModel),
		Changed:     []dtos.ZoneAddressChange{},
		WithoutZone: []dtos.ZoneAddressChange{},
	}

	var changedModel []models.Address

	for _, address := range addressesModel {

		change := dtos.ZoneAddressChange{
			ID:        address.ID,
			Street:    address.Street,
			Number:    address.Number,
			OldIDZone: address.IDZone,
		}

		idZone, ok := handlers.GetZoneByLatLng(address.Lat, address.Lng)

		if !ok {
			change.NewIDZone = address.IDZone
			response.WithoutZone = append(response.WithoutZone, change)
			continue
		}

		if idZone != address.IDZone {
			change.NewIDZone = idZone
			response.Changed = append(response.Changed, change)

			address.IDZone = idZone
			changedModel = append(changedModel, address)
		}
	}

	if !dryRun && len(changedModel) > 0 {

		status, err := dbAddress.UpdateAddressesZone(changedModel)

		if err != nil {
			http.Error(rw, "Ocurrio un error al actualizar las zonas de las direcciones "+err.Error(), http.StatusInternalServerError)
			return
		}

		if !status {
			http.Error(rw, "no se ha logrado actualizar las zonas de las direcciones en la BD", http.StatusInternalServerError)
			return
		}
	}

	rw.Header().Set("Content-type", "application/json")
	if dryRun {
		rw.WriteHeader(http.StatusOK)
	} else {
		rw.WriteHeader(http.StatusCreated)
	}
	json.NewEncoder(rw).Encode(response)

}
