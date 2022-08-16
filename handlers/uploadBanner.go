package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
	"viandasApp/db"
	"viandasApp/models"
)

/*subir el avatar al servidor*/
func UploadBanner(w http.ResponseWriter, r *http.Request) {

	var locationModel models.LocationImg

	w.Header().Add("content-type", "application/json")

	file, handle, err := r.FormFile("banner")
	if err != nil {
		return
	}

	defer file.Close()

	filename := md5.New()

	_, err = io.Copy(filename, file) //funcion que copia el hash de file a la variable filename
	if err != nil {
		panic(err)
	}

	var extension = strings.Split(handle.Filename, ".")[1] //saco la extension del archivo de imagen

	hash := filename.Sum(nil) // guardo el valor de hash md5 en la varialbe hash

	locationModel.Location = "uploads/banners/" + hex.EncodeToString(hash[:]) + "." + extension //la ubicacion esta compuesta por el la ruta + el hash convertido a string + la extension del archivo

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
