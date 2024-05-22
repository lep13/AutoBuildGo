// create_git_repo.go

package Repo

import (
    "bytes"
    "encoding/json"
    "errors"
    "net/http"
)

// GitHubRepo represents the JSON structure for creating a repository on GitHub
type GitHubRepo struct {
    Name        string `json:"name"`
    Description string `json:"description"`
    Private     bool   `json:"private"`
}

// CreateGitHubRepo creates a new GitHub repository
func CreateGitHubRepo(token, gitrepoName, description string, isPrivate bool) error {
    // Create JSON payload for repository creation
    repo := GitHubRepo{
        Name:        gitrepoName,
        Description: description,
        Private:     isPrivate,
    }
    payload, err := json.Marshal(repo)
    if err != nil {
        return err
    }

    // Create HTTP client
    client := &http.Client{}

    // Prepare request
    req, err := http.NewRequest("POST", "https://api.github.com/user/repos", bytes.NewBuffer(payload))
    if err != nil {
        return err
    }

    // Set authorization header
    req.Header.Set("Authorization", "token "+token)
    req.Header.Set("Content-Type", "application/json")

    // Send request
    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Check response status
    if resp.StatusCode != http.StatusCreated {
        return errors.New("failed to create repository")
    }

    return nil
}
