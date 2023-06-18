package routes

import (
	"github.com/arencloud/antarticum/internal/controllers"
	"github.com/gin-gonic/gin"
)

func InitializePatientsRoutes(router *gin.Engine) {
	patients := router.Group("/patients")

	patients.GET("", controllers.GetAllPatients)

	patients.POST("", controllers.CreatePatient)

	patients.GET("/:id", controllers.GetPatientByID)

	patients.PUT("/:id", controllers.UpdatePatient)
	patients.DELETE("/:id", controllers.DeletePatient)
}
