package gitsetup

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// CommandRunner defines an interface for running commands.
type CommandRunner interface {
	Run(cmd *exec.Cmd) error
	Output(cmd *exec.Cmd) ([]byte, error)
}

// DefaultCommandRunner is the default implementation of CommandRunner.
type DefaultCommandRunner struct{}

// Run executes a command.
func (r *DefaultCommandRunner) Run(cmd *exec.Cmd) error {
	return cmd.Run()
}

// Output gets the output of a command.
func (r *DefaultCommandRunner) Output(cmd *exec.Cmd) ([]byte, error) {
	return cmd.Output()
}

var runner CommandRunner = &DefaultCommandRunner{}

// FetchSecretToken retrieves the GitHub token from the local Git configuration using the git credential system.
func FetchSecretToken() (string, error) {
	cmd := exec.Command("git", "credential", "fill")
	cmdInput := bytes.NewBufferString("url=https://github.com\n")
	cmd.Stdin = cmdInput

	var out bytes.Buffer
	cmd.Stdout = &out
	err := runner.Run(cmd)
	if err != nil {
		return "", fmt.Errorf("error running git credential fill: %v", err)
	}

	scanner := bufio.NewScanner(&out)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "password=") {
			return strings.TrimPrefix(line, "password="), nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading credential output: %v", err)
	}

	return "", fmt.Errorf("no token found in Git credentials")
}
