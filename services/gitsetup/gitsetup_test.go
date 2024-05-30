package gitsetup

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockHTTPClient is a mock of the HTTPClient interface
type MockHTTPClient struct {
	mock.Mock
}

// Do is a mock implementation of the Do method
func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	resp, _ := args.Get(0).(*http.Response)
	return resp, args.Error(1)
}

// MockFetchSecretFunc is a mock implementation of FetchSecretFunc
func MockFetchSecretFunc() (string, error) {
	return "mocked_token", nil
}

func TestCreateGitRepository(t *testing.T) {
	t.Run("CreateGitRepository_Success", func(t *testing.T) {
		mockHTTPClient := new(MockHTTPClient)
		mockResponse := &http.Response{
			StatusCode: http.StatusCreated,
			Body:       io.NopCloser(bytes.NewBufferString(`{}`)),
		}
		mockHTTPClient.On("Do", mock.Anything).Return(mockResponse, nil)

		client := &GitClient{
			HTTPClient:      mockHTTPClient,
			FetchSecretFunc: MockFetchSecretFunc,
		}

		config := RepoConfig{
			Name:        "test-repo",
			Description: "A test repository",
			Private:     true,
			TemplateURL: "https://api.github.com/repos/user/template-repo/generate",
		}

		err := client.CreateGitRepository(config)
		assert.NoError(t, err)
		mockHTTPClient.AssertExpectations(t)
	})

	t.Run("CreateGitRepository_FetchTokenFailure", func(t *testing.T) {
		mockHTTPClient := new(MockHTTPClient)
		client := &GitClient{
			HTTPClient: mockHTTPClient,
			FetchSecretFunc: func() (string, error) {
				return "", errors.New("failed to fetch token")
			},
		}

		config := RepoConfig{
			Name:        "test-repo",
			Description: "A test repository",
			Private:     true,
			TemplateURL: "https://api.github.com/repos/user/template-repo/generate",
		}

		err := client.CreateGitRepository(config)
		assert.EqualError(t, err, "failed to fetch token")
	})

	t.Run("CreateGitRepository_RequestFailure", func(t *testing.T) {
		mockHTTPClient := new(MockHTTPClient)
		mockHTTPClient.On("Do", mock.Anything).Return(new(http.Response), errors.New("http request failed"))

		client := &GitClient{
			HTTPClient:      mockHTTPClient,
			FetchSecretFunc: MockFetchSecretFunc,
		}

		config := RepoConfig{
			Name:        "test-repo",
			Description: "A test repository",
			Private:     true,
			TemplateURL: "https://api.github.com/repos/user/template-repo/generate",
		}

		err := client.CreateGitRepository(config)
		assert.EqualError(t, err, "http request failed")
		mockHTTPClient.AssertExpectations(t)
	})

	t.Run("CreateGitRepository_Non201StatusCode", func(t *testing.T) {
		mockHTTPClient := new(MockHTTPClient)
		mockResponse := &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:       io.NopCloser(bytes.NewBufferString(`{"message":"Bad Request"}`)),
		}
		mockHTTPClient.On("Do", mock.Anything).Return(mockResponse, nil)

		client := &GitClient{
			HTTPClient:      mockHTTPClient,
			FetchSecretFunc: MockFetchSecretFunc,
		}

		config := RepoConfig{
			Name:        "test-repo",
			Description: "A test repository",
			Private:     true,
			TemplateURL: "https://api.github.com/repos/user/template-repo/generate",
		}

		err := client.CreateGitRepository(config)
		assert.EqualError(t, err, "failed to create repository, status code: 400, response: {\"message\":\"Bad Request\"}")
		mockHTTPClient.AssertExpectations(t)
	})
}

func TestCreateRepositoryWithTemplate(t *testing.T) {
	t.Run("CreateRepositoryWithTemplate_Success", func(t *testing.T) {
		mockHTTPClient := new(MockHTTPClient)
		mockResponse := &http.Response{
			StatusCode: http.StatusCreated,
			Body:       io.NopCloser(bytes.NewBufferString(`{}`)),
		}
		mockHTTPClient.On("Do", mock.Anything).Return(mockResponse, nil)

		client := &GitClient{
			HTTPClient: mockHTTPClient,
		}

		config := RepoConfig{
			Name:        "test-repo",
			Description: "A test repository",
			Private:     true,
			TemplateURL: "https://api.github.com/repos/user/template-repo/generate",
		}
		token := "mocked_token"

		err := client.createRepositoryWithTemplate(config, token)
		assert.NoError(t, err)
		mockHTTPClient.AssertExpectations(t)
	})

	t.Run("CreateRepositoryWithTemplate_RequestFailure", func(t *testing.T) {
		mockHTTPClient := new(MockHTTPClient)
		mockHTTPClient.On("Do", mock.Anything).Return(new(http.Response), errors.New("http request failed"))

		client := &GitClient{
			HTTPClient: mockHTTPClient,
		}

		config := RepoConfig{
			Name:        "test-repo",
			Description: "A test repository",
			Private:     true,
			TemplateURL: "https://api.github.com/repos/user/template-repo/generate",
		}
		token := "mocked_token"

		err := client.createRepositoryWithTemplate(config, token)
		assert.EqualError(t, err, "http request failed")
		mockHTTPClient.AssertExpectations(t)
	})

	t.Run("CreateRepositoryWithTemplate_Non201StatusCode", func(t *testing.T) {
		mockHTTPClient := new(MockHTTPClient)
		mockResponse := &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:       io.NopCloser(bytes.NewBufferString(`{"message":"Bad Request"}`)),
		}
		mockHTTPClient.On("Do", mock.Anything).Return(mockResponse, nil)

		client := &GitClient{
			HTTPClient: mockHTTPClient,
		}

		config := RepoConfig{
			Name:        "test-repo",
			Description: "A test repository",
			Private:     true,
			TemplateURL: "https://api.github.com/repos/user/template-repo/generate",
		}
		token := "mocked_token"

		err := client.createRepositoryWithTemplate(config, token)
		assert.EqualError(t, err, "failed to create repository, status code: 400, response: {\"message\":\"Bad Request\"}")
		mockHTTPClient.AssertExpectations(t)
	})
}
