package gitsetup

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// HttpClient interface for making HTTP requests
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// RealHttpClient implements HttpClient using http.DefaultClient
type RealHttpClient struct{}

func (c *RealHttpClient) Do(req *http.Request) (*http.Response, error) {
	return http.DefaultClient.Do(req)
}

// NewGitHub repository using the specified configuration.
func CreateGitRepository(client HttpClient, config RepoConfig, executor CommandExecutor) error {
	token, err := FetchSecretToken(executor)
	if err != nil {
		return err
	}
	return createRepositoryWithTemplate(client, config, token)
}

// Sends a request to GitHub API to create a repository from a template.
func createRepositoryWithTemplate(client HttpClient, config RepoConfig, token string) error {
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

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create repository, status code: %d, response: %s", resp.StatusCode, string(body))
	}

	return nil
}
