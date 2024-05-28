package ecr

import (
	"context"
	"errors"
	// "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Mock your config loader implementation
		mockLoader := mockConfigLoader{
			ShouldSucceed: true,
			Error: assert.AnError,
		}

		// Call your LoadConfig function
		_, err := LoadConfig(mockLoader)

		// Check if the error is nil
		assert.Nil(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		// Mock your config loader implementation to return an error
		mockLoader := mockConfigLoader{
			Error: assert.AnError,
			ShouldSucceed: false,
		}

		// Call your LoadConfig function
		_, err := LoadConfig(mockLoader)

		// Check if the error is not nil
		assert.Error(t, err)
		// Check if the error message matches the expected error message
		assert.EqualError(t, err, "error loading config")
	})
}

// Mock implementation of ConfigLoader interface for testing
type mockConfigLoader struct {
	Error error
	ShouldSucceed bool // Flag to control whether to simulate success or failure
}

// Mock implementation of LoadDefaultConfig method for testing
func (m mockConfigLoader) LoadDefaultConfig(ctx context.Context) (interface{}, error) {
	if m.ShouldSucceed {
		// Mock your desired behavior here for success
		return struct{}{}, nil
	}
	// Mock your desired behavior here for failure
	return nil, errors.New("error loading config")
}
