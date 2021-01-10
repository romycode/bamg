package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/romycode/bank-manager/database/repositories"
	"github.com/romycode/bank-manager/errors"
	"github.com/romycode/bank-manager/models"
)

type AccountController struct {
	repository repositories.AccountRepository
}

func NewAccountController(repository repositories.AccountRepository) AccountController {
	return AccountController{repository: repository}
}

func (ac AccountController) GetAllAccounts(c echo.Context) error {
	return c.JSON(
		http.StatusOK,
		ac.repository.All(),
	)
}

func (ac AccountController) CreateAccount(c echo.Context) error {
	a := new(models.Account)
	err := c.Bind(a)
	errors.HandleError(err)

	a.IBAN = models.NewIban()

	ac.repository.Save(a)

	return c.JSON(
		http.StatusCreated,
		a,
	)
}

func (ac AccountController) DeleteAccount(c echo.Context) error {
	id := c.Param("id")

	ac.repository.Delete(id)

	return c.JSON(
		http.StatusOK,
		nil,
	)
}
