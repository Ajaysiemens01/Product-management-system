package services

import (
	"strconv"
	"sync"
	"github.com/xuri/excelize/v2"
	"report-service/models"
	"errors"
)

var mutex sync.Mutex
var filePath = "/app/data/products.xlsx"

// GetInventoryReport fetches all product data
func GetInventoryReport(restockThreshold int) ([]*models.ProductReport, error) {
	mutex.Lock()
	defer mutex.Unlock()

	file, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	if restockThreshold < 0 {
		return nil ,errors.New("invalid restock threshold")
	} 
	rows, _ := file.GetRows("Sheet1")
	var report []*models.ProductReport

	for i, row := range rows {
		if i == 0 {
			continue
		}
		price, _ := strconv.ParseFloat(row[3], 64)
		quantity, _ := strconv.Atoi(row[4])
		report = append(report, &models.ProductReport{
			ProductID:   row[0],
			Name:        row[1],
			Price:       price,
			Quantity:    quantity,
			NeedsRestock: quantity < restockThreshold,
		})
	}
	return report, nil
}
