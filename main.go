package main

import (
	"flag"
	"log"

	"github.com/lep13/AutoBuildGo/ecr"
)

func main() {
	// Define a flag for repository name
	repoName := flag.String("repo", "", "Repository name")
	flag.Parse()

	// Check if repository name is provided
	if *repoName == "" {
		log.Fatal("Repository name is required. Usage: go run main.go -repo=<reponame>")
	}

	err := ecr.CreateRepo(*repoName)
	if err != nil {
		log.Fatalf("Failed to create repository: %v", err)
	}

	log.Println("Repository created successfully")
}
