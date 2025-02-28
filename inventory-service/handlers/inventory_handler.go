package handlers

import (
	"inventory-service/models"
	"inventory-service/services"
	"sync"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/go-playground/validator/v10"
	"os"
	"fmt"
)

var wg sync.WaitGroup

func GetExcelFilePath() string {
    path := os.Getenv("EXCEL_FILE_PATH")
    if path == "" {
        path = "../data/products.xlsx" // Default path for local testing
    }
    fmt.Println("Using Excel file path:", path)
    return path
}
 var filePath = GetExcelFilePath()

// UpdateProductDetails handles product updates concurrently
func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	var update *models.InventoryUpdate
    id := mux.Vars(r)["product_id"]
    var err error
    update, err = ParseRequestBody(r)
    if err != nil {
        handleError(w, "Invalid request", http.StatusBadRequest)
        return
    }
    
   // Validate input
   if err := validator.New().Struct(update); err != nil {
	errs := BundleValidationErrors(err,update)
	if errs != nil{
		handleErrors(w, errs, http.StatusBadRequest)
		return
	}
    }
    if id != "" {
		update.ProductID = id
	}

    wg.Add(1)
    go func() {
        defer wg.Done()
        if err := services.UpdateStock(update,filePath); err != nil {
            handleError(w, err.Error(), http.StatusBadRequest)
            return
        }
        response := models.Response{
            Message: "Stock updated successfully",
        }
        if err := RespondWithJsonApi(w, &response); err != nil {
            handleError(w, "Failed to respond with products", http.StatusInternalServerError)
        }
    }()

    wg.Wait()
}

