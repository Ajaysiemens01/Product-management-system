package services

import (
	"fmt"
	"strconv"
	"sync"
	"github.com/xuri/excelize/v2"
	"inventory-service/models"
	"errors"
)

var mutex    sync.Mutex


// UpdateStock modifies product attributes based on provided fields
func UpdateStock(update *models.InventoryUpdate,filePath string) error {
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
			continue // Skip header
		}
		if row[0] == update.ProductID {
			// Update fields if they are provided
			if  update.StockAdded > 0{
				currentQty, _ := strconv.Atoi(row[4])
				newQty := currentQty + update.StockAdded
				file.SetCellValue(sheet, fmt.Sprintf("E%d", i+1), newQty)
			} else if update.StockAdded <0 {
				return errors.New("validation Error: Stock must be greaterthan zero")
			}
			if update.Name != "" {
				file.SetCellValue(sheet, fmt.Sprintf("B%d", i+1), update.Name)
			}
			if update.Description != "" {
				file.SetCellValue(sheet, fmt.Sprintf("C%d", i+1), update.Description)
			} 
			if update.Price > 0 {
				file.SetCellValue(sheet, fmt.Sprintf("D%d", i+1), update.Price)
			} else if update.Price < 0{
				return errors.New("validation Error: Price must be greaterthan zero")
			}
			if update.Quantity > 0 {
				file.SetCellValue(sheet, fmt.Sprintf("E%d", i+1), update.Quantity)
			} else if update.Quantity < 0  {
				return errors.New("validation Error: Quantity must be greaterthan zero")
			}

			// Save changes
			return file.SaveAs(filePath)
		}
	}

	return errors.New("product not found")
}

