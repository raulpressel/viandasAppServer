package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetReportDeliveryDriver(id int) (models.DeliveryDriver, error) {

	db := db.GetDB()

	var deliveryDriverModel models.DeliveryDriver

	err := db.First(&deliveryDriverModel, id).Error

	return deliveryDriverModel, err

}


/* select *
from delivery_drivers
left join tandas on tandas.delivery_driver_id = delivery_drivers.id
left join tanda_addresses ON tanda_addresses.tanda_id = tandas.id
left join day_orders ON day_orders.address_id = tanda_addresses.address_id
left join day_menus ON day_orders.day_menu_id = day_menus.id
where  tandas.delivery_driver_id is not null
AND day_menus.date BETWEEN "2023-06-12" AND "2023-06-16" 
AND delivery_drivers.id = 1; */