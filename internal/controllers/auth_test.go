package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestVerifyPassword(t *testing.T) {
	// Test case 1: Correct password
	password := "password123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	err := VerifyPassword(password, string(hashedPassword))
	if err != nil {
		t.Errorf("Expected password verification to succeed, but got error: %v", err)
	}

	// Test case 2: Incorrect password
	incorrectPassword := "wrongpassword"

	err = VerifyPassword(incorrectPassword, string(hashedPassword))
	if err == nil {
		t.Error("Expected password verification to fail, but got no error")
	}
}

func TestLogin(t *testing.T) {
	router := gin.Default()

	router.POST("/login", Login)

	payload := `{"email": "star@gmail.com", "password": "123456stacy"}`

	req, err := http.NewRequest("POST", "/login", strings.NewReader(payload))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, res.Code)
	}
}

func TestLoginCheck(t *testing.T) {
	user := RegisterInput{
		ID:       1,
		Name:     "Stacy Bale",
		Age:      22,
		Email:    "star@gmail.com",
		Address:  "221 BB",
		Password: "$2a$10$9IgfVRe/YqSVTcHJHCsTiOxZJrdTtXx5ylf5wvqC1pwtiWqUK1pUe", // hashed password for "123456stacy"
	}

	users = append(users, user)

	// Test case 1: Correct email and password
	token, err := LoginCheck("star@gmail.com", "123456stacy")
	if err != nil {
		t.Errorf("Expected login check to succeed, but got error: %v", err)
	}

	// TODO: Verify the generated token
	fmt.Println(token)

	// Test case 2: Incorrect email
	_, err = LoginCheck("wrongemail@gmail.com", "123456stacy")
	if err == nil {
		t.Error("Expected login check to fail for incorrect email, but got no error")
	}

	// Test case 3: Incorrect password
	_, err = LoginCheck("star@gmail.com", "wrongpassword")
	if err == nil {
		t.Error("Expected login check to fail for incorrect password, but got no error")
	}
}

func TestRegister(t *testing.T) {
	router := gin.Default()

	router.POST("/register", Register)

	payload := map[string]interface{}{
		"fullName":  "Stacy Bale",
		"age":       float64(22), // Convert to float64
		"email":     "star@gmail.com",
		"address":   "221 BB",
		"password":  "123456stacy",
		"patientId": float64(1), // Convert to float64
	}

	jsonPayload, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonPayload))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, res.Code)
	}

	var response map[string]interface{}
	err = json.Unmarshal(res.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	expectedMessage := payload
	actualMessage := response["message"].(map[string]interface{})

	if len(expectedMessage) != len(actualMessage) {
		t.Errorf("Expected message to be %v but got %v", expectedMessage, actualMessage)
	}

	for key, expectedValue := range expectedMessage {
		actualValue, ok := actualMessage[key]
		if !ok || !reflect.DeepEqual(expectedValue, actualValue) {
			t.Errorf("Expected message[%s] to be %v but got %v", key, expectedValue, actualValue)
		}
	}
}
