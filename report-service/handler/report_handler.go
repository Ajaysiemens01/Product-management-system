package handler

import (
	"fmt"
	"net/http"
	"os"
	"report-service/services"
	"strconv"
	"sync"

	"github.com/google/jsonapi"
)

//Get the product file path
func GetExcelFilePath() string {
    path := os.Getenv("EXCEL_FILE_PATH")
    if path == "" {
        path = "../data/products.xlsx" // Default path for local testing
    }
    fmt.Println("Using Excel file path:", path)
    return path
}

//Get the product report file path
func GetReportFilePath() string {
    path := os.Getenv("REPORT_FILE_PATH")
    if path == "" {
        path = "../data/products_report.xlsx" // Default path for local testing
    }
    fmt.Println("Using Excel file path:", path)
    return path
}

var (
	filePath = GetExcelFilePath()
	reportFilePath = GetReportFilePath()
)
// RespondWithJsonApi sends the response
func RespondWithJsonApi(w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	return jsonapi.MarshalPayload(w, response)
}

// handleError is a helper function to handle errors
func handleError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(statusCode)
	var errors []*jsonapi.ErrorObject

		errors = append(errors, &jsonapi.ErrorObject{
			Status: fmt.Sprintf("%d", statusCode), 
			Title:  http.StatusText(statusCode),
			Detail: message,
		})
	if err := jsonapi.MarshalErrors(w, errors); err != nil {
		http.Error(w, "Error encoding error response", http.StatusInternalServerError)
	}
}

// Get full inventory report
func GetInventoryReportHandler(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		restockThresholdStr := r.URL.Query().Get("restock_threshold")
		restockThreshold, err := strconv.Atoi(restockThresholdStr)
		if err != nil {
			handleError(w, err.Error(), http.StatusInternalServerError)
		}
		report, err := services.GetInventoryReport(restockThreshold,filePath,reportFilePath)
		if err != nil {
			handleError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := RespondWithJsonApi(w, report); err != nil {
			handleError(w, "Failed to send response", http.StatusInternalServerError)
		}
	}()
	wg.Wait()
}
