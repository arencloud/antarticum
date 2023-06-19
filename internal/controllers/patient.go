package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllPatients(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Fetching all patients"})
}

func createPatient(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Creating a new patient"})
}

func getPatientByID(c *gin.Context) {
	patientID := c.Param("id")

	c.JSON(http.StatusOK, gin.H{"message": "Fetching patient with ID " + patientID})
}
