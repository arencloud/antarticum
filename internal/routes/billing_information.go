package routes

import (
	"github.com/arencloud/antarticum/internal/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeBillingInformationRoutes(router *gin.Engine) {
	patients := router.Group("/billing_information")

	patients.GET("", controllers.GetAllBillingInformation)

	patients.POST("", controllers.CreateBillingInformation)

	patients.PUT("/:id", controllers.UpdateBillingInformation)

	patients.DELETE("/:id", controllers.DeleteBillingInformation)
}
