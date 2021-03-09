package account

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/romycode/bank-manager/internal/bank_manager_api/errors"
	"github.com/romycode/bank-manager/internal/bank_manager_api/models"
)

type AccountHandler struct {
	repository models.AccountRepository
}

func NewAccountController(repository models.AccountRepository) AccountHandler {
	return AccountHandler{repository: repository}
}

func GetAllAccounts(repository models.AccountRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusOK,
			repository.All(),
		)
	}
}

func CreateAccount(repository models.AccountRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		a := new(models.Account)
		err := ctx.Bind(a)
		errors.HandleError(err)
		repository.Save(*a)

		ctx.JSON(
			http.StatusCreated,
			a,
		)
	}
}

func DeleteAccount(repository models.AccountRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		repository.Delete(id)
		ctx.Status(
			http.StatusCreated,
		)
	}
}
