package gitsetup

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/stretchr/testify/assert"
)

// Manual mock for ConfigLoader
type MockConfigLoader struct {
	Call int
	Cfg  aws.Config
	Err  error
}

func (m *MockConfigLoader) LoadDefaultConfig(ctx context.Context, opts ...func(*config.LoadOptions) error) (aws.Config, error) {
	m.Call++
	return m.Cfg, m.Err
}

// Manual mock for SecretsManagerClient
type MockSecretsManagerClient struct {
	Call   int
	Result *secretsmanager.GetSecretValueOutput
	Err    error
}

func (m *MockSecretsManagerClient) GetSecretValue(ctx context.Context, params *secretsmanager.GetSecretValueInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.GetSecretValueOutput, error) {
	m.Call++
	return m.Result, m.Err
}

func TestFetchSecretToken(t *testing.T) {
	// Create mocks
	mockConfigLoader := &MockConfigLoader{}
	mockSecretsManagerClient := &MockSecretsManagerClient{}

	// Positive test case
	mockConfigLoader.Cfg = aws.Config{Region: "us-east-1"} // Ensure this matches your AWS service region
	mockConfigLoader.Err = nil
	mockSecretsManagerClient.Result = &secretsmanager.GetSecretValueOutput{
		SecretString: aws.String(`{"GITHUB_TOKEN":"test_token"}`),
	}

	// Inject mocks
	configLoader = mockConfigLoader
	secretsManagerClient = mockSecretsManagerClient

	// Call the function under test
	token, err := FetchSecretToken()
	assert.NoError(t, err)
	assert.Equal(t, "test_token", token)

	// Negative test case - error fetching AWS config
	mockConfigLoader.Err = errors.New("AWS config error")
	_, err = FetchSecretToken()
	assert.Error(t, err)

	// Reset for the next test
	mockConfigLoader.Err = nil

	// Negative test case - error fetching secret value
	mockSecretsManagerClient.Err = errors.New("Secrets Manager error")
	_, err = FetchSecretToken()
	assert.Error(t, err)

	// Ensure LoadDefaultConfig was called once
	assert.Equal(t, 1, mockConfigLoader.Call)
}
