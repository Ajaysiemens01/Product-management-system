package test

import (
	"sync"
	"testing"
	"purchase-service/services"
	"github.com/stretchr/testify/assert"
	"os"
	"fmt"
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
testMutex sync.Mutex
filePath = GetExcelFilePath()
)

//Test UpdateStock with valid credentials
func TestUpdateStockValid(t *testing.T) {
	testMutex.Lock()
	defer testMutex.Unlock()
	id := "7a95839e-7075-40c8-9c46-a5990084fb46"
	err := services.UpdateStock(id,1,filePath)
	assert.NoError(t, err, "Expected no error for valid input")
}

//Test UpdateStock with invalid credentials
func TestUpdateStockInvalid(t *testing.T) {
	testMutex.Lock()
	defer testMutex.Unlock()
	id := "7a95839e-7075-40c8-9c46-a5990084fb46"
	err := services.UpdateStock(id,-1,filePath)
	assert.Error(t, err, "Expected Validation error for change")
}




