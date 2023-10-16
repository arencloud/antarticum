package middlewares

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"net/http"
	"strings"
)

type CommonDataStructure struct {
	ID       int    `json:"patientId" binding:"required"`
	Name     string `json:"fullName" binding:"required"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

func FormatMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		format := strings.ToLower(c.DefaultQuery("format", "json"))
		switch format {
		case "json":
			c.Request.Header.Set("Accept", "application/json")
		case "xml":
			c.Request.Header.Set("Accept", "application/xml")
		case "yaml":
			c.Request.Header.Set("Accept", "application/yaml")
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported format"})
			c.Abort()
			return
		}

		c.Next()

		switch format {
		case "json":
			// Do nothing, response is already in JSON format
		case "xml":
			var data CommonDataStructure
			if err := c.ShouldBindJSON(&data); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
			xmlData, err := xml.Marshal(data)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
			c.Data(http.StatusOK, "application/xml", xmlData)
		case "yaml":
			var data CommonDataStructure
			if err := c.ShouldBindJSON(&data); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
			yamlData, err := yaml.Marshal(data)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
			c.Data(http.StatusOK, "application/yaml", yamlData)
		}
	}
}
