package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/romycode/bank-manager/internal/bank_manager_api/errors"
	"github.com/romycode/bank-manager/internal/bank_manager_api/models"
)

func GetAllUsers(repository models.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusOK,
			repository.All(),
		)
	}
}

func CreateUser(repository models.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u := new(models.User)
		err := ctx.Bind(u)
		errors.HandleError(err)

		repository.Save(*u)

		ctx.JSON(
			http.StatusCreated,
			u,
		)
	}
}

func DeleteUser(repository models.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		repository.Delete(id)

		ctx.Status(
			http.StatusOK,
		)
	}
}
