package routes

import (
	"github.com/gin-gonic/gin"
)

type UniversalController interface {
	GetAll(c *gin.Context)
	Create(c *gin.Context)
	GetByID(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

func InitializeRoutes(router *gin.Engine, entityName string, controller UniversalController) {
	entityRoutes := router.Group("/" + entityName)

	entityRoutes.GET("", controller.GetAll)
	entityRoutes.POST("", controller.Create)
	entityRoutes.GET("/:id", controller.GetByID)
	entityRoutes.PUT("/:id", controller.Update)
	entityRoutes.DELETE("/:id", controller.Delete)
}
