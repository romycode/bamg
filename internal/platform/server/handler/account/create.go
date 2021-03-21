package account

import (
	"log"
	"net/http"

	noter "github.com/romycode/bank-manager/internal"

	"github.com/gin-gonic/gin"
)

func CreateAccountHandler(repository noter.AccountRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		a := new(noter.Account)
		err := ctx.BindJSON(a)
		if err != nil {
			log.Fatal(err)
		}

		repository.Save(ctx, *a)

		ctx.JSON(
			http.StatusCreated,
			a,
		)
	}
}
