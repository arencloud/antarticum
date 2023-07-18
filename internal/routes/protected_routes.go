package routes

import (
	"github.com/arencloud/antarticum/internal/controllers"
	"github.com/arencloud/antarticum/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupProtectedRoutes(router *gin.Engine) {
	protected := router.Group("/api/v1/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)
}
