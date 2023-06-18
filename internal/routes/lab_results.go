package routes

import (
	"github.com/arencloud/antarticum/internal/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeLabResultsRoutes(router *gin.Engine) {
	patients := router.Group("/lab_results")

	patients.GET("", controllers.GetAllLabResults)

	patients.POST("", controllers.CreateLabResult)

	patients.PUT("/:id", controllers.UpdateLabResult)

	patients.DELETE("/:id", controllers.DeleteLabResult)
}
