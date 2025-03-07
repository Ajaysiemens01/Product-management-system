package handler

import (
	"fmt"
	"net/http"
	"os"
	"product-service/models"
	"product-service/services"
	"reflect"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/google/jsonapi"
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

func ParseRequestBody(r *http.Request) ([]*models.Product, error) {
	// Decode JSON:API payload manually
	nodes, err := jsonapi.UnmarshalManyPayload(r.Body, reflect.TypeOf(new(models.Product)))
	if err != nil {
		return nil, err
	}
	var products []*models.Product
	// Convert nodes to []*models.Product
	for _, node := range nodes {
		if product, ok := node.(*models.Product); ok {
			products = append(products, product)
		} else {
			return nil, fmt.Errorf("failed to parse product")
		}
	}
	return products, nil
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
		products, err := services.GetProducts(filePath)
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
	var products []*models.Product
	var err error

	if products, err = ParseRequestBody(r); err != nil {
		handleError(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, p := range products {
			if err := validator.New().Struct(p); err != nil {
				handleError(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
		if err := services.SaveProduct(products, filePath); err != nil {
			handleError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)

		response := &models.Response{
			Message: "products added successfully",
		}
		if err := RespondWithJsonApi(w, response); err != nil {
			handleError(w, "Failed to respond with products", http.StatusInternalServerError)
			return
		}
	}()

	wg.Wait()
}
