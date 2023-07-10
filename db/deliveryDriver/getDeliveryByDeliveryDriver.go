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

func GetDeliveryByDeliveryDriver(id int, dateStart time.Time, dateEnd time.Time) (*dtos.DeliveryResponse, error) {

	db := db.GetDB()

	deliveryModel := []models.Delivery{}

	query := db.Model(&deliveryModel)

	query = query.Where("date(deliveries.delivery_date) BETWEEN ? AND ?", dateStart.Format("2006-01-02"), dateEnd.Format("2006-01-02"))

	query = query.Where("address_id <> 100")

	if id != 0 {
		query = query.Where("delivery_driver_id = ?", id)
	}

	if err := query.Find(&deliveryModel).Error; err != nil {
		return nil, db.Error
	}

	occurrences := countOccurrences(deliveryModel)

	var deliveryRes []dtos.DeliveryRes

	var deliRes dtos.DeliveryRes

	for num := range occurrences {

		deliveryDriverModel, err := GetDeliveryDriverByID(num)
		if err != nil {
			return nil, db.Error
		}

		deliRes.DeliveryDriver = dtos.DeliveryByDeliveryDriver{

			ID:       deliveryDriverModel.ID,
			DNI:      deliveryDriverModel.DNI,
			Name:     deliveryDriverModel.Name,
			LastName: deliveryDriverModel.LastName,
			Phone:    deliveryDriverModel.Phone,
		}

		for _, delivery := range deliveryModel {

			if delivery.DeliveryDriverID != nil && int(*delivery.DeliveryDriverID) == num {

				var modelOrder models.Order

				db.First(&modelOrder, delivery.OrderID)

				clientModel, err := dbClient.GetClientById(modelOrder.ClientID)
				if err != nil {
					return nil, db.Error
				}

				addressModel, err := dbAdd.GetAddressById(delivery.AddressID)
				if err != nil {
					return nil, db.Error
				}

				cityModel, err := dbCity.GetCityById(addressModel.CityID)
				if err != nil {
					return nil, db.Error
				}

				deli := dtos.Delivery{
					Deli: dtos.Deli{
						ID:      delivery.ID,
						IdOrder: delivery.OrderID,
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

				deliRes.DeliveryDriver.Delivery = append(deliRes.DeliveryDriver.Delivery, deli)

			}
		}

		deliveryRes = append(deliveryRes, deliRes)

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
