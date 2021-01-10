package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/romycode/bank-manager/database/repositories"
	"github.com/romycode/bank-manager/errors"
	"github.com/romycode/bank-manager/models"
)

type UserController struct {
	repository repositories.UserRepository
}

func NewUserController(repository repositories.UserRepository) UserController {
	return UserController{repository: repository}
}

func (uc *UserController) GetAllUsers(c echo.Context) error {
	responseBody, err := json.Marshal(uc.repository.All())
	errors.HandleError(err)

	return c.JSONBlob(
		http.StatusOK,
		responseBody,
	)
}

func (uc *UserController) CreateUser(c echo.Context) error {
	u := new(models.User)
	err := c.Bind(u)
	if err != nil {
		fmt.Println(err)
	}

	uc.repository.Save(u)

	return c.JSON(
		http.StatusCreated,
		u,
	)
}

func (uc *UserController) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	uc.repository.Delete(id)

	return c.JSON(
		http.StatusOK,
		nil,
	)
}
