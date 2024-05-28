package gitsetup

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockHTTPClient is a mock implementation of an HTTP client.
type MockHTTPClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func TestCreateGitRepositorySuccess(t *testing.T) {
	// mocks a successful http client response
	mockClient := &MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			resp := &http.Response{
				StatusCode: http.StatusCreated,
				Body:       io.NopCloser(bytes.NewBufferString("")),
			}
			return resp, nil
		},
	}

	// Mock the FetchSecretToken function
	mockFetchSecretFunc := func() (string, error) {
		return "mock-token", nil
	}

	client := &GitClient{
		HTTPClient:      mockClient,
		FetchSecretFunc: mockFetchSecretFunc,
	}

	config := RepoConfig{
		Name:        "test-repo",
		Description: "Test repository",
		Private:     true,
		TemplateURL: "https://api.github.com/repos/lep13/ServiceTemplate/generate",
	}

	err := client.CreateGitRepository(config)
	assert.NoError(t, err)
}

func TestCreateGitRepositoryFetchTokenError(t *testing.T) {
	// Mock the FetchSecretToken function to return an error
	mockFetchSecretFunc := func() (string, error) {
		return "", errors.New("failed to fetch token")
	}

	client := &GitClient{
		HTTPClient:      &http.Client{},
		FetchSecretFunc: mockFetchSecretFunc,
	}

	config := RepoConfig{
		Name:        "test-repo",
		Description: "Test repository",
		Private:     true,
		TemplateURL: "https://api.github.com/repos/lep13/ServiceTemplate/generate",
	}

	err := client.CreateGitRepository(config)
	assert.Error(t, err)
	assert.Equal(t, "failed to fetch token", err.Error())
}

func TestCreateRepositoryWithTemplateSuccess(t *testing.T) {
	// Mock the HTTP client
	mockClient := &MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			resp := &http.Response{
				StatusCode: http.StatusCreated,
				Body:       io.NopCloser(bytes.NewBufferString("")),
			}
			return resp, nil
		},
	}

	client := &GitClient{
		HTTPClient:      mockClient,
		FetchSecretFunc: FetchSecretToken,
	}

	config := RepoConfig{
		Name:        "test-repo",
		Description: "Test repository",
		Private:     true,
		TemplateURL: "https://api.github.com/repos/lep13/ServiceTemplate/generate",
	}
	token := "mock-token"

	err := client.createRepositoryWithTemplate(config, token)
	assert.NoError(t, err)
}

func TestCreateRepositoryWithTemplateError(t *testing.T) {
	// Mock the HTTP client to return an error response
	mockClient := &MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			resp := &http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       io.NopCloser(bytes.NewBufferString("Internal Server Error")),
			}
			return resp, nil
		},
	}

	client := &GitClient{
		HTTPClient:      mockClient,
		FetchSecretFunc: FetchSecretToken,
	}

	config := RepoConfig{
		Name:        "test-repo",
		Description: "Test repository",
		Private:     true,
		TemplateURL: "https://api.github.com/repos/lep13/ServiceTemplate/generate",
	}
	token := "mock-token"

	err := client.createRepositoryWithTemplate(config, token)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to create repository")
}

func TestCreateRepositoryWithTemplateRequestError(t *testing.T) {
	// Mock the HTTP client to simulate a request error
	mockClient := &MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("request error")
		},
	}

	client := &GitClient{
		HTTPClient:      mockClient,
		FetchSecretFunc: FetchSecretToken,
	}

	config := RepoConfig{
		Name:        "test-repo",
		Description: "Test repository",
		Private:     true,
		TemplateURL: "https://api.github.com/repos/lep13/ServiceTemplate/generate",
	}
	token := "mock-token"

	err := client.createRepositoryWithTemplate(config, token)
	assert.Error(t, err)
	assert.Equal(t, "request error", err.Error())
}
