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

// Function to get file path
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

// Test saving multiple products
func TestSaveMultipleProducts(t *testing.T) {
	testMutex.Lock()
	defer testMutex.Unlock()

	testProducts := []*models.Product{
		{Name: "Test Product 1", Description: "Test Desc 1", Price: 10.0, Quantity: 5},
		{Name: "Test Product 2", Description: "Test Desc 2", Price: 20.0, Quantity: 10},
	}

	   err := services.SaveProduct(testProducts,filePath)
		assert.NoError(t, err, "Failed to save product")
    r:= len(testProducts)

    for r>0{
		deleteTestProduct()
		r--
	}
}

// Test fetching products after adding multiple products
func TestGetProducts(t *testing.T) {
	testMutex.Lock()
	defer testMutex.Unlock()

	products, err := services.GetProducts(filePath)
	assert.NoError(t, err, "Failed to fetch products")
	assert.NotEmpty(t, products, "Products list is empty")
}

// Test saving multiple empty products
func TestSaveEmptyProducts(t *testing.T) {
	testMutex.Lock()
	defer testMutex.Unlock()

	testProducts := []*models.Product{
		{},
		{},
	}
	
		err := services.SaveProduct(testProducts,filePath)
		assert.Error(t, err, "Expected failure when saving invalid product")

}

// Test saving duplicate products
func TestSaveDuplicateProducts(t *testing.T) {
	testMutex.Lock()
	defer testMutex.Unlock()
    var testProducts []*models.Product
	testProduct := &models.Product{Name: "Widget D", Description: "Test Desc", Price: 50.0, Quantity: 20}
    testProducts = append(testProducts, testProduct)
	err := services.SaveProduct(testProducts,filePath)
	assert.Error(t, err, "Expected failure when saving duplicate product")
}
