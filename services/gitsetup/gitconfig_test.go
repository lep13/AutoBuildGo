package gitsetup

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadEnvSuccess(t *testing.T) {
	// sets the environment variable for the test
	expectedURL := "https://api.github.com/repos/lep13/ServiceTemplate/generate"
	os.Setenv("TEMPLATE_URL", expectedURL)
	defer os.Unsetenv("TEMPLATE_URL") // Clean up after the test

	// invoke the function that loads the environment
	loadEnv()

	// check if the environment variable is correctly set
	result := os.Getenv("TEMPLATE_URL")
	assert.Equal(t, expectedURL, result, "TEMPLATE_URL did not match expected value")
}


func TestLoadEnvFailNoEnvFile(t *testing.T) {
    // Clear TEMPLATE_URL to simulate a missing .env and unset environment variable
    originalValue := os.Getenv("TEMPLATE_URL")
    os.Unsetenv("TEMPLATE_URL")
    defer os.Setenv("TEMPLATE_URL", originalValue) // Restore after test

    // This should panic because TEMPLATE_URL is not set and no .env will be found
    assert.Panics(t, func() {
        loadEnv()
    }, "Expected panic due to missing TEMPLATE_URL and no .env file")
}


func TestCheckTemplateURLPanics(t *testing.T) {
	// Ensure no TEMPLATE_URL is set
	os.Unsetenv("TEMPLATE_URL")
	defer os.Unsetenv("TEMPLATE_URL") // Clean up after the test

	// This should panic because TEMPLATE_URL is required
	assert.Panics(t, func() {
		checkTemplateURL()
	}, "Expected panic due to missing TEMPLATE_URL")
}

func TestCheckTemplateURLSuccess(t *testing.T) {
	// Set a valid TEMPLATE_URL
	os.Setenv("TEMPLATE_URL", "https://example.com/template")
	defer os.Unsetenv("TEMPLATE_URL") // Clean up after the test

	// This should not panic
	assert.NotPanics(t, func() {
		checkTemplateURL()
	}, "Did not expect panic with valid TEMPLATE_URL set")
}

func TestDefaultRepoConfigSuccess(t *testing.T) {
    // Setup: Set TEMPLATE_URL in environment
    expectedURL := "https://example.com/template"
    os.Setenv("TEMPLATE_URL", expectedURL)
    defer os.Unsetenv("TEMPLATE_URL")

    repoName := "test-repo"
    description := "Test repository"

    // Call the function
    config := DefaultRepoConfig(repoName, description)

    // Assertions
    assert.Equal(t, repoName, config.Name)
    assert.Equal(t, description, config.Description)
    assert.Equal(t, true, config.Private)
    assert.Equal(t, true, config.AutoInit)
    assert.Equal(t, expectedURL, config.TemplateURL)
}

func TestDefaultRepoConfigMissingTemplateURL(t *testing.T) {
    // Ensure TEMPLATE_URL is not set
    os.Unsetenv("TEMPLATE_URL")

    repoName := "test-repo"
    description := "Test repository"

    // Expected to panic because TEMPLATE_URL is not set
    assert.Panics(t, func() {
        DefaultRepoConfig(repoName, description)
    }, "Expected panic due to missing TEMPLATE_URL")
}
