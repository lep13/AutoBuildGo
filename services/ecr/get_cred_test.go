package ecr

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	// "github.com/aws/aws-sdk-go-v2/config"
)


type MockAWSCredentialsRetriever struct {
    SimulateSuccess bool
    
}

func (m *MockAWSCredentialsRetriever) Retrieve(ctx context.Context) (aws.Credentials, error) {
    if m.SimulateSuccess {
        // Simulate a successful retrieval by returning dummy credentials
        return aws.Credentials{
            AccessKeyID:     "access_key_id",
            SecretAccessKey: "secret_access_key",
            SessionToken:    "session_token",
        }, nil
    }
    // Simulate an error by returning a mock error
    return aws.Credentials{}, errors.New("mock error")
}
func TestRetrieveMethod(t *testing.T) {
    t.Run("Positive case", func(t *testing.T) {
        // Initialize a mock object with success simulation
        mockRetriever := &MockAWSCredentialsRetriever{
            SimulateSuccess: true,
        }

        // Call the Retrieve method with the mock object
        _, err := mockRetriever.Retrieve(context.Background())

        // Check if there is no error
        if err != nil {
            t.Errorf("Unexpected error: %v", err)
        }
    })

    t.Run("Negative case", func(t *testing.T) {
        // Initialize a mock object with failure simulation
        mockRetriever := &MockAWSCredentialsRetriever{
            SimulateSuccess: false,
        }
  
        // Call the Retrieve method with the mock object
        _, err := mockRetriever.Retrieve(context.Background())
       
        // Check if the error is as expected
        if err == nil || err.Error() != "mock error" {
            t.Errorf("Expected 'mock error', got '%v'", err)
        }
    })
}