package services

import (
	"fmt"
	"strconv"
	"sync"
	"os"
	"github.com/xuri/excelize/v2"
)
func GetExcelFilePath() string {
    path := os.Getenv("EXCEL_FILE_PATH")
    if path == "" {
        path = "../../data/products.xlsx" // Default path for local testing
    }
    fmt.Println("Using Excel file path:", path)
    return path
}

var (
	filePath = GetExcelFilePath()
	mutex    sync.Mutex
)


// UpdateStock modifies product quantity based on change (+restock, -sale)
func UpdateStock(productID string, change int) error {
	mutex.Lock()
	defer mutex.Unlock()

	file, err := excelize.OpenFile(filePath)
	if err != nil {
		return err
	}

	sheet := "Sheet1"
	rows, _ := file.GetRows(sheet)

	for i, row := range rows {
		if i == 0 {
			continue
		}
		if row[0] == productID {
			currentQty, _ := strconv.Atoi(row[4])
			newQty := currentQty + change
			if newQty < 0 {
				return fmt.Errorf("insufficient stock")
			}
			file.SetCellValue(sheet, fmt.Sprintf("E%d", i+1), newQty)
			return file.SaveAs(filePath)
		}
	}

	return fmt.Errorf("product not found")
}
