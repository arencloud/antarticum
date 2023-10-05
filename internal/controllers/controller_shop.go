package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ShopController struct {
}

func (hc *ShopController) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetAll hospitals"})
}

func (hc *ShopController) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Create hospital"})
}

func (hc *ShopController) GetByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get hospital by ID"})
}

func (hc *ShopController) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Update hospital"})
}

func (hc *ShopController) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Delete hospital"})
}
