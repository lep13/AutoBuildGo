package gitsetup

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func init() {
	// Find the root directory
	root, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting working directory: %s", err)
	}
	// Attempt to load the .env file
	err = godotenv.Load(filepath.Join(root, ".env"))
	if err != nil {
		log.Fatalf("No .env file found: %s", err)
	}
}

// default repository configuration with a dynamic description.
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
