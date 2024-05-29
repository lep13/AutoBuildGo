package ecr

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func GetAWSConfig() (aws.Config, error) {
	cfg, err := globalAWSConfigLoader.LoadDefaultConfig(context.Background())
	if err != nil {
		return aws.Config{}, errors.New("failed to load AWS config")
	}
	return cfg, nil
}


// AWSConfigLoader defines the interface for loading AWS configurations
type AWSConfigLoader interface {
	LoadDefaultConfig(ctx context.Context, optFns ...func(*config.LoadOptions) error) (aws.Config, error)
}

// DefaultAWSConfigLoader is the default implementation of AWSConfigLoader
type DefaultAWSConfigLoader struct{}

// LoadDefaultConfig loads the default AWS configuration
func (d DefaultAWSConfigLoader) LoadDefaultConfig(ctx context.Context, optFns ...func(*config.LoadOptions) error) (aws.Config, error) {
	return config.LoadDefaultConfig(ctx, optFns...)
}

// globalAWSConfigLoader is a global instance of the AWS config loader
var globalAWSConfigLoader AWSConfigLoader = DefaultAWSConfigLoader{}

// GetAWSConfig retrieves AWS configuration using the global loader
