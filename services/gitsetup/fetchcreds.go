package gitsetup

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type ConfigLoader interface {
	LoadDefaultConfig(ctx context.Context, options ...func(*config.LoadOptions) error) (aws.Config, error)
}

var configLoader ConfigLoader = &defaultConfigLoader{}

type defaultConfigLoader struct{}

func (l *defaultConfigLoader) LoadDefaultConfig(ctx context.Context, options ...func(*config.LoadOptions) error) (aws.Config, error) {
	return config.LoadDefaultConfig(ctx, options...)
}

type SecretsManagerClient interface {
	GetSecretValue(ctx context.Context, params *secretsmanager.GetSecretValueInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.GetSecretValueOutput, error)
}

var secretsManagerClient SecretsManagerClient

func init() {
	cfg, err := configLoader.LoadDefaultConfig(context.Background(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	secretsManagerClient = secretsmanager.NewFromConfig(cfg)
}

type CommandRunner interface {
	Run(cmd *exec.Cmd) error
	Output(cmd *exec.Cmd) ([]byte, error)
}

type DefaultCommandRunner struct{}

func (r *DefaultCommandRunner) Run(cmd *exec.Cmd) error {
	return cmd.Run()
}

func (r *DefaultCommandRunner) Output(cmd *exec.Cmd) ([]byte, error) {
	return cmd.Output()
}

var runner CommandRunner = &DefaultCommandRunner{}

var secretCache = struct {
	sync.Mutex
	data map[string]string
}{data: make(map[string]string)}

func FetchSecretToken() (string, error) {
	const secretName = "github_token"

	secretCache.Lock()
	if token, found := secretCache.data[secretName]; found {
		secretCache.Unlock()
		return token, nil
	}
	secretCache.Unlock()

	_, err := configLoader.LoadDefaultConfig(context.Background())
	if err != nil {
		return "", fmt.Errorf("error loading AWS config: %v", err)
	}

	client := secretsManagerClient
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
