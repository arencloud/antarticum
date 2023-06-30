package routes

import (
	"github.com/arencloud/antarticum/internal/controllers"
	"github.com/gin-gonic/gin"
)

func InitializePatientsRoutesTest(router *gin.Engine, patientController controllers.PatientController) {
	patients := router.Group("/patients")

	patients.GET("", patientController.GetAllPatients)

	patients.POST("", patientController.CreatePatient)

	patients.GET("/:id", patientController.GetPatientByID)

	patients.PUT("/:id", patientController.UpdatePatient)
	patients.DELETE("/:id", patientController.DeletePatient)
}
