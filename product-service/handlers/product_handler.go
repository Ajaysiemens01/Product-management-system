package handlers

import (
	"fmt"
	"net/http"
	"product-service/models"
	"product-service/services"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/google/jsonapi"
)

var wg sync.WaitGroup

// ParseRequestBody parses the request body
func ParseRequestBody(r *http.Request) (*models.Product, error) {
	var request models.Product
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

// Get all products
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {

	wg.Add(1)
	go func() {
		defer wg.Done()
		products, err := services.GetProducts()
		if err != nil {
			handleError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := RespondWithJsonApi(w, products); err != nil {
			handleError(w, "Failed to respond with products", http.StatusInternalServerError)
			return
		}
	}()
	wg.Wait()
}

// Add a new product
func AddProductHandler(w http.ResponseWriter, r *http.Request) {
	//var wg sync.WaitGroup
	var p *models.Product
	var err error

	if p, err = ParseRequestBody(r); err != nil {
		handleError(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if err := validator.New().Struct(p); err != nil {
		handleError(w, err.Error(), http.StatusBadRequest)
		return
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := services.SaveProduct(*p); err != nil {
			handleError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)

		response := &models.Response{
			Message: "product added successfully",
		}
		if err := RespondWithJsonApi(w, response); err != nil {
			handleError(w, "Failed to respond with products", http.StatusInternalServerError)
			return
		}
	}()
	wg.Wait()
}
