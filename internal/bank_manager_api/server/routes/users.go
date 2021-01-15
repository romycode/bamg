package routes

import (
	"net/http"

	"github.com/romycode/bank-manager/internal/bank_manager_api/controllers"
	"github.com/romycode/bank-manager/internal/bank_manager_api/database/repositories"
)

var userRepository = repositories.NewSqliteUserRepository(db)
var userController = controllers.NewUserController(userRepository)

var usersPath = "/v1/users"
var userRoutes = []Endpoint{
	{
		Method:  http.MethodGet,
		Group:   usersPath,
		Path:    "",
		Handler: userController.GetAllUsers,
	},
	{
		Method:  http.MethodPost,
		Group:   usersPath,
		Path:    "",
		Handler: userController.CreateUser,
	},
	{
		Method:  http.MethodDelete,
		Group:   usersPath,
		Path:    "/:id",
		Handler: userController.DeleteUser,
	},
}