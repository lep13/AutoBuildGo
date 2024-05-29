package ecr

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/stretchr/testify/assert"
)

func TestCreateECRClient(t *testing.T) {
    tests := []struct {
        name          string
        mockFunc      func() (aws.Config, error)
        expectError   bool
    }{
        {
            name:          "PositiveCase",
            mockFunc:      MockGetAWSConfig,
            expectError:   true,
        },
        {
            name:          "NegativeCase",
            mockFunc:      GetAWSConfig,
            expectError:   false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Save the original function
            originalGetAWSConfigFunc := getAWSConfigFunc
            // Replace the original function with the mock function
            getAWSConfigFunc = tt.mockFunc
            // Restore the original function after the test
            defer func() {
                getAWSConfigFunc = originalGetAWSConfigFunc
            }()

            // Call the CreateECRClient function
            _, err := CreateECRClient()

            // Check if we expect an error or not
            if tt.expectError {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
            }
        })
    }
}
