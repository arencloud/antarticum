package main

import (
	"os"
	"testing"

	"github.com/arencloud/antarticum/internal/controllers"
	"github.com/arencloud/antarticum/internal/middlewares"
	"github.com/arencloud/antarticum/internal/routes"
	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()

	mockPatientController := &controllers.MockPatientController{}

	routes.InitializePatientsRoutesTest(router, mockPatientController)

	router.Use(middlewares.MockJwtAuthMiddleware())

	exitCode := m.Run()

	os.Exit(exitCode)
}
