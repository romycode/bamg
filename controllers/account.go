package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/romycode/bank-manager/errors"
	"github.com/romycode/bank-manager/models"
)

type AccountController struct {
	repository models.AccountRepository
}

func NewAccountController(repository models.AccountRepository) AccountController {
	return AccountController{repository: repository}
}

func (ac AccountController) GetAllAccounts(c echo.Context) error {
	a, _ := json.Marshal(ac.repository.All())

	return c.JSONBlob(
		http.StatusOK,
		a,
	)
}

func (ac AccountController) CreateAccount(c echo.Context) error {
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

func (ac AccountController) DeleteAccount(c echo.Context) error {
	id := c.Param("id")

	ac.repository.Delete(id)

	return c.NoContent(
		http.StatusOK,
	)
}
