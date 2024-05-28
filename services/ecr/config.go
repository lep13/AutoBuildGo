package ecr

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)


// NewAWSConfig creates a new AWS configuration with provided credentials.
func NewAWSConfig(creds AWSCredentials) (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithCredentialsProvider(
			aws.NewCredentialsCache(
				aws.CredentialsProviderFunc(func(context.Context) (aws.Credentials, error) {
					return aws.Credentials{
						AccessKeyID:     creds.AccessKeyID,
						SecretAccessKey: creds.SecretAccessKey,
						SessionToken:    creds.SessionToken,
					}, nil
				}),
			),
		),
	)
	if err != nil {
		return aws.Config{}, err
	}
	return cfg, nil
}
