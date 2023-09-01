package endpoints_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetPatients(t *testing.T) {
	// Mock your data and behavior here
	mockData := []Patient{
		{ID: "1", FullName: "John Doe", Age: 30, Address: "123 Main St"},
		{ID: "2", FullName: "Jane Doe", Age: 25, Address: "456 Elm St"},
	}

	
	getPatientsHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockData)
	})

	req := httptest.NewRequest("GET", "/patients", nil)
	w := httptest.NewRecorder()
	getPatientsHandler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	var response []Patient
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Error decoding response JSON: %v", err)
	}

	// Add assertions to verify the response data
	if len(response) != len(mockData) {
		t.Errorf("Expected %d patients, but got %d", len(mockData), len(response))
	}
	
}
