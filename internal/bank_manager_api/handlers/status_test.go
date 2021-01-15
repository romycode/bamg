package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckController_HealthCheck(t *testing.T) {
	// Setup
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	expected := `{"data":"¡Alive!"}`

	sut := NewHealthCheckController()

	// Assertions
	if assert.NoError(t, sut.HealthCheck(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	}
}
