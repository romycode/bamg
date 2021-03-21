package server

import (
	"fmt"
	"log"

	noter "github.com/romycode/bank-manager/internal"
	"github.com/romycode/bank-manager/internal/bank_manager_api/database"
	"github.com/romycode/bank-manager/internal/platform/server/handler/account"
	"github.com/romycode/bank-manager/internal/platform/server/handler/health"
	"github.com/romycode/bank-manager/internal/platform/server/handler/user"
	"github.com/romycode/bank-manager/internal/platform/storage/sqlite"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	// deps
	userRepository    noter.UserRepository
	accountRepository noter.AccountRepository
}

func New(host string, port uint) Server {
	db := database.GetConnection()

	accountRepository := sqlite.NewAccountRepository(db)
	userRepository := sqlite.NewUserRepository(db, accountRepository)

	srv := Server{
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		engine:   gin.New(),
		// deps
		accountRepository: accountRepository,
		userRepository:    userRepository,
	}
	srv.registerRoutes()

	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {

	v1 := s.engine.Group("v1")
	v1.GET("/health", health.HealthHandler())

	accounts := v1.Group("/accounts")
	accounts.GET("", account.FetchAllAccountsHandler(s.accountRepository))
	accounts.POST("", account.CreateAccountHandler(s.accountRepository))
	accounts.DELETE("/:id", account.DeleteAccountHandler(s.accountRepository))

	users := v1.Group("/users")
	users.GET("", user.FetchAllUsersHandler(s.userRepository))
	users.POST("", user.CreateUserHandler(s.userRepository))
	users.DELETE("/:id", user.DeleteUserHandler(s.userRepository))
}
