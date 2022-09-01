package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"viandasApp/db"
	carouseldb "viandasApp/db/carousel"
	"viandasApp/handlers"
	"viandasApp/models"
)

/*subir el avatar al servidor*/
func UploadBanner(w http.ResponseWriter, r *http.Request) {

	if err := os.MkdirAll("/var/www/default/htdocs/public/banners", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	var locationModel models.LocationImg

	w.Header().Add("content-type", "application/json")

	file, handle, err := r.FormFile("banner")
	if err != nil {
		return
	}

	locationModel.Location = "/var/www/default/htdocs/public/banners/" + handle.Filename

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

	defer file.Close()

	locationModel.Location = handlers.GetHash(locationModel.Location)

	var bannerModel models.Banner

	bannerModel.ID, _ = strconv.Atoi(r.FormValue("id"))

	bannerModel.Title = r.FormValue("title")
	bannerModel.DateStart, err = time.Parse("Mon, 02 Jan 2006 15:04:05 MST", r.FormValue("dateStart"))
	if err != nil {
		fmt.Println(err)
		return
	}
	bannerModel.DateEnd, err = time.Parse("Mon, 02 Jan 2006 15:04:05 MST", r.FormValue("dateEnd"))
	if err != nil {
		fmt.Println(err)
		return
	}
	bannerModel.Active = true

	db.ExistTable(bannerModel)
	db.ExistTable(locationModel)

	status, err := carouseldb.UploadBanner(bannerModel, locationModel)
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
