package ecr

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/stretchr/testify/assert"
)

// MockAWSConfigLoader is a mock implementation of AWSConfigLoader for testing
type MockAWSConfigLoader struct{}

// LoadDefaultConfig mocks the LoadDefaultConfig method and returns the provided config and error
func (m MockAWSConfigLoader) LoadDefaultConfig(ctx context.Context, optFns ...func(*config.LoadOptions) error) (aws.Config, error) {
	// Mocked config for testing
	return aws.Config{Region: "us-west-2"}, nil
}

// MockAWSConfigLoaderError is a mock implementation of AWSConfigLoader that returns an error
type MockAWSConfigLoaderError struct{}

// LoadDefaultConfig mocks the LoadDefaultConfig method and returns an error
func (m MockAWSConfigLoaderError) LoadDefaultConfig(ctx context.Context, optFns ...func(*config.LoadOptions) error) (aws.Config, error) {
	// Mocked error for testing
	return aws.Config{}, errors.New("failed to load AWS config")
}

func TestGetAWSConfig(t *testing.T) {
	// Backup the original globalAWSConfigLoader and defer restoring it
	originalLoader := globalAWSConfigLoader
	defer func() { globalAWSConfigLoader = originalLoader }()

	// Positive test case
	t.Run("LoadDefaultConfig_Success", func(t *testing.T) {
		// Replace the globalAWSConfigLoader with the mock loader
		globalAWSConfigLoader = MockAWSConfigLoader{}

		// Call the function under test
		config, err := GetAWSConfig()

		// Assert the result
		assert.NoError(t, err)
		assert.Equal(t, "us-west-2", config.Region)
	})

	// Negative test case
	t.Run("LoadDefaultConfig_Failure", func(t *testing.T) {
		// Replace the globalAWSConfigLoader with a mock loader that returns an error
		globalAWSConfigLoader = MockAWSConfigLoaderError{}

		// Call the function under test
		config, err := GetAWSConfig()

		// Assert the error
		assert.Error(t, err)
		assert.EqualError(t, err, "failed to load AWS config")
		assert.Equal(t, aws.Config{}, config)
	})
}
