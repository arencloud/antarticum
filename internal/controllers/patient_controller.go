package controllers

import "github.com/gin-gonic/gin"

type PatientController interface {
	GetAllPatients(ctx *gin.Context)
	CreatePatient(ctx *gin.Context)
	GetPatientByID(ctx *gin.Context)
	UpdatePatient(ctx *gin.Context)
	DeletePatient(ctx *gin.Context)
}
