package handlers

import (
	"fmt"
	"inventory-service/models"
	"inventory-service/services"
	"net/http"
	"sync"
	"github.com/google/jsonapi"
	"github.com/gorilla/mux"
)

// ParseRequestBody parses the request body
func ParseRequestBody(r *http.Request) (*models.InventoryUpdate, error) {
	var request models.InventoryUpdate
	if err := jsonapi.UnmarshalPayload(r.Body, &request); err != nil {
		return nil, err
	}
	return &request, nil
}

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
var wg sync.WaitGroup

// UpdateStockHandler handles stock updates concurrently
func UpdateStockHandler(w http.ResponseWriter, r *http.Request) {
    var update *models.InventoryUpdate
    id := mux.Vars(r)["product_id"]
    var err error

    update, err = ParseRequestBody(r)
    if err != nil {
        handleError(w, "Invalid request", http.StatusBadRequest)
        return
    }
    update.ProductID = id

    wg.Add(1)
    go func() {
        defer wg.Done()
        if err := services.UpdateStock(update.ProductID, update.Change); err != nil {
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
