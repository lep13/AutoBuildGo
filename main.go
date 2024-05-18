package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lep13/AutoBuildGo/cmd/api/routes"
)

const (
	host = "localhost"
	port = "8080"
)

func main() {
	http.HandleFunc("/", routes.HandleRequests)

	address := fmt.Sprintf("%s:%s", host, port)
	log.Printf("Starting server on %s\n", address)

	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
