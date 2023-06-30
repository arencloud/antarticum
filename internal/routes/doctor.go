package routes

import (
	"github.com/arencloud/antarticum/internal/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeDoctorsRoutes(router *gin.Engine) {
	patients := router.Group("/doctors")

	patients.GET("", controllers.GetAllDoctors)

	patients.POST("", controllers.CreateDoctor)

	patients.DELETE("/:id", controllers.DeleteDoctor)
}
