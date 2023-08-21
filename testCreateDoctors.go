package endpoints_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateDoctor(t *testing.T) {
	// Mock your data and behavior here
	mockDoctorInput := DoctorInput{
		Name:           "Dr. Test Doctor",
		Specialization: "Test Specialization",
	}

	
	createDoctorHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var doctor DoctorInput
		if err := json.NewDecoder(r.Body).Decode(&doctor); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}


		responseDoctor := Doctor{
			ID:             "newDoctorID",
			Name:           doctor.Name,
			Specialization: doctor.Specialization,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(responseDoctor)
	})

	requestBody := `{"name": "Dr. Test Doctor", "specialization": "Test Specialization"}`
	req := httptest.NewRequest("POST", "/doctors", strings.NewReader(requestBody))
	w := httptest.NewRecorder()
	createDoctorHandler.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, w.Code)
	}

	var response Doctor
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Error decoding response JSON: %v", err)
	}

	// Add assertions to verify the response data
	if response.Name != mockDoctorInput.Name {
		t.Errorf("Expected doctor name %s, but got %s", mockDoctorInput.Name, response.Name)
	}
	
}


