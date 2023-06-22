package main

import (
	"github.com/arencloud/antarticum/internal/controllers"
	"github.com/arencloud/antarticum/internal/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	public := router.Group("/api/v1")

	public.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)
	public.GET("/patients", controllers.GetUsers)

	protected := router.Group("/api/v1/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	// Start the server
	err := router.Run(":8000")
	if err != nil {
		return
	}
}
