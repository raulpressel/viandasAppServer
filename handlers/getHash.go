package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func GetHash(path string) string {

	h := md5.New()
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	//defer f.Close()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	var path_aux = strings.Split(path, "/var/www/default/htdocs/")[1]

	var folder = strings.Split(path_aux, "/")[1]

	path_aux = path_aux[(len(path_aux) - 5):]

	var extension = strings.Split(path_aux, ".")[1] //saco la extension del archivo de imagen

	dst := "/var/www/default/htdocs/public/" + folder + "/" + hex.EncodeToString(h.Sum(nil)) + "." + extension

	// read original file
	origFile, _ := os.ReadFile(path)

	// create new file with a different name
	newFile, _ := os.Create(dst)

	// print data from original file to new file.
	fmt.Fprintf(newFile, "%s", string(origFile))

	err = os.Remove(path)
	if err != nil {
		fmt.Println("no se pudo borrar la imagen " + err.Error())
	}
	dest := strings.Split(dst, "/var/www/default/htdocs")[1]

	return dest

}
