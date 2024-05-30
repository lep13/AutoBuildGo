package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/lep13/AutoBuildGo/services/ecr"
	"github.com/lep13/AutoBuildGo/services/gitsetup"
)

func main() {
	if len(os.Args) > 1 {
		handleCLI()
	} else {
		gitsetup.HandleWebServer()
	}
}

func handleCLI() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <repo-name> [\"optional description\"]")
	}
	repoName := os.Args[1]
	description := "Created from a template via automated setup" // Default description if none provided

	if len(os.Args) > 2 {
		description = strings.Join(os.Args[2:], " ") // Combine all arguments after repoName as description
	}

	// Create ECR client
	ecrClient, err := ecr.CreateECRClient()
	if err != nil {
		log.Fatalf("Failed to create ECR client: %v", err)
	}

	// Create ECR Repository
	if err := ecr.CreateRepo(repoName, ecrClient); err != nil {
		log.Fatalf("Failed to create ECR repository: %v", err)
	}

	// Ensure environment is loaded
	gitsetup.LoadEnv()

	// Create Git Repository
	config := gitsetup.DefaultRepoConfig(repoName, description)
	gitClient := gitsetup.NewGitClient() // Create an instance of GitClient

	if err := gitClient.CreateGitRepository(config); err != nil {
		log.Fatalf("Failed to create Git repository: %v", err)
	}

	log.Println("ECR and Git repositories created successfully")

	// 20 second time delay
	time.Sleep(20 * time.Second)

	// Clone the repo, update go.mod, and push changes
	if err := gitsetup.CloneAndPushRepo(repoName); err != nil {
		log.Fatalf("Failed to clone and push repository: %v", err)
	}
}
