package main

import (
	"flag"
	"log"

	"github.com/lep13/AutoBuildGo/services/ecr"
	// "github.com/lep13/AutoBuildGo/services/gitsetup"
)

func main() {
	repoName := flag.String("repo", "", "The name of the repository to create")
	flag.Parse()

	if *repoName == "" {
		log.Fatal("Repository name is required. Usage: go run main.go -repo=<repo-name>")
	}

	// Create ECR repository
	err := ecr.CreateRepo(*repoName)
	if err != nil {
		log.Fatalf("Failed to create ECR repository: %v", err)
	}

	// Create Git repository
	// config := gitsetup.RepoConfig{
	// 	Name:        *repoName,
	// 	Description: "Created from a template via automated setup",
	// 	Private:     true,
	// 	AutoInit:    true,
	// }
	// gitsetup.CreateRepository(config)

	log.Println("ECR and Git repositories created successfully")
}
