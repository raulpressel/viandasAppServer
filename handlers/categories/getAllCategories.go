package handlers

import (
	"encoding/json"
	"net/http"
	db "viandasApp/db/categories"
	"viandasApp/dtos"
)

/*Obtenerbanner envia el banner al http*/

func GetAllCategories(rw http.ResponseWriter, r *http.Request) {

	var responseModel []dtos.AllCategoryResponse

	responseModel, err := db.GetAllCategory()

	if err != nil {
		http.Error(rw, "no fue posible recuperar las categorias", http.StatusBadRequest)
		return
	}
	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(responseModel)

}
