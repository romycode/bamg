package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthCheckHandler struct{}

func NewHealthCheckController() HealthCheckHandler {
	return HealthCheckHandler{}
}

func (hc HealthCheckHandler) HealthCheck(c echo.Context) error {
	return c.JSONBlob(http.StatusOK, []byte("{\"data\":\"¡Alive!\"}"))
}
