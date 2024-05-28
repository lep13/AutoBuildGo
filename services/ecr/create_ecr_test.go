package ecr

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAWSClient is a mock implementation of AWSClient interface for testing purposes.
type MockAWSClient struct {
	mock.Mock
}

// CreateRepository mocks the CreateRepository method.
func (m *MockAWSClient) CreateRepository(ctx context.Context, params *ecr.CreateRepositoryInput, optFns ...func(*ecr.Options)) (*ecr.CreateRepositoryOutput, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(*ecr.CreateRepositoryOutput), args.Error(1)
}

func TestCreateRepo(t *testing.T) {
	t.Run("CreateRepository_Success", func(t *testing.T) {
		mockClient := new(MockAWSClient)
		repoName := "test-repo"
		input := &ecr.CreateRepositoryInput{
			RepositoryName:     aws.String(repoName),
			ImageTagMutability: types.ImageTagMutabilityImmutable,
			ImageScanningConfiguration: &types.ImageScanningConfiguration{
				ScanOnPush: true,
			},
		}
		mockClient.On("CreateRepository", context.Background(), input).Return(&ecr.CreateRepositoryOutput{}, nil)

		err := CreateRepo(repoName, mockClient)

		assert.NoError(t, err, "Expected no error for creating a new repository")
		mockClient.AssertExpectations(t)
	})

	t.Run("CreateRepository_AlreadyExists", func(t *testing.T) {
		mockClient := new(MockAWSClient)
		repoName := "test-repo"
		input := &ecr.CreateRepositoryInput{
			RepositoryName:     aws.String(repoName),
			ImageTagMutability: types.ImageTagMutabilityImmutable,
			ImageScanningConfiguration: &types.ImageScanningConfiguration{
				ScanOnPush: true,
			},
		}
	
		// Configure mockClient to return a non-nil output and nil error for RepositoryAlreadyExistsException
		mockClient.On("CreateRepository", context.Background(), input).Return(&ecr.CreateRepositoryOutput{}, nil)
	
		err := CreateRepo(repoName, mockClient)
	
		assert.NoError(t, err, "Expected no error for existing repository")
		mockClient.AssertExpectations(t)
	})


	t.Run("CreateRepository_Failure", func(t *testing.T) {
		mockClient := new(MockAWSClient)
		repoName := "test-repo"
		input := &ecr.CreateRepositoryInput{
			RepositoryName:     aws.String(repoName),
			ImageTagMutability: types.ImageTagMutabilityImmutable,
			ImageScanningConfiguration: &types.ImageScanningConfiguration{
				ScanOnPush: true,
			},
		}
		expectedErr := errors.New("some error")
		
		// Configure mockClient to return nil output and a non-nil error
		mockClient.On("CreateRepository", context.Background(), input).Return(&ecr.CreateRepositoryOutput{}, expectedErr)
	
		err := CreateRepo(repoName, mockClient)
	
		assert.Error(t, err, "Expected an error for unknown repository creation error")
		assert.Equal(t, expectedErr, err, "Expected the same error returned by AWS client")
		mockClient.AssertExpectations(t)
	})
	
	
}
