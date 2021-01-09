package routes

import (
	"net/http"

	"github.com/romycode/bank-manager/controllers"
)

var statusRoutes = []Endpoint{
	{
		Method:  http.MethodGet,
		Group:   "",
		Path:    "/health-check",
		Handler: controllers.HealthCheck(),
	},
}
