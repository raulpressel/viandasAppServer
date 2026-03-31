package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	dbAddress "viandasApp/db/address"
	dbCategories "viandasApp/db/categories"
	dbClient "viandasApp/db/client"
	dbMenu "viandasApp/db/menu"
	dbOrder "viandasApp/db/order"
	dbSetting "viandasApp/db/setting"
	dbTanda "viandasApp/db/tanda"
	"viandasApp/dtos"
	"viandasApp/models"
)

type response struct {
	OrderId      int                  `json:"idOrder"`
	Total        float32              `json:"total"`
	CantDelivery int                  `json:"cantDelivery"`
	Categories   []dtos.CategoryTable `json:"categories"`
}
type categoriesCant struct {
	Cant        int    `json:"cant"`
	Description string `json:"description"`
}

func UploadOrder(rw http.ResponseWriter, r *http.Request) {

	var orderDto dtos.OrderRequest

	var orderModel models.Order

	var dOrderModel models.DayOrder

	var dayOrderModel []models.DayOrder

	var deliveryModel models.Delivery

	var deliverysModel []models.Delivery

	err := json.NewDecoder(r.Body).Decode(&orderDto)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	clientModel, err := dbClient.GetClientById(orderDto.IDClient)

	if err != nil {
		http.Error(rw, "Ocurrio un error al obtener el ID del cliente "+err.Error(), http.StatusInternalServerError)
		return
	}

	if clientModel.ID == 0 {
		http.Error(rw, "El Cliente enviado no existe ", http.StatusBadRequest)
		return
	}

	auxCalc, valid := ProcessOrder(orderDto)

	if !valid {
		http.Error(rw, "Ocurrio un error al procesar los calculos de la orden "+err.Error(), http.StatusInternalServerError)
		return
	}

	orderModel.ClientID = orderDto.IDClient

	orderModel.Observation = orderDto.Observation

	orderModel.Total = auxCalc.Total

	orderModel.StatusOrderID = 1 //se da de alta orden y queda con estado 1 - Activa

	orderModel.Paid = false

	orderModel.OrderDate, err = time.Parse(time.RFC3339, orderDto.Date)
	if err != nil {
		http.Error(rw, "Error en el formato de fecha recibido "+err.Error(), http.StatusBadRequest)
		return
	}

	for _, day := range orderDto.DaysOrderRequest {

		if day.Amount > 0 {

			dOrderModel.Amount = day.Amount

			deliveryModel.DeliveryMenuAmount = day.Amount

			dayMenuModel, err := dbMenu.GetDayMenuById(day.IDDayFood)

			if err != nil {
				http.Error(rw, "Ocurrio un error al obtener el ID del menu "+err.Error(), http.StatusInternalServerError)
				return
			}

			if dayMenuModel.ID == 0 {
				http.Error(rw, "El Day Menu enviado no existe ", http.StatusBadRequest)
				return
			}

			dOrderModel.DayMenuID = day.IDDayFood

			dOrderModel.Observation = day.Observation

			//dOrderModel.Active = true

			dOrderModel.Status = true

			addressModel, err := dbAddress.GetAddressById(day.IDAddress)

			if err != nil {
				http.Error(rw, "Ocurrio un error al obtener el ID de la direccion "+err.Error(), http.StatusInternalServerError)
				return
			}

			if addressModel.ID == 0 {
				http.Error(rw, "La direccion enviada no existe ", http.StatusBadRequest)
				return
			}

			dOrderModel.AddressID = day.IDAddress

			deliveryModel.AddressID = dOrderModel.AddressID

			deliveryModel.Status = false

			categoryModel, err := dbCategories.GetCategoryById(dayMenuModel.CategoryID)

			if err != nil {
				http.Error(rw, "Ocurrio un error al obtener el ID de la categoria "+err.Error(), http.StatusInternalServerError)
				return
			}

			deliveryModel.DeliveryMenuPrice = categoryModel.Price * float32(dOrderModel.Amount)

			deliveryModel.DeliveryDate = dayMenuModel.Date

			idTanda, err := dbTanda.CheckExistTandaByAddressId(deliveryModel.AddressID)

			if err != nil {
				http.Error(rw, "Ocurrio un error al obtener el ID de la Tanda "+err.Error(), http.StatusInternalServerError)
				return
			}

			if idTanda > 0 {
				deliveryDriverId, err := dbTanda.GetDeliveryDriverIdByTandaId(idTanda)
				if err != nil {
					http.Error(rw, "Ocurrio un error al obtener el ID del Cadete "+err.Error(), http.StatusInternalServerError)
					return
				}

				tempID := uint(deliveryDriverId)

				deliveryModel.DeliveryDriverID = &tempID
			} else {
				deliveryModel.DeliveryDriverID = nil
			}

			deliveryModel.PercentageDiscount = auxCalc.Percentage

			dayOrderModel = append(dayOrderModel, dOrderModel)

			//if deliveryModel.AddressID != 100 {
			deliverysModel = append(deliverysModel, deliveryModel)
			//}

		}
	}

	deliveriesModel, valid := processDeliveryOrder(deliverysModel)

	if !valid {
		http.Error(rw, "Ocurrio un error al calcular los valores del delivery order", http.StatusInternalServerError)
		return
	}

	status, err, orderId := dbOrder.UploadOrder(orderModel, dayOrderModel, deliveriesModel)

	if err != nil {
		http.Error(rw, "Ocurrio un error al intentar ingresar el pedido "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "no se ha logrado cargar el pedido en la BD", http.StatusInternalServerError)
		return
	}

	categoriesCant := calcAmountsCategories(dayOrderModel)

	cantDelivery := calcCantDeliveries(deliveriesModel)

	res := response{
		OrderId:      orderId.IDOrder,
		Total:        orderModel.Total,
		CantDelivery: cantDelivery,
		Categories:   *categoriesCant,
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(res)

}

func calcCantDeliveries(deliveriesModel []models.Delivery) int {

	var cantDeliveries int

	for _, deli := range deliveriesModel {

		if deli.AddressID != 100 {
			cantDeliveries++
		}

	}

	return cantDeliveries
}

func calcAmountsCategories(dayOrderModel []models.DayOrder) *[]dtos.CategoryTable {

	var arr []int

	var categories []dtos.CategoryTable

	for _, day := range dayOrderModel {

		dayMenu, err := dbMenu.GetDayMenuById(day.DayMenuID)

		if err != nil {
			return nil
		}

		if day.Amount > 1 {
			array := multiplyElement(dayMenu.CategoryID, day.Amount)
			arr = append(arr, array...)
		} else {
			arr = append(arr, dayMenu.CategoryID)
		}

	}

	counts := countOccurrences(arr)

	for num, count := range counts {
		categoryModel, err := dbCategories.GetCategoryById(num)
		if err != nil {
			return nil
		}
		category := dtos.CategoryTable{
			Cant: count,
			Category: dtos.CategoryResponse{
				ID:    categoryModel.ID,
				Title: categoryModel.Title,
			},
		}
		categories = append(categories, category)
	}

	return &categories

}

func countOccurrences(arr []int) map[int]int {
	counts := make(map[int]int)
	for _, num := range arr {
		counts[num]++
	}
	return counts
}

func multiplyElement(element, multiplier int) []int {
	arr := make([]int, multiplier)
	for i := range arr {
		arr[i] = element
	}
	return arr
}

func processDeliveryOrder(deliverysModel []models.Delivery) ([]models.Delivery, bool) {

	accumulatedValues := make(map[time.Time]models.Delivery)

	// Recorrer el array de deliveries
	for _, delivery := range deliverysModel {
		// Verificar si la fecha ya existe en el map
		if accumulatedDelivery, ok := accumulatedValues[delivery.DeliveryDate]; ok {
			// La fecha ya existe, acumular los valores
			accumulatedDelivery.DeliveryMenuPrice += delivery.DeliveryMenuPrice
			accumulatedDelivery.DeliveryMenuAmount += delivery.DeliveryMenuAmount
			accumulatedValues[delivery.DeliveryDate] = accumulatedDelivery
		} else {
			// La fecha no existe, agregar el delivery completo al map
			accumulatedValues[delivery.DeliveryDate] = delivery
		}
	}

	// Crear un nuevo array con los valores acumulados
	uniqueDeliveries := make([]models.Delivery, 0, len(accumulatedValues))
	for _, accumulatedDelivery := range accumulatedValues {
		uniqueDeliveries = append(uniqueDeliveries, accumulatedDelivery)
	}

	for i := range uniqueDeliveries {
		if uniqueDeliveries[i].AddressID != 100 {
			addressModel, err := dbAddress.GetAddressById(uniqueDeliveries[i].AddressID)

			if err != nil {
				return uniqueDeliveries, false
			}
			zoneModel, err := dbSetting.GetZoneById(addressModel.IDZone)
			if err != nil {

			}
			priceFactor := PriceFactor(uniqueDeliveries[i].DeliveryMenuAmount)
			uniqueDeliveries[i].DeliveryPrice = zoneModel.Price * priceFactor
		}

	}

	return uniqueDeliveries, true
}
