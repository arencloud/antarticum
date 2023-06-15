package routes

import (
	"github.com/gin-gonic/gin"
)

func InitializeHospitalRoutes(router *gin.Engine) {
	router.GET("/departments", controllers.getAllDepartments)
}
