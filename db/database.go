package db

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

type DataBase struct {
	Connector *gorm.DB
}

var db *gorm.DB

/* type ConnDB interface {
	GetDB() DataBase

} */

func GetDB() *gorm.DB {

	/* 	sqlDB, err := db.DB()
	   	if err != nil {
	   		return nil
	   	}
	   	defer sqlDB.Close()

	   	err = sqlDB.Ping()
	   	if err != nil {
	   		return nil
	   	} */

	return db
}

func ConnectDB(dsn string) (d *gorm.DB, err error) {

	d, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db = d

	fmt.Println("conexion existosa")

	return db, nil

}

func GetKeyDB(key string) (string, error) {
	keyDB := os.Getenv(key)
	if keyDB == "" {
		return keyDB, errors.New("missing key")
	}
	return keyDB, nil
}
