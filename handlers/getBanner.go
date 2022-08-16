package handlers

//	"github.com/raulpressel/twittor/bd"

/*Obtenerbanner envia el banner al http*/

/* func getBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "debe enviar el parametro ID", http.StatusBadRequest)
		return
	}



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
	}

} */
