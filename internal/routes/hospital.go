package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeHospitalRoutes(router *gin.Engine) {
	router.GET("/departments", getAllDepartments)
}

func getAllDepartments(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Fetching all departments"})
}
