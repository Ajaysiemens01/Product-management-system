package test

import (
	"sync"
	"testing"
	"inventory-service/services"
	"github.com/stretchr/testify/assert"
	"inventory-service/models"
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
// Test UpdateStock with valid credentials
func TestUpdateStockValid(t *testing.T) {
	testMutex.Lock()
	defer testMutex.Unlock()
    update := &models.InventoryUpdate{
			ProductID:"0ea6502c-d8c3-4381-b5d9-96fa2c76edf4",
			Price : 50,
			StockAdded:1,
		}
	  
	err := services.UpdateStock(update,filePath)
    
	assert.NoError(t, err, "Failed to update product")
}


// Test UpdateStock with invalid credentials
func TestUpdateStockInvalid(t *testing.T) {
	testMutex.Lock()
	defer testMutex.Unlock()
	update := &models.InventoryUpdate{
		ProductID:"e4815cc3-2fd5-4a46-a4b6-8d3c5bf03dd8",
		Price : 50,
		StockAdded:-10,
	}
	err := services.UpdateStock(update,filePath)
	assert.Error(t, err, "Expected validation error for stock added")
}


