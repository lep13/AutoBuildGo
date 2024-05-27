package ecr

import (
	"context"
	"log"
	"github.com/aws/aws-sdk-go-v2/config"
)

type AWSCredentials struct {
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
}

func GetAWSCredentials() (AWSCredentials, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Printf("Failed to load AWS SDK config: %v", err)
		return AWSCredentials{}, err
	}

	creds, err := cfg.Credentials.Retrieve(context.Background())
	if err != nil {
		log.Printf("Failed to retrieve AWS credentials: %v", err)
		return AWSCredentials{}, err
	}

	return AWSCredentials{
		AccessKeyID:     creds.AccessKeyID,
		SecretAccessKey: creds.SecretAccessKey,
		SessionToken:    creds.SessionToken,
	}, nil
}
