package gitsetup

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type CommandExecutor interface {
	ExecuteCommand(name string, arg ...string) ([]byte, error)
}

type RealCommandExecutor struct{}

func (e RealCommandExecutor) ExecuteCommand(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.Bytes(), err
}

func FetchSecretToken(executor CommandExecutor) (string, error) {
	output, err := executor.ExecuteCommand("git", "credential", "fill")
	if err != nil {
		return "", fmt.Errorf("error running git credential fill: %v", err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(output))
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
