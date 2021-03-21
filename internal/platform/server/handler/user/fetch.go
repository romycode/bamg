package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	noter "github.com/romycode/bank-manager/internal"
)

func FetchAllUsersHandler(repository noter.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusOK,
			repository.All(ctx),
		)
	}
}