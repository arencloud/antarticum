package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HospitalController struct {
}

func (hc *HospitalController) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetAll hospitals"})
}

func (hc *HospitalController) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Create hospital"})
}

func (hc *HospitalController) GetByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get hospital by ID"})
}

func (hc *HospitalController) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Update hospital"})
}

func (hc *HospitalController) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Delete hospital"})
}
