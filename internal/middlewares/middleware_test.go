package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func MockJwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func TestJwtAuthMiddleware(t *testing.T) {
	router := gin.Default()

	router.GET("/protected", MockJwtAuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Protected route",
		})
	})

	req, _ := http.NewRequest("GET", "/protected", nil)
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusUnauthorized, res.Code)

	reqWithToken, _ := http.NewRequest("GET", "/protected", nil)
	reqWithToken.Header.Set("Authorization", "Bearer YOUR_JWT_TOKEN")
	resWithToken := httptest.NewRecorder()

	router.ServeHTTP(resWithToken, reqWithToken)

	assert.Equal(t, http.StatusOK, resWithToken.Code)

}
