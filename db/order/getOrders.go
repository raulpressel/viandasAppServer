package db

import (
	"time"
	"viandasApp/db"
	dbAddress "viandasApp/db/address"
	dbCat "viandasApp/db/categories"
	dbCity "viandasApp/db/city"
	dbClient "viandasApp/db/client"
	dbDeliveryDriver "viandasApp/db/deliveryDriver"
	dbFood "viandasApp/db/food"
	dbMenu "viandasApp/db/menu"

	dbVehicle "viandasApp/db/vehicle"
	"viandasApp/dtos"
	"viandasApp/models"
)

type Resp struct {
	Description string
	Total       int
}

func GetOrders(date time.Time) (*dtos.OrdersResponse, error) {

	db := db.GetDB()

	//var modelOrder models.Order

	//var response dtos.OrdersResponse

	var modelDayOrder []models.DayOrder

	modelTanda := []models.Tanda{}

	//modelTandaAddress := []models.TandaAddress{}

	if err := db.Find(&modelTanda, "active = 1").Error; err != nil {
		return nil, err
	}

	//var tandaTable dtos.Tanda
	var tandasTable []dtos.Tanda

	//var tandaRes dtos.TandaRes

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

		//tandasRes = append(tandasRes, tandaRes)

		/* if err := db.Find(&modelTandaAddress, "tanda_id = ?", tanda.ID).Error; err != nil {
			return nil, err
		} */

		if err := db.Table("day_orders").
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

			// reemplazar por FIRST o funciones que ya existen

			/* var modelOrder models.Order

			if err := db.Find(&modelOrder, "id = ?", dayOrder.OrderID).Error; err != nil {
				return nil, err
			} */

			modelOrder, err := GetModelOrderById(dayOrder.OrderID)
			if err != nil {
				return nil, err
			}

			modelDayMenu, err := dbMenu.GetDayMenuById(dayOrder.DayMenuID)
			if err != nil {
				return nil, err
			}

			modelFood, err := dbFood.GetFoodById(modelDayMenu.FoodID)
			if err != nil {
				return nil, err
			}

			modelCategory, err := dbCat.GetCategoryById(modelDayMenu.CategoryID)
			if err != nil {
				return nil, err
			}

			modelClient, err := dbClient.GetClientById(modelOrder.ClientID)
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
					Pathologies:    nil, //me falta esto//pathology hacer consulta q traiga
				},
				DayOrder: []dtos.DayOrderResponse{
					dtos.DayOrderResponse{
						ID:          dayOrder.ID,
						Amount:      dayOrder.Amount,
						Observation: dayOrder.Observation,
						Status:      dayOrder.Status,
						Date:        modelDayMenu.Date,
						Food: dtos.FoodResponse{
							ID:          modelFood.ID,
							Title:       modelFood.Title,
							Description: modelFood.Description,
						},
						Category: dtos.CategoryResponse{
							ID:          modelCategory.ID,
							Description: modelCategory.Description,
							Title:       modelCategory.Description,
							Price:       modelCategory.Price,
							Color:       modelCategory.Color,
							//Location:    nil,
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
					},
				},
			}

			ordersRes = append(ordersRes, orderRes)

		}

		var resCategoryCant []dtos.ResCantDB

		if err := db.Table("tanda_addresses ").
			Select("categories.id, categories.description, categories.title, categories.color, categories.price, sum(day_orders.amount) as cant").
			Joins("left join day_orders ON tanda_addresses.address_id = day_orders.address_id ").
			Joins("left join day_menus ON day_menus.id = day_orders.day_menu_id ").
			Joins("left join categories ON categories.id = day_menus.category_id").
			Where("day_menus.date = ?", date.Format("2006-01-02")).
			Where("tanda_addresses.tanda_id = ?", tanda.ID).
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

		//select  category_id, sum(amount)
		//from tanda_addresses
		//left join day_orders ON tanda_addresses.address_id = day_orders.address_id
		//left join day_menus ON day_menus.id = day_orders.day_menu_id
		//where day_menus.date = "2023-01-25" group by category_id;

	}

	var resCategoryCant []dtos.ResCantDB

	if err := db.Table("tanda_addresses ").
		Select("categories.id, categories.description, categories.title, categories.color, categories.price, sum(day_orders.amount) as cant").
		Joins("left join day_orders ON tanda_addresses.address_id = day_orders.address_id ").
		Joins("left join day_menus ON day_menus.id = day_orders.day_menu_id ").
		Joins("left join categories ON categories.id = day_menus.category_id").
		Where("day_menus.date = ?", date.Format("2006-01-02")).
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
			TandaTable: tandasTable,
		},
		CategoryTable: cantTotalTable,
	}

	/*


		if err := db.Find(&modelDayMenu, "date = ?", date.Format("2006-01-02")).Error; err != nil {
			return nil, err
		}

		for _, dayMenu := range modelDayMenu {

			if err := db.Find(&modelDayOrder, "day_menu_id = ?", dayMenu.ID).Error; err != nil {
				return nil, err
			}

			for _, dayOrder := range modelDayOrder {
				dayOrder.AddressID
			}


		}

			err = db.Table("day_orders").
				Select("day_orders.id, day_orders.observation, day_orders.amount, day_orders.status, day_orders.address_id, day_orders.day_menu_id").
				Where("day_orders.order_id = ?", ord.ID).
				Scan(&daysOrderModel).Error

			dayMenuModel, err := dbMenu.GetDayMenuById(daysOrderModel.DayMenuID)
			if err != nil {
				return dtoOrder, err
			}

			menuModel, err := dbMenu.GetMenuByTurnMenuID(dayMenuModel.TurnMenuID)
			if err != nil {
				return dtoOrder, err
			}

			orders.ID = ord.ID
			orders.OrderDate = ord.OrderDate
			orders.Observation = ord.Observation
			orders.Status = ord.Status
			orders.Total = ord.Total
			orders.DateStart = menuModel.DateStart
			orders.DateEnd = menuModel.DateEnd

			dtoOrder.Order = append(dtoOrder.Order, orders)
		}

		/* 	for _, valor := range modelAllMenu {
			id, err := GetIdMenuActive(valor.Menuid)
			if err != nil {
				return responseAllMenu, err
			}
			if id > 0 {
				valor.IsCurrent = true
			}
			responseAllMenu = append(responseAllMenu, *valor.ToAllMenuResponse())
		} */

	return &response, db.Error

}
