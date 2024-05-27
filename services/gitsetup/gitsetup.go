package gitsetup

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

// CreateGitRepository creates a new GitHub repository using the specified configuration.
func CreateGitRepository(config RepoConfig) error {
    token, err := FetchSecretToken("github_token")
    if err != nil {
        return err
    }
    return createRepositoryWithTemplate(config, token)
}

// createRepositoryWithTemplate sends a request to GitHub API to create a repository from a template.
func createRepositoryWithTemplate(config RepoConfig, token string) error {
    data, err := json.Marshal(map[string]interface{}{
        "name":        config.Name,
        "description": config.Description,
        "private":     config.Private,
    })
    if err != nil {
        return err
    }

    client := &http.Client{}
    req, err := http.NewRequest(http.MethodPost, config.TemplateURL, bytes.NewBuffer(data))
    if err != nil {
        return err
    }

    req.Header.Set("Authorization", "token " + token)
    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusCreated {
        return nil
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return fmt.Errorf("failed to read response body: %w", err)
    }

    return fmt.Errorf("failed to create repository, status code: %d, response: %s", resp.StatusCode, string(body))
}
