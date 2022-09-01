package handlers

import (
	"net/http"
	"strconv"
	db "viandasApp/db/carousel"
	"viandasApp/models"
)

/*subir el avatar al servidor*/
func DeleteBanner(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("idBanner")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var bannerModel models.Banner

	_ID, err := strconv.Atoi(ID)

	bannerModel, err = db.GetBannerById(_ID)
	if err != nil {
		http.Error(w, "No se pudo obtener el banner "+err.Error(), 400)
		return
	}

	bannerModel.Active = false

	status, err := db.DeleteBanner(bannerModel)
	if err != nil {
		http.Error(w, "No se pudo guardar el mensaje en la base de datos "+err.Error(), 400)
		return
	}

	if !status { //esto es igual a !status == false
		http.Error(w, "no se ha logrado insertar el registro  // status = false ", 400)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
