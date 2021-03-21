package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	noter "github.com/romycode/bank-manager/internal"
)

func DeleteUserHandler(repository noter.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		repository.Delete(ctx, id)

		ctx.Status(
			http.StatusOK,
		)
	}
}