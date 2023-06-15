package routes

import (
	"github.com/gin-gonic/gin"
	"internal/controllers"
	//"../controllers"
)

func InitializePatientsRoutes(router *gin.Engine) {
	patients := router.Group("/patients")

	patients.GET("", controllers.getAllPatients)

	patients.POST("", controllers.createPatient)

	patients.GET("/:id", controllers.getPatientByID)
}
