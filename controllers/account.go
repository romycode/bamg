package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/romycode/bank-manager/database/repositories"
	"github.com/romycode/bank-manager/models"
)

func GetAllAccounts() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			repositories.GetAllAccounts(),
		)
	}
}

func CreateAccount() echo.HandlerFunc {
	return func(c echo.Context) error {
		a := new(models.Account)
		err := c.Bind(a)
		if err != nil {
			fmt.Println(err)
		}
		a.IBAN = models.NewIban()

		repositories.SaveAccount(a)

		return c.JSON(
			http.StatusCreated,
			a,
		)
	}
}

func DeleteAccount() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		repositories.DeleteAccount(id)

		return c.JSON(
			http.StatusOK,
			nil,
		)
	}
}
