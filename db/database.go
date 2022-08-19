package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//realiza la conexion
var dsn = "root:Aerolavelarata66@tcp(localhost:3306)/viandas_db?charset=utf8mb4&parseTime=True&loc=Local" //falta pass

func ConnectDB() *gorm.DB {
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("error en la conexion", err)
		panic(err)
	} else {
		fmt.Println("conexion existosa")
		return db
	}
}

func CheckConnection() int {
	var db = ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	err := sqlDB.Ping()
	if err != nil {
		return 0
	}
	return 1

}
