package routes

import (
	"net/http"

	"github.com/romycode/bank-manager/controllers"
)

var accountsPath = "/v1/accounts"

var accountRoutes = []Endpoint{
	{
		Method:  http.MethodGet,
		Group:   accountsPath,
		Path:    "",
		Handler: controllers.GetAllAccounts,
	},
	{
		Method:  http.MethodPost,
		Group:   accountsPath,
		Path:    "",
		Handler: controllers.CreateAccount,
	},
	{
		Method:  http.MethodDelete,
		Group:   accountsPath,
		Path:    "/:id",
		Handler: controllers.DeleteAccount,
	},
}
