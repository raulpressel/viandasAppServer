package db

import (
	"time"
	"viandasApp/db"
	"viandasApp/models"
)

func GetDeliveryByOrderIDandDate(id int, date time.Time) (models.Delivery, error) {
	db := db.GetDB()

	var deliveryModel models.Delivery

	err := db.First(&deliveryModel, "order_id = ? and date(delivery_date) = ?", id, date.Format("2006-01-02")).Error

	/* query := db.Model(&deliveryModel)

	query = query.Where("date(deliveries.delivery_date) = ?", dateStart.Format("2006-01-02"), dateEnd.Format("2006-01-02"))

	query = query.Where("deliveries.order_id = ?")

	if err := query.Find(&deliveryModel).Error; err != nil {
		return deliveryModel, db.Error
	} */

	return deliveryModel, err

}
