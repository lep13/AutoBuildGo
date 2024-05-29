package ecr

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)
func NewAWSConfig(creds AWSCredentials) (aws.Config, error) {
	cfg, err := loadAWSConfig(context.TODO(),
		config.WithCredentialsProvider(
			NewCredentialsCache(creds),
		),
	)
	if err != nil {
		return aws.Config{}, err
	}
	return cfg, nil
}

// NewCredentialsCache creates a new credentials cache with provided credentials.
func NewCredentialsCache(creds AWSCredentials) aws.CredentialsProvider {
	return aws.NewCredentialsCache(
		aws.CredentialsProviderFunc(func(context.Context) (aws.Credentials, error) {
			return aws.Credentials{
				AccessKeyID:     creds.AccessKeyID,
				SecretAccessKey: creds.SecretAccessKey,
				SessionToken:    creds.SessionToken,
			}, nil
		}),
	)
}

// loadAWSConfig is a function variable that loads AWS configuration.
var loadAWSConfig = func(ctx context.Context, optFns ...func(*config.LoadOptions) error) (aws.Config, error) {
	return config.LoadDefaultConfig(ctx, optFns...)
}

// NewAWSConfig creates a new AWS configuration with provided credentials.
