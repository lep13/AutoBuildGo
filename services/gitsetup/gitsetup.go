package gitsetup

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func CreateGitRepository(config RepoConfig) error {
	token, err := FetchSecretToken("github_token")
	if err != nil {
		return err
	}
	return createRepositoryWithTemplate(config, token)
}

func createRepositoryWithTemplate(config RepoConfig, token string) error {
	data, err := json.Marshal(map[string]interface{}{
		"name":        config.Name,
		"description": config.Description,
		"private":     config.Private,
	})
	if err != nil {
		return err
	}

	url := "https://api.github.com/repos/lep13/ServiceTemplate/generate"
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		log.Fatalf("Failed to create repository, status code: %d, response: %s", resp.StatusCode, string(body))
	}
	return nil
}
