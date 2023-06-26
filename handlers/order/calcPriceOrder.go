package handlers

import (
	"encoding/json"
	"math"
	"net/http"
	"time"

	dbAddress "viandasApp/db/address"
	dbCategory "viandasApp/db/categories"
	dbClient "viandasApp/db/client"
	dbMenu "viandasApp/db/menu"
	dbSetting "viandasApp/db/setting"
	"viandasApp/dtos"
)

var priceTable = []struct {
	MinAmount   int
	MaxAmount   int
	PriceFactor float32
}{
	{1, 2, 1.0},
	{3, 4, 1.5},
	{5, 6, 2.0},
	{7, 8, 2.5},
	{9, 10, 3.0},
	{11, 12, 3.5},
	{13, 14, 4.0},
	{15, 16, 4.5},
	{17, 18, 5.0},
	{19, 20, 5.5},
	{21, 22, 6.0},
	{23, 24, 6.5},
	{25, 26, 7.0},
	{27, 28, 7.5},
	{29, 30, 8.0},
	{31, 32, 8.5},
	{33, 34, 9.0},
	{35, 36, 9.5},
	{37, 38, 10.0},
	{39, 40, 10.5},
	{41, 42, 11.0},
	{43, 44, 11.5},
	{45, 46, 12.0},
	{47, 48, 12.5},
	{49, 50, 13.0},
}

type responseTotal struct {
	SubTotal   float32 `json:"subTotal"`
	Discount   float32 `json:"discount"`
	Percentage float32 `json:"percentage"`
	Total      float32 `json:"total"`
	Delivery   float32 `json:"delivery"`
}

func CalcPriceOrder(rw http.ResponseWriter, r *http.Request) {

	var orderDto dtos.OrderRequest

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

	response, valid := ProcessOrder(orderDto)

	if !valid {
		http.Error(rw, "Ocurrio un error al procesar los calculos de la orden "+err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(response)

}

func allEqual(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[0] {
			return false
		}
	}
	return true
}

func isWeekday(t time.Time) bool {
	return t.Weekday() >= time.Monday && t.Weekday() <= time.Friday
}

func worksDays(year, week int) []time.Time {

	t := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)

	daysToAdd := time.Duration((week-1)*7-int(t.Weekday())) * 24 * time.Hour

	firstDay := t.Add(daysToAdd)

	weekdays := make([]time.Time, 0)
	for i := 0; i < 7; i++ {
		if isWeekday(firstDay) {
			weekdays = append(weekdays, firstDay)
		}
		firstDay = firstDay.AddDate(0, 0, 1)
	}
	return weekdays

}

func PriceFactor(amount int) float32 {
	var priceFactor float32 = 1.0
	for _, entry := range priceTable {
		if amount >= entry.MinAmount && amount <= entry.MaxAmount {
			priceFactor = entry.PriceFactor
			break
		}
	}

	return priceFactor

}

type auxCalcDeliveryDriver struct {
	amount    int
	date      time.Time
	idAddress int
}

