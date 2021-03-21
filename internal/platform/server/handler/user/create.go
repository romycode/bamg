package user

import (
	"net/http"

	noter "github.com/romycode/bank-manager/internal"

	"github.com/gin-gonic/gin"

	"log"
)

func CreateUserHandler(repository noter.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u := new(noter.User)
		err := ctx.Bind(u)
		if err != nil {
		log.Fatal(err)
	}

		repository.Save(ctx, *u)

		ctx.JSON(
			http.StatusCreated,
			u,
		)
	}
}