package main

import (
	"github.com/arencloud/antarticum/internal/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func TestRoutes(t *testing.T) {
	router := gin.Default()

	routes.InitializePatientsRoutes(router)
	routes.InitializeDoctorsRoutes(router)
	routes.InitializeLabResultsRoutes(router)
	routes.InitializeAppointmentsRoutes(router)
	routes.InitializeBillingInformationRoutes(router)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
}
