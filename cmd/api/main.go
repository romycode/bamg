package main

import (
	"log"

	"github.com/romycode/bank-manager/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
