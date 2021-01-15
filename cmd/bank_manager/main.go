package main

import (
	"github.com/romycode/bank-manager/internal/bank_manager_api/server"
)

func main() {
	server.Serve("8080")
}
