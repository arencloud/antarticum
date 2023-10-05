package main

import (
	"github.com/arencloud/antarticum/internal/controllers"
	"github.com/arencloud/antarticum/internal/middlewares"
	"github.com/arencloud/antarticum/internal/routes"
	"github.com/gin-gonic/gin"
	"os"
	"testing"
)

func TestUniversal(m *testing.M) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()

	hospitalController := &controllers.HospitalController{}
	routes.InitializeRoutes(r, "hospitals", hospitalController)

	shopController := &controllers.ShopController{}
	routes.InitializeRoutes(r, "shops", shopController)

	factoryController := &controllers.FactoryController{}
	routes.InitializeRoutes(r, "chemical factory", factoryController)

	r.Use(middlewares.MockJwtAuthMiddleware())

	exitCode := m.Run()

	os.Exit(exitCode)
}
