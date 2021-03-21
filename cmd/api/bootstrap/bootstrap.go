package bootstrap

import "github.com/romycode/bank-manager/internal/platform/server"

const (
	host = "0.0.0.0"
	port = 8080
)

func Run() error {
	srv := server.New(host, port)
	return srv.Run()
}
