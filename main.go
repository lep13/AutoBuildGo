package main

import (
	"log"

	"github.com/lep13/AutoBuildGo/ecr"
)

func main() {
    err := ecr.CreateRepo()
    if err != nil {
        log.Fatalf("Failed to create repository: %v", err)
    }

    log.Println("Repository created successfully")
}
