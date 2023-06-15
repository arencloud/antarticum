package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arencloud/antarticum/internal/routes"
	"github.com/gin-gonic/gin"
)

type Plugin interface {
	Setup()
}

type ExamplePlugin struct{}

func (p *ExamplePlugin) Setup() {
	fmt.Println("Plugin setup")
}

func main() {
	router := gin.Default()

	routes.InitializePatientsRoutes(router)
	routes.InitializeHospitalRoutes(router)

	plugins := []Plugin{
		&ExamplePlugin{},
	}

	for _, plugin := range plugins {
		plugin.Setup()
	}

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
