package gitsetup

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockExecutor struct {
	mock.Mock
}

func (m *MockExecutor) ExecuteCommand(name string, arg ...string) ([]byte, error) {
	args := m.Called(name, arg)
	return args.Get(0).([]byte), args.Error(1)
}

type MockHttpClient struct {
	mock.Mock
}

func (c *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	args := c.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}


func TestCreateGitRepository(t *testing.T) {
	executor := new(MockExecutor)
	client := new(MockHttpClient)
	executor.On("ExecuteCommand", "git", mock.Anything).Return([]byte("password=mySecretToken\n"), nil)

	// Mock server response
	resp := &http.Response{
		StatusCode: http.StatusCreated,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{}`))),
	}
	client.On("Do", mock.Anything).Return(resp, nil)

	config := RepoConfig{
		Name:        "testrepo",
		Description: "A test repository",
		Private:     true,
		TemplateURL: "http://api.github.com/repos/template/generate",
	}

	err := CreateGitRepository(client, config, executor)
	assert.NoError(t, err)

	// Failure scenarios
	executor.On("ExecuteCommand", "git", mock.Anything).Return(nil, errors.New("command failed"))
	err = CreateGitRepository(client, config, executor)
	assert.Error(t, err)

	client.On("Do", mock.Anything).Return(nil, errors.New("HTTP request failed"))
	err = CreateGitRepository(client, config, executor)
	assert.Error(t, err)
}
