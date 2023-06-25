package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arencloud/antarticum/internal/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPatientRoutes(t *testing.T) {
	router := gin.Default()

	mockPatientController := &controllers.MockPatientController{}

	InitializePatientsRoutesTest(router, mockPatientController)

	req, _ := http.NewRequest("GET", "/patients", nil)
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

}
