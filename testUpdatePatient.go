package endpoints_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Adapt this template for the UpdatePatient endpoint
func TestUpdatePatient(t *testing.T) {
	// Mock your data and behavior here
	mockPatientID := "123" // Replace with an actual patient ID
	mockPatientInput := PatientInput{
		FullName: "Updated John Doe",
		Age:      35,
		Address:  "Updated Address",
	}

	updatePatientHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract patient ID from the request URL
		patientIDToUpdate := r.URL.Path[len("/patients/"):] // Assuming the path is "/patients/{patientId}"

		// Compare the extracted ID with the mock data
		if patientIDToUpdate != mockPatientID {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// Parse the request body
		var input PatientInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Perform the update logic (this could involve database interactions, etc.)
		// Here, we simply update the mock data
		mockPatientInput.FullName = input.FullName
		mockPatientInput.Age = input.Age
		mockPatientInput.Address = input.Address

		// Return the updated patient as the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockPatientInput)
	})

	// Create a request with mock patient input
	requestBody := `{"fullName": "Updated John Doe", "age": 35, "address": "Updated Address"}`
	req := httptest.NewRequest("PUT", "/patients/"+mockPatientID, strings.NewReader(requestBody))
	w := httptest.NewRecorder()
	updatePatientHandler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	var response PatientInput
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Error decoding response JSON: %v", err)
	}

	// Add assertions to verify the response data
	if response.FullName != mockPatientInput.FullName {
		t.Errorf("Expected full name %s, but got %s", mockPatientInput.FullName, response.FullName)
	}
	if response.Age != mockPatientInput.Age {
		t.Errorf("Expected age %d, but got %d", mockPatientInput.Age, response.Age)
	}
	if response.Address != mockPatientInput.Address {
		t.Errorf("Expected address %s, but got %s", mockPatientInput.Address, response.Address)
	}
	
}

