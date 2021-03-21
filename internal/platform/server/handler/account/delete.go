package account

import (
	"net/http"

	noter "github.com/romycode/bank-manager/internal"

	"github.com/gin-gonic/gin"
)

func DeleteAccountHandler(repository noter.AccountRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		repository.Delete(ctx, id)

		ctx.Status(
			http.StatusOK,
		)
	}
}
