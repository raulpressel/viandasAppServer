package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	db "viandasApp/db/food"
)

func GetImageByCategory(rw http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("idCategory")

	if len(ID) < 1 {
		http.Error(rw, "debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	idCategory, err := strconv.Atoi(ID)

	if err != nil {
		http.Error(rw, "Error al convertir el ID", http.StatusInternalServerError)
		return
	}

	responseModelFood, err := db.GetImgByCategoryId(idCategory)

	if err != nil {
		http.Error(rw, "no fue posible recuperar las imagenes de los platos", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(responseModelFood)

}
