package endpoints

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAuthenticateUser(t *testing.T) {
	// Mock your data and behavior here
	mockUserCredentials := UserCredentials{
		Username: "testuser",
		Password: "testpassword",
	}

	// Replace the following with your actual implementation of AuthenticateUser
	authHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract user credentials from the request body
		var creds UserCredentials
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Compare credentials with mock data
		if creds.Username != mockUserCredentials.Username || creds.Password != mockUserCredentials.Password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Return a mock auth token
		authToken := AuthToken{Token: "mockToken"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(authToken)
	})

	// Create a request with mock user credentials
	requestBody := `{"username": "testuser", "password": "testpassword"}`
	req := httptest.NewRequest("POST", "/authentication", strings.NewReader(requestBody))
	w := httptest.NewRecorder()
	authHandler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	var response AuthToken
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Error decoding response JSON: %v", err)
	}

	// Add assertions to verify the response data
	if response.Token != "mockToken" {
		t.Errorf("Expected auth token %s, but got %s", "mockToken", response.Token)
	}
	
}
