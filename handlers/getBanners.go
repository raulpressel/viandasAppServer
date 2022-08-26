package handlers

import (
	"encoding/json"
	"net/http"
	"viandasApp/db"
)

/*Obtenerbanner envia el banner al http*/

func GetBanners(rw http.ResponseWriter, r *http.Request) {



	/* 	err := json.NewDecoder(r.Body).Decode(&onlyActiveModel) //body es un obj string de solo lectura, una vez q se utiliza body se destruye
	   	if err != nil {
	   		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
	   		return
	   	} */

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(db.GetBanners())
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
