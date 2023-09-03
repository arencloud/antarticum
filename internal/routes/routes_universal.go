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

	entityRoutes.GET("/:org/:action/:any_attributes", controller.GetAll)
	entityRoutes.POST("/:org/:action/:any_attributes", controller.Create)
	entityRoutes.GET("/:org/:action/:any_attributes/:id", controller.GetByID)
	entityRoutes.PUT("/:org/:action/:any_attributes/:id", controller.Update)
	entityRoutes.DELETE("/:org/:action/:any_attributes/:id", controller.Delete)

}
