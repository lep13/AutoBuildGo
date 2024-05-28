package gitsetup

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// HTTPClient is an interface that defines the Do method used by http.Client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// GitClient is a structure that holds dependencies for making HTTP requests.
type GitClient struct {
	HTTPClient      HTTPClient
	FetchSecretFunc func() (string, error)
}

// NewGitClient returns an instance of GitClient with default dependencies.
func NewGitClient() *GitClient {
	return &GitClient{
		HTTPClient:      &http.Client{},
		FetchSecretFunc: FetchSecretToken,
	}
}

// CreateGitRepository creates a new GitHub repository using the specified configuration.
func (client *GitClient) CreateGitRepository(config RepoConfig) error {
	// Fetch the token using the FetchSecretToken function.
	token, err := client.FetchSecretFunc()
	if err != nil {
		return err
	}
	return client.createRepositoryWithTemplate(config, token)
}

// createRepositoryWithTemplate sends a request to GitHub API to create a repository from a template.
func (client *GitClient) createRepositoryWithTemplate(config RepoConfig, token string) error {
	data, err := json.Marshal(map[string]interface{}{
		"name":        config.Name,
		"description": config.Description,
		"private":     config.Private,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, config.TemplateURL, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	return fmt.Errorf("failed to create repository, status code: %d, response: %s", resp.StatusCode, string(body))
}
