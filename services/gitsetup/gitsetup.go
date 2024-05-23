package gitsetup

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

type RepoConfig struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
	AutoInit    bool   `json:"auto_init"`
}

type SecretData struct {
	GITHUB_TOKEN string `json:"GITHUB_TOKEN"`
}

func fetchSecretToken(secretName string) (string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		log.Fatalf("Error creating AWS session: %v", err)
	}

	svc := secretsmanager.New(sess)
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		return "", err
	}

	var secretData SecretData
	if err := json.Unmarshal([]byte(*result.SecretString), &secretData); err != nil {
		return "", err
	}

	return secretData.GITHUB_TOKEN, nil
}

func CreateRepository(config RepoConfig) {
	token, err := fetchSecretToken("github_token")
	if err != nil {
		log.Fatalf("Failed to fetch secret: %v", err)
	}

	data, err := json.Marshal(map[string]interface{}{
		"name":        config.Name,
		"description": config.Description,
		"private":     config.Private,
	})
	if err != nil {
		log.Fatalf("Error marshaling config: %v", err)
	}

	url := "https://api.github.com/repos/lep13/ServiceTemplate/generate"
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Failed to read response body: %v", err)
		}
		log.Fatalf("Failed to create repository from template, status code: %d, response: %s", resp.StatusCode, string(body))
	}
}
