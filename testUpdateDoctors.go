package endpoints_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUpdateDoctor(t *testing.T) {
	// Mock your data and behavior here
	mockDoctorID := "123" // Replace with an actual doctor ID
	mockDoctorInput := DoctorInput{
		Name:           "Updated Dr. John Doe",
		Specialization: "Updated Cardiology",
	}

	// Replace the following with your actual implementation of UpdateDoctor
	updateDoctorHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract doctor ID from the request URL
		doctorIDToUpdate := r.URL.Path[len("/doctors/"):] // Assuming the path is "/doctors/{doctorId}"

		// Compare the extracted ID with the mock data
		if doctorIDToUpdate != mockDoctorID {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// Parse the request body
		var input DoctorInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Perform the update logic (this could involve database interactions, etc.)
		// Here, we simply update the mock data
		mockDoctorInput.Name = input.Name
		mockDoctorInput.Specialization = input.Specialization

		// Return the updated doctor as the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockDoctorInput)
	})

	// Create a request with mock doctor input
	requestBody := `{"name": "Updated Dr. John Doe", "specialization": "Updated Cardiology"}`
	req := httptest.NewRequest("PUT", "/doctors/"+mockDoctorID, strings.NewReader(requestBody))
	w := httptest.NewRecorder()
	updateDoctorHandler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	var response DoctorInput
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Error decoding response JSON: %v", err)
	}

	// Add assertions to verify the response data
	if response.Name != mockDoctorInput.Name {
		t.Errorf("Expected name %s, but got %s", mockDoctorInput.Name, response.Name)
	}
	if response.Specialization != mockDoctorInput.Specialization {
		t.Errorf("Expected specialization %s, but got %s", mockDoctorInput.Specialization, response.Specialization)
	}
	
}
