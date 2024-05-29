package ecr

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func TestNewAWSConfig(t *testing.T) {
	t.Run("Positive case - Successful configuration", func(t *testing.T) {
		mockLoadConfig := func(ctx context.Context, optFns ...func(*config.LoadOptions) error) (aws.Config, error) {
			return aws.Config{
				Region: "us-west-2", // Assuming Region is a required field for a valid config
			}, nil
		}

		loadAWSConfig = mockLoadConfig

		creds := AWSCredentials{
			AccessKeyID:     "test-access-key-id",
			SecretAccessKey: "test-secret-access-key",
			SessionToken:    "test-session-token",
		}

		// Call the NewAWSConfig function
		cfg, err := NewAWSConfig(creds)

		// Check if there is no error
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		// Check if the configuration is populated
		if cfg.Region == "" {
			t.Errorf("Expected populated configuration, got empty")
		}
	})

	t.Run("Negative case - Failed to load configuration", func(t *testing.T) {
		mockLoadConfig := func(ctx context.Context, optFns ...func(*config.LoadOptions) error) (aws.Config, error) {
			return aws.Config{}, errors.New("mock error")
		}

		loadAWSConfig = mockLoadConfig

		creds := AWSCredentials{
			AccessKeyID:     "test-access-key-id",
			SecretAccessKey: "test-secret-access-key",
			SessionToken:    "test-session-token",
		}

		// Call the NewAWSConfig function
		_, err := NewAWSConfig(creds)

		// Check if the error is as expected
		if err == nil || err.Error() != "mock error" {
			t.Errorf("Expected 'mock error', got '%v'", err)
		}
	})
}

func TestNewCredentialsCache(t *testing.T) {
	// Define mock credentials
	mockCreds := aws.Credentials{
		AccessKeyID:     "mock-access-key-id",
		SecretAccessKey: "mock-secret-access-key",
		SessionToken:    "mock-session-token",
	}

	// Convert mock credentials to type AWSCredentials
	mockAWSCreds := AWSCredentials{
		AccessKeyID:     mockCreds.AccessKeyID,
		SecretAccessKey: mockCreds.SecretAccessKey,
		SessionToken:    mockCreds.SessionToken,
	}

	// Call the NewCredentialsCache function
	provider := NewCredentialsCache(mockAWSCreds)

	// Retrieve credentials from the provider
	creds, err := provider.Retrieve(context.Background())

	// Check if there is no error
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the retrieved credentials match the mock credentials
	if creds.AccessKeyID != mockCreds.AccessKeyID || creds.SecretAccessKey != mockCreds.SecretAccessKey || creds.SessionToken != mockCreds.SessionToken {
		t.Errorf("Expected credentials: %v, got: %v", mockCreds, creds)
	}
}