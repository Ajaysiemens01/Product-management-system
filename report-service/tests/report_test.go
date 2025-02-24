package test

import (
	"sync"
	"testing"
	"report-service/services"
	"github.com/stretchr/testify/assert"
)

var testMutex sync.Mutex

//Test GetInventoryReport with valid credentials
func TestGetInventorReport(t *testing.T) {
	testMutex.Lock()
	defer testMutex.Unlock()

	_,err := services.GetInventoryReport(5)
	assert.NoError(t, err, "Failed to get report")
}

//Test GetInventoryReport with invalid credentials
func TestGetInventorReportNegitiveLimit(t *testing.T) {
	testMutex.Lock()
	defer testMutex.Unlock()

	report,err := services.GetInventoryReport(-5)
	assert.Error(t, err, "Expected failure with invalid limit")
	assert.Nil(t, report, "Expected Empty Products report")
}



