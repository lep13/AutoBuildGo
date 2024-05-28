package gitsetup

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Set up the environment by explicitly loading the .env file

	absPath := `C:\Users\Relanto\AutoBuildGo\.env` 
    err := godotenv.Load(absPath)
    if err != nil {
        log.Fatalf("Failed to load .env file for tests: %v", err)
    }


	// err := godotenv.Load("../../.env")
	// if err != nil {
	// 	log.Fatalf("Failed to load .env file for tests: %v", err)
	// }

	// cwd, err := os.Getwd()
	// if err != nil {
	// 	log.Fatalf("Failed to get current working directory: %v", err)
	// }
	// log.Println("Current Working Directory during tests:", cwd)



	// Run tests
	code := m.Run()

	// Clean up / tear down environment if necessary
	os.Exit(code)
}

// tests the successful creation of a repo config.
func TestDefaultRepoConfigSuccess(t *testing.T) {
	repoName := "testRepo"
	description := "A test repository"
	config := DefaultRepoConfig(repoName, description)

	assert.Equal(t, repoName, config.Name)
	assert.Equal(t, description, config.Description)
	assert.True(t, config.Private)
	assert.True(t, config.AutoInit)
	assert.Equal(t, "http://example.com/template", config.TemplateURL)
}

// tests the panic when TEMPLATE_URL is not set.
func TestDefaultRepoConfigNoTemplateURL(t *testing.T) {
	// Backup current environment variable
	originalURL := os.Getenv("TEMPLATE_URL")
	defer os.Setenv("TEMPLATE_URL", originalURL) // Ensure environment is restored after the test

	// Clear TEMPLATE_URL for the error condition
	os.Unsetenv("TEMPLATE_URL")

	assert.Panics(t, func() {
		DefaultRepoConfig("testRepo", "A test repository")
	}, "The code did not panic when TEMPLATE_URL was not set")
}

// tests that environment variables are properly loaded and used.
func TestLoadEnvironmentVariables(t *testing.T) {
	templateURL := os.Getenv("TEMPLATE_URL")
	assert.NotEmpty(t, templateURL, "TEMPLATE_URL should not be empty after loading environment variables")
}
