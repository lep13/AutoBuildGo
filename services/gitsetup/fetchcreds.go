package gitsetup

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// retrieves the GitHub token from the local Git configuration.
func FetchSecretToken() (string, error) {
	// git credential fill command to fetch stored credentials
	cmd := exec.Command("git", "credential", "fill")
	cmdInput := bytes.NewBufferString("url=https://github.com\n")
	cmd.Stdin = cmdInput

	// output from the command
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error running git credential fill: %v", err)
	}

	// Read the output and parse the password field
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
