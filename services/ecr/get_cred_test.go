package ecr

import (
    "context"
    "errors"
    "testing"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

// MockAWSCredentialsRetriever is a mock implementation of AWSCredentialsRetriever interface
type MockAWSCredentialsRetriever struct {
    mock.Mock
}

// Retrieve mocks the Retrieve method of AWSCredentialsRetriever interface
func (m *MockAWSCredentialsRetriever) Retrieve(ctx context.Context) (aws.Credentials, error) {
    args := m.Called(ctx)
    return args.Get(0).(aws.Credentials), args.Error(1)
}


func TestGetAWSCredentials(t *testing.T) {
	t.Run("Positive case", func(t *testing.T) {
		// Create a mock retriever for positive case
		mockRetriever := new(MockAWSCredentialsRetriever)
		creds := aws.Credentials{
			AccessKeyID:     "mock_access_key",
			SecretAccessKey: "mock_secret_key",
			SessionToken:    "mock_session_token",
		}
		mockRetriever.On("Retrieve", mock.Anything).Return(creds, nil)

		// Call the function under test with the mock retriever
		awsCreds, err := GetAWSCredentials(mockRetriever)

		// Assert the results
		assert.NoError(t, err)
		assert.Equal(t, creds.AccessKeyID, awsCreds.AccessKeyID)
		assert.Equal(t, creds.SecretAccessKey, awsCreds.SecretAccessKey)
		assert.Equal(t, creds.SessionToken, awsCreds.SessionToken)

		// Verify that the Retrieve method of the mock was called
		mockRetriever.AssertCalled(t, "Retrieve", mock.Anything)
	})

	t.Run("Negative case", func(t *testing.T) {
		// Create a mock retriever for negative case
		mockRetriever := new(MockAWSCredentialsRetriever)
		mockRetriever.On("Retrieve", mock.Anything).Return(aws.Credentials{}, errors.New("mock error"))

		// Call the function under test with the mock retriever
		_, err := GetAWSCredentials(mockRetriever)

		// Assert the error
		assert.Error(t, err)
		assert.Equal(t, "failed to retrieve AWS credentials", err.Error())

		// Verify that the Retrieve method of the mock was called
		mockRetriever.AssertCalled(t, "Retrieve", mock.Anything)
	})
}