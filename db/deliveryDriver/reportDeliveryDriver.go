package db

import (
	"time"
	"viandasApp/db"
	dbAdd "viandasApp/db/address"
	dbCity "viandasApp/db/city"
	dbClient "viandasApp/db/client"
	"viandasApp/dtos"
	"viandasApp/models"
)

func GetReportDeliveryDriver(id int, dateStart time.Time, dateEnd time.Time) (*dtos.DeliveryResponse, error) {

	db := db.GetDB()

	deliveryModel := []models.Delivery{}

	query := db.Model(&deliveryModel)

	query = query.Where("deliveries.delivery_date BETWEEN ? AND ?", dateStart.Format("2006-01-02"), dateEnd.Format("2006-01-02"))

	query = query.Where("address_id <> 100")

	if id != 0 {
		query = query.Where("delivery_driver_id = ?", id)
	}

	if err := query.Find(&deliveryModel).Error; err != nil {
		return nil, db.Error
	}

	occurrences := countOccurrences(deliveryModel)

	var deliveryRes []dtos.DeliveryRes

	for num := range occurrences {

		deliveryDriverModel, _ := GetDeliveryDriverByID(num)

		deliveryDriver := dtos.DeliveryDriverRes{

			ID:       deliveryDriverModel.ID,
			DNI:      deliveryDriverModel.DNI,
			Name:     deliveryDriverModel.Name,
			LastName: deliveryDriverModel.LastName,
			Phone:    deliveryDriverModel.Phone,
		}

		var deliveries dtos.DeliveryRes

		for _, delivery := range deliveryModel {

			if delivery.DeliveryDriverID != nil && int(*delivery.DeliveryDriverID) == num {

				//modelOrder, _ := .GetModelOrderById(delivery.OrderID)

				var modelOrder models.Order

				db.First(&modelOrder, delivery.OrderID)

				clientModel, _ := dbClient.GetClientById(modelOrder.ClientID)

				addressModel, _ := dbAdd.GetAddressById(delivery.AddressID)

				cityModel, _ := dbCity.GetCityById(addressModel.CityID)

				deli := dtos.Delivery{
					Deli: dtos.Deli{
						ID:      delivery.ID,
						IdOrden: delivery.OrderID,
						Price:   delivery.DeliveryPrice,
						Date:    delivery.DeliveryDate,
						Client: dtos.Client{
							ID:       clientModel.ID,
							Name:     clientModel.Name,
							LastName: clientModel.LastName,
						},
						Address: dtos.AddressRespone{
							ID:          addressModel.ID,
							Street:      addressModel.Street,
							Number:      addressModel.Number,
							Floor:       addressModel.Floor,
							Departament: addressModel.Departament,
							IDZone:      addressModel.IDZone,
							Favourite:   addressModel.Favourite,
							Lat:         addressModel.Lat,
							Lng:         addressModel.Lng,
							Observation: addressModel.Observation,
							City: dtos.AllCityResponse{
								ID:          cityModel.ID,
								Description: cityModel.Description,
								CP:          cityModel.CP,
							},
						},
					},
				}

				deliveries.Delivery = append(deliveries.Delivery, deli)
			}
		}

		deliveries.DeliveryDriverRes = deliveryDriver

		deliveryRes = append(deliveryRes, deliveries)

	}

	var deliveryResponse dtos.DeliveryResponse

	deliveryResponse.DeliveryRes = deliveryRes

	return &deliveryResponse, nil

}

func countOccurrences(arr []models.Delivery) map[int]int {
	occurrences := make(map[int]int)

	for _, num := range arr {
		if num.DeliveryDriverID != nil {
			occurrences[int(*num.DeliveryDriverID)]++
		}
	}

	return occurrences
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
