package routes

import (
	"net/http"

	"github.com/romycode/bank-manager/internal/bank_manager_api/controllers"
)

var healthCheckController = controllers.NewHealthCheckController()

var statusRoutes = []Endpoint{
	{
		Method:  http.MethodGet,
		Group:   "",
		Path:    "/health-check",
		Handler: healthCheckController.HealthCheck,
	},
}
