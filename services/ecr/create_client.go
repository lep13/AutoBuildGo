package ecr

import (
	// "context"
	// "log"

	// "github.com/aws/aws-sdk-go-v2/aws"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
)

// CreateECRClient creates and returns an ECR client using the provided AWS credentials.
var getAWSConfigFunc = GetAWSConfig

// CreateECRClient creates and returns an ECR client using the provided AWS credentials.
func CreateECRClient() (*ecr.Client, error) {
    cfg, err := getAWSConfigFunc()
    if err != nil {
        return nil, err
    }
    return ecr.NewFromConfig(cfg), nil
}
func MockGetAWSConfig() (aws.Config, error) {
    // Mocked implementation for testing
    return aws.Config{}, errors.New("mocked error")
}