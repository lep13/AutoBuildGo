package main

import (
	"fmt"
	"log"

	"github.com/lep13/AutoBuildGo/cmd/api"
)

const (
	host = "localhost"
	port = "8080"
)

func main() {

	err := api.ServeHTTP(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("[ERROR]: %s", err)
	}
}
