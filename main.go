package main

import (
	// "context"
	"log"
	"os"

	// "strings"

	// "github.com/aws/aws-sdk-go-v2/config"
	"github.com/lep13/AutoBuildGo/services/ecr"
	// "github.com/lep13/AutoBuildGo/services/gitsetup"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <repo-name> [\"optional description\"]")
	}
	repoName := os.Args[1]
	description := "Created from a template via automated setup" // Default description if none provided

	// if len(os.Args) > 2 {
	// 	description = strings.Join(os.Args[2:], " ") // Combine all arguments after repoName as description
	// }

	// Create AWS client

	var client ecr.AWSClient
	// Create ECR Repository
	if err := ecr.CreateRepo(repoName, client); err != nil {
		log.Fatalf("Failed to create ECR repository: %v", err)
	}

	// Create Git Repository
	config := gitsetup.DefaultRepoConfig(repoName, description)
	gitClient := gitsetup.NewGitClient() // Create an instance of GitClient

	if err := gitClient.CreateGitRepository(config); err != nil {
		log.Fatalf("Failed to create Git repository: %v", err)
	}

	log.Println("ECR and Git repositories created successfully")
}
