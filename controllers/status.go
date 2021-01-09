package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HealthCheck() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSONBlob(http.StatusOK, []byte("{\"data\":\"¡Alive!\"}"))
	}
}
