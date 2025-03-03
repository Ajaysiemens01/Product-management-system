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
	id := "0ea6502c-d8c3-4381-b5d9-96fa2c76edf4"
	err := services.UpdateStock(id,1,filePath)
	assert.NoError(t, err, "Expected no error for valid input")
}

//Test UpdateStock with invalid credentials
func TestUpdateStockInvalid(t *testing.T) {
	testMutex.Lock()
	defer testMutex.Unlock()
	id := "0ea6502c-d8c3-4381-b5d9-96fa2c76edf4"
	err := services.UpdateStock(id,-1,filePath)
	assert.Error(t, err, "Expected Validation error for change")
}




