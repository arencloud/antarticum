package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type FactoryController struct {
}

func (hc *FactoryController) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetAll hospitals"})
}

func (hc *FactoryController) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Create hospital"})
}

func (hc *FactoryController) GetByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get hospital by ID"})
}

func (hc *FactoryController) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Update hospital"})
}

func (hc *FactoryController) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Delete hospital"})
}
