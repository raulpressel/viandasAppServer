package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	db "viandasApp/db/client"
	"viandasApp/handlers"
)

type TokenKC struct {
	AccessToken string `json:"access_token"`
}

type KC struct {
	Path       string
	Pass       string
	Usr        string
	Client_ID  string
	Grant_Type string
}

func GetPathKC(key, pass, usr, client_id, grant_type string) (*KC, error) {

	var kc KC

	kc.Path = os.Getenv(key)
	if kc.Path == "" {
		return nil, errors.New("missing key")
	}
	kc.Pass = os.Getenv(pass)
	if kc.Pass == "" {
		return nil, errors.New("missing key")
	}
	kc.Usr = os.Getenv(usr)
	if kc.Usr == "" {
		return nil, errors.New("missing key")
	}
	kc.Client_ID = os.Getenv(client_id)
	if kc.Client_ID == "" {
		return nil, errors.New("missing key")
	}
	kc.Grant_Type = os.Getenv(grant_type)
	if kc.Grant_Type == "" {
		return nil, errors.New("missing key")
	}

	return &kc, nil
}

func DeleteClient(rw http.ResponseWriter, r *http.Request) {

	idUserKL := r.URL.Query().Get("idUser")

	if len(idUserKL) < 1 {
		http.Error(rw, "debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	usr := handlers.GetUser()

	if !usr.Admin {
		http.Error(rw, "No tienes los permisos para ver esta información", http.StatusBadRequest)
		return
	}

	client, valid := db.CheckExistClient(idUserKL)

	if !valid {
		http.Error(rw, "No existe un cliente con los datos solicitados ", http.StatusBadRequest)
		return
	}

	client.Active = false

	status, err := db.DeleteClient(client)

	if err != nil {
		http.Error(rw, "Ocurrio un error al eliminar el cliente "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "no se ha logrado eliminar el cliente de la BD", http.StatusInternalServerError)
		return
	}

	kc, err := GetPathKC("KC", "PASSKC", "USERKC", "CLIENT_ID", "GRANT_TYPE")

	if err != nil {
		log.Fatal("Key incorrecta PATH")
		return
	}

	token := getTokenAdminKC(*kc)

	if token == nil {
		http.Error(rw, "No tienes los permisos para ver esta información", http.StatusBadRequest)
	}

	req, err := http.NewRequest("DELETE", kc.Path+"auth/admin/realms/viandas/users/"+idUserKL, nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}

	defer resp.Body.Close()

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)

}

func getTokenAdminKC(kc KC) *TokenKC {

	var token TokenKC

	url := kc.Path + "auth/realms/master/protocol/openid-connect/token"
	method := "POST"

	payload := strings.NewReader("client_id=" + kc.Client_ID + "&username=" + kc.Usr + "&password=" + kc.Pass + "&grant_type=" + kc.Grant_Type)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {

		return nil
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {

		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {

		return nil
	}

	err = json.Unmarshal([]byte(body), &token)
	if err != nil {
		fmt.Println("Error al decodificar JSON:", err)
		return nil
	}

	return &token
}
