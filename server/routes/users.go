package routes

import (
	"net/http"

	"github.com/romycode/bank-manager/controllers"
	"github.com/romycode/bank-manager/database"
	"github.com/romycode/bank-manager/database/repositories"
)

var db = database.GetConnection()
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
