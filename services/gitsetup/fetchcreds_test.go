package gitsetup

import (
	"fmt"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockCommandRunner is a mock implementation of the CommandRunner interface.
type MockCommandRunner struct {
	RunFunc    func(cmd *exec.Cmd) error
	OutputFunc func(cmd *exec.Cmd) ([]byte, error)
}

func (m *MockCommandRunner) Run(cmd *exec.Cmd) error {
	return m.RunFunc(cmd)
}

func (m *MockCommandRunner) Output(cmd *exec.Cmd) ([]byte, error) {
	return m.OutputFunc(cmd)
}

func TestFetchSecretTokenSuccess(t *testing.T) {
	// Mock the command runner
	mockRunner := &MockCommandRunner{
		RunFunc: func(cmd *exec.Cmd) error {
			cmd.Stdout.Write([]byte("password=mock-token\n"))
			return nil
		},
	}

	// Override the global runner with the mock runner
	originalRunner := runner
	runner = mockRunner
	defer func() { runner = originalRunner }()

	token, err := FetchSecretToken()
	assert.NoError(t, err)
	assert.Equal(t, "mock-token", token)
}

func TestFetchSecretTokenRunError(t *testing.T) {
	// Mock the command runner to return an error
	mockRunner := &MockCommandRunner{
		RunFunc: func(cmd *exec.Cmd) error {
			return fmt.Errorf("command error")
		},
	}

	// Override the global runner with the mock runner
	originalRunner := runner
	runner = mockRunner
	defer func() { runner = originalRunner }()

	token, err := FetchSecretToken()
	assert.Error(t, err)
	assert.Equal(t, "", token)
	assert.Contains(t, err.Error(), "error running git credential fill")
}

func TestFetchSecretTokenNoPassword(t *testing.T) {
	// Mock the command runner with no password in the output
	mockRunner := &MockCommandRunner{
		RunFunc: func(cmd *exec.Cmd) error {
			cmd.Stdout.Write([]byte("no password here\n"))
			return nil
		},
	}

	// Override the global runner with the mock runner
	originalRunner := runner
	runner = mockRunner
	defer func() { runner = originalRunner }()

	token, err := FetchSecretToken()
	assert.Error(t, err)
	assert.Equal(t, "", token)
	assert.Contains(t, err.Error(), "no token found in Git credentials")
}

func TestFetchSecretTokenScannerError(t *testing.T) {
	// Mock the command runner with an output that causes a scanner error
	mockRunner := &MockCommandRunner{
		RunFunc: func(cmd *exec.Cmd) error {
			cmd.Stdout.Write([]byte("\x00\x00\x00\x00"))
			return nil
		},
	}

	// Override the global runner with the mock runner
	originalRunner := runner
	runner = mockRunner
	defer func() { runner = originalRunner }()

	token, err := FetchSecretToken()
	assert.Error(t, err)
	assert.Equal(t, "", token)
	assert.Contains(t, err.Error(), "error reading credential output")
}
