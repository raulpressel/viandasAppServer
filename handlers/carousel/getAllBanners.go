package handlers

import (
	"encoding/json"
	"net/http"

	db "viandasApp/db/carousel"
	"viandasApp/dtos"
)

/*Obtenerbanner envia el banner al http*/

func GetAllBanners(rw http.ResponseWriter, r *http.Request) {

	var bannerModel []dtos.AllBannersResponse

	bannerModel, err := db.GetAllBanners()

	if err != nil {
		http.Error(rw, "no se pudo recuperar los banners", http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(bannerModel)

}
