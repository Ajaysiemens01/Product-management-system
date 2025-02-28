package services

import (
	"strconv"
	"sync"
	"github.com/xuri/excelize/v2"
	"report-service/models"
	"errors"
)


var mutex sync.Mutex


// GetInventoryReport fetches all product data
func GetInventoryReport(restockThreshold int, filePath string) ([]*models.ProductReport, error) {
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
		if quantity < restockThreshold {
		report = append(report, &models.ProductReport{
			ProductID:   row[0],
			Name:        row[1],
			Price:       price,
			Quantity:    quantity,
		})
		}
	}
	return report, nil
}
