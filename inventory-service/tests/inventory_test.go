package test

import (
	"sync"
	"testing"
	"inventory-service/services"
	"github.com/stretchr/testify/assert"
)

var testMutex sync.Mutex

// Test UpdateStock with valid credentials
func TestUpdateStock(t *testing.T) {
	testMutex.Lock()
	defer testMutex.Unlock()

	err := services.UpdateStock("e4815cc3-2fd5-4a46-a4b6-8d3c5bf03dd8", 5)
	assert.NoError(t, err, "Failed to update product")
	err = services.UpdateStock("e4815cc3-2fd5-4a46-a4b6-8d3c5bf03dd8", -5)
	assert.NoError(t, err, "Failed to update product")
}


