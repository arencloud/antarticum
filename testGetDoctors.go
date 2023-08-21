package endpoints_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetDoctors(t *testing.T) {
	// Mock your data and behavior here
	mockData := []Doctor{
		{ID: "1", Name: "Dr. John Doe", Specialization: "Cardiology"},
		{ID: "2", Name: "Dr. Jane Doe", Specialization: "Dermatology"},
	}

	
	getDoctorsHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockData)
	})

	req := httptest.NewRequest("GET", "/doctors", nil)
	w := httptest.NewRecorder()
	getDoctorsHandler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	var response []Doctor
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Error decoding response JSON: %v", err)
	}

	// Add assertions to verify the response data
	if len(response) != len(mockData) {
		t.Errorf("Expected %d doctors, but got %d", len(mockData), len(response))
	}
	
}
