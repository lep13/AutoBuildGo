package ecr

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
)

// mockAWSClient is a mock implementation of the AWSClient interface for testing purposes.
type mockAWSClient struct {
	CreateRepositoryFunc func(ctx context.Context, params *ecr.CreateRepositoryInput, optFns ...func(*ecr.Options)) (*ecr.CreateRepositoryOutput, error)
}

func (m *mockAWSClient) CreateRepository(ctx context.Context, params *ecr.CreateRepositoryInput, optFns ...func(*ecr.Options)) (*ecr.CreateRepositoryOutput, error) {
	return m.CreateRepositoryFunc(ctx, params, optFns...)
}

func TestCreateRepo(t *testing.T) {
	t.Run("Positive case - Successful creation", func(t *testing.T) {
		// Initialize a mock object with success simulation
		mockClient := &mockAWSClient{
			CreateRepositoryFunc: func(ctx context.Context, params *ecr.CreateRepositoryInput, optFns ...func(*ecr.Options)) (*ecr.CreateRepositoryOutput, error) {
				return &ecr.CreateRepositoryOutput{}, nil
			},
		}

		// Call the CreateRepo function with the mock client
		err := CreateRepo("test-repo", mockClient)

		// Check if there is no error
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	})

	t.Run("Negative case - Repository already exists", func(t *testing.T) {
		// Initialize a mock object with repository already exists simulation
		mockClient := &mockAWSClient{
			CreateRepositoryFunc: func(ctx context.Context, params *ecr.CreateRepositoryInput, optFns ...func(*ecr.Options)) (*ecr.CreateRepositoryOutput, error) {
				return nil, &types.RepositoryAlreadyExistsException{}
			},
		}

		// Call the CreateRepo function with the mock client
		err := CreateRepo("existing-repo", mockClient)

		// Check if there is no error (since it should handle already exists case gracefully)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	})

	t.Run("Negative case - Error creating repository", func(t *testing.T) {
		// Initialize a mock object with failure simulation
		mockClient := &mockAWSClient{
			CreateRepositoryFunc: func(ctx context.Context, params *ecr.CreateRepositoryInput, optFns ...func(*ecr.Options)) (*ecr.CreateRepositoryOutput, error) {
				return nil, errors.New("create repository failed")
			},
		}

		// Call the CreateRepo function with the mock client
		err := CreateRepo("error-repo", mockClient)

		// Check if the error is as expected
		if err == nil || err.Error() != "create repository failed" {
			t.Errorf("Expected 'create repository failed' error, got '%v'", err)
		}
	})
}
