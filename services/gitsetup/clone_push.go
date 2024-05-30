package gitsetup

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

// CloneAndPushRepo clones the repository, updates the go.mod file, and pushes the changes back to GitHub.
func CloneAndPushRepo(repoName string) error {
	// Fetch GitHub token
	token, err := FetchSecretToken()
	if err != nil {
		return fmt.Errorf("error fetching GitHub token: %v", err)
	}

	// Fetch GitHub username
	username, err := FetchGitHubUsername(token)
	if err != nil {
		return fmt.Errorf("error fetching GitHub username: %v", err)
	}

	// Clone the repository
	repoURL := fmt.Sprintf("https://%s@github.com/%s/%s.git", token, username, repoName)
	cmd := exec.Command("git", "clone", repoURL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error cloning repository: %v", err)
	}

	// Change directory to the cloned repository
	if err := os.Chdir(repoName); err != nil {
		return fmt.Errorf("error changing directory to cloned repository: %v", err)
	}

	// Update go.mod file
	goModFile := "go.mod"
	input, err := os.ReadFile(goModFile)
	if err != nil {
		return fmt.Errorf("error reading go.mod file: %v", err)
	}

	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, "module") {
			lines[i] = fmt.Sprintf("module github.com/%s/%s", username, repoName)
			break
		}
	}
	output := strings.Join(lines, "\n")
	if err := os.WriteFile(goModFile, []byte(output), 0644); err != nil {
		return fmt.Errorf("error writing to go.mod file: %v", err)
	}

	// Commit and push changes
	cmd = exec.Command("git", "add", goModFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error adding go.mod file to git: %v", err)
	}

	cmd = exec.Command("git", "commit", "-m", "Update go.mod module path")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error committing changes: %v", err)
	}

	cmd = exec.Command("git", "push")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error pushing changes: %v", err)
	}

	// Go back to the previous directory
	if err := os.Chdir(".."); err != nil {
		return fmt.Errorf("error changing back to the parent directory: %v", err)
	}

	// Remove the cloned repository
	if err := os.RemoveAll(repoName); err != nil {
		return fmt.Errorf("error removing the cloned repository: %v", err)
	}

	return nil
}

// FetchGitHubUsername fetches the GitHub username of the authenticated user.
func FetchGitHubUsername(token string) (string, error) {
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "token "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch GitHub username, status code: %d", resp.StatusCode)
	}

	var result struct {
		Login string `json:"login"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Login, nil
}
