package middlewares

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"net/http"
	"strings"
)

type CommonDataStructure struct {
	ID       int    `json:"patientId" xml:"patientId" yaml:"patientId" binding:"required"`
	Name     string `json:"fullName" xml:"fullName" yaml:"fullName" binding:"required"`
	Age      int    `json:"age" xml:"age" yaml:"age"`
	Address  string `json:"address" xml:"address" yaml:"address"`
	Email    string `json:"email" xml:"email" yaml:"email" binding:"required"`
	Password string `json:"password" xml:"password" yaml:"password" binding:"required"`
	Role     string `json:"role" xml:"role" yaml:"role" binding:"required"`
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
