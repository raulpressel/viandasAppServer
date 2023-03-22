package handlers

import (
	"encoding/json"
	"net/http"
	"viandasApp/db"
	settingDB "viandasApp/db/setting"
	"viandasApp/dtos"
	"viandasApp/models"
)

func UploadDiscount(rw http.ResponseWriter, r *http.Request) {

	var discountDto dtos.DiscountRequest

	var discountModel models.Discount

	err := json.NewDecoder(r.Body).Decode(&discountDto)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	db.ExistTable(discountModel)

	discountModel.Description = discountDto.Discount.Description

	if discountDto.Discount.Cant < 1 {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	discountModel.Cant = discountDto.Discount.Cant

	if discountDto.Discount.Percentage < 1 {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}
	if discountDto.Discount.Percentage > 100 {
		http.Error(rw, "Error en los datos recibidos, no puede ser mayor a 100 "+err.Error(), http.StatusBadRequest)
		return
	}

	discountModel.Percentage = discountDto.Discount.Percentage

	discountModel.Active = true

	status, err := settingDB.UploadDiscount(discountModel)

	if err != nil {
		http.Error(rw, "Ocurrio un error al intentar registrar el descuento "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "no se ha logrado registrar el descuento en la BD", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
