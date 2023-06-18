package main

import (
	"log"
	"net/http"

	"github.com/arencloud/antarticum/internal/middlewares"
	"github.com/arencloud/antarticum/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Pluggable support
	router.Use(middlewares.JwtAuthMiddleware())

	routes.InitializePatientsRoutes(router)
	routes.InitializeDoctorsRoutes(router)
	routes.InitializeLabResultsRoutes(router)
	routes.InitializeAppointmentsRoutes(router)
	routes.InitializeBillingInformationRoutes(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Successful",
		})
	})

	// Start the server
	err := router.Run(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
