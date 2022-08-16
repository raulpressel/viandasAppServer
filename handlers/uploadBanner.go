package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"viandasApp/db"
	"viandasApp/models"
)

/*subir el avatar al servidor*/
func UploadBanner(w http.ResponseWriter, r *http.Request) {

	var locationModel models.LocationImg

	w.Header().Add("content-type", "application/json")

	file, handle, err2 := r.FormFile("banner")
	if err2 != nil {
		return
	}

	locationModel.Location = "uploads/banners/" + handle.Filename

	f, err := os.OpenFile(locationModel.Location, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "error al subir banner "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)

	if err != nil {
		http.Error(w, "error al copiar  banner "+err.Error(), http.StatusBadRequest)
		return
	}

	var bannerModel models.Banner

	bannerModel.Title = r.FormValue("titulo")
	bannerModel.DateStart, err = time.Parse("Mon, 02 Jan 2006 15:04:05 MST", r.FormValue("fechaDesde"))
	if err != nil {
		fmt.Println(err)
		return
	}
	bannerModel.DateEnd, err = time.Parse("Mon, 02 Jan 2006 15:04:05 MST", r.FormValue("fechaHasta"))
	if err != nil {
		fmt.Println(err)
		return
	}
	bannerModel.Status = true

	db.ExistTable(bannerModel)
	db.ExistTable(locationModel)

	status, err := db.UploadBanner(bannerModel, locationModel)
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
