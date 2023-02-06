package db

import (
	"time"
	"viandasApp/db"
	dbAddress "viandasApp/db/address"
	dbCity "viandasApp/db/city"
	dbClient "viandasApp/db/client"
	dbDeliveryDriver "viandasApp/db/deliveryDriver"

	dbVehicle "viandasApp/db/vehicle"
	"viandasApp/dtos"
	"viandasApp/models"
)

func GetOrders(date time.Time) (*dtos.OrdersResponse, error) {

	db := db.GetDB()

	//var modelDayOrder []models.DayOrder

	modelTanda := []models.Tanda{}

	if err := db.Table("tandas").
		Select("tandas.id, tandas.description, tandas.hour_start, tandas.hour_end, tandas.delivery_driver_id").
		Where("tandas.active = 1").
		Where("exists (select tanda_addresses.id from tanda_addresses where tanda_addresses.tanda_id = tandas.id)").
		Where("exists (select id from day_orders where day_orders.address_id IN (select tanda_addresses.address_id from tanda_addresses where tanda_addresses.tanda_id = tandas.id))").
		Scan(&modelTanda).
		Error; err != nil {
		return nil, err
	}

	var tandasTable []dtos.Tanda

	for _, tanda := range modelTanda {

		deliveryDriverModel, err := dbDeliveryDriver.GetDeliveryDriverByID(tanda.DeliveryDriverID)
		if err != nil {
			return nil, err
		}

		addressModel, err := dbAddress.GetAddressById(deliveryDriverModel.AddressID)
		if err != nil {
			return nil, err
		}

		vehicleModel, err := dbVehicle.GetVehicleByID(deliveryDriverModel.VehicleID)
		if err != nil {
			return nil, err
		}

		cityModel, err := dbCity.GetCityById(addressModel.CityID)
		if err != nil {
			return nil, err
		}

		tandaRes := dtos.TandaRes{
			ID:          tanda.ID,
			Description: tanda.Description,
			HourStart:   tanda.HourStart,
			HourEnd:     tanda.HourEnd,
			DeliveryDriver: dtos.DeliveryDriverRes{
				ID:       deliveryDriverModel.ID,
				DNI:      deliveryDriverModel.DNI,
				Name:     deliveryDriverModel.Name,
				LastName: deliveryDriverModel.LastName,
				Phone:    deliveryDriverModel.Phone,
				BornDate: deliveryDriverModel.BornDate,
				Vehicle: dtos.Vehicle{
					ID:     vehicleModel.ID,
					Brand:  vehicleModel.Brand,
					Models: vehicleModel.Models,
					Patent: vehicleModel.Patent,
					Year:   vehicleModel.Year,
				},
				Address: dtos.AddressRespone{
					ID:          addressModel.ID,
					Street:      addressModel.Street,
					Number:      addressModel.Number,
					Floor:       addressModel.Floor,
					Departament: addressModel.Departament,
					Observation: addressModel.Observation,
					Favourite:   addressModel.Favourite,
					City: dtos.AllCityResponse{
						ID:          cityModel.ID,
						Description: cityModel.Description,
						CP:          cityModel.CP,
					},
				},
			},
		}

		var modelDayOrder []models.DayOrder

		if err := db.Table("day_orders").
			Select("distinct day_orders.order_id, day_orders.address_id ").
			Joins("left join day_menus ON day_menus.id = day_orders.day_menu_id").
			Joins("left join orders ON orders.id = day_orders.order_id").
			Where("address_id IN (select tanda_addresses.address_id from tanda_addresses where tanda_addresses.tanda_id = ?)", tanda.ID).
			Where("day_menus.date = ?", date.Format("2006-01-02")).
			Scan(&modelDayOrder).
			Error; err != nil {
			return nil, db.Error
		}

		var ordersRes []dtos.OrdersRes

		for _, dayOrder := range modelDayOrder {

			modelOrder, err := GetModelOrderById(dayOrder.OrderID)
			if err != nil {
				return nil, err
			}

			modelClient, err := dbClient.GetClientById(modelOrder.ClientID)
			if err != nil {
				return nil, err
			}

			pathologies, err := dbClient.GetPathologiesClient(modelClient.ID)
			if err != nil {
				return nil, err
			}

			addressOrderModel, err := dbAddress.GetAddressById(dayOrder.AddressID)
			if err != nil {
				return nil, err
			}

			cityOrderModel, err := dbCity.GetCityById(addressOrderModel.CityID)
			if err != nil {
				return nil, err
			}

			var resCategoryCant []dtos.ResCantDB

			if err := db.Table("day_orders").
				Select("categories.id, categories.description, categories.title, categories.color, categories.price, sum(day_orders.amount) as cant").
				Joins("left join day_menus ON day_menus.id = day_orders.day_menu_id ").
				Joins("left join categories ON categories.id = day_menus.category_id").
				Joins("left join orders ON orders.id = day_orders.order_id").
				Where("day_menus.date = ?", date.Format("2006-01-02")).
				Where("day_orders.address_id IN (select tanda_addresses.address_id from tanda_addresses where tanda_addresses.tanda_id = ?)", tanda.ID).
				Where("categories.active = 1").
				Where("orders.client_id = ?", modelClient.ID).
				Where("orders.id = ?", modelOrder.ID).
				Group("day_menus.category_id").
				Find(&resCategoryCant).Error; err != nil {
				return nil, err
			}

			var cantClientTable []dtos.CategoryTable

			for _, res := range resCategoryCant {

				catCliTable := dtos.CategoryTable{
					Category: dtos.CategoryResponse{
						ID:          res.ID,
						Description: res.Description,
						Title:       res.Title,
						Price:       res.Price,
						Color:       res.Color,
					},
					Cant: res.Cant,
				}

				cantClientTable = append(cantClientTable, catCliTable)
			}

			var obsDayOrderClient []string

			if err := db.Table("day_orders").
				Select("day_orders.observation").
				Joins("left join day_menus ON day_menus.id = day_orders.day_menu_id ").
				Where("day_menus.date = ?", date.Format("2006-01-02")).
				Where("day_orders.order_id = ?", modelOrder.ID).
				Find(&obsDayOrderClient).Error; err != nil {
				return nil, err
			}

			orderRes := dtos.OrdersRes{
				ID:          modelOrder.ID,
				OrderDate:   modelOrder.OrderDate,
				Observation: modelOrder.Observation,
				Total:       modelOrder.Total,
				Status:      modelOrder.Status,
				Client: dtos.Client{
					ID:             modelClient.ID,
					Name:           modelClient.Name,
					LastName:       modelClient.LastName,
					Email:          modelClient.Email,
					PhonePrimary:   modelClient.PhonePrimary,
					PhoneSecondary: modelClient.PhoneSecondary,
					ObsClient:      modelClient.Observation,
					BornDate:       modelClient.BornDate,
					Address:        nil,
					Pathologies:    pathologies,
				},
				Address: dtos.AddressRespone{
					ID:          addressOrderModel.ID,
					Street:      addressOrderModel.Street,
					Number:      addressOrderModel.Number,
					Floor:       addressOrderModel.Floor,
					Departament: addressOrderModel.Departament,
					Observation: addressOrderModel.Observation,
					Favourite:   addressOrderModel.Favourite,
					City: dtos.AllCityResponse{
						ID:          cityOrderModel.ID,
						Description: cityOrderModel.Description,
						CP:          cityOrderModel.CP,
					},
				},
				CategoryTable: cantClientTable,
				Observations:  obsDayOrderClient,
			}

			ordersRes = append(ordersRes, orderRes)

		}

		var resCategoryCant []dtos.ResCantDB

		if err := db.Table("day_orders").
			Select("categories.id, categories.description, categories.title, categories.color, categories.price, sum(day_orders.amount) as cant").
			Joins("left join day_menus ON day_menus.id = day_orders.day_menu_id ").
			Joins("left join categories ON categories.id = day_menus.category_id").
			Where("day_menus.date = ?", date.Format("2006-01-02")).
			Where("day_orders.address_id IN (select tanda_addresses.address_id from tanda_addresses where tanda_addresses.tanda_id = ?)", tanda.ID).
			Where("categories.active = 1").
			Group("day_menus.category_id").
			Find(&resCategoryCant).Error; err != nil {
			return nil, err
		}

		var catArrTable []dtos.CategoryTable

		for _, res := range resCategoryCant {

			catTable := dtos.CategoryTable{
				Category: dtos.CategoryResponse{
					ID:          res.ID,
					Description: res.Description,
					Title:       res.Title,
					Price:       res.Price,
					Color:       res.Color,
				},
				Cant: res.Cant,
			}

			catArrTable = append(catArrTable, catTable)

		}

		tandaTable := dtos.Tanda{
			Tanda:         tandaRes,
			CategoryTable: catArrTable,
			OrderRes:      ordersRes,
		}

		tandasTable = append(tandasTable, tandaTable)

	}

	var resCategoryCant []dtos.ResCantDB

	if err := db.Table("day_orders ").
		Select("categories.id, categories.description, categories.title, categories.color, categories.price, sum(day_orders.amount) as cant").
		Joins("left join day_menus ON day_menus.id = day_orders.day_menu_id ").
		Joins("left join categories ON categories.id = day_menus.category_id").
		Where("day_menus.date = ?", date.Format("2006-01-02")).
		Where("day_orders.address_id IN (select tanda_addresses.address_id from tanda_addresses)").
		Where("categories.active = 1").
		Group("day_menus.category_id").
		Find(&resCategoryCant).Error; err != nil {
		return nil, err
	}

	var cantTotalTable []dtos.CategoryTable

	for _, res := range resCategoryCant {

		cantTotTable := dtos.CategoryTable{
			Category: dtos.CategoryResponse{
				ID:          res.ID,
				Description: res.Description,
				Title:       res.Title,
				Price:       res.Price,
				Color:       res.Color,
			},
			Cant: res.Cant,
		}

		cantTotalTable = append(cantTotalTable, cantTotTable)

	}

	response := dtos.OrdersResponse{
		TandasTable: dtos.TandaTable{
			TandaTable:    tandasTable,
			CategoryTable: cantTotalTable,
		},
	}

	return &response, db.Error

}
