package main

import (
	"log"
	"os"

	"github.com/lep13/AutoBuildGo/services/ecr"
	"github.com/lep13/AutoBuildGo/services/gitsetup"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: go run main.go <repo-name>")
	}
	repoName := os.Args[1]

	// Create ECR Repository
	if err := ecr.CreateRepo(repoName); err != nil {
		log.Fatalf("Failed to create ECR repository: %v", err)
	}

	// Create Git Repository
	config := gitsetup.DefaultRepoConfig(repoName)

	if err := gitsetup.CreateGitRepository(config); err != nil {
		log.Fatalf("Failed to create Git repository: %v", err)
	}

	log.Println("ECR and Git repositories created successfully")
}
