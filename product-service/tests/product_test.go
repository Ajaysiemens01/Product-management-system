package test

import (
	"product-service/models"
	"sync"
	"testing"
	"product-service/services"
	"github.com/stretchr/testify/assert"
	"log"
	"github.com/xuri/excelize/v2"
	"os"
	"fmt"
)

var testMutex sync.Mutex

//Function to get file path
func GetExcelFilePath() string {
    path := os.Getenv("EXCEL_FILE_PATH")
    if path == "" {
        path = "../../data/products.xlsx" // Default path for local testing
    }
    fmt.Println("Using Excel file path:", path)
    return path
}
var filePath = GetExcelFilePath()

func deleteTestProduct() {
	// Open the Excel file
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	sheetName := file.GetSheetName(file.GetActiveSheetIndex()) // Get active sheet name
	rows, _ := file.GetRows(sheetName)
	lastRow := len(rows)

	// Delete last row by clearing its values
	for colIdx := 1; colIdx <= len(rows[lastRow-1]); colIdx++ {
		cellName, _ := excelize.CoordinatesToCellName(colIdx, lastRow)
		file.SetCellValue(sheetName, cellName, "")
	}

	// Save the modified file
	if err := file.SaveAs(filePath); err != nil {
		log.Fatalf("Failed to save file: %v", err)
	}
}


// Test SaveProduct with valid product credentials
func TestSaveProduct(t *testing.T) {
	testMutex.Lock()
	defer testMutex.Unlock()

	product := models.Product{
		Name:        "Test Product",
		Description: "A test description",
		Price:       99.99,
		Quantity:    10,
	}

	err := services.SaveProduct(product)
	assert.NoError(t, err, "Failed to save product")
	deleteTestProduct()
}
// Test GetProduct
func TestGetProducts(t *testing.T) {
	testMutex.Lock()
	defer testMutex.Unlock()


	products, err := services.GetProducts()
	assert.NoError(t, err, "Failed to fetch products")
	assert.NotEmpty(t, products, "Products list is empty")
}


// Test SaveProduct if product is empty
func TestSaveEmptyProduct(t *testing.T) {
	testMutex.Lock()
	defer testMutex.Unlock()

	product := models.Product{}
	err := services.SaveProduct(product)
	assert.Error(t, err, "Expected failure when saving invalid product")
}

// Test SaveProduct if product already exist
func TestSaveExistedProduct(t *testing.T) {
	testMutex.Lock()
	defer testMutex.Unlock()

	product := models.Product{
		Name: "Widget A",
	Description: "A test description",
	Price:       99.99,
	Quantity:    10,}
	err := services.SaveProduct(product)
	assert.Error(t, err, "Expected failure when saving invalid product")
}
