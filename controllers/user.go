package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/romycode/bank-manager/database/repositories"
	"github.com/romycode/bank-manager/models"
)

func GetAllUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			repositories.GetAllUsers(),
		)
	}
}

func CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(models.User)
		err := c.Bind(u)
		if err != nil {
			fmt.Println(err)
		}

		repositories.SaveUser(u)

		return c.JSON(
			http.StatusCreated,
			u,
		)
	}
}

func DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		repositories.DeleteUser(id)

		return c.JSON(
			http.StatusOK,
			nil,
		)
	}
}
