package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/romycode/bank-manager/internal/bank_manager_api/errors"
	"github.com/romycode/bank-manager/internal/bank_manager_api/models"
)

type UserController struct {
	repository models.UserRepository
}

func NewUserController(repository models.UserRepository) UserController {
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
	errors.HandleError(err)

	uc.repository.Save(*u)

	res, _ := json.Marshal(u)
	return c.JSONBlob(
		http.StatusCreated,
		res,
	)
}

func (uc *UserController) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	uc.repository.Delete(id)

	return c.NoContent(
		http.StatusOK,
	)
}
