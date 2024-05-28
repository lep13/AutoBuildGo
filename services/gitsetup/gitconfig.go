package gitsetup

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	loadEnv()
}

func loadEnv() {
	// Attempt to load .env file first
	err := godotenv.Load()
	if err != nil {
		log.Println("Debug: .env file not loaded")
	}
}

func checkTemplateURL() {
	if os.Getenv("TEMPLATE_URL") == "" {
		panic("TEMPLATE_URL must be set in the environment")
	}
}

// DefaultRepoConfig constructs a default repository configuration.
func DefaultRepoConfig(repoName string, description string) RepoConfig {
	templateURL := os.Getenv("TEMPLATE_URL")
	if templateURL == "" {
		panic("TEMPLATE_URL must be set in the environment")
	}

	return RepoConfig{
		Name:        repoName,
		Description: description,
		Private:     true,
		AutoInit:    true,
		TemplateURL: templateURL,
	}
}
