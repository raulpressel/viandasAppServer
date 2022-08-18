package db

import (
	"fmt"
	"time"
	"viandasApp/dtos"
	"viandasApp/models"
)

func GetBanners(onlyActive bool) []dtos.BannersResponse {
	var db = ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	locationModel := models.LocationImg{}
	responseModel := []dtos.BannersResponse{}

	var dateTime time.Time = time.Now()

	if onlyActive {

		err := db.Model(&locationModel).
			Select("banners.id, banners.title, banners.date_start, banners.date_end, location_imgs.location").
			Joins("JOIN banners ON banners.location_id = location_imgs.id").
			Where("? BETWEEN banners.date_start AND banners.date_end", dateTime).
			Scan(&responseModel).Error

		if err != nil {
			fmt.Println(err)
			return responseModel
		}

	} else {
		err := db.Table("location_imgs").
			Select("banners.id, banners.title, banners.date_start, banners.date_end, location_imgs.location").
			Joins("JOIN banners ON banners.location_id = location_imgs.id").
			Scan(&responseModel).Error

		if err != nil {
			fmt.Println(err)
			return responseModel
		}
	}

	return responseModel

}
