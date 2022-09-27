package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	db "viandasApp/db/menu"
	"viandasApp/dtos"
)

/*Obtenerbanner envia el banner al http*/

func GetMenuByCategory(rw http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("idCategory")

	if len(ID) < 1 {
		http.Error(rw, "debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	idCategory, _ := strconv.Atoi(ID)

	responseMenuFood := dtos.MenuViewer{}

	responseMenuFood, err := db.GetMenuByCategory(idCategory)

	if err != nil {
		http.Error(rw, "Menu no encontrado", http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(responseMenuFood)
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
