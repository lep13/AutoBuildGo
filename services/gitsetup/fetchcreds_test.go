package gitsetup

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCommandExecutor struct {
	mock.Mock
}

func (m *MockCommandExecutor) ExecuteCommand(name string, arg ...string) ([]byte, error) {
	args := m.Called(name, arg)
	return args.Get(0).([]byte), args.Error(1)
}

func TestFetchSecretToken(t *testing.T) {
	mockExec := new(MockCommandExecutor)
	// executor := RealCommandExecutor{}

	// Test case: successful token retrieval
	mockExec.On("ExecuteCommand", "git", []string{"credential", "fill"}).Return([]byte("password=mySecretToken\n"), nil)
	token, err := FetchSecretToken(mockExec)
	assert.NoError(t, err)
	assert.Equal(t, "mySecretToken", token)

	// Test case: command execution error
	mockExec.On("ExecuteCommand", "git", []string{"credential", "fill"}).Return(nil, fmt.Errorf("failed to run command"))
	_, err = FetchSecretToken(mockExec)
	assert.Error(t, err)

	// Test case: no token found
	mockExec.On("ExecuteCommand", "git", []string{"credential", "fill"}).Return([]byte("username=user\n"), nil)
	_, err = FetchSecretToken(mockExec)
	assert.Error(t, err)
}
