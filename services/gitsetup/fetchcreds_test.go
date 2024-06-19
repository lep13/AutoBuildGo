package gitsetup

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockConfigLoader is a mock for ConfigLoader
type MockConfigLoader struct {
	mock.Mock
}

func (m *MockConfigLoader) LoadDefaultConfig(ctx context.Context, options ...func(*config.LoadOptions) error) (aws.Config, error) {
	args := m.Called(ctx, options)
	return args.Get(0).(aws.Config), args.Error(1)
}

// MockSecretsManagerClient is a mock for SecretsManagerClient
type MockSecretsManagerClient struct {
	mock.Mock
}

func (m *MockSecretsManagerClient) GetSecretValue(ctx context.Context, params *secretsmanager.GetSecretValueInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.GetSecretValueOutput, error) {
	args := m.Called(ctx, params, optFns)
	return args.Get(0).(*secretsmanager.GetSecretValueOutput), args.Error(1)
}

func TestFetchSecretToken(t *testing.T) {
	// Backup the real instances
	realConfigLoader := configLoader
	realSecretsManagerClient := secretsManagerClient

	// Restore the real instances after the test completes
	defer func() {
		configLoader = realConfigLoader
		secretsManagerClient = realSecretsManagerClient
	}()

	tests := []struct {
		name          string
		mockConfig    func(*MockConfigLoader)
		mockSecret    func(*MockSecretsManagerClient)
		expectedToken string
		expectedError string
	}{
		{
			name: "success",
			mockConfig: func(m *MockConfigLoader) {
				m.On("LoadDefaultConfig", mock.Anything, mock.Anything).Return(aws.Config{Region: "us-east-1"}, nil).Once()
			},
			mockSecret: func(m *MockSecretsManagerClient) {
				secretString := `{"GITHUB_TOKEN": "test_token"}`
				m.On("GetSecretValue", mock.Anything, mock.Anything, mock.Anything).Return(&secretsmanager.GetSecretValueOutput{
					SecretString: &secretString,
				}, nil).Once()
			},
			expectedToken: "test_token",
		},
		{
			name: "config error",
			mockConfig: func(m *MockConfigLoader) {
				m.On("LoadDefaultConfig", mock.Anything, mock.Anything).Return(aws.Config{}, errors.New("config error")).Once()
			},
			expectedError: "error loading AWS config: config error",
		},
		{
			name: "secret error",
			mockConfig: func(m *MockConfigLoader) {
				m.On("LoadDefaultConfig", mock.Anything, mock.Anything).Return(aws.Config{Region: "us-east-1"}, nil).Once()
			},
			mockSecret: func(m *MockSecretsManagerClient) {
				m.On("GetSecretValue", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("secret error")).Once()
			},
			expectedError: "error fetching secret value: secret error",
		},
		{
			name: "unmarshal error",
			mockConfig: func(m *MockConfigLoader) {
				m.On("LoadDefaultConfig", mock.Anything, mock.Anything).Return(aws.Config{Region: "us-east-1"}, nil).Once()
			},
			mockSecret: func(m *MockSecretsManagerClient) {
				invalidSecretString := `{"INVALID_JSON"}`
				m.On("GetSecretValue", mock.Anything, mock.Anything, mock.Anything).Return(&secretsmanager.GetSecretValueOutput{
					SecretString: &invalidSecretString,
				}, nil).Once()
			},
			expectedError: "error unmarshalling secret value: invalid character 'I' looking for beginning of object key string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create new mocks for each test
			mockConfigLoader := new(MockConfigLoader)
			mockSecretsManagerClient := new(MockSecretsManagerClient)

			// Assign the new mocks to the global variables
			configLoader = mockConfigLoader
			secretsManagerClient = mockSecretsManagerClient

			if tt.mockConfig != nil {
				tt.mockConfig(mockConfigLoader)
			}
			if tt.mockSecret != nil {
				tt.mockSecret(mockSecretsManagerClient)
			}

			token, err := FetchSecretToken()

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedToken, token)
			}

			mockConfigLoader.AssertExpectations(t)
			mockSecretsManagerClient.AssertExpectations(t)
		})
	}
}
