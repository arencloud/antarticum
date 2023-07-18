package middlewares

import (
	"net/http"

	"github.com/arencloud/antarticum/internal/helpers"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := helpers.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		errRole := helpers.AdminRoleValid(c)
		if errRole != nil {
			c.String(http.StatusUnauthorized, "Only Administrator is allowed to perform this action")
			c.Abort()
			return
		}
		c.Next()
	}
}
