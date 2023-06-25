package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Patient struct {
	ID        string
	Name      string
	Age       int
	Gender    string
	Condition string
}

func CreateDummyPatient() *Patient {
	return &Patient{
		ID:        "123",
		Name:      "John Doe",
		Age:       30,
		Gender:    "Male",
		Condition: "Healthy",
	}
}

type MockPatientController struct{}

func (c *MockPatientController) GetAllPatients(ctx *gin.Context) {
	dummyPatient := CreateDummyPatient()

	ctx.JSON(http.StatusOK, dummyPatient)
}

func (c *MockPatientController) CreatePatient(ctx *gin.Context) {
	dummyPatient := CreateDummyPatient()

	ctx.JSON(http.StatusOK, dummyPatient)
}

func (c *MockPatientController) GetPatientByID(ctx *gin.Context) {
	dummyPatient := CreateDummyPatient()

	ctx.JSON(http.StatusOK, dummyPatient)
}

func (c *MockPatientController) UpdatePatient(ctx *gin.Context) {
	dummyPatient := CreateDummyPatient()

	ctx.JSON(http.StatusOK, dummyPatient)
}

func (c *MockPatientController) DeletePatient(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "DeletePatient mock response",
	})
}
