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
func UpdateBanner(w http.ResponseWriter, r *http.Request) {

	var locationModel models.LocationImg

	var bannerModel models.Banner

	bannerModel.ID, _ = strconv.Atoi(r.FormValue("id"))

	bannerModel, _ = carouseldb.GetBannerById(bannerModel.ID)

	w.Header().Add("content-type", "application/json")

	file, handle, err := r.FormFile("banner")

	switch err {
	case nil:
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

		locationModel.Location = handlers.GetHash(locationModel.Location)
		file.Close()
	case http.ErrMissingFile:
		locationModel, _ = imgdb.GetLocationImgById(bannerModel.LocationID)
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

	locationModel.ID = bannerModel.LocationID

	status, err := carouseldb.UpdateBanner(bannerModel, locationModel)
	if err != nil {
		http.Error(w, "No se pudo guardar el mensaje en la base de datos "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "no se ha logrado insertar el registro  // status = false ", 400)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
