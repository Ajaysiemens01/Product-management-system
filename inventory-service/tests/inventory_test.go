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
			ProductID:"7a95839e-7075-40c8-9c46-a5990084fb46",
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
		ProductID:"7a95839e-7075-40c8-9c46-a5990084fb46",
		Price : 50,
		StockAdded:-10,
	}
	err := services.UpdateStock(update,filePath)
	assert.Error(t, err, "Expected validation error for stock added")
}


