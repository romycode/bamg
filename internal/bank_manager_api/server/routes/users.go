package routes

import (
	"net/http"

	"github.com/romycode/bank-manager/internal/bank_manager_api/database/repositories"
	"github.com/romycode/bank-manager/internal/platform/server/handler/health"
)

var ac = repositories.NewSqliteAccountRepository(db)
var userRepository = repositories.NewSqliteUserRepository(db, ac)
var userController = health.NewUserController(userRepository)

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
