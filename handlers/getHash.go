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

	var extension = strings.Split(path, ".")[1] //saco la extension del archivo de imagen

	dst := "uploads/banners/" + hex.EncodeToString(h.Sum(nil)) + "." + extension

	//path = "C:/Users/Raul/Documents/github.com/raulpressel/viandasAppServer/" + path

	//os.Rename(path, dst)

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

	return dst

}
