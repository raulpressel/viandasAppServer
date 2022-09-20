package handlers

import (
	"encoding/json"
	"net/http"
	db "viandasApp/db/food"
	"viandasApp/dtos"
)

/*Obtenerbanner envia el banner al http*/

func GetAllFood(rw http.ResponseWriter, r *http.Request) {

	responseModelFood := []dtos.AllFoodResponse{}

	responseModelFood, err := db.GetAllFood()

	if err != nil {
		http.Error(rw, "no fue posible recuperar los platos", http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(responseModelFood)
	/* output, _ := json.Marshal(db.GetBanners(onlyActiveModel.OnlyActive))
	fmt.Fprintln(rw, string(output)) */

	/*
		perfil, err := bd.BuscoPerfil(ID)
		if err != nil {
			http.Error(w, "usuario no encontrado", http.StatusBadRequest)
			return
		}

		OpenFile, err := os.Open("uploads/banners" + perfil.Banner)
		if err != nil {
			http.Error(w, "banner no encontrado", http.StatusBadRequest)
			return
		}

		_, err = io.Copy(w, OpenFile)

		if err != nil {
			http.Error(w, "error al copiar banner", http.StatusBadRequest)
		} */

}
