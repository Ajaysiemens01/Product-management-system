package test

import (
	"sync"
	"testing"
	"report-service/services"
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
func GetReportFilePath() string {
    path := os.Getenv("EXCEL_FILE_PATH")
    if path == "" {
        path = "../../data/products_report.xlsx" // Default path for local testing
    }
    fmt.Println("Using Excel file path:", path)
    return path
}
var (
testMutex sync.Mutex
filePath = GetExcelFilePath()
reportFilePath = GetReportFilePath()
)

//Test GetInventoryReport with valid credentials
func TestGetInventorReport(t *testing.T) {
	testMutex.Lock()
	defer testMutex.Unlock()

	_,err := services.GetInventoryReport(5, filePath, reportFilePath)
	assert.NoError(t, err, "Failed to get report")
}

//Test GetInventoryReport with invalid credentials
func TestGetInventorReportNegitiveLimit(t *testing.T) {
	testMutex.Lock()
	defer testMutex.Unlock()

	report,err := services.GetInventoryReport(-5, filePath, reportFilePath)
	assert.Error(t, err, "Expected failure with invalid limit")
	assert.Nil(t, report, "Expected Empty Products report")
}



