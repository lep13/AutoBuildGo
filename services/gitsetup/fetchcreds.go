package gitsetup

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
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

var secretCache = struct {
	sync.Mutex
	data map[string]string
}{data: make(map[string]string)}

// FetchSecretToken retrieves the GitHub token from AWS Secrets Manager.
func FetchSecretToken() (string, error) {
	const secretName = "github_token"

	secretCache.Lock()
	if token, found := secretCache.data[secretName]; found {
		secretCache.Unlock()
		return token, nil
	}
	secretCache.Unlock()

	creds, err := GetAWSCredentials()
	if err != nil {
		return "", fmt.Errorf("error retrieving AWS credentials: %v", err)
	}

	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithCredentialsProvider(creds))
	if err != nil {
		return "", fmt.Errorf("error loading AWS config: %v", err)
	}

	client := secretsmanager.NewFromConfig(cfg)
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	result, err := client.GetSecretValue(context.Background(), input)
	if err != nil {
		return "", fmt.Errorf("error fetching secret value: %v", err)
	}

	var secretData SecretData
	err = json.Unmarshal([]byte(*result.SecretString), &secretData)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling secret value: %v", err)
	}

	secretCache.Lock()
	secretCache.data[secretName] = secretData.GITHUB_TOKEN
	secretCache.Unlock()

	return secretData.GITHUB_TOKEN, nil
}

type AWSCredentials struct {
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
}

// Retrieve returns the AWS credentials.
func (c AWSCredentials) Retrieve(ctx context.Context) (aws.Credentials, error) {
	return aws.Credentials{
		AccessKeyID:     c.AccessKeyID,
		SecretAccessKey: c.SecretAccessKey,
		SessionToken:    c.SessionToken,
	}, nil
}

// GetAWSCredentials retrieves AWS credentials
func GetAWSCredentials() (AWSCredentials, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return AWSCredentials{}, errors.New("failed to load AWS config")
	}

	creds, err := cfg.Credentials.Retrieve(context.Background())
	if err != nil {
		return AWSCredentials{}, errors.New("failed to retrieve AWS credentials")
	}

	return AWSCredentials{
		AccessKeyID:     creds.AccessKeyID,
		SecretAccessKey: creds.SecretAccessKey,
		SessionToken:    creds.SessionToken,
	}, nil
}
