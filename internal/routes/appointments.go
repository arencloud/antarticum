package routes

import (
	"github.com/arencloud/antarticum/internal/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeAppointmentsRoutes(router *gin.Engine) {
	patients := router.Group("/appointments")

	patients.GET("", controllers.GetAllAppointments)

	patients.POST("", controllers.CreateAppointment)

	patients.GET("/:id", controllers.GetAppointmentById)

	patients.PUT("/:id", controllers.UpdateAppointment)

	patients.DELETE("/:id", controllers.DeleteAppointment)
}
