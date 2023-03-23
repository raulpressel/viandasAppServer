package handlers

import (
	"encoding/json"
	"net/http"
	settingDB "viandasApp/db/setting"
	"viandasApp/dtos"
)

func UpdateDiscount(rw http.ResponseWriter, r *http.Request) {

	var discountDto dtos.DiscountRequest

	err := json.NewDecoder(r.Body).Decode(&discountDto)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	discountModel, err := settingDB.GetDiscountById(discountDto.Discount.ID)

	if err != nil {
		http.Error(rw, "no fue posible recuperar el descuento de la BD", http.StatusInternalServerError)
		return
	}

	if discountModel.ID == 0 {
		http.Error(rw, "no fue posible recuperar el descuento de la BD", http.StatusBadRequest)
		return
	}

	discountModel.Description = discountDto.Discount.Description

	if discountDto.Discount.Cant < 1 {
		http.Error(rw, "Error en los datos recibidos, no puede ser menor a 1 "+err.Error(), http.StatusBadRequest)
		return
	}

	discountModel.Cant = discountDto.Discount.Cant

	if discountDto.Discount.Percentage < 1 {
		http.Error(rw, "Error en los datos recibidos, no puede ser menor a 1 "+err.Error(), http.StatusBadRequest)
		return
	}
	if discountDto.Discount.Percentage > 100 {
		http.Error(rw, "Error en los datos recibidos, no puede ser mayor a 100 "+err.Error(), http.StatusBadRequest)
		return
	}

	discountModel.Percentage = discountDto.Discount.Percentage

	status, err := settingDB.UploadDiscount(discountModel)

	if err != nil {
		http.Error(rw, "Ocurrio un error al intentar actualizar el descuento "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "no se ha logrado actualizar el descuento en la BD", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