func ProcessOrder(orderDto dtos.OrderRequest) (responseTotal, bool) {

	var auxCalc responseTotal

	var dates []time.Time

	var amount int = 0

	var price float32 = 0.0

	var auxsCalcsDD []auxCalcDeliveryDriver

	var auxCalcDD auxCalcDeliveryDriver

	for _, day := range orderDto.DaysOrderRequest {

		if day.Amount > 0 {

			amount = amount + day.Amount

			dayMenuModel, err := dbMenu.GetDayMenuById(day.IDDayFood)

			if err != nil {
				return auxCalc, false
			}

			dates = append(dates, dayMenuModel.Date)

			cat, err := dbCategory.GetCategoryById(dayMenuModel.CategoryID)
			if err != nil {
				return auxCalc, false
			}

			price = price + (cat.Price * float32(day.Amount))

			auxCalcDD.amount = day.Amount
			auxCalcDD.date = dayMenuModel.Date
			auxCalcDD.idAddress = day.IDAddress

			auxsCalcsDD = append(auxsCalcsDD, auxCalcDD)

		}

	}

	if len(auxsCalcsDD) > 0 {

		var valid bool

		auxCalc.Delivery, valid = calcDeliveryPrice(auxsCalcsDD)

		if !valid {
			return auxCalc, false
		}
	}

	discounts, err := dbSetting.GetDiscounts()
	if err != nil {
		return auxCalc, false
	}

	auxCalc.SubTotal = price

	auxCalc.Discount = 0.0

	var filterDates []time.Time

	if len(dates) >= 3 && len(dates) < 5 {

		years := make([]int, amount)
		weeks := make([]int, amount)

		for i, date := range dates {
			years[i], weeks[i] = date.ISOWeek()
		}
		if allEqual(years) {
			if allEqual(weeks) {
				filterDates = append(filterDates, dates...)
			} else if allEqual(weeks[:3]) {
				if len(weeks[:3]) > 2 {
					filterDates = append(filterDates, dates[:3]...)
				}
			} else if allEqual(weeks[1:]) {
				if len(weeks[:1]) > 2 {
					filterDates = append(filterDates, dates[1:]...)
				}
			}
		}

		if len(filterDates) > 0 {

			workdays := worksDays(filterDates[0].ISOWeek())

			var diff []string

			for _, a := range workdays {
				found := false
				for _, b := range filterDates {
					if a.Format("2006-01-02") == b.Format("2006-01-02") {
						found = true
						break
					}
				}
				if !found {
					diff = append(diff, a.Format("2006-01-02"))
				}

			}

			if len(diff) > 0 {

				check, _ := dbMenu.CheckFeriado(diff)

				if check {
					amount = amount + len(diff)
				}

			}
		}

	}

	for i := range discounts {
		if amount >= discounts[i].Cant {
			auxCalc.Discount = price * (discounts[i].Percentage / 100)
			auxCalc.Percentage = discounts[i].Percentage
			break
		} else if i < len(discounts)-1 && amount >= discounts[i+1].Cant && amount < discounts[i].Cant {
			auxCalc.Discount = price * (discounts[i+1].Percentage / 100)
			auxCalc.Percentage = discounts[i+1].Percentage
			break
		}
	}

	redon := math.RoundToEven(float64(auxCalc.Discount))

	auxCalc.Discount = float32(redon)

	auxCalc.Total = (auxCalc.SubTotal - auxCalc.Discount) + auxCalc.Delivery

	return auxCalc, true

}

func calcDeliveryPrice(aux []auxCalcDeliveryDriver) (float32, bool) {

	var delivery float32

	accumulatedValues := make(map[time.Time]auxCalcDeliveryDriver)

	for _, delivery := range aux {

		if accumulatedDelivery, ok := accumulatedValues[delivery.date]; ok {

			accumulatedDelivery.amount += delivery.amount

			accumulatedValues[delivery.date] = accumulatedDelivery
		} else {

			accumulatedValues[delivery.date] = delivery
		}
	}

	uniqueDeliveries := make([]auxCalcDeliveryDriver, 0, len(accumulatedValues))
	for _, accumulatedDelivery := range accumulatedValues {
		uniqueDeliveries = append(uniqueDeliveries, accumulatedDelivery)
	}

	for i := range uniqueDeliveries {
		if uniqueDeliveries[i].idAddress != 100 {
			addressModel, err := dbAddress.GetAddressById(uniqueDeliveries[i].idAddress)

			if err != nil {
				return 0.0, false
			}

			zoneModel, err := dbSetting.GetZoneById(addressModel.IDZone)
			if err != nil {
				return 0.0, false
			}

			priceFactor := PriceFactor(uniqueDeliveries[i].amount)

			zoneModel.Price *= priceFactor

			delivery = delivery + zoneModel.Price

		}
	}

	return delivery, true
}
