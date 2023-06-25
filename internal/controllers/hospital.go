package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllDepartments(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Fetching all departments"})
}
