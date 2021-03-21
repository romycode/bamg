package account

import (
	"net/http"

	"github.com/gin-gonic/gin"

	noter "github.com/romycode/bank-manager/internal"
)

func FetchAllAccountsHandler(repository noter.AccountRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusOK,
			repository.All(ctx),
		)
	}
}