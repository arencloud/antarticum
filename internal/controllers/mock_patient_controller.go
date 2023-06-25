package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Patient represents a patient entity.
type Patient struct {
	ID        string
	Name      string
	Age       int
	Gender    string
	Condition string
}

// CreateDummyPatient creates a dummy patient instance for testing purposes.
func CreateDummyPatient() *Patient {
	return &Patient{
		ID:        "123",
		Name:      "John Doe",
		Age:       30,
		Gender:    "Male",
		Condition: "Healthy",
	}
}

// MockPatientController is a mock implementation of the patient controller.
type MockPatientController struct{}

// GetAllPatients is a mock implementation of the GetAllPatients handler.
func (c *MockPatientController) GetAllPatients(ctx *gin.Context) {
	// Create a dummy patient
	dummyPatient := CreateDummyPatient()

	// Return the dummy patient as a mock response
	ctx.JSON(http.StatusOK, dummyPatient)
}

// CreatePatient is a mock implementation of the CreatePatient handler.
func (c *MockPatientController) CreatePatient(ctx *gin.Context) {
	// Create a dummy patient
	dummyPatient := CreateDummyPatient()

	// Return the dummy patient as a mock response
	ctx.JSON(http.StatusOK, dummyPatient)
}

// GetPatientByID is a mock implementation of the GetPatientByID handler.
func (c *MockPatientController) GetPatientByID(ctx *gin.Context) {
	// Create a dummy patient
	dummyPatient := CreateDummyPatient()

	// Return the dummy patient as a mock response
	ctx.JSON(http.StatusOK, dummyPatient)
}

// UpdatePatient is a mock implementation of the UpdatePatient handler.
func (c *MockPatientController) UpdatePatient(ctx *gin.Context) {
	// Create a dummy patient
	dummyPatient := CreateDummyPatient()

	// Return the dummy patient as a mock response
	ctx.JSON(http.StatusOK, dummyPatient)
}

// DeletePatient is a mock implementation of the DeletePatient handler.
func (c *MockPatientController) DeletePatient(ctx *gin.Context) {
	// Return a mock response
	ctx.JSON(http.StatusOK, gin.H{
		"message": "DeletePatient mock response",
	})
}
