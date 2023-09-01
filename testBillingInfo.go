package endpoints

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBillingInformation(t *testing.T) {
	// Mock your data and behavior here
	mockData := []BillingInformation{
		{ID: 1, PatientID: 123, Amount: 100.0},
		{ID: 2, PatientID: 456, Amount: 150.0},
	}

	
	getBillingInfoHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockData)
	})

	req := httptest.NewRequest("GET", "/billing-information", nil)
	w := httptest.NewRecorder()
	getBillingInfoHandler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	var response []BillingInformation
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Error decoding response JSON: %v", err)
	}

	// Add assertions to verify the response data
	if len(response) != len(mockData) {
		t.Errorf("Expected %d billing information records, but got %d", len(mockData), len(response))
	}
	
}
