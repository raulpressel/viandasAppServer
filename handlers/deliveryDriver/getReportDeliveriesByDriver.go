package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
	addressDB "viandasApp/db/address"
	clientDB "viandasApp/db/client"
	deliDriverDB "viandasApp/db/deliveryDriver"
	orderDB "viandasApp/db/order"

	"viandasApp/dtos"
	"viandasApp/models"

	"github.com/xuri/excelize/v2"
)

type response struct {
	Path string `json:"path"`
}

func GetReportDeliveriesByDriver(rw http.ResponseWriter, r *http.Request) {

	var deliveryDto dtos.DeliveryRequest

	var deliveryDriverModel models.DeliveryDriver

	err := json.NewDecoder(r.Body).Decode(&deliveryDto)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	if deliveryDto.DeliveryDriverID == nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	deliveryDriverModel, err = deliDriverDB.GetDeliveryDriverByID(*deliveryDto.DeliveryDriverID)
	if err != nil {
		http.Error(rw, "Error al obtener el delivery driver", http.StatusInternalServerError)
		return
	}
	if deliveryDriverModel.ID == 0 {
		http.Error(rw, "No existe un Delivery Driver con ese ID ", http.StatusBadRequest)
		return
	}

	dateStart, err := time.Parse(time.RFC3339, deliveryDto.DateStart)
	if err != nil {
		http.Error(rw, "Error en el formato de fecha recibido "+err.Error(), http.StatusBadRequest)
		return
	}

	dateEnd, err := time.Parse(time.RFC3339, deliveryDto.DateEnd)
	if err != nil {
		http.Error(rw, "Error en el formato de fecha recibido "+err.Error(), http.StatusBadRequest)
		return
	}

	deliveries, err := deliDriverDB.GetReportDeliveryDriver(deliveryDriverModel.ID, dateStart, dateEnd)

	if err != nil {
		http.Error(rw, "Error al recuperar los datos para el reporte  "+err.Error(), http.StatusInternalServerError)
		return
	}

	deliveryExcel := processDeliveryOrder(*deliveries)

	path, valid := gerateXLSX(deliveryExcel, dateStart, dateEnd)

	if !valid {
		http.Error(rw, "Error al generar el reporte  "+err.Error(), http.StatusInternalServerError)
		return
	}

	var res response

	res.Path = path

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(res)
}

type DeliveryExcel struct {
	DeliveryDriver string
	ClientExcel    ClientExcel
}

type ClientExcel struct {
	OrderID        int
	Client         string
	AddressExcel   []AddressExcel
	DatePriceExcel []DatePriceExcel
}

type AddressExcel struct {
	Address string
}

type DatePriceExcel struct {
	Date  time.Time
	Price int
}

func processDeliveryOrder(deliveries []models.Delivery) []DeliveryExcel {

	var deliveriesExcel []DeliveryExcel

	var currentOrderID int

	var currentAddressID int

	for _, delivery := range deliveries {
		if delivery.OrderID != currentOrderID {

			deliveryDriverModel, _ := deliDriverDB.GetDeliveryDriverByID(int(*delivery.DeliveryDriverID))
			orderModel, _ := orderDB.GetModelOrderById(delivery.OrderID)
			clientModel, _ := clientDB.GetClientById(orderModel.ClientID)

			deliveriesExcel = append(deliveriesExcel, DeliveryExcel{
				DeliveryDriver: deliveryDriverModel.Name + " " + deliveryDriverModel.LastName,
				ClientExcel: ClientExcel{
					OrderID:        delivery.OrderID,
					Client:         clientModel.Name + " " + clientModel.LastName,
					AddressExcel:   []AddressExcel{},
					DatePriceExcel: []DatePriceExcel{},
				},
			})
			currentOrderID = delivery.OrderID

			currentAddressID = 0
		}

		if delivery.AddressID != currentAddressID {
			addressModel, _ := addressDB.GetAddressById(delivery.AddressID)
			deliveriesExcel[len(deliveriesExcel)-1].ClientExcel.AddressExcel = append(deliveriesExcel[len(deliveriesExcel)-1].ClientExcel.AddressExcel, AddressExcel{
				Address: addressModel.Street + " " + addressModel.Number + " " + addressModel.Observation,
			})
			currentAddressID = delivery.AddressID
		}

		// Agregar una nueva DatePrice al último OtroDelivery agregado
		deliveriesExcel[len(deliveriesExcel)-1].ClientExcel.DatePriceExcel = append(deliveriesExcel[len(deliveriesExcel)-1].ClientExcel.DatePriceExcel, DatePriceExcel{
			Date:  delivery.DeliveryDate,
			Price: int(delivery.DeliveryPrice),
		})
	}

	return deliveriesExcel

}

