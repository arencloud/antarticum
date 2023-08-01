package routes

import (
	"github.com/arencloud/antarticum/internal/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupPublicRoutes(router *gin.Engine) {
	public := router.Group("/api/v1")
	public.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)
	public.GET("/patients", controllers.GetUsers)
}
