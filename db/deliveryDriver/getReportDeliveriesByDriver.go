package db

import (
	"time"
	"viandasApp/db"
	"viandasApp/models"
)

func GetReportDeliveryDriver(id int, dateStart time.Time, dateEnd time.Time) (*[]models.Delivery, error) {

	db := db.GetDB()

	deliveryModel := []models.Delivery{}

	query := db.Model(&deliveryModel)

	query = query.Where("date(deliveries.delivery_date) BETWEEN ? AND ?", dateStart.Format("2006-01-02"), dateEnd.Format("2006-01-02"))

	query = query.Where("address_id <> 100")

	if id != 0 {
		query = query.Where("delivery_driver_id = ?", id)
	}

	query = query.Order("deliveries.order_id asc, deliveries.delivery_date asc")

	if err := query.Find(&deliveryModel).Error; err != nil {
		return nil, db.Error
	}

	return &deliveryModel, nil

}
