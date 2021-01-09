package routes

import (
	"net/http"

	"github.com/romycode/bank-manager/controllers"
)

var usersPath = "/v1/users"

var userRoutes = []Endpoint{
	{
		Method:  http.MethodGet,
		Group:   usersPath,
		Path:    "",
		Handler: controllers.GetAllUsers(),
	},
	{
		Method:  http.MethodPost,
		Group:   usersPath,
		Path:    "",
		Handler: controllers.CreateUser(),
	},
	{
		Method:  http.MethodDelete,
		Group:   usersPath,
		Path:    "/:id",
		Handler: controllers.DeleteUser(),
	},
}
