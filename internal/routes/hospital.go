package routes

import (
	"github.com/gin-gonic/gin"
	"internal/controllers"
)

func InitializeHospitalRoutes(router *gin.Engine) {
	router.GET("/departments", controllers.getAllDepartments)
}
