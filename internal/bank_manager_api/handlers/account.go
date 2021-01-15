package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/romycode/bank-manager/internal/bank_manager_api/errors"
	"github.com/romycode/bank-manager/internal/bank_manager_api/models"
)

type AccountHandler struct {
	repository models.AccountRepository
}

func NewAccountController(repository models.AccountRepository) AccountHandler {
	return AccountHandler{repository: repository}
}

func (ac AccountHandler) GetAllAccounts(c echo.Context) error {
	a, _ := json.Marshal(ac.repository.All())

	return c.JSONBlob(
		http.StatusOK,
		a,
	)
}

func (ac AccountHandler) CreateAccount(c echo.Context) error {
	a := new(models.Account)
	err := c.Bind(a)
	errors.HandleError(err)
	ac.repository.Save(*a)

	res, _ := json.Marshal(a)
	return c.JSONBlob(
		http.StatusCreated,
		res,
	)
}

func (ac AccountHandler) DeleteAccount(c echo.Context) error {
	id := c.Param("id")
	ac.repository.Delete(id)

	return c.NoContent(
		http.StatusOK,
	)
}