func gerateXLSX(deliveriesExcel []DeliveryExcel, dateStart time.Time, dateEnd time.Time) (string, bool) {

	fileDir := "/var/www/default/htdocs/public/reports"

	if err := os.MkdirAll(fileDir, os.ModePerm); err != nil {
		return "", false
	}

	file := excelize.NewFile()
	defer func() {
		if err := file.Close(); err != nil {
			return
		}
	}()

	file.SetColWidth("Sheet1", "A", "A", 13)
	file.SetRowHeight("Sheet1", 1, 20)
	file.SetColWidth("Sheet1", "B", "B", 40)
	file.SetRowHeight("Sheet1", 2, 20)
	file.SetColWidth("Sheet1", "C", "C", 40)

	styleA1, _ := file.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Size: 14, Bold: true, Color: "070808"},
		Alignment: &excelize.Alignment{Vertical: "center"},
		Border: []excelize.Border{
			{
				Type:  "left",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "top",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "bottom",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "right",
				Color: "#000000",
				Style: 1,
			},
		},
	})

	styleA2, _ := file.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Size: 14, Bold: true, Color: "070808"},
		Alignment: &excelize.Alignment{Vertical: "center", Horizontal: "center"},
		Border: []excelize.Border{
			{
				Type:  "left",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "top",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "bottom",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "right",
				Color: "#000000",
				Style: 1,
			},
		},
	})

	styleTotal, _ := file.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Size: 14, Bold: true, Color: "070808"},
		Alignment: &excelize.Alignment{Vertical: "center", Horizontal: "right"},
	})

	var cantDays int

	var cleanDays []time.Time
	currentDate := dateStart
	for currentDate.Before(dateEnd) || currentDate.Equal(dateEnd) {
		if currentDate.Weekday() != time.Saturday && currentDate.Weekday() != time.Sunday {
			cantDays++
			cleanDays = append(cleanDays, currentDate)
		}
		currentDate = currentDate.AddDate(0, 0, 1)

	}

	startCell := "A1"
	endCell, _ := excelize.CoordinatesToCellName(3+cantDays, 2)

	column := string(endCell[0]) + "%d"

	file.SetColWidth("Sheet1", "D", string(endCell[0]), 15)

	file.MergeCell("Sheet1", "A1", fmt.Sprintf(column, 1))

	file.SetCellStyle("Sheet1", startCell, endCell, styleA1)

	file.SetCellStyle("Sheet1", "A2", endCell, styleA2)

	fileName := "Reporte " + deliveriesExcel[0].DeliveryDriver + " " + dateStart.Format("02-01") + " al " + dateEnd.Format("02-01")

	file.SetCellValue("Sheet1", "A1", "REPARTIDOR: "+deliveriesExcel[0].DeliveryDriver)
	file.SetCellValue("Sheet1", "A2", "Nº ORDEN")
	file.SetCellValue("Sheet1", "B2", "NOMBRE Y APELLIDO")
	file.SetCellValue("Sheet1", "C2", "DIRECCIÓN")

	for i, day := range cleanDays {
		cell, _ := excelize.CoordinatesToCellName(4+i, 2)

		formatDate := day.Format("02/01/2006")

		file.SetCellValue("Sheet1", cell, formatDate)
	}

	var row int

	for i, delivery := range deliveriesExcel {

		row = i + 3

		file.SetRowHeight("Sheet1", row, 18)

		var fill string
		if i%2 == 0 {
			fill = "F3F3F3"
		} else {
			fill = "FFFFFF"
		}

		styleData, _ := file.NewStyle(&excelize.Style{
			Fill:      excelize.Fill{Type: "pattern", Pattern: 1, Color: []string{fill}},
			Font:      &excelize.Font{Size: 13, Color: "0a0a0a"},
			Alignment: &excelize.Alignment{Vertical: "center"},
			Border: []excelize.Border{
				{
					Type:  "left",
					Color: "#000000",
					Style: 1,
				}, {
					Type:  "top",
					Color: "#000000",
					Style: 1,
				}, {
					Type:  "bottom",
					Color: "#000000",
					Style: 1,
				}, {
					Type:  "right",
					Color: "#000000",
					Style: 1,
				},
			},
		})

		file.SetCellStyle("Sheet1", fmt.Sprintf("A%d", row), fmt.Sprintf(column, row), styleData)

		file.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), delivery.ClientExcel.OrderID)
		file.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), delivery.ClientExcel.Client)
		file.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), delivery.ClientExcel.AddressExcel[0].Address)

		for i, day := range cleanDays {
			cell2, _ := excelize.CoordinatesToCellName(4+i, 2)

			formatDate := day.Format("02/01/2006")

			file.SetCellValue("Sheet1", cell2, formatDate)

			cell, _ := excelize.CoordinatesToCellName(4+i, row)
			for _, datePrice := range delivery.ClientExcel.DatePriceExcel {

				if day.UTC() == datePrice.Date.UTC() {
					file.SetCellValue("Sheet1", cell, datePrice.Price)
				}

			}
		}

	}

	rowInit := (row - len(deliveriesExcel)) + 1

	for i := 0; i < cantDays; i++ {
		cell, _ := excelize.CoordinatesToCellName(4+i, row+1)

		cellSum1, _ := excelize.CoordinatesToCellName(4+i, rowInit)

		cellSum2, _ := excelize.CoordinatesToCellName(4+i, row)

		formulaType := excelize.STCellFormulaTypeDataTable
		if err := file.SetCellFormula("Sheet1", cell, "=SUM("+cellSum1+":"+cellSum2+")",
			excelize.FormulaOpts{Type: &formulaType}); err != nil {
			fmt.Println(err)

		}

	}
	cellTotal, _ := excelize.CoordinatesToCellName(3, row+3)

	file.SetCellValue("Sheet1", cellTotal, "TOTAL: ")

	cell, _ := excelize.CoordinatesToCellName(4, row+3)

	cellSum1, _ := excelize.CoordinatesToCellName(4, row+1)

	cellSum2, _ := excelize.CoordinatesToCellName(3+cantDays, row+1)

	formulaType := excelize.STCellFormulaTypeDataTable
	if err := file.SetCellFormula("Sheet1", cell, "=SUM("+cellSum1+":"+cellSum2+")",
		excelize.FormulaOpts{Type: &formulaType}); err != nil {
		fmt.Println(err)

	}

	file.SetRowHeight("Sheet1", row+1, 20)
	file.SetCellStyle("Sheet1", cellSum1, cellSum2, styleA1)

	file.SetRowHeight("Sheet1", row+3, 22)

	file.SetCellStyle("Sheet1", cellTotal, cell, styleTotal)

	//cell, _ := excelize.CoordinatesToCellName(4, row+1)

	//file.SetCellFormula("Sheet1", cell, "=SUM(D3,D7)")

	fileDir = fileDir + "/" + fileName + ".xlsx"

	if err := file.SaveAs(fileDir); err != nil {
		return "", false
	}

	pathDownload := "/public/reports/" + fileName + ".xlsx"

	return pathDownload, true

}
