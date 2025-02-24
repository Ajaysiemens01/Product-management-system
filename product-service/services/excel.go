package services

import (
	"errors"
	"fmt"
	"product-service/models"
	"strconv"
	"sync"
	"time"
	"os"
	"github.com/google/uuid"
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

// SaveProduct writes a new product to the Excel file but returns an error if the product already exists
func SaveProduct(p models.Product) error {
	mutex.Lock()
	defer mutex.Unlock()

	file, err := excelize.OpenFile(filePath)
	if err != nil {
		// Create a new file if it doesn't existm
		file = excelize.NewFile()
		file.SetSheetRow("Sheet1", "A1", &[]string{"ID", "Name", "Description", "Price", "Quantity"})
	}
	time.Sleep(5 * time.Second) // Simulate a slow operation
	sheet := "Sheet1"
	rows, _ := file.GetRows(sheet)

	// Check if the product already exists (based on Name)
	for i, row := range rows {
		if i == 0 || len(row) < 2 { // Skip header and empty rows
			continue
		}

		existingName := row[1] // Column B (Product Name)
		if existingName == p.Name {
			// Product already exists, return an error
			return errors.New("product with the same name already exists")
		}
	}
    
	// Add a new product if it doesn't exist
	p.ID = uuid.New().String()
	rowCount := len(rows) + 1
	file.SetCellValue(sheet, fmt.Sprintf("A%d", rowCount), p.ID)          
	if p.Name != "" && p.Price > 0 && p.Quantity >= 0 {
		file.SetCellValue(sheet, fmt.Sprintf("B%d", rowCount), p.Name)
	} else {
		return errors.New("invalid Product - validation failed")
	}       
	file.SetCellValue(sheet, fmt.Sprintf("C%d", rowCount), p.Description) 
	file.SetCellValue(sheet, fmt.Sprintf("D%d", rowCount), p.Price)       
	file.SetCellValue(sheet, fmt.Sprintf("E%d", rowCount), p.Quantity)    

	return file.SaveAs(filePath)
}

// GetProducts retrieves all products from the Excel file
func GetProducts() ([]*models.Product, error) {
	mutex.Lock()
	defer mutex.Unlock()

	file, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}

	rows, _ := file.GetRows("Sheet1")
	var products []*models.Product
	for i, row := range rows {
		if i == 0 {
			continue // Skip header row
		}
		price, _ := strconv.ParseFloat(row[3], 64)
		quantity, _ := strconv.Atoi(row[4])
		products = append(products, &models.Product{
			ID:          row[0],
			Name:        row[1],
			Description: row[2],
			Price:       price,
			Quantity:    quantity,
		})
	}
	return products, nil
}
