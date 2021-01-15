package routes

import (
	"net/http"

	"github.com/romycode/bank-manager/internal/bank_manager_api/database/repositories"
	"github.com/romycode/bank-manager/internal/bank_manager_api/handlers"
)

var accountRepository = repositories.NewSqliteAccountRepository(db)
var accountController = handlers.NewAccountController(accountRepository)

var accountsPath = "/v1/accounts"
var accountRoutes = []Endpoint{
	{
		Method:  http.MethodGet,
		Group:   accountsPath,
		Path:    "",
		Handler: accountController.GetAllAccounts,
	},
	{
		Method:  http.MethodPost,
		Group:   accountsPath,
		Path:    "",
		Handler: accountController.CreateAccount,
	},
	{
		Method:  http.MethodDelete,
		Group:   accountsPath,
		Path:    "/:id",
		Handler: accountController.DeleteAccount,
	},
}
