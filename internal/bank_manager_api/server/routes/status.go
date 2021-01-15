package routes

import (
	"net/http"

	"github.com/romycode/bank-manager/internal/bank_manager_api/handlers"
)

var healthCheckController = handlers.NewHealthCheckController()

var statusRoutes = []Endpoint{
	{
		Method:  http.MethodGet,
		Group:   "",
		Path:    "/health-check",
		Handler: healthCheckController.HealthCheck,
	},
}
