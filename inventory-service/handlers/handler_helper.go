package handlers

import (
	"fmt"
	"inventory-service/models"
	"net/http"
	"github.com/google/jsonapi"
	"github.com/go-playground/validator/v10"
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

// handleError is a helper function to handle errors
func handleErrors(w http.ResponseWriter, messages []string, statusCode int) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(statusCode)

	var errors []*jsonapi.ErrorObject
	for _,message := range messages{
		errors = append(errors, &jsonapi.ErrorObject{
			Status: fmt.Sprintf("%d", statusCode), 
			Title:  http.StatusText(statusCode),
			Detail: message,
		})
	}
	if err := jsonapi.MarshalErrors(w, errors); err != nil {
		http.Error(w, "Error encoding error response", http.StatusInternalServerError)
	}
}

// BundleValidationErrors maps validation errors to custom messages
func BundleValidationErrors(err error,update *models.InventoryUpdate) []string {
	var errorMessages []string

	// Loop through validation errors
	for _, e := range err.(validator.ValidationErrors) {
		switch e.Field() {
		case "Price":
			if e.Tag() == "gt" && update.Price < 0{
				errorMessages = append(errorMessages, "Price must be greater than 0")
			}
		case "Quantity":
			if e.Tag() == "gt" && update.Quantity < 0{
				errorMessages = append(errorMessages, "Quantity must be greater than 0")
			}
		case "StockAdded":
			if e.Tag() == "gt" && update.StockAdded < 0 {
				errorMessages = append(errorMessages, "Stock added must be greater than 0")
			}
		default:
			errorMessages = append(errorMessages, fmt.Sprintf("%s is invalid", e.Field()))
		}
	}
  return errorMessages
}
