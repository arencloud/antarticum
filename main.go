package main

import (
	"github.com/arencloud/antarticum/internal/routes"
	"github.com/gin-gonic/gin"
)

// @title           Hospital API Service
// @version         1.0
// @description     A hospital service management API in Go using Gin framework.

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /api/v1

func main() {
	// Create a new Gin router
	router := gin.Default()

	routes.SetupPublicRoutes(router)
	routes.SetupProtectedRoutes(router)

	// Start the server
	err := router.Run(":8000")
	if err != nil {
		return
	}
}
