package handlers

import (
	"encoding/json"
	"net/http"
	db "viandasApp/db/carousel"
)

/*Obtenerbanner envia el banner al http*/

func GetBanners(rw http.ResponseWriter, r *http.Request) {

	bannerModel, err := db.GetBanners()

	if err != nil {
		http.Error(rw, "no se pudo recuperar los banners ", http.StatusInternalServerError)

	}
	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(bannerModel)

}
