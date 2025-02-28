package services

import (
	"strconv"
	"sync"
	"github.com/xuri/excelize/v2"
	"report-service/models"
	"errors"
	"fmt"
	"github.com/google/uuid"
)


var mutex sync.Mutex

// SaveProduct writes a new product to the report file 
func SaveProduct(products []*models.ProductReport, ReportfilePath string) error {
	// Create a new file if it doesn't existm
	file := excelize.NewFile()
		file.SetSheetRow("Sheet1", "A1", &[]string{"ID", "Name", "Description", "Price", "Quantity"})

	sheet := "Sheet1"
	rows, _ := file.GetRows(sheet)
	rowCount := len(rows)
	for _,p := range products{
        
	    // Add a new product to report
	    p.ProductID = uuid.New().String()
	    rowCount++
	    file.SetCellValue(sheet, fmt.Sprintf("A%d", rowCount), p.ProductID)          
	    file.SetCellValue(sheet, fmt.Sprintf("B%d", rowCount), p.Name)       
	    file.SetCellValue(sheet, fmt.Sprintf("C%d", rowCount), p.Description) 
	    file.SetCellValue(sheet, fmt.Sprintf("D%d", rowCount), p.Price)       
	    file.SetCellValue(sheet, fmt.Sprintf("E%d", rowCount), p.Quantity)    
    
	}
  return file.SaveAs(ReportfilePath)
}

// GetInventoryReport fetches all product data
func GetInventoryReport(restockThreshold int, filePath string, reportFilePath string) ([]*models.ProductReport, error) {
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
			Description: row[2],
			Price:       price,
			Quantity:    quantity,
		})
		}
	}
	if err := SaveProduct(report, reportFilePath); err != nil {
		return report,err
	}
	return report, nil
}
