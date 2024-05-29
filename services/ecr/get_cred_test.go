package ecr

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

// mockAWSCredentialsRetriever is a mock implementation of the AWSCredentialsRetriever interface.
type mockAWSCredentialsRetriever struct {
	RetrieveFunc func(ctx context.Context) (aws.Credentials, error)
}

func (m *mockAWSCredentialsRetriever) Retrieve(ctx context.Context) (aws.Credentials, error) {
	return m.RetrieveFunc(ctx)
}

// mockLoadDefaultConfig is a mock implementation of the loadDefaultConfig function.
type mockLoadDefaultConfig struct {
	LoadDefaultConfigFunc func(ctx context.Context, optFns ...func(*config.LoadOptions) error) (aws.Config, error)
}

func (m *mockLoadDefaultConfig) LoadDefaultConfig(ctx context.Context, optFns ...func(*config.LoadOptions) error) (aws.Config, error) {
	return m.LoadDefaultConfigFunc(ctx, optFns...)
}

func TestGetAWSCredentials(t *testing.T) {
	originalLoadDefaultConfig := loadDefaultConfig
	defer func() { loadDefaultConfig = originalLoadDefaultConfig }()

	t.Run("Positive case - Successful retrieval", func(t *testing.T) {
		// Initialize a mock object with success simulation
		mockRetriever := &mockAWSCredentialsRetriever{
			RetrieveFunc: func(ctx context.Context) (aws.Credentials, error) {
				return aws.Credentials{
					AccessKeyID:     "test-access-key-id",
					SecretAccessKey: "test-secret-access-key",
					SessionToken:    "test-session-token",
				}, nil
			},
		}

		mockLoadConfig := &mockLoadDefaultConfig{
			LoadDefaultConfigFunc: func(ctx context.Context, optFns ...func(*config.LoadOptions) error) (aws.Config, error) {
				return aws.Config{Credentials: aws.NewCredentialsCache(mockRetriever)}, nil
			},
		}

		loadDefaultConfig = mockLoadConfig.LoadDefaultConfig

		// Call the GetAWSCredentials function
		creds, err := GetAWSCredentials()

		// Check if there is no error
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		// Check if the credentials are as expected
		if creds.AccessKeyID != "test-access-key-id" || creds.SecretAccessKey != "test-secret-access-key" || creds.SessionToken != "test-session-token" {
			t.Errorf("Unexpected credentials: %v", creds)
		}
	})

	t.Run("Negative case - Failed to retrieve credentials", func(t *testing.T) {
		// Initialize a mock object with failure simulation
		mockRetriever := &mockAWSCredentialsRetriever{
			RetrieveFunc: func(ctx context.Context) (aws.Credentials, error) {
				return aws.Credentials{}, errors.New("mock error")
			},
		}

		mockLoadConfig := &mockLoadDefaultConfig{
			LoadDefaultConfigFunc: func(ctx context.Context, optFns ...func(*config.LoadOptions) error) (aws.Config, error) {
				return aws.Config{Credentials: aws.NewCredentialsCache(mockRetriever)}, nil
			},
		}

		loadDefaultConfig = mockLoadConfig.LoadDefaultConfig

		// Call the GetAWSCredentials function
		_, err := GetAWSCredentials()

		// Check if the error is as expected
		if err == nil || err.Error() != "failed to retrieve AWS credentials" {
			t.Errorf("Expected 'failed to retrieve AWS credentials' error, got '%v'", err)
		}
	})
}
