package endpoints_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetLabResults(t *testing.T) {
	// Mock your data and behavior here
	mockData := []LabResult{
		{ID: "1", PatientID: "123", Result: "Positive"},
		{ID: "2", PatientID: "456", Result: "Negative"},
	}

	
	getLabResultsHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockData)
	})

	req := httptest.NewRequest("GET", "/lab-results", nil)
	w := httptest.NewRecorder()
	getLabResultsHandler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	var response []LabResult
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Error decoding response JSON: %v", err)
	}

	// Add assertions to verify the response data
	if len(response) != len(mockData) {
		t.Errorf("Expected %d lab results, but got %d", len(mockData), len(response))
	}

}
