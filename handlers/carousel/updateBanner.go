package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	carouseldb "viandasApp/db/carousel"
	imgdb "viandasApp/db/img"
	"viandasApp/handlers"
	"viandasApp/models"
)

/*subir el avatar al servidor*/
func UpdateBanner(rw http.ResponseWriter, r *http.Request) {

	var locationModel models.LocationImg

	var bannerModel models.Banner

	_ID, err := strconv.Atoi(r.FormValue("id"))
	if _ID < 1 {
		http.Error(rw, "debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(rw, "Error al convertir el ID", http.StatusInternalServerError)
		return
	}

	bannerModel, err = carouseldb.GetBannerById(_ID)
	if err != nil {
		http.Error(rw, "no fue posible recuperar el banner por ID", http.StatusInternalServerError)
		return
	}

	if bannerModel.LocationID != nil {

		locationModel, err = imgdb.GetLocationImgById(*bannerModel.LocationID)
		if err != nil {
			http.Error(rw, "no fue posible recuperar la imagen por ID", http.StatusInternalServerError)
			return
		}
	}

	rw.Header().Add("content-type", "application/json")

	file, handle, err := r.FormFile("banner")

	switch err {
	case nil:
		locationModel.Location = "/var/www/default/htdocs/public/banners/" + handle.Filename

		f, err := os.OpenFile(locationModel.Location, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			http.Error(rw, "error al subir banner "+err.Error(), http.StatusBadRequest)
			return
		}

		_, err = io.Copy(f, file)

		if err != nil {
			http.Error(rw, "error al copiar  banner "+err.Error(), http.StatusBadRequest)
			return
		}

		locationModel.Location = handlers.GetHash(locationModel.Location)

		if bannerModel.LocationID != nil {

			locationModel.ID = *bannerModel.LocationID
		} else {
			numero := 0
			bannerModel.LocationID = &numero
		}
		file.Close()
	case http.ErrMissingFile:
		if locationModel.Location != "" {

			bannerModel.LocationID = nil

		}
	default:
		log.Println(err)
	}

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

	status, err := carouseldb.UpdateBanner(bannerModel, locationModel)
	if err != nil {
		http.Error(rw, "No se pudo guardar el mensaje en la base de datos "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(rw, "no se ha logrado insertar el registro  // status = false ", 400)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
