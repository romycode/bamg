package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthCheckController struct{}

func NewHealthCheckController() HealthCheckController {
	return HealthCheckController{}
}

func (hc HealthCheckController) HealthCheck(c echo.Context) error {
	return c.JSONBlob(http.StatusOK, []byte("{\"data\":\"¡Alive!\"}"))
}
